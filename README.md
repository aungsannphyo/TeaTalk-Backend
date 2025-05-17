
```
ywar-talk-app/
│
├── cmd/
│   └── main.go                      # App entry point
│
├── internal/
│   ├── domain/
│   │   ├── model/                   # Entity structs: User, Message, Group
│   │   └── repository/             # Interfaces for repositories
│   │
│   ├── repository_impl/                 # DB implementation using MariaDB (via sql.DB)
│   │   ├── user_repository.go
│   │   ├── message_repository.go
│   │   └── group_repository.go
│   │
│   ├── service/                    # Business logic (auth, chat, etc.)
│   │   ├── auth_service.go
│   │   ├── chat_service.go
│   │   └── post_service.go
│   │
│   ├── handler/                    # HTTP handlers
│   │   ├── auth_handler.go
│   │   ├── message_handler.go
│   │   └── group_handler.go
│   │
│   ├── middleware/
│   │   └── auth_middleware.go      # JWT Auth
│
├── pkg/
│   ├── config/                     # DB and environment config
│   ├── utils/                      # JWT, password hashing
│   └── database/
│       └── mysql.go                # Open & manage DB connection
│
├── go.mod
└── go.sum
```


dispatcher.GroupObserverNotify(events.GroupEvent{
	GroupID: conversationID,
	Type:    events.GroupUserAdded,  // or GroupUserRemoved
})


package websocket

import (
"net/http"

```
s "github.com/aungsannphyo/ywartalk/internal/domain/service"
"github.com/gin-gonic/gin"
"github.com/gorilla/websocket"
```

)

var upgrader = websocket.Upgrader{
CheckOrigin: func(r \*http.Request) bool { return true },
}

type WebSocketHandler struct {
hub        \*Hub
msgService s.MessageService
uService   s.UserService
}

func NewWebSocketHandler(
hub \*Hub,
msgS s.MessageService,
userS s.UserService,
) \*WebSocketHandler {
return \&WebSocketHandler{
hub:        hub,
msgService: msgS,
uService:   userS,
}
}

func (h \*WebSocketHandler) WebSocketHandler(c \*gin.Context) {
userID := c.GetString("userId")
conversationID := c.Query("group\_id")

```
conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
if err != nil {
	return
}
client := &Client{
	Hub:            h.hub,
	Conn:           conn,
	Send:           make(chan []byte, 256),
	UserID:         userID,
	MessageService: h.msgService,
	UserService:    h.uService,
}

SyncGroupUsersToHub(c.Request.Context(), h.hub, conversationID, userID, client, h.uService.GetGroupUsers)

client.Hub.register <- client

go client.WritePump()
go client.ReadPump()
```

}

package websocket

import (
"context"
"log"
"sync"
"time"

```
"github.com/aungsannphyo/ywartalk/internal/domain/service"
"github.com/aungsannphyo/ywartalk/pkg/events"
```

)

type GroupObserver interface {
OnGroupChanged(ctx context.Context, conversationID string)
}

type Hub struct {
clients        map\[string]\*Client
groups         map\[string]map\[string]\*Client
register       chan \*Client
unregister     chan \*Client
privateMessage chan PrivateMessage
groupMessage   chan GroupMessage
mu             sync.RWMutex
// debounce control per group
debounceMu     sync.Mutex
debounceTimers map\[string]\*time.Timer
userService    service.UserService
}

func NewHub(userService service.UserService) \*Hub {
return \&Hub{
clients:        make(map\[string]\*Client),
groups:         make(map\[string]map\[string]\*Client),
register:       make(chan \*Client),
unregister:     make(chan \*Client),
privateMessage: make(chan PrivateMessage),
groupMessage:   make(chan GroupMessage),
userService:    userService,
}
}

func (h \*Hub) Run() {
for {
select {
case c := <-h.register:
h.mu.Lock()
h.clients\[c.UserID] = c
h.mu.Unlock()
log.Printf("User %s connected", c.UserID)

```
	case c := <-h.unregister:
		h.mu.Lock()
		if _, ok := h.clients[c.UserID]; ok {
			delete(h.clients, c.UserID)
			close(c.Send)
			for groupID := range h.groups {
				h.RemoveUserFromGroup(groupID, c.UserID)
			}
			log.Printf("User %s disconnected", c.UserID)
		}
		h.mu.Unlock()

	case pm := <-h.privateMessage:
		h.mu.RLock()
		recipient, ok := h.clients[pm.ToUserID]
		h.mu.RUnlock()
		if ok {
			recipient.Send <- pm.Content
		}

	case gm := <-h.groupMessage:
		h.mu.RLock()
		members, ok := h.groups[gm.GroupID]
		h.mu.RUnlock()
		if ok {
			for uid, member := range members {
				if uid != gm.FromUserID {
					member.Send <- gm.Content
				}
			}
		}
	}
}
```

}

func (h \*Hub) AddUserToGroup(groupID, userID string, c \*Client) {
h.mu.Lock()
defer h.mu.Unlock()
if h.groups\[groupID] == nil {
h.groups\[groupID] = make(map\[string]\*Client)
}
h.groups\[groupID]\[userID] = c
}

func (h \*Hub) RemoveUserFromGroup(groupID, userID string) {
h.mu.Lock()
defer h.mu.Unlock()
if members, ok := h.groups\[groupID]; ok {
delete(members, userID)
if len(members) == 0 {
delete(h.groups, groupID)
}
}
}

func (h \*Hub) ReplaceGroupMembers(groupID string, clients map\[string]\*Client) {
h.mu.Lock()
defer h.mu.Unlock()
h.groups\[groupID] = clients
}

func (h \*Hub) OnGroupChanged(event events.GroupEvent) {
h.debounceMu.Lock()
defer h.debounceMu.Unlock()

```
if timer, exists := h.debounceTimers[event.GroupID]; exists {
	timer.Stop()
}
h.debounceTimers[event.GroupID] = time.AfterFunc(1*time.Second, func() {
	h.syncGroupUsers(event.GroupID)
})
```

}

func (h \*Hub) syncGroupUsers(groupID string) {
ctx := context.Background()

```
users, err := h.userService.GetGroupUsers(ctx, groupID)
if err != nil {
	log.Printf("[Hub] Failed to fetch users for group %s: %v", groupID, err)
	return
}

newMembers := make(map[string]*Client)
h.mu.RLock()
for _, user := range users {
	if client, ok := h.clients[user.ID]; ok {
		newMembers[user.ID] = client
	}
}
h.mu.RUnlock()

h.ReplaceGroupMembers(groupID, newMembers)

log.Printf("[Hub] Synced group %s users, total %d members", groupID, len(newMembers))
```

}
package websocket

import (
"context"
"log"

```
"github.com/aungsannphyo/ywartalk/internal/domain/models"
```

)

func SyncGroupUsersToHub(
ctx context.Context,
hub \*Hub,
conversationID string,
currentUserID string,
client \*Client,
getGroupUsers func(ctx context.Context, conversationID string) (\[]models.User, error),
) {
// Add the current user to the group
hub.AddUserToGroup(conversationID, currentUserID, client)

```
// Get all users in the group
users, err := getGroupUsers(ctx, conversationID)

if err != nil {
	log.Printf("Error fetching group users for conversation %s: %v", conversationID, err)
	return
}

hub.mu.RLock()
defer hub.mu.RUnlock()

for _, user := range users {
	if user.ID == currentUserID {
		continue
	}

	if connectedClient, ok := hub.clients[user.ID]; ok {
		hub.AddUserToGroup(conversationID, user.ID, connectedClient)
	}
}

log.Printf("[Hub] Synced group %s users, total %d members", conversationID, len(users))
```

}
package events

type GroupEventType string

const (
GroupUserAdded   GroupEventType = "GROUP\_USER\_ADDED"
GroupUserRemoved GroupEventType = "GROUP\_USER\_REMOVED"
)

type GroupEvent struct {
GroupID string
Type    GroupEventType
}

type GroupObserver interface {
OnGroupChanged(event GroupEvent)
}

type GroupEventDispatcher struct {
observers \[]GroupObserver
}

func NewGroupEventDispatcher() \*GroupEventDispatcher {
return \&GroupEventDispatcher{
observers: \[]GroupObserver{},
}
}

func (d \*GroupEventDispatcher) GroupObserverRegister(observer GroupObserver) {
d.observers = append(d.observers, observer)
}

func (d \*GroupEventDispatcher) GroupObserverNotify(event GroupEvent) {
for \_, o := range d.observers {
go o.OnGroupChanged(event)
}
}
dispatcher := events.NewGroupEventDispatcher()
hub := websocket.NewHub(serviceFactory.UserService())
dispatcher.GroupObserverRegister(hub)
go hub.Run()

```
websocketHandler := ws.NewWebSocketHandler(
	hub,
	serviceFactory.MessageService(),
	serviceFactory.UserService())
```

my goal is fetch for the first time for group user and then during the websocket active
user leave or new user coming i triger the observer and update with debounce

here is my problem
image i have 3 user batman superman and flash
batman is first connect is fine
and then superman is try to connect that not connect
why? can you check and analyze my code why

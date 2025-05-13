
go-chat-app/
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

package common

type ForbiddenError struct {
	Message string
}

func (e *ForbiddenError) Error() string {
	return e.Message
}

type BadRequestError struct {
	Message string
}

func (e *BadRequestError) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	return e.Message
}

type UnAuthorizedError struct {
	Message string
}

func (e *UnAuthorizedError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type ConflictError struct {
	Message string
}

func (e *ConflictError) Error() string {
	return e.Message
}

type UnexpectedRowCountError struct {
	Message string
}

func (e *UnexpectedRowCountError) Error() string {
	return e.Message
}

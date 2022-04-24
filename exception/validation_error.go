package exception

type ValidationError struct {
	Message string
}

type AuthorizedError struct {
	Status  string
	Message string
}

type ConflictError struct {
	Status  string
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}

func (authorizedError AuthorizedError) Error() string {
	return authorizedError.Message
}

func (conflictError ConflictError) Error() string {
	return conflictError.Message
}

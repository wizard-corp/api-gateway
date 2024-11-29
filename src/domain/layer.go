package domain

// Error interface implementation for all custom error types

type PresentationError struct {
	Title   string
	Message string
}

type ApplicationError struct {
	Title   string
	Message string
}

type InfrastructureError struct {
	Title   string
	Message string
}

type DomainError struct {
	Title   string
	Message string
}

// Implement the Error() method for each error type
func (e *PresentationError) Error() string {
	return e.Title + ": " + e.Message
}

func (e *ApplicationError) Error() string {
	return e.Title + ": " + e.Message
}

func (e *InfrastructureError) Error() string {
	return e.Title + ": " + e.Message
}

func (e *DomainError) Error() string {
	return e.Title + ": " + e.Message
}

// Constructor functions for creating new errors

func NewPresentationError(title string, message string) *PresentationError {
	return &PresentationError{Title: title, Message: message}
}

func NewApplicationError(title string, message string) *ApplicationError {
	return &ApplicationError{Title: title, Message: message}
}

func NewInfrastructureError(title string, message string) *InfrastructureError {
	return &InfrastructureError{Title: title, Message: message}
}

func NewDomainError(title string, message string) *DomainError {
	return &DomainError{Title: title, Message: message}
}

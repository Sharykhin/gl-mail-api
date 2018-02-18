package contract

// InputValidation - an interface for all request structs
type InputValidation interface {
	Validate() error
}

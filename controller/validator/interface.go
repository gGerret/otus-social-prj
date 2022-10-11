package validator

type FieldError struct {
	Field string `json:"field"`
	Err   string `json:"error"`
}

type Validator interface {
	Validate() []FieldError
}

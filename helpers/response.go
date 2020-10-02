package helpers

var (
	// EmptyValue define an empty value for error message
	EmptyValue = make([]int, 0)
)

// Response defines the structure of the response entity.
type Response struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

// Status defines the structure of the status entity for the response.
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SetResponse set the structure of the JSON response.
func SetResponse(code int, message string, data interface{}) *Response {
	res := &Response{
		Data: data,
		Status: Status{
			Code:    code,
			Message: message,
		},
	}

	return res
}

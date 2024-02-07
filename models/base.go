package models

type Request [D interface{}] struct {
	Data D `json:"data"` 
}

type Response struct {
	Data     	 interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Message	   interface{} `json:"message,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

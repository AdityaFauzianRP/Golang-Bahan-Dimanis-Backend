package model

type APIRoute struct {
	ID           int    `json:"id"`
	Path         string `json:"path"`
	Method       string `json:"method"`
	FunctionName string `json:"function_name"`
	Middleware   bool   `json:"middleware"`
}

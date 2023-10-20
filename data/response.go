package response

type ResponseResult struct {
	Code   int         `json:"Code"`
	Status string      `json:"Status"`
	Data   interface{} `json:"data,omitempty"`
}

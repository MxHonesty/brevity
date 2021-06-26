package response

// A Response type is sent by the server to the client. The data field contains
// the data that the operation returns. This data will be cast to the appropriate
// type. This can be done because every request will know it's response return
// type.
type Response struct {
	Data interface{}
}

// Create a new Response.
func NewResponse(data interface{}) Response {
	return Response{Data: data}
}

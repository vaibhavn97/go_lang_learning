package main

import "fmt"

type HttpResponse int

const (
	OK          HttpResponse = 202
	BAD_REQUEST              = 400
	NOT_FOUND                = 404
	FORBIDDEN                = 401
)


func (code HttpResponse) String() string{
	return "Implemented";
}


func main() {
	fmt.Println(getResponse(400))
}

func getResponse(code HttpResponse) HttpResponse {
	switch code {
	case OK:
		return OK
	case BAD_REQUEST:
		return BAD_REQUEST

	default:
		panic(fmt.Errorf("Unundetified Error"))
	}
}

package httpserver

type HttpServer struct {
	userService UserService
}

func NewHttpServer(userService UserService) HttpServer {
	return HttpServer{
		userService: userService,
	}
}

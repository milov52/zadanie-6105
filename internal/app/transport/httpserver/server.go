package httpserver

type HttpServer struct {
	userService   UserService
	tenderService TenderService
	bidService    BidService
}

func NewHttpServer(userService UserService, tenderService TenderService, bidService BidService) HttpServer {
	return HttpServer{
		userService:   userService,
		tenderService: tenderService,
		bidService:    bidService,
	}
}

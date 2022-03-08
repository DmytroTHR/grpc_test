package service

import (
	"context"
	"log"
	"strings"
	"svcpod-server/proto"
)

type MyService struct {
	*proto.UnimplementedSvcPodServServer
}

func NewMyService() *MyService {
	return &MyService{}
}

func (svc *MyService) AskForHelp(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	responseText := workWithRequest(request.Message)
	log.Println(request.Message, "=>", responseText)
	return &proto.Response{Message: responseText}, nil
}

func workWithRequest(req string) string {
	if strings.Contains(strings.ToLower(req), "русск") ||
		strings.Contains(strings.ToLower(req), "russi") {
		return "GO FUCK YOURSELF! (ИДИ НА Х*Й!)"
	}

	return "You're welcome!"
}

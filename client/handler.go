package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"svcpod-client/proto"
)

type requestHandler struct {
	ClientGRPC proto.SvcPodServClient
}

func NewRequestHandler(client proto.SvcPodServClient) *requestHandler {
	return &requestHandler{ClientGRPC: client}
}

func (rh *requestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		decodedQuery, err := url.QueryUnescape(r.URL.RawQuery)
		if err != nil {
			fmt.Fprintln(w, "Error getting response:", err)
			return
		}
		response, err := rh.ClientGRPC.AskForHelp(context.TODO(), &proto.Request{Message: decodedQuery})
		if err != nil {
			fmt.Fprintln(w, "Error getting response:", err)
			return
		}
		fmt.Fprintf(w, "Your response will be: %v\n", response.Message)
	} else {
		fmt.Fprintln(w, "ERROR! Serving only GET requests...")
	}
}

package handler

import (
	"context"
	"fmt"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) Handle(ctx context.Context) {
	fmt.Println("Hello!")
	return
}

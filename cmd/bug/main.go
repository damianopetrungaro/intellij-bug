package main

import (
	"context"
	internalHandler "github.com/damianopetrungaro/intellij-bug/cmd/bug/internal/handler"
	"github.com/damianopetrungaro/intellij-bug/handler"
)

func main() {
	var h internalHandler.Handler
	h = handler.NewHandler()
	h.Handle(context.Background())
}

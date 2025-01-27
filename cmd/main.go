package main

import (
	"context"
	"github.com/ViciousKit/course-chat-server/internal/app"
)

func main() {
	ctx := context.Background()
	application, _ := app.NewApp(ctx)
	_ = application.Run()
}

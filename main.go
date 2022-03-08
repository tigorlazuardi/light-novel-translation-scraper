package main

import (
	"context"

	"github.com/tigorlazuardi/light-novel-translation-scraper/cmd"
	"github.com/tigorlazuardi/light-novel-translation-scraper/logger"
)

func main() {
	ctx := context.Background()
	err := cmd.Exec(ctx)
	if err != nil {
		logger.Error.Fatal(err)
	}
}

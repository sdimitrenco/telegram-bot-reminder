package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"telegram-bot-reminder/controlers"
	"telegram-bot-reminder/database"
	"telegram-bot-reminder/parsers"
	"telegram-bot-reminder/telegram"
)

var ctx context.Context

func main() {
	// listen signal termination
	terminateC := make(chan os.Signal)
	stopC := make(chan struct{})
	signal.Notify(terminateC, os.Interrupt, syscall.SIGTERM)
	go terminationListening(terminateC, stopC)

	//load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env file")
	}

	//add database to context
	ctx = database.Boot(context.Background())

	//parser daily text
	go controlers.GetDailyText(ctx)

	//parse list users
	go parsers.RunTableParser(ctx)

	//telegram server
	go telegram.Run(ctx)

	// Run listener
	<-stopC

}

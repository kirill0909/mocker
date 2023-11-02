package main

import (
	"context"
	"log"
	"mocker/config"
	"mocker/internal/mocker/repository"
	"mocker/internal/mocker/usecase"
	"mocker/pkg/storage/postgres"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfgFile, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config loaded")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pgDB, err := postgres.InitPGDB(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("PostgreSQL connection stats: %#v", pgDB.Stats())
	}

	mockerRepo := repository.NewMockerPGRepo(cfg, pgDB)
	mockerUC := usecase.NewMockerUC(cfg, mockerRepo)

	exitCh := make(chan os.Signal)
	go func() {
		if err := mockerUC.Mock(ctx); err != nil {
			log.Println(err)
		}
		exitCh <- os.Interrupt
	}()

	signal.Notify(exitCh, os.Interrupt, syscall.SIGINT)
	<-exitCh

	log.Println("Mocking is commplited")

}

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/antondoinikov/calendar-server/hw12_13_14_15_calendar/internal/config"
	"github.com/antondoinikov/calendar-server/hw12_13_14_15_calendar/internal/version"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.json", "Path to configuration file")
}

func main() {

	flag.Parse()
	if flag.Arg(0) == "version" {
		version.PrintVersion()
		return
	}

	config, err := config.NewConfig(configFile)
	if (err != nil) && (config == nil) {
		log.Println("No config error: ", err)
		return
	}
	fmt.Printf("Config: %s \n", *config)
	//logg := logger.New(config.Logger.LogLevel)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		list, err := net.Listen("tcp", net.JoinHostPort(config.GRPC.Host, config.GRPC.Port))
		if err != nil {
			log.Fatalf("failed to mapping port: %s", err.Error())
		}

		fmt.Println("GRPC server is running on port: ", config.GRPC.Port)

		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed the serve: %s", err.Error())
			return
		}
	}()

	wg.Wait()
	//storage := memorystorage.New()
	//calendar := app.New(logg, storage)
	//
	//server := internalhttp.NewServer(logg, calendar)
	//
	//ctx, cancel := signal.NotifyContext(context.Background(),
	//	syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	//defer cancel()
	//
	//go func() {
	//	<-ctx.Done()
	//
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	//	defer cancel()
	//
	//	if err := server.Stop(ctx); err != nil {
	//		logg.Error("failed to stop http server: " + err.Error())
	//	}
	//}()
	//
	//logg.Info("calendar is running...")
	//
	//if err := server.Start(ctx); err != nil {
	//	logg.Error("failed to start http server: " + err.Error())
	//	cancel()
	//	os.Exit(1) //nolint:gocritic
	//}
}

package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cdb "github.com/uhhc/sdk-common-go/db"
	"github.com/uhhc/sdk-common-go/log"
	"github.com/uhhc/sdk-common-go/mongodb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	exgrpc "github.com/uhhc/amf/pkg/application/example/transport/grpc"
	exhttp "github.com/uhhc/amf/pkg/application/example/transport/http"
	"github.com/uhhc/amf/pkg/grpc/pb"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Init server",
	Long:  `Init server`,
	Run: func(cmd *cobra.Command, args []string) {
		// Init logger
		logger := log.NewLogger()
		defer logger.FlushLogger()

		// Init log info
		logger.Infow("service started")
		defer logger.Infow("service ended")

		// Connect to Database
		db, err := cdb.NewDB(*logger, nil)
		if err != nil {
			logger.Errorw("exit", "error", err)
			os.Exit(-1)
		}

		// Connect to MongoDB
		_, err = mongodb.NewMongoClient(*logger)
		if err != nil {
			logger.Errorw("exit", "error", err)
			os.Exit(-1)
		}

		// Deal handlers
		mux := http.NewServeMux()
		mux.Handle("/v1/examples", exhttp.MakeHTTPHandler(db, *logger))
		mux.Handle("/v1/examples/", exhttp.MakeHTTPHandler(db, *logger))
		http.Handle("/", mux)

		// Add swagger handler
		fs := http.FileServer(http.Dir("./assets/swagger-ui"))
		http.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

		// Start HTTP server
		errs := make(chan error)
		go func() {
			port := ":" + viper.GetString("HTTP_PORT")
			logger.Infow("", "transport", "HTTP", "addr", port)
			errs <- http.ListenAndServe(port, nil)
		}()
		go func() {
			c := make(chan os.Signal)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			errs <- fmt.Errorf("%s", <-c)
		}()

		// Start GRPC server
		go func() {
			port := ":" + viper.GetString("GRPC_SERVER_PORT")
			listener, err := net.Listen("tcp", port)
			if err != nil {
				errs <- err
				return
			}

			ctx := context.Background()
			gRPCServer := grpc.NewServer()
			pb.RegisterExampleServiceServer(gRPCServer, exgrpc.NewServer(ctx, db, *logger))
			reflection.Register(gRPCServer)

			logger.Infow("", "transport", "GRPC", "addr", port)
			errs <- gRPCServer.Serve(listener)
		}()

		logger.Errorw("exit", "error", <-errs)
	},
}

// corsControl add CORS header
func corsControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

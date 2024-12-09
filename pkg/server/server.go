package server

import (
	"fmt"

	// apitrace "go.opentelemetry.io/otel/api/trace"
	// "go.opentelemetry.io/otel/instrumentation/grpctrace"

	//graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/kumarabd/service-template/internal/logger"
	"github.com/kumarabd/service-template/internal/metrics"
	"github.com/kumarabd/service-template/pkg/service"
)

type Config struct {
	HTTP *HTTPServerConfig `json:"http" yaml:"http"`
	GRPC *GRPCServerConfig `json:"grpc" yaml:"grpc"`
}

type Handler struct {
	HTTPServer *HTTPServer
	GRPCServer *GRPCServer
	config     *Config
	log        *logger.Handler
}

// panicHandler is the handler function to handle panic errors
func panicHandler(r interface{}) error {
	fmt.Println("600 Error: ", r)
	return fmt.Errorf("internal server error")
}

func New(name string, l *logger.Handler, m *metrics.Handler, config *Config, service *service.Handler) (*Handler, error) {
	// Initiate HTTP Server object
	httpObj := &HTTPServer{
		log:     l,
		service: service,
		metric:  m,
	}
	httpObj.handler = gin.New()
	// Global middleware
	httpObj.handler.Use(gin.Recovery())
	//httpObj.handler.Use(traces.NewGinMiddleware(name))
	gin.SetMode(gin.ReleaseMode)
	httpObj.handler.GET("/healthz", httpObj.HealthHandler)
	httpObj.handler.GET("/readyz", httpObj.HealthHandler)
	httpObj.handler.GET("/metrics", httpObj.MetricsHandler)

	//// Initiate Graphql Server object
	//graphqlObj := &Resolver{
	//	log:     l,
	//	service: service,
	//	metric:  m,
	//}
	//resolvers := graphql_handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graphqlObj}))
	//httpObj.handler.GET("/graphql", ginGraphQLHandler(resolvers))
	//httpObj.handler.POST("/graphql", ginGraphQLHandler(resolvers))
	//httpObj.handler.GET("/playground", ginPlaygroundHandler())

	// Initiate GRPC Server object
	grpcObj := &GRPCServer{
		log:     l,
		service: service,
		metric:  m,
	}
	//middlewares := middleware.ChainUnaryServer(
	// 		grpctrace.UnaryServerInterceptor(tr.Tracer(s.Name).(apitrace.Tracer)),
	//)
	return &Handler{
		HTTPServer: httpObj,
		GRPCServer: grpcObj,
		config:     config,
		log:        l,
	}, nil
}

func (h *Handler) Start(ch chan struct{}) {
	// Start the HTTP server
	go func() {
		h.log.Info().Msgf("started http server on port: %s", h.config.HTTP.Port)
		err := h.HTTPServer.handler.Run(fmt.Sprintf("0.0.0.0:%s", h.config.HTTP.Port))
		h.log.Error().Err(err).Msg("server stopped")
		ch <- struct{}{}
	}()

	//// Start the GRPC server
	//go func() {
	//	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", h.config.GRPC.Port))
	//	if err != nil {
	//		h.log.Error().Err(err).Msg("server stopped")
	//	}

	//	h.log.Info().Msgf("started grpc server on port: %s", h.config.GRPC.Port)
	//	err = h.GRPCServer.handler.Serve(listener)
	//	if err != nil {
	//		h.log.Error().Err(err).Msg("server stopped")
	//	}
	//	ch <- struct{}{}
	//}()
}

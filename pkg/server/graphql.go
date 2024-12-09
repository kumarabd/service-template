package server

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

//import (
//	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
//	"github.com/99designs/gqlgen/graphql/playground"
//	"github.com/gin-gonic/gin"
//	"github.com/kumarabd/service-template/api/graphql/generated"
//	"github.com/kumarabd/service-template/internal/logger"
//	"github.com/kumarabd/service-template/internal/metrics"
//	"github.com/kumarabd/service-template/pkg/service"
//	"google.golang.org/grpc"
//)

//type Resolver struct {
//	handler *grpc.Server
//	service *service.Handler
//	log     *logger.Handler
//	metric  *metrics.Handler
//}

//// Mutation returns generated.MutationResolver implementation.
//func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

//// Query returns generated.QueryResolver implementation.
//func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

//// Subscription returns generated.SubscriptionResolver implementation.
//func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

//type mutationResolver struct{ *Resolver }
//type queryResolver struct{ *Resolver }
//type subscriptionResolver struct{ *Resolver }

//// ginGraphQLHandler wraps the GraphQL handler to make it compatible with Gin
//func ginGraphQLHandler(h *graphql_handler.Server) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		h.ServeHTTP(c.Writer, c.Request)
//	}
//}

//// ginPlaygroundHandler creates a Gin-compatible handler for the GraphQL playground
//func ginPlaygroundHandler() gin.HandlerFunc {
//	h := playground.Handler("GraphQL playground", "/graphql")
//	return func(c *gin.Context) {
//		h.ServeHTTP(c.Writer, c.Request)
//	}
//}

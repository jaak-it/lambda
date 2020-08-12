package helpers

import (
    "context"
    "github.com/aws/aws-lambda-go/events"
    ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/jaak-it/lambda/commons/middleware"
    "github.com/sirupsen/logrus"
)

type Lambda struct {
	Engine *gin.Engine
}

func (lambda *Lambda) HandlerLambda(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ginLambda := ginadapter.New(lambda.Engine)
	return ginLambda.ProxyWithContext(ctx, req)
}

func (lambda *Lambda) Run() error {
	err := lambda.Engine.Run()
	if err != nil {
		return err
	}
	return nil
}

func NewLambda() *Lambda {
	return &Lambda{
		Engine: createServerGin(),
	}
}

func createServerGin() *gin.Engine {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logrus.SetLevel(logrus.TraceLevel)

	if IsProduction() {
		logger.SetLevel(logrus.InfoLevel)
		logrus.SetLevel(logrus.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(cors.Default())
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware(logger))
	r.Use(middleware.AuthMiddleware())

	return r
}

package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth/pb"
)

func AuthRequired(ctx *gin.Context) {
	fmt.Println("AuthRequired --------------")
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c := &ServiceClient{
		Client: InitServiceClient(),
	}

	fmt.Println("token", token[1])

	res, err := c.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	fmt.Println("res", res)
	fmt.Println("err", err)

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.AuthId)

	ctx.Next()
}

package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nusaib44/go-grpc-api-gateway/pkg/auth/models"
	"github.com/nusaib44/go-grpc-api-gateway/pkg/auth/pb"
)

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	body := models.User{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(401, err)
		return

	}
	fmt.Println(body)
	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)

}

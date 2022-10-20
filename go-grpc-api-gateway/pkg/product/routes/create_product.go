package routes

import (
	"context"
	"fmt"
	"go-grpc-api-gateway/pkg/product/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	body := CreateProductRequestBody{}
	fmt.Println("@@@@@@@ 11111")
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println("@@@@@@@ 22222")

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})
	fmt.Println("@@@@@@@ 33333")
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

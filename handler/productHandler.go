package handler

import (
	"chapter3-sesi2/dto"
	"chapter3-sesi2/entity"
	"chapter3-sesi2/pkg/errs"
	"chapter3-sesi2/pkg/helpers"
	"chapter3-sesi2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) productHandler {
	return productHandler{
		productService: productService,
	}
}
func (p productHandler) CreateProduct(ctx *gin.Context) {
	var productRequest dto.NewProductRequest
	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	user := ctx.MustGet("userData").(entity.User)
	newProduct, err := p.productService.CreateProduct(user.Id, productRequest)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusCreated, newProduct)
}
func (p productHandler) UpdateProductById(ctx *gin.Context) {
	var productRequest dto.NewProductRequest

	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	productId, err := helpers.GetParamId(ctx, "productId")
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.productService.UpdateProductById(productId, productRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

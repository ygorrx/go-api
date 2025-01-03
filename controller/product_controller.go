package controller

import (
	"go-api/model"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

type productController struct {
	ProductUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		ProductUseCase: usecase,
	}
}

func (p *productController) GetProducts(c *gin.Context) {
	products, err := p.ProductUseCase.GetProducts()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, products)
}

func (p *productController) CreateProduct(c *gin.Context) {
	var product model.Product
	c.BindJSON(&product)
	createdProduct, err := p.ProductUseCase.CreateProduct(product)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, createdProduct)
}
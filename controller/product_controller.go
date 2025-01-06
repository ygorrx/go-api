package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

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

func (p *productController) GetProductById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response := model.Response{
			Message: "id is required",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "id needs to be a number",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.ProductUseCase.GetProductById(productID)
	if err != nil {
		c.JSON(500, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "product not found",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(200, product)
}

func (p *productController) CreateProduct(c *gin.Context) {
	var product model.Product
	c.BindJSON(&product)
	createdProduct, err := p.ProductUseCase.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, createdProduct)
}
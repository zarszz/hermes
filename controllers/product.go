package controllers

import (
	"hermes/controllers/entity"
	"hermes/models/product/repo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jmoiron/sqlx"
)

type Product struct {
	productService *repo.ProductService
}

func NewProduct(conn *sqlx.DB) *Product {
	return &Product{productService: repo.NewProductService(conn)}
}

func (s *Product) Get(c *gin.Context) {
	products, err := s.productService.GetAll()
	if err != nil {		
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": err.Error()},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"products": &products.Products},
	)
}

func (s *Product) GetByPK(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := s.productService.GetByPK(id)
	if err != nil {		
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": err.Error()},
		)
		return
	}
	if product == nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "product not found",
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"product": product},
	)
}

func (s *Product) Create(c *gin.Context) {
	var form entity.ProductInput
	if c.ShouldBindWith(&form, binding.JSON) != nil {
		c.JSON(
			http.StatusNotAcceptable,
			gin.H{"message": "invalid data"},
		)
		c.Abort()
		return
	}
	err := s.productService.InsertOne(form)
	if err != nil {		
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": err.Error()},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"message": "successfully created"},
	)
}

func (s *Product) Update(c *gin.Context) {
	var form entity.ProductInput
	id := c.Param("id")
	if c.ShouldBindWith(&form, binding.JSON) != nil {
		c.JSON(
			http.StatusNotAcceptable,
			gin.H{"message": "invalid data"},
		)
		c.Abort()
		return
	}
	form.Id, _ = strconv.Atoi(id)
	err := s.productService.UpdateOne(form)
	if err != nil {		
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": err.Error()},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"message": "successfully updated"},
	)
}

func (s *Product) DeleteByPK(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := s.productService.GetByPK(id)
	if err != nil {		
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": err.Error()},
		)
		return
	}
	if product == nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "product not found",
			},
		)
		return
	}

	errDelete := s.productService.DeleteOne(product.Id)

	if errDelete != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": err.Error()},
		)
		return
	}
	if product == nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "product not found",
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"product": product},
	)
}

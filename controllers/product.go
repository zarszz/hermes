package controllers

import (
	"hermes/controllers/entity"
	"hermes/models/product/repo"
	"hermes/utils"
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
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, err)
		return
	}
	utils.HandleResponse(c, "successfully retrieve products", http.StatusOK, c.Request.Method, products.Products)
}

func (s *Product) GetByPK(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := s.productService.GetByPK(id)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, err)
		return
	}
	if product == nil {
		utils.HandleErrorResponse(c, http.StatusNotFound, c.Request.Method, "product not found")
		return
	}
	utils.HandleResponse(c, "successfully retrieve product", http.StatusOK, c.Request.Method, product)
}

func (s *Product) Create(c *gin.Context) {
	var form entity.ProductInput
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "invalid data")
		return
	}
	err := s.productService.InsertOne(form)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, err)
		return
	}
	utils.HandleResponse(c, "successfully created", http.StatusCreated, c.Request.Method, "")
}

func (s *Product) Update(c *gin.Context) {
	var form entity.ProductInput
	id := c.Param("id")
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "invalid data")
		return
	}
	form.Id, _ = strconv.Atoi(id)
	err := s.productService.UpdateOne(form)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, err)
		return
	}
	utils.HandleResponse(c, "successfully updated", http.StatusOK, c.Request.Method, "")

}

func (s *Product) DeleteByPK(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := s.productService.GetByPK(id)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, err)
		return
	}
	if product == nil {
		utils.HandleErrorResponse(c, http.StatusNotFound, c.Request.Method, "product not found")
		return
	}

	errDelete := s.productService.DeleteOne(product.Id)

	if errDelete != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, err)
		return
	}
	if product == nil {
		utils.HandleErrorResponse(c, http.StatusNotFound, c.Request.Method, err)
		return
	}
	utils.HandleResponse(c, "successfully deleted", http.StatusOK, c.Request.Method, "")

}

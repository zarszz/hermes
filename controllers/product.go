package controllers

import (
	"fmt"
	"hermes/controllers/entity"
	"hermes/models/product/repo"
	"hermes/utils"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
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
	sku := c.PostForm("sku")
	name := c.PostForm("name")
	image, _ := c.FormFile("display")
	if sku == "" {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "sku data required")
		return
	}
	if name == "" {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "name data required")
		return
	}
	if image == nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "image data required")
		return
	}
	product, err := s.productService.GetBySKU(sku)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, c.Request.Method, err.Error())
		return
	}
	if product != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "sku is already used")
		return
	}
	workDirPath, _ := os.Getwd()
	dst := filepath.Join(workDirPath, "images", image.Filename)
	if image != nil {
		errSaveImage := c.SaveUploadedFile(image, dst)
		if errSaveImage != nil {
			fmt.Printf("an error occured when save image : %s", errSaveImage.Error())
		}
	}

	form.Display = os.Getenv("ROOT_URI") + "/images/" + image.Filename
	form.Name = name
	form.Sku = sku

	errInsert := s.productService.InsertOne(form)
	if errInsert != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, errInsert)
		return
	}
	utils.HandleResponse(c, "successfully created", http.StatusCreated, c.Request.Method, "")
}

func (s *Product) Update(c *gin.Context) {
	var form entity.ProductInput
	id := c.Param("id")
	sku := c.PostForm("sku")
	name := c.PostForm("name")
	image, _ := c.FormFile("display")
	if sku == "" {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "sku data required")
		return
	}
	if name == "" {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "name data required")
		return
	}
	if image == nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "image data required")
		return
	}
	product, finSkuErr := s.productService.GetBySKU(sku)
	if finSkuErr != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, c.Request.Method, finSkuErr.Error())
		return
	}
	if product != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "sku is already used")
		return
	}
	workDirPath, _ := os.Getwd()
	dst := filepath.Join(workDirPath, "images", image.Filename)
	if image != nil {
		errSaveImage := c.SaveUploadedFile(image, dst)
		if errSaveImage != nil {
			fmt.Printf("an error occured when save image : %s", errSaveImage.Error())
		}
	}

	form.Display = os.Getenv("ROOT_URI") + "/images/" + image.Filename
	form.Name = name
	form.Sku = sku
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

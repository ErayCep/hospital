package handlers

import (
	"hospital/db"
	"hospital/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetHospitalHandler(c *gin.Context) {
	id_string := c.Param("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get id from request",
		})
	}
	hospital, err := db.GetHospitalWithID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get hospital",
		})
	}

	c.JSON(http.StatusOK, hospital)
}

func (h *Handlers) GetHospitalsHandler(c *gin.Context) {
	hospitals, err := h.Storage.GetHospitals()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get hospitals",
		})
	}

	c.JSON(http.StatusOK, hospitals)
}

func (h *Handlers) PostHospitalHandler(c *gin.Context) {
	var hospital models.Hospital
	if c.Bind(&hospital) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}

	err := h.Storage.AddHospital(&hospital)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create hospital",
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handlers) DeleteHospitalHandler(c *gin.Context) {
	id_string := c.Param("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read id from body",
		})
	}

	err = h.Storage.DeleteHospital(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete hospital",
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}

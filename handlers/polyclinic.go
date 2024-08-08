package handlers

import (
	"hospital/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetPolyclinicsHandler(c *gin.Context) {
	polyclinics, err := h.Storage.GetPolyclinics()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get polyclinic",
		})

		return
	}

	c.JSON(http.StatusOK, polyclinics)
}

func (h *Handlers) GetPolyclinicHandler(c *gin.Context) {
	id_string := c.Param("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get hospital id",
		})

		return
	}

	p_id_string := c.Param("p_id")
	p_id, err := strconv.Atoi(p_id_string)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get polyclinic id",
		})
	}

	polyclinic, err := h.Storage.GetPolyclinic(id, p_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get polyclinic",
		})

		return
	}

	c.JSON(http.StatusOK, polyclinic)
}

func (h *Handlers) PostPolyclinicHandler(c *gin.Context) {
	var body struct {
		ID         int
		Name       string
		City       string
		County     string
		Address    string
		TotalStaff uint32
	}

	id_string := c.Param("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get hospital id",
		})

		return
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	var polyclinic = models.Polyclinic{City: body.City, County: body.County, Address: body.Address, TotalStaff: body.TotalStaff, HospitalID: id}
	err = h.Storage.AddPolyclinic(polyclinic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create polyclinic",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handlers) DeletePolyclinicHandler(c *gin.Context) {
	id_string := c.Param("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read id from request",
		})

		return
	}

	err = h.Storage.DeletePolyclinic(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to delete polyclinic",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handlers) GetPolyclinicHandler2(c *gin.Context) {
	id_string := c.Param("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read id from request",
		})

		return
	}

	var polyclinic models.Polyclinic
	polyclinic, err = h.Storage.GetPolyclinic2(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get polyclinic",
		})

		return
	}

	c.JSON(http.StatusOK, polyclinic)
}

func (h *Handlers) PostPolyclinicHandler2(c *gin.Context) {
	var body struct {
		ID         int
		Name       string
		City       string
		County     string
		Address    string
		TotalStaff uint32
		HospitalID int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	var polyclinic = models.Polyclinic{City: body.City, County: body.County, Address: body.Address, TotalStaff: body.TotalStaff, HospitalID: body.HospitalID}
	err := h.Storage.AddPolyclinic(polyclinic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to add new polyclinic",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

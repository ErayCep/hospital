package handlers

import (
	"hospital/helpers"
	"hospital/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) Signup(c *gin.Context) {
	// Get email, phone, tc and password from request body and check for errors in request.
	var body struct {
		Email    string
		Phone    string
		TC       string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// Add new staff to database
	err = h.Storage.AddStaff(body.Email, body.Phone, body.TC, string(hash), 2)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create staff",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handlers) Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var staff models.Staff
	//h.Storage.DB.First(&staff, "email = ?", body.Email)
	h.Storage.GetStaffWithEmail(&staff, body.Email)

	if staff.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := helpers.GenerateJWTWithID(staff.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate jwt token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (h *Handlers) Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"info": "logged in",
	})
}

func (h *Handlers) GetPolyclinicStaffHandler(c *gin.Context) {
	id_string := c.Param("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read id from request",
		})

		return
	}

	var staff []models.Staff
	staff, err = h.Storage.GetPolyclinicStaff(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get staff",
		})

		return
	}

	c.JSON(http.StatusOK, staff)
}

func (h *Handlers) PostPolyclinicStaffHandler(c *gin.Context) {
	id, err := helpers.ReadIDFromRequest(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read id from request",
		})
	}

	var body struct {
		FirstName  string
		LastName   string
		Email      string
		Phone      string
		Password   string
		TC         string
		Privileged bool
		Title      string
		Skill      string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	var staff = models.Staff{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Phone: body.Phone, Password: string(hash), TC: body.TC, Privileged: body.Privileged, PolyclinicID: id, Title: body.Title, Skill: body.Skill}
	err = h.Storage.PostStaff(staff)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to add new staff",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handlers) DeletePolyclinicStaffHandler(c *gin.Context) {
	id, err := helpers.ReadIDFromRequest("id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read id from request body",
		})

		return
	}

	staff_id, err := helpers.ReadIDFromRequest("s_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read id from request body",
		})

		return
	}

	err = h.Storage.DeletePolyclinicStaff(id, staff_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to delete staff",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handlers) GetStaffHandler(c *gin.Context) {
	id, err := helpers.ReadIDFromRequest("id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read id from request body",
		})

		return
	}

	var staff models.Staff
	staff, err = h.Storage.GetStaffWithID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get staff",
		})

		return
	}

	c.JSON(http.StatusOK, staff)
}

func (h *Handlers) GetStaffsHandler(c *gin.Context) {
	var staffs []models.Staff
	staffs, err := h.Storage.GetStaffs()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get staffs",
		})

		return
	}

	c.JSON(http.StatusOK, staffs)
}

func (h *Handlers) ChangePasswordHandler(c *gin.Context) {
	id, err := helpers.ReadIDFromRequest("id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read id from request body",
		})

		return
	}

	staff, err := h.Storage.GetStaffWithID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to find staff",
		})
	}

	var body struct {
		OldPassword      string
		NewPassword      string
		NewPasswordAgain string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	if body.NewPassword != body.NewPasswordAgain {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "new passwords don't match",
		})

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(body.OldPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong password",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	result := h.Storage.DB.Model(&staff).Where("id = ?", id).Update("password", hash)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to update password",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handlers) DeleteStaffHandler(c *gin.Context) {
	id, err := helpers.ReadIDFromRequest("id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read id from request body",
		})

		return
	}

	err = h.Storage.DeleteStaff(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to delete staff",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

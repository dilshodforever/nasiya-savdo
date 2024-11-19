package handler

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

// Check handles user Check
// @Summary Check a user
// @Description Check a user with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} string "Check successful"
// @Failure 401 {string} string "Unauthorized"
// @Router /check/location [post]
func (h *Handler)Check(ctx *gin.Context) {
	deviceInfo, location, err := getClientDetails(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve client details"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"device_info": deviceInfo,
		"location":    location,
	})
}

// getClientDetails retrieves the device information and location based on IP.
func getClientDetails(r *http.Request) (string, string, error) {
	// Qurilma haqida ma'lumot
	ua := user_agent.New(r.Header.Get("User-Agent"))
	browserName, browserVersion := ua.Browser()
	deviceInfo := fmt.Sprintf("OS: %s, Device: %s, Browser: %s %s", ua.OS(), ua.Platform(), browserName, browserVersion)

	// IP manzilni olish
	clientIP := getClientIP(r)

	// Geolokatsiya ma'lumotlarini olish
	location, err := getLocationFromIP(clientIP)
	if err != nil {
		return deviceInfo, "Unknown", err
	}

	return deviceInfo, location, nil
}

func getClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	// Agar IP localhost bo'lsa, sinov uchun soxta IP o'rnatiladi
	if ip == "127.0.0.1" || ip == "::1" {
		ip = "8.8.8.8" // Google DNS sinov IP
	}

	fmt.Println("Detected Client IP:", ip)
	return strings.TrimSpace(ip)
}

// getLocationFromIP fetches geolocation data using an external API.
func getLocationFromIP(ip string) (string, error) {
	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Country string `json:"country"`
		City    string `json:"city"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s, %s", result.City, result.Country), nil
}

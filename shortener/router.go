package shortener

import (
	"errors"
	"net/http"
	"url-shortener/random"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
)

type ShortenerRequest struct {
	OriginalURL string `json:"originalUrl"`
	ShortenURL  string `json:"shortenUrl"`
}

func initRouter(router *gin.Engine) {
	router.POST("/shortener", func(c *gin.Context) {
		var req ShortenerRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": ErrInvalidRequest.Error(),
			})
			return
		}

		shortener := Shortener{
			OriginalURL: req.OriginalURL,
			ShortenURL:  req.ShortenURL,
		}

		if err := shortener.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if shortener.ShortenURL == "" {
			shortener.ShortenURL = random.GenerateRandomString(8)
		}

		err := StoreShortener(c, shortener)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	})
}

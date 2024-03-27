package shortener

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	"url-shortener/random"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
)

type ShortenerRequest struct {
	OriginalURL string `json:"originalUrl"`
	ShortenURL  string `json:"shortenUrl"`
}

func initRouter(router *gin.Engine) {
	router.GET("/:shorten_url", func(c *gin.Context) {
		shorten := c.Param("shorten_url")
		shortener, err := GetShorten(c, shorten)
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "shorten url not found",
			})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}

		time.Sleep(2 * time.Second)

		c.Redirect(http.StatusTemporaryRedirect, shortener.OriginalURL)
	})

	router.POST("/shortener", func(c *gin.Context) {
		var req ShortenerRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": ErrInvalidRequest.Error(),
			})
			return
		}

		shortenerReq := Shortener{
			OriginalURL: req.OriginalURL,
			ShortenURL:  req.ShortenURL,
		}

		if err := shortenerReq.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if shortenerReq.ShortenURL == "" {
			shortenerReq.ShortenURL = random.GenerateRandomString(8)
		}

		shorten, err := GetShorten(c, req.ShortenURL)
		if err != nil && err != pgx.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}

		// if shortener is empty then create new shortener
		if shorten.UUID == "" {
			err = StoreShortener(c, shortenerReq)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "internal server error",
				})
				return
			}
		} else if shorten.OriginalURL != shortenerReq.OriginalURL {
			err = UpdateShortener(c, shortenerReq.OriginalURL, shorten.ShortenURL)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "internal server error",
				})
				return
			}

			err = StoreShortenerLog(c, shorten.UUID, shortenerReq.OriginalURL, shorten.OriginalURL)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "internal server error",
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message":      "success shorten the requested URL",
			"shortenedUrl": fmt.Sprintf("%s/%s", c.Request.Host, shortenerReq.ShortenURL),
		})
	})
}

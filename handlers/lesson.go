package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"education/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/russross/blackfriday/v2"
)

func GetLecture(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid lecture ID"})
        }

        var lesson models.Lesson
        if err := db.First(&lesson, id).Error; err != nil {
            fmt.Print(lesson)
            return c.JSON(http.StatusNotFound, map[string]string{"error": "lecture not found"})
        }

        lesson.Content = string(blackfriday.Run([]byte(lesson.Content)))

        return c.JSON(http.StatusOK, lesson)
    }
}

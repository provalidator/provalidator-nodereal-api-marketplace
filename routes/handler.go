package routes

import (
	"os"

	"github.com/provalidator-nodereal-api-marketplace/dblayer"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db dblayer.DBLayer
}

func (h *Handler) GetMemberData(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "dsn 오류"})
		return
	}
	contents, err := h.db.GetMemberDataFromToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contents)
}

type HandlerInterface interface {
	GetMemberData(c *gin.Context)
}

// HandlerInterface 생성자 정의
func NewHandler() (HandlerInterface, error) {
	dsn := os.Getenv("MYSQL_USER_ID") + ":" + os.Getenv("MYSQL_USER_PW") + "@tcp(" + os.Getenv("MYSQL_SERVER_URL") + ":" + os.Getenv("MYSQL_SERVER_PORT") + ")/" + os.Getenv("MYSQL_SELECT_DB_NAME") + "?charset=utf8"
	// DBORM 초기화 - DBORM 생성자 호출
	db, err := dblayer.NewORM("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}

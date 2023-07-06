package dblayer

import "github.com/provalidator-nodereal-api-marketplace/models"

// DB 인터페이스 정의
type DBLayer interface {
	GetMemberDataFromToken() ([]models.ApiUsage, error)
}

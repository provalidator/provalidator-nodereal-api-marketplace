package dblayer

import (
	"database/sql"

	"github.com/provalidator-nodereal-api-marketplace/log"
	"github.com/provalidator-nodereal-api-marketplace/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ----------------------------------
// DBORM 구조체 정의
// *gorm.DB 타입을 임베드
type DBORM struct {
	*gorm.DB
}

// DBORM 생성자 정의
func NewORM(dbengine string, dsn string) (*DBORM, error) {
	sqlDB, err := sql.Open(dbengine, dsn)
	if err != nil {
		log.Logger.Error.Println(err)
	}
	// gorm.Open은 *gorm.DB 타입을 초기화한다.
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return &DBORM{
		DB: gormDB,
	}, err
}

// DBORM 함수 정의
func (db *DBORM) GetMemberDataFromToken() (contents []models.ApiUsage, err error) {
	return contents, db.Find(&contents).Error
}

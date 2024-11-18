package migrations

import (
	"authSSO/internal/config"
	rlog "authSSO/internal/log"
	"authSSO/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Migrations struct {
	Db  *gorm.DB
	Log *rlog.Logger
}

func NewMigrations(config *config.Config, log *rlog.Logger) *Migrations {
	db, err := gorm.Open(postgres.Open(config.DbLink), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	m := &Migrations{
		Db:  db,
		Log: log,
	}

	return m

}

func (m *Migrations) Up() error {
	err := m.Db.AutoMigrate(&models.User{}, &models.Role{}, &models.UserRole{})

	if err != nil {

		return err
	}
	return nil
}

func (m *Migrations) Down() error {
	err := m.Db.Migrator().DropTable(&models.User{}, &models.Role{}, &models.UserRole{})

	if err != nil {
		return err
	}
	return nil
}

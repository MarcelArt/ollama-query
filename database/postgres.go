package database

import (
	"fmt"
	"strconv"

	"github.com/MarcelArt/ollama-query/config"
	"github.com/MarcelArt/ollama-query/models"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func ConnectDB() {
	p := config.Env.DBPort
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable", config.Env.DBHost, port, config.Env.DBUser, config.Env.DBPassword, config.Env.DBName, config.Env.DBSchema)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// MigrateDB()
}

func GetDB() *gorm.DB {
	return db
}

func MigrateDB() error {
	err := db.AutoMigrate(
		// models.User{},
		// models.AuthorizedDevice{},
		models.Class{},
		models.Teacher{},
		models.ClassTeacher{},
		models.Student{},
	)
	fmt.Println("Database Migrated")

	return err
}

func DropDB() error {
	err := db.Migrator().DropTable(
		// models.User{},
		// models.AuthorizedDevice{},
		models.Class{},
		models.Teacher{},
		models.ClassTeacher{},
		models.Student{},
	)
	fmt.Println("Database Droped")

	return err
}

func SeedDB() error {
	tx := db.Begin()
	defer tx.Rollback()

	for i := 1; i <= 12; i++ {
		class := models.Class{
			Name: fmt.Sprintf("Grade %d", i),
		}
		if err := tx.Create(&class).Error; err != nil {
			return err
		}

		var teacher models.Teacher
		if err := faker.FakeData(&teacher, options.WithFieldsToIgnore("ID", "CreatedAt", "UpdatedAt", "DeletedAt")); err != nil {
			return err
		}

		if err := tx.Create(&teacher).Error; err != nil {
			return err
		}

		classTeacher := models.ClassTeacher{
			ClassID:   class.ID,
			TeacherID: teacher.ID,
		}
		if err := tx.Create(&classTeacher).Error; err != nil {
			return err
		}

		for j := 0; j < 10; j++ {
			var student models.StudentDTO
			if err := faker.FakeData(&student, options.WithFieldsToIgnore("ID", "CreatedAt", "UpdatedAt", "DeletedAt")); err != nil {
				return err
			}

			student.ClassID = class.ID
			if err := tx.Create(&student).Error; err != nil {
				return err
			}
		}
	}

	return tx.Commit().Error

}

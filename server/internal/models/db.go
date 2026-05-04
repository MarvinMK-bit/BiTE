package models

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	gormDB *gorm.DB
	once   sync.Once
)

func Db() *gorm.DB {
	once.Do(func() {
		var err error

		env := os.Getenv("GO_ENV")
		log.Println("GO_ENV:", env)

		if env == "development" {
			err = godotenv.Load()
			if err != nil {
				log.Fatalf("Error loading .env file")
			}
			log.Println("Loaded .env var file")
		}

		switch env {
		case "development":
			gormDB, err = gorm.Open(postgres.Open(os.Getenv("BiTE_DEV_DSN")), &gorm.Config{
				SkipDefaultTransaction: true,
				PrepareStmt:            true,
				// Logger:                 logger.Default.LogMode(logger.Info),
			})
			if err != nil {
				log.Fatal("Failed to connect to the database: ", err)
			}
		case "production":
			gormDB, err = gorm.Open(postgres.Open(os.Getenv("BiTE_PROD_DSN")), &gorm.Config{
				SkipDefaultTransaction: true, PrepareStmt: true,
			})
			if err != nil {
				log.Fatal("Failed to connect to the database: ", err)
			}
		default:
			log.Fatal("Unrecognized GO_ENV:", env)
		}

		log.Println("Connected to postgres successfully")

		err = gormDB.AutoMigrate(&User{}, &OTP{}, &Session{}, &Location{},
			&QuizCategory{}, &Quiz{}, &QuizDuplicate{}, &Question{}, &Answer{},
			&Attempt{}, &AttemptStatus{}, &AttemptDuration{}, &AIPreview{},
			&Ranking{}, &SiteVisit{}, &QuizUserProgress{}, &Attachment{}, &SatsReward{},
			&SatsRewardAddress{}, &SatsRewardTransaction{}, &SatsRewardOperation{},
			&CategoryCertificate{}, &CategoryCertificateQuizzes{}, &CertificateAwarded{},
			&ChessPuzzle{}, &ChessPuzzleRound{}, &ChessUserPuzzleRating{})
		if err != nil {
			log.Fatal("Failed to make auto migration", err)
		}
		log.Println("Auto Migration successful")
	})

	return gormDB
}

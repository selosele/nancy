package main

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
	"github.com/selosele/nancy/handlers"
	"github.com/selosele/nancy/models"
)

func main() {

	// 환경변수 로드
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Cloudinary 클라이언트 초기화
	cld, err := cloudinary.New()
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary, %v\n", err)
	}

	ctx := context.Background()

	// Handler 구조체 생성
	h := models.Handler{
		Cld: cld,
		Ctx: ctx,
	}

	// 라우트 설정
	handlers.Setup(h)
}

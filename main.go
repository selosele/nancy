package main

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
	"github.com/selosele/nancy/handlers"
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

	// 라우트 설정
	handlers.Init(cld)
}

dev:
	docker-compose up -d
	
dev-down:
	docker-compose down

install-modules:
	go get github.com/gofiber/fiber/v2
	go get github.com/go-playground/validator/v10
	go get -u gorm.io/gorm
	go get gorm.io/driver/postgres
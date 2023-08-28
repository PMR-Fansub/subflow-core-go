package swagger

//go:generate go run -mod=mod github.com/swaggo/swag/cmd/swag fmt -g api.go --dir=../../internal/api

//go:generate go run -mod=mod github.com/swaggo/swag/cmd/swag init -g api.go --dir=../../internal/api --output ../../docs

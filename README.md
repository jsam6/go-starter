go run main.go
go build -o binary main.go

go get -u github.com/go-chi/chi/v5

go mod tidy # when u dont want to use go get and add the import library and use go mod tidy to tell go to find for new library use
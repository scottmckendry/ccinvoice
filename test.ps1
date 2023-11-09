go test -v -covermode=count -coverprofile=coverage
go tool cover -html=coverage


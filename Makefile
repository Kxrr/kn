buildMain:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o scripts/kn_main scripts/main.go

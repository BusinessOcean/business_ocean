//go:build tools
// +build tools

//go:generate golangci-lint run ./...
//go:generate mockgen -source=source.go -destination=mocks/mock_source.go -package=mocks

package tools

import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	// Add other tools as needed
)

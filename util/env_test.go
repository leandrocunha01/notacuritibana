package util

import (
	"fmt"
	"os"
	"testing"
)

func TestEnvFile(t *testing.T) {
	test := os.Getenv("CERT_KEY")
	fmt.Println(test)
}

package test

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func TestUuid(t *testing.T) {
	u1 := uuid.Must(uuid.NewV4(), nil)
	fmt.Printf("UUIDv4: %s\n", u1)
}

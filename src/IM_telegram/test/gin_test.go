package test

import (
	"go_project/src/IM_telegram/router"
	"testing"
)

func TestGin(t *testing.T) {
	r := router.Router()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

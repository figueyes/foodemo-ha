package utils

import (
	"go-course/demo/app/shared/log"
	"os"
)

func GetEnvironments() {
	for _, pair := range os.Environ() {
		log.Info(pair)
	}
}

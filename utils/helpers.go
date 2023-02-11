package utils

import (
	"log"
)

func Sptr(s string) *string {
	return &s
}

func Iptr(i int) *int {
	return &i
}

func ToString(array []string) {
	for key, item := range array {
		log.Printf("key: %v, value: %v", key, item)
	}
}

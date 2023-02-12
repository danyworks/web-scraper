package common

import (
	"context"
	"errors"
	"github.com/go-rod/rod"
	"log"
)

func HandleError(err error) {
	var evalErr *rod.ErrEval
	if errors.Is(err, context.DeadlineExceeded) { // timeout error
		log.Println("timeout err")
	} else if errors.As(err, &evalErr) { // eval error
		log.Println(evalErr.LineNumber)
	} else if err != nil {
		log.Println("can't handle", err)
	}
}

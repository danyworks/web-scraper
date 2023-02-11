package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/daniyalibrahim/notybot/internal/models"
	"io"
	"log"
	"net/http"
	"strings"
)

/*
Not Required for the Webhook but could be used for custom api requests!!
*/
func ParseErrors(w http.ResponseWriter, err error) models.Response {
	response := &models.Response{}
	response.Status = "failure"

	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	log.Printf("Error: %s", err.Error())
	w.WriteHeader(http.StatusBadRequest)
	switch {
	case errors.As(err, &syntaxError):
		response.Message = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
	case errors.Is(err, io.ErrUnexpectedEOF):
		response.Message = fmt.Sprintf("Request body contains badly-formed JSON")
	case errors.As(err, &unmarshalTypeError):
		response.Message = fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		response.Message = fmt.Sprintf("Request body contains unknown field %s", fieldName)
	case errors.Is(err, io.EOF):
		response.Message = "Request body must not be empty"
	case err.Error() == "http: request body too large":
		response.Message = "Request body must not be larger than 1MB"
		w.WriteHeader(http.StatusRequestEntityTooLarge)
	default:
		response.Message = "Oops! Something went wrong"
		w.WriteHeader(http.StatusInternalServerError)
	}

	return *response
}

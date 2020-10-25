package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const IdQueryparam = "id"

func SendJsonError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"errorCode": http.StatusBadRequest,
		"error":     "Body is not a valid JSON",
	})
}

func SendInternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"errorCode": http.StatusInternalServerError,
		"error":     err.Error(),
	})
}

func SendIntegerParsingError(ctx *gin.Context, i string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"errorCode": http.StatusBadRequest,
		"error":     fmt.Sprintf("Id is not a valid integer : %s", i),
	})
}

func SendBadRequestError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"errorCode": http.StatusBadRequest,
		"error":     err.Error(),
	})
}

func SendEntityCreatedResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "")
}

func GetIncludeFields(include string, knownFields []string) (*[]string, error) {
	var fieldsToInclude []string
	if include != "" {
		additionalFields := strings.Split(include, ",")
		var unknownIncl []string
		for _, element := range additionalFields {
			for _, knownField := range knownFields {
				if element != knownField {
					unknownIncl = append(unknownIncl, element)
				} else {
					fieldsToInclude = append(fieldsToInclude, element)
				}
			}
		}
		if len(unknownIncl) > 0 {
			errorMessage := "Unknown fields to include: "
			for index, el := range unknownIncl {
				if index == 0 {
					errorMessage = errorMessage + el
				} else {
					errorMessage = ", " + errorMessage + el
				}
			}

			return nil, errors.New(errorMessage)
		}
	}
	return &fieldsToInclude, nil
}

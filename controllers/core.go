package controllers

import (
	"errors"
	"fmt"
	"strings"
)

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
		fmt.Printf("%v", len(unknownIncl))
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

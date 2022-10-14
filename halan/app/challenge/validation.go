package challenge

import (
	"github.com/gin-gonic/gin"
	"halan/common/exchange"
	"halan/common/helper"
	"regexp"
	"strings"
)

func WordValidation(g *gin.Context) (bool, exchange.Word) {
	var request exchange.Word
	_ = g.ShouldBindJSON(&request)
	request.Word = strings.TrimSpace(request.Word)
	if request.Word != "" {
		return true, request
	}
	helper.BadRequestResponse(g, helper.Required("word"), nil)
	return false, request
}

func RLEWordValidation(g *gin.Context) (bool, exchange.Word) {
	valid, request := WordValidation(g)
	if !valid {
		return valid, request
	}
	ok := regexp.MustCompile(`\d`).MatchString(request.Word)
	if !ok {
		return valid, request
	}
	helper.BadRequestResponse(g, "word field must has no numbers", nil)
	return false, request
}

func Number2DArrayValidation(g *gin.Context) (bool, exchange.Number2DArray) {
	var request exchange.Number2DArray
	_ = g.ShouldBindJSON(&request)
	if len(request.Numbers2D) == 0 {
		helper.BadRequestResponse(g, helper.Required("numbers_2d"), nil)
		return false, request
	}
	return true, request
}

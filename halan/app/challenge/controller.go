package challenge

import (
	"github.com/gin-gonic/gin"
	"halan/common/exchange"
	"halan/common/helper"
)

// Problem1 godoc
// @Summary Get String Palindrome
// @Param request body exchange.Word true "required request filters"
// @Produce json
// @Success 200
// @failure 400
// @Router /challenge/problem-1 [post]
func Problem1(g *gin.Context) {
	valid, request := WordValidation(g)
	if !valid {
		return
	}
	ok := isPalindrome(request.Word)
	if !ok {
		helper.BadRequestResponse(g, request.Word+" Is Not Palindrome", nil)
		return
	}
	helper.OkResponse(g, request.Word+" Is Palindrome", nil)
}

// Problem2 godoc
// @Summary Write the Linux command needed to change a file name from original.txt to changed.txt
// @Produce json
// @Success 200
// @Router /challenge/problem-2 [post]
func Problem2(g *gin.Context) {
	command := GetLinuxCommandLineFileNameChange()
	helper.OkResponse(g, command, nil)
}

// Problem3RLE godoc
// @Summary  implement a function runLengthEncode
// @Description Given a string containing characters (a-z), implement a function runLengthEncode that compresses repeated ‘runs’ of the same character by storing the length of that run
// @Param request body exchange.Word true "required request filters"
// @Produce json
// @Success 200 {object} exchange.Word
// @failure 400
// @Router /challenge/problem-3/rle [post]
func Problem3RLE(g *gin.Context) {
	valid, request := RLEWordValidation(g)
	if !valid {
		return
	}
	str := RunLengthEncode(request.Word)
	helper.OkResponse(g, "", exchange.Word{Word: str})
}

// Problem3RLD godoc
// @Summary provide a function runLengthDecode
// @Description provide a function runLengthDecode to reverse the compression. The output can be anything, as long as you can recreate the input with it
// @Param request body exchange.Word true "required request filters"
// @Produce json
// @Success 200
// @failure 400 {object} exchange.Word
// @Router /challenge/problem-3/rld [post]
func Problem3RLD(g *gin.Context) {
	valid, request := WordValidation(g)
	if !valid {
		return
	}
	str := RunLengthDecode(request.Word)
	helper.OkResponse(g, "", exchange.Word{Word: str})
}

// Problem4 godoc
// @Summary Define a function compose that implements composition.
// @Description Let f and g be two one-argument functions. The composition f after g is defined to be the function . Define a function compose that implements composition. For example, if inc is a function that adds 1 to its argument, and square is a function that squares its argument
// @Param request body exchange.Number true "required request filters"
// @Produce json
// @Success 200 {object} exchange.Value
// @Router /challenge/problem-4 [post]
func Problem4(g *gin.Context) {
	var request exchange.Number
	_ = g.ShouldBindJSON(&request)
	rslt := CompositionService(request.Number)
	helper.OkResponse(g, "", exchange.Value{Value: rslt})
}

// Problem5 godoc
// @Summary function unique
// @Description Write a function unique that takes an array of strings as input and returns an array of the unique entries in the input
// @Param request body exchange.WordArray true "required request filters"
// @Produce json
// @Success 200 {object} exchange.WordArray
// @Router /challenge/problem-5 [post]
func Problem5(g *gin.Context) {
	var request exchange.WordArray
	_ = g.ShouldBindJSON(&request)
	rslt := GetUniqueFromStringList(request.Words)
	helper.OkResponse(g, "", exchange.WordArray{Words: rslt})
}

// Problem6 godoc
// @Summary the transpose of a matrix
// @Description In linear algebra, the transpose of a matrix is another matrix created by writing the rows of as the columns of a t,
// @Param request body exchange.Number2DArray true "required request filters"
// @Produce json
// @Success 200 {object} exchange.Number2DArray
// @Router /challenge/problem-6 [post]
func Problem6(g *gin.Context) {
	valid, request := Number2DArrayValidation(g)
	if !valid {
		return
	}
	rslt := Transpose(request.Numbers2D)
	helper.OkResponse(g, "", exchange.Number2DArray{Numbers2D: rslt})
}

// Problem7 godoc
// @Summary Containers A, B
// @Description You are given 2 containers: A and B. Container A can hold 5 litres of water, while container B can hold 3 litres. You are also given a source of water that you can use as you wish. Show how you can use the containers and the water source to put exactly 4 litres of water in container
// @Produce json
// @Success 200 {object} exchange.Word
// @Router /challenge/problem-7 [post]
func Problem7(g *gin.Context) {
	rslt := ContainersABWithWater()
	helper.OkResponse(g, "", exchange.Word{Word: rslt})
}

// Problem8 godoc
// @Summary Find the index of the first duplicate element in the array
// @Description Given an integer array of length n, find the index of the first duplicate element in the array and state the runtime and space complexity of your implementation
// @Param request body exchange.NumberArray true "required request filters"
// @Produce json
// @Success 200 {object} exchange.IndexWithTimeAndSpaceComplexity
// @failure 400
// @Router /challenge/problem-8 [post]
func Problem8(g *gin.Context) {
	var request exchange.NumberArray
	_ = g.ShouldBindJSON(&request)
	rslt := GetFirstDuplicateIndex(request.Numbers)
	if rslt == -1 {
		helper.BadRequestResponse(g, "no duplications", nil)
		return
	}
	helper.OkResponse(g, "", exchange.IndexWithTimeAndSpaceComplexity{Index: rslt, TimeComplexity: "O(N)", SpaceComplexity: "O(N)"})
}

// Problem9 godoc
// @Summary sum of integers for the whole tree represented
// @Description Given the below tree structure, write a function sum that accepts a node and returns the sum of integers for the whole tree represented by the given node argument
// @Param request body exchange.Node true "required request filters"
// @Produce json
// @Success 200 {object} exchange.Value
// @Router /challenge/problem-9 [post]
func Problem9(g *gin.Context) {
	var request exchange.Node
	_ = g.ShouldBindJSON(&request)
	rslt := GetTreeSum(request)
	helper.OkResponse(g, "", exchange.Value{Value: rslt})
}

package challenge

import (
	"halan/common/exchange"
	"halan/common/helper"
	"regexp"
	"strconv"
	"strings"
)

func isPalindrome(str string) bool {
	str = strings.ToLower(str)
	reversedStr := helper.ReverseString(str)
	if str != reversedStr {
		return false
	}
	return true
}

func GetLinuxCommandLineFileNameChange() string {
	return "mv original.txt changed.txt"
}

func RunLengthEncode(str string) string {
	strOutput := ""
	currentCounter := 0
	currentChar := ""
	for i, c := range str {
		if string(c) == currentChar {
			currentCounter++
		} else {
			if currentCounter != 0 {
				strOutput += currentChar + strconv.Itoa(currentCounter)
			}
			currentChar = string(c)
			currentCounter = 1
		}
		if i == len(str)-1 {
			if currentCounter != 0 {
				strOutput += currentChar + strconv.Itoa(currentCounter)
			}
		}
	}
	return strOutput
}

func RunLengthDecode(str string) string {
	extractNumbers := regexp.MustCompile("[0-9]+").FindAllString(str, -1)
	targetStr := ""
	nextNumberIndex := 0
	for _, ele := range extractNumbers {
		nextNumberIndex = strings.Index(str[nextNumberIndex:], ele) + nextNumberIndex
		intVar, _ := strconv.Atoi(ele)
		targetStr += strings.Repeat(string(str[nextNumberIndex-1]), intVar)
		nextNumberIndex += len(ele)
	}
	return targetStr
}

func CompositionService(value int) int {
	return compose(inc, square)(value)
}

func GetUniqueFromStringList(arr []string) []string {
	arrFrequencies := GetStringsFrequencies(arr)
	return GetStringsWithFrequent(arrFrequencies, 1)
}

func Transpose(arr2D [][]int) [][]int {
	rowsNum, colsNum := len(arr2D[0]), len(arr2D)
	result := Initialize2DIntArray(rowsNum, colsNum)
	for i := 0; i < rowsNum; i++ {
		for j := 0; j < colsNum; j++ {
			result[i][j] = arr2D[j][i]
		}
	}
	return result
}

func ContainersABWithWater() string {
	steps := "1- fill container A with 5 litres from the source. \n"
	steps += "2- pour water from container A to container B  until B is full \n"
	steps += " -  after step 2 the container A contain 2 liters and container B is completely full with 3 liters \n"
	steps += "3- empty the container B \n"
	steps += "4- pour the rest 2 liters of water from container A to container B \n"
	steps += "5- fill container A with 5 litres from the source. \n"
	steps += "6- pour water from container A to container B  until B is full  \n"
	steps += " -  after step 6 the container A contain 4 liters and container B is completely full with 3 liters \n"
	return steps
}

func GetFirstDuplicateIndex(arr []int) int {
	keys := make(map[int]bool)
	for i, ele := range arr {
		if _, value := keys[ele]; value {
			return i
		} else {
			keys[ele] = true
		}

	}
	return -1
}

func GetTreeSum(tree exchange.Node) int {
	if len(tree.Children) == 0 {
		return tree.Value
	}
	sum := tree.Value
	for _, obj := range tree.Children {
		sum += GetTreeSum(obj)
	}
	return sum
}

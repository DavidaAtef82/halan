package challenge

func inc(i int) int {
	return i + 1
}

func square(i int) int {
	return i * i
}

func compose(manyv ...func(int) int) func(int) int {
	return func(i int) int {
		for _, v := range manyv {
			i = v(i)
		}
		return i
	}
}

func GetStringsFrequencies(arr []string) map[string]int {
	numbers := make(map[string]int, 0)
	for _, str := range arr {
		numbers[str]++
	}
	return numbers
}

func GetStringsWithFrequent(stringsFrequencies map[string]int, frequent int) []string {
	list := make([]string, 0)
	for str, strFrequent := range stringsFrequencies {
		if strFrequent == frequent {
			list = append(list, str)
		}
	}
	return list
}

func Initialize2DIntArray(rowsNum, colsNum int) [][]int {
	result := make([][]int, rowsNum)
	for i := range result {
		result[i] = make([]int, colsNum)
	}
	return result
}

package exchange

type Word struct {
	Word string `json:"word"`
}

type WordArray struct {
	Words []string `json:"words"`
}

type Number struct {
	Number int `json:"number"`
}

type Value struct {
	Value int `json:"value"`
}

type IndexWithTimeAndSpaceComplexity struct {
	Index           int    `json:"index"`
	TimeComplexity  string `json:"time_complexity"`
	SpaceComplexity string `json:"space_complexity"`
}

type NumberArray struct {
	Numbers []int `json:"numbers"`
}

type Number2DArray struct {
	Numbers2D [][]int `json:"numbers_2d"`
}

type Node struct {
	Value    int    `json:"value"`
	Children []Node `json:"children"`
}

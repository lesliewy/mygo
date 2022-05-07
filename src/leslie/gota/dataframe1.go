package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

/**
   参考: https://golangrepo.com/repo/go-gota-gota
 */
func main() {
	test1()
	testSubset1()
	testSelect1()
	testFilter1()
}

func test1() {
	fmt.Println("======test1======")
	// series 创建.
	df := dataframe.New(
		series.New([]string{"b", "a"}, series.String, "COL.1"),
		series.New([]int{1, 2}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0}, series.Float, "COL.3"),
	)
	fmt.Println(df)

	// LoadRecords
	df = dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	fmt.Println(df)

	// LoadStructs
	type User struct {
		Name     string
		Age      int
		Accuracy float64
		ignored  bool // ignored since unexported
	}
	users := []User{
		{"Aram", 17, 0.2, true},
		{"Juan", 18, 0.8, true},
		{"Ana", 22, 0.5, true},
	}
	df = dataframe.LoadStructs(users)
	fmt.Println(df)
}

func testSubset1() {
	fmt.Println("=====testSubset1=====")
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	sub := df.Subset([]int{0, 2})
	fmt.Println(sub)
}

func testSelect1() {
	fmt.Println("=====testSelect1=====")
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	sel1 := df.Select([]int{0, 2})
	sel2 := df.Select([]string{"A", "C"})
	fmt.Printf("sel1: %s\n", sel1)
	fmt.Printf("sel2: %s\n", sel2)
}

func testFilter1() {
	fmt.Println("=====testFilter1=====")
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	// Filter()内是OR
	fil := df.Filter(
		dataframe.F{Colname: "A", Comparator: series.Eq, Comparando: "a"},
		dataframe.F{Colname: "B", Comparator: series.Greater, Comparando: 4},
	)
	fmt.Printf("fil: %s\n", fil)

	// AND
	fil2 := df.Filter(dataframe.F{Colname: "A", Comparator: series.Eq, Comparando: "a"}).Filter(dataframe.F{Colname: "B", Comparator: series.Greater, Comparando: 3})
	fmt.Printf("fil2: %s\n", fil2)

	//fil2 := fil.Filter(
	//	dataframe.F{"D", series.Eq, true},
	//)
}

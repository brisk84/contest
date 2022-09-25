package main

import "fmt"

type hotel struct {
	votes int
	star  int
}

func main() {

	hotels := make([]hotel, 26)

	inp := []int{5, 12, 6, 0, 9, 0, 13, 6, 4, 17, 9, 5, 4, 13, 5, 13, 6, 5, 4, 5, 4, 17, 3, 5, 21, 3}
	for k, v := range inp {
		hotels[k].votes += v
	}

	for k, v := range hotels {
		fmt.Println(k, v.votes)
	}

	// in := bufio.NewReader(os.Stdin)

	// f, err := os.Open("/Users/brisk/vscode/contest/con05/in.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// in := bufio.NewReader(f)

	// total := 0
	// fmt.Fscan(in, &total)

}

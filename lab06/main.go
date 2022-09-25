package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var flDebug = false

type PotFriend struct {
	Num    int
	Weight int
}

type User struct {
	Friends    []int
	PotFriends []PotFriend
}

func main() {
	in := bufio.NewReader(os.Stdin)
	n := 0
	m := 0
	s := ""
	if !flDebug {
		fmt.Fscan(in, &n, &m)
	}

	friends := make([][2]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &friends[i][0])
		fmt.Fscan(in, &friends[i][1])
	}

	users := make([]User, n)
	for _, v := range friends {
		users[v[0]-1].Friends = append(users[v[0]-1].Friends, v[1])
		users[v[1]-1].Friends = append(users[v[1]-1].Friends, v[0])
	}

	if !flDebug {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &friends[j][0])
			fmt.Fscan(in, &friends[j][1])
		}
	}

	for k, v := range users {
		for _, vv := range v.Friends {
			for _, vvv := range users[vv-1].Friends {
				if k == vvv-1 {
					continue
				}
				contains := false
				for _, fr := range users[k].Friends {
					if fr == vvv {
						contains = true
						break
					}
				}
				if contains {
					continue
				}
				for kkkk, vvvv := range users[k].PotFriends {
					if vvvv.Num == vvv {
						contains = true
						users[k].PotFriends[kkkk].Weight++
						break
					}
				}
				if !contains {
					pot := PotFriend{
						Num:    vvv,
						Weight: 1,
					}
					users[k].PotFriends = append(users[k].PotFriends, pot)
				}
			}
		}
	}

	for _, v := range users {
		s = ""
		// fmt.Println("user:", k+1, "friends:", v.Friends, "\t\t", "potential:", v.PotFriends)
		maxWeight := 0
		for _, vv := range v.PotFriends {
			if vv.Weight > maxWeight {
				maxWeight = vv.Weight
			}
		}
		printed := false
		sort.Slice(v.PotFriends, func(i, j int) bool {
			return v.PotFriends[i].Num < v.PotFriends[j].Num
		})
		for _, vv := range v.PotFriends {
			if vv.Weight == maxWeight {
				if printed {
					s += fmt.Sprint(" ")
				}
				s += fmt.Sprint(vv.Num)
				printed = true
			}
		}
		if !printed {
			s += fmt.Sprint(0)
		}
		fmt.Println(s)
	}
	// fmt.Println(s)
}

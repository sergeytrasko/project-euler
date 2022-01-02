package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type cell = struct {
	pos     int
	visited int
}

func roll(r *rand.Rand) (int, int) {
	return r.Intn(4) + 1, r.Intn(4) + 1
}

func playMove(field *[]cell, pos *int, communityChest *[]int, chance *[]int, r *rand.Rand) {
	a, b := roll(r)
	*pos = (*pos + a + b) % 40
	if *pos == 30 {
		// go to jail
		*pos = 10
	} else if *pos == 2 || *pos == 17 || *pos == 33 {
		// community chest
		move := (*communityChest)[0]
		*communityChest = (*communityChest)[1:]
		*communityChest = append(*communityChest, move)
		if move == 1 {
			//advance to go
			*pos = 0
		} else if move == 2 {
			// go to jail
			*pos = 10
		}
	} else if *pos == 7 || *pos == 22 || *pos == 36 {
		move := (*chance)[0]
		*chance = (*chance)[1:]
		*chance = append(*chance, move)
		if move == 1 {
			//advance to go
			*pos = 0
		} else if move == 2 {
			// go to jail
			*pos = 10
		} else if move == 3 {
			// go to C1 (11)
			*pos = 11
		} else if move == 4 {
			// go to E3 (24)
			*pos = 24
		} else if move == 5 {
			// go to H2 (39)
			*pos = 39
		} else if move == 6 {
			// go to R1 (5)
			*pos = 5
		} else if move == 7 || move == 8 {
			// go to next R
			if *pos == 7 {
				*pos = 15
			} else if *pos == 22 {
				*pos = 25
			} else if *pos == 36 {
				*pos = 5
			}
		} else if move == 9 {
			// go to next U
			if *pos == 7 {
				*pos = 12
			} else if *pos == 22 || *pos == 36 {
				*pos = 28
			}
		} else if move == 10 {
			*pos = *pos - 3
		}
	}
	(*field)[*pos].visited++
}

func shuffle(x *[]int, r *rand.Rand) {
	sort.Slice(*x, func(i, j int) bool {
		return r.Float64() < 0.5
	})
}

func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	field := []cell{}
	for i := 0; i < 40; i++ {
		field = append(field, cell{
			pos:     i,
			visited: 0,
		})
	}
	communityChest := []int{}
	for i := 0; i < 16; i++ {
		communityChest = append(communityChest, i)
	}
	chance := []int{}
	for i := 0; i < 16; i++ {
		chance = append(chance, i)
	}
	shuffle(&communityChest, r)
	shuffle(&chance, r)
	pos := 0
	for moves := 0; moves < 1_000_000; moves++ {
		playMove(&field, &pos, &communityChest, &chance, r)
	}
	sort.Slice(field, func(i, j int) bool {
		return field[i].visited > field[j].visited
	})
	fmt.Printf("%0.2d%0.2d%0.2d\n", field[0].pos, field[1].pos, field[2].pos)
}

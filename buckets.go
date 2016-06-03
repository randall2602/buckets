package main

import (
	"net/http"
	"strconv"
    "sync"
    "math/rand"
    "time"
)

type color struct {
    red, green, blue uint16
}

var (
    red = color{255,0,0}
    yellow = color{255,255,0}
    green = color{0,255,0}
    cyan = color{0,255,255}
    blue = color{0,0,255}
    magenta = color{255,0,255}
    
    dimRed = color{127,0,0}
    dimYellow = color{127,127,0}
    dimGreen = color{0,127,0}
    dimCyan = color{0,127,127}
    dimBlue = color{0,0,127}
    dimMagenta = color{127,0,127}
)

type bucket struct {
    col int
    row int
    color color
    address int
}

type screen [][]bucket

func addr(col, row int) int {
    return 200 + col + ((row-1) * 10)
}

func NewScreen(cols, rows, start int) screen {
    s := [cols][rows]bucket
    for c := 0; c < cols; c++ {
        for r := 0; r < rows; r++ {
            s[c][r].address = start + (c + 1) + ((r) *rows)
        }
    }
}

func (b *bucket) update(c color) {
	bk := strconv.Itoa(b.address)
	r := strconv.Itoa(c.red)
	g := strconv.Itoa(c.green)
	b := strconv.Itoa(c.blue)
	url := "http://192.168.16." + bk + "/?r=" + r + "&g=" + g + "&b=" + b
    http.Get(url)
}


//////////////////////////////////////////////////////

func matrix(col int) {
    stack := [5]int{0,0,0,0,0}
    state := [5]int{0,0,0,0,0}
    
    for row := range stack {
        stack[row] = addr(col,row+1)
    }
    
    rem := 5 //buckets left to light
    for l := 0; l <= 10; l++ {
        for i, v := range stack {
            if state[i] > 0 {
                state[i] -= 20
            }
            if state[i] == 0 && rem > 0 {
                rem--
                state[i] = 120
                break
            }
            update(v,0,state[i],0)
        }
    }
}

func main() {
    rand.Seed(434545)
	var wg sync.WaitGroup
    for i := 1; i <= 100; i++ {
        wg.Add(1)
		go func(col int) {
			matrix(col)
            defer wg.Done()
		}(rand.Intn(10)+1)
        delay := rand.Intn(20)*50
        time.Sleep(time.Duration(delay) * time.Millisecond)
	}
    wg.Wait()
}

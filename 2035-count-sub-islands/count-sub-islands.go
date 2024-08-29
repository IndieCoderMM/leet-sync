type Coord struct {
    x int
    y int
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    paths := make(map[Coord]bool)
    islands := [][]Coord{}
    var ans int = 0

    var travel func(x, y int, coords *[]Coord) 

    travel = func(x, y int, coords *[]Coord) {
        if y < 0 || x < 0 {
            return 
        }
        if  y >= len(grid2) || x >= len(grid2[y]) {
            return
        }
        if grid2[y][x] == 0 {
            return
        }
        if _, ok := paths[Coord{x, y}]; ok {
            return
        }

        *coords = append(*coords, Coord{x, y})
        paths[Coord{x, y}] = true

        travel(x+1, y, coords)
        travel(x-1, y, coords)
        travel(x, y+1, coords)
        travel(x, y-1, coords)
    }

    for y, row := range grid2 {
        for x, v := range row {
            if v == 0 {
                continue
            }
            if visited := paths[Coord{x, y}]; visited {
                continue
            }
            coords := []Coord{}
            travel(x, y, &coords)
            islands = append(islands, coords)
        }
    }

    for _, coords := range islands {
        isValid := true
       for _, c := range coords {
        if grid1[c.y][c.x] == 0 {
            isValid = false
            break
        }
       } 
       if isValid {
        ans++
       }
    }

    return ans
}

// Grid 1                  Grid 2 
// [1,1,1,1,0,0],           [1,1,1,1,0,1],
// [1,1,0,1,0,0],           [0,0,1,0,1,0]
// [1,0,0,1,1,1],           [1,1,1,1,1,1],    
// [1,1,1,0,0,1],           [0,1,1,1,1,1],
// [1,1,1,1,1,0],           [1,1,1,0,1,0],
// [1,0,1,0,1,0],           [0,1,1,1,1,1],
// [0,1,1,1,0,1],           [1,1,0,1,1,1],
// [1,0,0,0,1,1],           [1,0,0,1,0,1],
// [1,0,0,0,1,0],           [1,1,1,1,1,1],
// [1,1,1,1,1,0]            [1,0,0,1,0,0]

// 1 [{0 0} {1 0} {2 0} {3 0} {2 1} {2 2} {3 2} {4 2} {5 2} {5 3} {4 3} {4 4} {4 5} {5 5} {5 6} {5 7} {5 8} {4 6} {3 3} {2 3} {2 4} {2 5} {3 5} {3 6} {3 7} {3 8} {4 8} {3 9}] 
// 2 [{5 0}] [{4 1}] [{0 2} {1 2} {1 3} {1 4} {1 5} {1 6}] [{0 4}] 
// 3 [{0 6} {0 7} {0 8} {1 8} {2 8} {0 9}]]
package viz

import (
	"fmt"
	"strings"
	"sync"
)

const MaxColorValue = 255

func CanvasToPPM(c Canvas) string {
	res := []string{"P3", fmt.Sprintf("%d %d", c.Width, c.Height), fmt.Sprint(MaxColorValue)}
	rows := make(map[int][]string)
	wg := sync.WaitGroup{}
	wg.Add(c.Height)
	for i := 0; i < c.Height; i++ {
		rowI := i
		go func() {
			row := []string{}
			var curStr string
			for j := 0; j < c.Width; j++ {
				r, g, b := ppmScaledColor(c.Pixel(j, rowI))
				for _, v := range []int{r, g, b} {
					vStr := fmt.Sprint(v)
					if len(curStr)+len(vStr) > 70 {
						row = append(row, strings.TrimSpace(curStr))
						curStr = vStr + " "
					} else {
						curStr += vStr + " "
					}
				}
			}
			row = append(row, strings.TrimSpace(curStr))
			rows[rowI] = row
			wg.Done()
		}()
	}
	wg.Wait()
	for i := 0; i < c.Height; i++ {
		row := rows[i]
		for _, s := range row {
			res = append(res, s)
		}
	}
	return strings.Join(res, "\n") + "\n"
}

func ppmScaledColor(c Color) (int, int, int) {
	return ScaledColorValue(c.R()), ScaledColorValue(c.G()), ScaledColorValue(c.B())
}

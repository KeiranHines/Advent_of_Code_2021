package main

import "strconv"

type oPoint struct {
	x int
	y int
}

type oStats struct {
	energy  int
	flashed bool
}

func SimulateOctopi(input []string, steps int) (map[oPoint]oStats, int) {
	octopi := buildOctpopiMap(input)
	count := 0
	for i := 0; i < steps; i++ {
		count += step(octopi)
		for k, v := range octopi {
			v.flashed = false
			octopi[k] = v
		}
	}
	return octopi, count
}

func FindFirstSyncPoint(input []string) int {
	octopi := buildOctpopiMap(input)
	all := len(octopi)
	count := 1
	for {
		flashed := step(octopi)
		if flashed == all {
			return count
		}
		count++
		for k, v := range octopi {
			v.flashed = false
			octopi[k] = v
		}

	}
}

func buildOctpopiMap(input []string) map[oPoint]oStats {
	mapping := make(map[oPoint]oStats, 0)
	for y, line := range input {
		for x, e := range line {
			energy, _ := strconv.Atoi(string(e))
			o := oPoint{x, y}
			mapping[o] = oStats{energy, false}
		}
	}
	return mapping
}

func step(octopi map[oPoint]oStats) int {
	ready := make(map[oPoint]bool, 0)
	for oct, stats := range octopi {
		stats.energy += 1
		octopi[oct] = stats
		if stats.energy > 9 && !stats.flashed {
			ready[oct] = false
		}
	}
	return flash(octopi, ready)
}

func flash(octopi map[oPoint]oStats, ready map[oPoint]bool) int {
	og := len(ready)
	new := make(map[oPoint]bool, 0)
	for oct := range ready {
		x, _ := octopi[oct]
		if !x.flashed {
			x.flashed = true
			octopi[oct] = x
			adjacent := toTrigger(octopi, oct)
			for _, a := range adjacent {
				stats, ok := octopi[a]
				if ok {
					stats.energy += 1
					octopi[a] = stats
					if stats.energy > 9 && !stats.flashed {
						new[a] = false
					}
				}
			}
		}
	}
	for key := range new {
		ready[key] = false
	}
	if len(ready) == og {
		// No for new ones
		for flashed := range ready {
			s := octopi[flashed]
			if s.energy >= 9 {
				s.energy = 0
			}
			octopi[flashed] = s
		}
		return len(ready)
	}
	return flash(octopi, ready)

}

func toTrigger(octopi map[oPoint]oStats, oct oPoint) []oPoint {
	t := make([]oPoint, 0)
	x := oct.x
	y := oct.y
	t = append(t, oPoint{x - 1, y})
	t = append(t, oPoint{x + 1, y})
	t = append(t, oPoint{x, y - 1})
	t = append(t, oPoint{x, y + 1})
	t = append(t, oPoint{x - 1, y - 1})
	t = append(t, oPoint{x - 1, y + 1})
	t = append(t, oPoint{x + 1, y - 1})
	t = append(t, oPoint{x + 1, y + 1})
	return t
}

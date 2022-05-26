package models

type Location struct {
	height int
	x      int
	y      int
}

func (loc Location) GetHeight() int {
	return loc.height
}

func (loc Location) GetX() int {
	return loc.x
}

func (loc Location) GetY() int {
	return loc.y
}

func (loc Location) GetScore() int {
	return loc.height + 1
}

func (loc Location) IsHigherThanOrEqualTo(i int) bool {
	return loc.height >= i
}

func (loc Location) GetNeighbours(hMap HeightMap) (neighbours []*Location) {
	if loc.x != 0 {
		neighbours = append(neighbours, &hMap.Locations[loc.y][loc.x-1])
	}
	if loc.x != len(hMap.Locations[loc.y])-1 {
		neighbours = append(neighbours, &hMap.Locations[loc.y][loc.x+1])
	}
	if loc.y != 0 {
		neighbours = append(neighbours, &hMap.Locations[loc.y-1][loc.x])
	}
	if loc.y != len(hMap.Locations)-1 {
		neighbours = append(neighbours, &hMap.Locations[loc.y+1][loc.x])
	}
	return
}

func MakeLocation(height int, x int, y int) Location {
	return Location{height, x, y}
}

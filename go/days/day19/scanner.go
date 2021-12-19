package day19

type aocScanner struct {
	id              int
	originalBeacons map[string][3]int
	vectors         [24]map[string][3]int
}

type aocCoordinates struct {
	x int
	y int
	z int
}

func generateSwaps(c aocCoordinates) []aocCoordinates {
	swaps := make([]aocCoordinates, 0, 24)
	X, Y, Z := c.x, c.y, c.z
	swaps = append(swaps, aocCoordinates{x: X, y: Y, z: Z})
	swaps = append(swaps, aocCoordinates{x: -Z, y: Y, z: X})
	swaps = append(swaps, aocCoordinates{x: -X, y: Y, z: -Z})
	swaps = append(swaps, aocCoordinates{x: Z, y: Y, z: -X})
	swaps = append(swaps, aocCoordinates{x: X, y: -Z, z: Y})
	swaps = append(swaps, aocCoordinates{x: X, y: -Y, z: -Z})
	swaps = append(swaps, aocCoordinates{x: X, y: Z, z: -Y})
	swaps = append(swaps, aocCoordinates{x: -Y, y: X, z: Z})
	swaps = append(swaps, aocCoordinates{x: -X, y: -Y, z: Z})
	swaps = append(swaps, aocCoordinates{x: Y, y: -X, z: Z})

	swaps = append(swaps, aocCoordinates{x: Z, y: X, z: Y})

	swaps = append(swaps, aocCoordinates{x: Z, y: -Y, z: X})
	swaps = append(swaps, aocCoordinates{x: Z, y: -X, z: -Y})
	swaps = append(swaps, aocCoordinates{x: -Z, y: X, z: -Y})
	swaps = append(swaps, aocCoordinates{x: Y, y: X, z: -Z})
	swaps = append(swaps, aocCoordinates{x: -X, y: Z, z: Y})
	swaps = append(swaps, aocCoordinates{x: -Z, y: -X, z: Y})

	swaps = append(swaps, aocCoordinates{x: Y, y: Z, z: X})

	swaps = append(swaps, aocCoordinates{x: Y, y: -X, z: Z})
	swaps = append(swaps, aocCoordinates{x: Y, y: -Z, z: -X})
	swaps = append(swaps, aocCoordinates{x: -Z, y: Y, z: X})
	swaps = append(swaps, aocCoordinates{x: -Y, y: -Z, z: X})
	swaps = append(swaps, aocCoordinates{x: -Y, y: Z, z: -X})
	swaps = append(swaps, aocCoordinates{x: X, y: Z, z: -Y})

	return swaps
}

func (scanner *aocScanner) populateVectors() {

}

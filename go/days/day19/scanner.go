package day19

type aocScanner struct {
	id               int
	originalBeacons  map[string]aocCoordinates
	beaconsPerSystem [24]map[string]aocCoordinates
	vectors          map[string]*[24]*[]aocVector
}

type aocCoordinates struct {
	x int
	y int
	z int
}

type aocVector struct {
	x int
	y int
	z int
}

func (scanner *aocScanner) populateSystems() {
	for system := 0; system < 24; system++ {
		beacons := make(map[string]aocCoordinates)
		for key, c := range scanner.originalBeacons {
			beacons[key] = getCoordinatesInSystem(c, system)
		}
		scanner.beaconsPerSystem[system] = beacons
	}
}

func (scanner *aocScanner) populateVectors() {
	for refKey := range scanner.originalBeacons {
		scanner.vectors[refKey] = &[24]*[]aocVector{}
	}
	for refKey, refC := range scanner.originalBeacons {
		for system := 0; system < 24; system++ {
			vectors := make([]aocVector, 0, len(scanner.originalBeacons))
			for targetKey, targetC := range scanner.originalBeacons {
				if refKey != targetKey {
					vectors = append(vectors, aocVector{
						x: targetC.x - refC.x,
						y: targetC.y - refC.y,
						z: targetC.z - refC.z,
					})
				}
			}
			myArray := scanner.vectors[refKey]
			myArray[system] = &vectors
		}
	}
}

func getCoordinatesInSystem(c aocCoordinates, system int) aocCoordinates {
	X, Y, Z := c.x, c.y, c.z
	switch system {
	case 0:
		return aocCoordinates{x: X, y: Y, z: Z}
	case 1:
		return aocCoordinates{x: -Z, y: Y, z: X}
	case 2:
		return aocCoordinates{x: -X, y: Y, z: -Z}
	case 3:
		return aocCoordinates{x: Z, y: Y, z: -X}
	case 4:
		return aocCoordinates{x: X, y: -Z, z: Y}
	case 5:
		return aocCoordinates{x: X, y: -Y, z: -Z}
	case 6:
		return aocCoordinates{x: X, y: Z, z: -Y}
	case 7:
		return aocCoordinates{x: -Y, y: X, z: Z}
	case 8:
		return aocCoordinates{x: -X, y: -Y, z: Z}
	case 9:
		return aocCoordinates{x: Y, y: -X, z: Z}
	case 10:
		return aocCoordinates{x: Z, y: X, z: Y}
	case 11:
		return aocCoordinates{x: Z, y: -Y, z: X}
	case 12:
		return aocCoordinates{x: Z, y: -X, z: -Y}
	case 13:
		return aocCoordinates{x: -Z, y: X, z: -Y}
	case 14:
		return aocCoordinates{x: Y, y: X, z: -Z}
	case 15:
		return aocCoordinates{x: -X, y: Z, z: Y}
	case 16:
		return aocCoordinates{x: -Z, y: -X, z: Y}
	case 17:
		return aocCoordinates{x: Y, y: Z, z: X}
	case 18:
		return aocCoordinates{x: Y, y: -X, z: Z}
	case 19:
		return aocCoordinates{x: Y, y: -Z, z: -X}
	case 20:
		return aocCoordinates{x: -Z, y: Y, z: X}
	case 21:
		return aocCoordinates{x: -Y, y: -Z, z: X}
	case 22:
		return aocCoordinates{x: -Y, y: Z, z: -X}
	case 23:
		return aocCoordinates{x: X, y: Z, z: -Y}
	default:
		return aocCoordinates{x: -9999999, y: -9999999, z: -99999999}
	}
}

package day19

import "fmt"

type aocScanner struct {
	id               int
	originalBeacons  map[string]aocCoordinates
	beaconsPerSystem [24]map[string]aocCoordinates
	vectors          map[string]*[24]map[string]aocVector
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

func (scanner *aocScanner) hasEnoughCommonPoints(targetScanner *aocScanner) bool {
	for system := 0; system < 24; system++ {
		for beaconKey := range scanner.originalBeacons {
			for targetBeaconKey := range targetScanner.originalBeacons {
				if targetBeaconKey == "s2b0" && beaconKey == "s0b0" && system == 23 {
					fmt.Printf("toto")
				}
				numCommon := numCommonVectors(
					scanner.vectors[beaconKey][system],
					targetScanner.vectors[targetBeaconKey][0])
				if numCommon >= 5 {
					return true
				}
			}
		}
	}
	return false

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

func (scanner *aocScanner) print(system int) {
	fmt.Printf("##########################\n")
	for key, c := range scanner.beaconsPerSystem[system] {
		fmt.Printf("System #%v %v %v (System #0: %v)\n", system, key, c, scanner.beaconsPerSystem[0][key])
	}
	fmt.Printf("##########################\n")
}

func (scanner *aocScanner) populateVectors() {
	for refKey := range scanner.originalBeacons {
		systemArray := [24]map[string]aocVector{}
		for system := 0; system < 24; system++ {
			vectorMap := make(map[string]aocVector)
			systemArray[system] = vectorMap
		}
		scanner.vectors[refKey] = &systemArray
	}
	for refKey := range scanner.originalBeacons {
		for system := 0; system < 24; system++ {
			for targetKey := range scanner.originalBeacons {
				if refKey != targetKey {
					refCInSystem := scanner.beaconsPerSystem[system][refKey]
					targetCInSystem := scanner.beaconsPerSystem[system][targetKey]
					scanner.vectors[refKey][system][targetKey] = aocVector{
						x: targetCInSystem.x - refCInSystem.x,
						y: targetCInSystem.y - refCInSystem.y,
						z: targetCInSystem.z - refCInSystem.z,
					}
				}
			}
		}
	}
}

func getCoordinatesInSystem(c aocCoordinates, system int) aocCoordinates {
	X, Y, Z := c.x, c.y, c.z
	switch system {
	case 0:
		return aocCoordinates{x: X, y: Y, z: Z}
	case 1:
		return aocCoordinates{x: X, y: -Z, z: Y}
	case 2:
		return aocCoordinates{x: X, y: -Y, z: -Z}
	case 3:
		return aocCoordinates{x: X, y: Z, z: -Y}
	case 4:
		return aocCoordinates{x: -X, y: -Y, z: Z}
	case 5:
		return aocCoordinates{x: -X, y: -Z, z: -Y}
	case 6:
		return aocCoordinates{x: -X, y: Y, z: -Z}
	case 7:
		return aocCoordinates{x: -X, y: Z, z: Y}
	case 8:
		return aocCoordinates{x: Y, y: Z, z: X}
	case 9:
		return aocCoordinates{x: Y, y: -X, z: Z}
	case 10:
		return aocCoordinates{x: Y, y: -Z, z: -X}
	case 11:
		return aocCoordinates{x: Y, y: X, z: -Z}
	case 12:
		return aocCoordinates{x: -Y, y: -Z, z: X}
	case 13:
		return aocCoordinates{x: -Y, y: -X, z: -Z}
	case 14:
		return aocCoordinates{x: -Y, y: Z, z: -X}
	case 15:
		return aocCoordinates{x: -Y, y: X, z: Z}
	case 16:
		return aocCoordinates{x: Z, y: X, z: Y}
	case 17:
		return aocCoordinates{x: Z, y: -Y, z: X}
	case 18:
		return aocCoordinates{x: Z, y: -X, z: -Y}
	case 19:
		return aocCoordinates{x: Z, y: Y, z: -X}
	case 20:
		return aocCoordinates{x: -Z, y: -X, z: Y}
	case 21:
		return aocCoordinates{x: -Z, y: -Y, z: -X}
	case 22:
		return aocCoordinates{x: -Z, y: X, z: -Y}
	case 23:
		return aocCoordinates{x: -Z, y: Y, z: X}
	default:
		return aocCoordinates{x: -9999999, y: -9999999, z: -99999999}
	}
}

package day16

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/utils"
)

var charToHex = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

var hexToChar = map[string]rune{
	"0000": '0',
	"0001": '1',
	"0010": '2',
	"0011": '3',
	"0100": '4',
	"0101": '5',
	"0110": '6',
	"0111": '7',
	"1000": '8',
	"1001": '9',
	"1010": 'A',
	"1011": 'B',
	"1100": 'C',
	"1101": 'D',
	"1110": 'E',
	"1111": 'F',
}

type bitsPacket struct {
	// General
	sequence     string
	version      int
	typeId       int
	indexPayload int
	totalLength  int

	// Literal
	literal int

	// Composite
	lengthType int
	subPackets []*bitsPacket
}

func createBitsPacket(sequence string) *bitsPacket {
	packet := bitsPacket{
		sequence: sequence,
	}
	packet.parse()
	return &packet
}

func (packet *bitsPacket) parse() {
	packet.version = utils.RuneToInt(hexToChar["0"+packet.sequence[:3]])
	packet.typeId = utils.RuneToInt(hexToChar["0"+packet.sequence[3:6]])

	switch packet.typeId {
	case 4:
		packet.indexPayload = 6
		literal, length := getLiteral(packet.sequence[6:])
		packet.literal = literal
		packet.totalLength = 6 + length

	default:
		packet.totalLength = 0
		packet.lengthType = utils.RuneToInt(rune(packet.sequence[6]))
		if packet.lengthType == 0 {
			length := utils.BitsToInt(packet.sequence[7 : 7+15])
			packet.indexPayload = 22

			offset := 0
			for {
				if offset > length-1 {
					break
				}
				subPacket := createBitsPacket(packet.sequence[packet.indexPayload+offset:])
				packet.subPackets = append(packet.subPackets, subPacket)
				offset += subPacket.totalLength
			}

			packet.totalLength += packet.indexPayload + length

		} else {
			numSubPackets := utils.BitsToInt(packet.sequence[7 : 7+11])
			packet.indexPayload = 18
			offset := 0
			for numProcessed := 0; numProcessed < numSubPackets; numProcessed++ {
				subPacket := createBitsPacket(packet.sequence[packet.indexPayload+offset:])
				packet.subPackets = append(packet.subPackets, subPacket)
				offset += subPacket.totalLength
			}
			packet.totalLength += packet.indexPayload + offset
		}
	}
}

func (packet *bitsPacket) sumVersions() int {
	sumSubVersions := 0
	for _, subPacket := range packet.subPackets {
		sumSubVersions += subPacket.sumVersions()
	}
	return packet.version + sumSubVersions
}

func (packet *bitsPacket) evaluate() int {
	result := 0
	values := make([]int, 0)
	for _, subPacket := range packet.subPackets {
		values = append(values, subPacket.evaluate())
	}

	switch packet.typeId {
	case 0:
		result = 0
		for _, value := range values {
			result = result + value
		}
	case 1:
		result = 1
		for _, value := range values {
			result = result * value
		}
	case 2:
		result = utils.MinSlice(values)
	case 3:
		result = utils.MaxSlice(values)
	case 4:
		result = packet.literal
	case 5:
		if values[0] > values[1] {
			result = 1
		} else {
			result = 0
		}
	case 6:
		if values[0] < values[1] {
			result = 1
		} else {
			result = 0
		}
	case 7:
		if values[0] == values[1] {
			result = 1
		} else {
			result = 0
		}
	default:
		fmt.Printf("Error!!!!! %v\n", packet.typeId)
	}
	return result
}

func getLiteral(bits string) (int, int) {
	literalAsString := ""
	length := 0
	for i := 0; true; i += 5 {
		literalAsString = literalAsString + bits[i+1:i+5]
		length += 5
		if rune(bits[i]) == '0' {
			break
		}
	}
	return utils.BitsToInt(literalAsString), length
}

func doPart1(packet *bitsPacket) int {
	return packet.sumVersions()
}

func doPart2(packet *bitsPacket) int {
	return packet.evaluate()
}

func getSequence(text []rune) string {
	sequence := ""
	for _, letter := range text {
		sequence = sequence + charToHex[letter]
	}
	return sequence
}

func Run(path string) {
	sequence := getSequence(utils.ReadFileAsSliceOfRuneSlices(path)[0])
	packet := createBitsPacket(sequence)
	answer1 := doPart1(packet)
	answer2 := doPart2(packet)
	fmt.Printf("Part 1 answer: %v\n", answer1)
	fmt.Printf("Part 2 answer: %v\n", answer2)
}

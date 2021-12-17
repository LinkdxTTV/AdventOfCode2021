package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()

	binStr := convertHexToBin(line)

	packets := []packetLiteral{}

	for binStr != "" {
		if len(binStr) < 10 {
			if binToDec(binStr) == 0 {
				break
			}
		}
		var packet packetLiteral
		packet, binStr = parsePacket(binStr)
		packets = append(packets, packet)
	}

	versionSum := 0
	recursivePacketValue(packets, &versionSum)
	fmt.Println("Part 1:", versionSum)

	// Part 2
	sum := 0
	for _, packet := range packets {
		sum += returnPacketValue(packet)
	}
	fmt.Println("Part 2:", sum)
}

func recursivePacketValue(packets []packetLiteral, version *int) {
	for _, packet := range packets {
		*version += packet.version
		recursivePacketValue(packet.subpackets, version)
	}
}

func returnPacketValue(packet packetLiteral) int {

	switch packet.typeID {
	case 0:
		sum := 0
		for _, subpacket := range packet.subpackets {
			sum += returnPacketValue(subpacket)
		}
		return sum
	case 1:
		product := 1
		for _, subpacket := range packet.subpackets {
			product = product * returnPacketValue(subpacket)
		}
		return product
	case 2:
		min := math.MaxInt64
		for _, subpacket := range packet.subpackets {
			value := returnPacketValue(subpacket)
			if value < min {
				min = value
			}
		}
		return min
	case 3:
		max := 0
		for _, subpacket := range packet.subpackets {
			value := returnPacketValue(subpacket)
			if value > max {
				max = value
			}
		}
		return max
	case 4:
		return *packet.value
	case 5:
		if returnPacketValue(packet.subpackets[0]) > returnPacketValue(packet.subpackets[1]) {
			return 1
		}
		return 0
	case 6:
		if returnPacketValue(packet.subpackets[0]) < returnPacketValue(packet.subpackets[1]) {
			return 1
		}
		return 0
	case 7:
		if returnPacketValue(packet.subpackets[0]) == returnPacketValue(packet.subpackets[1]) {
			return 1
		}
		return 0
	default:
		panic("what the hell")
	}
}

type packetLiteral struct {
	version    int
	typeID     int
	value      *int
	subpackets []packetLiteral
}

func parsePacket(binStr string) (packetLiteral, string) { // Returns a parsed packet and leftovers
	version, pType, temp := checkPacketHeader(binStr)
	binStr = temp
	switch pType {
	case 4: // Literal
		var packet packetLiteral
		packet, binStr = parseLiteralPacket(binStr)
		packet.version = version
		packet.typeID = pType
		return packet, binStr
	default: // Operators
		var packet packetLiteral
		packet, binStr = parseOperatorPacket(binStr)
		packet.version = version
		packet.typeID = pType
		return packet, binStr
	}
}
func checkPacketHeader(binStr string) (int, int, string) {
	version, binStr := binToDec(binStr[:3]), binStr[3:]
	pType, binStr := binToDec(binStr[:3]), binStr[3:]
	return version, pType, binStr
}

func parseLiteralPacket(binStr string) (packetLiteral, string) {
	stop := false
	outBin := ""
	for {
		first := binStr[0:1]
		binStr = binStr[1:]
		if first == "0" {
			stop = true
		}
		next4 := binStr[:4]
		binStr = binStr[4:]
		outBin += next4
		if stop {
			break
		}
	}
	value := binToDec(outBin)
	return packetLiteral{42, 42, &value, []packetLiteral{}}, binStr
}

func parseOperatorPacket(binStr string) (packetLiteral, string) {
	lengthType := binStr[0:1]
	binStr = binStr[1:]
	switch lengthType {
	case "0":
		length := binToDec(binStr[0:15])
		binStr = binStr[15:]
		outPacket := packetLiteral{42, 42, nil, []packetLiteral{}}
		subPackets := binStr[0:length]
		binStr = binStr[length:]
		for subPackets != "" {
			var packet packetLiteral
			packet, subPackets = parsePacket(subPackets)
			outPacket.subpackets = append(outPacket.subpackets, packet)
		}
		return outPacket, binStr
	case "1":
		numSubPackets := binToDec(binStr[0:11])
		binStr = binStr[11:]
		outPacket := packetLiteral{42, 42, nil, []packetLiteral{}}
		for i := 0; i < numSubPackets; i++ {
			var packet packetLiteral
			packet, binStr = parsePacket(binStr)
			outPacket.subpackets = append(outPacket.subpackets, packet)
		}
		return outPacket, binStr

	default:
		log.Fatal("we messed up (parse operator packet lengthType)")
		return packetLiteral{}, ""
	}
}

func convertHexToBin(in string) string {
	out := ""
	for _, char := range in {
		bin, err := strconv.ParseUint(string(char), 16, 4)
		if err != nil {
			log.Fatal(err)
		}
		out += fmt.Sprintf("%04b", bin)
	}
	return out
}

func binToDec(in string) int {
	num, err := strconv.ParseInt(in, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}

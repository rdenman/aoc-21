package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rdenman/aoc-21/util"
)

func main() {
	sol1 := solution1()
	sol2 := solution2()

	fmt.Printf("Solution #1: %d\n", sol1)
	fmt.Printf("Solution #2: %d\n", sol2)
}

func solution1() int {
	binary := readInputData()

	packet, _ := newPacket(binary)
	return packet.versionSum()
}

func solution2() int {
	binary := readInputData()

	packet, _ := newPacket(binary)
	return packet.evaluateExpression()
}

func readInputData() Binary {
	input, err := os.ReadFile("./input.txt")
	util.CheckError(err)

	return hexToBinary(string(input))
}

func hexToBinary(hexStr string) (binary Binary) {
	hex, err := hex.DecodeString(hexStr)
	util.CheckError(err)
	for _, bit := range hex {
		binary += Binary(fmt.Sprintf("%08b", int(bit)))
	}
	return binary
}

func binaryToInt(binary string) int {
	val, err := strconv.ParseInt(string(binary), 2, 64)
	util.CheckError(err)
	return int(val)
}

type Binary string
func (binary Binary) splitAt(i int) (Binary, Binary) {
	return binary[:i], binary[i:]
}

type TypeId int
const (
	TYPE_SUM TypeId = iota
	TYPE_PRODUCT
	TYPE_MIN
	TYPE_MAX
	TYPE_LITERAL
	TYPE_GT
	TYPE_LT
	TYPE_EQ
)

type LengthTypeId int
const (
	LENGTH_BITS LengthTypeId = iota
	LENGTH_PACKETS
)


type LiteralValuePacket struct {
	value int
}

type OperatorPacket struct {
	lengthTypeId LengthTypeId
	packets      []Packet
}

type Packet struct {
	version int
	typeId  TypeId
	LiteralValuePacket
	OperatorPacket
}

func newPacket(binary Binary) (packet *Packet, tail Binary) {
	head, tail := binary.splitAt(3)
	version := binaryToInt(string(head))

	head, tail = tail.splitAt(3)
	typeId := binaryToInt(string(head))

	packet = &Packet{
		version: version,
		typeId:  TypeId(typeId),
	}

	if packet.typeId == TYPE_LITERAL {
		tail = packet.parseLiteralValue(tail)
	} else {
		tail = packet.parseOperator(tail)
	}

	return packet, tail
}

func (packet *Packet) parseLiteralValue(binary Binary) (tail Binary) {
	chunkSize := 5
	var valueBin Binary
	head, tail := binary.splitAt(chunkSize)

	for len(head) == 5 {
		first, rest := head.splitAt(1)
		valueBin += rest

		if first == "0" {
			break
		}

		head, tail = tail.splitAt(chunkSize)
	}

	value := binaryToInt(string(valueBin))
	packet.value = value
	return tail
}

func (packet *Packet) parseOperator(binary Binary) (tail Binary) {
	head, tail := binary.splitAt(1)
	if head == "0" {
		packet.lengthTypeId = LENGTH_BITS
		tail = packet.parseOperatorBitCount(tail)
	} else {
		packet.lengthTypeId = LENGTH_PACKETS
		tail = packet.parseOperatorPacketCount(tail)
	}

	return tail
}

func (packet *Packet) parseOperatorBitCount(binary Binary) (tail Binary) {
	head, tail := binary.splitAt(15)
	bitCount := binaryToInt(string(head))

	subpackets, tail := tail.splitAt(bitCount)
	for subpackets != Binary(strings.Repeat("0", len(subpackets))) {
		subpacket, nextPackets := newPacket(subpackets)
		packet.packets = append(packet.packets, *subpacket)
		subpackets = nextPackets
	}

	return tail
}

func (packet *Packet) parseOperatorPacketCount(binary Binary) (tail Binary) {
	head, tail := binary.splitAt(11)
	packetCount := binaryToInt(string(head))

	for packetCount > 0 {
		packetCount--
		subpacket, nextTail := newPacket(tail)
		packet.packets = append(packet.packets, *subpacket)
		tail = nextTail
	}

	return tail
}

func (packet *Packet) versionSum() (sum int) {
	sum += packet.version
	for _, subpacket := range packet.packets {
		sum += subpacket.versionSum()
	}
	return
}

func (packet *Packet) evaluateExpression() int {
	if packet.typeId == TYPE_LITERAL {
		return packet.value
	}

	values := []int{}
	for _, subpacket := range packet.packets {
		values = append(values, subpacket.evaluateExpression())
	}

	switch packet.typeId {
	case TYPE_SUM:
		return sum(values...)
	case TYPE_PRODUCT:
		product := 1
		for _, val := range values {
			product *= val
		}
		return product
	case TYPE_MIN:
		return min(values...)
	case TYPE_MAX:
		return max(values...)
	case TYPE_GT:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case TYPE_LT:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case TYPE_EQ:
		if values[0] == values[1] {
			return 1
		}
		return 0
	}
	return 0
}

func sum(nums ...int) (s int) {
	for _, val := range nums {
		s += val
	}
	return s
}

func min(nums ...int) int {
	m := nums[0]
	for _, val := range nums {
		if val < m {
			m = val
		}
	}
	return m
}

func max(nums ...int) int {
	m := nums[0]
	for _, val := range nums {
		if val > m {
			m = val
		}
	}
	return m
}

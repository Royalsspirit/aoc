package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func packet(packet string) {
	version := packet[0:3]
	typeId := packet[3:6]

	bits := packet[6]
	var subPacket string
	if bits == 1 {

		subPacket = packet[7:18]
	} else {

		subPacket = packet[7:22]
	}
	fmt.Println("version id", version, "typeId", typeId, "bits", bits, "subpacket", subPacket)

}
func subPacket(subPacket string) {
	version := subPacket[0:3]
	typeId := subPacket[3:6]
	markBit := subPacket[6]
	/**
	* if marketBit == 1 then we know that next 4 bits are not the last
	* otherwise we know that is the last 4 bits then start to parse new packet
	**/

}
func part1(bits []string) {
	packet := bits[len(bits)-1]
	version := packet[0:3]
	typeId := packet[3:6]
	markBit := packet[6]
	value, _ := strconv.Atoi("1")
	fmt.Println("part1", value)
	fmt.Println("bits", bits)
}
func part2() {
}
func main() {
	f, err := os.Open("input-ex.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var bits []string
	for scanner.Scan() {

		currentValue := scanner.Text()
		var binary []string
		for _, v := range currentValue {
			r, _ := strconv.ParseUint(string(v), 16, 64)
			binary = append(binary, fmt.Sprintf("%04b", r))
		}
		binaryString := strings.Join(binary, "")
		bits = append(bits, binaryString[0:3])
		bits = append(bits, binaryString[3:6])

		bits = append(bits, string(binaryString[6]))
		if binaryString[6] == 1 {

			bits = append(bits, string(binaryString[7:18]))
		} else {

			bits = append(bits, string(binaryString[7:22]))
		}
		bits = append(bits, binaryString[22:])

	}
	if err != nil {
		fmt.Println("error", err)
	}
	part1(bits)
	//	part2(template, pairs)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}

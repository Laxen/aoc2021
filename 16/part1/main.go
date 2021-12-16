package main

import "fmt"

type packet struct {
	version, id  int
	number       int
	lengthTypeId int
	subpackages  []*packet
}

func (p packet) toSingleString() string {
	s := ""
	if p.id == 4 {
		s = fmt.Sprintf("Literal packet v%d containing %d\n", p.version, p.number)
	} else {
		s = fmt.Sprintf("Operator packet v%d\n", p.version)
	}
	return s
}

func (p packet) toString(indent int) string {
	s := ""
	prefix := ""
	for i := 0; i < indent; i++ {
		prefix += " "
	}

	if p.id == 4 {
		s = fmt.Sprintf("%sLiteral packet v%d containing %d\n", prefix, p.version, p.number)
	} else {
		s = fmt.Sprintf("%sOperator packet v%d\n", prefix, p.version)
		for _, sp := range p.subpackages {
			s += sp.toString(indent + 2)
		}
	}

	return s
}

func hexToBin(hex string) string {
	ret := ""

	for _, d := range hex {
		switch string(d) {
		case "0":
			ret += "0000"
		case "1":
			ret += "0001"
		case "2":
			ret += "0010"
		case "3":
			ret += "0011"
		case "4":
			ret += "0100"
		case "5":
			ret += "0101"
		case "6":
			ret += "0110"
		case "7":
			ret += "0111"
		case "8":
			ret += "1000"
		case "9":
			ret += "1001"
		case "A":
			ret += "1010"
		case "B":
			ret += "1011"
		case "C":
			ret += "1100"
		case "D":
			ret += "1101"
		case "E":
			ret += "1110"
		case "F":
			ret += "1111"
		}
	}

	return ret
}

func pow(a int, b int) int {
	ret := 1
	for i := 0; i < b; i++ {
		ret *= a
	}
	return ret
}

func binToDec(bin string) int {
	ret := 0
	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == '1' {
			ret += pow(2, len(bin)-1-i)
		}
	}
	return ret
}

func decodeLiteralPacket(bin string) (string, packet) {
	p := packet{}
	p.version = binToDec(bin[:3])
	p.id = binToDec(bin[3:6])

	i := 6
	number := ""
	for true {
		group := bin[i : i+5]
		number += group[1:5]
		i += 5

		if group[0] == '0' {
			break
		}
	}

	p.number = binToDec(number)
	return bin[i:], p
}

func decodeOperatorPacket(bin string) (string, packet) {
	p := packet{}
	p.version = binToDec(bin[:3])
	p.id = binToDec(bin[3:6])
	p.lengthTypeId = binToDec(string(bin[6]))

	if p.lengthTypeId == 0 {
		bitLength := binToDec(bin[7:22])
		subPkts := bin[22 : 22+bitLength]
		pkts := decodePackets(subPkts)
		p.subpackages = pkts

		return bin[22+bitLength:], p
	}

	// id == 1
	nSubPkts := binToDec(bin[7:18])
	bin = bin[18:]
	for i := 0; i < nSubPkts; i++ {
		newBin, pkt := decodePacket(bin)
		bin = newBin
		p.subpackages = append(p.subpackages, &pkt)
	}

	return bin, p
}

func hasData(bin string) bool {
	for _, b := range bin {
		if b == '1' {
			return true
		}
	}
	return false
}

func decodePacket(bin string) (string, packet) {
	id := binToDec(bin[3:6])

	if id == 4 {
		newBin, p := decodeLiteralPacket(bin)

		return newBin, p
	}

	// id != 4
	newBin, p := decodeOperatorPacket(bin)
	return newBin, p
}

func decodePackets(bin string) []*packet {
	// version := binToDec(bin[:3])

	packets := []*packet{}
	for hasData(bin) {
		newBin, p := decodePacket(bin)
		bin = newBin
		packets = append(packets, &p)
	}

	return packets
}

func getPacketsInHierarchy(p packet) []*packet {
	packets := []*packet{}
	for _, p := range p.subpackages {
		pkts := getPacketsInHierarchy(*p)
		packets = append(packets, pkts...)
	}
	return packets
}

func getAllPackets(pktList []*packet) []*packet {
	ret := pktList

	for _, pkt := range pktList {
		subpkts := getAllPackets(pkt.subpackages)
		ret = append(ret, subpkts...)
	}

	return ret
}

func main() {
	// literal := "D2FE28"
	// operator := "38006F45291200"
	// operator := "EE00D40C823060"
	// operator := "8A004A801A8002F478"
	// operator := "620080001611562C8802118E34"
	// operator := "C0015000016115A2E0802F182340"
	// operator := "A0016C880162017C3686B18A3D4780"
	input := "E0529D18025800ABCA6996534CB22E4C00FB48E233BAEC947A8AA010CE1249DB51A02CC7DB67EF33D4002AE6ACDC40101CF0449AE4D9E4C071802D400F84BD21CAF3C8F2C35295EF3E0A600848F77893360066C200F476841040401C88908A19B001FD35CCF0B40012992AC81E3B980553659366736653A931018027C87332011E2771FFC3CEEC0630A80126007B0152E2005280186004101060C03C0200DA66006B8018200538012C01F3300660401433801A6007380132DD993100A4DC01AB0803B1FE2343500042E24C338B33F5852C3E002749803B0422EC782004221A41A8CE600EC2F8F11FD0037196CF19A67AA926892D2C643675A0C013C00CC0401F82F1BA168803510E3942E969C389C40193CFD27C32E005F271CE4B95906C151003A7BD229300362D1802727056C00556769101921F200AC74015960E97EC3F2D03C2430046C0119A3E9A3F95FD3AFE40132CEC52F4017995D9993A90060729EFCA52D3168021223F2236600ECC874E10CC1F9802F3A71C00964EC46E6580402291FE59E0FCF2B4EC31C9C7A6860094B2C4D2E880592F1AD7782992D204A82C954EA5A52E8030064D02A6C1E4EA852FE83D49CB4AE4020CD80272D3B4AA552D3B4AA5B356F77BF1630056C0119FF16C5192901CEDFB77A200E9E65EAC01693C0BCA76FEBE73487CC64DEC804659274A00CDC401F8B51CE3F8803B05217C2E40041A72E2516A663F119AC72250A00F44A98893C453005E57415A00BCD5F1DD66F3448D2600AC66F005246500C9194039C01986B317CDB10890C94BF68E6DF950C0802B09496E8A3600BCB15CA44425279539B089EB7774DDA33642012DA6B1E15B005C0010C8C917A2B880391160944D30074401D845172180803D1AA3045F00042630C5B866200CC2A9A5091C43BBD964D7F5D8914B46F040"

	bin := hexToBin(input)

	packets := decodePackets(bin)
	for _, packet := range packets {
		fmt.Println(packet.toString(0))
	}
	allPkts := getAllPackets(packets)
	version := 0
	for _, pkt := range allPkts {
		// fmt.Println(pkt.toSingleString())
		version += pkt.version
	}
	fmt.Println(version)
}

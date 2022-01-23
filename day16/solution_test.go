package day16

import (
	"strings"
	"testing"
)

func TestBitReader(t *testing.T) {
	reader := NewPacketReader(strings.NewReader("D2FE28"))
	got, err := reader.readBits(3)
	if err != nil {
		t.Error(err)
	}
	if got != 6 {
		t.Errorf("expected 6, got %d", got)
	}
	got, err = reader.readBits(3)
	if err != nil {
		t.Error(err)
	}
	if got != 4 {
		t.Errorf("expected 4, got %d", got)
	}
}

func TestLiteralPacket(t *testing.T) {
	reader := NewPacketReader(strings.NewReader("D2FE28"))
	got, err := reader.ReadPacket()
	if err != nil {
		t.Error(err)
	}
	if got.version != 6 || got.typeID != 4 || got.literal != 2021 {
		t.Errorf("bad packet: %v", got)
	}
}

func TestBitSubPacket(t *testing.T) {
	reader := NewPacketReader(strings.NewReader("38006F45291200"))
	packet, err := reader.ReadPacket()
	if err != nil {
		t.Error(err)
	}
	if packet.version != 1 || packet.typeID != 6 {
		t.Errorf("unexpected: %v", packet)
	}
	if len(packet.subPackets) != 2 {
		t.Errorf("expected 2 subpackets")
	}
	if packet.subPackets[0].literal != 10 || packet.subPackets[1].literal != 20 {
		t.Errorf("subpackets borked: %v", packet.subPackets)
	}
}

func TestCountSubPacket(t *testing.T) {
	reader := NewPacketReader(strings.NewReader("EE00D40C823060"))
	packet, err := reader.ReadPacket()
	if err != nil {
		t.Error(err)
	}
	if packet.version != 7 || packet.typeID != 3 {
		t.Errorf("unexpected: %v", packet)
	}
	if len(packet.subPackets) != 3 {
		t.Errorf("expected 3 subpackets")
	}
	if packet.subPackets[0].literal != 1 ||
		packet.subPackets[1].literal != 2 ||
		packet.subPackets[2].literal != 3 {
		t.Errorf("subpackets borked: %v", packet.subPackets)
	}
}

func TestSumVersion(t *testing.T) {
	tests := []struct {
		input string
		sum   int
	}{
		{input: "8A004A801A8002F478", sum: 16},
		{input: "620080001611562C8802118E34", sum: 12},
		{input: "C0015000016115A2E0802F182340", sum: 23},
		{input: "A0016C880162017C3686B18A3D4780", sum: 31},
	}
	for _, test := range tests {
		reader := NewPacketReader(strings.NewReader(test.input))
		packet, err := reader.ReadPacket()
		if err != nil {
			t.Error(err)
		}
		got := packet.SumVersion()
		if got != test.sum {
			t.Errorf("for input %s, got %d expected %d", test.input, got, test.sum)
		}
	}
}

func TestPart1(t *testing.T) {
	value := part1("example-input.txt")
	if value != 31 {
		t.Errorf("expected 31, got %d", value)
	}
}

func TestExecutePacket(t *testing.T) {
	tests := []struct {
		input string
		value int
	}{
		{input: "C200B40A82", value: 3},                 // sum
		{input: "04005AC33890", value: 54},              // product
		{input: "880086C3E88112", value: 7},             // minimum
		{input: "CE00C43D881120", value: 9},             // maximum
		{input: "D8005AC2A8F0", value: 1},               // <
		{input: "F600BC2D8F", value: 0},                 // >
		{input: "9C005AC2F8F0", value: 0},               // =
		{input: "9C0141080250320F1802104A08", value: 1}, // = (with subexpressions)
	}
	for _, test := range tests {
		reader := NewPacketReader(strings.NewReader(test.input))
		packet, err := reader.ReadPacket()
		if err != nil {
			t.Error(err)
		}
		got := packet.Execute()
		if got != test.value {
			t.Errorf("for input %s, got %d expected %d", test.input, got, test.value)
		}
	}
}

func TestPart2(t *testing.T) {
	value := part2("example-input.txt")
	if value != 54 {
		t.Errorf("expected 54, got %d", value)
	}
}

package day16

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Run(inputPath string) {
	fmt.Printf("Part 1: %d\n", part1(inputPath))
	fmt.Printf("Part 2: %d\n", part2(inputPath))
}

func part1(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	reader := NewPacketReader(file)
	packet, err := reader.ReadPacket()
	if err != nil {
		panic(err)
	}
	return packet.SumVersion()
}

func part2(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	reader := NewPacketReader(file)
	packet, err := reader.ReadPacket()
	if err != nil {
		panic(err)
	}
	return packet.Execute()
}

func NewPacketReader(reader io.Reader) *PacketReader {
	return &PacketReader{
		input: bufio.NewReader(reader),
	}
}

const (
	SumPacket         = 0
	ProductPacket     = 1
	MinimumPacket     = 2
	MaximumPacket     = 3
	LiteralPacket     = 4
	GreaterThanPacket = 5
	LessThanPacket    = 6
	EqualToPacket     = 7
)

type Packet struct {
	version    int
	typeID     int
	literal    int
	subPackets []Packet
}

func (p Packet) SumVersion() int {
	sum := p.version
	for _, sp := range p.subPackets {
		sum += sp.SumVersion()
	}
	return sum
}

func (p Packet) Execute() int {
	value := 0
	switch p.typeID {
	case SumPacket:
		for _, sp := range p.subPackets {
			value += sp.Execute()
		}
	case ProductPacket:
		value = 1
		for _, sp := range p.subPackets {
			value *= sp.Execute()
		}
	case MinimumPacket:
		value = -1
		for _, sp := range p.subPackets {
			v := sp.Execute()
			if value == -1 || v < value {
				value = v
			}
		}
	case MaximumPacket:
		value = -1
		for _, sp := range p.subPackets {
			v := sp.Execute()
			if value == -1 || v > value {
				value = v
			}
		}
	case LiteralPacket:
		value = p.literal
	case GreaterThanPacket:
		v1 := p.subPackets[0].Execute()
		v2 := p.subPackets[1].Execute()
		if v1 > v2 {
			value = 1
		}
	case LessThanPacket:
		v1 := p.subPackets[0].Execute()
		v2 := p.subPackets[1].Execute()
		if v1 < v2 {
			value = 1
		}
	case EqualToPacket:
		v1 := p.subPackets[0].Execute()
		v2 := p.subPackets[1].Execute()
		if v1 == v2 {
			value = 1
		}
	default:
		panic(fmt.Sprintf("bad packet type id: %d", p.typeID))
	}
	return value
}

type PacketReader struct {
	input      *bufio.Reader
	current    byte
	currentBit byte
	bitsRead   int
}

// readBits reads the next 'count' bits as an int
func (r *PacketReader) readBits(count int) (int, error) {
	value := 0
	for i := 0; i < count; i++ {
		bit, err := r.nextBit()
		if err != nil {
			return 0, err
		}
		value = value<<1 + bit
	}
	return value, nil
}

// Read the next bit from the input, automatically moving between the ASCII hex digits
// as needed.  Updates the total number of bits read for use in parsing bit-length
// sub packets.
func (r *PacketReader) nextBit() (int, error) {
	// if mask is zero it's time to read in next hex digit
	if r.currentBit == 0 {
		current, err := r.input.ReadByte()
		if err != nil {
			return 0, err
		}
		if current >= '0' && current <= '9' {
			current = current - '0'
		} else if current >= 'A' && current <= 'F' {
			current = current - 'A' + 10
		} else {
			return 0, fmt.Errorf("bad input: %c", rune(current))
		}
		r.current = current
		r.currentBit = 0b1000
	}
	bit := 0
	if r.current&r.currentBit != 0 {
		bit = 1
	}
	r.currentBit >>= 1
	r.bitsRead++
	return bit, nil
}

// ReadPacket reads the next packet and any subpackets from the underlying Reader.
// If any error occurs the reader state is undefined and no further reading should be done.
func (r *PacketReader) ReadPacket() (Packet, error) {
	return r.readPacket2(true)
}

// Read a packet and sub packets.  The parameter specifies whether this is the top-level
// packet parsing ie whether extra padded bits at the end should be discarded.
func (r *PacketReader) readPacket2(discardExtraBits bool) (Packet, error) {
	packet := Packet{
		subPackets: []Packet{},
	}
	version, err := r.readBits(3)
	if err != nil {
		return packet, err
	}
	packet.version = version
	typeID, err := r.readBits(3)
	if err != nil {
		return packet, err
	}
	packet.typeID = typeID
	if packet.typeID == LiteralPacket {
		literal, err := r.readLiteral()
		if err != nil {
			return packet, err
		}
		packet.literal = literal
	} else {
		// Operator packets contain one or more packets.
		lengthTypeID, err := r.nextBit()
		if err != nil {
			return packet, nil
		}
		if lengthTypeID == 0 {
			// Read packets until numBits bits have been read
			numBits, err := r.readBits(15)
			if err != nil {
				return packet, nil
			}
			stopBit := r.bitsRead + numBits - 1
			for r.bitsRead < stopBit {
				subPacket, err := r.readPacket2(false)
				if err != nil {
					return packet, err
				}
				packet.subPackets = append(packet.subPackets, subPacket)
			}
		} else {
			numPackets, err := r.readBits(11)
			if err != nil {
				return packet, nil
			}
			for i := 0; i < numPackets; i++ {
				subPacket, err := r.readPacket2(false)
				if err != nil {
					return packet, err
				}
				packet.subPackets = append(packet.subPackets, subPacket)
			}
		}
	}
	if discardExtraBits {
		r.currentBit = 0 // this is end of packet, any unread bits should be discarded
	}
	return packet, nil
}

// read a literal value from the reader, based on the 5 bit encoding described in README
func (r *PacketReader) readLiteral() (int, error) {
	value := 0
	more := true
	for more {
		group, err := r.readBits(5)
		if err != nil {
			return value, err
		}
		more = group&0b_0001_0000 != 0
		group &= 0b_0000_1111
		value = value<<4 | group
	}
	return value, nil
}

package packet

const (
	// NullPacketPid is the PID for a null packet = 0x1FFF
	NullPacketPid = 8191
)

// Packet is a basic unit in a transport stream. It correspond to a set
// of 188 bytes, beggining with a sync_byte = 0x47
type Packet [188]byte

// New parses a new packet, receiving its bytes.
func New(content [188]byte) (p Packet) {
	p = Packet(content)
	return
}

// Pid returns the packet identification number.
func (p Packet) Pid() uint16 {
	return uint16(p[1]&0x1f)<<8 | uint16(p[2])
}

// IsNull returns true if the packet correspond to a Null Packet.
func (p Packet) IsNull() bool {
	return p.Pid() == NullPacketPid
}

// IsStreamTypeC returns true if the packet has any of the pids received.
// Currently, that's the only way to find stream type C packets, since
// this project isn't reading PMT tables.
func (p Packet) IsStreamTypeC(pids []uint) bool {
	for i := 0; i < len(pids); i++ {
		if p.Pid() == uint16(pids[i]) {
			return true
		}
	}
	return false
}

// TransportErrorIndicator returns the Transport Error Indicator for the
// packet.
func (p Packet) TransportErrorIndicator() bool {
	return p.getBit(1, 0x80)
}

// getBit returns true if a bit in a packet is set to 1.
func (p Packet) getBit(index int, mask byte) bool {
	return p[index]&mask != 0
}

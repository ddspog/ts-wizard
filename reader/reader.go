package reader

import "github.com/ddsgok/ts-wizard/packet"

// Reader correspond to a Reader with skill to execute reading operations
// on a transport stream file.
type Reader interface {
	ActualPacket() packet.Packet
	ActualPacketID() int
	StillReading() bool
	ReadNextPacket() error
	SaveActualPacket() error
}

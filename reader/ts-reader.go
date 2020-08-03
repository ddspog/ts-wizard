package reader

import (
	"errors"
	"os"

	"github.com/ddsgok/ts-wizard/packet"
)

type tsReader struct {
	OrigFile        *os.File
	DestFile        *os.File
	stillReading    bool
	currentPacketID int
	actualPacket    packet.Packet

	numNullPackets        int
	maxNullPackets        int
	numStreamTypeCPackets int
	maxStreamTypeCPackets int

	streamTypeCPIDs []uint
}

func newTSReader(of, df *os.File, mn, ms int, stcPids []uint) (ts Reader) {
	ts = &tsReader{
		OrigFile:              of,
		DestFile:              df,
		stillReading:          true,
		currentPacketID:       -1,
		numNullPackets:        0,
		maxNullPackets:        mn,
		numStreamTypeCPackets: 0,
		maxStreamTypeCPackets: ms,
		streamTypeCPIDs:       stcPids,
	}

	return
}

func newQuickTSReader(of *os.File) (ts Reader) {
	ts = &tsReader{
		OrigFile:              of,
		stillReading:          true,
		currentPacketID:       -1,
		numNullPackets:        0,
		maxNullPackets:        -1,
		numStreamTypeCPackets: 0,
		maxStreamTypeCPackets: -1,
	}

	return
}

// ActualPacket returns the current packet the tsReader is reading.
func (ts *tsReader) ActualPacket() (p packet.Packet) {
	p = ts.actualPacket
	return
}

// ActualPacket find the id for the current packet that the tsReader is
// reading.
func (ts *tsReader) ActualPacketID() (id int) {
	id = ts.currentPacketID
	return
}

// StillReading returns true if the tsReader is still reading a ts file.
func (ts *tsReader) StillReading() (b bool) {
	b = ts.stillReading
	return
}

// ReadNextPacket tris to read a new packet from the ts file that the
// tsReader reads.
func (ts *tsReader) ReadNextPacket() (err error) {
	var p packet.Packet
	var nBytesRead int
	if nBytesRead, err = ts.OrigFile.Read(p[:]); err != nil {
		ts.stillReading = false
		return
	}

	if nBytesRead != 188 {
		ts.stillReading = false
		err = errors.New("Couldn't read the whole packet")
		return
	}

	ts.currentPacketID++
	ts.actualPacket = p
	return
}

// SaveActualPacket takes the current packet the tsReader reads, and save
// to the destiny file.
func (ts *tsReader) SaveActualPacket() (err error) {
	if ts.maxNullPackets != -1 && ts.actualPacket.IsNull() {
		ts.numNullPackets++

		if ts.numNullPackets > ts.maxNullPackets {
			return
		}
	}

	if ts.maxStreamTypeCPackets != -1 && len(ts.streamTypeCPIDs) > 0 && ts.actualPacket.IsStreamTypeC(ts.streamTypeCPIDs) {
		ts.numStreamTypeCPackets++

		if ts.numStreamTypeCPackets > ts.maxStreamTypeCPackets {
			return
		}
	}

	var nBytesWritten int
	if nBytesWritten, err = ts.DestFile.Write(ts.actualPacket[:]); err != nil {
		return
	}

	if nBytesWritten != 188 {
		err = errors.New("Couldn't write the whole packet")
		return
	}

	return
}

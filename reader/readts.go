package reader

import (
	"os"
)

// ReadTS takes the file on [filepath], and transfer its packets to a new
// file on [outputPath], filtering a number [mn] of null packets, and a
// number of other packets [ms] contained on a set of Pids [stcPids].
func ReadTS(filepath, outputPath string, mn, ms int, stcPids []uint) (err error) {
	var origF *os.File
	if origF, err = os.Open(filepath); err != nil {
		return
	}
	defer origF.Close()

	var destF *os.File
	if destF, err = os.OpenFile(outputPath, os.O_CREATE, 0777); err != nil {
		return
	}
	defer destF.Close()

	var rd = newTSReader(origF, destF, mn, ms, stcPids)

	for rd.StillReading() {
		if err = rd.ReadNextPacket(); err != nil {
			return
		}
		if err = rd.SaveActualPacket(); err != nil {
			return
		}
	}

	return
}

// FindOnTS takes the file on [filepath], and find how many corrupted packets
// there are on the transport stream.
func FindOnTS(filepath string, corrupted bool) (packetNumber int, err error) {
	var origF *os.File
	if origF, err = os.Open(filepath); err != nil {
		return
	}
	defer origF.Close()

	var rd = newQuickTSReader(origF)

	packetNumber = -1
	for rd.StillReading() {
		if err = rd.ReadNextPacket(); err != nil {
			return
		}

		if corrupted {
			if rd.ActualPacket().TransportErrorIndicator() == true {
				packetNumber = rd.ActualPacketID()
				return
			}
		}
	}

	return
}

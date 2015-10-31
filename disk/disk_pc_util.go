package disk

type diskIoTime struct {
	tstamp time.Time
	iotime uint64
}

var lastCheck = make(map[string]*diskIoTime)

func countPcUtil(tstamp time.Time, name string, iotime uint64) (pc float64) {
	last, ok := lastCheck[name]

	if !ok {
		disksStat[name] = &diskIoTime{tstamp, iotime}
		return
	}

	passed := tstamp.Sub(last.tstamp)

	if passed > 0 {
		pc = float64(iotime - last.iotime) / 1000 / passed.Seconds() * 100
		switch {
			case pc > 100 : pc = 100
			case pc <   0 : pc = 0
		}
		last.tstamp = tstamp
		last.iotime = iotime
	}

	return
}

package main

import (
	. "github.com/elamre/go_beagle_api/beagle_bindings/bindings"
	"log"
	"sync"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

var running = false
var wg sync.WaitGroup

func Stop() {
	running = false
	wg.Wait()
}

func dumpAll(beagle Beagle) {
	if running {
		log.Println("Already running, stopping and starting again")
		Stop()
	}
	wg.Add(1)
	running = true
	const bufferSize = 512
	timingSize, err := BgBitTimingSize(BG_PROTOCOL_SPI, bufferSize)
	checkError(err)
	mosi := make([]uint8, bufferSize)
	miso := make([]uint8, bufferSize)
	timing := make([]uint32, timingSize)
	samplerateKhz, err := BgSampleRate(beagle, 0)
	checkError(err)
	beagle.Enable(BG_PROTOCOL_SPI)

	go func() {
		for ; running; {
			count, status, timeSop, timeDuration, timeDataOffset, err := BgSpiRead(beagle, &mosi, &miso, &timing)
			_, _, _ = timeSop, timeDuration, timeDataOffset
			timeSopNs := (uint64(timeSop * 1000)) / (uint64(samplerateKhz / 1000))
			_ = timeSopNs
			if err != nil {
				log.Printf("error: %s", err)
				continue
			}

			if status != 0 {
				//if status != BG_READ_TIMEOUT {
					log.Println(status.String())
				//}
				continue
			}
			log.Printf("[%0 2x][%0 2x]", mosi[:count], miso[:count])
		}
		beagle.Disable()
		wg.Done()
	}()
}

func main() {
	ids, uniques := FindDevices()
	log.Printf("%+v %+v", ids, uniques)

	port := 0
	samplerate := int32(10000) // in kHz
	timeout := uint32(5000)    // in milliseconds
	latency := uint32(200)     // in milliseconds
	//length := int32(0)
	var beagle Beagle
	var err error

	beagle, err = NewBeagle(port)
	checkError(err)
	defer func() { checkError(beagle.Close()) }()

	samplerate, err = BgSampleRate(beagle, samplerate)
	checkError(err)
	beagle.SetTimeout(timeout)
	beagle.SetLatency(latency)
	log.Println(beagle.GetVersion().String())
	if beagle.IsFullSpeed() {
		log.Println("full speed")
	} else {
		log.Println("high speed")
	}
	beagle.SpiConfigure(BG_SPI_SS_ACTIVE_LOW, BG_SPI_SCK_SAMPLING_EDGE_RISING, BG_SPI_BITORDER_MSB)
	beagle.TargetPower(BG_TARGET_POWER_OFF)
	beagle.Enable(BG_PROTOCOL_SPI)
	beagle.Disable()
	dumpAll(beagle)
	time.Sleep(10000 * time.Millisecond)
	running = false
	Stop()
}

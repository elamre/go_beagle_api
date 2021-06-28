package main

import (
	"github.com/elamre/go_beagle_api/pkg/beagle"
	"log"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ids, uniques := beagle.FindDevices()
	log.Printf("%+v %+v", ids, uniques)

	port := 0
	samplerate := int32(10000) // in kHz
	timeout := uint32(5000)    // in milliseconds
	latency := uint32(200)     // in milliseconds
	//length := int32(0)
	var beagleDevice beagle.Beagle
	var err error

	beagleDevice, err = beagle.NewBeagle(port)
	checkError(err)
	defer func() { checkError(beagleDevice.Close()) }()

	samplerate, err = beagle.BgSampleRate(beagleDevice, samplerate)
	checkError(err)
	beagleDevice.SetTimeout(timeout)
	beagleDevice.SetLatency(latency)
	if beagleDevice.IsFullSpeed() {
		log.Println("full speed")
	} else {
		log.Println("high speed")
	}
	beagleDevice.SpiConfigure(beagle.BG_SPI_SS_ACTIVE_LOW, beagle.BG_SPI_SCK_SAMPLING_EDGE_RISING, beagle.BG_SPI_BITORDER_MSB)
	beagleDevice.TargetPower(beagle.BG_TARGET_POWER_OFF)
	beagle.ReadToCsvLoop(beagleDevice, beagle.BG_PROTOCOL_SPI, "Infineon", "GtpGo", "output")
	time.Sleep(10000 * time.Millisecond)
	beagle.StopReadToCsvLoop()
}

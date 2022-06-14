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
	samplerate := int32(50000) // in kHz
	timeout := uint32(5000)    // in milliseconds
	latency := uint32(200)     // in milliseconds
	//length := int32(0)
	var beagleDevice beagle.Beagle
	var err error

	beagleDevice, err = beagle.NewBeagle(port)
	checkError(err)
	defer func() { checkError(beagleDevice.Close()) }()

	beagleDevice.SetTimeout(timeout)
	beagleDevice.SetLatency(latency)
	if beagleDevice.IsFullSpeed() {
		log.Println("full speed")
	} else {
		log.Println("high speed")
	}
	beagleDevice.TargetPower(beagle.BG_TARGET_POWER_OFF)
	log.Printf("Samplerate: %d", beagleDevice.SetSampleRate(samplerate))
	beagle.ReadI2cToCsvLoop(beagleDevice, "Infineon", "GtpGo", "output")
	time.Sleep(10000 * time.Millisecond)
	beagle.StopReadToCsvLoop()
}

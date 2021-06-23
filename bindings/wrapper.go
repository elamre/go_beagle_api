package bindings

/*
#cgo linux LDFLAGS: -ldl
#include "beagle.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func checkRetVal(retVal C.int) error {
	if retVal >= 0 {
		return nil
	}
	return fmt.Errorf((C.BeagleStatus)(retVal).String())
}

func BgOpen(port int) (Beagle, error) {
	retVal := C.bg_open(C.int(port))
	err := checkRetVal(retVal)
	return Beagle(retVal), err
}

func BgClose(beagle Beagle) error {
	return checkRetVal(C.bg_close(C.int(beagle)))
}

func FindDevices() ([]uint16, []uint32) {
	const bufferSize = 32
	idBuffer := make([]uint16, bufferSize)
	uniqueIdBuffer := make([]uint32, bufferSize)
	connected := C.bg_find_devices_ext(C.int(bufferSize), (*C.ushort)(unsafe.Pointer(&idBuffer[0])), C.int(bufferSize), (*C.uint)(unsafe.Pointer(&uniqueIdBuffer[0])))
	err := checkRetVal(connected)
	if err != nil {
		panic(err)
	}
	return idBuffer[:connected], uniqueIdBuffer[:connected]
}

func BgEnable(beagle Beagle, beagleProtocol BeagleProtocol) error {
	return checkRetVal(C.bg_enable((C.Beagle)(beagle), (C.BeagleProtocol)(beagleProtocol)))
}

func BgDisable(beagle Beagle) error {
	return checkRetVal(C.bg_disable((C.Beagle)(beagle)))
}

func BgSampleRate(beagle Beagle, sampleRateKhz int32) (int32, error) {
	retVal := C.bg_samplerate(C.int(beagle), C.int(sampleRateKhz))
	err := checkRetVal(retVal)
	return int32(retVal), err
}

/* Bg_status_string function as declared in beagle/beagle.h:407
func Bg_status_string(Status int32) string {
	cStatus, cStatusAllocMap := (C.int)(Status), cgoAllocsUnknown
	__ret := C.bg_status_string(cStatus)
	runtime.KeepAlive(cStatusAllocMap)
	__v := packPCharString(__ret)
	return __v
}*/

// Bg_version function as declared in beagle/beagle.h:417
func BgVersion(beagle Beagle) (BeagleVersion, error) {
	version := BeagleVersion{}
	return version, checkRetVal(C.bg_version(C.int(beagle), (*C.BeagleVersion)(unsafe.Pointer(&version))))
}

func BgLatency(beagle Beagle, latencyMs uint32) error {
	return checkRetVal(C.bg_latency(C.int(beagle), C.u32(latencyMs)))
}

func BgTimeout(beagle Beagle, timeoutMs uint32) error {
	return checkRetVal(C.bg_timeout(C.int(beagle), C.u32(timeoutMs)))
}

func BgSleepMs(sleepMs uint32) uint32 {
	return uint32(C.bg_sleep_ms(C.u32(sleepMs)))
}

func BgTargetPower(beagle Beagle, powerFlag byte) error {
	return checkRetVal(C.bg_target_power(C.int(beagle), (C.u08)(powerFlag)))
}

func BgHostIfceSpeed(beagle Beagle) bool {
	if int32(C.bg_host_ifce_speed(C.int(beagle))) <= 0 {
		return true
	} else {
		return false
	}
}

func BgSpiConfigure(beagle Beagle, ss_polarity BeagleSpiSSPolarity, sck_sampling_edge BeagleSpiSckSamplingEdge, bitorder BeagleSpiBitorder) error {
	return checkRetVal(C.bg_spi_configure(C.int(beagle), (C.BeagleSpiSSPolarity)(ss_polarity), (C.BeagleSpiSckSamplingEdge)(sck_sampling_edge), (C.BeagleSpiBitorder)(bitorder)))
}

func BgBitTimingSize(beagleProtocol BeagleProtocol, numDataBytes int32) (int32, error) {
	retVal := C.bg_bit_timing_size(C.BeagleProtocol(beagleProtocol), C.int(numDataBytes))
	return int32(retVal), checkRetVal(retVal)
}

func BgSpiRead(beagle Beagle, mosi *[]byte, miso *[]byte, timings *[]uint32) (count int32, status ReadStatus, timeSop uint64, timeDuration uint64, timeDataOffset uint32, err error) {
	retVal := C.bg_spi_read_data_timing(
		C.int(beagle),
		(*C.u32)(unsafe.Pointer(&status)),
		(*C.u64)(unsafe.Pointer(&timeSop)),
		(*C.u64)(unsafe.Pointer(&timeDuration)),
		(*C.u32)(unsafe.Pointer(&timeDataOffset)),
		C.int(len(*mosi)),
		(*C.u08)(unsafe.Pointer(&(*mosi)[0])),
		C.int(len(*miso)),
		(*C.u08)(unsafe.Pointer(&(*miso)[0])),
		C.int(len(*timings)),
		(*C.u32)(unsafe.Pointer(&(*timings)[0])),
	)
	count = int32(retVal)
	return
}

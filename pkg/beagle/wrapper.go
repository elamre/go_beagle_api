package beagle

/*
#cgo linux LDFLAGS: -ldl
#cgo CFLAGS: -I. -I${SRCDIR}/
#include "beagle.h"
*/
import "C"
import (
	"bufio"
	"fmt"
	util2 "github.com/elamre/go_beagle_api/pkg/util"
	"log"
	"os"
	"strings"
	"sync"
	"time"
	"unsafe"
)

var running = false
var wg sync.WaitGroup

func ReadToCsvLoop(beagle Beagle, protocol BeagleProtocol, copyright, csvDescription, filename string) {
	if running {
		log.Println("Already running, stopping and starting again")
		StopReadToCsvLoop()
	}
	wg.Add(1)
	running = true
	const bufferSize = 512
	timingSize, err := BgBitTimingSize(BG_PROTOCOL_SPI, bufferSize)
	util2.CheckError(err)
	mosi := make([]uint8, bufferSize)
	miso := make([]uint8, bufferSize)
	timing := make([]uint32, timingSize)
	samplerateKhz, err := BgSampleRate(beagle, beagle.GetSampleRate())
	util2.CheckError(err)
	beagle.Enable(protocol)
	if !strings.HasSuffix(filename, ".csv") {
		filename += ".csv"
	}
	file, err := os.Create(filename)
	util2.CheckError(err)
	writer := bufio.NewWriter(file)
	util2.CheckErrorRetVal(writer.WriteString("# " + csvDescription + "\n"))
	util2.CheckErrorRetVal(writer.WriteString("# " + time.Now().String()))
	util2.CheckErrorRetVal(writer.WriteString("# " + copyright + "\n"))
	util2.CheckErrorRetVal(writer.WriteString("# beagle version: " + beagle.GetVersion().String() + "\n"))
	util2.CheckErrorRetVal(writer.WriteString("#\n"))
	util2.CheckErrorRetVal(writer.WriteString("#\n"))
	util2.CheckErrorRetVal(writer.WriteString("# Level,Index,m:s.ms.us,Dur,Len,Err,Record,Data\n"))
	util2.CheckErrorRetVal(writer.WriteString("0,0,0:00.000.000,,,,Capture started," + time.Now().String() + "\n"))
	level := 1 // No idea what this is for

	go func() {
		for ; running; {
			count, status, timeSop, timeDuration, timeDataOffset, err := BgSpiRead(beagle, &mosi, &miso, &timing)
			_, _, _ = timeSop, timeDuration, timeDataOffset
			timeSopms := (uint64(timeSop)) / (uint64(samplerateKhz / 1000))
			min := timeSopms / 60000000
			rest := timeSopms % 60000000
			sec := rest / 1000000
			rest = rest % 1000000
			ms := rest / 1000
			rest = rest % 1000
			us := rest
			if err != nil {
				log.Printf("error: %s", err)
				continue
			}

			if status != 0 {
				if status != BG_READ_TIMEOUT {
					log.Println(status.String())
				}
				continue
			}
			if count <= 0 {
				continue
			}
			dataString := ""
			for i := int32(0); i < count; i++ {
				dataString += fmt.Sprintf("%02x%02x ", mosi[i], miso[i])
			}
			timeString := fmt.Sprintf("%d:%d.%d.%d", min, sec, ms, us)
			durationString := ""
			w := fmt.Sprintf("0,%d,%s,%s,%d B,,Transaction,%s\n", level, timeString, durationString, count, dataString)
			util2.CheckErrorRetVal(writer.WriteString(w))
			level += 3
		}
		util2.CheckError(writer.Flush())
		beagle.Disable()
		wg.Done()
	}()
}

func StopReadToCsvLoop() {
	running = false
	wg.Wait()
}

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

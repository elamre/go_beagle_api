# go_beagle_api
Small subset of bindings working with the Beagle logic analyzer api

# Instructions
Make sure that the correct library is available when running your application. The library makes use of the official dynamic libraries that can be found a: https://www.totalphase.com/products/beagle-software-api/ .

The wrapper is far from complete, only very basic functionality is working. I rebuilt the Spi dump demo as well as the find devices demo which can be found in the official archive on TotalPhase's website. 
I have tested it on a windows pc as well as a Raspberry pi 4, where it worked fine

# Example
To print out all the attached devices
```
import gobeagle "github.com/elamre/go_beagle_api/bindings"

func main() {
	ids, uniques := gobeagle.FindDevices()
	log.Printf("%+v %+v", ids, uniques)
}
```
Show the version of the beagle device
```
import gobeagle "github.com/elamre/go_beagle_api/bindings"

func main() {
	beagle, err := gobeagle.NewBeagle(0)     // Port 0
	if err != nil{
	    panic(err.Error())
	}
	defer func() { _ = beagle.Close() }()     // Not doing anything on error
	log.Println(beagle.GetVersion().String())
}
```

# Supported functions
 - bg_open
 - bg_close
 - bg_find_devices_ext
 - bg_enable
 - bg_disable
 - bg_samplerate
 - bg_version
 - bg_latency
 - bg_timeout
 - bg_sleep_ms
 - bg_target_power
 - bg_host_ifce_speed
 - bg_spi_configure
 - bg_bit_timing_size
 - bg_spi_read_data_timing
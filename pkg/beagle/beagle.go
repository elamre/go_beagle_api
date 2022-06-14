package beagle

import "github.com/elamre/go_beagle_api/pkg/util"

type Beagle int32

var samplerate int32

func (b Beagle) GetSampleRate() int32 {
	if samplerate == 0 {
		samplerateKhz, err := BgSampleRate(b, 50000)
		util.CheckError(err)
		samplerate = samplerateKhz
	}
	return samplerate
}

func (b Beagle) SetSampleRate(rate int32) int32 {
	samplerate = rate
	samplerateKhz, err := BgSampleRate(b, 50000)
	util.CheckError(err)
	return samplerateKhz

}

func NewBeagle(port int) (Beagle, error) {
	return BgOpen(port)
}

func (b Beagle) GetVersion() BeagleVersion {
	v, err := BgVersion(b)
	if err != nil {
		panic(err)
	}
	return v
}

func (b Beagle) Close() error {
	return BgClose(b)
}

func (b Beagle) SetTimeout(ms uint32) {
	if err := BgTimeout(b, ms); err != nil {
		panic(err)
	}
}
func (b Beagle) SetLatency(ms uint32) {
	if err := BgLatency(b, ms); err != nil {
		panic(err)
	}
}
func (b Beagle) I2cSetPullup(pullup BeagleI2cPullup) {
	if err := BgI2cPullup(b, pullup); err != nil {
		panic(err)
	}
}
func (b Beagle) SpiConfigure(polarity BeagleSpiSSPolarity, samplingEdge BeagleSpiSckSamplingEdge, order BeagleSpiBitorder) {
	if err := BgSpiConfigure(b, polarity, samplingEdge, order); err != nil {
		panic(err)
	}
}
func (b Beagle) TargetPower(powerFlag byte) {
	if err := BgTargetPower(b, powerFlag); err != nil {
		panic(err)
	}
}
func (b Beagle) Enable(protocol BeagleProtocol) {
	if err := BgEnable(b, protocol); err != nil {
		panic(err)
	}
}
func (b Beagle) Disable() {
	if err := BgDisable(b); err != nil {
		panic(err)
	}
}
func (b Beagle) IsFullSpeed() bool {
	return BgHostIfceSpeed(b)
}

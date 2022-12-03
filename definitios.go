/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/12/3 12:04
 * @Desc:
 */

package go_ads1015

import (
	"github.com/AghostPrj/go-i2c"
	"time"
)

const (
	startSingleConversion      = uint16(0b1)
	startSingleConversionShift = 15

	Input0Base1   = uint16(0b000)
	Input0Base3   = uint16(0b001)
	Input1Base3   = uint16(0b010)
	Input2Base3   = uint16(0b011)
	Input0BaseGnd = uint16(0b100)
	Input1BaseGnd = uint16(0b101)
	Input2BaseGnd = uint16(0b110)
	Input3BaseGnd = uint16(0b111)
	inputMulShift = 12

	Fsr6144mV = uint16(0b000)
	Fsr4096mV = uint16(0b001)
	Fsr2048mV = uint16(0b010)
	Fsr1024mV = uint16(0b011)
	Fsr512mV  = uint16(0b100)
	Fsr256mV  = uint16(0b101)
	fsrShift  = 9

	OperationModeContinuous = uint16(0b0)
	OperationModeSingleShot = uint16(0b1)
	operationModeShift      = 8

	DataRate128Sps  = uint16(0b000)
	DataRate250Sps  = uint16(0b001)
	DataRate490Sps  = uint16(0b010)
	DataRate920Sps  = uint16(0b011)
	DataRate1600Sps = uint16(0b100)
	DataRate2400Sps = uint16(0b101)
	DataRate3300Sps = uint16(0b110)
	dataRateShift   = 5

	dataRate128SpsDelay  = time.Microsecond * 7830 * 11 / 10
	dataRate250SpsDelay  = time.Microsecond * 4010 * 11 / 10
	dataRate490SpsDelay  = time.Microsecond * 2060 * 11 / 10
	dataRate920SpsDelay  = time.Microsecond * 1077 * 11 / 10
	dataRate1600SpsDelay = time.Microsecond * 636 * 11 / 10
	dataRate2400SpsDelay = time.Microsecond * 427 * 11 / 10
	dataRate3300SpsDelay = time.Microsecond * 314 * 11 / 10

	ComparatorModeTraditional = uint16(0b0)
	ComparatorModeWindow      = uint16(0b1)
	comparatorModeShift       = 4

	comparatorPolarityLow   = uint16(0b0)
	comparatorPolarityShift = 3

	comparatorLatchingNoLatching = uint16(0b0)
	comparatorLatchingShift      = 2

	comparatorQueueDisable = uint16(0b11)
	comparatorQueueShift   = 0

	adcResultShift = 4

	operationDelay = time.Millisecond

	configReg = byte(0x1)
	resultReg = byte(0x0)
)

type Ads1015Config struct {
	Input          uint16 `json:"input"`
	Fsr            uint16 `json:"fsr"`
	OperationMode  uint16 `json:"operation_mode"`
	DataRate       uint16 `json:"data_rate"`
	ComparatorMode uint16 `json:"comparator_mode"`
}

type Ads1015 struct {
	fp     *i2c.I2C
	Config *Ads1015Config
}

func (conf *Ads1015Config) Marshal() uint16 {
	result := uint16(0)

	result |= startSingleConversion << startSingleConversionShift

	if checkValue(conf.Input, Input0Base1, Input0Base3, Input1Base3, Input2Base3,
		Input0BaseGnd, Input1BaseGnd, Input2BaseGnd, Input3BaseGnd) {
		result |= conf.Input << inputMulShift
	} else {
		result |= Input0BaseGnd << inputMulShift
	}

	if checkValue(conf.Fsr, Fsr6144mV, Fsr4096mV, Fsr2048mV, Fsr1024mV,
		Fsr512mV, Fsr256mV) {
		result |= conf.Fsr << fsrShift
	} else {
		result |= Fsr2048mV << fsrShift
	}

	if checkValue(conf.OperationMode, OperationModeContinuous, OperationModeSingleShot) {
		result |= conf.OperationMode << operationModeShift
	} else {
		result |= OperationModeSingleShot << operationModeShift
	}

	if checkValue(conf.DataRate, DataRate128Sps, DataRate250Sps, DataRate490Sps, DataRate920Sps,
		DataRate2400Sps, DataRate3300Sps) {
		result |= conf.DataRate << dataRateShift
	} else {
		result |= DataRate1600Sps << dataRateShift
	}

	if checkValue(conf.ComparatorMode, ComparatorModeTraditional, ComparatorModeWindow) {
		result |= conf.ComparatorMode << comparatorModeShift
	} else {
		result |= ComparatorModeTraditional << comparatorModeShift
	}

	result |= comparatorPolarityLow << comparatorPolarityShift
	result |= comparatorLatchingNoLatching << comparatorLatchingShift
	result |= comparatorQueueDisable << comparatorQueueShift

	return result
}

func checkValue(val uint16, data ...uint16) bool {
	for _, d := range data {
		if val == d {
			return true
		}
	}

	return false
}

func (a *Ads1015) Close() error {
	return a.fp.Close()
}

func (a *Ads1015) checkRunning() (bool, error) {
	data, err := a.fp.ReadRegU16BEWithDelay(configReg, operationDelay)
	if err != nil {
		return false, err
	}
	return ((data >> startSingleConversionShift) & 0b1) == 0b0, nil
}

func (a *Ads1015) RunComparator() (int16, error) {
	err := a.fp.WriteRegU16BE(configReg, a.Config.Marshal())
	if err != nil {
		return 0, err
	}

	switch a.Config.DataRate {
	case DataRate128Sps:
		time.Sleep(dataRate128SpsDelay)
		break
	case DataRate250Sps:
		time.Sleep(dataRate250SpsDelay)
		break
	case DataRate490Sps:
		time.Sleep(dataRate490SpsDelay)
		break
	case DataRate920Sps:
		time.Sleep(dataRate920SpsDelay)
		break
	case DataRate1600Sps:
		time.Sleep(dataRate1600SpsDelay)
		break
	case DataRate2400Sps:
		time.Sleep(dataRate2400SpsDelay)
		break
	case DataRate3300Sps:
		time.Sleep(dataRate3300SpsDelay)
		break
	default:
		time.Sleep(dataRate128SpsDelay)
		break
	}

	data, err := a.fp.ReadRegU16LEWithDelay(resultReg, operationDelay)
	if err != nil {
		return 0, err
	}

	data = data << adcResultShift

	isNeg := data<<11 == 0b1

	data &= 0b011111111111

	result := int16(data)

	switch a.Config.Fsr {
	case Fsr6144mV:
		result *= 3
		break
	case Fsr4096mV:
		result *= 2
		break
	case Fsr2048mV:
		break
	case Fsr1024mV:
		result /= 2
		break
	case Fsr512mV:
		result /= 4
		break
	case Fsr256mV:
		result /= 8
		break
	default:
		break
	}

	if isNeg {
		result *= -1
	}

	return result, nil
}
func (a *Ads1015) RunComparatorInput0() (int16, error) {
	a.Config.Input = Input0BaseGnd
	return a.RunComparator()
}
func (a *Ads1015) RunComparatorInput1() (int16, error) {
	a.Config.Input = Input1BaseGnd
	return a.RunComparator()
}
func (a *Ads1015) RunComparatorInput2() (int16, error) {
	a.Config.Input = Input2BaseGnd
	return a.RunComparator()
}
func (a *Ads1015) RunComparatorInput3() (int16, error) {
	a.Config.Input = Input3BaseGnd
	return a.RunComparator()
}
func (a *Ads1015) RunComparatorInput0BaseInput1() (int16, error) {
	a.Config.Input = Input0Base1
	return a.RunComparator()
}
func (a *Ads1015) RunComparatorInput0BaseInput3() (int16, error) {
	a.Config.Input = Input0Base3
	return a.RunComparator()
}
func (a *Ads1015) RunComparatorInput1BaseInput3() (int16, error) {
	a.Config.Input = Input1Base3
	return a.RunComparator()
}
func (a *Ads1015) RunComparatorInput2BaseInput3() (int16, error) {
	a.Config.Input = Input2Base3
	return a.RunComparator()
}

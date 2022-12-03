/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/12/3 12:42
 * @Desc:
 */

package go_ads1015

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestAds1015Config_Marshal(t *testing.T) {
	conf := new(Ads1015Config)

	conf.Input = Input1BaseGnd
	conf.Fsr = Fsr4096mV
	conf.OperationMode = OperationModeSingleShot
	conf.DataRate = DataRate3300Sps
	conf.ComparatorMode = ComparatorModeTraditional

	rawConf := conf.Marshal()
	if rawConf != uint16(0b1101001111000011) {
		t.Errorf("marshal raw config error: 0b%b/0x%x", rawConf, rawConf)
	}
	logrus.
		WithField("raw_conf", fmt.Sprintf("0x%x", rawConf)).
		WithField("op", "test").
		Info()
}
func TestAds1015(t *testing.T) {
	ads1015, err := NewAds1015(0x48, 2, &Ads1015Config{
		Input:          Input1BaseGnd,
		Fsr:            Fsr2048mV,
		OperationMode:  OperationModeSingleShot,
		DataRate:       DataRate3300Sps,
		ComparatorMode: ComparatorModeTraditional,
	})
	if err != nil {
		t.Errorf("get hdc1080 client error: %s", err)
	}

	defer func() {
		_ = ads1015.Close()
	}()

	for j := 0; j < 500; j++ {
		dataArr := make([]int16, 0)

		startTime := time.Now().UnixMicro()
		for i := 0; i < 100; i++ {

			d, err := ads1015.RunComparatorInput0()
			if err != nil {
				t.Errorf("get ads1015 result error: %s", err)
			}

			dataArr = append(dataArr, d)
		}

		endTime := time.Now().UnixMicro()

		total := int64(0)

		for _, d := range dataArr {
			total += int64(d)
		}

		logrus.
			WithField("data", total/int64(len(dataArr))).
			WithField("time", endTime-startTime).
			WithField("op", "test").
			Info()
	}
}

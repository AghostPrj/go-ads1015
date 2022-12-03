/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/12/3 12:04
 * @Desc:
 */

package go_ads1015

import "github.com/AghostPrj/go-i2c"

func NewAds1015(addr uint8, bus int, conf *Ads1015Config) (*Ads1015, error) {
	c, err := i2c.NewI2C(addr, bus)
	if err != nil {
		return nil, err
	}

	result := Ads1015{
		fp:     c,
		Config: conf,
	}

	return &result, err

}

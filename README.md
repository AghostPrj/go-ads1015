# go-ads1015

德州仪器ads1015 adc golang库  
ti ads1015 adc golang library


----------------------------

这个库使用 [golang](https://golang.org/) 编写，用于操作ads1015。

This library written in [Go programming language](https://golang.org/) intended to operation ads1015 chip.

-------------------

## Usage

```go
func main(){
ads1015, err := NewAds1015(0x48, 2, &Ads1015Config{
Input:          Input1BaseGnd,
Fsr:            Fsr2048mV,
OperationMode:  OperationModeSingleShot,
DataRate:       DataRate3300Sps,
ComparatorMode: ComparatorModeTraditional,
})
if err != nil {
// ....
}
defer ads1015.Close()
d, err := ads1015.RunComparatorInput0()
}
```

详细使用方法参考 [ads1015_test.go](./pkg/go_ads1015/ads1015_test.go)  




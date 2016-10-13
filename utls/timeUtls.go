package utls

import "time"

type timeUtls struct {

}


var TimeUtls = timeUtls{}

//获取GTM的时间戳
func (this timeUtls) GetGTMTimeTamp() int{
	currTime := time.Now()
	_, offset := currTime.Zone()

	return int(currTime.Unix()) - offset
}

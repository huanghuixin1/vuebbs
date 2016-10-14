package utls

import (
	"time"
	"strconv"
)

type timeUtls struct {

}

var TimeUtls = timeUtls{}

//获取GTM的时间戳
func (this timeUtls) GetGTMTimeTamp() int {
	currTime := time.Now()
	_, offset := currTime.Zone()

	return int(currTime.Unix()) - offset
}

//获取当前时间的GMT时间
func (this timeUtls) GetNow4GMT() time.Time {
	currTime := time.Now()

	_, offset := currTime.Zone()

	durationOffset, _ := time.ParseDuration(strconv.Itoa(offset) + "s")
	return currTime.Add(durationOffset * -1)
}
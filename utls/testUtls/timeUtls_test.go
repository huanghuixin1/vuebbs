package testUtls

import "testing"
import "../timeUtls"


//[ArticlesDAL.GetDetail]
func Test_GetGTMTimeTamp(t *testing.T) {
	t.Log(timeUtls.TimeUtls.GetGTMTimeTamp())
}


func Test_GetNow4GMT(t *testing.T) {
	t.Log(timeUtls.TimeUtls.GetNow4GMT())
}
package utls

import "testing"


//[ArticlesDAL.GetDetail]
func Test_GetGTMTimeTamp(t *testing.T) {
	t.Log(TimeUtls.GetGTMTimeTamp())
}


func Test_GetNow4GMT(t *testing.T) {
	t.Log(TimeUtls.GetNow4GMT())
}
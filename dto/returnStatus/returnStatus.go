package returnStatus

/**
返回的结果数据
 */
type ReturnStatus struct {
	Status int //状态码
	Data interface{} //返回的内容
}
//
//func (this *ReturnStatus) New() ReturnStatus {
//
//}

const(
	Ok = 0 //一切正常
	Error = 1 //出现错误
	ParamError = 2 //参数错误
	NoData = 100 //没有数据


	//用户相关
	UserExist = 100001
)
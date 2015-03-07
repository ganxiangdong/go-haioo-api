package code

type CodeItem struct {
	Code        int32  //code编号
	Description string //code描述
}
type codeList struct {
	//成功处理
	SUCCESS CodeItem

	//系统异常
	CATCH_EXCEPTION CodeItem

	//系统繁忙
	SYSTEM_BUSY CodeItem

	//无效的参数
	INVALID_PARAMETER CodeItem
}

var Response codeList //声明供外部使用的变量
func init() {
	//系统层编号
	Response.SUCCESS = CodeItem{0, "success"}
	Response.CATCH_EXCEPTION = CodeItem{-1, "catch exception"}
	Response.SYSTEM_BUSY = CodeItem{-3, "system busy"}
	//逻辑层编号，从-10001递减
	Response.INVALID_PARAMETER = CodeItem{-10001, "invalid parameter"}
}

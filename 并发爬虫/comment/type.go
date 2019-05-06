package comment
//请求需要地址和对请求内容的解析方法
type Request struct {
	Url 	string
	ParseFunc	func([]byte) ParseResult	//解析方法需要对http请求返回的切片内容进行解析，最后返回解析结果
}
//分析结果需要的到返回的城市、用户名称、详情等具体信息，和对应的请求
type ParseResult struct {
	Iterm	[]interface{}
	Request	[]Request
}
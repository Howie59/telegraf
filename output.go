package telegraf

// 普通output输出插件
type Output interface {
	PluginDescriber

	// Connect初始化到目的端的链接，例如链接存储服务器等
	Connect() error
	// Close关闭到目的端的链接，telegraf退出的时候执行
	Close() error
	// Write接收时序信息然后输出到目的端，telegraf会缓存interval时长的数据，批量写入
	// interval默认配置在telegraf.conf > agent > interval
	// 可以自己在插件中新增这个配置项，插件struct中这个配置必须是interval才可以被telegraf探测到自定义配置
	// 执行时长过长会报错
	Write(metrics []Metric) error
}

// AggregatingOutput添加函数式方法到output。如果output一段时间内仅接收固定的aggregations
// 这些方法可以被并发调用
type AggregatingOutput interface {
	Output

	// 添加metrics到aggregator
	Add(in Metric)
	// Push返回聚合的metrics并每隔flush interval被调用一次
	Push() []Metric
	// Reset在结束的时候signal一次
	Reset()
}

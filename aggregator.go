package telegraf

// Aggregator插件会对一段时间内的指标进行聚合/计算，然后产生新的指标输出。可以计算最大值，最小值，均值，标准差等。
// 插件可以指定对应的原始指标项是否抛弃，通过Add方法的返回值实现。插件的接口比较多，telegraf会保证不会并发执行
// 运行在Processor插件之后，在聚合完成之后，telegraf会让新产生的指标重新过一遍Processor插件
type Aggregator interface {
	PluginDescriber

	// Add 不断往聚合器中添加metrics
	Add(in Metric)

	// Push 将汇总出来的结果上报到聚合器
	Push(acc Accumulator)

	// Reset 将上一周期的计算结果清空，请求重置插件
	Reset()
}

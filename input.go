package telegraf

type Input interface {
	PluginDescriber

	// Gather将收集到的指标上报到Accumulator
	// 这个函数是周期性被telegraf调用的，默认配置在telegraf.conf > agent.interval 指定
	// 也可以自定义interval，需要在配置项中声明interval变量，返回整数时间，单位是秒
	Gather(Accumulator) error
}

// 特殊的input插件，serviceInput
// 可以用来启动一个并行的服务，并自行决定什么时候上报数据
type ServiceInput interface {
	Input

	// ServiceInput插件会被显式的调用Start接口来启动服务
	// Accumulator是数据上报的接口，ServiceInput这个插件可以将这个保存下来
	// 通过Accumulator，插件可以自己决定什么时候上报什么数据，也可以通过Gather接口被动上报
	Start(Accumulator) error

	// Stop为停止ServiceInput插件，主要是清理不再需要的资源，关闭链接，关闭channel
	Stop()
}

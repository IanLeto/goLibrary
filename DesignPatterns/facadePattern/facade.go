package facadePattern

// 外观模式，将多个业务封装到一个中，对外暴露
type Device interface {
	start() bool
}

type CPU struct {
}

func (c CPU) start() bool {
	return true
}

type Disk struct {
}

func (d Disk) start() bool {
	return true
}

type Memory struct {
}

func (m Memory) start() bool {
	return true
}

// 开机启动项
type StartButton struct {
}

func (s StartButton) start() bool {
	var (
		cpu    = &CPU{}
		memory = &Memory{}
		disk   = &Disk{}
	)
	if cpu.start() && memory.start() && disk.start() {
		return true
	}
	return false
}

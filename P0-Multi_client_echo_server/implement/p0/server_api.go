package p0

type MultiEchoServer interface {
	Start(port int) error
	Count() int
	Close()
}

package tss

type Request interface {
	Type() uint8
}

type Response interface {
	Type() uint8
}

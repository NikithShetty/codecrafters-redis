package eventloop

type CmdStruct struct {
	In  []byte
	Out chan []byte
}

type CmdQ = chan CmdStruct

type ReadCmdQ = <-chan CmdStruct
type WriteCmdQ = chan<- CmdStruct

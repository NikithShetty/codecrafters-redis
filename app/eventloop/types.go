package eventloop

import "codecrafters-redis/app/commands"

type CmdStruct struct {
	In  *commands.Command
	Out chan *commands.CmdResult
}

type CmdQ = chan CmdStruct

type ReadCmdQ = <-chan CmdStruct
type WriteCmdQ = chan<- CmdStruct

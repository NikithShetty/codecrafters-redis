package commands

import "errors"

type CmdType string

const (
	ECHO  CmdType = "echo"
	PING  CmdType = "ping"
	GET   CmdType = "get"
	SET   CmdType = "set"
	SETPX CmdType = "setpx"
)

type Command struct {
	CmdType CmdType
	Data    []string
}

type RedisDataType string

const (
	Error        RedisDataType = "Error"
	SimpleString RedisDataType = "SimpleString"
	BulkString   RedisDataType = "BulkString"
	Integer      RedisDataType = "Integer"
	Array        RedisDataType = "Array"
)

type CmdResult struct {
	DataType     RedisDataType
	err          string
	simpleString string

	// Using pointers to add support for nil values
	bulkString *string
	array      *[]string
}

func NewErrResult(err string) *CmdResult {
	return &CmdResult{DataType: Error, err: err}
}

func NewSimpleStringResult(str string) *CmdResult {
	return &CmdResult{DataType: SimpleString, simpleString: str}
}

func NewBulkStringResult(bulkStr *string) *CmdResult {
	return &CmdResult{DataType: BulkString, bulkString: bulkStr}
}

func NewArrResult(arr *[]string) *CmdResult {
	return &CmdResult{DataType: Array, array: arr}
}

func (cmdRes *CmdResult) GetCmdResDataType() RedisDataType {
	return cmdRes.DataType
}

func (cmdRes *CmdResult) GetErr() (string, error) {
	if cmdRes.DataType == Error {
		return cmdRes.err, nil
	}

	return "", errors.New("illegal get operation on data type" + string(cmdRes.DataType))
}

func (cmdRes *CmdResult) GetSimpleStr() (string, error) {
	if cmdRes.DataType == SimpleString {
		return cmdRes.simpleString, nil
	}

	return "", errors.New("illegal get operation on data type" + string(cmdRes.DataType))
}

func (cmdRes *CmdResult) GetBlkStr() (*string, error) {
	if cmdRes.DataType == BulkString {
		return cmdRes.bulkString, nil
	}

	return nil, errors.New("illegal get operation on data type" + string(cmdRes.DataType))
}

func (cmdRes *CmdResult) GetArr() (*[]string, error) {
	if cmdRes.DataType == Array {
		return cmdRes.array, nil
	}

	return nil, errors.New("illegal get operation on data type" + string(cmdRes.DataType))
}

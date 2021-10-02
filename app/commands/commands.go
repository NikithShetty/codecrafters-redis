package commands

import "errors"

type CmdType string

const (
	ECHO CmdType = "ECHO"
	PING CmdType = "PING"
	GET  CmdType = "GET"
	SET  CmdType = "SET"
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
	bulkString   string
	array        []string
}

func NewErrResult(resType RedisDataType, err string) *CmdResult {
	return &CmdResult{DataType: resType, err: err}
}

func NewSimpleStringResult(resType RedisDataType, str string) *CmdResult {
	return &CmdResult{DataType: resType, simpleString: str}
}

func NewBulkStringResult(resType RedisDataType, bulkStr string) *CmdResult {
	return &CmdResult{DataType: resType, bulkString: bulkStr}
}

func NewArrResult(resType RedisDataType, arr []string) *CmdResult {
	return &CmdResult{DataType: resType, array: arr}
}

func (cmdRes *CmdResult) GetCmdResDataType() RedisDataType {
	return cmdRes.DataType
}

func (cmdRes *CmdResult) GetErr() (string, error) {
	if cmdRes.DataType == Error {
		return cmdRes.err, nil
	}

	return "", errors.New("Ilegal get operation on data type" + string(cmdRes.DataType))
}

func (cmdRes *CmdResult) GetSimpleStr() (string, error) {
	if cmdRes.DataType == SimpleString {
		return cmdRes.simpleString, nil
	}

	return "", errors.New("Ilegal get operation on data type" + string(cmdRes.DataType))
}

func (cmdRes *CmdResult) GetBlkStr() (string, error) {
	if cmdRes.DataType == BulkString {
		return cmdRes.bulkString, nil
	}

	return "", errors.New("Ilegal get operation on data type" + string(cmdRes.DataType))
}

func (cmdRes *CmdResult) GetArr() ([]string, error) {
	if cmdRes.DataType == Array {
		return cmdRes.array, nil
	}

	return []string{}, errors.New("Ilegal get operation on data type" + string(cmdRes.DataType))
}

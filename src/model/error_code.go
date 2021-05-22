package model

const (
	Ok               = 0
	ReqParamError    = 10000 + iota
	MysqlInsertError = 10000 + iota
	MysqlQueryError  = 10000 + iota
)

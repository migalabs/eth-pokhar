package models

type ModelType int8

const (
	TransactionModel ModelType = iota
	BeaconDepositModel
)

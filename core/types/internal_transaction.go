package types

import (
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

//TODO: Add BlockNumber
type InternalTransaction struct {
	TxHash     common.Hash    `json:"txHash"`
	TxIndex    uint           `json:"txIndex"`
	BlockHash  common.Hash    `json:"blockHash"`
	Address    common.Address `json:"address"`
	To         common.Address `json:"to"`
	From       common.Address `json:"from"`
	Index      uint           `json:"index"`
	StackDepth int            `json:"stackDepth"`
	Value      *big.Int       `json:"value"`
	Input      []byte         `json:"input"`
}

//go:generate go run ../../rlp/rlpgen -type rlpInternalTransaction -out gen_internal_transaction_rlp.go

type rlpInternalTransaction struct {
	Address common.Address
	To      common.Address
	From    common.Address
	Input   []byte
}

func (t *InternalTransaction) EncodeRLP(w io.Writer) error {
	rl := rlpInternalTransaction{Address: t.Address, To: t.To, From: t.From, Input: t.Input}
	return rlp.Encode(w, &rl)
}

func (t *InternalTransaction) DecodeRLP(s *rlp.Stream) error {
	var dec rlpInternalTransaction
	err := s.Decode(&dec)
	if err == nil {
		t.Address, t.To, t.From, t.Input = dec.Address, dec.To, dec.From, dec.Input
	}
	return err
}

type InternalTransactionForStorage InternalTransaction

func (t *InternalTransactionForStorage) EncodeRLP(w io.Writer) error {
	rl := rlpInternalTransaction{Address: t.Address, To: t.To, From: t.From, Input: t.Input}
	return rlp.Encode(w, &rl)
}

func (t *InternalTransactionForStorage) DecodeRLP(s *rlp.Stream) error {
	blob, err := s.Raw()
	if err != nil {
		return err
	}
	var dec rlpInternalTransaction
	err = rlp.DecodeBytes(blob, &dec)
	if err == nil {
		*t = InternalTransactionForStorage{
			Address: dec.Address,
			To:      dec.To,
			From:    dec.From,
			Input:   dec.Input,
		}
	}

	return err
}

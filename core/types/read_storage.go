package types

import (
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

//Storage     map[[32]byte][]byte `json:"storage"`
type ReadStorage struct {
	TxHash      common.Hash    `json:"txHash"`
	TxIndex     uint           `json:"txIndex"`
	BlockHash   common.Hash    `json:"blockHash"`
	BlockNumber uint64         `json:"blockNumber"`
	Index       uint           `json:"index"`
	Address     common.Address `json:"address"`
	Slot        *big.Int       `json:"slot"`
	Value       *big.Int       `json:"value"`
}

//go:generate go run ../../rlp/rlpgen -type rlpReadStorage -out gen_read_storage_rlp.go
type rlpReadStorage struct {
	Address common.Address
	Slot    *big.Int
	Value   *big.Int
}

func (r *ReadStorage) EncodeRLP(w io.Writer) error {
	rl := rlpReadStorage{
		Address: r.Address,
		Slot:    r.Slot,
		Value:   r.Value,
	}
	return rlp.Encode(w, &rl)
}

func (r *ReadStorage) DecodeRLP(s *rlp.Stream) error {
	var dec rlpReadStorage
	err := s.Decode(&dec)
	if err == nil {
		r.Address, r.Slot, r.Value = dec.Address, dec.Slot, dec.Value
	}
	return err
}

type ReadStorageForStorage ReadStorage

func (r *ReadStorageForStorage) EncodeRLP(w io.Writer) error {
	rl := rlpReadStorage{Address: r.Address, Slot: r.Slot, Value: r.Value}
	return rlp.Encode(w, &rl)
}

func (r *ReadStorageForStorage) DecodeRLP(s *rlp.Stream) error {
	blob, err := s.Raw()
	if err != nil {
		return err
	}
	var dec rlpReadStorage
	err = rlp.DecodeBytes(blob, &dec)
	if err == nil {
		*r = ReadStorageForStorage{
			Address: dec.Address,
			Slot:    dec.Slot,
			Value:   dec.Value,
		}
	}

	return err
}

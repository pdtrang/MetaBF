package ppt_filter

import (
	"fmt"
)

type Bacteria struct {
	Signatures *Int64Set
	Threshold float32
	Reported bool
}

func NewBacteria(t float32) *Bacteria {
	return &Bacteria {
		Signatures: NewInt64Set(),
		Threshold: t,
		Reported: false,
	}
}

func (b *Bacteria) AddSignature(j int64) {
	if !b.Signatures.Has(j) {
		b.Signatures.Add(j)
	}
}

func (b *Bacteria) PrintBacteria() {
	fmt.Println("Signatures: ", b.Signatures)
	fmt.Println("Threshold: ", b.Threshold)
	fmt.Println("Reported: ", b.Reported)
}

func (b *Bacteria) ReachThreshold() bool {
	if float32(b.Signatures.Size()) >= b.Threshold {
		return true
	} else {
		return false
	}
}
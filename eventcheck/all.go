package eventcheck

import (
	"github.com/motif-foundation/motif-blockchain/eventcheck/basiccheck"
	"github.com/motif-foundation/motif-blockchain/eventcheck/epochcheck"
	"github.com/motif-foundation/motif-blockchain/eventcheck/gaspowercheck"
	"github.com/motif-foundation/motif-blockchain/eventcheck/heavycheck"
	"github.com/motif-foundation/motif-blockchain/eventcheck/parentscheck"
	"github.com/motif-foundation/motif-blockchain/inter"
)

// Checkers is collection of all the checkers
type Checkers struct {
	Basiccheck    *basiccheck.Checker
	Epochcheck    *epochcheck.Checker
	Parentscheck  *parentscheck.Checker
	Gaspowercheck *gaspowercheck.Checker
	Heavycheck    *heavycheck.Checker
}

// Validate runs all the checks except Poset-related
func (v *Checkers) Validate(e inter.EventPayloadI, parents inter.EventIs) error {
	if err := v.Basiccheck.Validate(e); err != nil {
		return err
	}
	if err := v.Epochcheck.Validate(e); err != nil {
		return err
	}
	if err := v.Parentscheck.Validate(e, parents); err != nil {
		return err
	}
	var selfParent inter.EventI
	if e.SelfParent() != nil {
		selfParent = parents[0]
	}
	if err := v.Gaspowercheck.Validate(e, selfParent); err != nil {
		return err
	}
	if err := v.Heavycheck.Validate(e); err != nil {
		return err
	}
	return nil
}

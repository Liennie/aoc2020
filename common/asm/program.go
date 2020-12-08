package asm

import (
	"fmt"
)

type Registers struct {
	Acc int
	Ip  int
}

type Program struct {
	reg          Registers
	instructions []Instruction
}

func (p *Program) Next() Instruction {
	return p.instructions[p.reg.Ip]
}

func (p *Program) Step(callbacks ...Callback) error {
	i := p.Next()

	for _, callback := range callbacks {
		err := callback.Pre(&i, p.reg)
		if err != nil {
			return err
		}
	}

	p.reg.Ip++
	p.Exec(i)

	for _, callback := range callbacks {
		err := callback.Post(&i, p.reg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Program) Reg() Registers {
	return p.reg
}

func (p *Program) Reset() {
	p.reg = Registers{}
}

func (p *Program) Run(callbacks ...Callback) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%+v %w", p.reg, err)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	for err == nil && p.reg.Ip != len(p.instructions) {
		err = p.Step(callbacks...)
	}

	return err
}

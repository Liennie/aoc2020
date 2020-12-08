package asm

type Registers struct {
	Acc int
	Ip  int
}

type Program struct {
	reg          Registers
	instructions []Instruction
}

func (p *Program) Next() DebugInfo {
	return p.instructions[p.reg.Ip].Debug()
}

func (p *Program) Step() {
	i := p.instructions[p.reg.Ip]
	p.reg.Ip++
	i.Exec(p)
}

func (p *Program) Reg() Registers {
	return p.reg
}

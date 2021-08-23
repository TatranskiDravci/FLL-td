package main

import "time"

type Dropper struct {
	shifter 	Shifter
	id			int
	speed		int
}

func InitDropper(shifter Shifter, id int, speed int) Dropper {
	return Dropper {
		shifter : shifter,
		id : id,
		speed : speed,
	}
}

func (m Dropper) Release() {
	m.shifter.AwaitTo()
	if m.shifter.current != m.id {
		m.shifter.To(m.id)
	}
	m.shifter.Run(-300, m.speed)
	time.Sleep(time.Millisecond * 500)
}

type Lifter struct {
	shifter		Shifter
	id			int
	speed		int
	absPercent	int
}

func InitLifter(shifter Shifter, id, speed, pct int) Lifter {
	return Lifter {
		shifter : shifter,
		id : id,
		speed : speed,
		absPercent : pct,
	}
}

func (m *Lifter) To(target int) {
	m.shifter.AwaitTo()
	if m.shifter.current != m.id {
		m.shifter.To(m.id)
	}
	dPercent := target - m.absPercent
	m.shifter.Run(-57 * dPercent, m.speed)
	time.Sleep(time.Millisecond * 250)
	m.absPercent = target
}

func (m *Lifter) ToAsync(target int) {
	m.shifter.AwaitTo()
	if m.shifter.current != m.id {
		m.shifter.To(m.id)
	}
	dPercent := target - m.absPercent
	m.absPercent = target
	m.shifter.RunAsync(-57 * dPercent, m.speed)
}

func (m Lifter) AwaitTo() {
	m.shifter.AwaitRun()
}

type Column struct {
	shifter		Shifter
	id			int
	speed		int
	absPercent	int
}

func InitColumn(shifter Shifter, id int) Column {
	return Column {
		shifter : shifter,
		id : id,
	}
}

func (m Column) To(target int, speed int) {
	m.shifter.AwaitTo()
	if m.shifter.current != m.id {
		m.shifter.To(m.id)
	}
	m.shifter.Run(target, speed)
	time.Sleep(time.Millisecond * 250)
}

type Slide struct {
	shifter		Shifter
	id			int
	speed		int
}

func InitSlide(shifter Shifter, id int, speed int) Slide {
	return Slide {
		shifter : shifter,
		id : id,
		speed : speed,
	}
}

func (m Slide) Unlock() {
	m.shifter.AwaitTo()
	if m.shifter.current != m.id {
		m.shifter.To(m.id)
	}
	m.shifter.Run(-100, m.speed)
	time.Sleep(time.Millisecond * 250)
}

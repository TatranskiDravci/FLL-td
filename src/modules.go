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
	m.shifter.Run(-100, m.speed)
	time.Sleep(time.Millisecond * 500)
}
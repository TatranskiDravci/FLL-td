package main

type Lift struct {
	shifter		Shifter
	height		float64
	id			int
}

func InitLift(shifter Shifter, id int) Lift {
	liftr := Lift {
		shifter		: shifter,
		height		: 7.0,
		id			: id,
	}
	return liftr
}

type Carrier struct {
	shifter		Shifter
	contents	int
	id			int
}


func (m *Lift) To(target float64, speed int) {
	m.shifter.To(m.id)
	targetAngf := -(180.0*(target - 7.0))
	m.shifter.RunToAbs(int(targetAngf), speed)
	m.height = target
}
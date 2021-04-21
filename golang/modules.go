package main

// General module struct format:
//		type Module struct {
//			shifter		Shifter		<--- shifter it's connected to
//			id			int			<--- identifier of the given model on the shifter
// 			additional	...			<--- additional properties
// 		}


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

// target : cm (perpendicular distance of lifter pole from the ground)
// speed : motor units [unknown]
func (m *Lift) To(target float64, speed int) {
	m.shifter.To(m.id)
	targetAngf := -(180.0*(target - 7.0)) + (180.0*(m.height - 7.0))
	m.shifter.Run(int(targetAngf), speed)
	m.height = target
}


type Carrier struct {
	shifter		Shifter
	contents	int
	id			int
}

func InitCarrier(shifter Shifter, id int, contents int) Carrier {
	carrierr := Carrier {
		shifter		: shifter,
		contents	: contents,
		id			: id,
	}
	return carrierr
}

// count : unitless (number of bricks to release)
// speed : motor units [unknown]
func (m *Carrier) Release(count int, speed int) {
	m.shifter.To(m.id)
	targetAng := 0;
	if count <= m.contents {
		targetAng = -920*count
	} else {
		targetAng = -920*m.contents
	}
	m.shifter.Run(targetAng, speed)
	m.contents -= count
}

type Wheel struct {
	shifter 	Shifter
	id			int
}

func InitWheel(shifter Shifter, id int) Wheel {
	wheelr := Wheel {
		shifter		: shifter,
		id			: id,
	}
	return wheelr
}

// target : motor angle
// speed : motor units [unknown]
func (m Wheel) Run(target int, speed int) {
	m.shifter.To(m.id)
	m.shifter.Run(target, speed)
}
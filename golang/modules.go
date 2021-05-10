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
		height		: 3.5,
		id			: id,
	}
	return liftr
}

// speed : motor units [unknown]
func (m *Lift) To(target float64, speed int) {
	m.shifter.To(m.id)
	targetAngf := -540.0*target
	m.shifter.Run(int(targetAngf), speed)
	m.height = target
}

func (m *Lift) ToUnbounded(target float64, speed int) {
	m.shifter.To(m.id)
	targetAngf := -540.0*target
	m.shifter.RunUnbounded(int(targetAngf), speed)
	m.height = target
}


type Carrier struct {
	shifter		Shifter
	id			int
}

func InitCarrier(shifter Shifter, id int) Carrier {
	carrierr := Carrier {
		shifter		: shifter,
		id			: id,
	}
	return carrierr
}

// count : unitless (number of bricks to release)
// speed : motor units [unknown]
func (m *Carrier) Release(speed int) {
	m.shifter.To(m.id)
	m.shifter.Run(360, speed)
}

type Box struct {
	shifter Shifter
	id		int
}

func InitBox(shifter Shifter, id int) Box {
	boxr := Box {
		shifter		: shifter,
		id			: id,
	}
	return boxr
}

func (m *Box) Release(speed int) {
	m.shifter.To(m.id)
	m.shifter.Run(-150, speed)
}
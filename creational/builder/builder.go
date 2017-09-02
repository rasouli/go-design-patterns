package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

///////////////////////////////////////////////

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

////////////////////////////////////////////////

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

////////////////////////////////////////////////
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

///////////////////////////////////////////////
//////////////////////////////////////////////////////
type MotorBike struct {
	m VehicleProduct
}

func (c *MotorBike) SetWheels() BuildProcess {
	c.m.Wheels = 2
	return c
}

func (c *MotorBike) SetSeats() BuildProcess {
	c.m.Seats = 1
	return c
}

func (c *MotorBike) SetStructure() BuildProcess {
	c.m.Structure = "motorbike"
	return c
}

func (c *MotorBike) GetVehicle() VehicleProduct {
	return c.m
}

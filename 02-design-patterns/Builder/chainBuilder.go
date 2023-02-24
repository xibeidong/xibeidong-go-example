package Builder

//链式builder

//汽车由发动机、方向盘、车身构成，但这些部件也多种多样：

type Car struct {
	Engine   string //引擎
	Rack     string //架子
	Steering string //方向盘
}

type CarBuilder struct {
	Car
}

func (build *CarBuilder) EngineBuild(engine string) *CarBuilder {
	build.Engine = engine
	return build
}
func (build *CarBuilder) RackBuild(rack string) *CarBuilder {
	build.Rack = rack
	return build
}
func (build *CarBuilder) SteeringBuild(steering string) *CarBuilder {
	build.Steering = steering
	return build
}

func (build *CarBuilder) Build() Car {
	return Car{
		Engine:   build.Engine,
		Rack:     build.Rack,
		Steering: build.Steering,
	}
}

package main

type RoadLogistics struct {
	name string
}

func (r *RoadLogistics) setName(name string) {
	r.name = name
}

func (r *RoadLogistics) getName() string {
	return r.name
}

// var roadInstance ILogistics = (*RoadLogistics)(nil)  // 验证RoadLogistics是否实习了接口iLogistics
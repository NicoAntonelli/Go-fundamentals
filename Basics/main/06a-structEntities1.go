package main

type Student struct {
	name     string
	age      uint8
	subjects []string
	grades   []Grade
	University
}

func (s Student) getAverageGrade() float32 {
	var sum float32
	for _, grade := range s.grades {
		sum += float32(grade.value)
	}
	return sum / float32(len(s.grades))
}

type Grade struct {
	subject string
	value   uint8
}

type University struct {
	univName string
}

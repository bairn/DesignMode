package Observer

type Subject struct {
	obs []Observer
	context string
}

func NewSubject()*Subject {
	return &Subject{obs:make([]Observer, 0)}
}

func (s *Subject) Attch(o Observer) {
	s.obs = append(s.obs, o)
}

func (s *Subject) notify() {
	for _, o := range s.obs {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

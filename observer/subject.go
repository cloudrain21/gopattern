package observer

type Subject struct {
	observerList []Observer
}

func NewSubject() *Subject {
	return &Subject{
		observerList: make([]Observer,0),
	}
}

func (s *Subject)RegisterObserver(o Observer) {
	s.observerList = append(s.observerList, o)
}

func (s *Subject)notifyObservers(i int) {
	for _, o := range s.observerList {
		o.Notify(i)
	}
}

func (s *Subject)Run() {
	for i:=0; i<2; i++ {
		s.notifyObservers(i)
	}
}
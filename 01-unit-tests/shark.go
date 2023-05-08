package hunt

import (
	"errors"
)

type Shark struct {
	hungry bool
	tired  bool
	speed  int
}

type Prey struct {
	name  string
	speed int
}

var (
	ErrPreyNil        = errors.New("prey is nil")
	ErrSharkTired     = errors.New("cannot hunt, i am really tired")
	ErrSharkNotHungry = errors.New("cannot hunt, i am not hungry")
	ErrSharkNotCatch  = errors.New("could not catch it")
)

func (s *Shark) Hunt(p *Prey) error {
	if p == nil {
		return ErrPreyNil //fmt.Errorf("prey is nil")
	}
	if s.tired {
		return ErrSharkTired //fmt.Errorf("cannot hunt, i am really tired")
	}
	if !s.hungry {
		return ErrSharkNotHungry //fmt.Errorf("cannot hunt, i am not hungry")
	}
	if p.speed >= s.speed {
		s.tired = true
		return ErrSharkNotCatch //fmt.Errorf("could not catch it")
	}

	s.hungry = false
	s.tired = true
	return nil
}

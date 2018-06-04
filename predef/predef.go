package predef

type Tuple1 struct {
  _1 interface{}
}

type Tuple2 struct {
  _1, _2 interface{}
}

type Tuple3 struct {
  _1, _2, _3 interface{}
}

type Tuple4 struct {
  _1, _2, _3, _4 interface{}
}

type A interface{}
type B interface{}
type C interface{}
type T interface{}
type U interface{}
type V interface{}
type X interface{}
type Y interface{}
type Z interface{}

type Iterator struct {
  fmapThunk func(x interface{}, f interface{}) interface{}
  ch chan interface{}
}

func BuildIterator(f func(x interface{}, f interface{}) interface{}) Iterator {
  return Iterator{f, make(chan interface{})}
}

func (p Iterator) Filter(f interface{}, q Iterator) Iterator {
  go func() {
    for x := range p.ch {
      if p.fmapThunk(x, f).(bool) {
        q.ch <- x
      }
    }
    close(q.ch)
  } ()
  return q
}

func (p Iterator) Fmap(f interface{}, q Iterator) Iterator {
  go func() {
    for x := range p.ch {
      q.ch <- p.fmapThunk(x, f)
    }
    close(q.ch)
  } ()
  return q
}

func (p Iterator) Build(f func(ch chan interface{}) interface{} ) interface{} {
  return f(p.ch)
}

func (p Iterator) From(f func() chan interface{}) Iterator {
  p.ch = f()
  return p
}

func (p Iterator) Collect(f interface{}, q Iterator) Iterator {
  go func() {
    for x := range p.ch {
      if v := p.fmapThunk(x, f); v != nil {
        q.ch <- v
      }
    }
    close(q.ch)
  } ()
  return q
}


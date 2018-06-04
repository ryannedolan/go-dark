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
  elems []interface{}
}
func (p Iterator) next() interface{} {
  return p.elems[0]
}

func (p Iterator) hasNext() bool {
  return len(p.elems) > 0
}

func (p Iterator) push(x interface{}) {
  p.elems = append(p.elems, x)
}

func BuildIterator(f func(x interface{}, f interface{}) interface{}) Iterator {
  return Iterator{f, make([]interface{}, 0)}
}

func (p Iterator) Filter(f interface{}, q Iterator) Iterator {
  for p.hasNext() {
    x := p.next()
    if p.fmapThunk(x, f).(bool) {
      q.push(x)
    }
  }
  return q
}

func (p Iterator) Fmap(f interface{}, q Iterator) Iterator {
  for p.hasNext() {
    x := p.next()
    q.push(p.fmapThunk(x, f))
  }
  return q
}

func (p Iterator) Build(f func(ch []interface{}) interface{}) interface{} {
  return f(p.elems)
}

func (p Iterator) From(f func() []interface{}) Iterator {
  p.elems = f()
  return p
}

func (p Iterator) Collect(f interface{}, q Iterator) Iterator {
  for p.hasNext() {
    x := p.next()
    if v := p.fmapThunk(x, f); v != nil {
      q.push(v)
    }
  }
  return q
}


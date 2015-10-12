package filter

import (
	"fmt"
)

/*
path := id | id/path
field := path | path(fields)
fields := field | field,fields
fields' : = eps | fields
*/

type Tree struct {
	Children map[string]*Tree
}

type Path struct {
	Path []string
}

func Compile(str string) (t *Tree, e error) {
	t = &Tree{}
	p, e := t.compileFields_(str, 0)
	if e == nil && p != len(str) {
		e = fmt.Errorf("expect EOF at %v", p)
	}
	return
}

func (this *Tree) String() (s string) {
	f := true
	for k, v := range this.Children {
		if !f {
			s += ","
		}
		s += k
		if v.Children != nil {
			s += "("
			s += v.String()
			s += ")"
		}
		f = false
	}
	return
}

func (this *Tree) compileFields_(str string, pos int) (p int, e error) {
	p = pos
	if p == len(str) {
		return
	} else {
		p, e = this.compileFields(str, p)
		return
	}
}

func (this *Tree) compileFields(str string, pos int) (p int, e error) {
	p = pos
	for {
		p, e = this.compileField(str, p)
		if e != nil {
			break
		}
		if p < len(str) && str[p] == ',' {
			p++
		} else {
			break
		}
	}
	return
}

func (this *Tree) compileField(str string, pos int) (p int, e error) {
	p = pos
	var path Path
	p, e = path.compilePath(str, p)
	if e != nil {
		return
	}
	t := this
	for _, p := range path.Path {
		if t.Children == nil {
			t.Children = make(map[string]*Tree)
		}
		if _, ok := t.Children[p]; !ok {
			t.Children[p] = &Tree{}
		}
		t = t.Children[p]
	}
	if p == len(str) || str[p] != '(' {
		return
	} else {
		p++
	}
	p, e = t.compileFields(str, p)
	if e != nil {
		return
	}
	if p == len(str) || str[p] != ')' {
		e = fmt.Errorf("expect ) at %v", p)
		return
	} else {
		p++
		return
	}
}

func (this *Path) compilePath(str string, pos int) (p int, e error) {
	p = pos
	l := p
	for {
		for !(p == len(str) || str[p] == ',' || str[p] == '(' || str[p] == ')' || str[p] == '/') {
			p++
		}
		if l == p {
			e = fmt.Errorf("identifier is epsilon at %v", p)
			break
		}
		this.Path = append(this.Path, str[l:p])
		if p != len(str) && str[p] == '/' {
			p++
			l = p
		} else {
			break
		}
	}
	return
}

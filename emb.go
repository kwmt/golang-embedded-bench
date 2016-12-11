package embeded


import (
//	"fmt"
)

type RouterGroup struct {	
	root     bool
}

type Engine struct {
	RouterGroup
}

var b bool

func Embedded(engine *Engine) {
	
	b = engine.root
	if b {
		//
	}
}

func EmbeddedNo(engine *Engine) {
	
	b =  engine.RouterGroup.root
	if b {
		//
	}

}

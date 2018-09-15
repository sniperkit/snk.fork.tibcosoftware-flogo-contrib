/*
Sniperkit-Bot
- Status: analyzed
*/

package linker

import (
	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/core/mapper/exprmapper/expression/function"
)

type IsDefine struct {
}

func init() {
	function.Registry(&IsDefine{})
}

func (s *IsDefine) GetName() string {
	return "isDefined"
}

func (s *IsDefine) GetCategory() string {
	return ""
}

func (s *IsDefine) Eval(in interface{}) bool {
	return in != nil
}

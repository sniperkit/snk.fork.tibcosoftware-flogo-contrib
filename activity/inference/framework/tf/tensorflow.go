/*
Sniperkit-Bot
- Status: analyzed
*/

package tf

import (
	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-contrib/activity/inference/framework"
)

// TensorflowModel is
type TensorflowModel struct {
	frameworkTyp string
}

func init() {
	instance := new()
	framework.Register(instance)
}

func new() *TensorflowModel {
	return &TensorflowModel{frameworkTyp: "Tensorflow"}
}

func (a *TensorflowModel) FrameworkTyp() string {
	return a.frameworkTyp
}

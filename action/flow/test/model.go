/*
Sniperkit-Bot
- Status: analyzed
*/

package test

import (
	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-contrib/action/flow/model"
	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-contrib/action/flow/model/simple"
)

func init() {
	model.Register(NewTestModel())
}

func NewTestModel() *model.FlowModel {
	m := model.New("test")
	m.RegisterFlowBehavior(&simple.FlowBehavior{})
	m.RegisterDefaultTaskBehavior("basic", &simple.TaskBehavior{})

	return m
}

/*
Sniperkit-Bot
- Status: analyzed
*/

package simple

import (
	"testing"

	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-contrib/action/flow/model"
)

func TestRegistered(t *testing.T) {
	act := model.Get("tibco-simple")

	if act == nil {
		t.Error("Model Not Registered")
		t.Fail()
		return
	}
}

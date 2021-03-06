/*
Sniperkit-Bot
- Status: analyzed
*/

package channel

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/core/activity"
	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/engine/channels"

	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-contrib/action/flow/test"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	channels.New("test", 5)
	ch := channels.GetChannel("test")

	//setup attrs
	tc.SetSetting(sChannel, "test")
	tc.SetInput(ivValue, 2)

	done, err := act.Eval(tc)

	if !done {
		t.Error("activity should be done")
		return
	}

	if err != nil {
		t.Error("activity has an error: ", err)
		return
	}

	expected := 2
	var found interface{}

	ch.RegisterCallback(func(msg interface{}) {
		found = msg
	})

	channels.Start()

	time.Sleep(100 * time.Millisecond)

	if found != expected {
		t.Errorf("Expected %s, found %s", expected, found)
	}

	channels.Close()
}

/*
Sniperkit-Bot
- Status: analyzed
*/

package gpio

import (
	"io/ioutil"
	"testing"

	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/core/activity"

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

func TestReadState(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("method", "Read State")
	tc.SetInput("pin number", 10)
	//eval
	_, err := act.Eval(tc)
	if err != nil {
		log.Errorf("Error occured: %+v", err)
	}
	val := tc.GetOutput("result")
	log.Debugf("Resut %s", val)

}

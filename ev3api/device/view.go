package device
/*
import (
	"github.com/fuzzycow/ev32go/ev3api"
	"encoding/json"
)

type View struct {
	ev3api.Device
	Ravp map[string]interface{}
	Wavp map[string]interface{}
	attrs []string
}

func NewView(R,W []string) *View {
	return &View{
		R: R,
		W: W}
}

func (view View) Marshall() error {
	avps := make(map[string]interface{})
	for _,a := range view.attrs {
		avps[a] = view.GetAttrString(a)
	}
}


func (view View) Unmarshal(data []byte, v interface{}) {
	avps := make(map[string]interface{})
	json.Unmarshal(avps)
	for _,a := range view.R {
		avps[a] = json.Unmarshal()
	}
}
*/
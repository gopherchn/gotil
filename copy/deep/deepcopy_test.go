package deep

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

type Student struct {
	Name     string
	Age      int
	Birthday time.Time
}

var s = Student{
	Name:     "abc",
	Age:      12,
	Birthday: time.Now(),
}
// 


func Test_deepcopyWithMarshal(t *testing.T) {
	s1 := Student{}
	b, err := json.Marshal(s)
	if err != nil {
		t.Error(err)
	}
	if err := json.Unmarshal(b, s1); err != nil {
		t.Error(err)
	}
	t.Logf("%+v, %+v", s, s1)
}

func Test_DeepCopyWithGob(t *testing.T) {
	s1 := Student{}
	DeepcopyWithGob(&s1, &s)
	t.Logf("%+v, \n %+v", s, s1)
	if(!reflect.DeepEqual(s1, s)) {
		t.Error("dont't deep equal: ", s1, s)
	}
}

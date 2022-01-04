package deep

import (
	"bytes"
	"encoding/gob"
)

func GobDeepCopy(dst, src interface{}) error {
    var buf bytes.Buffer
    if err := gob.NewEncoder(&buf).Encode(src); err != nil {
        return err
    }
    return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
 
// DeepcopyWithGob Deepcopy with gob 
func DeepcopyWithGob(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	} 
	// return gob.NewDecoder(&buf).Decode(dst)
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func DeepCopyWithMarshal(dst, src interface{}) error {
	return nil
}

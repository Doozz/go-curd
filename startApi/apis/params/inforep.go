package params

import "encoding/json"


type InfoRep struct {
    Code int
    Data  interface{}
}

func (in *InfoRep) Result(code int, msg interface{}) ([]byte, error) {
    in.Code = code
    in.Data = msg

    return json.Marshal(in)
}
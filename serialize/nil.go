package serialize

import "github.com/shamaton/msgpack/def"

func (s *serializer) writeNil(offset int) (int, error) {
	offset = s.setByte1Int(def.Nil, offset)
	return offset, nil
}

package light

import (
	"bytes"
	"encoding/json"

	"ilysa/pkg/beatsaber"
)

type ID []int

func NewID(ids ...int) ID {
	return append(ID{}, ids...)
}

func (id *ID) UnmarshalJSON(raw []byte) error {
	if bytes.ContainsRune(raw, '[') {
		return json.Unmarshal(raw, &id)
	}
	var singleLightID int
	if err := json.Unmarshal(raw, &singleLightID); err != nil {
		return err
	}

	*id = ID{singleLightID}
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	if len(id) == 1 {
		return json.Marshal(id[0])
	}
	return json.Marshal([]int(id))
}

type Light interface {
	EventType(index int) beatsaber.EventType
	LightID(i int) ID
	LightIDMin() int
	LightIDMax() int
	LightIDLen() int
}

type Instance struct {

	eventType beatsaber.EventType
	id        ID
}

type Basic struct {
	eventType beatsaber.EventType
	id        ID
}

func (l Basic) EventType(index int) beatsaber.EventType {
	return l.eventType
}

func (l Basic) LightID(i int) ID {
	return ID{l.id[i-1]}
}

func (l Basic) LightIDMin() int {
	return 1
}

func (l Basic) LightIDMax() int {
	return len(l.id)
}

func (l Basic) LightIDLen() int {
	return len(l.id)
}

type SetBuilder func(ID) *Set

func Identity(id ID) *Set {
	return NewSet(id)
}

func DivideSingle(id ID) *Set {
	set := NewSet()
	for _, lightID := range id {
		set.Add(ID{lightID})
	}
	return set
}

func Divide(divisor int) SetBuilder {
	return func(id ID) *Set {
		set := NewSet()

		for len(id) > divisor {
			set.Add(id[0:divisor])
			id = id[divisor:]
		}

		set.AppendToIndex(set.Len()-1, id...)

		return set
	}
}

func Fan(groupCount int) SetBuilder {
	return func(id ID) *Set {
		set := make(Set, groupCount)

		for i, lightID := range id {
			set[i%groupCount] = append(set[i%groupCount], lightID)
		}

		return &set
	}
}

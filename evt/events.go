package evt

type Option interface {
}

type Events []Event

func (events Events) Apply(opts ...Option) {

}

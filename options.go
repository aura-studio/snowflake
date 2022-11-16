package snowflake

const (
	WaitMethodSleep = iota
	WaitMethodSpin
)

type Options struct {
	pattern    *Pattern
	safe       bool
	waitMethod int
}

func (o *Options) Apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

var defaultOptions = Options{
	pattern:    DefaultPattern,
	safe:       true,
	waitMethod: WaitMethodSleep,
}

type Option func(*Options)

func WithPattern(p *Pattern) Option {
	return func(o *Options) {
		o.pattern = p
	}
}

func WithSafe(safe bool) Option {
	return func(o *Options) {
		o.safe = safe
	}
}

func WithWaitMethod(method int) Option {
	return func(o *Options) {
		o.waitMethod = method
	}
}

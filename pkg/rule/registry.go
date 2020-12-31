package rule

type Register interface {
	Add(def Definition)
}

type RegisterLister interface {
	List() []Definition
}

type registry struct {
	registeredRules []Definition
}

func (r *registry) Add(def Definition) {
	r.registeredRules = append(r.registeredRules, def)
}

func (r *registry) List() []Definition {
	return r.registeredRules
}

var globalRegistry registry

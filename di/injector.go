package di

type Injector struct {
	Dependencies []Dependencies
}

func (i *Injector) Inject(modules []Module) []Module {
	for _, module := range modules {
		module.Init(i.Dependencies)
	}

	return modules
}

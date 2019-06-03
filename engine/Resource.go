package travian

type Resource struct {
	Lumber int
	Clay   int
	Iron   int
	Crop   int
}

func (a *Village) Has(resource Resource) bool {

	return a.resource.Lumber >= resource.Lumber &&
		a.resource.Clay >= resource.Clay &&
		a.resource.Iron >= resource.Iron &&
		a.resource.Crop >= resource.Crop
}

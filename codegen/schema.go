package codegen

type SystemProperty struct {
	Description []string `json:"description"`
	Name        string   `json:"name"`
	ReadAccess  bool     `json:"readAccess"`
	SystemName  string   `json:"systemName"`
	Type        string   `json:"type"`
	WriteAccess bool     `json:"writeAccess"`
}

type PropertyValue struct {
	PropertyName string `json:"propertyName"`
	Values       []struct {
		Description []string `json:"description"`
		Name        string   `json:"name"`
	} `json:"values"`
}

type Device struct {
	Description                []string `json:"description"`
	DocsLink                   string   `json:"docsLink"`
	FriendlyName               string   `json:"friendlyName"`
	SystemClassName            string `json:"systemClassName"`
	SystemDeviceNameConvention string `json:"systemDeviceNameConvention"`
	Inheritance					[]string `json:"inheritance"`
	SystemProperties           []*SystemProperty `json:"systemProperties"`
	PropertyValues             []*PropertyValue `json:"propertyValues"`
}

type Spec struct {
	Classes map[string]*Device `json:"classes"`
	Meta struct {
				SpecRevision    int    `json:"specRevision"`
				SupportedKernel string `json:"supportedKernel"`
				Version         string `json:"version"`
			} `json:"meta"`
}


package kinggen

type Endpoint int

const (
	ProfileEndpoint = iota
	AltEndpoint
)

func (endpoint Endpoint) Build(key string) string {
	return "https://kinggen.wtf/api/v2/" + endpoint.ToString() + "?key=" + key
}

func (endpoint Endpoint) ToString() string {
	return []string{"profile", "alt"}[endpoint]
}

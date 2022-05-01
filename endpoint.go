package kinggen

type Endpoint int

const (
	profileEndpoint = iota
	altEndpoint
)

func (endpoint Endpoint) build(key string) string {
	return "https://kinggen.wtf/api/v2/" + endpoint.toString() + "?key=" + key
}

func (endpoint Endpoint) toString() string {
	return []string{"profile", "alt"}[endpoint]
}

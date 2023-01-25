package util

import "net/url"

func BuildBody(mapInfo map[string]string) string {
	u, _ := url.Parse("")
	query := u.Query()
	for k, v := range mapInfo {
		query.Add(k, v)
	}
	u.RawQuery = query.Encode()
	return u.String()[1:]
}

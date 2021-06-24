package sw

import (
	"fmt"
	"strings"
)

type ScrutRepo struct {
	Provider string
	Owner    string
	Name     string
}

func (sR ScrutRepo) FromString(s string) (ScrutRepo, error) {
	s = strings.TrimPrefix(s, "/")
	s = strings.TrimSuffix(s, "/")
	sz := strings.Split(s, "/")

	if len(sz) != 3 {
		return ScrutRepo{}, fmt.Errorf("specified string is not a valid scrutinizer repository string")
	}

	sR.Provider = sz[0]
	sR.Owner = sz[1]
	sR.Name = sz[2]

	return sR, nil
}

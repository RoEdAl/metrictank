package util

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParseIngestFromFlags(ingestFromStr string) (map[uint32]int64, error) {
	if len(ingestFromStr) == 0 {
		return nil, nil
	}

	ingestFrom := make(map[uint32]int64)

	ingestFromStrPerOrgs := strings.Split(ingestFromStr, ",")
	for _, ingestFromStrPerOrg := range ingestFromStrPerOrgs {

		if len(ingestFromStr) == 0 {
			continue
		}
		parts := strings.Split(ingestFromStrPerOrg, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("could not parse section %q from %q", ingestFromStrPerOrg, ingestFromStr)
		}
		orgID, err := strconv.ParseUint(parts[0], 0, 32)
		if err != nil {
			return nil, fmt.Errorf("could not parse org id %q: %s", parts[0], err.Error())
		}
		timestamp, err := strconv.ParseInt(parts[1], 0, 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse timestamp %q: %s", parts[1], err.Error())
		}
		if timestamp <= 0 {
			return nil, errors.New("timestamp must be > 0")
		}

		ingestFrom[uint32(orgID)] = timestamp
	}
	return ingestFrom, nil
}

package tag

import (
	"github.com/skycoin/bbs/src/misc/boo"
	"strconv"
	"strings"
)

/*
	<<< CHECKING FUNCTIONS >>>
*/

// CheckSeed ensures validity of seed. TODO
func CheckSeed(seed string) error {
	return nil
}

// CheckName ensures validity of board/thread/post name. TODO
func CheckName(name string) error {
	return nil
}

// CheckBody ensures validity of board/thread/post description. TODO
func CheckBody(body string) error {
	return nil
}

// CheckAddress ensures validity of address. TODO
func CheckAddress(address string) error {
	pts := strings.Split(address, ":")
	port, err := strconv.ParseUint(pts[len(pts)-1], 10, 16)
	if err != nil {
		return boo.Newf(boo.InvalidInput,
			"address '%s' is invalid", string(address))
	}
	if port < 0 || port > 65535 {
		return boo.Newf(boo.InvalidInput,
			"address:%s port is invalid ", string(address))
	}
	return nil
}

// CheckMode check's the vote's mode.
func CheckMode(mode int8) error {
	switch mode {
	case -1, 0, +1:
		return nil
	default:
		return boo.Newf(boo.InvalidInput,
			"invalid vote mode of %d provided", mode)
	}
}


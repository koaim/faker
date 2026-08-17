package faker

import (
	"math"
	"math/rand/v2"
	"net"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func fillURL(rv reflect.Value, opt Option) {
	u := url.URL{
		Scheme: randIn(opt.Url.AvailableSchemes...),
		Host:   randStr(opt.AllowedRunes, opt.StrLen) + "." + randIn(opt.Url.AvailableDomainZones...),
		Path:   randStr(opt.AllowedRunes, opt.Url.PathLen),
	}

	var sb strings.Builder
	for i := range opt.Url.QueryCount {
		sb.WriteString(randStr(opt.AllowedRunes, 5))
		sb.WriteString("=")
		sb.WriteString(randStr(opt.AllowedRunes, 3))

		if i != opt.Url.QueryCount-1 {
			sb.WriteString("&")
		}
	}
	if sb.Len() != 0 {
		u.RawQuery = sb.String()
	}

	if opt.Url.WithUserInfo {
		u.User = url.UserPassword(randStr(opt.AllowedRunes, 5), randStr(opt.AllowedRunes, 5))
	}

	rv.Set(reflect.ValueOf(u))
}

func fillIP(rv reflect.Value) {
	ip := net.IPv4(
		byte(randInRange(1, math.MaxUint8)),
		byte(randInRange(1, math.MaxUint8)),
		byte(randInRange(1, math.MaxUint8)),
		byte(randInRange(1, math.MaxUint8)),
	)

	rv.Set(reflect.ValueOf(ip))
}

func fillIPNet(rv reflect.Value) {
}

func fillIPMask(rv reflect.Value) {
	mask := net.IPv4Mask(
		byte(randInRange(1, math.MaxUint8)),
		byte(randInRange(1, math.MaxUint8)),
		byte(randInRange(1, math.MaxUint8)),
		byte(randInRange(1, math.MaxUint8)),
	)

	rv.Set(reflect.ValueOf(mask))
}

func fillTime(rv reflect.Value) {
	minUnix := time.Now().AddDate(-yearsForward, 0, 0).Unix()
	maxUnix := time.Now().AddDate(yearsForward, 0, 0).Unix()
	sec := minUnix + rand.Int64N(maxUnix-minUnix+1)

	rv.Set(reflect.ValueOf(time.Unix(sec, rand.Int64N(1e9))))
}

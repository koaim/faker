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
	ip := createIP()
	mask := net.IPv4(ip[0], ip[1], ip[2], ip[3])

	rv.Set(reflect.ValueOf(mask))
}

func fillIPNet(rv reflect.Value) {
	ip := createIP()
	mask := createIP()

	ipNet := net.IPNet{
		IP:   net.IPv4(ip[0], ip[1], ip[2], ip[3]),
		Mask: net.IPv4Mask(mask[0], mask[1], mask[2], mask[3]),
	}

	rv.Set(reflect.ValueOf(ipNet))
}

func fillIPMask(rv reflect.Value) {
	ip := createIP()
	mask := net.IPv4Mask(ip[0], ip[1], ip[2], ip[3])

	rv.Set(reflect.ValueOf(mask))
}

func fillTime(rv reflect.Value) {
	minUnix := time.Now().AddDate(-yearsForward, 0, 0).Unix()
	maxUnix := time.Now().AddDate(yearsForward, 0, 0).Unix()
	sec := minUnix + rand.Int64N(maxUnix-minUnix+1)

	rv.Set(reflect.ValueOf(time.Unix(sec, rand.Int64N(1e9))))
}

func createIP() []byte {
	var res []byte
	for range 4 {
		res = append(res, byte(randInRange(1, math.MaxUint8)))
	}

	return res
}

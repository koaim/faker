package faker

import (
	"math"
	"math/rand/v2"
	"net"
	"net/url"
	"reflect"
	"time"
)

func fillURL(rv reflect.Value, opt Option) {
	u := url.URL{
		Scheme: randIn(opt.Url.AllowedSchemes...),
		Host:   randStr(opt.AllowedRunes, opt.StrLen) + "." + randIn(opt.Url.AllowedDomainZones...),
		Path:   randStr(opt.AllowedRunes, opt.Url.PathLen),
	}

	query := make(url.Values)
	for range opt.Url.QueryCount {
		query.Add(randStr(opt.AllowedRunes, 5), randStr(opt.AllowedRunes, 3))
	}
	u.RawQuery = query.Encode()

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
	mask := net.CIDRMask(rand.IntN(33), 32)
	address := net.IPv4(ip[0], ip[1], ip[2], ip[3])

	ipNet := net.IPNet{
		IP:   address.Mask(mask),
		Mask: mask,
	}

	rv.Set(reflect.ValueOf(ipNet))
}

func fillIPMask(rv reflect.Value) {
	mask := net.CIDRMask(rand.IntN(33), 32)

	rv.Set(reflect.ValueOf(mask))
}

func fillHardwareAddr(rv reflect.Value) {
	var mac []byte
	for range 6 {
		mac = append(mac, byte(randInRange(0, math.MaxUint8)))
	}

	rv.Set(reflect.ValueOf(net.HardwareAddr(mac)))
}

func fillTime(rv reflect.Value, opt Option) {
	now := time.Now()
	minUnix := now.AddDate(-opt.Time.YearsBack, 0, 0).Unix()
	maxUnix := now.AddDate(opt.Time.YearsForward, 0, 0).Unix()
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

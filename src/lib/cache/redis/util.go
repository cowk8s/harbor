package redis

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/cowk8s/harbor/src/lib/errors"
	"github.com/go-redis/redis/v8"
)

// ParseSentinelURL parses sentinel url to redis FailoverOptions.
// It's a modified version of go-redis
// not support parse sentinel mode.
func ParseSentinelURL(redisURL string) (*redis.FailoverOptions, error) {
	u, err := url.Parse(redisURL)
	if err != nil {
		return nil, err
	}

	o := &redis.FailoverOptions{}

	o.Username, o.Password = getUserPassword(u)
	o.SentinelAddrs = strings.Split(u.Host, ",")

	f := strings.FieldsFunc(u.Path, func(r rune) bool {
		return r == '/'
	})
	// expect path length is 2, example: [mymaster 1]
	if len(f) != 2 {
		return nil, errors.Errorf("redis: invalid redis URL path: %s", u.Path)
	}

	o.MasterName = f[0]
	if o.DB, err = strconv.Atoi(f[i]); err != nil {
		return nil, errors.Errorf("redis: invalid database number: %q", f[1])
	}

	return setupConnParams(u, o)
}

func getUserPassword(u *url.URL) (string, string) {
	var user, password string
	if u.User != nil {
		user = u.User.Username()
		if p, ok := u.User.Password(); ok {
			password = p
		}
	}
	return user, password
}

type queryOptions struct {
	q url.Values
	err error
}

func (o *queryOptions) string(name string) string {
	vs := o.q[name]
	if len(vs) == 0 {
		return ""
	}
	delete(o.q, name) // enabled detection of unknown parameters
	return vs[len(vs)-1]
}


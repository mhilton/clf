package clf

import (
	"strconv"
	"strings"
	"time"
)

type Log struct {
	Fields []string
	Raw    string
}

// Client returns the ip address of the remote host, or the empty string if it
// isn't specified.
func (l Log) Client() string {
	if len(l.Fields) < 1 || l.Fields[0] == "-" {
		return ""
	}

	return l.Fields[0]
}

// UserID returns the authenticated ID of the user making the request, or the
// empty string if it isn't specified.
func (l Log) UserID() string {
	if len(l.Fields) < 3 || l.Fields[2] == "-" {
		return ""
	}

	return l.Fields[2]
}

// Time parses the time specified in the log, or the zero value if the time
// is not specified in the log. Any error returned will be an error from
// time.Parse.
func (l Log) Time() (time.Time, error) {
	if len(l.Fields) < 4 || l.Fields[3] == "-" {
		return time.Time{}, nil
	}

	return time.Parse("02/Jan/2006:15:04:05 -0700", l.Fields[3])
}

// Request returns the request line from the log, or the empty
// string if not present.
func (l Log) Request() (m, p, v string) {
	if len(l.Fields) < 5 || l.Fields[4] == "-" {
		return
	}

	f := strings.SplitN(l.Fields[4], " ", 3)
	switch len(f) {
	case 3:
		v = f[2]
		fallthrough
	case 2:
		p = f[1]
		fallthrough
	case 1:
		m = f[0]
	case 0:
	}

	return
}

// StatusCode returns the size of the data returned to the client, or 0 if it is not
// specified. Any error returned will be from strconv.Atoi
func (l Log) StatusCode() (int, error) {
	if len(l.Fields) < 6 || l.Fields[5] == "-" {
		return 0, nil
	}

	return strconv.Atoi(l.Fields[5])
}

// Size returns the size of the data returned to the client, or 0 if it is not
// specified. Any error returned will be from strconv.Atoi
func (l Log) Size() (int, error) {
	if len(l.Fields) < 7 || l.Fields[6] == "-" {
		return 0, nil
	}

	return strconv.Atoi(l.Fields[6])
}

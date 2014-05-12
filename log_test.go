package clf

import (
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	var l Log

	l.Fields = []string{"192.168.0.1"}

	c := l.Client()

	if c != "192.168.0.1" {
		t.Errorf("Incorrect client address returned: expected \"192.168.0.1\" but got \"%s\".", c)
	}
}

func TestClientEmpty(t *testing.T) {
	var l Log

	l.Fields = []string{}

	c := l.Client()

	if c != "" {
		t.Errorf("Incorrect client address returned: expected \"\" but got \"%s\".", c)
	}
}

func TestClientHyphen(t *testing.T) {
	var l Log

	l.Fields = []string{"-"}

	c := l.Client()

	if c != "" {
		t.Errorf("Incorrect client address returned: expected \"\" but got \"%s\".", c)
	}
}

func TestUserID(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "martin"}

	u := l.UserID()

	if u != "martin" {
		t.Errorf("Incorrect user ID returned: expected \"martin\" but got \"%s\".", u)
	}
}

func TestUserIDEmpty(t *testing.T) {
	var l Log

	l.Fields = []string{}

	u := l.UserID()

	if u != "" {
		t.Errorf("Incorrect user id returned: expected \"\" but got \"%s\".", u)
	}
}

func TestUserIDHyphen(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "-"}

	u := l.UserID()

	if u != "" {
		t.Errorf("Incorrect user id returned: expected \"\" but got \"%s\".", u)
	}
}

func TestTime(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "USERID", "25/Dec/2014:01:00:00 +0000"}

	tm, err := l.Time()
	if err != nil {
		t.Fatalf("Unexpected error %s when parsing time", err)
	}

	texp := time.Date(2014, time.December, 25, 1, 0, 0, 0, time.UTC)
	if !tm.Equal(texp) {
		t.Errorf("Incorrect time returned: expected \"%s\" but got \"%s\".", texp, tm)
	}
}

func TestTimeParseError(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "USERID", "25/12/2014:01:00:00 +0000"}

	tm, err := l.Time()
	if err == nil {
		t.Fatalf("Time parsed as %s should have produced an error", tm)
	}
}

func TestTimeEmpty(t *testing.T) {
	var l Log

	l.Fields = []string{}

	tm, err := l.Time()
	if err != nil {
		t.Fatalf("Unexpected error %s when parsing time", err)
	}

	texp := time.Time{}
	if !tm.Equal(texp) {
		t.Errorf("Incorrect time returned: expected \"%s\" but got \"%s\".", texp, tm)
	}

}

func TestTimeHyphen(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "USERID", "-"}

	tm, err := l.Time()
	if err != nil {
		t.Fatalf("Unexpected error %s when parsing time", err)
	}

	texp := time.Time{}
	if !tm.Equal(texp) {
		t.Errorf("Incorrect time returned: expected \"%s\" but got \"%s\".", texp, tm)
	}
}

func TestRequest(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "USERID", "TIME", "GET / HTTP/1.0"}

	m, p, v := l.Request()

	if m != "GET" {
		t.Errorf("Incorrect method returned: expected \"GET\" but got \"%s\".", m)
	}
	if p != "/" {
		t.Errorf("Incorrect path returned: expected \"/\" but got \"%s\".", p)
	}
	if v != "HTTP/1.0" {
		t.Errorf("Incorrect version returned: expected \"HTTP/1.0\" but got \"%s\".", v)
	}
}

func TestRequestEmpty(t *testing.T) {
	var l Log

	l.Fields = []string{}

	m, p, v := l.Request()

	if m != "" {
		t.Errorf("Incorrect method returned: expected \"\" but got \"%s\".", m)
	}
	if p != "" {
		t.Errorf("Incorrect path returned: expected \"\" but got \"%s\".", p)
	}
	if v != "" {
		t.Errorf("Incorrect version returned: expected \"\" but got \"%s\".", v)
	}

}

func TestRequestHyphen(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "USERID", "TIME", "-"}

	m, p, v := l.Request()

	if m != "" {
		t.Errorf("Incorrect method returned: expected \"\" but got \"%s\".", m)
	}
	if p != "" {
		t.Errorf("Incorrect path returned: expected \"\" but got \"%s\".", p)
	}
	if v != "" {
		t.Errorf("Incorrect version returned: expected \"\" but got \"%s\".", v)
	}
}

func TestStatusCode(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "USERID", "TIME", "REQUEST", "200"}

	sc, err := l.StatusCode()
	if err != nil {
		t.Fatalf("Unexpected error %s when parsing status code", err)
	}

	if sc != 200 {
		t.Errorf("Incorrect status code returned: expected 200 but got %d.", sc)
	}
}

func TestStatusCodeParseError(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "USERID", "TIME", "REQUEST", "SC"}

	sc, err := l.StatusCode()
	if err == nil {
		t.Fatalf("Status code parsed as %d should have produced an error", sc)
	}
}

func TestStatusCodeEmpty(t *testing.T) {
	var l Log

	l.Fields = []string{}

	sc, err := l.StatusCode()
	if err != nil {
		t.Fatalf("Unexpected error %s when parsing status code", err)
	}

	if sc != 0 {
		t.Errorf("Incorrect status code returned: expected 0 but got %d.", sc)
	}
}

func TestStatusCodeHyphen(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "USERID", "TIME", "REQUEST", "-"}

	sc, err := l.StatusCode()
	if err != nil {
		t.Fatalf("Unexpected error %s when parsing status code", err)
	}

	if sc != 0 {
		t.Errorf("Incorrect status code returned: expected 0 but got %d.", sc)
	}
}

func TestSize(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "USERID", "TIME", "REQUEST", "STATUS", "1234"}

	s, err := l.Size()
	if err != nil {
		t.Fatalf("Unexpected error %s when parsing size", err)
	}

	if s != 1234 {
		t.Errorf("Incorrect size returned: expected 1234 but got %d.", s)
	}
}

func TestSizeParseError(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "USERID", "TIME", "REQUEST", "STATUS", "SIZE"}

	s, err := l.Size()
	if err == nil {
		t.Fatalf("Size parsed as %d should have produced an error", s)
	}
}

func TestSizeEmpty(t *testing.T) {
	var l Log

	l.Fields = []string{}

	s, err := l.Size()
	if err != nil {
		t.Fatalf("Unexpected error %s when parsing size", err)
	}

	if s != 0 {
		t.Errorf("Incorrect size returned: expected 0 but got %d.", s)
	}
}

func TestSizeHyphen(t *testing.T) {
	var l Log

	l.Fields = []string{"CLIENT", "IDENT", "USERID", "TIME", "REQUEST", "STATUS", "-"}

	s, err := l.Size()
	if err != nil {
		t.Fatalf("Unexpected error %s when parsing status code", err)
	}

	if s != 0 {
		t.Errorf("Incorrect size returned: expected 0 but got %d.", s)
	}
}

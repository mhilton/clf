package clf

import (
	"testing"
)

type scanTest struct {
	raw    string
	fields []string
}

var scanTests []scanTest = []scanTest{
	{"", []string{}},
	{"- - - - - - -", []string{"-", "-", "-", "-", "-", "-", "-"}},
	{"- - - - - - -\n", []string{"-", "-", "-", "-", "-", "-", "-"}},
	{"127.0.0.1 - jdoe [25/Dec/2013:07:00:00 +0000] \"GET / HTTP/1.1\" 200 0",
		[]string{"127.0.0.1", "-", "jdoe", "25/Dec/2013:07:00:00 +0000", "GET / HTTP/1.1", "200", "0"}},
	{"127.0.0.1 - jdoe [25/Dec/2013:07:00:00 +0000\n",
		[]string{"127.0.0.1", "-", "jdoe", "25/Dec/2013:07:00:00 +0000"}},
	{"127.0.0.1 - jdoe [25/Dec/2013:07:00:00 +0000] \"GET / HTTP/1.1\n",
		[]string{"127.0.0.1", "-", "jdoe", "25/Dec/2013:07:00:00 +0000", "GET / HTTP/1.1"}},
}

func TestScan(t *testing.T) {
	for n, st := range scanTests {
		var l Log
		l.Raw = st.raw

		scan(&l)

		if len(l.Fields) != len(st.fields) {
			t.Errorf("Scan test %d: expected %d fields but got %d", n, len(st.fields), len(l.Fields))
		}

		for i, f := range st.fields {
			if i < len(l.Fields) {
				if f != l.Fields[i] {
					t.Errorf("Scan Test %d: field %d expected \"%s\" but got \"%s\"", n, i, f, l.Fields[i])
				}
			}
		}
	}
}

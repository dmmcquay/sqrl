package sqrl

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCPU(t *testing.T) {
	fmt.Println("Testing CPU function...")
	c, err := GetCPUInfo()
	if err != nil {
		t.Fatalf("Could not retrieve CPU info.")
	}
	c.LogWriter()
	fmt.Println(c)
}

func TestRam(t *testing.T) {
	fmt.Println("Testing Ram function...")
	r, _, _, err := RamOSInfo()
	if err != nil {
		t.Fatalf("Could not retrieve RAM info.")
	}
	r.LogWriter()
	fmt.Println(r)
}

func TestOS(t *testing.T) {
	fmt.Println("Testing OS function...")
	_, o, _, err := RamOSInfo()
	if err != nil {
		t.Fatalf("Could not retrieve OS info.")
	}
	o.LogWriter()
	fmt.Println(o)
}

func TestSwap(t *testing.T) {
	fmt.Println("Testing Swap function...")
	_, _, s, err := RamOSInfo()
	if err != nil {
		t.Fatalf("Could not retrieve swap info.")
	}
	s.LogWriter()
	fmt.Println(s)
}

func TestNet(t *testing.T) {
	fmt.Println("Testing Network function...")
	i, err := GetInterfaces()
	if err != nil {
		t.Fatalf("Could not retrieve network info.")
	}
	i.LogWriter()
	fmt.Println(i)
}

func TestDaemonAll(t *testing.T) {
	fmt.Println("Testing all route...")
	sm := http.NewServeMux()
	AddRoutes(sm)
	ts := httptest.NewServer(sm)

	u := fmt.Sprintf("%s%s", ts.URL, prefix["all"])
	fmt.Println(u)
}

func TestDaemonCom(t *testing.T) {
	fmt.Println("Testing com route...")
	sm := http.NewServeMux()
	AddRoutes(sm)
	ts := httptest.NewServer(sm)

	u := fmt.Sprintf("%s%s", ts.URL, prefix["com"])
	fmt.Println(u)
}

func TestDaemonRAM(t *testing.T) {
	fmt.Println("Testing ram route...")
	sm := http.NewServeMux()
	AddRoutes(sm)
	ts := httptest.NewServer(sm)

	u := fmt.Sprintf("%s%s", ts.URL, prefix["ram"])
	fmt.Println(u)
}

func TestDaemonOS(t *testing.T) {
	fmt.Println("Testing os route...")
	sm := http.NewServeMux()
	AddRoutes(sm)
	ts := httptest.NewServer(sm)

	u := fmt.Sprintf("%s%s", ts.URL, prefix["ops"])
	fmt.Println(u)
}

func TestDaemonSwap(t *testing.T) {
	fmt.Println("Testing swap route...")
	sm := http.NewServeMux()
	AddRoutes(sm)
	ts := httptest.NewServer(sm)

	u := fmt.Sprintf("%s%s", ts.URL, prefix["swap"])
	fmt.Println(u)
}

func TestDaemonNet(t *testing.T) {
	fmt.Println("Testing network route...")
	sm := http.NewServeMux()
	AddRoutes(sm)
	ts := httptest.NewServer(sm)

	u := fmt.Sprintf("%s%s", ts.URL, prefix["network"])
	fmt.Println(u)
}

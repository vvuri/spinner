package profibook

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func TestHttpHello(t *testing.T) {
	http.HandleFunc("/hello", hello)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/hello", nil)
	http.DefaultServeMux.ServeHTTP(writer, request)
	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
	if expected, actual := "Hello World!", writer.Body.String(); expected != actual {
		t.Errorf("Response body is %v", actual)
	}
}

// WEB SRV -------------------------------------------------------

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func TestHttpIndex(t *testing.T) {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8000", nil)
}

// NET Interfaces---------------------------------------------------

func TestGetInterface(t *testing.T) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n", i.Name)

		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}

		addresses, err := byName.Addrs()
		if err != nil {
			fmt.Println(err)
		}

		for k, v := range addresses {
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}
		fmt.Println()
	}
}

func TestGetInterfaceInfo(t *testing.T) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Name: %v\n", i.Name)
		fmt.Println("Interface Flags:", i.Flags.String())
		fmt.Println("Interface MTU:", i.MTU)
		fmt.Println("Interface Hardware Address:", i.HardwareAddr)
		fmt.Println()
	}
}

// DNS --------------------------------------------------------

func lookIP(address string) ([]string, error) {
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}
func lookHostname(hostname string) ([]string, error) {
	IPs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	return IPs, nil
}

func TestDNSlookup(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{"yandex.ru"},
		{"8.8.8.8"},
		{"127.0.0.1"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			input := testCase.name

			IPaddress := net.ParseIP(input)

			if IPaddress == nil {
				IPs, err := lookHostname(input)
				if err == nil {
					for _, singleIP := range IPs {
						fmt.Println(input, singleIP)
					}
				}
			} else {
				hosts, err := lookIP(input)
				if err == nil {
					for _, hostname := range hosts {
						fmt.Println(input, hostname)
					}
				}
			}
		})
	}
}

func TestGetNSRecord(t *testing.T) {
	domain := "google.com"

	NSs, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, NS := range NSs {
		IPs, _ := lookHostname(NS.Host)
		fmt.Println(NS.Host, IPs)
	}
}

func TestGetMailServers(t *testing.T) {
	domain := "yandex.ru"

	MXs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, MX := range MXs {
		fmt.Println(MX.Host)
	}
}

// ---------------------------------------------------------

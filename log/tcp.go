package log

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/micro/go-platform/log"
)

var DefaultRAddress net.TCPAddr = net.TCPAddr{
	IP:   net.ParseIP("127.0.0.1"),
	Port: 5000,
	Zone: "",
}

type OutputOptions struct {
	RAddress net.IP
	RPort    int
	LAddress net.IP
}

type OutputOption func(*OutputOptions)

type TcpOutput struct {
	options  OutputOptions
	conn     net.Conn
	raddress *net.TCPAddr
	laddress *net.TCPAddr
	err      error
}

// connect to the sink.  TODO make this use a pool? auto reconnect?
func (to *TcpOutput) connect() error {
	to.conn, to.err = net.DialTCP("tcp", to.laddress, to.raddress)
	fmt.Println("connecting", to.conn, to.err)
	return to.err
}

// Send an event
func (to *TcpOutput) Send(e *log.Event) error {
	if to.conn == nil {
		fmt.Println("TCP output connection is nil")
		if err := to.connect(); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return json.NewEncoder(to.conn).Encode(e)
	// var err error
	// for i := 0; i < 3; i++ {
	// 	go func() {
	// 		multiWriter := io.MultiWriter(to.conn, os.Stderr)
	// 		errCh <- json.NewEncoder(multiWriter).Encode(e)
	// 	}()
	// 	if err = <-errCh; err == nil {
	// 		break
	// 	}
	// }
	// return err
}

// Flush any buffered events
func (to *TcpOutput) Flush() error { return nil }

// Discard the output
func (to *TcpOutput) Close() error {
	if to.conn == nil {
		return to.err
	}
	return to.Close()
}

// Name of output
func (to *TcpOutput) String() string { return "tcp-output" }

// output options

func RemoteAddress(address string) OutputOption {
	return func(o *OutputOptions) {
		if addr, err := net.LookupIP(address); err == nil {
			o.RAddress = addr[0]
			return
		}
		o.RAddress = net.ParseIP(address)
	}
}

func RemotePort(port int) OutputOption {
	return func(o *OutputOptions) {
		o.RPort = port
	}
}

func LocalAddress(address string) OutputOption {
	return func(o *OutputOptions) {
		o.LAddress = net.ParseIP(address)
	}
}

func NewOutput(opts ...OutputOption) log.Output {
	var options OutputOptions
	var laddr *net.TCPAddr = nil
	for _, o := range opts {
		o(&options)
	}

	if len(options.RAddress) == 0 {
		options.RAddress = DefaultRAddress.IP
	}

	if options.RPort == 0 {
		options.RPort = DefaultRAddress.Port
	}
	if len(options.LAddress) != 0 {
		laddr = &net.TCPAddr{options.LAddress, 0, ""}
	}

	tcpOutput := &TcpOutput{
		options:  options,
		raddress: &net.TCPAddr{options.RAddress, options.RPort, ""},
		laddress: laddr,
	}
	tcpOutput.connect()
	return tcpOutput
}

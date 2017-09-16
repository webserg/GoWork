package main
import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"errors"
)
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}
type Arith interface {
	Multiply(args *Args, reply *int) error
	Divide(args *Args, quo *Quotient) error
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func registerArith(server *rpc.Server, arith Arith) {
	// registers Arith interface by name of `Arithmetic`.
	// If you want this name to be same as the type name, you
	// can use server.Register instead.
	server.RegisterName("Arithmetic", arith)
}

func main() {

	//Creating an instance of struct which implement Arith interface
	arith := new(Arith)

	// Register a new rpc server (In most cases, you will use default server only)
	// And register struct we created above by name "Arith"
	// The wrapper method here ensures that only structs which implement Arith interface
	// are allowed to register themselves.
	server := rpc.NewServer()
	registerArith(server, arith)

	// registers an HTTP handler for RPC messages on rpcPath, and a debugging handler on debugPath
	server.HandleHTTP("/", "/debug")

	// Listen for incoming tcp packets on specified port.
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	// This statement starts go's http server on
	// socket specified by l.
	http.Serve(l, nil)
}

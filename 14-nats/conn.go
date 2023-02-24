package _4_nats

import (
	"github.com/nats-io/nats.go"
	"log"
	"strings"
	"time"
)

func NewConn() (*nats.Conn, error) {
	opts := []nats.Option{nats.Name("NATS Sample Subscriber")}
	opts = setupConnOptions(opts)

	services := []string{
		"192.168.124.23:14222",
		"192.168.124.23:24222",
		"192.168.124.23:34222"}
	return nats.Connect(strings.Join(services, ","), opts...)
	//return nats.Connect("192.168.124.38:4222", opts...)
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Printf("Exiting: %v", nc.LastError())
	}))

	//// Use UserCredentials
	//if *userCreds != "" {
	//	opts = append(opts, nats.UserCredentials(*userCreds))
	//}
	//
	//// Use TLS client authentication
	//if *tlsClientCert != "" && *tlsClientKey != "" {
	//	opts = append(opts, nats.ClientCert(*tlsClientCert, *tlsClientKey))
	//}
	//
	//// Use specific CA certificate
	//if *tlsCACert != "" {
	//	opts = append(opts, nats.RootCAs(*tlsCACert))
	//}
	//
	//// Use Nkey authentication.
	//if *nkeyFile != "" {
	//	opt, err := nats.NkeyOptionFromSeed(*nkeyFile)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	opts = append(opts, opt)
	//}

	return opts
}

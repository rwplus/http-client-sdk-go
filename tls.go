package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
)

func GetTLSConfig(opt Options) (*tls.Config, error) {
	if opt.TLSOption == nil {
		return &tls.Config{}, nil
	}

	// clone,avoid mutate it
	tlsCfg := opt.TLSOption

	config := &tls.Config{
		InsecureSkipVerify: tlsCfg.InsecureSkipVerify,
		ServerName:         tlsCfg.ServerName,
	}

	// set up ca certification
	if len(tlsCfg.CAFile) > 0 {
		caFile := x509.NewCertPool()
		if ok := caFile.AppendCertsFromPEM([]byte(tlsCfg.CAFile)); !ok {
			return nil, fmt.Errorf("failed to parse ca pem file")
		}

		config.RootCAs = caFile
	}

	// set up key/cert certification
	if len(tlsCfg.KeyFile) > 0 && len(tlsCfg.CertFile) > 0 {
		cert, err := tls.X509KeyPair([]byte(tlsCfg.CertFile), []byte(tlsCfg.KeyFile))
		if err != nil {
			return nil, err
		}

		config.Certificates = []tls.Certificate{cert}
	}

	return config, nil
}

package es_test

import (
	"fmt"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EsOfficeSuite struct {
	suite.Suite
	client *elasticsearch7.Client
}

func (s *EsOfficeSuite) SetupTest() {
	cfg := elasticsearch7.Config{
		Addresses: []string{
			"http://49.233.61.57:9211",
		},
		Username:                "",
		Password:                "",
		CloudID:                 "",
		APIKey:                  "",
		ServiceToken:            "",
		CertificateFingerprint:  "",
		Header:                  nil,
		CACert:                  nil,
		RetryOnStatus:           nil,
		DisableRetry:            false,
		EnableRetryOnTimeout:    false,
		MaxRetries:              0,
		CompressRequestBody:     false,
		DiscoverNodesOnStart:    false,
		DiscoverNodesInterval:   0,
		EnableMetrics:           false,
		EnableDebugLogger:       false,
		EnableCompatibilityMode: false,
		DisableMetaHeader:       false,
		UseResponseCheckOnly:    false,
		RetryBackoff:            nil,
		Transport:               nil,
		Logger:                  nil,
		Selector:                nil,
		ConnectionPoolFunc:      nil,
	}
	var err error
	s.client, err = elasticsearch7.NewClient(cfg)
	s.NoError(err)
	res, err := s.client.Info()
	s.NoError(err)
	defer func() { _ = res.Body.Close() }()
	fmt.Println(s.client.Transport.(*estransport.Client).URLs())
	fmt.Println(cast.ToString(res.Body))
}

// TestMarshal :
func (s *EsOfficeSuite) TestHelloWorld() {
	fn := s.client.API.Cat.Nodes.WithHuman()
	fn(&esapi.CatNodesRequest{
		Bytes:                   "",
		Format:                  "",
		FullID:                  nil,
		H:                       nil,
		Help:                    nil,
		IncludeUnloadedSegments: nil,
		Local:                   nil,
		MasterTimeout:           0,
		S:                       nil,
		Time:                    "",
		V:                       nil,
		Pretty:                  false,
		Human:                   false,
		ErrorTrace:              false,
		FilterPath:              nil,
		Header:                  nil,
	})
	//s.client.Search.

}

func TestEsOfficeConfiguration(t *testing.T) {
	suite.Run(t, new(EsOfficeSuite))
}

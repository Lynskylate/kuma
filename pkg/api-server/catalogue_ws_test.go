package api_server_test

import (
	"fmt"
	config "github.com/Kong/kuma/pkg/config/api-server"
	"github.com/Kong/kuma/pkg/plugins/resources/memory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
)

var _ = Describe("Catalogue WS", func() {

	It("should return the api catalogue", func() {
		// given
		cfg := config.DefaultApiServerConfig()
		cfg.Catalogue.DataplaneToken.LocalUrl = "http://localhost:1111"
		cfg.Catalogue.DataplaneToken.PublicUrl = "https://kuma.internal:2222"
		cfg.Catalogue.Bootstrap.Url = "http://kuma.internal:3333"

		// setup
		resourceStore := memory.NewStore()
		apiServer := createTestApiServer(resourceStore, cfg)

		stop := make(chan struct{})
		go func() {
			defer GinkgoRecover()
			err := apiServer.Start(stop)
			Expect(err).ToNot(HaveOccurred())
		}()

		// wait for the server
		Eventually(func() error {
			_, err := http.Get(fmt.Sprintf("http://localhost%s/catalogue", apiServer.Address()))
			return err
		}, "3s").ShouldNot(HaveOccurred())

		// when
		resp, err := http.Get(fmt.Sprintf("http://localhost%s/catalogue", apiServer.Address()))
		Expect(err).ToNot(HaveOccurred())

		// then
		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).ToNot(HaveOccurred())

		expected := `
		{
			"apis": {
				"bootstrap": {
					"url": "http://kuma.internal:3333"
				},
				"dataplaneToken": {
					"localUrl": "http://localhost:1111",
					"publicUrl": "https://kuma.internal:2222"
				}
			}
		}
`
		Expect(body).To(MatchJSON(expected))
	})
})
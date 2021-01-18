package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// go test lambda-basic-auth-golang/functions/secured_greeting -v

var _ = Describe("Call greeting", func() {
	var (
		response events.APIGatewayProxyResponse
		request  events.APIGatewayProxyRequest
		err      error
	)

	JustBeforeEach(func() {
		request = events.APIGatewayProxyRequest{
			RequestContext: events.APIGatewayProxyRequestContext{
				Authorizer: map[string]interface{}{
					"username": "Steven",
				},
			},
		}
		response, err = Handler(request)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		request = events.APIGatewayProxyRequest{}
		response = events.APIGatewayProxyResponse{}
	})

	Context("with username Steven", func() {
		It("Fails auth", func() {
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(response.Body).To(Equal("{\"message\":\"Hello from secured lambda function\",\"your_username\":\"Steven\"}"))
		})
	})
})

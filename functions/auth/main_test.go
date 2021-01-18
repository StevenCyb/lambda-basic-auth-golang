package main

import (
	b64 "encoding/base64"
	"os"

	"github.com/aws/aws-lambda-go/events"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// go test lambda-basic-auth-golang/functions/auth -v

var _ = Describe("Authenticate", func() {
	var (
		response events.APIGatewayCustomAuthorizerResponse
		request  events.APIGatewayCustomAuthorizerRequest
		err      error
	)

	JustBeforeEach(func() {
		os.Setenv("USERNAME", "admin")
		os.Setenv("PASSWORD", "gfaho32feoi")

		response, err = Handler(request)
	})

	AfterEach(func() {
		request = events.APIGatewayCustomAuthorizerRequest{}
		response = events.APIGatewayCustomAuthorizerResponse{}
	})

	Context("without auth bearer setted", func() {
		It("Fails auth", func() {
			Expect(err).To(MatchError("Unauthorized"))
			Expect(response).To(Equal(events.APIGatewayCustomAuthorizerResponse{}))
		})
	})

	Context("with invalid auth bearer setted", func() {
		Context("and wrong credentials", func() {
			BeforeEach(func() {
				request = events.APIGatewayCustomAuthorizerRequest{
					AuthorizationToken: "Basic token",
					MethodArn:          "testARN",
				}
			})

			It("Fails auth", func() {
				Expect(err).To(MatchError("Unauthorized"))
				Expect(response).To(Equal(events.APIGatewayCustomAuthorizerResponse{}))
			})
		})

		Context("with wrong credentials", func() {
			BeforeEach(func() {
				request = events.APIGatewayCustomAuthorizerRequest{
					AuthorizationToken: "Basic " + b64.StdEncoding.EncodeToString([]byte("admin:gewgdedwf3eds")),
					MethodArn:          "testARN",
				}
			})

			It("Fails auth", func() {
				Expect(err).To(MatchError("Unauthorized"))
				Expect(response).To(Equal(events.APIGatewayCustomAuthorizerResponse{}))
			})
		})

		Context("with valid credentials", func() {
			BeforeEach(func() {
				request = events.APIGatewayCustomAuthorizerRequest{
					AuthorizationToken: "Basic " + b64.StdEncoding.EncodeToString([]byte(os.Getenv("USERNAME")+":"+os.Getenv("PASSWORD"))),
					MethodArn:          "testARN",
				}
			})

			It("authorizes", func() {
				Expect(err).To(BeNil())
				Expect(response.PolicyDocument.Version).To(Equal("2012-10-17"))
				Expect(response.PolicyDocument.Statement).To(Equal([]events.IAMPolicyStatement{
					{
						Action:   []string{"execute-api:Invoke"},
						Effect:   "Allow",
						Resource: []string{"testARN"},
					},
				}))
			})
		})
	})
})

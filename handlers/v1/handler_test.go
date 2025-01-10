package v1_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/Snehashish1609/validator-api/config"
	"github.com/Snehashish1609/validator-api/models"

	"github.com/Snehashish1609/validator-api/middlewares"

	v1 "github.com/Snehashish1609/validator-api/handlers/v1"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("APIHandler", func() {
	var (
		r           *gin.Engine
		c           *config.Config
		recorder    *httptest.ResponseRecorder
		userHandler *models.UserHandler
		apiHandler  *v1.APIHandler
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		r = gin.Default()
		r.Use(middlewares.LatencyLogger())
		c = config.InitConfig("test", "test")

		userHandler = models.NewUserHandler()
		apiHandler = v1.NewAPIHandler(c, userHandler)

		// routes
		r.POST("/validate-user", apiHandler.ValidateUser)

		recorder = httptest.NewRecorder()
	})

	Describe("Validate User Payload", func() {
		Context("can validate when the payload is valid", func() {
			It("should return 200 OK success", func() {
				payload := `{"name":"foo bar","pan":"ABCDE1234F","mobile":"1234567890","email":"foobar@example.com"}`
				req, _ := http.NewRequest("POST", "/validate-user", bytes.NewBufferString(payload))
				r.ServeHTTP(recorder, req)

				Expect(recorder.Code).To(Equal(http.StatusOK))
				Expect(recorder.Body.String()).To(ContainSubstring("payload is valid"))
			})
		})

		Context("throw error when the PAN format is invalid", func() {
			It("should return 400 error", func() {
				payload := `{"name":"foo bar","pan":"Something123","mobile":"1234567890","email":"foobar@example.com"}`
				req, _ := http.NewRequest("POST", "/validate-user", bytes.NewBufferString(payload))
				r.ServeHTTP(recorder, req)
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
				Expect(recorder.Body.String()).To(ContainSubstring("pan"))
			})
		})

		Context("throw error when the mobile number format is invalid", func() {
			It("should return 400 error", func() {
				payload := `{"name":"foo bar","pan":"ABCDE1234F","mobile":"12345","email":"foobar@example.com"}`
				req, _ := http.NewRequest("POST", "/validate-user", bytes.NewBufferString(payload))
				r.ServeHTTP(recorder, req)

				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
				Expect(recorder.Body.String()).To(ContainSubstring("mobile"))
			})
		})

		Context("throw error when the payload is empty", func() {
			It("should return 400 error", func() {
				req, _ := http.NewRequest("POST", "/validate-user", bytes.NewBufferString("{}"))
				r.ServeHTTP(recorder, req)
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
				Expect(recorder.Body.String()).To(ContainSubstring("required"))
			})
		})

		Context("can validate when the json is malformed", func() {
			It("should return 400 error", func() {
				payload := `{"name":"foo bar", "pan":"ABCDE1234F"`
				req, _ := http.NewRequest("POST", "/validate-user", bytes.NewBufferString(payload))

				r.ServeHTTP(recorder, req)
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

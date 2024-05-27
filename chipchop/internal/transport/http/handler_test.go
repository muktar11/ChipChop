package http_test

import (
    "testing"
    "github.com/gin-gonic/gin"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "net/http"
    "net/http/httptest"
)

func TestAPI(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "API Suite")
}

var _ = Describe("API", func() {
    var router *gin.Engine

    BeforeEach(func() {
        // Setup router and service mocks here
    })

    Describe("GET /estimations", func() {
        Context("when there are no estimations", func() {
            It("should return an empty list", func() {
                w := httptest.NewRecorder()
                req, _ := http.NewRequest("GET", "/estimations", nil)
                router.ServeHTTP(w, req)
                Expect(w.Code).To(Equal(http.StatusOK))
                Expect(w.Body.String()).To(MatchJSON(`[]`))
            })
        })
    })
})

package handler_test

import (
	"context"
	"errors"
	"github.com/cobbinma/example-go-api/cmd/api/handler"
	"github.com/cobbinma/example-go-api/pkg/models"
	mockModels "github.com/cobbinma/example-go-api/pkg/models/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("GetPet", func() {
	var (
		ctx        context.Context
		ctrl       *gomock.Controller
		repository *mockModels.MockRepository
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		repository = mockModels.NewMockRepository(ctrl)
	})

	Context("with valid pet id", func() {
		Context("get pet does not error", func() {
			pet := models.NewPet(1, "tim", "")
			expected := `{"id":1,"name":"tim"}
`
			BeforeEach(func() {
				repository.EXPECT().GetPet(ctx, 1).Return(pet, error(nil))
			})
			It("should return a 200 status", func() {
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/pet/:id")
				c.SetParamNames("id")
				c.SetParamValues("1")

				h := handler.NewHandler(repository)

				err := h.GetPet(c)
				Expect(rec.Code).To(Equal(http.StatusOK))
				Expect(err).To(BeNil())
			})
			It("should return expected pet", func() {
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/pet/:id")
				c.SetParamNames("id")
				c.SetParamValues("1")

				h := handler.NewHandler(repository)

				err := h.GetPet(c)
				Expect(rec.Body.String()).To(Equal(expected))
				Expect(err).To(BeNil())
			})
		})
		Context("when get pet returns error", func() {
			BeforeEach(func() {
				repository.EXPECT().GetPet(ctx, 1).Return(nil, models.NewPetError(errors.New("test error"), "test error", 0))
			})
			It("should return a 500 internal server error status", func() {
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/pet/:id")
				c.SetParamNames("id")
				c.SetParamValues("1")

				h := handler.NewHandler(repository)

				err := h.GetPet(c)
				Expect(rec.Code).To(Equal(http.StatusInternalServerError))
				Expect(err).To(BeNil())
			})
		})
	})
})

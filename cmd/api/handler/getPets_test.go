package handler_test

import (
	"context"
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

var _ = Describe("CreatePet", func() {
	var (
		ctx        context.Context
		ctrl       *gomock.Controller
		repository *mockModels.MockRepository
		e          *echo.Echo
	)

	BeforeEach(func() {
		ctx = context.Background()
		e = echo.New()
		ctrl = gomock.NewController(GinkgoT())
		repository = mockModels.NewMockRepository(ctrl)
	})

	Context("get pets does not error", func() {
		pets := []*models.Pet{models.NewPet(1, "tim", "")}
		expected := `[{"id":1,"name":"tim"}]
`
		BeforeEach(func() {
			repository.EXPECT().GetPets(ctx, 100, 0).Return(pets, error(nil))
		})
		It("should return a 200 status", func() {
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			h := handler.NewHandler(repository)

			err := h.GetPets(c)
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(err).To(BeNil())
		})
		It("should return expected pets", func() {
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			h := handler.NewHandler(repository)

			err := h.GetPets(c)
			Expect(rec.Body.String()).To(Equal(expected))
			Expect(err).To(BeNil())
		})
	})
})

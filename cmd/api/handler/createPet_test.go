package handler_test

import (
	"context"
	"github.com/cobbinma/example-go-api/cmd/api/handler"
	"github.com/cobbinma/example-go-api/models"
	mockModels "github.com/cobbinma/example-go-api/models/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
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

	Context("with valid input", func() {
		Context("without tag", func() {
			req := `{
	"id": 1,
	"name": "tim"
}`
			pet := models.NewPet(1, "tim", "")
			BeforeEach(func() {
				repository.EXPECT().CreatePet(gomock.Eq(ctx), gomock.Eq(pet)).Return(error(nil))
			})
			It("should return a status no content to client", func() {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(req))
				rec := httptest.NewRecorder()

				c := e.NewContext(req, rec)
				h := handler.NewHandler(repository)

				err := h.CreatePet(c)
				Expect(rec.Code).To(Equal(http.StatusCreated))
				Expect(err).To(BeNil())
			})
		})
		Context("with tag", func() {
			req := `{
	"id": 1,
	"name": "tim",
	"tag": "dog"
}`
			pet := models.NewPet(1, "tim", "dog")
			BeforeEach(func() {
				repository.EXPECT().CreatePet(gomock.Eq(ctx), gomock.Eq(pet)).Return(error(nil))
			})
			It("should return a status no content to client", func() {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(req))
				rec := httptest.NewRecorder()

				c := e.NewContext(req, rec)
				h := handler.NewHandler(repository)

				err := h.CreatePet(c)
				Expect(rec.Code).To(Equal(http.StatusCreated))
				Expect(err).To(BeNil())
			})
		})
	})
	Context("with invalid input", func() {
		Context("of missing id", func() {
			req := `{
	"name": "tim"
}`
			It("should return a status no content to client", func() {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(req))
				rec := httptest.NewRecorder()

				c := e.NewContext(req, rec)
				h := handler.NewHandler(repository)

				err := h.CreatePet(c)
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
				Expect(err).To(BeNil())
			})
		})
		Context("of missing name", func() {
			req := `{
	"id": 1
}`
			It("should return a status no content to client", func() {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(req))
				rec := httptest.NewRecorder()

				c := e.NewContext(req, rec)
				h := handler.NewHandler(repository)

				err := h.CreatePet(c)
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
				Expect(err).To(BeNil())
			})
		})
		Context("of empty request", func() {
			It("should return a status no content to client", func() {
				req := httptest.NewRequest(http.MethodPost, "/", nil)
				rec := httptest.NewRecorder()

				c := e.NewContext(req, rec)
				h := handler.NewHandler(repository)

				err := h.CreatePet(c)
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
				Expect(err).To(BeNil())
			})
		})
	})
})

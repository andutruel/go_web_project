package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/andutruel/go_transacciones/internal/transactions"
	"github.com/andutruel/go_transacciones/pkg/web"
	"github.com/gin-gonic/gin"
)

type requestTransaccion struct {
	CodigoTransaccion string  `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_transaccion"`
}

type TransaccionHandler struct {
	service transactions.Service
}

// ListTransactions godoc
// @Summary Lists transactions
// @Tags Transactions
// @Description get transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions [get]
func (c *TransaccionHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		t, err := c.service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}
		ctx.JSON(http.StatusOK, t)
	}
}

// StoreTransactions godoc
// @Summary Store transactions
// @Tags Transactions
// @Description store transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions [post]
func (c *TransaccionHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		var req requestTransaccion

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		t, err := c.service.Store(req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaTransaccion)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		ctx.JSON(http.StatusOK, t)
	}
}

// UpdateTransactions godoc
// @Summary Update transactions
// @Tags Transactions
// @Description update transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions/:id [post]
func (c *TransaccionHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid ID"))
			return
		}

		var req requestTransaccion

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		//validaciones para campos requeridos
		if req.CodigoTransaccion == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El código de transacción no puede estar vacío."))
			return
		}

		t, err := c.service.Update(id, req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaTransaccion)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, t, ""))
	}
}

func NewTransaccion(t transactions.Service) *TransaccionHandler {
	return &TransaccionHandler{
		service: t,
	}
}

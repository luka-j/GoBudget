package Routes

import (
	"GoBudget/Database/Entities"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type CreateEnvelopeRequest struct {
	Name     string          `json:"name" binding:"required"`
	Interval int16           `json:"type" binding:"required"`
	Budget   decimal.Decimal `json:"budget" binding:"required"`
}

type EnvelopeResponse struct {
	Id               string          `json:"id"`
	Name             string          `json:"name"`
	EnvelopeInterval int16           `json:"interval"`
	BudgetedAmount   decimal.Decimal `json:"budget"`
	CurrentAmount    decimal.Decimal `json:"currentAmount"`
}

func createEnvelope(c *gin.Context) {
	var request CreateEnvelopeRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}
	tokenClaims := jwt.ExtractClaims(c)
	envelope, err := Entities.CreateEnvelope(request.Name, request.Interval, request.Budget, tokenClaims["sub"].(string))
	if err != nil {
		panic(err)
	}
	c.JSON(200, buildEnvelopeResponse(envelope))
}

func getEnvelopes(c *gin.Context) {
	userId := jwt.ExtractClaims(c)["sub"].(string)
	envelopes, err := Entities.GetEnvelopesByUser(userId)
	if err != nil {
		panic(err)
	}
	envelopesResponse := make([]EnvelopeResponse, len(envelopes))
	for i, envelope := range envelopes {
		envelopesResponse[i] = buildEnvelopeResponse(envelope)
	}
	c.JSON(200, envelopesResponse)
}

func patchEnvelope(c *gin.Context) {
	var request PatchRequest
	id := c.Param("id")
	if err := c.BindJSON(&request); err != nil {
		return
	}
	userId := jwt.ExtractClaims(c)["sub"].(string)
	envelope, err := Entities.EditEnvelope(id, userId, request.Field, request.Value)
	if err == Entities.ERROR_NO_SUCH_FIELD {
		c.JSON(400, ErrorResponse{Code: "NO_SUCH_FIELD", Message: "No field " + request.Field + " present for envelopes."})
		return
	}
	if err != nil {
		panic(err)
	}
	c.JSON(200, buildEnvelopeResponse(envelope))
}

func deleteEnvelope(c *gin.Context) {
	id := c.Param("id")
	userId := jwt.ExtractClaims(c)["sub"].(string)
	if err := Entities.DeleteEnvelope(id, userId); err != nil {
		panic(err)
	}
	c.Status(200)
}

func buildEnvelopeResponse(envelope Entities.Envelope) EnvelopeResponse {
	return EnvelopeResponse{
		envelope.ID.String(),
		envelope.Name,
		envelope.EnvelopeInterval,
		envelope.BudgetedAmount,
		envelope.CurrentAmount,
	}
}

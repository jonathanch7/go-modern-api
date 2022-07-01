package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseBuilder struct {
	ctx    *gin.Context
	status int
	entity interface{}
	err    error
}

// Create the builder of the response
func NewResponseBuilder(ctx *gin.Context, entity interface{}) *ResponseBuilder {
	return &ResponseBuilder{
		ctx:    ctx,
		entity: entity,
	}
}

// Status code of response
func (b *ResponseBuilder) Status(status int) *ResponseBuilder {
	b.status = status
	return b
}

// Add error to response
func (b *ResponseBuilder) Error(err error) *ResponseBuilder {
	b.err = err
	return b
}

// Constructs the response for the entity and submits the response
// When the status code is not specified by default it will be 200,
// except when the interface is nil, there it will be by default 204.
func (b *ResponseBuilder) Build() {

	if b.err != nil {
		if b.status == 0 {
			b.status = http.StatusBadRequest
		}
		b.ctx.JSON(b.status, gin.H{"message": b.err.Error()})
		return
	}

	if b.entity == nil {
		if b.status == 0 {
			b.status = http.StatusNoContent
		}
		b.ctx.Status(b.status)
		return
	}
	if b.status == 0 {
		b.status = http.StatusOK
	}

	b.ctx.JSON(b.status, b.entity)

	return
}

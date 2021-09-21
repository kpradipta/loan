package loan

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lending/internal/core/domain"
	"github.com/lending/internal/core/ports"
	"net/http"
	"strings"
)

type HTTPHandler struct {
	loanServices ports.LoanService
}

func NewHTTPHandler(loanServices ports.LoanService) *HTTPHandler {
	return &HTTPHandler{
		loanServices: loanServices,
	}
}

func (hdl *HTTPHandler) Create(c *gin.Context) {
	res := map[string]string{}
	var loan domain.Loan
	if err := c.BindJSON(&loan); err != nil {
		fmt.Printf( "error parsing request: %s", err.Error())
		res["message"]= "error while marshalling response"
		res["code"]= "300"

		c.AbortWithStatusJSON(http.StatusBadRequest, res )
		return
	}

	id, err := hdl.loanServices.Create(loan)

	if err != nil {
		fmt.Println(err.Error())
		splitMessage := strings.Split(err.Error(), ":")
		res["message"]= splitMessage[1]
		res["code"]= splitMessage[0]
		c.AbortWithStatusJSON(http.StatusInternalServerError, res )
	}
	res["message"]= "Success"
	res["code"]= "0"
	res["loan_id"]= id
	c.JSON(http.StatusCreated, res)
}

func (hdl *HTTPHandler) Read(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id : ",id)
	loan, err := hdl.loanServices.Read(id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, loan)
}

func (hdl *HTTPHandler) Update(c *gin.Context) {
	res := make(map[string]interface{})
	var loan domain.Loan
	if err := c.BindJSON(&loan); err != nil {
		fmt.Printf( "error parsing request: %s", err.Error())
		res["message"]= "error while parsing request"
		res["code"]= "300"

		c.AbortWithStatusJSON(http.StatusBadRequest, res )
		return
	}

	loan, err := hdl.loanServices.Update(loan)

	if err != nil {
		fmt.Println(err.Error())
		splitMessage := strings.Split(err.Error(), ":")
		res["message"]= splitMessage[1]
		res["code"]= splitMessage[0]
		c.AbortWithStatusJSON(http.StatusInternalServerError, res )
		return
	}

	res["message"]= "Success"
	res["code"]= "0"
	res["data"]= loan
	c.JSON(http.StatusCreated, res)
}

func (hdl *HTTPHandler) Approve(c *gin.Context) {

}
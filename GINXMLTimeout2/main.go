package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	XMLName        xml.Name       `xml:"Body"`
	ConfirmRequest ConfirmRequest `xml:"ConfirmRequest"`
}

type ConfirmRequest struct {
	XMLName    xml.Name `xml:"ConfirmRequest"`
	TxType     string   `xml:"txType"`
	ReasonCode string   `xml:"reasonCode"`
	ReasonDesc string   `xml:"reasonDesc"`
	SourceTxNo string   `xml:"sourceTxNo"`
	Timestamp  string   `xml:"timestamp"`
	Ref_no1    string   `xml:"ref_no1"`
	Ref_no2    string   `xml:"ref_no2"`
}

func main() {
	router := gin.Default()
	count := 0
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Envelope
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if count < 5 {
			time.Sleep(20 * time.Second)
		} else {
			time.Sleep(5 * time.Second)
		}
		fmt.Println(count)
		count++
		//c.JSON(http.StatusOK, gin.H{"status": "you are logged in", "Type": xml.Body.ConfirmRequest.TxType})
		xml.Body.ConfirmRequest.ReasonCode = "0001"
		xml.Body.ConfirmRequest.ReasonDesc = "Alreadypaid"
		//c.JSON(http.StatusOK, xml)

		var RetXML EnvelopeS
		RetXML.Body.ConfirmRequest = xml.Body.ConfirmRequest
		c.XML(http.StatusOK, RetXML)
	})
	// Listen and serve on 0.0.0.0:8080
	router.Run(":8090")
}

type EnvelopeS struct {
	XMLName xml.Name `xml:"S:Envelope"`
	Body    BodyS    `xml:"S:Body"`
}

type BodyS struct {
	XMLName        xml.Name       `xml:"S:Body"`
	ConfirmRequest ConfirmRequest `xml:"ConfirmRequest"`
}

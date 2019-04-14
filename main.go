package main

import (
	"log"
	"net/http"
	"os"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

var pairDescription = map[string]string{
	"btc_mxn": "bitcoin en pesos mexicanos",
	"eth_mxn": "ethereum en pesos mexicanos",
	"xrp_mxn": "ripple en pesos mexicanos",
}

var applications = map[string]interface{}{
	"/health": alexa.StdApplication{
		Methods: "GET",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Ok!"))
		},
	},
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	appid := os.Getenv("APPID")
	if appid == "" {
		log.Fatal("$APPID must be set")
	}

	applications["/echo/helloworld"] = alexa.EchoApplication{ // Route
		AppID:    appid, // Echo App ID from Amazon Dashboard
		OnIntent: echoIntentHandler,
		OnLaunch: echoIntentHandler,
	}

	alexa.Run(applications, port)
}

func echoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {

	var response *alexa.EchoResponse

	switch echoReq.GetIntentName() {
	case "getbtcmxn":
		bitsoResponse := getBitsoResponse("btc_mxn")
		response = echoResp.OutputSpeech(bitsoResponse).Card("BITSO", bitsoResponse)
	case "getethmxn":
		bitsoResponse := getBitsoResponse("eth_mxn")
		response = echoResp.OutputSpeech(bitsoResponse).Card("BITSO", bitsoResponse)
	case "getxrpmxn":
		bitsoResponse := getBitsoResponse("xrp_mxn")
		response = echoResp.OutputSpeech(bitsoResponse).Card("BITSO", bitsoResponse)
	case "getbtcusdt":
		binanceResponse := MarketPrices()
		response = echoResp.OutputSpeech(binanceResponse).Card("BINANCE", binanceResponse)
	case "AMAZON.HelpIntent":
		response = handleHelpIntent()
	default:
		response = handleAboutIntent(echoReq)
	}

	if response == nil {
		response = alexa.NewEchoResponse()
		response.OutputSpeech("Disculpa, algo no salio bien.")
	}

	*echoResp = *response
}

func handleAboutIntent(echoReq *alexa.EchoRequest) *alexa.EchoResponse {

	response := alexa.NewEchoResponse()
	response.OutputSpeech("Te puedo decir los precios de bitso, solo dime alexa, abre bitso, y dame etereum")
	response.SimpleCard("About", "Aplicación no oficial")
	MarketPrices()
	return response
}

func handleHelpIntent() *alexa.EchoResponse {

	response := alexa.NewEchoResponse()
	builder := alexa.NewSSMLTextBuilder()

	builder.AppendSentence("Aqui hay algunas cosas que puedes preguntar: ")
	builder.AppendSentence("Dame valor bitcoin.")
	builder.AppendSentence("Dame valor ethereum.")
	builder.AppendSentence("Dame valor ripple.")

	return response.OutputSpeechSSML(builder.Build())
}

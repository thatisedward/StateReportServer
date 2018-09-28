package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type PostDuridRsp struct {
	Result PostDuridValue `json:"result"`
}

type PostDuridValue struct {
	Received int `json:"received"`
	Sent     int `json:"sent"`
}

func handlingReport(packetChan *chan string) {

	for {
		parcelCapacity := Configurations.Control.parcelCapacity

		parcel := bytes.NewBufferString("")

		for i := 0; i < parcelCapacity-1; i++ {
			item := <-*packetChan

			if len(item) > 0 {
				parcel.WriteString(<-*packetChan)
				parcel.WriteString("\n")
			}
		}
		parcel.WriteString(<-*packetChan)

		dayLog.Debug("Parcel is full, ready for posting!")

		if parcel.Len() == 0 {
			report.errors(POST_PARCEL_EMPTY)
			continue
		}

		doneCheck := false

		for i := 0; i < Configurations.Control.postRetry && !doneCheck; i++ {

			doneCheck = posting(parcel)
			dayLog.Debug("The", i, " th try", "| postFlag is ", doneCheck)

			if i > 0 {
				report.errors(POST_RETRY)
			}
		}

		if !doneCheck {
			report.errors(POST_FAILED)
			dayLog.Debug("Reposting all failed.| postFlag still remains: ", doneCheck)
		}
	}
}

func posting(parcel *bytes.Buffer) bool {
	client := &http.Client{}
	resp, err := client.Post(Configurations.Durid.url, "text/plain", parcel)

	if err != nil {
		dayLog.Error("Failures! Client.Post failed ->", err)
		report.errors(POST_HTTP_ERROR)
		return false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		dayLog.Error("Failures! Ioutil.ReadAll failed ->", err)
		return false
	}

	postResp := &PostDuridRsp{}
	parsing_error := json.Unmarshal(body, postResp)

	if parsing_error != nil {
		dayLog.Error("Failure! Parsing postResp failed ->", parsing_error, "parcel is: ", parcel)
		report.errors(DURID_PARSE_ERROR)
		return false
	}

	dayLog.Debug("Success! HTTP Received: ", string(body))

	if postResp.Result.Received > 0 && postResp.Result.Sent > 0 {
		//*doneCheck = true
		dayLog.Info("Successfully posted and parsed by Durid! "+
			"| postResp.Result.Received: ", postResp.Result.Received,
			"| postResp.Result.Sent: ", postResp.Result.Sent)

		report.packet(postResp.Result.Sent)
	} else {
		dayLog.Error("WRONG! failures when Durid was receiving!",
			"| 1th postResp.Result.Received: ", postResp.Result.Received,
			"| 2th postResp.Result.Sent: ", postResp.Result.Sent,
			"| parcel is:", parcel.String())
		return false
	}
	return true
}

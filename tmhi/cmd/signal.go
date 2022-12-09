package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var signalCmd = &cobra.Command{
	Use:   "signal",
	Short: "View the gateway's current signal metrics",
	RunE: func(cmd *cobra.Command, args []string) error {

		req, err := http.NewRequest("GET", "http://192.168.12.1/TMI/v1/gateway?get=signal", nil)
		if err != nil {
			return err
		}

		req.Header.Add("Authorization", "Bearer "+viper.GetString("password"))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var respJSON signalResp

		json.Unmarshal(body, &respJSON)

		displaySignalData(respJSON.Signal.Signal4G, "4G")
		fmt.Println() // print an empty line
		displaySignalData(respJSON.Signal.Signal5G, "5G")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(signalCmd)
}

func displaySignalData(data signalMetrics, gen string) {

	////
	// Format data
	////

	// coverage bars
	barsInt := int(math.Round(data.Bars))
	bars := strings.Repeat("▮", barsInt)
	bars = bars + strings.Repeat("▯", 5-barsInt)

	// rsrq - reference signal received quality
	rsrq := data.RSRQ
	var rsrqGrade string
	switch {
	case rsrq >= -8:
		rsrqGrade = "great"
	case rsrq >= -15:
		rsrqGrade = "good"
	case rsrq >= -21:
		rsrqGrade = "poor"
	default:
		rsrqGrade = "useless"
	}

	// sinr - sigal to interference and noise ratio
	sinr := data.SINR
	var sinrGrade string
	switch {
	case sinr >= 20:
		sinrGrade = "great"
	case sinr >= 13:
		sinrGrade = "good"
	case sinr >= 0:
		sinrGrade = "poor"
	default:
		sinrGrade = "useless"
	}

	// eNB ID - tower ID
	var towerID string
	if data.ENBID != 0 {
		towerID = strconv.Itoa(data.ENBID)
	} else {
		towerID = "n/a"
	}

	// present
	fmt.Printf("%s Signal:\n", gen)
	fmt.Printf("Bars: %s (%d/5)\n", bars, barsInt)
	fmt.Printf("RSRQ: %s (%d)\n", rsrqGrade, rsrq)
	fmt.Printf("SINR: %s (%d)\n", sinrGrade, sinr)
	fmt.Printf("Band: %s\n", data.Bands[0])
	fmt.Printf("Tower ID: %s\n", towerID)
}

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Display information about local T-Mobile gateway",
	RunE: func(cmd *cobra.Command, args []string) error {

		// all is used because we need time in addition to device
		req, err := http.NewRequest("GET", "http://192.168.12.1/TMI/v1/gateway?get=all", nil)
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

		var respJSON gatewayResp

		json.Unmarshal(body, &respJSON)

		// prepare
		dur, err := time.ParseDuration(strconv.Itoa(respJSON.Time.Uptime) + "s")
		if err != nil {
			return err
		}

		fmt.Println("Gateway data:")
		fmt.Printf("Name: %s\n", respJSON.Device.FriendlyName)
		fmt.Printf("Device: %s %s\n", respJSON.Device.Manufacturer, respJSON.Device.Model)
		fmt.Printf("Serial No.: %s\n", respJSON.Device.Serial)
		fmt.Printf("Uptime: %s\n", dur)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(gatewayCmd)
}

// Copyright (C) 2021 github.com/V4NSH4J
//
// This source code has been released under the GNU Affero General Public
// License v3.0. A copy of this license is available at
// https://www.gnu.org/licenses/agpl-3.0.en.html

package directmessage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/V4NSH4J/discord-mass-dm-GO/utilities"
)


// "Opens" the channel with a discord account and outputs the Channel ID or the Channel Snowflake
func OpenChannel(authorization string, recepientUID string, cookie string, fingerprint string) string {
	url := "https://discord.com/api/v9/users/@me/channels"

	json_data := []byte("{\"recipients\":[\"" + recepientUID + "\"]}")
	//Cookie := utilities.GetCookie()
	//if Cookie.Dcfduid == "" && Cookie.Sdcfduid == "" {
	//	fmt.Println("ERR: Empty cookie")
	//	return ""
	//}

	// Cookies := "__dcfduid=" + Cookie.Dcfduid + "; " + "__sdcfduid=" + Cookie.Sdcfduid + "; " + " locale=us" + "; __cfruid=d2f75b0a2c63c38e6b3ab5226909e5184b1acb3e-1634536904"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println("Error while making request")
		return ""
	}
	req.Header.Set("authorization", authorization)
	req.Header.Set("Cookie", cookie)
	req.Header.Set("x-fingerprint", fingerprint)
	httpClient := &http.Client{}
	resp, err := httpClient.Do(utilities.CommonHeaders(req))

	if err != nil {
		fmt.Printf("Error while sending Open channel request  %v", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("[SNOWFLAKE] %v", string(body))
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("[%v]Invalid Status Code while sending request %v \n",time.Now().Format("15:05:04"), resp.StatusCode)
		return ""
	}
	type responseBody struct {
		ID string `json:"id,omitempty"`
	}

	var channelSnowflake responseBody
	json.Unmarshal(body, &channelSnowflake)

	return channelSnowflake.ID
}

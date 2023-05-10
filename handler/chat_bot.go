package handler

import (
	"chatBot/until/zaplog"
	"fmt"
	"github.com/chatgp/chatgpt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func GetChatReplay(c *gin.Context) {
	jData := map[string]interface{}{
		"ret":   0,
		"value": c.Query("aa"),
	}
	// new chatgpt client
	token := `eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIn0..OjwGx52_vBVySNQ9.4mWWGJyCsgnbC3mpvfdw48hpkZDNcFgJmrlX5MdqFxR0WCZezx7Ip8evEvtj03X87pwqbO2Eh0nVy1x362qKuf46Gf5aJFDjgBMA9sHXiJun9UP-C01lUKd7AmIyfhtm8xVIPcYW1-WqX5m_Vze8vzm8FjBKB-lNeT6QfVW3uhZnPG1jR9PWDEtLVSqmyZT7mpD7_8_U4YXSd0rbyb-TyFy0yL1yooYPJlpiJmB3dGR_PnJvienH-91JRVIlsEc9UoUjUN4Eh3chnsAU8dRJu_ddRSG8Qzo-2dwVP7jDFZOPDiIKqx_1XsybgKaimsv-rJ-0G7Se57d-YYyFLQd0rGmJW_A_S0W1v-rY0GF9nhn5-7CsuAaWKfAl7KLsYB0KBNDaHF1FRLGHBPZVnvUSF3Al70xcax3BxviladMGCiQ-ec0v0YN60cgbNzfkLjbN-jZ5NIABIUOerxwcpgQyrND1Enz8FRu86zDR1fRLKKhj89iVlSWwKP-X72p3Ui9RRmfF_wEOFLlQfmrR8Z6-nj53aqpYfEIbHUEw_haMQ5yWz2bxdzSlNP6FXWzxtXx-MMRYUFlGe69ycjW_-uAzIZEHELhxU369AA4FiHLR_Zycb91LFsNBJS7LQf7x9s7Iu0SrWvJ2nu6Vi0GbBEWNvdkF4zAeysE90vd3-evPoW9l4pEdhfiWeUnzyZco74ntUAOiw0jIEwLqLUtbzOKuDovjZtX_t_qrNkhbanJ6cGJnqyERcTLSua4al1hJVCy7A-uYIITsVNtDNiBlvcLT4p80GM2RIpOpiOKEF2EWYHbLklBfyZC-DdBWzuU-L6JU3rORJbAv5szdh-sucxqQSe-XT2lfbLXxUJaStc4XcgXYOeqyrZMDNL3YVDQoaeeupumme6mS3EVve8A2m5jHDFotg3_7DPo9rYPuE_XaujA-UNx-bOEi9oqicSGkuA5BsIlXDBIxlwgpuxA8VNVLBzvYokK8rnDGTvexhMVfiZEcF1KxfjZfan4mPyRyUYuLhb9sIlU9EP9gl5A1kHuO67bSsxdq_c_8xVG9cMNjFsX-xwAdt83YlgDl0ZpTii9Dwu4g98YknbMt1NoO7LpBRmd93xwxYKapNysJPG9pW29vJoWYbDKMW_ODVudTkVu0B3ZnT_qP3aW2-Cy2bam7R4uGjbFgblQ1aBvvE1xy4FpnwhhP8MiH3_YT0T_6mf70J5uTw4ilGQ-sJxaXMQxqr_FNCD_4DKOUUuUSn8k9mBOpYd9A8Pz1k4NWWluhFXV33YPcnLOk46qslftjOLQ3ryLsJe9bDzUmULC1A4KfnajHk8ILODsRKnDneJV-ACFyWBVR5DOBf-lmL5frMSbXST0BqQsxR3fWjQCKca7aC5uizzKJtYq1toqJ7YQ-7jAaRpfY9oYuIbVLapENxE3RLdmzCHrl5vlSDE9ideWPgRigDH-KX5wi-QZC9pDLViTtys4JlXqh3DENKeAtB3xXI_RSfb-D6pq3RbHj5VPN7kEO2CnIr7LobAHwwkVjQKPp-BDc-LF4j0KcJSM7wmfMcp_czZLN7Kfw_t3GJS-_iznYV7MMbvF0K1or7g0L2LppHK0vA4M9DZNkI_ZDSDMCfJQUlg1iPFQdiwW7ri7-d2uhP8g3rrEhhWPDCDYmWmN1ZRGmOsdrz8VmA7nUt7-bjv74Vs5_x5ZfneJWnbuObdyTxgqbDC_tjbJ8kj2BFm5pLpST1GehyFIugk8qeSV1b4H93c6qL3h_pQUssEEqAT_KkDwHmxBLo0cHfc9PVnyhrTjePz1IRdNPBxdIO4uNz9xojWlaRwipp_q2oi1nLan9m--r8Y9m8R6M_bZL5in3TMVglGGuLuLmbvkTHy2IvFdaQsaZzH4lZBpp7rJyWMtQ-Fn0UWrZRM3CITZswza6GAbUZa0mmVlKkugmK5lRW3Sfz9I6fTJ6XLIgXgGwdBvmsOgrEBGGzGgsRMhPg0zBUIqsfEOjl_UBBz8-nXb_LdCf1CCbBOizM7uxna06rLcTssoMY6Zl8DxlwSOfguIGYp3BVbqqOsfvcLwA7SUjqcm2oK8G7t-5CR11SxRWba9s7CM8NIwgVGBeLlznVosJBl2xJDqQHDmRutK-z0jjb_tPZDK-7b2kXnpyYnv9vfAxc8ut-0XeBZ_Zf3MYfou4wiYKzi9k4dYwo3fbcoB4ghjyRzMXEhH2hWNGwpuGq5CUluNEx9GoFb8jNyN0bA5LnqqwMCFMx2t4FwJjdrM986pj-jQjePwxcKPQoWaHmaIJsEGYQVqn_aHUjXQlCGoZVRVdgb41C5g47QObetME39ImXFqe9iZciXgQhvpqB8aBGr6fwRmJwOScJGTOzpE3qoFQ9ZJ7o5flennJzonkBDzrC09RB8rVKJsJYIixXbKWxMc-zWnDw5Nxe6oss_Pr6YKxdfNOndlG-qHNXjUpUzrINUjvR_8nM5iduE4dcCsV851B2w3lloI_tnZXBXkKnw_5Q46Q-DDCfr7w22CI3xKWnQEFmfCOE7VIFpG_nVMw1vgi71qCFJ4_CNx551gzuQG8YlaZpUZEgJk6LB5STpe9JhIUKVAqwa2a92NctzXqukq8b6nM2NAq0fd9BvTIuikrTo5TGH59cw4UxVap-C5Vn5ppylP57JejcTCZuCgMgYtNknXByxTTR3ONMCV9_wbamWQ9-RZU801ffBK68Vjgy86_3QQKOr2XW15AIhrTztBqsAB3cY5g1sSHurbhMcK2at5wcB5KGGOpqLs2GvkMnxD9PSLp0-Seb2g7LBZzqWXoGw.RGBZoS32vC8OVkL1Mn4nyw`
	cfValue := "oUmd8bAzkC.dBgBWwld1g.Gh22P4bibJJ0l0M9hoekg-1683699191-0-1-bb5f92c7.7e47f054.ec5ecd27-160"

	cookies := []*http.Cookie{
		{
			Name:  "__Secure-next-auth.session-token",
			Value: token,
		},
		{
			Name:  "cf_clearance",
			Value: cfValue,
		},
	}

	cli := chatgpt.NewClient(
		chatgpt.WithDebug(true),
		chatgpt.WithTimeout(60*time.Second),
		chatgpt.WithCookies(cookies),
	)

	// first message
	message := c.Query("question")
	text, err := cli.GetChatText(message)
	if err != nil {
		jData["ret"] = 10050
		jData["msg"] = fmt.Sprintf("get chat text failed: %v", err)
		return
	}

	// continue conversation with new message
	conversationID := text.ConversationID
	parentMessage := text.MessageID
	newMessage := "again"

	newText, err := cli.GetChatText(newMessage, conversationID, parentMessage)
	if err != nil {
		jData["ret"] = 10051
		jData["msg"] = fmt.Sprintf("get chat text failed: %v", err)
		return
	}

	zaplog.Trace("GetChatReplayRequest").Info("GetGpt", zap.Any("sReq", jData), zap.Any("err", err), zap.Error(err))
	if err != nil {
		jData["ret"] = 10052
		jData["errMsg"] = err.Error()
		c.JSONP(200, jData)
		return
	}
	jData["value"] = newText
	c.JSONP(200, jData)
	return
}

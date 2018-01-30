package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"fmt"
	"math/rand"
	"time"
)

func get_rand_ip() string {
	rand.Seed(time.Now().UnixNano())
	ip_t := [...]string{"218", "218", "66", "66", "218", "218", "60", "60", "202", "204", "66", "66", "66", "59", "61", "60", "222", "221", "66", "59", "60", "60", "66", "218", "218", "62", "63", "64", "66", "66", "122", "211"}
	a := ip_t[rand.Intn(len(ip_t)-1)]
	b := rand.Intn(254)
	c := rand.Intn(254)
	d := rand.Intn(254)
	return fmt.Sprintf("%s.%d.%d.%d", a, b, c, d)
}

func post() {
	var r http.Request
	r.ParseForm()
	r.Form.Add("id", "MTkz")
	//r.Form.Add("keyCode", "6a5694b8028ae5271bd58acab03f6aabe6814762")
	//r.Form.Add("keyCode", "414553ed80a79ab61bdc3fdbed86457afea043a0")
	bodystr := strings.TrimSpace(r.Form.Encode())

	request, err := http.NewRequest("POST", "http://www.test.com/vote.php", strings.NewReader(bodystr))
	//request.Header.Set("Cookie", "ctt_vote_cookie=213dc9e224e2368cbc44421a14c052bb1e18ab6c; PHPSESSID=u8o8r9lhtupio88in0aepv9o56; CTRACKID=u8o8r9lhtupio88in0aepv9o56-1516696541000; UM_distinctid=16122289c00907-0db752d8bbc13a-32657403-240000-16122289c012c6; Hm_lvt_addae925595b77324e7a3c1e940f3ad0=1516696544; PHPSESSID=u8o8r9lhtupio88in0aepv9o56; CNZZDATA1256335252=999783365-1516693464-%7C1516844673; share_userid=580131; share_token=7dd3df020399f5919f00078c4e8fbebb; share_username=%E4%B8%80%E8%B4%AB; share_nicename=%E4%B8%80%E8%B4%AB; display_name=%E4%B8%80%E8%B4%AB; share_avatar=%2Fuploads%2Favatar%2F000%2F58%2F01%2Fsmall_31.jpg; Hm_lpvt_addae925595b77324e7a3c1e940f3ad0=1516849730")
	if err != nil {
		log.Println(err)
	}
	request.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("X-Forwarded-For", get_rand_ip())
	request.Header.Set("Origin", "Keep-Alive")
	request.Header.Set("Connection", "http://www.test.com")
	request.Header.Set("Referer", "http://www.test.com/vote2.php")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Mobile Safari/537.36")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	log.Println(string(b))

}
func main() {
	forever := make(chan bool)
	for i := 0; i < 200; i++ {
		go post()
	}
	<-forever
}
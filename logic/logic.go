package logic

import (
	"net/url"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"url_shortner/configure"
)

type URL struct{
	org_url string;
	tiny_url string;
}

func getrequest(tiny string)string{
	str1 ,err := http.Get(tiny)
	if err!=nil{
		log.Fatal(err)
	}
	defer str1.Body.Close()
	content, err:= ioutil.ReadAll(str1.Body)

	if err!=nil{
		log.Fatal(err)
	}
	return string(content)
}




func usingTinyURL(urlorg string) (string,string){
	temp:= url.QueryEscape(urlorg)
	tinyurl:= fmt.Sprintf("http://tinyurl.com/api-create.php?url=%s", temp)
	str1:= getrequest(tinyurl)
	return str1 ,urlorg

}

func (u *URL) urlshortner(urlorg string) *URL{

		s,l:=usingTinyURL(urlorg)
		u.org_url=l
		u.tiny_url=s
		return u
}



func API(org string) string{
	u:=URL{}
	u.urlshortner(org)
	temp:=u.tiny_url
	z:=strings.Split(temp,"/")
	fmt.Printf("Original URL : %s\nTiny URL : %s\n",u.org_url,z[3])
	configure.Insert(u.org_url,z[3])

	return u.tiny_url
}





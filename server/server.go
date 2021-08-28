package server

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"url_shortner/logic"
	"strings"
	"url_shortner/configure"
)

func Run(){


	router:=mux.NewRouter()



	router.HandleFunc("/convert/{email_id}", func(w http.ResponseWriter, r *http.Request) {

		x:=mux.Vars(r)
		temp:=string(logic.API(x["email_id"]))
		z:=strings.Split(temp,"/")

		fmt.Fprintf(w,"Original URL : %s\nTiny URL : %s\n",x["email_id"],z[3])
	})

	router.HandleFunc("/reverse/{tiny_id}", func(w http.ResponseWriter, r *http.Request) {

		x:=mux.Vars(r)
		temp:=string(configure.Find(x["tiny_id"]))

		fmt.Fprintf(w,"Original URL : %s\nTiny URL : %s\n",temp,x["tiny_id"])
	})


	router.HandleFunc("/showall", func(w http.ResponseWriter, r *http.Request) {

			temp:= configure.Show()

			fmt.Fprintf(w,"List:\n")

			for i,j:= range temp{
				fmt.Fprintf(w,"%s : %s\n",i,j)

			}

	})

	http.ListenAndServe(":8080",router)

}


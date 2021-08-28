package configure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)


var (
	 org string
	 tiny string
)



func Show()map[string]string{
	db, err := sql.Open("mysql", "root:@/dummy")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	//var org, tiny string
	var tinyurl = map[string]string{}



	stmt,err:=db.Query("select org,tiny from tinyurl")

	defer stmt.Close()

	for stmt.Next(){
		err:= stmt.Scan(&org, &tiny)
		if err != nil {
			log.Fatal(err)
		}

		tinyurl[org]=tiny

	}
	return tinyurl
}

func Insert(orgurl ,tinyurl string){

	db, err := sql.Open("mysql", "root:@/dummy")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()


	fmt.Println(orgurl,tinyurl)

	stmt,err:=db.Prepare("insert into tinyurl(org,tiny) VALUES(?,?)")
	if err != nil {

		log.Fatal(err)
	}
	res, err :=stmt.Exec(orgurl, tinyurl)
	fmt.Println(res.RowsAffected())
	if err != nil {
		log.Fatalln("duplicate key")
	}
	defer stmt.Close()

}


func Find(t string)string{
	db, err := sql.Open("mysql", "root:@/dummy")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	var org string
	stmt,err:= db.Query("select org from tinyurl where tiny=?",t)

	defer stmt.Close()
	if err!=nil{
		log.Fatal(err)
	}
	for stmt.Next(){
		err := stmt.Scan(&org)
		if err != nil {
			log.Fatal(err)
		}
		return org
	}


	return "error not found"
}


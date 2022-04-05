package main

import (
	"database/sql"
	"net/http"
	"log"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func main()  {
	http.HandleFunc("/", Inicio)
	log.Println("Server Running...")
	http.ListenAndServe(":9000", nil)
}


func conexionDB()(conexion *sql.DB){
	Driver:="mysql"
	Usuario:="root"
	Nombre:="sopes1P2"
	Contra:=""
	conexion, err:= sql.Open(Driver, Usuario+":"+Contra+"@tcp(34.136.54.2:4000)/" + Nombre)
	if err != nil{
		panic(err.Error())
	}
	return conexion;
}


func Inicio(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola SIUUUU")
	conexionEstablecida := conexionDB()
	nombre := "Juan solares"
	edad := 23
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO Test(nombre, edad) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insertarRegistros.Exec(nombre, edad)
}
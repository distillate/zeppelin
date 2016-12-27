package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Mux() http.Handler {
	router := httprouter.New()

	// Pack operations
	router.GET("/pack/:id", GetPack)
	router.POST("/pack", CreatePack)
	router.PUT("/pack/:id", UpdatePack)
	router.DELETE("/pack/:id", DeletePack)
	router.POST("/pack/:id/launch", LaunchPack)

	// Launch operations
	router.GET("/launch", GetLaunches)
	router.GET("/launch/:id", GetLaunch)
	router.DELETE("/launch/:id", RevokeLaunch)

	// Suite operations
	router.DELETE("/suite/:id", RevokeSuite)
	router.GET("/suite/:id/logs", GetSuiteLogs)

	// Report operations
	router.POST("/launch/:id/report", RequestReport)

	return router
}

func CreatePack(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notImplemented(w)
}

func UpdatePack(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//packId := ps.ByName("id")
	notImplemented(w)
}

func GetPack(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//packId := ps.ByName("id")
	notImplemented(w)
}

func DeletePack(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//packId := ps.ByName("id")
	notImplemented(w)
}

func LaunchPack(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//packId := ps.ByName("id")
	notImplemented(w)
}

func GetLaunches(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	notImplemented(w)
}

func GetLaunch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//launchId := ps.ByName("id")
	notImplemented(w)
}

func RevokeLaunch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//launchId := ps.ByName("id")
	notImplemented(w)
}

func RevokeSuite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//suiteId := ps.ByName("id")
	notImplemented(w)
}

func GetSuiteLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//suiteId := ps.ByName("id")
	notImplemented(w)
}

func RequestReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//launchId := ps.ByName("id")
	notImplemented(w)
}

func notImplemented(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("This method is not implemented..."))
}

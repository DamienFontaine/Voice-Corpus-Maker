// Copyright (c) - Damien Fontaine <damien.fontaine@lineolia.net>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>

package main

import (
	"net/http"

	"github.com/DamienFontaine/lunarc/datasource/mongo"
	"github.com/DamienFontaine/lunarc/web"
	"github.com/gorilla/mux"
	"gitlab.lineolia.net/meda/voice-corpus-maker/record"
)

//Route to Server
func Route(s *web.Server, m mongo.Mongo) {
	r := mux.NewRouter()

	//Record
	recordService := record.NewService(mongo.Service{Mongo: m})
	recordController := record.NewController(&recordService)
	prh := http.HandlerFunc(recordController.Post)
	r.HandleFunc("/record", prh).Methods("POST")
	erh := http.HandlerFunc(recordController.Export)
	r.HandleFunc("/record/export", erh).Methods("GET")
	urh := http.HandlerFunc(recordController.Update)
	r.HandleFunc("/record/{id}", urh).Methods("POST")
	crh := http.HandlerFunc(recordController.Count)
	r.HandleFunc("/record/count", crh).Methods("GET")
	gpprh := http.HandlerFunc(recordController.GetPerPage)
	r.HandleFunc("/record/page/{page}/{limit}", gpprh).Methods("GET")
	grh := http.HandlerFunc(recordController.Get)
	r.HandleFunc("/record", grh).Methods("GET")
	uprh := http.HandlerFunc(recordController.Upload)
	r.HandleFunc("/upload/{id}", uprh).Methods("GET")
	drh := http.HandlerFunc(recordController.Delete)
	r.HandleFunc("/record/{id}", drh).Methods("DELETE")

	// Angular 7
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	r.PathPrefix("/").Handler(web.SingleFile("public/index.html"))

	mux := s.Handler.(*web.LoggingServeMux)
	mux.Handle("/", r)
}

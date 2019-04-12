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
	"flag"
	"log"

	"github.com/DamienFontaine/lunarc/datasource/mongo"
	"github.com/DamienFontaine/lunarc/web"
)

var env string

func init() {
	flag.StringVar(&env, "env", "development", "Set development, test, staging or production environment")
}

func main() {

	flag.Parse()

	s, err := web.NewServer("config.yml", env)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	m, err := mongo.NewMongo("config.yml", env)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	Route(s, *m)
	go s.Start()

	select {
	case <-s.Done:
		log.Println("Server shutdown")
		return
	case <-s.Error:
		log.Println("Error: server terminate")
		return
	}
}

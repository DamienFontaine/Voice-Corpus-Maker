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

package record

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Controller listen /record
type Controller struct {
	manager Manager
}

//NewController constructs new Controller
func NewController(m Manager) *Controller {
	c := Controller{manager: m}
	return &c
}

//Post creates a new record.
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer file.Close()

	text := r.FormValue("text")
	set := r.FormValue("set")

	record := Record{Metadata: Metadata{Text: text, Set: set}}
	record, err = c.manager.Add(record, file)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	json, _ := json.Marshal(record)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// Count records in DataSource
func (c *Controller) Count(w http.ResponseWriter, r *http.Request) {
	count, err := c.manager.Count()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	w.Write([]byte(fmt.Sprintf("%v", count)))
}

//Get returns all records.
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	records, err := c.manager.FindAll()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	json, err := json.Marshal(records)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

//GetPerPage list records per page.
func (c *Controller) GetPerPage(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.ParseInt(mux.Vars(r)["page"], 10, 32)
	limit, err := strconv.ParseInt(mux.Vars(r)["limit"], 10, 32)
	records, err := c.manager.FindPerPage(int32(page), int32(limit))
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	json, err := json.Marshal(records)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// Update a record
func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	set := r.FormValue("set")

	record, err := c.manager.Update(id, set)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	json, err := json.Marshal(record)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

//Upload a record on client
func (c *Controller) Upload(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	reader, err := c.manager.Upload(id)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	io.Copy(w, reader)
}

//Delete a record in DataSource
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := c.manager.Delete(id)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
}

// Export all records
func (c *Controller) Export(w http.ResponseWriter, r *http.Request) {
	f, err := c.manager.Export()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	io.Copy(w, f)

	w.Header().Set("Content-Description", "File Transfer")
	w.Header().Set("Content-Type", "application/octet-stream")
}

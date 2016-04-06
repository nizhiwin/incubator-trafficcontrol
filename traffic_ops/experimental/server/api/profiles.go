// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Profiles struct {
	Name        string        `db:"name" json:"name"`
	Description string        `db:"description" json:"description"`
	CreatedAt   time.Time     `db:"created_at" json:"createdAt"`
	Links       ProfilesLinks `json:"_links" db:-`
}

type ProfilesLinks struct {
	Self string `db:"self" json:"_self"`
}

type ProfilesLink struct {
	ID  string `db:"profile" json:"name"`
	Ref string `db:"profiles_name_ref" json:"_ref"`
}

// @Title getProfilesById
// @Description retrieves the profiles information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Profiles
// @Resource /api/2.0
// @Router /api/2.0/profiles/{id} [get]
func getProfilesById(name string, db *sqlx.DB) (interface{}, error) {
	ret := []Profiles{}
	arg := Profiles{}
	arg.Name = name
	queryStr := "select *, concat('" + API_PATH + "profiles/', name) as self"
	queryStr += " from profiles WHERE name=:name"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getProfiless
// @Description retrieves the profiles
// @Accept  application/json
// @Success 200 {array}    Profiles
// @Resource /api/2.0
// @Router /api/2.0/profiles [get]
func getProfiless(db *sqlx.DB) (interface{}, error) {
	ret := []Profiles{}
	queryStr := "select *, concat('" + API_PATH + "profiles/', name) as self"
	queryStr += " from profiles"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postProfiles
// @Description enter a new profiles
// @Accept  application/json
// @Param                 Body body     Profiles   true "Profiles object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/profiles [post]
func postProfiles(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Profiles
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "INSERT INTO profiles("
	sqlString += "name"
	sqlString += ",description"
	sqlString += ",created_at"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:description"
	sqlString += ",:created_at"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putProfiles
// @Description modify an existing profilesentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Profiles   true "Profiles object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/profiles/{id}  [put]
func putProfiles(name string, payload []byte, db *sqlx.DB) (interface{}, error) {
	var arg Profiles
	err := json.Unmarshal(payload, &arg)
	arg.Name = name
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "UPDATE profiles SET "
	sqlString += "name = :name"
	sqlString += ",description = :description"
	sqlString += ",created_at = :created_at"
	sqlString += " WHERE name=:name"
	result, err := db.NamedExec(sqlString, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delProfilesById
// @Description deletes profiles information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Profiles
// @Resource /api/2.0
// @Router /api/2.0/profiles/{id} [delete]
func delProfiles(name string, db *sqlx.DB) (interface{}, error) {
	arg := Profiles{}
	arg.Name = name
	result, err := db.NamedExec("DELETE FROM profiles WHERE name=:name", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
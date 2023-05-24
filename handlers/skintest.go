package handlers

import (
	"encoding/json"
	"github.com/araquach/apiClientForms/helpers"
	"github.com/araquach/apiClientForms/models"
	db "github.com/araquach/dbService"
	"github.com/gorilla/mux"
	"github.com/signintech/pdft"
	gopdf "github.com/signintech/pdft/minigopdf"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Error error

func ApiSkinTestCreate(w http.ResponseWriter, r *http.Request) {
	var error helpers.Error

	decoder := json.NewDecoder(r.Body)

	var data models.SkinTest
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	db.DB.Create(&data)
	if err != nil {
		error.Message = "Server error."
		helpers.RespondWithError(w, http.StatusInternalServerError, error)
		return
	}

	w.WriteHeader(http.StatusOK)

	now := time.Now()
	d := now.Format("2006-01-02 15:04:05")

	w.Header().Set("Content-Type", "application/json")
	helpers.ResponseJSON(w, data)

	createSkinTestPDF(data)
	helpers.SendSkinTestEmail(data)
	helpers.SaveToS3("output/skinTest/skintest.pdf", data.LastName+" "+data.FirstName+" "+d+".pdf")

	return
}

func ApiGetTestedClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	s, _ := strconv.Atoi(vars["salon"])
	fn := vars["first_name"]
	ln := vars["last_name"]

	var skinTests []models.ApiSkinTest

	if fn == "0" && ln == "0" {
		db.DB.Where("salon", s).Model(&models.SkinTest{}).Limit(12).Order("id desc").Find(&skinTests)
	} else {
		db.DB.Where("first_name ILIKE ? AND last_name ILIKE ?", fn+"%", ln+"%").Model(&models.SkinTest{}).Find(&skinTests)
	}

	json, err := json.Marshal(skinTests)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func ApiSkinTestDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	param := vars["id"]

	var skintest models.ApiSkinTest
	db.DB.Where("id", param).Model(&models.SkinTest{}).Find(&skintest)

	json, err := json.Marshal(skintest)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func createSkinTestPDF(r models.SkinTest) {
	var err error
	var pt pdft.PDFt
	var l string
	var a string

	now := time.Now()
	date := now.Format("02/01/2006")

	switch r.Salon {
	case 1:
		l = "forms/logos/logo_jakata.png"
		a = "forms/address/address_jakata.png"
	case 2:
		l = "forms/logos/logo_pk.png"
		a = "forms/address/address_pk.png"
	case 3:
		l = "forms/logos/logo_base.png"
		a = "forms/address/address_base.png"
	}

	logo := pngToBase64(l)
	logo = strings.Split(logo, ",")[1]

	address := pngToBase64(a)
	address = strings.Split(address, ",")[1]

	client := r.FirstName + " " + r.LastName
	signature := strings.Split(r.Signature, ",")[1]

	if r.Type == "renewal" {
		err = pt.Open("forms/skinTest/skin_test_renewal.pdf")
		if err != nil {
			panic("Couldn't open pdf.")
		}
	}
	if r.Type == "new" {
		err = pt.Open("forms/skinTest/skin_test_new.pdf")
		if err != nil {
			panic("Couldn't open pdf.")
		}
	}
	if r.Type == "minor" {
		err = pt.Open("forms/skinTest/skin_test_minor.pdf")
		if err != nil {
			panic("Couldn't open pdf.")
		}
	}

	err = pt.AddFont("helvetica", "fonts/Helvetica.ttf")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = pt.AddFont("helvetica-bold", "fonts/Helvetica-Bold.ttf")
	if err != nil {
		log.Fatal(err)
		return
	}

	//insert the bits to pdf

	// Logo
	err = pt.InsertImgBase64(logo, 1, 30, 30, 321.26, 94.49)
	if err != nil {
		panic(err)
	}

	// Address
	err = pt.InsertImgBase64(address, 1, 440, 40, 120, 71.25)
	if err != nil {
		panic(err)
	}

	err = pt.SetFont("helvetica-bold", "", 13)
	if err != nil {
		panic(err)
	}

	// Client Name:
	err = pt.Insert("Client Name: "+client, 1, 35, 160, 100, 100, gopdf.Left|gopdf.Top)
	if err != nil {
		panic(err)
	}

	err = pt.SetFont("helvetica", "", 13)
	if err != nil {
		panic(err)
	}

	// Date:
	err = pt.Insert(date, 3, 160, 125, 100, 100, gopdf.Left|gopdf.Top)
	if err != nil {
		panic(err)
	}

	// First Name:
	err = pt.Insert(r.FirstName, 3, 160, 153, 100, 100, gopdf.Left|gopdf.Top)
	if err != nil {
		panic(err)
	}

	// Last Name:
	err = pt.Insert(r.LastName, 3, 160, 181, 100, 100, gopdf.Left|gopdf.Top)
	if err != nil {
		panic(err)
	}

	// Signature:
	err = pt.Insert(r.Email, 3, 160, 208, 260, 100, gopdf.Left|gopdf.Top)
	if err != nil {
		panic(err)
	}

	err = pt.InsertImgBase64(signature, 3, 20, 400, 175, 75)

	err = pt.SetFont("helvetica", "", 13)
	if err != nil {
		panic(err)
	}

	// Client Name:
	err = pt.Insert(client, 3, 35, 490, 100, 100, gopdf.Left|gopdf.Top)
	if err != nil {
		panic(err)
	}

	//save within the apps output folder
	err = pt.Save("output/skinTest/skintest.pdf")
	if err != nil {
		log.Fatal("Couldn't save output pdf")
	}
}

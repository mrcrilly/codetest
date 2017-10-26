// file mux_ex.go
package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/leekchan/accounting"
	"strconv"
)

// Simple global variable to act as a runtime cache
// I would ideally place the data in a database, but that
// assumes you want to manipulate it
var jsonCache = &PoorMansCache{}

// newTemplate() is a helper function for attaching
// additional functions to the templating engine without
// having to repeat the process for each handler.
func newTemplate() *template.Template {

	// This probably is ideal each and every time.
	// Precompiled templates would be better
	t := template.New("")

	t.Funcs(template.FuncMap{
		// Having nicer dates makes a big difference even in this trivial
		// example
		"prettyDate": func(date string) string {
			trimmedDate := strings.Split(date, "+")[0]
			t, err := time.Parse("2006-01-02T15:04:05", trimmedDate)
			if err != nil {
				panic(err)
				return date
			}

			return t.Format("02 Jan 2006")
		},

		// Pretty printing the currency makes a difference too
		"prettyCurrency": func(amount string) string {
			asFloat, err := strconv.ParseFloat(amount, 32)
			if err != nil {
				panic(err)
				return amount
			}

			// Might be a bit heavy bringing in a whole lib to do this
			// Another option would be to copy/paste the code from the lib
			// that does just enough for our needs in a helper function
			ac := accounting.Accounting{Symbol: "$", Precision: 2}
			return ac.FormatMoney(asFloat)
		},
	})

	return t
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	t := newTemplate()

	// Grabbing the templates from disk each time is not efficient
	// but it does enable fast development during the MVP stage.
	t, err := t.ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		// It would be better to use error pages here, of course
		w.Write([]byte("Unable to render template."))
		return
	}

	t.ExecuteTemplate(w, "base", jsonCache)
}

func handlerLotteryView(w http.ResponseWriter, r *http.Request) {
	t := newTemplate()

	// Grabbing the templates from disk each time is not efficient
	// but it does enable fast development during the MVP stage.
	t, err := t.ParseFiles("templates/base.html", "templates/view_lottery.html")
	if err != nil {
		w.Write([]byte("Unable to render template."))
		return
	}

	vars := mux.Vars(r)
	if key, OK := vars["key"]; OK {
		data := jsonCache.getLotteryKey(key)
		if data != nil {
			t.ExecuteTemplate(w, "base", jsonCache.getLotteryKey(key))
			return
		}

		// It would be better to use error pages here, of course
		w.Write([]byte("Unable to find that key in the JSON cache"))
	} else {
		// It would be better to use error pages here, of course
		w.Write([]byte("Failed to find key in parameters"))
	}
}

func handlerRaffleView(w http.ResponseWriter, r *http.Request) {
	t := newTemplate()

	// Grabbing the templates from disk each time is not efficient
	// but it does enable fast development during the MVP stage.
	t, err := t.ParseFiles("templates/base.html", "templates/view_raffle.html")
	if err != nil {
		w.Write([]byte("Unable to render template."))
		return
	}

	vars := mux.Vars(r)
	if key, OK := vars["key"]; OK {
		data := jsonCache.getRaffleKey(key)
		if data != nil {
			t.ExecuteTemplate(w, "base", jsonCache.getRaffleKey(key))
			return
		}

		// It would be better to use error pages here, of course
		w.Write([]byte("Unable to find that key in the JSON cache"))
	} else {
		// It would be better to use error pages here, of course
		w.Write([]byte("Failed to find key in parameters"))
	}
}

func main() {
	fillCache()

	r := mux.NewRouter()
	r.HandleFunc("/", handlerIndex)
	r.HandleFunc("/lottery/{key}", handlerLotteryView)
	r.HandleFunc("/raffle/{key}", handlerRaffleView)
	log.Fatal(http.ListenAndServe(":8080", r))
}

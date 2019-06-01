package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"helloworld/travian/engine"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jetbasrawi/go.cqrs"
)

var (
	readModel  travian.ReadModelFacade
	dispatcher ycq.Dispatcher
	repo       travian.VillageRepository

	villageView = parseTemplate("dorf1.html")
)


func init() {
	// CQRS Infrastructure configuration

	// Configure the read model
	// Create a readModel instance
	readModel = travian.NewMap(200, 10);

	// we have several projection that we need to init

	// todo init all the projection that we require (this is extendable)
	eventBus := ycq.NewInternalEventBus()

	// add a projection for the Resources per village.
	resources := travian.NewResourceProjection()
	eventBus.AddHandler(resources,
		&travian.VillageEstablished{},
		&travian.FieldUpgraded{},
	)

	// Here we use an in memory event repository.
	repo = travian.NewInMemoryRepo(eventBus)

	// create command handlers!
	villageCommandHandler := travian.NewVillageCommandHandlers(repo)

	// Create a dispatcher
	dispatcher = ycq.NewInMemoryDispatcher()
	// Register the inventory command handlers instance as a command handler
	// for the events specified.
	err := dispatcher.RegisterHandler(villageCommandHandler,
		&travian.EstablishVillage{},
		&travian.UpgradeField{},
	)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	setupHandlers()
	log.Fatal(http.ListenAndServe(":8088", nil))
}

var epoch = time.Unix(0, 0).Format(time.RFC1123)

var noCacheHeaders = map[string]string{
	"Expires":         epoch,
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

var etagHeaders = []string{
	"ETag",
	"If-Modified-Since",
	"If-Match",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

func NoCache(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Delete any ETag headers that may have been set
		for _, v := range etagHeaders {
			if r.Header.Get(v) != "" {
				r.Header.Del(v)
			}
		}
		// Set our NoCache headers
		for k, v := range noCacheHeaders {
			w.Header().Set(k, v)
		}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func setupHandlers() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", NoCache(http.StripPrefix("/assets/", fs)))

	r.Methods("GET").Path("/village").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//params := mux.Vars(r)
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				log.Fatal(err)
			}

			id := r.Form.Get("ID")
			//id := ycq.NewUUID()
			em := ycq.NewCommandMessage(id, &travian.EstablishVillage{
				X:     0,
				Y:     0,
				Owner: r.Form.Get("name"),
			})

			err = dispatcher.Dispatch(em)
			if err != nil {
				log.Println(err)
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		d := struct {
			Title string
		}{"Village 1"}

		if err := villageView.Execute(w, r, d); err != nil {
			log.Println(err)
		}
	})

	r.Methods("GET").Path("/village/{id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		fmt.Print(params)
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				log.Fatal(err)
			}

			id := r.Form.Get("ID")
			//id := ycq.NewUUID()
			em := ycq.NewCommandMessage(id, &travian.EstablishVillage{
				X:     0,
				Y:     0,
				Owner: r.Form.Get("name"),
			})

			err = dispatcher.Dispatch(em)
			if err != nil {
				log.Println(err)
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
//!= nil {
//			log.Fatal(err)
//		}
		//if err != nil {
		//	log.Fatal(err)
		//}
	})

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))

	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//
	//	Items := readModel.GetVillages();
	//	err := t.ExecuteTemplate(w, "index", Items)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//})
	//
	//mux.HandleFunc("/add",
	//
	//mux.HandleFunc("/village/{id}/build/{index}", func(w http.ResponseWriter, r *http.Request) {
	//
	//	if r.Method == http.MethodPost {
	//		err := r.ParseForm()
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		id := ycq.NewUUID()
	//		em := ycq.NewCommandMessage(id, &travian.UpgradeField{
	//			ID :   "123",
	//			Index: 1,
	//		})
	//
	//		err = dispatcher.Dispatch(em)
	//		if err != nil {
	//			log.Println(err)
	//		}
	//
	//		http.Redirect(w, r, "/", http.StatusSeeOther)
	//	}
	//
	//	err := t.ExecuteTemplate(w, "build", nil)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//})
	//
	//mux.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
	//	staticFile := r.URL.Path[len("/assets/"):]
	//	if len(staticFile) != 0 {
	//		f, err := http.Dir("assets/").Open(staticFile)
	//		if err == nil {
	//			content := io.ReadSeeker(f)
	//			http.ServeContent(w, r, staticFile, time.Now(), content)
	//			return
	//		}
	//	}
	//	http.NotFound(w, r)
	//})
}

package main

import (
	"encoding/json"
	travian "helloworld/travian/engine"
	"math"
	"strconv"

	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/jetbasrawi/go.cqrs"
)

var (
	world       travian.ReadModelFacade
	dispatcher  ycq.Dispatcher
	repo        travian.VillageRepository
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	villageView = parseTemplate("dorf1.html")
	karteView   = parseTemplate("karte.html")
)

func init() {
	rand.Seed(time.Now().UnixNano())

	// CQRS Infrastructure configuration
	//speed := flag.Int("speed", 1, "set the server speed")
	//size := flag.Int("size", 1, "set the server speed")

	// Configure the read model
	// Create a world instance
	world = travian.NewMap(200, 256)

	fmt.Print(world.Coordinate(78, -74).Id())
	//fmt.Print(world.)
	dispatcher = ycq.NewInMemoryDispatcher()
	// we have several projection that we need to init

	// todo init all the projection that we require (this is extendable)
	eventBus := ycq.NewInternalEventBus()
	// SAGAS
	//buildingQueue := queues.NewBuildingQueue(dispatcher)
	//eventBus.AddHandler(buildingQueue,
	//	&travian.VillageEstablished{},
	//	&travian.EnqueuedBuilding{},
	//	&travian.CompletedBuilding{},
	//	&travian.AbortedBuilding{},
	//)

	// PROJECTIONS
	// add a projection for the Resources per village.
	resources := travian.NewResourceProjection()
	eventBus.AddHandler(resources,
		&travian.VillageEstablished{},
		&travian.FieldUpgraded{},
	)

	// Here we use an in memory event repository.
	repo = travian.NewInMemoryRepo(eventBus)

	//COMMAND HANDLERS

	// create command handlers!
	villageCommandHandler := travian.NewVillageCommandHandlers(repo)

	// Register the inventory command handlers instance as a command handler
	// for the events specified.
	err := dispatcher.RegisterHandler(villageCommandHandler,
		&travian.EstablishVillage{},
		&travian.UpgradeField{},
	)
	if err != nil {
		log.Fatal(err)
	}

	//id := ycq.NewUUID()
	em := ycq.NewCommandMessage("1", &travian.EstablishVillage{
		X:     0,
		Y:     0,
		Owner: "admin",
	})

	err = dispatcher.Dispatch(em)

}

func main() {
	setupHandlers()
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func setupHandlers() {
	r := mux.NewRouter()

	//assets with ETAG headers
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", Etag(http.StripPrefix("/assets/", fs)))

	r.Methods("GET").Path("/ajax.php").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()
		x, _ := strconv.Atoi(v.Get("x"))
		xx, _ := strconv.Atoi(v.Get("xx"))
		y, _ := strconv.Atoi(v.Get("y"))

		size := int(math.Abs(float64(x)-float64(xx))) + 1
		centerCoordinate := world.Coordinate(x+(size/2), y+(size/2))
		data := world.Karte(centerCoordinate.Id(), size)
		z, _ := json.Marshal(data)
		w.Write(z)
	})

	r.Methods("GET").Path("/karte.php").Queries("d", "{\\d+}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//params := mux.Vars(r)
		//params := mux.Vars(r)

		center, _ := strconv.Atoi(r.FormValue("d"))

		tiles := world.Karte(center, 7)
		//fmt.Print(tiles);
		d := struct {
			Tiles      [][]*travian.Tile
			Coordinate travian.Coordinate
		}{tiles, world.CoordinateForId(center)}

		if err := karteView.Execute(w, r, d); err != nil {
			log.Println(err)
		}
	})

	r.Methods("GET").Path("/village").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//params := mux.Vars(r)

		id := ycq.NewUUID()
		em := ycq.NewCommandMessage(id, &travian.EstablishVillage{
			X:     0,
			Y:     0,
			Owner: RandStringRunes(20),
		})

		if err := dispatcher.Dispatch(em); err != nil {
			log.Println(err)
		}

		//http.Redirect(w, r, "/village", http.StatusSeeOther)

		village := world.GetVillage("1")

		fmt.Println(village)
		//d := struct {
		//	Title string
		//}{"Village 1"}

		if err := villageView.Execute(w, r, village); err != nil {
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

}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

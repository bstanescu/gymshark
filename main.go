package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var packs []int
var target int

func closestMinPacks(packs []int, target int) map[int]int {
	// sort packs descending, first one will be the largest one
	sort.Sort(sort.Reverse(sort.IntSlice(packs)))
	upperTarget := (target/packs[0] + 1) * packs[0]

	// prepare the dynamic programming matrix
	dp := make([]int, upperTarget)
	packUsed := make([]int, upperTarget)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= upperTarget-1; i++ {
		for _, pack := range packs {
			if i < pack {
				continue
			}
			subRes := dp[i-pack]
			if subRes != math.MaxInt32 && subRes+1 < dp[i] {
				dp[i] = subRes + 1
				packUsed[i] = pack
			}
		}
	}

	// find the closest greater target
	closestGreaterTarget := target
	for i := target; i < len(dp); i++ {
		if dp[i] < dp[closestGreaterTarget] {
			closestGreaterTarget = i
		}
	}

	// Reconstruct the selected packs for the closest greater target.
	selectedPacks := make(map[int]int, len(packs))
	current := closestGreaterTarget
	for current > 0 {
		selectedPacks[packUsed[current]]++
		current -= packUsed[current]
	}

	return selectedPacks
}

func postRoute(w http.ResponseWriter, r *http.Request) {
	if packsInputStr := r.FormValue("packs"); packsInputStr != "" {
		tmp := strings.Split(packsInputStr, ",")
		packsInput := make([]int, 0, len(tmp))
		for _, raw := range tmp {
			v, err := strconv.Atoi(strings.TrimSpace(raw))
			if err != nil {
				log.Print(err)
				continue
			}
			packsInput = append(packsInput, v)
		}
		if len(packsInput) > 0 {
			packs = packsInput
		}
	}

	if targetStr := r.FormValue("target"); targetStr != "" {
		target, _ = strconv.Atoi(strings.TrimSpace(targetStr))
	}

	type Data struct {
		Target   int
		Packs    string
		Solution string
	}

	p := make([]string, 0)
	for pack, count := range closestMinPacks(packs, target) {
		p = append(p, fmt.Sprintf("[%d x %d]", pack, count))
	}

	tmp, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	tmp.Execute(w, Data{
		Target:   target,
		Packs:    strings.Trim(strings.Join(strings.Fields(fmt.Sprint(packs)), ", "), "[]"),
		Solution: strings.Join(p, ", "),
	})
}

func main() {
	// set some defaults
	packs = []int{2000, 500, 250, 1000}
	target = 63

	router := mux.NewRouter()
	router.HandleFunc("/", postRoute).Methods(http.MethodPost, http.MethodGet)
	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", router)
}

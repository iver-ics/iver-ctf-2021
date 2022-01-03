package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"

	"git.mills.io/prologic/go-gopher"
	"github.com/google/uuid"
)

type handlers struct {
	firstID    uuid.UUID
	firstIDStr string
	flagIDStr  string
	flagStr    string
	probs      map[uuid.UUID]mathProbLink
	cfg        Config
}

func (h *handlers) handleRFC1436(w gopher.ResponseWriter, r *gopher.Request) {
	w.Write(assetsRFC1436)
}

func (h *handlers) handleMath(w gopher.ResponseWriter, r *gopher.Request) {
	var id uuid.UUID
	if strings.HasSuffix(r.Selector, h.firstIDStr) {
		w.WriteInfo("Welcome to the math challenge(s)!")
		w.WriteInfo("To proceed, you'll have to answer the question correctly.")
		w.WriteInfo("")
		w.WriteInfo("Let's begin!")
		id = h.firstID
	} else {
		idStr := strings.TrimPrefix(r.Selector, "/math/")
		var err error
		id, err = uuid.Parse(idStr)
		if err != nil {
			w.WriteError(fmt.Sprintf("Failed to parse math problem UUID: %q", idStr))
			w.WriteError("Error: " + err.Error())
			w.WriteInfo("")
			writeDirectoryLink(w, h.cfg, "/", "Back to start page.")
			return
		}
	}
	prob, ok := h.probs[id]
	if !ok {
		w.WriteError("Wrong answer!")
		w.WriteInfo("")
		writeDirectoryLink(w, h.cfg, "/", "Back to start page.")
		return
	}
	if id != h.firstID {
		w.WriteInfo("Correct! Nice counting!")
		w.WriteInfo("OK, next up:")
	}
	w.WriteInfo("")
	w.WriteInfo(fmt.Sprintf("%d %s %d = ?", prob.LeftOperand, prob.Operator, prob.RightOperand))
	w.WriteInfo("")
	type resultLink struct {
		description string
		selector    string
	}
	var results []resultLink
	for _, invRes := range prob.InvalidResults {
		results = append(results, resultLink{
			description: fmt.Sprintf("= %d", invRes.Result),
			selector:    "/math/" + invRes.IDStr,
		})
	}
	if prob.HasNext {
		results = append(results, resultLink{
			description: fmt.Sprintf("= %d", prob.Result),
			selector:    "/math/" + prob.NextIDStr,
		})
	} else {
		results = append(results, resultLink{
			description: fmt.Sprintf("= %d", prob.Result),
			selector:    "/flag/" + h.flagIDStr,
		})
	}
	rand.Shuffle(len(results), func(i, j int) {
		results[i], results[j] = results[j], results[i]
	})
	for _, res := range results {
		writeDirectoryLink(w, h.cfg, res.selector, res.description)
	}
	w.WriteInfo("")
	writeDirectoryLink(w, h.cfg, "/", "Back to start page.")
}

func (h *handlers) handleFakeFlag(w gopher.ResponseWriter, r *gopher.Request) {
	w.WriteInfo("Hah! Nope!")
	w.WriteInfo("You're not getting away that easily.")
	w.WriteInfo("")
	writeDirectoryLink(w, h.cfg, "/", "Back to start page.")
}

func (h *handlers) handleRealFlag(w gopher.ResponseWriter, r *gopher.Request) {
	flagB64 := base64.StdEncoding.EncodeToString([]byte(h.flagStr))
	w.WriteInfo("Good job! Santa's transaction can now finally be committed.")
	w.WriteInfo("It seems to be way late to the party though, and the whole")
	w.WriteInfo("process needs to be repeated.")
	w.WriteInfo("")
	w.WriteInfo("Anyway, the transaction Santa tried to commit was:")
	w.WriteInfo("")
	w.WriteInfo("  TRAN ID:  " + flagB64)
	w.WriteInfo("  AMOUNT:   0,000000521 BTC")
	w.WriteInfo("  RECEIVER: SantaMoneyLaundering AB")
	w.WriteInfo("")
	writeDirectoryLink(w, h.cfg, "/", "Back to start page.")
}

func (h *handlers) handleProtocolInfo(w gopher.ResponseWriter, r *gopher.Request) {
	w.WriteInfo("What you are seeing is an old text transfer protocol.")
	w.WriteInfo("But not that Hyper-Text Transfer Protocol everyone is talking about.")
	w.WriteInfo("")
	w.WriteInfo("This is the predecessor. This is:")
	w.WriteInfo("")
	w.WriteInfo("                   The Internet Gopher Protocol")
	w.WriteInfo("      (a distributed document search and retrieval protocol)")
	w.WriteInfo("")
	w.WriteInfo("A full specification can be found in the IETF RFC-1436.")
	w.WriteInfo("For convinience, we've included the raw-text version in this server:")
	w.WriteInfo("")
	w.WriteItem(&gopher.Item{
		Type:        gopher.FILE,
		Selector:    "/rfc/rfc1436.txt",
		Description: "Open file /rfc/rfc1436.txt",
	})
	w.WriteInfo("")
	w.WriteInfo("For completeness sake, the server side is of course written in Go. :)")
	w.WriteInfo("")
	writeDirectoryLink(w, h.cfg, "/", "Back to start page.")
}

func (h *handlers) handleIndex(w gopher.ResponseWriter, r *gopher.Request) {
	if r.Selector != "/" {
		h.handleNotFound(w, r)
		return
	}
	writeInfoMultiline(w, assetsBanner)
	w.WriteInfo("")
	w.WriteInfo("Santa has lost his APU! Help him calculate the missing")
	w.WriteInfo("equations to finalize his blockchain hash.")
	w.WriteInfo("")
	w.WriteInfo(fmt.Sprintf("Our intel suggests that we have %d equations to solve.", len(h.probs)))
	w.WriteInfo("Get going! Before he loses his transaction!")
	w.WriteInfo("")
	writeDirectoryLink(w, h.cfg, "/protocol", `What protocol is this? Tip: Try sending TCP packet "/protocol<CR><LF>"`)
	writeDirectoryLink(w, h.cfg, "/math/"+h.firstID.String(), "Giv me da math")
	writeDirectoryLink(w, h.cfg, "/flag", "No, just give me the CTF flag")
}

func (h *handlers) handleNotFound(w gopher.ResponseWriter, r *gopher.Request) {
	writeInfoMultiline(w, assetsOhNo)
	w.WriteError(fmt.Sprintf(`Nothing found at "%s".`, r.Selector))
	w.WriteInfo("")
	writeDirectoryLink(w, h.cfg, "/", "Back to start page.")
}

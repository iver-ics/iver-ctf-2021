package main

import (
	"math/rand"
	"os"

	"git.mills.io/prologic/go-gopher"
	"github.com/google/uuid"
	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/iver-wharf/wharf-core/pkg/logger/consolepretty"
)

var log = logger.NewScoped("main")

type mathProbLink struct {
	MathProblem
	ID             uuid.UUID
	IDStr          string
	NextID         uuid.UUID
	NextIDStr      string
	HasNext        bool
	InvalidResults invalidResultList
}

type invalidResult struct {
	Result int64
	ID     uuid.UUID
	IDStr  string
}

type invalidResultList [5]invalidResult

func genInvalidResults() invalidResultList {
	var res invalidResultList
	for i := range res {
		id := uuid.New()
		res[i] = invalidResult{
			Result: RandInt64(-halfInt64Size, halfInt64Size),
			ID:     id,
			IDStr:  id.String(),
		}
	}
	return res
}

func main() {
	logCfg := consolepretty.DefaultConfig
	logCfg.CallerMaxLength = 18
	logCfg.CallerMinLength = 18
	logger.AddOutput(logger.LevelDebug, consolepretty.New(logCfg))

	cfg, err := loadConfig()
	if err != nil {
		log.Error().WithError(err).Message("Failed to load config.")
		os.Exit(1)
	} else {
		log.Debug().
			WithString("flag", cfg.flag).
			Message("Loaded config and flag.")
	}

	if cfg.RngSeed == nil {
		if err := cryptSeedRand(); err != nil {
			log.Error().WithError(err).Message("Failed to seed random store.")
			os.Exit(1)
		}
	} else {
		rand.Seed(*cfg.RngSeed)
	}

	var probs = map[uuid.UUID]mathProbLink{}
	var prevID uuid.UUID
	var prevIDStr string
	for i := 0; i < cfg.Equations; i++ {
		id := uuid.New()
		idStr := id.String()
		probs[id] = mathProbLink{
			MathProblem:    GenMathProblem(),
			ID:             id,
			IDStr:          idStr,
			NextID:         prevID,
			NextIDStr:      prevIDStr,
			HasNext:        prevIDStr != "",
			InvalidResults: genInvalidResults(),
		}
		prevID = id
		prevIDStr = idStr
	}

	flagIDStr := uuid.NewString()

	h := &handlers{
		firstID:    prevID,
		firstIDStr: prevIDStr,
		flagIDStr:  flagIDStr,
		flagStr:    cfg.flag,
		probs:      probs,
		cfg:        cfg,
	}
	gopher.HandleFunc("/", h.handleIndex)
	gopher.HandleFunc("/math/", h.handleMath)
	gopher.HandleFunc("/flag", h.handleFakeFlag)
	gopher.HandleFunc("/flag/"+flagIDStr, h.handleRealFlag)
	gopher.HandleFunc("/protocol", h.handleProtocolInfo)
	gopher.HandleFunc("/rfc/rfc1436.txt", h.handleRFC1436)

	log.Debug().WithStringf("selector", "/flag/%s", flagIDStr).Message("Flag hole")

	log.Info().WithString("address", cfg.BindAddress).Message("Starting Gopher server.")
	if err := gopher.ListenAndServe(cfg.BindAddress, logHandler{}); err != nil {
		log.Error().WithError(err).WithString("address", cfg.BindAddress).Message("Failed to start server.")
		os.Exit(1)
	}
}

type logHandler struct{}

func (logHandler) ServeGopher(w gopher.ResponseWriter, r *gopher.Request) {
	log.Debug().
		WithString("selector", r.Selector).
		Message("Incoming request.")
	gopher.DefaultServeMux.ServeGopher(w, r)
}

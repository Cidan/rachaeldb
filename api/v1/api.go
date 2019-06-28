package v1

import (
	context "context"
	"math/rand"
	"time"

	"github.com/allegro/bigcache"
)

type API struct {
	cache     *bigcache.BigCache
	setSass   []string
	errorSass []string
	getSass   []string
}

func New() *API {
	rand.Seed(time.Now().Unix())
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	return &API{
		cache: cache,
		setSass: []string{
			"Okay, I'll remember that. I guess",
			"Why didn't you put this in a doc?",
			"There's this amazing gSuite product called 'sheets' you should try sometime.",
		},
		errorSass: []string{
			"Look, I don't know what happened. That's your job.",
			"Why didn't you make a backup?",
			"You don't know what happened? That's literally insane.",
		},
		getSass: []string{
			"Here.",
			"Did you try Googling it?",
			"Why didn't you look at the logs?",
			"I THINK this is it.",
		},
	}
}

// Get a key
func (a *API) Get(ctx context.Context, in *Record) (*Record, error) {
	v, err := a.cache.Get(in.Key)
	if err != nil {
		in.Sass = a.errorSass[rand.Intn(len(a.errorSass))]
		return in, err
	}
	in.Sass = a.getSass[rand.Intn(len(a.getSass))]
	in.Data = v
	return in, nil
}

// Set a key
func (a *API) Set(ctx context.Context, in *Record) (*Record, error) {
	err := a.cache.Set(in.Key, in.Data)
	if err != nil {
		in.Sass = a.errorSass[rand.Intn(len(a.errorSass))]
		return in, err
	}
	in.Sass = a.setSass[rand.Intn(len(a.setSass))]
	return in, nil
}

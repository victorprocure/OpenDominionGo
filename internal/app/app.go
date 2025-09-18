package app

import (
	"database/sql"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories/heroes"
	"github.com/victorprocure/opendominiongo/internal/repositories/races"
	"github.com/victorprocure/opendominiongo/internal/repositories/spells"
	"github.com/victorprocure/opendominiongo/internal/repositories/tech"
	"github.com/victorprocure/opendominiongo/internal/repositories/wonders"
)

// App holds repo instances and shared dependencies.
type App struct {
	DB      *sql.DB
	Log     *slog.Logger
	Spells  *spells.SpellsRepo
	Tech    *tech.TechRepo
	Wonders *wonders.WondersRepo
	Races   *races.RacesRepo
	Heroes  *heroes.HeroesRepo
}

// New creates an App by constructing repositories.
func New(db *sql.DB, log *slog.Logger) *App {
	return &App{
		DB:      db,
		Log:     log,
		Spells:  spells.NewSpellsRepository(db, log),
		Tech:    tech.NewTechRepo(db, log),
		Wonders: wonders.NewWondersRepo(db, log),
		Races:   races.NewRacesRepository(db, log),
		Heroes:  heroes.NewHeroesRepo(db, log),
	}
}

// Syncers constructors for main to use
func (a *App) NewTechSync() any        { return nil }
func (a *App) NewRacesSync() any       { return nil }
func (a *App) NewSpellsSync() any      { return nil }
func (a *App) NewWondersSync() any     { return nil }
func (a *App) NewHeroUpgradeSync() any { return nil }

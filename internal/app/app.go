package app

import (
	"database/sql"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/datasync"
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
func (a *App) NewTechSync() datasync.Syncer        { return datasync.NewTechSync(a.DB, a.Log) }
func (a *App) NewRacesSync() datasync.Syncer       { return datasync.NewRacesSync(a.DB, a.Log) }
func (a *App) NewSpellsSync() datasync.Syncer      { return datasync.NewSpellsSync(a.DB, a.Log) }
func (a *App) NewWondersSync() datasync.Syncer     { return datasync.NewWondersSync(a.DB, a.Log) }
func (a *App) NewHeroUpgradeSync() datasync.Syncer { return datasync.NewHeroesSync(a.DB, a.Log) }

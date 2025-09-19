package app

import (
	"database/sql"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/datasync"
	"github.com/victorprocure/opendominiongo/internal/repositories/hero/upgrade"
	racerepo "github.com/victorprocure/opendominiongo/internal/repositories/race"
	spellrepo "github.com/victorprocure/opendominiongo/internal/repositories/spell"
	"github.com/victorprocure/opendominiongo/internal/repositories/tech"
	wonderrepo "github.com/victorprocure/opendominiongo/internal/repositories/wonder"
)

// App holds repo instances and shared dependencies.
type App struct {
	DB      *sql.DB
	Log     *slog.Logger
	Spells  *spellrepo.Repo
	Tech    *tech.Repo
	Wonders *wonderrepo.Repo
	Races   *racerepo.Repo
	Heroes  *upgrade.Repo
}

// New creates an App by constructing repositories.
func New(db *sql.DB, log *slog.Logger) *App {
	return &App{
		DB:      db,
		Log:     log,
		Spells:  spellrepo.NewSpellRepo(db, log),
		Tech:    tech.NewTechRepo(db, log),
		Wonders: wonderrepo.NewWonderRepo(db, log),
		Races:   racerepo.NewRaceRepo(db, log),
		Heroes:  upgrade.NewHeroUpgradeRepo(db, log),
	}
}

// Syncers constructors for main to use
func (a *App) NewTechSync() datasync.Syncer        { return datasync.NewTechSync(a.DB, a.Log) }
func (a *App) NewRacesSync() datasync.Syncer       { return datasync.NewRacesSync(a.DB, a.Log) }
func (a *App) NewSpellsSync() datasync.Syncer      { return datasync.NewSpellsSync(a.DB, a.Log) }
func (a *App) NewWondersSync() datasync.Syncer     { return datasync.NewWondersSync(a.DB, a.Log) }
func (a *App) NewHeroUpgradeSync() datasync.Syncer { return datasync.NewHeroesSync(a.DB, a.Log) }

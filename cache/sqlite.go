package cache

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	Instance *sqlx.DB
)

// Init initializes the database
func dbliteInit(ctx context.Context) error {
	err := connect()
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}

	err = os.MkdirAll("cache/", 0755)
	if err != nil {
		return fmt.Errorf("make cache: %w", err)
	}
	scopes := []string{
		"npc",
		"item",
		"player",
		"npc_loot",
		"npc_merchant",
		"npc_spawn",
		"npc_faction",
		"npc_spell",
	}

	for _, scope := range scopes {
		query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (key INTEGER PRIMARY KEY, data TEXT, expiration INTEGER)", scope)
		_, err = Instance.ExecContext(ctx, query)
		if err != nil {
			return fmt.Errorf("create table: %w", err)
		}
	}

	return nil
}

func connect() error {
	var err error
	var db *sqlx.DB
	tlog.Debugf("Connecting to sqlite database cache/sqlite.db")

	db, err = sqlx.Open("sqlite3", "cache/sqlite.db")
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}
	db.SetMaxOpenConns(1)
	Instance = db
	return nil
}

func readSqliteCache(path string) (model.CacheIdentifier, bool) {
	if !config.Get().SqliteCache.IsEnabled {
		return nil, false
	}

	var data model.CacheIdentifier

	scope := strings.TrimPrefix(path, "cache/")
	key := scope
	if strings.Contains(scope, "/") {
		records := strings.Split(scope, "/")
		scope = records[0]
		key = records[1]
	}
	key = strings.TrimSuffix(key, ".yaml")

	switch scope {
	case "npc":
		data = &model.Npc{}
	case "item":
		data = &model.Item{}
	case "player":
		data = &model.Player{}
	case "npc_loot":
		data = &model.NpcLoot{}
	case "npc_merchant":
		data = &model.NpcMerchant{}
	case "npc_spawn":
		data = &model.NpcSpawn{}
	case "npc_faction":
		data = &model.NpcFaction{}
	case "npc_spell":
		data = &model.NpcSpell{}
	default:
		tlog.Warnf("Unknown cache scope: %s", scope)
		return nil, false
	}

	query := fmt.Sprintf("SELECT data FROM %s WHERE key = :key AND expiration > :expiration", scope)

	row := Instance.QueryRowxContext(context.Background(),
		query,
		key,
		time.Now().Unix())
	if row.Err() != nil {
		tlog.Warnf("Query cache: %v", row.Err())
		return nil, false
	}

	rawData := ""
	err := row.Scan(&rawData)
	if err != nil {
		tlog.Warnf("rows.StructScan: %v", err)
		return nil, false
	}

	err = data.Deserialize(rawData)
	if err != nil {
		tlog.Warnf("Deserialize: %v", err)
		return nil, false
	}
	return data, true
}

func writeSqliteCache(ctx context.Context, path string, data model.CacheIdentifier) error {
	if !config.Get().SqliteCache.IsEnabled {
		return nil
	}

	scope := strings.TrimPrefix(path, "cache/")
	key := scope
	if strings.Contains(scope, "/") {
		records := strings.Split(scope, "/")
		scope = records[0]
		key = records[1]
	}
	key = strings.TrimSuffix(key, ".yaml")
	tlog.Debugf("Sqlite read: %s %s", scope, key)

	query := fmt.Sprintf("REPLACE INTO %s (key, data, expiration) VALUES (?, ?, ?)", data.Identifier())

	data.SetExpiration(int64(time.Now().Add(time.Minute * time.Duration(config.Get().SqliteCache.Expiration)).Unix()))

	_, err := Instance.ExecContext(ctx,
		query,
		key,
		data.Serialize(),
		data.Expiration(),
	)
	if err != nil {
		return err
	}

	tlog.Debugf("Sqlitecache write: %s %s", scope, key)
	return nil
}

func truncateSqliteCache() {
	if !config.Get().SqliteCache.IsEnabled {
		return
	}
	time.Sleep(time.Duration(config.Get().SqliteCache.TruncateSchedule))
	tlog.Debugf("Sqlitecache truncate schedule running...")
	start := time.Now()

	scopes := []string{"npc", "item", "player"}

	for _, scope := range scopes {
		query := fmt.Sprintf("DELETE FROM %s WHERE expiration < :expiration", scope)
		_, err := Instance.ExecContext(context.Background(),

			query,
			map[string]interface{}{
				"expiration": time.Now().Unix(),
			})
		if err != nil {
			tlog.Warnf("Query cache: %v", err)
			return
		}
	}

	tlog.Debugf("Sqlitecache truncate schedule complete in %s", time.Since(start))
}

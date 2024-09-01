// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sqlite-cache-queries.sql

package sqlitecachec

import (
	"context"
)

const itemByItemID = `-- name: ItemByItemID :one
SELECT gob FROM item WHERE item_id = ? AND expiration > ?
`

type ItemByItemIDParams struct {
	ItemID     int64
	Expiration int64
}

func (q *Queries) ItemByItemID(ctx context.Context, arg ItemByItemIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, itemByItemID, arg.ItemID, arg.Expiration)
	var gob string
	err := row.Scan(&gob)
	return gob, err
}

const itemCreate = `-- name: ItemCreate :exec
CREATE TABLE IF NOT EXISTS item (item_id INTEGER PRIMARY KEY, gob TEXT NOT NULL, expiration INTEGER NOT NULL)
`

func (q *Queries) ItemCreate(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, itemCreate)
	return err
}

const itemRecipeByItemID = `-- name: ItemRecipeByItemID :one
SELECT gob FROM item_recipe WHERE item_id = ? AND expiration > ?
`

type ItemRecipeByItemIDParams struct {
	ItemID     int64
	Expiration int64
}

func (q *Queries) ItemRecipeByItemID(ctx context.Context, arg ItemRecipeByItemIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, itemRecipeByItemID, arg.ItemID, arg.Expiration)
	var gob string
	err := row.Scan(&gob)
	return gob, err
}

const itemRecipeCreate = `-- name: ItemRecipeCreate :exec
CREATE TABLE IF NOT EXISTS item_recipe (item_id INTEGER PRIMARY KEY, gob TEXT NOT NULL, expiration INTEGER NOT NULL)
`

func (q *Queries) ItemRecipeCreate(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, itemRecipeCreate)
	return err
}

const itemRecipeReplace = `-- name: ItemRecipeReplace :exec
REPLACE INTO item_recipe (item_id, gob, expiration) VALUES (?, ?, ?)
`

type ItemRecipeReplaceParams struct {
	ItemID     int64
	Gob        string
	Expiration int64
}

func (q *Queries) ItemRecipeReplace(ctx context.Context, arg ItemRecipeReplaceParams) error {
	_, err := q.db.ExecContext(ctx, itemRecipeReplace, arg.ItemID, arg.Gob, arg.Expiration)
	return err
}

const itemRecipeTruncate = `-- name: ItemRecipeTruncate :exec
DELETE FROM item_recipe WHERE expiration < ?
`

func (q *Queries) ItemRecipeTruncate(ctx context.Context, expiration int64) error {
	_, err := q.db.ExecContext(ctx, itemRecipeTruncate, expiration)
	return err
}

const itemReplace = `-- name: ItemReplace :exec
REPLACE INTO item (item_id, gob, expiration) VALUES (?, ?, ?)
`

type ItemReplaceParams struct {
	ItemID     int64
	Gob        string
	Expiration int64
}

func (q *Queries) ItemReplace(ctx context.Context, arg ItemReplaceParams) error {
	_, err := q.db.ExecContext(ctx, itemReplace, arg.ItemID, arg.Gob, arg.Expiration)
	return err
}

const itemTruncate = `-- name: ItemTruncate :exec
DELETE FROM item WHERE expiration < ?
`

func (q *Queries) ItemTruncate(ctx context.Context, expiration int64) error {
	_, err := q.db.ExecContext(ctx, itemTruncate, expiration)
	return err
}

const npcByNpcID = `-- name: NpcByNpcID :one
SELECT gob FROM npc WHERE npc_id = ? AND expiration > ?
`

type NpcByNpcIDParams struct {
	NpcID      int64
	Expiration int64
}

func (q *Queries) NpcByNpcID(ctx context.Context, arg NpcByNpcIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, npcByNpcID, arg.NpcID, arg.Expiration)
	var gob string
	err := row.Scan(&gob)
	return gob, err
}

const npcCreate = `-- name: NpcCreate :exec
CREATE TABLE IF NOT EXISTS npc (npc_id INTEGER PRIMARY KEY, gob TEXT NOT NULL, expiration INTEGER NOT NULL)
`

func (q *Queries) NpcCreate(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, npcCreate)
	return err
}

const npcFactionByFactionID = `-- name: NpcFactionByFactionID :one
SELECT gob FROM npc_faction WHERE faction_id = ? AND expiration > ?
`

type NpcFactionByFactionIDParams struct {
	FactionID  int64
	Expiration int64
}

func (q *Queries) NpcFactionByFactionID(ctx context.Context, arg NpcFactionByFactionIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, npcFactionByFactionID, arg.FactionID, arg.Expiration)
	var gob string
	err := row.Scan(&gob)
	return gob, err
}

const npcFactionCreate = `-- name: NpcFactionCreate :exec
CREATE TABLE IF NOT EXISTS npc_faction (faction_id INTEGER PRIMARY KEY, gob TEXT NOT NULL, expiration INTEGER NOT NULL)
`

func (q *Queries) NpcFactionCreate(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, npcFactionCreate)
	return err
}

const npcFactionReplace = `-- name: NpcFactionReplace :exec
REPLACE INTO npc_faction (faction_id, gob, expiration) VALUES (?, ?, ?)
`

type NpcFactionReplaceParams struct {
	FactionID  int64
	Gob        string
	Expiration int64
}

func (q *Queries) NpcFactionReplace(ctx context.Context, arg NpcFactionReplaceParams) error {
	_, err := q.db.ExecContext(ctx, npcFactionReplace, arg.FactionID, arg.Gob, arg.Expiration)
	return err
}

const npcFactionTruncate = `-- name: NpcFactionTruncate :exec
DELETE FROM npc_faction WHERE expiration < ?
`

func (q *Queries) NpcFactionTruncate(ctx context.Context, expiration int64) error {
	_, err := q.db.ExecContext(ctx, npcFactionTruncate, expiration)
	return err
}

const npcLootByLootTableID = `-- name: NpcLootByLootTableID :one
SELECT gob FROM npc_loot WHERE loot_table_id = ? AND expiration > ?
`

type NpcLootByLootTableIDParams struct {
	LootTableID int64
	Expiration  int64
}

func (q *Queries) NpcLootByLootTableID(ctx context.Context, arg NpcLootByLootTableIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, npcLootByLootTableID, arg.LootTableID, arg.Expiration)
	var gob string
	err := row.Scan(&gob)
	return gob, err
}

const npcLootCreate = `-- name: NpcLootCreate :exec
CREATE TABLE IF NOT EXISTS npc_loot (loot_table_id INTEGER PRIMARY KEY, gob TEXT NOT NULL, expiration INTEGER NOT NULL)
`

func (q *Queries) NpcLootCreate(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, npcLootCreate)
	return err
}

const npcLootReplace = `-- name: NpcLootReplace :exec
REPLACE INTO npc_loot (loot_table_id, gob, expiration) VALUES (?, ?, ?)
`

type NpcLootReplaceParams struct {
	LootTableID int64
	Gob         string
	Expiration  int64
}

func (q *Queries) NpcLootReplace(ctx context.Context, arg NpcLootReplaceParams) error {
	_, err := q.db.ExecContext(ctx, npcLootReplace, arg.LootTableID, arg.Gob, arg.Expiration)
	return err
}

const npcLootTruncate = `-- name: NpcLootTruncate :exec
DELETE FROM npc_loot WHERE expiration < ?
`

func (q *Queries) NpcLootTruncate(ctx context.Context, expiration int64) error {
	_, err := q.db.ExecContext(ctx, npcLootTruncate, expiration)
	return err
}

const npcMerchantByNpcID = `-- name: NpcMerchantByNpcID :one
SELECT gob FROM npc_merchant WHERE npc_id = ? AND expiration > ?
`

type NpcMerchantByNpcIDParams struct {
	NpcID      int64
	Expiration int64
}

func (q *Queries) NpcMerchantByNpcID(ctx context.Context, arg NpcMerchantByNpcIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, npcMerchantByNpcID, arg.NpcID, arg.Expiration)
	var gob string
	err := row.Scan(&gob)
	return gob, err
}

const npcMerchantCreate = `-- name: NpcMerchantCreate :exec
CREATE TABLE IF NOT EXISTS npc_merchant (npc_id INTEGER PRIMARY KEY, gob TEXT NOT NULL, expiration INTEGER NOT NULL)
`

func (q *Queries) NpcMerchantCreate(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, npcMerchantCreate)
	return err
}

const npcMerchantReplace = `-- name: NpcMerchantReplace :exec
REPLACE INTO npc_merchant (npc_id, gob, expiration) VALUES (?, ?, ?)
`

type NpcMerchantReplaceParams struct {
	NpcID      int64
	Gob        string
	Expiration int64
}

func (q *Queries) NpcMerchantReplace(ctx context.Context, arg NpcMerchantReplaceParams) error {
	_, err := q.db.ExecContext(ctx, npcMerchantReplace, arg.NpcID, arg.Gob, arg.Expiration)
	return err
}

const npcMerchantTruncate = `-- name: NpcMerchantTruncate :exec
DELETE FROM npc_merchant WHERE expiration < ?
`

func (q *Queries) NpcMerchantTruncate(ctx context.Context, expiration int64) error {
	_, err := q.db.ExecContext(ctx, npcMerchantTruncate, expiration)
	return err
}

const npcQuestByNpcID = `-- name: NpcQuestByNpcID :one
SELECT gob FROM npc_quest WHERE npc_id = ? AND expiration > ?
`

type NpcQuestByNpcIDParams struct {
	NpcID      int64
	Expiration int64
}

func (q *Queries) NpcQuestByNpcID(ctx context.Context, arg NpcQuestByNpcIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, npcQuestByNpcID, arg.NpcID, arg.Expiration)
	var gob string
	err := row.Scan(&gob)
	return gob, err
}

const npcQuestCreate = `-- name: NpcQuestCreate :exec
CREATE TABLE IF NOT EXISTS npc_quest (npc_id INTEGER PRIMARY KEY, gob TEXT NOT NULL, expiration INTEGER NOT NULL)
`

func (q *Queries) NpcQuestCreate(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, npcQuestCreate)
	return err
}

const npcQuestReplace = `-- name: NpcQuestReplace :exec
REPLACE INTO npc_quest (npc_id, gob, expiration) VALUES (?, ?, ?)
`

type NpcQuestReplaceParams struct {
	NpcID      int64
	Gob        string
	Expiration int64
}

func (q *Queries) NpcQuestReplace(ctx context.Context, arg NpcQuestReplaceParams) error {
	_, err := q.db.ExecContext(ctx, npcQuestReplace, arg.NpcID, arg.Gob, arg.Expiration)
	return err
}

const npcQuestTruncate = `-- name: NpcQuestTruncate :exec
DELETE FROM npc_quest WHERE expiration < ?
`

func (q *Queries) NpcQuestTruncate(ctx context.Context, expiration int64) error {
	_, err := q.db.ExecContext(ctx, npcQuestTruncate, expiration)
	return err
}

const npcReplace = `-- name: NpcReplace :exec
REPLACE INTO npc (npc_id, gob, expiration) VALUES (?, ?, ?)
`

type NpcReplaceParams struct {
	NpcID      int64
	Gob        string
	Expiration int64
}

func (q *Queries) NpcReplace(ctx context.Context, arg NpcReplaceParams) error {
	_, err := q.db.ExecContext(ctx, npcReplace, arg.NpcID, arg.Gob, arg.Expiration)
	return err
}

const npcSpawnByNpcID = `-- name: NpcSpawnByNpcID :one
SELECT gob FROM npc_spawn WHERE npc_id = ? AND expiration > ?
`

type NpcSpawnByNpcIDParams struct {
	NpcID      int64
	Expiration int64
}

func (q *Queries) NpcSpawnByNpcID(ctx context.Context, arg NpcSpawnByNpcIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, npcSpawnByNpcID, arg.NpcID, arg.Expiration)
	var gob string
	err := row.Scan(&gob)
	return gob, err
}

const npcSpawnCreate = `-- name: NpcSpawnCreate :exec
CREATE TABLE IF NOT EXISTS npc_spawn (npc_id INTEGER PRIMARY KEY, gob TEXT NOT NULL, expiration INTEGER NOT NULL)
`

func (q *Queries) NpcSpawnCreate(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, npcSpawnCreate)
	return err
}

const npcSpawnTruncate = `-- name: NpcSpawnTruncate :exec
DELETE FROM npc_spawn WHERE expiration < ?
`

func (q *Queries) NpcSpawnTruncate(ctx context.Context, expiration int64) error {
	_, err := q.db.ExecContext(ctx, npcSpawnTruncate, expiration)
	return err
}

const npcSpellByNpcSpellsID = `-- name: NpcSpellByNpcSpellsID :one
SELECT gob FROM npc_spell WHERE npc_spells_id = ? AND expiration > ?
`

type NpcSpellByNpcSpellsIDParams struct {
	NpcSpellsID int64
	Expiration  int64
}

func (q *Queries) NpcSpellByNpcSpellsID(ctx context.Context, arg NpcSpellByNpcSpellsIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, npcSpellByNpcSpellsID, arg.NpcSpellsID, arg.Expiration)
	var gob string
	err := row.Scan(&gob)
	return gob, err
}

const npcSpellCreate = `-- name: NpcSpellCreate :exec
CREATE TABLE IF NOT EXISTS npc_spell (npc_spells_id INTEGER PRIMARY KEY, gob TEXT NOT NULL, expiration INTEGER NOT NULL)
`

func (q *Queries) NpcSpellCreate(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, npcSpellCreate)
	return err
}

const npcSpellTruncate = `-- name: NpcSpellTruncate :exec
DELETE FROM npc_spell WHERE expiration < ?
`

func (q *Queries) NpcSpellTruncate(ctx context.Context, expiration int64) error {
	_, err := q.db.ExecContext(ctx, npcSpellTruncate, expiration)
	return err
}

const npcTruncate = `-- name: NpcTruncate :exec
DELETE FROM npc WHERE expiration < ?
`

func (q *Queries) NpcTruncate(ctx context.Context, expiration int64) error {
	_, err := q.db.ExecContext(ctx, npcTruncate, expiration)
	return err
}

const playerByCharacterID = `-- name: PlayerByCharacterID :one
SELECT gob FROM player WHERE character_id = ? AND expiration > ?
`

type PlayerByCharacterIDParams struct {
	CharacterID int64
	Expiration  int64
}

func (q *Queries) PlayerByCharacterID(ctx context.Context, arg PlayerByCharacterIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, playerByCharacterID, arg.CharacterID, arg.Expiration)
	var gob string
	err := row.Scan(&gob)
	return gob, err
}

const playerCreate = `-- name: PlayerCreate :exec
CREATE TABLE IF NOT EXISTS player (character_id INTEGER PRIMARY KEY, gob TEXT NOT NULL, expiration INTEGER NOT NULL)
`

func (q *Queries) PlayerCreate(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, playerCreate)
	return err
}

const playerReplace = `-- name: PlayerReplace :exec
REPLACE INTO player (character_id, gob, expiration) VALUES (?, ?, ?)
`

type PlayerReplaceParams struct {
	CharacterID int64
	Gob         string
	Expiration  int64
}

func (q *Queries) PlayerReplace(ctx context.Context, arg PlayerReplaceParams) error {
	_, err := q.db.ExecContext(ctx, playerReplace, arg.CharacterID, arg.Gob, arg.Expiration)
	return err
}

const playerTruncate = `-- name: PlayerTruncate :exec
DELETE FROM player WHERE expiration < ?
`

func (q *Queries) PlayerTruncate(ctx context.Context, expiration int64) error {
	_, err := q.db.ExecContext(ctx, playerTruncate, expiration)
	return err
}

const questByID = `-- name: QuestByID :one
SELECT gob FROM quest WHERE quest_id = ? AND expiration > ?
`

type QuestByIDParams struct {
	QuestID    int64
	Expiration int64
}

func (q *Queries) QuestByID(ctx context.Context, arg QuestByIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, questByID, arg.QuestID, arg.Expiration)
	var gob string
	err := row.Scan(&gob)
	return gob, err
}

const questCreate = `-- name: QuestCreate :exec
CREATE TABLE IF NOT EXISTS quest (quest_id INTEGER PRIMARY KEY, gob TEXT NOT NULL, expiration INTEGER NOT NULL)
`

func (q *Queries) QuestCreate(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, questCreate)
	return err
}

const questTruncate = `-- name: QuestTruncate :exec
DELETE FROM quest WHERE expiration < ?
`

func (q *Queries) QuestTruncate(ctx context.Context, expiration int64) error {
	_, err := q.db.ExecContext(ctx, questTruncate, expiration)
	return err
}
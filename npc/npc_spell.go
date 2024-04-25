package npc

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

func fetchNpcSpell(ctx context.Context, id int, level int) (*model.NpcSpell, error) {
	path := fmt.Sprintf("npc_spell/%d.yaml", id)
	cacheData, src, ok := cache.Read(ctx, path)
	if ok {
		cacheNpc, ok := cacheData.(*model.NpcSpell)
		if !ok {
			return nil, fmt.Errorf("cache read: invalid type, wanted *model.NpcSpell, got %T", cacheData)
		}
		if cacheNpc != nil {
			if src != cache.SourceCacheMemory {
				err := cache.WriteMemoryCache(ctx, path, cacheNpc)
				if err != nil {
					return nil, fmt.Errorf("cache write: %w", err)
				}
			}
			return cacheNpc, nil
		}
	}

	query := `SELECT id, name, proc_chance, attack_proc, range_proc, rproc_chance, defensive_proc, dproc_chance FROM npc_spells ns WHERE id=:id`

	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return nil, fmt.Errorf("query npcs: %w", err)
	}
	defer rows.Close()

	npcSpell := &model.NpcSpell{}

	for rows.Next() {
		npcSpellEntry := &model.NpcSpellEntry{}
		err = rows.StructScan(npcSpellEntry)
		if err != nil {
			return nil, fmt.Errorf("rows.StructScan: %w", err)
		}
		subEntries, err := fetchNpcSpellEntries(ctx, npcSpellEntry.ID, level)
		if err != nil {
			return nil, fmt.Errorf("fetchNpcSpellEntries: %w", err)
		}
		for _, subEntry := range subEntries.Entries {
			subEntry.ID = npcSpellEntry.ID
			subEntry.Name = npcSpellEntry.Name
			subEntry.Procchance = npcSpellEntry.Procchance
			subEntry.Attackproc = npcSpellEntry.Attackproc
			subEntry.Rangeproc = npcSpellEntry.Rangeproc
			subEntry.Rprocchance = npcSpellEntry.Rprocchance
			subEntry.Defensiveproc = npcSpellEntry.Defensiveproc
			subEntry.Dprocchance = npcSpellEntry.Dprocchance
			npcSpell.Entries = append(npcSpell.Entries, subEntry)
		}
	}

	err = cache.Write(ctx, path, npcSpell)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return npcSpell, nil
}

func fetchNpcSpellEntries(ctx context.Context, id int, level int) (*model.NpcSpell, error) {
	query := `SELECT nse.spellid 
FROM npc_spells_entries nse
WHERE nse.npc_spells_id = :id
AND nse.minlevel <= :level
AND nse.maxlevel >= :level
ORDER BY nse.priority DESC`

	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"id":    id,
			"level": level,
		})
	if err != nil {
		return nil, fmt.Errorf("query npcs: %w", err)
	}
	defer rows.Close()

	npcSpell := &model.NpcSpell{}

	for rows.Next() {
		npcSpellEntry := &model.NpcSpellEntry{}
		err = rows.StructScan(npcSpellEntry)
		if err != nil {
			return nil, fmt.Errorf("rows.StructScan: %w", err)
		}
		npcSpell.Entries = append(npcSpell.Entries, npcSpellEntry)
	}

	return npcSpell, nil
}

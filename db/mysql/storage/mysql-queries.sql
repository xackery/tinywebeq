-- name: ItemByID :one
SELECT * FROM items WHERE id = ? LIMIT 1;

-- name: ItemsAll :many
SELECT id, `name`, ac, reqlevel, reclevel, hp, damage, delay, mana FROM items;

-- name: ItemDiscoveredOnlyByID :one
SELECT * FROM items i, discovered_items di WHERE i.id=di.id AND di.item_id=? LIMIT 1;

-- name: ItemRecipeAll :many
SELECT tr.id recipe_id, tr.name recipe_name, 
tr.tradeskill, tr.trivial, tre.item_id, tre.iscontainer is_container,
tre.componentcount component_count, tre.successcount success_count
FROM tradeskill_recipe tr, tradeskill_recipe_entries tre
WHERE tr.id = tre.recipe_id
AND tr.enabled = 1
AND tre.componentcount > 0
ORDER by tre.item_id ASC;

-- name: ItemSearchByName :many
SELECT * FROM items WHERE `name` LIKE ? ORDER BY `name` ASC;

-- name: SpellsAll :many
SELECT * FROM spells_new;
-- name: SpellByID :one
SELECT * FROM spells_new WHERE id = ?;

-- name: SpellSearchByName :many
SELECT * FROM spells_new WHERE name LIKE ?;

-- name: NpcByNpcID :one
SELECT * FROM npc_types WHERE id = ? LIMIT 1;

-- name: NpcsAll :many
SELECT id, name, level FROM npc_types;

-- name: NpcSearchByName :many
SELECT id, name, level FROM npc_types WHERE name LIKE ? ORDER BY name ASC;

-- name: NpcFactionsByFactionID :many
SELECT fl.name, fl.id, fe.value
FROM faction_list fl, npc_faction_entries fe
WHERE fe.npc_faction_id = ?
AND fe.faction_id = fl.id
GROUP BY fl.id
ORDER BY fe.value DESC;

-- name: NpcLootsByLootTableID :many
SELECT i.id, i.name, i.itemtype, lde.chance, lte.probability, lte.lootdrop_id, lte.multiplier
FROM items i, loottable_entries lte, lootdrop_entries lde
WHERE lte.loottable_id = ?
AND lte.lootdrop_id = lde.lootdrop_id
AND lde.item_id = i.id;

-- name: NpcLootsDiscoveredOnlyByLootTableID :many
SELECT i.id, i.name, i.itemtype, lde.chance, lte.probability, lte.lootdrop_id, lte.multiplier
FROM items i, loottable_entries lte, lootdrop_entries lde, discovered_items di
WHERE lte.loottable_id = ?
AND lte.lootdrop_id = lde.lootdrop_id
AND lde.item_id = i.id AND di.item_id = i.id;

-- name: NpcMerchantsByMerchantID :many
SELECT i.id, i.Name name, i.price, i.ldonprice, i.icon
FROM items i, merchantlist ml
WHERE ml.merchantid = ?
AND ml.item = i.id
ORDER BY ml.slot;

-- name: NpcSpawnsByNpcID :many
SELECT z.long_name, z.short_name, s2.x, s2.y, s2.z, sg.name AS spawngroup, sg.id AS spawngroupid, s2.respawntime
FROM zone z, spawnentry se, spawn2 s2, spawngroup sg
WHERE se.npcID = ?
AND se.spawngroupID = s2.spawngroupID
AND s2.zone = z.short_name
AND se.spawngroupID = sg.id
ORDER BY z.long_name;

-- name: NpcSpellsByNpcSpellID :many
SELECT id, name, proc_chance, attack_proc, range_proc, rproc_chance, defensive_proc, dproc_chance
FROM npc_spells ns
WHERE id=?;
-- name: NpcSpellEntriesByNpcSpellID :many
SELECT nse.spellid 
FROM npc_spells_entries nse
WHERE nse.npc_spells_id = ?
AND nse.minlevel <= ?
AND nse.maxlevel >= ?
ORDER BY nse.priority DESC;


-- name: PlayerByCharacterID :one
SELECT * FROM character_data WHERE id = ? LIMIT 1;


-- name: ZoneByShortName :one
SELECT * FROM zone WHERE short_name = ? AND expansion <= ?;
-- name: ZonesAll :many
SELECT * FROM zone WHERE expansion <= ? ORDER by short_name ASC;
-- name: ZoneByZoneIDNumber :one
SELECT * FROM zone WHERE zoneidnumber = ? AND expansion <= ?; 

-- name: ZoneSearchByName :many
SELECT * FROM zone WHERE (short_name LIKE ? OR long_name LIKE ?) AND expansion <= ? ORDER BY short_name ASC;
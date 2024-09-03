package repo

import (
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"

	"github.com/xackery/tinywebeq/models"
)

const itemByID = `-- name: ItemByID :one
SELECT id, minstatus, name, aagi, ac, accuracy, acha, adex, aint, artifactflag, asta, astr, attack, augrestrict, augslot1type, augslot1visible, augslot2type, augslot2visible, augslot3type, augslot3visible, augslot4type, augslot4visible, augslot5type, augslot5visible, augslot6type, augslot6visible, augtype, avoidance, awis, bagsize, bagslots, bagtype, bagwr, banedmgamt, banedmgraceamt, banedmgbody, banedmgrace, bardtype, bardvalue, book, casttime, charmfile, charmfileid, classes, color, combateffects, extradmgskill, extradmgamt, price, cr, damage, damageshield, deity, delay, augdistiller, dotshielding, dr, clicktype, clicklevel2, elemdmgtype, elemdmgamt, endur, factionamt1, factionamt2, factionamt3, factionamt4, factionmod1, factionmod2, factionmod3, factionmod4, filename, focuseffect, fr, fvnodrop, haste, clicklevel, hp, regen, icon, idfile, itemclass, itemtype, ldonprice, ldontheme, ldonsold, light, lore, loregroup, magic, mana, manaregen, enduranceregen, material, herosforgemodel, maxcharges, mr, nodrop, norent, pendingloreflag, pr, procrate, races, ` + "`" + `range` + "`" + `, reclevel, recskill, reqlevel, sellrate, shielding, size, skillmodtype, skillmodvalue, slots, clickeffect, spellshield, strikethrough, stunresist, summonedflag, tradeskills, favor, weight, unk012, unk013, benefitflag, unk054, unk059, booktype, recastdelay, recasttype, guildfavor, unk123, unk124, attuneable, nopet, updated, comment, unk127, pointtype, potionbelt, potionbeltslots, stacksize, notransfer, stackable, unk134, unk137, proceffect, proctype, proclevel2, proclevel, unk142, worneffect, worntype, wornlevel2, wornlevel, unk147, focustype, focuslevel2, focuslevel, unk152, scrolleffect, scrolltype, scrolllevel2, scrolllevel, unk157, serialized, verified, serialization, source, unk033, lorefile, unk014, svcorruption, skillmodmax, unk060, augslot1unk2, augslot2unk2, augslot3unk2, augslot4unk2, augslot5unk2, augslot6unk2, unk120, unk121, questitemflag, unk132, clickunk5, clickunk6, clickunk7, procunk1, procunk2, procunk3, procunk4, procunk6, procunk7, wornunk1, wornunk2, wornunk3, wornunk4, wornunk5, wornunk6, wornunk7, focusunk1, focusunk2, focusunk3, focusunk4, focusunk5, focusunk6, focusunk7, scrollunk1, scrollunk2, scrollunk3, scrollunk4, scrollunk5, scrollunk6, scrollunk7, unk193, purity, evoitem, evoid, evolvinglevel, evomax, clickname, procname, wornname, focusname, scrollname, dsmitigation, heroic_str, heroic_int, heroic_wis, heroic_agi, heroic_dex, heroic_sta, heroic_cha, heroic_pr, heroic_dr, heroic_fr, heroic_cr, heroic_mr, heroic_svcorrup, healamt, spelldmg, clairvoyance, backstabdmg, created, elitematerial, ldonsellbackrate, scriptfileid, expendablearrow, powersourcecapacity, bardeffect, bardeffecttype, bardlevel2, bardlevel, bardunk1, bardunk2, bardunk3, bardunk4, bardunk5, bardname, bardunk7, unk214, subtype, unk220, unk221, heirloom, unk223, unk224, unk225, unk226, unk227, unk228, unk229, unk230, unk231, unk232, unk233, unk234, placeable, unk236, unk237, unk238, unk239, unk240, unk241, epicitem FROM items WHERE id = ? LIMIT 1
`

func (r *Repo) ItemByID(ctx context.Context, id int64) (*models.Item, error) {
	var err error

	r.logger.Debug("Executing ItemByID", zap.Int64("id", id))

	item := models.Item{}
	if err = r.db.QueryRowxContext(ctx, itemByID, id).StructScan(&item); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		r.logger.Error("error scanning item", err)
		return nil, err
	}

	return &item, nil
}

const itemDiscoveredOnlyByID = `
-- name: ItemDiscoveredOnlyByID :one
SELECT 
items.id, minstatus, name, aagi, ac, accuracy, acha, adex, aint, artifactflag, asta, astr, attack, augrestrict, augslot1type, 
augslot1visible, augslot2type, augslot2visible, augslot3type, augslot3visible, augslot4type, augslot4visible, 
augslot5type, augslot5visible, augslot6type, augslot6visible, augtype, avoidance, awis, bagsize, bagslots, bagtype, 
bagwr, banedmgamt, banedmgraceamt, banedmgbody, banedmgrace, bardtype, bardvalue, book, casttime, charmfile, 
charmfileid, classes, color, combateffects, extradmgskill, extradmgamt, price, cr, damage, damageshield, deity, delay, 
augdistiller, dotshielding, dr, clicktype, clicklevel2, elemdmgtype, elemdmgamt, endur, factionamt1, factionamt2, 
factionamt3, factionamt4, factionmod1, factionmod2, factionmod3, factionmod4, filename, focuseffect, fr, fvnodrop, 
haste, clicklevel, hp, regen, icon, idfile, itemclass, itemtype, ldonprice, ldontheme, ldonsold, light, lore, 
loregroup, magic, mana, manaregen, enduranceregen, material, herosforgemodel, maxcharges, mr, nodrop, norent, 
pendingloreflag, pr, procrate, races, ` + "`" + `range` + "`" + `, reclevel, recskill, reqlevel, sellrate, shielding, 
size, skillmodtype, skillmodvalue, slots, clickeffect, spellshield, strikethrough, stunresist, summonedflag, 
tradeskills, favor, weight, unk012, unk013, benefitflag, unk054, unk059, booktype, recastdelay, recasttype, guildfavor, 
unk123, unk124, attuneable, nopet, updated, comment, unk127, pointtype, potionbelt, potionbeltslots, stacksize, 
notransfer, stackable, unk134, unk137, proceffect, proctype, proclevel2, proclevel, unk142, worneffect, worntype, 
wornlevel2, wornlevel, unk147, focustype, focuslevel2, focuslevel, unk152, scrolleffect, scrolltype, scrolllevel2, 
scrolllevel, unk157, serialized, verified, serialization, source, unk033, lorefile, unk014, svcorruption, skillmodmax, 
unk060, augslot1unk2, augslot2unk2, augslot3unk2, augslot4unk2, augslot5unk2, augslot6unk2, unk120, unk121, 
questitemflag, unk132, clickunk5, clickunk6, clickunk7, procunk1, procunk2, procunk3, procunk4, procunk6, procunk7, 
wornunk1, wornunk2, wornunk3, wornunk4, wornunk5, wornunk6, wornunk7, focusunk1, focusunk2, focusunk3, focusunk4, 
focusunk5, focusunk6, focusunk7, scrollunk1, scrollunk2, scrollunk3, scrollunk4, scrollunk5, scrollunk6, scrollunk7, 
unk193, purity, evoitem, evoid, evolvinglevel, evomax, clickname, procname, wornname, focusname, scrollname, 
dsmitigation, items.heroic_str, heroic_int, heroic_wis, heroic_agi, heroic_dex, heroic_sta, heroic_cha, heroic_pr, heroic_dr, 
heroic_fr, heroic_cr, heroic_mr, heroic_svcorrup, healamt, spelldmg, clairvoyance, backstabdmg, created, elitematerial, 
ldonsellbackrate, scriptfileid, expendablearrow, powersourcecapacity, bardeffect, bardeffecttype, bardlevel2, 
bardlevel, bardunk1, bardunk2, bardunk3, bardunk4, bardunk5, bardname, bardunk7, unk214, subtype, unk220, unk221, 
heirloom, unk223, unk224, unk225, unk226, unk227, unk228, unk229, unk230, unk231, unk232, unk233, unk234, placeable, 
unk236, unk237, unk238, unk239, unk240, unk241, epicitem, item_id, char_name, discovered_date, account_status 
FROM
items, discovered_items 
WHERE items.id=discovered_items.item_id 
AND discovered_items.item_id=? 
LIMIT 1
`

func (r *Repo) DiscoveredItemByID(ctx context.Context, id int64) (*models.DiscoveredItem, error) {
	var err error

	r.logger.Debug("Executing DiscoveredItemByID ", "id ", id)

	item := models.DiscoveredItem{}
	if err = r.db.QueryRowxContext(ctx, itemDiscoveredOnlyByID, id).StructScan(&item); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if item.Item == nil {
		return nil, ErrNotFound
	}

	return &item, nil
}

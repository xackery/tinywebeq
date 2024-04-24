package model

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/xackery/tinywebeq/library"
)

type Player struct {
	key                   string
	expiration            int64
	ID                    int          `db:"id"`                      //	int(11) unsigned
	AccountID             int          `db:"account_id"`              //	int(11)
	Name                  string       `db:"name"`                    //	varchar(64)
	LastName              string       `db:"last_name"`               //	varchar(64)
	Title                 string       `db:"title"`                   //	varchar(32)
	Suffix                string       `db:"suffix"`                  //	varchar(32)
	ZoneId                int          `db:"zone_id"`                 //	int(11) unsigned
	ZoneInstance          int          `db:"zone_instance"`           //	int(11) unsigned
	Y                     float32      `db:"y"`                       //	float
	X                     float32      `db:"x"`                       //	float
	Z                     float32      `db:"z"`                       //	float
	Heading               float32      `db:"heading"`                 //	float
	Gender                int          `db:"gender"`                  //	tinyint(11) unsigned
	Race                  int          `db:"race"`                    //	smallint(11) unsigned
	Class                 int          `db:"class"`                   //	tinyint(11) unsigned
	Level                 int          `db:"level"`                   //	int(11) unsigned
	Deity                 int          `db:"deity"`                   //	int(11) unsigned
	Birthday              int          `db:"birthday"`                //	int(11) unsigned
	LastLogin             int          `db:"last_login"`              //	int(11) unsigned
	TimePlayed            int          `db:"time_played"`             //	int(11) unsigned
	Level2                int          `db:"level2"`                  //	tinyint(11) unsigned
	Anon                  int          `db:"anon"`                    //	tinyint(11) unsigned
	Gm                    int          `db:"gm"`                      //	tinyint(11) unsigned
	Face                  int          `db:"face"`                    //	int(11) unsigned
	HairColor             int          `db:"hair_color"`              //	tinyint(11) unsigned
	HairStyle             int          `db:"hair_style"`              //	tinyint(11) unsigned
	Beard                 int          `db:"beard"`                   //	tinyint(11) unsigned
	BeardColor            int          `db:"beard_color"`             //	tinyint(11) unsigned
	EyeColor1             int          `db:"eye_color_1"`             //	tinyint(11) unsigned
	EyeColor2             int          `db:"eye_color_2"`             //	tinyint(11) unsigned
	DrakkinHeritage       int          `db:"drakkin_heritage"`        //	int(11) unsigned
	DrakkinTattoo         int          `db:"drakkin_tattoo"`          //	int(11) unsigned
	DrakkinDetails        int          `db:"drakkin_details"`         //	int(11) unsigned
	AbilityTimeSeconds    int          `db:"ability_time_seconds"`    //	tinyint(11) unsigned
	AbilityNumber         int          `db:"ability_number"`          //	tinyint(11) unsigned
	AbilityTimeMinutes    int          `db:"ability_time_minutes"`    //	tinyint(11) unsigned
	AbilityTimeHours      int          `db:"ability_time_hours"`      //	tinyint(11) unsigned
	Exp                   int          `db:"exp"`                     //	int(11) unsigned
	ExpEnabled            int          `db:"exp_enabled"`             //	tinyint(1) unsigned
	AaPointsSpent         int          `db:"aa_points_spent"`         //	int(11) unsigned
	AaExp                 int          `db:"aa_exp"`                  //	int(11) unsigned
	AaPoints              int          `db:"aa_points"`               //	int(11) unsigned
	GroupLeadershipExp    int          `db:"group_leadership_exp"`    //	int(11) unsigned
	RaidLeadershipExp     int          `db:"raid_leadership_exp"`     //	int(11) unsigned
	GroupLeadershipPoints int          `db:"group_leadership_points"` //	int(11) unsigned
	RaidLeadershipPoints  int          `db:"raid_leadership_points"`  //	int(11) unsigned
	Points                int          `db:"points"`                  //	int(11) unsigned
	CurHp                 int          `db:"cur_hp"`                  //	int(11) unsigned
	Mana                  int          `db:"mana"`                    //	int(11) unsigned
	Endurance             int          `db:"endurance"`               //	int(11) unsigned
	Intoxication          int          `db:"intoxication"`            //	int(11) unsigned
	Str                   int          `db:"str"`                     //	int(11) unsigned
	Sta                   int          `db:"sta"`                     //	int(11) unsigned
	Cha                   int          `db:"cha"`                     //	int(11) unsigned
	Dex                   int          `db:"dex"`                     //	int(11) unsigned
	Int                   int          `db:"int"`                     //	int(11) unsigned
	Agi                   int          `db:"agi"`                     //	int(11) unsigned
	Wis                   int          `db:"wis"`                     //	int(11) unsigned
	ZoneChangeCount       int          `db:"zone_change_count"`       //	int(11) unsigned
	Toxicity              int          `db:"toxicity"`                //	int(11) unsigned
	HungerLevel           int          `db:"hunger_level"`            //	int(11) unsigned
	ThirstLevel           int          `db:"thirst_level"`            //	int(11) unsigned
	AbilityUp             int          `db:"ability_up"`              //	int(11) unsigned
	LdonPointsGuk         int          `db:"ldon_points_guk"`         //	int(11) unsigned
	LdonPointsMir         int          `db:"ldon_points_mir"`         //	int(11) unsigned
	LdonPointsMmc         int          `db:"ldon_points_mmc"`         //	int(11) unsigned
	LdonPointsRuj         int          `db:"ldon_points_ruj"`         //	int(11) unsigned
	LdonPointsTak         int          `db:"ldon_points_tak"`         //	int(11) unsigned
	LdonPointsAvailable   int          `db:"ldon_points_available"`   //	int(11) unsigned
	TributeTimeRemaining  int          `db:"tribute_time_remaining"`  //	int(11) unsigned
	CareerTributePoints   int          `db:"career_tribute_points"`   //	int(11) unsigned
	TributePoints         int          `db:"tribute_points"`          //	int(11) unsigned
	TributeActive         int          `db:"tribute_active"`          //	int(11) unsigned
	PvpStatus             int          `db:"pvp_status"`              //	tinyint(11) unsigned
	PvpKills              int          `db:"pvp_kills"`               //	int(11) unsigned
	PvpDeaths             int          `db:"pvp_deaths"`              //	int(11) unsigned
	PvpCurrentPoints      int          `db:"pvp_current_points"`      //	int(11) unsigned
	PvpCareerPoints       int          `db:"pvp_career_points"`       //	int(11) unsigned
	PvpBestKillStreak     int          `db:"pvp_best_kill_streak"`    //	int(11) unsigned
	PvpWorstDeathStreak   int          `db:"pvp_worst_death_streak"`  //	int(11) unsigned
	PvpCurrentKillStreak  int          `db:"pvp_current_kill_streak"` //	int(11) unsigned
	Pvp2                  int          `db:"pvp2"`                    //	int(11) unsigned
	PvpType               int          `db:"pvp_type"`                //	int(11) unsigned
	ShowHelm              int          `db:"show_helm"`               //	int(11) unsigned
	GroupAutoConsent      int          `db:"group_auto_consent"`      //	tinyint(11) unsigned
	RaidAutoConsent       int          `db:"raid_auto_consent"`       //	tinyint(11) unsigned
	GuildAutoConsent      int          `db:"guild_auto_consent"`      //	tinyint(11) unsigned
	LeadershipExpOn       int          `db:"leadership_exp_on"`       //	tinyint(11) unsigned
	RestTimer             int          `db:"RestTimer"`               //	int(11) unsigned
	AirRemaining          int          `db:"air_remaining"`           //	int(11) unsigned
	AutosplitEnabled      int          `db:"autosplit_enabled"`       //	int(11) unsigned
	Lfp                   int          `db:"lfp"`                     //	tinyint(1) unsigned
	Lfg                   int          `db:"lfg"`                     //	tinyint(1) unsigned
	Mailkey               string       `db:"mailkey"`                 //	char(16)
	Xtargets              int          `db:"xtargets"`                //	tinyint(3) unsigned
	Firstlogon            int          `db:"firstlogon"`              //	tinyint(3)
	EAaEffects            int          `db:"e_aa_effects"`            //	int(11) unsigned
	EPercentToAa          int          `db:"e_percent_to_aa"`         //	int(11) unsigned
	EExpendedAaSpent      int          `db:"e_expended_aa_spent"`     //	int(11) unsigned
	AaPointsSpentOld      int          `db:"aa_points_spent_old"`     //	int(11) unsigned
	AaPointsOld           int          `db:"aa_points_old"`           //	int(11) unsigned
	ELastInvsnapshot      int          `db:"e_last_invsnapshot"`      //	int(11) unsigned
	DeletedAt             sql.NullTime `db:"deleted_at"`              //	datetime
}

func (t *Player) Identifier() string {
	return "player"
}

func (t *Player) Key() string {
	return t.key
}

func (t *Player) SetKey(key string) {
	t.key = key
}

func (t *Player) SetExpiration(expiration int64) {
	t.expiration = expiration
}

func (t *Player) Expiration() int64 {
	return t.expiration
}

func (t *Player) ClassStr() string {
	return library.ClassStr(t.Class)
}

func (t *Player) RaceStr() string {
	out := ""

	out = strings.TrimSuffix(out, " ")
	return out
}

func (t *Player) IconUrl() string {
	return "https://www.eqitems.com/item_images/"
}

func (t *Player) Serialize() string {
	return serialize(t)
}

func (t *Player) Deserialize(data string) error {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return fmt.Errorf("base64 decode: %w", err)
	}
	buf := bytes.NewBuffer(decoded)
	d := gob.NewDecoder(buf)

	err = d.Decode(&t)
	if err != nil {
		return fmt.Errorf("gob decode: %w", err)
	}
	return nil
}

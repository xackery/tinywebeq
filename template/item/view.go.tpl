{{ define "data" }}

{{ if .IsItemSearchEnabled }}
<input type="text" name="name" autocomplete="off" placeholder="Search by name">
<div id="searchResults"></div>

<script>
document.addEventListener("DOMContentLoaded", function() {
    let searchTimeout = null;
    let searchResults = document.getElementById("searchResults");
    let searchInput = document.querySelector("input[name='name']");
    searchInput.value = '';
    searchInput.addEventListener("input", function() {
        if (searchTimeout) {
            clearTimeout(searchTimeout);
        }
        searchTimeout = setTimeout(function() {
            if (searchInput.value.length < 3) {
                searchResults.innerHTML = "";
                console.log("Search too short")
                return;
            } 
            fetch("/item/search?name=" + searchInput.value)
                .then(response => response.json())
                .then(data => {
                    searchResults.innerHTML = data.message + "<br>";

                    let items = data.items;
                    if (items === undefined) {
                        return;
                    }
                    if (items.length === 0) {
                        return;
                    }

                    for (let i = 0; i < items.length; i++) {
                        let item = items[i];
                        let link = document.createElement("a");
                        link.href = "/item/view?id=" + item.id;
                        link.innerText = item.name;
                        searchResults.appendChild(link);
                        let suffix = " (Lvl "+item.level+", ID " + item.id + ")";
                        searchResults.appendChild(document.createTextNode(suffix));
                        searchResults.appendChild(document.createElement("br"));
                        console.log("Adding link" + item.id)
                    }
                });
        }, 300);
    });
});
</script>
{{end}}
<br>

<table width="100%"><tr><td valign="top">
<h4 style="margin-top:0">{{ .Item.Name }}</h4></td>
<td><img src="{{ .Item.IconUrl }}item_{{ .Item.Icon }}.Item.png" align="right" valign="top"/></td></tr><tr><td>

<table width="100%" cellpadding="0" cellspacing="0">

    <tr> <td colspan="2" nowrap="1">{{ .Item.TagStr }}</td></tr>
    {{ if (gt .Item.Classes 0) }}<tr><td colspan="2"><b>Class: </b>{{ .Item.ClassesStr }}</td></tr>{{ end }}
    {{ if (gt .Item.Races 0) }}<tr><td colspan="2"><b>Race: </b>{{ .Item.RaceStr }}</td></tr>{{ end }}
    {{ if (gt .Item.Deity 0) }}<tr><td colspan="2" nowrap="1"><b>Deity: </b>{{ .Item.DeityStr }}</td></tr>{{ end }}
    {{ if (gt .Item.Slots 0) }}<tr><td colspan="2"><b>{{ .Item.SlotStr }}</b></td></tr>{{ end }}
    {{ if (eq .Item.Slots 0) }}<tr><td colspan="2"><b>NONE</b></td></tr>{{ end }}
    {{ if (gt .Item.Bagslots 0)}}
        <tr><td width="0%" nowrap="1"><b>Item Type: </b>Container</td></tr>
        <tr><td width="0%" nowrap="1"><b>Number of Slots: </b>{{ .Item.Bagslots }}</td></tr>
        {{ if (gt .Item.Bagtype 0) }}<tr><td width="0%" nowrap="1"><b>Trade Skill Container: </b>{{ .Item.BagtypeStr }}</td></tr>{{ end }}
        {{ if (gt .Item.Bagwr 0) }}<tr><td width="0%"  nowrap="1"><b>Weight Reduction: </b>{{ .Item.Bagwr }}%</td></tr>{{ end }}
        <tr><td width="0%" nowrap="1" colspan="2">This can hold {{ .Item.BagsizeStr }} and smaller items.</td></tr>
    {{ end }}
    </table>
    <br>
    <table>
        <tr valign="top"><td>
        <table style="width: 33%">
            <tr><td><b>Size: </b></td><td style="text-align:right">{{ .Item.SizeStr }}</td></tr>
            {{ if (ne .Item.Weight 0) }}<tr><td><b>Weight:</b></td><td style='text-align:right'>{{ .Item.Weight }}</td></tr>{{ end }}
            {{ if (eq .Item.Slots 0)}}
                <tr><td><b>{{ .Item.TypeStr }}: </b></td><td>Inventory
            {{else}}
                <tr><td><b>{{ .Item.TypeStr }}: </b></td><td><td style='text-align:right'>{{ .Item.ItemTypeStr }}
            {{ end }}
            {{ if (gt .Item.Stackable 0)}} (stackable){{ end }}
            </td></tr>
            {{ if (ne .Item.Reclevel 0) }}<tr><td><b>Rec Level:</b></td><td style='text-align:right'>{{ .Item.Reclevel }}</td></tr>{{ end }}
            {{ if (ne .Item.Reqlevel 0) }}<tr><td><b>Req Level:</b></td><td style='text-align:right'>{{ .Item.Reqlevel }}</td></tr>{{ end }}
        </table>
        </td><td>
        <table style="width: 33%">
            {{ if (ne .Item.Ac 0) }}<tr><td><b>AC: </b></td><td style="text-align:right">{{ .Item.Ac }}</td></tr>{{ end }}
            {{ if (ne .Item.Hp 0) }}<tr><td><b>HP: </b></td><td style="text-align:right">{{ .Item.Hp }}</td></tr>{{ end }}
            {{ if (ne .Item.Mana 0) }}<tr><td><b>Mana: </b></td><td style="text-align:right">{{ .Item.Mana }}</td></tr>{{ end }}
            {{ if (ne .Item.Endur 0) }}<tr><td><b>End: </b></td><td style="text-align:right">{{ .Item.Endur }}</td></tr>{{ end }}
            {{ if (ne .Item.Haste 0) }}<tr><td><b>Haste: </b></td><td style="text-align:right">{{ .Item.Haste }}%</td></tr>{{ end }}
        </table>
        </td><td>
        <table style="width: 33%">
            {{ if (ne .Item.Damage 0) }}<tr><td><b>Base Damage:</b></td><td style='text-align:right'>{{ .Item.Damage }}</td></tr>{{ end }}
            {{ if (ne .Item.Elemdmgamt 0) }}<tr><td><b>{{ .Item.EleDamageTypeStr }} Damage:</b></td><td style='text-align:right'>{{ .Item.Eledmgamt }}</td></tr>{{ end }}
            {{ if (ne .Item.Banedmgrace 0) }}<tr><td><b>Bane Damage ({{ .Item.BaneDamageTypeStr }})</b></td><td>{{ .Item.Banedmgamt }}</td></tr>{{ end }}
            {{ if (ne .Item.Banedmgamt 0) }}<tr><td><b>{{ .Item.BaneDamageTypeStr }} Damage:</b></td><td style='text-align:right'>{{ .Item.Banedmgamt }}</td></tr>{{ end }}
            {{ if (ne .Item.Backstabdmg 0) }}<tr><td><b>Backstab Damage:</b></td><td style='text-align:right'>{{ .Item.Backstabdmg }}</td></tr>{{ end }}
            {{ if (ne .Item.Delay 0) }}<tr><td><b>Delay:</b></td><td style='text-align:right'>{{ .Item.Delay }}</td></tr>{{ end }}
            {{ if (gt .Item.DamageBonus 0)}}<tr><td><b>Damage bonus: </b></td><td>{{ .Item.DamageBonus }}</td></tr>{{ end }}
            {{ if (ne .Item.Range 0) }}<tr><td><b>Range:</b></td><td style='text-align:right'>{{ .Item.Range }}</td></tr>{{ end }}
        </table>
        </td></tr><tr><td colspan="2">&nbsp;</td></td>
        <tr valign="top"><td>
        <table style="width:100%">
            {{ if (ne .Item.Astr 0) }}<tr><td><b>Strength:</b></td><td style='text-align:right'>{{ .Item.Astr }}{{ if .Item.Heroic_str }}+{{ .Item.Heroic_str }}{{ end }}</td></tr>{{ end }}
            {{ if (ne .Item.Asta 0) }}<tr><td><b>Stamina:</b></td><td style='text-align:right'>{{ .Item.Asta }}{{ if .Item.Heroic_sta }}+{{ .Item.Heroic_sta }}{{ end }}</td></tr>{{ end }}
            {{ if (ne .Item.Aint 0) }}<tr><td><b>Intelligence:</b></td><td style='text-align:right'>{{ .Item.Aint }}{{ if .Item.Heroic_int }}+{{ .Item.Heroic_int }}{{ end }}</td></tr>{{ end }}
            {{ if (ne .Item.Awis 0) }}<tr><td><b>Wisdom:</b></td><td style='text-align:right'>{{ .Item.Awis }}{{ if .Item.Heroic_wis }}+{{ .Item.Heroic_wis }}{{ end }}</td></tr>{{ end }}
            {{ if (ne .Item.Aagi 0) }}<tr><td><b>Agility:</b></td><td style='text-align:right'>{{ .Item.Aagi }}{{ if .Item.Heroic_agi }}+{{ .Item.Heroic_agi }}{{ end }}</td></tr>{{ end }}
            {{ if (ne .Item.Adex 0) }}<tr><td><b>Dexterity:</b></td><td style='text-align:right'>{{ .Item.Adex }}{{ if .Item.Heroic_dex }}+{{ .Item.Heroic_dex }}{{ end }}</td></tr>{{ end }}
            {{ if (ne .Item.Acha 0) }}<tr><td><b>Charisma:</b></td><td style='text-align:right'>{{ .Item.Acha }}{{ if .Item.Heroic_cha }}+{{ .Item.Heroic_cha }}{{ end }}</td></tr>{{ end }}
        </table>
        </td><td>
        <table style="width:100%">
            {{ if (ne .Item.Mr 0) }}<tr><td><b>Magic Resist:</b></td><td style='text-align:right'>{{ .Item.Mr }}{{ if .Item.Heroic_mr }}+{{ .Item.Heroic_mr }}{{ end }}</td></tr>{{ end }}
            {{ if (ne .Item.Fr 0) }}<tr><td><b>Fire Resist:</b></td><td style='text-align:right'>{{ .Item.Fr }}{{ if .Item.Heroic_fr }}+{{ .Item.Heroic_fr }}{{ end }}</td></tr>{{ end }}
            {{ if (ne .Item.Cr 0) }}<tr><td><b>Combat Effects:</b></td><td style='text-align:right'>{{ .Item.Cr }}{{ if .Item.Heroic_cr }}+{{ .Item.Heroic_cr }}{{ end }}</td></tr>{{ end }}
            {{ if (ne .Item.Dr 0) }}<tr><td><b>Disease Resist:</b></td><td style='text-align:right'>{{ .Item.Dr }}{{ if .Item.Heroic_dr }}+{{ .Item.Heroic_dr }}{{ end }}</td></tr>{{ end }}
            {{ if (ne .Item.Pr 0) }}<tr><td><b>Poison Resist:</b></td><td style='text-align:right'>{{ .Item.Pr }}{{ if .Item.Heroic_pr }}+{{ .Item.Heroic_pr }}{{ end }}</td></tr>{{ end }}
        </table>
        </td><td>
        <table style="width:100%">
            {{ if (ne .Item.Attack 0) }}<tr><td><b>Attack:</b></td><td style='text-align:right'>{{ .Item.Attack }}</td></tr>{{ end }}
            {{ if (ne .Item.Regen 0) }}<tr><td><b>HP Regen:</b></td><td style='text-align:right'>{{ .Item.Regen }}</td></tr>{{ end }}
            {{ if (ne .Item.Manaregen 0) }}<tr><td><b>Mana Regen:</b></td><td style='text-align:right'>{{ .Item.Manaregen }}</td></tr>{{ end }}
            {{ if (ne .Item.Enduranceregen 0) }}<tr><td><b>Endurance Regen:</b></td><td style='text-align:right'>{{ .Item.Enduranceregen }}</td></tr>{{ end }}
            {{ if (ne .Item.Spellshield 0) }}<tr><td><b>Spell Shielding:</b></td><td style='text-align:right'>{{ .Item.Spellshield }}</td></tr>{{ end }}
            {{ if (ne .Item.Dotshielding 0) }}<tr><td><b>Dot Shielding:</b></td><td style='text-align:right'>{{ .Item.Dotshielding }}</td></tr>{{ end }}
            {{ if (ne .Item.Avoidance 0) }}<tr><td><b>Avoidance:</b></td><td style='text-align:right'>{{ .Item.Avoidance }}</td></tr>{{ end }}
            {{ if (ne .Item.Accuracy 0) }}<tr><td><b>Accuracy:</b></td><td style='text-align:right'>{{ .Item.Accuracy }}</td></tr>{{ end }}
            {{ if (ne .Item.Stunresist 0) }}<tr><td><b>Stun Resist:</b></td><td style='text-align:right'>{{ .Item.Stunresist }}</td></tr>{{ end }}
            {{ if (ne .Item.Strikethrough 0) }}<tr><td><b>Strikethrough:</b></td><td style='text-align:right'>{{ .Item.Strikethrough }}</td></tr>{{ end }}
            {{ if (ne .Item.Damageshield 0) }}<tr><td><b>Damage Shield:</b></td><td style='text-align:right'>{{ .Item.Damageshield }}</td></tr>{{ end }}
            </td></tr>
        </table>
        </td></tr>
    </table><br>
    {{ if (ne .Item.Extradmgamt 0) }}<tr><td><b>{{ .Item.ExtraDamageSkillStr }} Damage: </b>{{ .Item.Extradmgamt }}</td></tr>{{ end }}
    {{ if and (ne .Item.Skillmodtype 0) (ne .Item.Skillmodvalue 0) }}<tr><td colspan="2" nowrap="1"><b>Skill Mod: {{ .Item.SkillModTypeStr }}: </b>{{ .Item.Skillmodvalue }}%</td></tr>{{ end }}
    {{ if (gt .Item.Augslot1type 0)  }}<tr><td width="0%" nowrap="1" colspan="2"><img src="images/icons/blank_slot.gif" style="width:auto;height:10px"> <b>Slot 1: </b>Type {{ .Item.Augslot1type }}</td></tr>{{ end }}
    {{ if (gt .Item.Augslot2type 0)  }}<tr><td width="0%" nowrap="1" colspan="2"><img src="images/icons/blank_slot.gif" style="width:auto;height:10px"> <b>Slot 2: </b>Type {{ .Item.Augslot2type }}</td></tr>{{ end }}
    {{ if (gt .Item.Augslot3type 0)  }}<tr><td width="0%" nowrap="1" colspan="2"><img src="images/icons/blank_slot.gif" style="width:auto;height:10px"> <b>Slot 3: </b>Type {{ .Item.Augslot3type }}</td></tr>{{ end }}
    {{ if (gt .Item.Augslot4type 0)  }}<tr><td width="0%" nowrap="1" colspan="2"><img src="images/icons/blank_slot.gif" style="width:auto;height:10px"> <b>Slot 4: </b>Type {{ .Item.Augslot4type }}</td></tr>{{ end }}
    {{ if (gt .Item.Augslot5type 0)  }}<tr><td width="0%" nowrap="1" colspan="2"><img src="images/icons/blank_slot.gif" style="width:auto;height:10px"> <b>Slot 5: </b>Type {{ .Item.Augslot5type }}</td></tr>{{ end }}
    <td><td>&nbsp;</td><td></tr>
    {{ if and (gt .Item.Proceffect 0) (ne .Item.Proceffect 65535) }}
        <tr><td colspan="2" nowrap="1">
        <b>Combat Effects: </b>{{ .Library.SpellName .Item.Proceffect }}
        {{ if (gt .Item.Proclevel2 0)}}<br><b>Level for effect: </b>{{ .Item.Proclevel2 }}{{ end }}
        <br><b>Effect chance modifier: {{ .Util.Add 100 .Item.Procrate }}%</b>
        </td></tr>
    {{ end }}
    </table>

    {{ if .ItemRecipe }}
        <br><b>Used in tradeskills:</b><br>
        {{ range .ItemRecipe.ComponentEntries }}
            {{ .RecipeName }}{{ if ne .ComponentCount 1}} ({{ .ComponentCount }}){{ end }}<br>
        {{ end }}
        <br><b>Result of tradeskills:</b><br>
        {{ range .ItemRecipe.SuccessEntries }}
            {{ .RecipeName }}{{ if ne .SuccessCount 1}} ({{ .SuccessCount }}){{ end }}<br>
        {{ end }}
    {{ end }}


    {{ if .ItemQuest }}
        <r><td><b>Obtained from Quest:</b><br>
        {{ range .ItemQuest.Entries }}
            {{ .NpcCleanName }} in {{ .ZoneLongName }}<br>
        {{ end}}
    {{ end }}

{{ end }}
{{ .Item.ID }}:
  - |
    <table width="500px">
        <tr><td><i class="item-{{ .Item.Icon }}"></i><h4 style="margin-top:0;display:inline">{{ .Item.Name }}</h4></td></tr>
        <table width="500px" cellpadding="0" cellspacing="0">
            <tr> <td colspan="2" nowrap="1">{{ .Item.TagStr }}</td></tr>
            {{- if (gt .Item.Classes 0) }}<tr><td colspan="2"><b>Class: </b>{{ .Item.ClassesStr }}</td></tr>{{- end }}
            {{- if (gt .Item.Races 0) }}<tr><td colspan="2"><b>Race: </b>{{ .Item.RaceStr }}</td></tr>{{- end }}
            {{- if (gt .Item.Deity 0) }}<tr><td colspan="2" nowrap="1"><b>Deity: </b>{{ .Item.DeityStr }}</td></tr>{{- end }}
            {{- if (gt .Item.Slots 0) }}<tr><td colspan="2"><b>{{ .Item.SlotStr }}</b></td></tr>{{- end }}
            {{- if (eq .Item.Slots 0) }}<tr><td colspan="2"><b>NONE</b></td></tr>{{- end }}
            {{- if (gt .Item.Bagslots 0)}}
                <tr><td width="0%" nowrap="1"><b>Item Type: </b>Container</td></tr>
                <tr><td width="0%" nowrap="1"><b>Number of Slots: </b>{{ .Item.Bagslots }}</td></tr>
                {{- if (gt .Item.Bagtype 0) }}<tr><td width="0%" nowrap="1"><b>Trade Skill Container: </b>{{ .Item.BagtypeStr }}</td></tr>{{- end }}
                {{- if (gt .Item.Bagwr 0) }}<tr><td width="0%"  nowrap="1"><b>Weight Reduction: </b>{{ .Item.Bagwr }}%</td></tr>{{- end }}
                <tr><td width="0%" nowrap="1" colspan="2">This can hold {{ .Item.BagsizeStr }} and smaller items.</td></tr>
            {{- end }}
        </table>
        <br>
        <table>
            <tr valign="top"><td>
            <table style="width: 99%">
                <tr><td><b>Size: </b></td><td style="text-align:right">{{ .Item.SizeStr }}</td></tr>
                {{- if (ne .Item.Weight 0) }}<tr><td><b>Weight:</b></td><td style='text-align:right'>{{ .Item.Weight }}</td></tr>{{- end }}
                {{- if (eq .Item.Slots 0)}}
                    <tr><td><b>{{ .Item.TypeStr }}: </b></td><td>Inventory
                {{else}}
                    <tr><td><b>{{ .Item.TypeStr }}: </b></td><td><td style='text-align:right'>{{ .Item.ItemTypeStr }}
                {{- end }}
                {{- if (gt .Item.Stackable 0)}} (stackable){{- end }}
                </td></tr>
                {{- if (ne .Item.Reclevel 0) }}<tr><td><b>Rec Level:</b></td><td style='text-align:right'>{{ .Item.Reclevel }}</td></tr>{{- end }}
                {{- if (ne .Item.Reqlevel 0) }}<tr><td><b>Req Level:</b></td><td style='text-align:right'>{{ .Item.Reqlevel }}</td></tr>{{- end }}
            </table>
        </td><td>
            <table style="width: 99%">
                {{- if (ne .Item.Ac 0) }}<tr><td><b>AC: </b></td><td style="text-align:right">{{ .Item.Ac }}</td></tr>{{- end }}
                {{- if (ne .Item.Hp 0) }}<tr><td><b>HP: </b></td><td style="text-align:right">{{ .Item.Hp }}</td></tr>{{- end }}
                {{- if (ne .Item.Mana 0) }}<tr><td><b>Mana: </b></td><td style="text-align:right">{{ .Item.Mana }}</td></tr>{{- end }}
                {{- if (ne .Item.Endur 0) }}<tr><td><b>End: </b></td><td style="text-align:right">{{ .Item.Endur }}</td></tr>{{- end }}
                {{- if (ne .Item.Haste 0) }}<tr><td><b>Haste: </b></td><td style="text-align:right">{{ .Item.Haste }}%</td></tr>{{- end }}
            </table>
        </td><td>
            <table style="width: 99%">
                {{- if (ne .Item.Damage 0) }}<tr><td><b>Base Damage:</b></td><td style='text-align:right'>{{ .Item.Damage }}</td></tr>{{- end }}
                {{- if (ne .Item.Elemdmgamt 0) }}<tr><td><b>{{ .Item.EleDamageTypeStr }} Damage:</b></td><td style='text-align:right'>{{ .Item.Elemdmgamt }}</td></tr>{{- end }}
                {{- if (ne .Item.Banedmgrace 0) }}<tr><td><b>Bane Damage ({{ .Item.BaneDamageTypeStr }})</b></td><td>{{ .Item.Banedmgamt }}</td></tr>{{- end }}
                {{- if (ne .Item.Banedmgamt 0) }}<tr><td><b>{{ .Item.BaneDamageTypeStr }} Damage:</b></td><td style='text-align:right'>{{ .Item.Banedmgamt }}</td></tr>{{- end }}
                {{- if (ne .Item.Backstabdmg 0) }}<tr><td><b>Backstab Damage:</b></td><td style='text-align:right'>{{ .Item.Backstabdmg }}</td></tr>{{- end }}
                {{- if (ne .Item.Delay 0) }}<tr><td><b>Delay:</b></td><td style='text-align:right'>{{ .Item.Delay }}</td></tr>{{- end }}
                {{- if (gt .Item.DamageBonus 0)}}<tr><td><b>Damage bonus: </b></td><td>{{ .Item.DamageBonus }}</td></tr>{{- end }}
                {{- if (ne .Item.Range 0) }}<tr><td><b>Range:</b></td><td style='text-align:right'>{{ .Item.Range }}</td></tr>{{- end }}
            </table>
        </td></tr><tr><td colspan="2">&nbsp;</td></td>
        <tr valign="top"><td>
            <table style="width:100%">
                {{- if (ne .Item.Astr 0) }}<tr><td><b>Strength:</b></td><td style='text-align:right'>{{ .Item.Astr }}{{- if .Item.HeroicStr }}+{{ .Item.HeroicStr }}{{- end }}</td></tr>{{- end }}
                {{- if (ne .Item.Asta 0) }}<tr><td><b>Stamina:</b></td><td style='text-align:right'>{{ .Item.Asta }}{{- if .Item.HeroicSta }}+{{ .Item.HeroicSta }}{{- end }}</td></tr>{{- end }}
                {{- if (ne .Item.Aint 0) }}<tr><td><b>Intelligence:</b></td><td style='text-align:right'>{{ .Item.Aint }}{{- if .Item.HeroicInt }}+{{ .Item.HeroicInt }}{{- end }}</td></tr>{{- end }}
                {{- if (ne .Item.Awis 0) }}<tr><td><b>Wisdom:</b></td><td style='text-align:right'>{{ .Item.Awis }}{{- if .Item.HeroicWis }}+{{ .Item.HeroicWis }}{{- end }}</td></tr>{{- end }}
                {{- if (ne .Item.Aagi 0) }}<tr><td><b>Agility:</b></td><td style='text-align:right'>{{ .Item.Aagi }}{{- if .Item.HeroicAgi }}+{{ .Item.HeroicAgi }}{{- end }}</td></tr>{{- end }}
                {{- if (ne .Item.Adex 0) }}<tr><td><b>Dexterity:</b></td><td style='text-align:right'>{{ .Item.Adex }}{{- if .Item.HeroicDex }}+{{ .Item.HeroicDex }}{{- end }}</td></tr>{{- end }}
                {{- if (ne .Item.Acha 0) }}<tr><td><b>Charisma:</b></td><td style='text-align:right'>{{ .Item.Acha }}{{- if .Item.HeroicCha }}+{{ .Item.HeroicCha }}{{- end }}</td></tr>{{- end }}
            </table>
        </td><td>
            <table style="width:100%">
                {{- if (ne .Item.Mr 0) }}<tr><td><b>Magic Resist:</b></td><td style='text-align:right'>{{ .Item.Mr }}{{- if .Item.HeroicMr }}+{{ .Item.HeroicMr }}{{- end }}</td></tr>{{- end }}
                {{- if (ne .Item.Fr 0) }}<tr><td><b>Fire Resist:</b></td><td style='text-align:right'>{{ .Item.Fr }}{{- if .Item.HeroicFr }}+{{ .Item.HeroicFr }}{{- end }}</td></tr>{{- end }}
                {{- if (ne .Item.Cr 0) }}<tr><td><b>Combat Effects:</b></td><td style='text-align:right'>{{ .Item.Cr }}{{- if .Item.HeroicCr }}+{{ .Item.HeroicCr }}{{- end }}</td></tr>{{- end }}
                {{- if (ne .Item.Dr 0) }}<tr><td><b>Disease Resist:</b></td><td style='text-align:right'>{{ .Item.Dr }}{{- if .Item.HeroicDr }}+{{ .Item.HeroicDr }}{{- end }}</td></tr>{{- end }}
                {{- if (ne .Item.Pr 0) }}<tr><td><b>Poison Resist:</b></td><td style='text-align:right'>{{ .Item.Pr }}{{- if .Item.HeroicPr }}+{{ .Item.HeroicPr }}{{- end }}</td></tr>{{- end }}
            </table>
        </td><td>
            <table style="width:100%">
                {{- if (ne .Item.Attack 0) }}<tr><td><b>Attack:</b></td><td style='text-align:right'>{{ .Item.Attack }}</td></tr>{{- end }}
                {{- if (ne .Item.Regen 0) }}<tr><td><b>HP Regen:</b></td><td style='text-align:right'>{{ .Item.Regen }}</td></tr>{{- end }}
                {{- if (ne .Item.Manaregen 0) }}<tr><td><b>Mana Regen:</b></td><td style='text-align:right'>{{ .Item.Manaregen }}</td></tr>{{- end }}
                {{- if (ne .Item.Enduranceregen 0) }}<tr><td><b>Endurance Regen:</b></td><td style='text-align:right'>{{ .Item.Enduranceregen }}</td></tr>{{- end }}
                {{- if (ne .Item.Spellshield 0) }}<tr><td><b>Spell Shielding:</b></td><td style='text-align:right'>{{ .Item.Spellshield }}</td></tr>{{- end }}
                {{- if (ne .Item.Dotshielding 0) }}<tr><td><b>Dot Shielding:</b></td><td style='text-align:right'>{{ .Item.Dotshielding }}</td></tr>{{- end }}
                {{- if (ne .Item.Avoidance 0) }}<tr><td><b>Avoidance:</b></td><td style='text-align:right'>{{ .Item.Avoidance }}</td></tr>{{- end }}
                {{- if (ne .Item.Accuracy 0) }}<tr><td><b>Accuracy:</b></td><td style='text-align:right'>{{ .Item.Accuracy }}</td></tr>{{- end }}
                {{- if (ne .Item.Stunresist 0) }}<tr><td><b>Stun Resist:</b></td><td style='text-align:right'>{{ .Item.Stunresist }}</td></tr>{{- end }}
                {{- if (ne .Item.Strikethrough 0) }}<tr><td><b>Strikethrough:</b></td><td style='text-align:right'>{{ .Item.Strikethrough }}</td></tr>{{- end }}
                {{- if (ne .Item.Damageshield 0) }}<tr><td><b>Damage Shield:</b></td><td style='text-align:right'>{{ .Item.Damageshield }}</td></tr>{{- end }}
                </td></tr>
            </table>
        </td></tr>
        </table>
    <br>
    {{- if (ne .Item.Extradmgamt 0) }}<tr><td><b>{{ .Item.ExtraDamageSkillStr }} Damage: </b>{{ .Item.Extradmgamt }}</td></tr>{{- end }}
    {{- if and (ne .Item.Skillmodtype 0) (ne .Item.Skillmodvalue 0) }}<tr><td colspan="2" nowrap="1"><b>Skill Mod: {{ .Item.SkillModTypeStr }}: </b>{{ .Item.Skillmodvalue }}%</td></tr>{{- end }}
    {{- if (gt .Item.Augslot1type 0)  }}<tr><td width="0%" nowrap="1" colspan="2"><img src="" style="width:auto;height:10px"> <b>Slot 1: </b>Type {{ .Item.Augslot1type }}</td></tr>{{- end }}
    {{- if (gt .Item.Augslot2type 0)  }}<tr><td width="0%" nowrap="1" colspan="2"><img src="" style="width:auto;height:10px"> <b>Slot 2: </b>Type {{ .Item.Augslot2type }}</td></tr>{{- end }}
    {{- if (gt .Item.Augslot3type 0)  }}<tr><td width="0%" nowrap="1" colspan="2"><img src="" style="width:auto;height:10px"> <b>Slot 3: </b>Type {{ .Item.Augslot3type }}</td></tr>{{- end }}
    {{- if (gt .Item.Augslot4type 0)  }}<tr><td width="0%" nowrap="1" colspan="2"><img src="" style="width:auto;height:10px"> <b>Slot 4: </b>Type {{ .Item.Augslot4type }}</td></tr>{{- end }}
    {{- if (gt .Item.Augslot5type 0)  }}<tr><td width="0%" nowrap="1" colspan="2"><img src="" style="width:auto;height:10px"> <b>Slot 5: </b>Type {{ .Item.Augslot5type }}</td></tr>{{- end }}
    <td><td>&nbsp;</td><td></tr>
    {{- if and (gt .Item.Proceffect 0) (ne .Item.Proceffect 65535) }}
        <tr><td colspan="2" nowrap="1">
        <b>Combat Effects: </b>{{ .Store.SpellName .Item.Proceffect }}
        {{- if (gt .Item.Proclevel2 0)}}<br><b>Level for effect: </b>{{ .Item.Proclevel2 }}{{- end }}
        <br><b>Effect chance modifier: {{ .Item.ProcRateTotal }}%</b>
        </td></tr>
    {{- end }}
    </table>

    {{- if .ItemRecipe }}
        {{- if gt (len .ItemRecipe.ComponentEntries) 1 }}
        <br><b>Used in tradeskills:</b><br>
        {{ range .ItemRecipe.ComponentEntries }}
            {{ .RecipeName }}{{- if ne .ComponentCount 1}} ({{ .ComponentCount }}){{- end }}<br>
        {{- end }}
        {{- end }}
        {{- if gt (len .ItemRecipe.SuccessEntries) 1 }}
        <br><b>Result of tradeskills:</b><br>
        {{ range .ItemRecipe.SuccessEntries }}
            {{ .RecipeName }}{{- if ne .SuccessCount 1}} ({{ .SuccessCount }}){{- end }}<br>
        {{- end }}
        {{- end }}
    {{- end }}


    {{- if .ItemQuest }}
        {{- if len .ItemQuest.RewardEntries}}
            <r><td><b>Result of Quests:</b><br>
            {{ range .ItemQuest.RewardEntries }}
                <a href="/quest/view?id={{ .QuestID }}">Quest {{ .QuestName }}</a> - {{ .NpcCleanName }} in {{ .Zone.LongName }}<br>
            {{- end}}
        {{- end }}

        {{- if len .ItemQuest.ComponentEntries}}
            <r><td><b>Used in Quests:</b><br>
            {{ range .ItemQuest.ComponentEntries }}
                <a href="/quest/view?id={{ .QuestID }}">Quest {{ .QuestName }}</a> - {{ .NpcCleanName }} in {{ .Zone.LongName }}<br>
            {{- end }}
        {{- end}}
    {{- end }}

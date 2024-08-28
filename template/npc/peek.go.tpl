{{ range .Entries }}
    {{ $id := .Npc.ID }}
    [[{{ $id }}]]
    name = "{{ .Npc.Name }}"
    cleanname = "{{ .Npc.CleanName }}"
    level = {{ .Npc.Level }}
    race = "{{ .Npc.RaceStr }}"
    class = "{{ .Npc.ClassStr }}"
    hp = {{ .Npc.Hp }}
    mindmg = {{ .Npc.Mindmg }}
    maxdmg = {{ .Npc.Maxdmg }}
    attackspeed = {{ .Npc.AttackSpeed }}
    specialattacks = "{{ .Npc.NpcSpecialAttacksStr }}"

    {{ if .NpcQuest }}[{{ $id}}.quests]
    {{ range .NpcQuest.Entries }}
        [[{{ $id}}.quests.entries]]
        id = {{ .QuestID }}
        name = "{{ .QuestName }}"
        zone = "{{ .ZoneLongName }}"
    {{ end }}{{ end }}

    {{ if .NpcSpawn }}[{{ $id}}.spawns]
    {{ range .NpcSpawn.Entries }}
        [[{{ $id}}.spawns.entries]]
        spawngroup = "{{ .Spawngroup }}"
        spawngroupid = {{ .Spawngroupid }}
        longname = "{{ .LongName }}"
        shortname = "{{ .ShortName }}"
        x = {{ .X }}
        y = {{ .Y }}
        z = {{ .Z }}
        respawntime = {{ .Respawntime }}
    {{ end }}{{ end }}

    {{ if .NpcFaction }}[{{ $id}}.factions]
    {{ range .NpcFaction.Entries }}
        [[{{ $id}}.factions.entries]]
        id = {{ .ID }}
        name = "{{ .Name }}"
        value = {{ .Value }}
    {{ end }}{{ end }}

    {{ if .NpcLoot }}[{{ $id}}.loot]
    {{ range .NpcLoot.Entries }}
        [[{{ $id}}.loot.entries]]
        id = {{ .ID }}
        name = "{{ .Name }}"
        itemtype = "{{ .ItemTypeStr }}"
        chance = {{ .Chance }}
        chanceglobal = {{ .ChanceGlobal }}
    {{ end }}{{ end }}

    {{ if .NpcMerchant }}[{{ $id}}.merchants]
    {{ range .NpcMerchant.Entries }}
        [[{{ $id}}.merchants.entries]]
        name = "{{ .Name }}"
        price = {{ .Price }}
        ldonprice = {{ .Ldonprice }}
    {{ end }}{{ end }}

    {{ if .NpcSpell }}[{{ $id}}.spells]
    {{ range .NpcSpell.Entries }}
        [[{{ $id}}.spells.entries]]
        name = "{{ .Name }}"
        spell = "{{ .Spell.Name }}"
        spellinfo = [
            {{ range .SpellInfo .Npc.Level }}
            "{{ . }}",
            {{ end }}
        ]
    {{ end }}{{ end }}
{{ end }}
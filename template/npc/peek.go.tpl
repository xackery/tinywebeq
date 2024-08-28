{{ .Npc.ID }}:
- "Level: {{ .Npc.Level }}"
- "Race: {{ .Npc.RaceStr }}"
- "Class: {{ .Npc.ClassStr }}"
- "Health points: {{ .Npc.Hp }}"
- "Damage: {{ .Npc.Mindmg }} to {{ .Npc.Maxdmg }}"
- "Attack speed: {{ .Npc.AttackSpeed }}%"
- "{{ if .Npc.NpcSpecialAttacksStr }}Special attacks: {{ .Npc.NpcSpecialAttacksStr }}{{ end }}"
- "{{ if .NpcQuest }} {{ range .NpcQuest.Entries }}{{ .QuestName }} in {{ .ZoneLongName }}{{ end }}{{ end }}"
{{ if .NpcSpawn }} {{ range .NpcSpawn.Entries }}
- "{{ .Spawngroup }} ({{ .Spawngroupid }}) {{ .LongName }} ({{ .ShortName.Value}}) {{ .X }}, {{ .Y }}, {{ .Z }} {{ .Respawntime }}"
{{ end }}{{ end }}
- "{{ if .NpcFaction }} {{ range .NpcFaction.Entries }}{{ .Name }} {{ .Value }}{{ end }}{{ end}}"
- "{{ if .NpcLoot }}{{ range .NpcLoot.Entries }}{{ .Name }} ({{ .ItemTypeStr }}) - "{{ .Chance }}% ({{ .ChanceGlobal }}% Global){{ end }}{{ end }}"
{{ if .NpcMerchant }} {{ range .NpcMerchant.Entries }}
- "{{ .Name }}"
- "Price: {{ .Price }}"
- "{{ if .Ldonprice }}LDoN Price: {{ .Ldonprice }}{{ end }}"
{{ end }}{{ end}}

{{ $level := .Npc.Level }}
{{ if .NpcSpell }} {{ range .NpcSpell.Entries }}
- "{{ .Name }} - "{{ .Spell.Name }}"
    {{ range .SpellInfo $level }}
- "{{ . }}"
    {{ end }}
{{ end }}{{ end}}
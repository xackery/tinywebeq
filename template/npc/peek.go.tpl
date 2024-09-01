<p>Name: {{ .Npc.Name }}</p>
<p>Level: {{ .Npc.Level }}</p>
<p>Race: {{ .Npc.RaceStr }}</p>
<p>Class: {{ .Npc.ClassStr }}</p>
<p>Health points: {{ .Npc.Hp }}</p>
<p>Damage: {{ .Npc.Mindmg }} to {{ .Npc.Maxdmg }}</p>
<p>Attack speed: {{ .Npc.AttackSpeed }}%</p>
{{ if .Npc.NpcSpecialAttacksStr }}<p>Special attacks: {{ .Npc.NpcSpecialAttacksStr }}</p>{{ end }}

{{ if .NpcQuest }} {{ range .NpcQuest.Entries }}
    <p><a href="/quest/view?id={{ .QuestID }}">{{ .QuestName }}</a> in {{ .ZoneID }}</p>
{{ end }}{{ end }}


{{ if .NpcSpawn }} {{ range .NpcSpawn.Entries }}
    <p>{{ .Spawngroup }} ({{ .Spawngroupid }}) {{ .LongName }} ({{ .ShortName}}) {{ .X }}, {{ .Y }}, {{ .Z }} {{ .Respawntime }} </p>
{{ end }}{{ end }}

{{ if .NpcFaction }} {{ range .NpcFaction.Entries }}
    <p><a href="/faction/view?id={{ .ID }}">{{ .Name }}</a> {{ .Value }}</p>
{{ end }}{{ end}}

{{ if .NpcMerchant }} {{ range .NpcMerchant.Entries }}
    <p>{{ .Name }}</p>
    <p>Price: {{ .Price }}</p>
    {{ if .Ldonprice }}<p>LDoN Price: {{ .Ldonprice }}</p>{{ end }}
{{ end }}{{ end}}

{{ $level := .Npc.Level }}
{{ if .NpcSpell }} {{ range .NpcSpell.Entries }}
    <p>{{ .Name }} - {{ .Spell.Name }} </p>
    {{ range .SpellInfo $level }}
        <p>{{ . }}</p>
    {{ end }}
{{ end }}{{ end}}

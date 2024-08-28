{{ define "data" }}
<h1>Npc</h1>
{{ if .IsNpcSearchEnabled }}
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
            fetch("/npc/search?name=" + searchInput.value)
                .then(response => response.json())
                .then(data => {
                    searchResults.innerHTML = data.message + "<br>";

                    let npcs = data.npcs;
                    if (npcs === undefined) {
                        return;
                    }
                    if (npcs.length === 0) {
                        return;
                    }

                    for (let i = 0; i < npcs.length; i++) {
                        let npc = npcs[i];
                        let link = document.createElement("a");
                        link.href = "/npc/view?id=" + npc.id;
                        link.innerText = npc.name;
                        searchResults.appendChild(link);
                        let suffix = " (Lvl "+npc.level+", ID " + npc.id + ")";
                        searchResults.appendChild(document.createTextNode(suffix));
                        searchResults.appendChild(document.createElement("br"));
                        console.log("Adding link" + npc.id)
                    }
                });
        }, 300);
    });
});
</script>
{{end}}
<br>
<sl-button disabled>Click me</sl-button>
<p>Name: {{ .Npc.Name }}</p>
<p>Level: {{ .Npc.Level }}</p>
<p>Race: {{ .Npc.RaceStr }}</p>
<p>Class: {{ .Npc.ClassStr }}</p>
<p>Health points: {{ .Npc.Hp }}</p>
<p>Damage: {{ .Npc.Mindmg }} to {{ .Npc.Maxdmg }}</p>
<p>Attack speed: {{ .Npc.AttackSpeed }}%</p>
{{ if .Npc.NpcSpecialAttacksStr }}<p>Special attacks: {{ .Npc.NpcSpecialAttacksStr }}</p>{{ end }}

{{ if .NpcQuest }} {{ range .NpcQuest.Entries }}
    <p><a href="/quest/view?id={{ .QuestID }}">{{ .QuestName }}</a> in {{ .ZoneLongName }}</p>
{{ end }}{{ end }}


{{ if .NpcSpawn }} {{ range .NpcSpawn.Entries }}
    <p>{{ .Spawngroup }} ({{ .Spawngroupid }}) {{ .LongName }} ({{ .ShortName}}) {{ .X }}, {{ .Y }}, {{ .Z }} {{ .Respawntime }} </p>
{{ end }}{{ end }}

{{ if .NpcFaction }} {{ range .NpcFaction.Entries }}
    <p><a href="/faction/view?id={{ .ID }}">{{ .Name }}</a> {{ .Value }}</p>
{{ end }}{{ end}}

{{ if .NpcLoot }}{{ range .NpcLoot.Entries }}
    <p><a href="/item/view?id={{ .ID }}">{{ .Name }}</a> ({{ .ItemTypeStr }}) - {{ .Chance }}% ({{ .ChanceGlobal }}% Global)</p>
{{ end }}{{ end }}

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

{{ end }}
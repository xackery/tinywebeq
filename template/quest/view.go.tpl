{{ define "data" }}
<h1>Quest</h1>
{{ if .IsQuestSearchEnabled }}
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
            fetch("/quest/search?name=" + searchInput.value)
                .then(response => response.json())
                .then(data => {
                    searchResults.innerHTML = data.message + "<br>";

                    let quests = data.quests;
                    if (quests === undefined) {
                        return;
                    }
                    if (quests.length === 0) {
                        return;
                    }

                    for (let i = 0; i < quests.length; i++) {
                        let quest = quests[i];
                        let link = document.createElement("a");
                        link.href = "/quest/view?id=" + quest.id;
                        link.innerText = quest.name;
                        searchResults.appendChild(link);
                        let suffix = " (Lvl "+quest.level+", ID " + quest.id + ")";
                        searchResults.appendChild(document.createTextNode(suffix));
                        searchResults.appendChild(document.createElement("br"));
                        console.log("Adding link" + quest.id)
                    }
                });
        }, 300);
    });
});
</script>
{{end}}
<br>
<p>Name: {{ .Quest.Name }}</p>
<p>Level: {{ .Quest.Level }}</p>
<p>Expansion: {{ .Quest.Expansion }}</p>
{{ range .Quest.Entries }}
<a href="/npc/view?id={{ .NpcID }}">{{ .NpcName }}</a> {{ .UseCase }} <a href="/items/view?id={{ .ItemID }}">{{ .ItemName }}</a><br>
{{ end }}

{{ end }}
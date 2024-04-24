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
<p>Name: {{ .Npc.Name }}</p>
{{ range $line := .NpcInfo }}
    <p>{{ $line }}</p>
{{ end }}
{{ end }}
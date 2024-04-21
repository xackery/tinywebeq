{{ define "data" }}
<h1>Spell</h1>
{{ if .IsSpellSearchEnabled }}
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
            fetch("/spell/search?name=" + searchInput.value)
                .then(response => response.json())
                .then(data => {
                    searchResults.innerHTML = data.message + "<br>";

                    let spells = data.spells;
                    if (spells === undefined) {
                        return;
                    }
                    if (spells.length === 0) {
                        return;
                    }

                    for (let i = 0; i < spells.length; i++) {
                        let spell = spells[i];
                        let link = document.createElement("a");
                        link.href = "/spell/view?id=" + spell.id;
                        link.innerText = spell.name;
                        searchResults.appendChild(link);
                        let suffix = " (Lvl "+spell.level+", ID " + spell.id + ")";
                        searchResults.appendChild(document.createTextNode(suffix));
                        searchResults.appendChild(document.createElement("br"));
                        console.log("Adding link" + spell.id)
                    }
                });
        }, 300);
    });
});
</script>
{{end}}
<br>
<p>Name: {{ .Spell.Name }}</p>
{{ range $line := .SpellInfo }}
    <p>{{ $line }}</p>
{{ end }}
{{ end }}
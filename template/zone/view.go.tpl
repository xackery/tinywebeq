{{ define "data" }}
<h1>Zone</h1>
{{ if .IsZoneSearchEnabled }}
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
            fetch("/zone/search?name=" + searchInput.value)
                .then(response => response.json())
                .then(data => {
                    searchResults.innerHTML = data.message + "<br>";

                    let zones = data.zones;
                    if (zones === undefined) {
                        return;
                    }
                    if (zones.length === 0) {
                        return;
                    }

                    for (let i = 0; i < zones.length; i++) {
                        let zone = zones[i];
                        let link = document.createElement("a");
                        link.href = "/zone/view?id=" + zone.ZoneIDNumber;
                        link.innerText = zone.LongName + " (" + zone.ShortName + ")";
                        searchResults.appendChild(link);
                        let suffix = " (Expansion "+zone.Expansion+", ID " + zone.ZoneIDNumber + ")";
                        searchResults.appendChild(document.createTextNode(suffix));
                        searchResults.appendChild(document.createElement("br"));
                        console.log("Adding link" + zone.ID)
                    }
                });
        }, 300);
    });
});
</script>
{{end}}
<br>
<p>ID: {{ .Zone.Zoneidnumber }}</p>
<p>Name: {{ .Zone.LongName }} ({{ .Zone.ShortName }})</p>
<p>Expansion: {{ .Zone.Expansion }}</p>

{{ end }}
{{ define "data" }}
<h1>Spell</h1>
<p>Name: {{ .Spell.Name }}</p>
{{ range $line := .SpellInfo }}
    <p>{{ $line }}</p>
{{ end }}
{{ end }}
{{ define "data" }}
<h1>Item</h1>
<p>ID: {{ .Item.ID }}</p>  
<p>Name: {{ .Item.Name }}</p>
<p>Ac: {{ .Item.Ac }}</p>
<p>Slots: {{ .Item.Slots }}</p>      
{{ end }}
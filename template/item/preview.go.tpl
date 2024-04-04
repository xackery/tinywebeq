{{ define "preview" }}
<table width="100%"><tr><td valign="top">
{{/* if (eq .Item.Name "test") */}}
<h4 style="margin-top:0">{{ .Item.Name }}</h4></td>
<td><img src="{{ .Item.IconUrl }}item_{{ .Item.Icon }}.Item.png" align="right" valign="top"/></td></tr><tr><td>
{{/* end */}}
    <table width="100%" cellpadding="0" cellspacing="0">

    <tr> <td colspan="2" nowrap="1">{{ .Item.TagStr }}</td></tr>
    {{ if (gt .Item.Classes 0) }}<tr><td colspan="2"><b>Class: </b>{{ .Item.ClassStr }}</td></tr>{{ end }}
    {{ if (gt .Item.Races 0) }}<tr><td colspan="2"><b>Race: </b>{{ .Item.RaceStr }}</td></tr>{{ end }}
    {{ if (gt .Item.Deity 0) }}<tr><td colspan="2" nowrap="1"><b>Deity: </b>{{ .Item.DeityStr }}</td></tr>{{ end }}
    {{ if (gt .Item.Slots 0) }}<tr><td colspan="2"><b>{{ .Item.SlotStr }}</b></td></tr>{{ end }}
    {{ if (eq .Item.Slots 0) }}<tr><td colspan="2"><b>NONE</b></td></tr>{{ end }}
    {{ if (gt .Item.Bagslots 0)}}
    <tr><td width="0%" nowrap="1"><b>Item Type: </b>Container</td></tr>
    <tr><td width="0%" nowrap="1"><b>Number of Slots: </b>{{ .Item.Bagslots }}</td></tr>
    {{ if (gt .Item.Bagtype 0) }}<tr><td width="0%" nowrap="1"><b>Trade Skill Container: </b>{{ .Item.BagtypeStr }}</td></tr>{{end}}
    {{ if (gt .Item.Bagwr 0) }}<tr><td width="0%"  nowrap="1"><b>Weight Reduction: </b>{{ .Item.Bagwr }}%</td></tr>{{end}}
    <tr><td width="0%" nowrap="1" colspan="2">This can hold {{ .Item.BagsizeStr }} and smaller items.Item.</td></tr>
    {{ end }}
    </table>
    <br>
    <table>
        <tr valign="top"><td>
        <table style="width: 125px;">
            <tr><td><b>Size: </b></td><td style="text-align:right">{{ .Item.SizeStr }}</td></tr>

{{end}}
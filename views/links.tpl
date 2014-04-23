{{ define "links_page" }}
  <html>
    <head>
      {{ template "assets" }}
    </head>
    <body>
      <table>
        {{ template "get_serial_form" }}
        <br>
        {{ $prefix := .Prefix }}
        <b> {{ .Title }}</b>
        {{ range .Episodes }}
          <tr>
            {{ $episode_title := .Title }}
            <td> {{ .Title }} </td>
            {{ range $quality, $link := .VideoLinks }}
              <td>
                <a href='{{ $link }}'>{{ $quality }}</a>
                <a class='down_link' data-episode='{{ $episode_title }}' data-link='{{ $link }}' data-prefix='{{ $prefix }}' href='#'>↓</a>
              </td>
            {{ end }}
          </tr>
        {{ end }}
        <br>
      </table>
    </body>
  </html>
{{ end }}

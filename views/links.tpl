{{ define "links_page" }}
  <html>
    <head>
      <script src='/assets/jquery.js' type='text/javascript'></script>
      <script src='/assets/application.js' type='text/javascript'></script>
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
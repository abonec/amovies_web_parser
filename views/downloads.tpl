{{ define "downloads_page" }}
  <html>
  <head></head>
    <body>
      {{ template "get_serial_form" }}
      <table>
        <tr>
          <td>File</td>
          <td>Link</td>
          <td>Progress</td>
          <td>Added</td>
        </tr>
        {{ range $download, $b := .Downloading }}
          <tr>
            <td>{{ $download.Filename }}</td>
            <td>{{ $download.Link }}</td>
            <td>{{ $download.Progress }}</td>
            <td>{{ $download.Added }}</td>
          </tr>
        {{ end }}
      </table>
      <table>
        <tr>
          <td>File</td>
          <td>Link</td>
          <td>Progress</td>
        </tr>
        {{ range $download, $b := .Downloaded }}
          <tr>
            <td><a href="{{ $download.Url }}">{{ $download.Filename }}</a></td>
            <td>{{ $download.Link }}</td>
            <td>{{ $download.Progress }}</td>
          </tr>
        {{ end }}
      </table>
    </body>
  </html>
{{ end }}

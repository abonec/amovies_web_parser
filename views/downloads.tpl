{{ define "downloads_page" }}
  <html>
  <head>
    {{ template "assets" }}
  </head>
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
          <td>Remove</td>
        </tr>
        {{ range $id, $download := .Downloaded }}
          <tr>
            <td><a href="{{ $download.Url }}">{{ $download.Filename }}</a></td>
            <td>{{ $download.Link }}</td>
            <td>{{ $download.Progress }}</td>
            <td><a href="#" class="remove_download" data-id="{{ $id }}">(X)</a></td>
          </tr>
        {{ end }}
      </table>
    </body>
  </html>
{{ end }}

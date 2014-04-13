{{ define "get_serial_form" }}
  <form action='/links' method='get'>
    <fieldset>
      <input name='url' type='text'>
    </fieldset>
    <input type='submit' value='Get Serials'>
  </form>
{{ end }}

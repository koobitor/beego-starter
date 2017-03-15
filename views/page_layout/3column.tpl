{{ template "layout/header.tpl" . }}

  <div class="left">
    {{ template "left" . }}
  </div>

  <div class="content">
    {{ template "content" . }}
  </div>

  <div class="right">
    {{ template "right" . }}
  </div>

{{ template "layout/footer.tpl" . }}
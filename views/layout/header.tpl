<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>{{.title}}</title>
  {{ template "meta" . }}

  <!-- stylesheet -->
  <link rel="stylesheet" href="/static/vendor/bootstrap/dist/css/bootstrap.min.css">
  <link rel="stylesheet" href="/static/vendor/bootstrap/dist/css/bootstrap-theme.min.css">
  <link rel="stylesheet" href="/static/css/style.css">
  {{ template "css" . }}
  
</head>
<body class="{{.classes}}">
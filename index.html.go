package main
const index = 
`<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width,initial-scale=1,shrink-to-fit=no">
    <title>Elements</title>

    <meta name="theme-color" content="#000000">
    <link rel="manifest" href="manifest.json">
    <link rel="shortcut icon" href="static/media/mdc_web_48dp.png">
    <!--
        <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
    -->
    <link rel="stylesheet" href="static/fonts/material-icons/material-icons.css">

    <!--
        <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Rubik:300,400,500">
        <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto+Mono">
        <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,600,700">
    -->
    <link rel="stylesheet" href="static/fonts/rubik/rubik.css">
    <link rel="stylesheet" href="static/fonts/roboto/roboto.css">

    <script type="text/javascript" src="static/js/mdc/material-components-web.min.js"></script>

    <link rel="stylesheet" href="static/css/normalize.min.css" />
    <link rel="stylesheet" href="static/css/material-components-web.min.css" />

    {{.Styles}}

</head>

<body class="mdc-typography mdc-theme--dark page-slide">
    <noscript>You need to enable JavaScript to run this app.</noscript>

    <!-- App Toolbar -->
    <!-- to shadows use class: mdc-toolbar--fixed -->
    <header class="mdc-top-app-bar mdc-top-app-bar--fixed">
        <div class="mdc-top-app-bar__row">
            <section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-start">
                <a href="#" class="back material-icons mdc-top-app-bar__navigation-icon mdc-icon-back">arrow_back</a>
                <span class="mdc-top-app-bar__title">Selecione...</span>
            </section>
            <section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-end" role="toolbar">
                <a href="#search" class="material-icons mdc-top-app-bar__action-item" aria-label="Search" alt="Search">search</a>
            </section>
        </div>
    </header>

    <main class="page">

        <div class="mdc-top-app-bar--fixed-adjust">

            {{with .Banks}}

            <ul id="banks" class="mdc-list" aria-orientation="vertical">{{range .}}
                <li><a class="mdc-list-item" href="#bank/{{.Id}}">
                    <span class="mdc-list-item__graphic material-icons" aria-hidden="true">table_chart</span>
                    {{.Id}} / {{.Name}}
                    <span class="mdc-list-item__meta material-icons" aria-hidden="true">chevron_right</span>
                </a></li>{{end}}
            </ul>

        </div>

    </main>

    {{- range .}}
    <section class="page page-bank{{.Id}} right">

        <div class="mdc-top-app-bar--fixed-adjust">

            <h3 class="mdc-list-item"><span class="mdc-list-item__graphic material-icons" aria-hidden="true">table_chart</span>{{.Id}} / {{.Name}}</h3>
            <ul id="bank{{.Id}}" class="mdc-list" aria-orientation="vertical">{{range .Instruments}}
                <li class="mdc-list-item">
                    <span class="mdc-list-item__graphic material-icons" aria-hidden="true">audiotrack</span>
                {{.Address}} / {{.Name}}</li>{{end}}
            </ul>

        </div>

    </section>
    {{- end}}

    {{- end}}

</body>

{{.Scripts}}

</html>`
package main
const styles =
    `<style>
    body{
        overflow-y: scroll;
    }
    
    /* THEME */
    :root {
        --mdc-theme-primary: #343D46;
        --mdc-theme-secondary: #25c6d8;  
        --mdc-theme-background: #414A52;
        --mdc-theme-surface: #464F58;
        --mdc-theme-on-primary: #fff;
        --mdc-theme-on-secondary: #fff;
        --mdc-theme-on-surface: #000;
        --mdc-theme-text-primary-on-background: rgba(255,255,255,.87);
        --mdc-theme-text-secondary-on-background: rgba(0,0,0,.54);
        --mdc-theme-text-hint-on-background: rgba(255,255,255,.38);
        --mdc-theme-text-disabled-on-background: rgba(255,255,255,.38);
        --mdc-theme-text-icon-on-background: rgba(255,255,255,.38);

        --mdc-theme-text-primary-on-light: rgba(0,0,0,.87);
        --mdc-theme-text-secondary-on-light: rgba(0,0,0,.54);
        --mdc-theme-text-hint-on-light: rgba(0,0,0,.38);
        --mdc-theme-text-disabled-on-light: rgba(0,0,0,.38);
        --mdc-theme-text-icon-on-light: rgba(0,0,0,.38);

        --mdc-theme-text-primary-on-dark: #fff;
        --mdc-theme-text-secondary-on-dark: hsla(0,0%,100%,.7);
        --mdc-theme-text-hint-on-dark: hsla(0,0%,100%,.5);
        --mdc-theme-text-disabled-on-dark: hsla(0,0%,100%,.5);
        --mdc-theme-text-icon-on-dark: hsla(100%,100%,100%,.5);
    }
    .mdc-theme--dark{
        background-color:var(--mdc-theme-background, #414A52);
        color:#fff;
    }
    .mdc-snackbar{
        background-color:#242D36;
    }



    /* BUTTON PANEL */
    .s-card__title, .s-card__subtitle{
        margin: 0;
    }
    .s-card__subtitle{
        color: rgba(255, 255, 255, 0.54);
        
    }
    .s-card__primary{       
        padding: 8px 16px 0 16px;
    }
    .s-card__secondary{
        padding: 0 1rem 8px 1rem;
    }
    .s-card__actions {
        padding: 0 16px 0 0;
    }
    .mdc-grid-tile {
        width : 70px;
    }

    /* HERO */
    .hero {
        display: -ms-flexbox;
        display: flex;
        -ms-flex-flow: row nowrap;
        flex-flow: row nowrap;
        -ms-flex-align: center;
        align-items: center;
        -ms-flex-pack: center;
        justify-content: center;
        padding: 16px;
    }
    .mdc-image-list {
        display: -ms-flexbox;
        display: flex;
        -ms-flex-wrap: wrap;
        flex-wrap: wrap;
        margin: 0 auto;
        padding: 0;
    }
    .hero-image-list {
        width: 300px;
        margin: 0;
    }
    .hero-image-list .mdc-image-list__item {
        width: calc(100% / 4 - 4.2px);
        margin: 2px;
        justify-content: center;
        align-items: center;
        display: flex;
        background-color: rgba(0,0,0,.26);
    }

    .hero .mdc-image-list__image-aspect-container {
        padding: calc(40% / 1);
    }
    .hero-image-list i{
        padding: 3px 7px;
        border-bottom: 2px solid var(--mdc-theme-background);
        border-right: 2px solid var(--mdc-theme-background);
    }
    .hero-image-list i::after{
        background-color: #00BCD4;
        padding: 7px 7px;
        /* border: 2px solid white; */
        /* border: 2px solid white; */
        position: absolute;
        content: '';
        /* bottom: 0; */
        left: 0;
        /* height: 10%; */
        /* width: 44%; */
        border: 2px solid #00BCD4;
        border-right: 2px solid #00BCD4;
        transform: skew(-30deg);
        z-index: -1;
    }

    @media(min-width: 768px) {
        .mdc--hide-desktop {
            display: none;
        }
    }


    /* SHDO THEME */

    .mdc-card{
        border: 1px solid #546E7A;
        border-radius: 4px;
        box-shadow: none;
    }

    #masterVolume_icon{
        color: var(--mdc-theme-secondary);
    }


/* PAGES - For SLIDE */

.page-slide {
  position: fixed;
  /*width: 100%;
  height: 100%;*/
  overflow: hidden;
}

.page-fixed {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  -webkit-transform: translate3d(0, 0, 0);
  transform: translate3d(0, 0, 0);
  text-align: left;
}

.page {
    position: absolute;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    -webkit-transform: translate3d(0, 0, 0);
    transform: translate3d(0, 0, 0);
    text-align: left;
    -webkit-transition-duration: .25s;
    transition-duration: .25s;
    background-color: var(--mdc-theme-background, #414A52);
    overflow-y: scroll;
}

.page.left {
  -webkit-transform: translate3d(-105%, 0, 0);
  transform: translate3d(-105%, 0, 0);
}

.page.center {
  -webkit-transform: translate3d(0, 0, 0);
  transform: translate3d(0, 0, 0);
}

.page.right {
  -webkit-transform: translate3d(105%, 0, 0);
  transform: translate3d(105%, 0, 0);
}

.page.transition {
  -webkit-transition-duration: .25s;
  transition-duration: .25s;
}

.page-mw{
  max-width: 400px;
  -webkit-box-shadow: 0 8px 10px -5px rgba(0,0,0,.2), 0 16px 24px 2px rgba(0,0,0,.14), 0 6px 30px 5px rgba(0,0,0,.12);
  box-shadow: 0 8px 10px -5px rgba(0,0,0,.2), 0 16px 24px 2px rgba(0,0,0,.14), 0 6px 30px 5px rgba(0,0,0,.12);
}

.tab-content{ 
  height: 93%;
  height: calc(100% - 56px);
  width: 100%;
  overflow-y: auto;
  margin-top: 56px;
}
</style>`
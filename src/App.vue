<template>
    <div id="app">
        <div class="wrapper">
            <router-view class="router" />
        </div>
        <Footer />
    </div>
</template>
<script>
    import Footer from './components/layout/Footer.vue';

    export default {
        components: { Footer },
    }
</script>
<style>
    :root {
        --background-primary: #FEF3ED;
        --background-secondary: #DEE8EA;
        --background-tertiary: #A1DED1;
        --emphasis-color: #F99170;
        --title-color: #404575;

        font-size: 0.2666666667vw;
    }

    * {
        margin: 0;
        padding: 0;
        outline: none;
        text-decoration: none;
        box-sizing: border-box;
    }

    html, body, #app, .router, .page {
        height: 100%;
    }

    body {
        font-family: 'Nunito', sans-serif;
        background-color: var(--background-primary);
    }

    .wrapper {
        position: relative;
        max-width: 550px;
        height: calc(100% - 50rem);
        margin: 0 auto;
        background-color: var(--background-secondary);
        border-radius: 0 0 40rem 40rem;
        box-shadow: 0 2rem 2rem #E3DCDC;
        overflow: hidden;
    }

    .container {
        position: relative;
        padding: 0 15rem;
    }

    a, button, input, textarea {
        font-family: inherit;
        color: inherit;
        font-size: inherit;
    }

    li {
        list-style-type: none;
    }

    @media (min-width: 550px) {
        :root {
            font-size: 1.466668px;
        }
    }
</style>
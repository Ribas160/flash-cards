<template>
    <footer>
        <div class="container">
            <router-link :to="{name: 'lists'}">
                <span class="material-icons-outlined">view_list</span>
            </router-link>
            <router-link :to="{name: 'lists'}">
                <span class="material-icons-outlined">rocket</span>
            </router-link>
            <router-link :to="{name: 'lists'}">
                <span class="material-icons-outlined">settings</span>
            </router-link>
        </div>
    </footer>
</template>
<style scoped>
    footer {
        position: absolute;
        left: 0;
        right: 0;
        bottom: 0;
        max-width: 550px;
        margin: 0 auto;
        padding: 10rem 0;
    }

    footer .container {
        display: flex;
        align-items: center;
        justify-content: space-around;
    }

    footer span {
        font-size: 30rem;
        color: var(--emphasis-color);
    }

    footer a {
        line-height: 1;
    }
</style>
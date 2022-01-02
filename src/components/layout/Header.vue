<template>
    <header>
        <div class="container">
            <router-link class="back" v-if="back" :to="{ name: back.to }">{{ back.name }}</router-link>
            <h1 v-if="title">{{ title }}</h1>
            <button class="edit" v-if="edit" @click="onClick">Править</button>
        </div>
    </header>
</template>
<script>
    export default {
        props: {
            back: {
                type: Object,
                required: false,
            },

            title: {
                type: String,
                required: false,
                default: '',
            },

            edit: {
                type: Boolean,
                required: false,
                default: false,
            }
        },
        methods: {
            onClick() {
                this.$emit('edit');
            }
        }
    }
</script>
<style scoped>
    header {
        height: 50rem;
        margin-bottom: 10rem;
    }

    header .container {
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    h1, button {
        color: var(--title-color);
        font-weight: 700;
    }

    h1 {
        position: absolute;
        top: 50%;
        transform: translateY(-50%);
        left: 0;
        right: 0;
        font-size: 25rem;
        text-align: center;
    }

    button {
        font-size: 15rem;
        background-color: transparent;
        border: none;
    }

    .back {
        margin-right: auto;
    }

    .edit {
        margin-left: auto;
    }
</style>
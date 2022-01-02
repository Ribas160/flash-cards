<template>
    <main>
        <div class="container">
            <ul>
                <li v-for="(item, index) in items" :key="index">
                    <router-link :to="{ name: 'list', params: {id: item.id} }">{{ item.name }}</router-link>
                </li>
            </ul>
        </div>
    </main>
</template>
<script>
    export default {
        props: {
            items: {
                type: Array,
                required: true,
            },
        },
    }
</script>
<style scoped>
    li {
        height: 45rem;
        margin-top: 20rem;
        padding: 0 15rem;
        font-size: 16rem;
        font-weight: 700;
        color: #fff;
        background-color: var(--background-tertiary);
        border-radius: 10rem;
        box-shadow: 0 0 1rem 1rem #B2D7D1;
    }

    li:first-of-type {
        margin-top: 0;
    }

    li a {
        height: 100%;
        display: block;
        display: flex;
        align-items: center;
    }
</style>
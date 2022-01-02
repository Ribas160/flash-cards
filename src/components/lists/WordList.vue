<template>
    <main>
        <div class="container">
            <ul>
                <li v-for="(item, index) in items" :key="index">
                    <span>{{ items.ru }}</span>
                    <span>{{ items.en }}</span>
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
        margin-top: 10rem;
        height: 75rem;
        padding: 0 15rem;
        display: flex;
        flex-direction: column;
        font-size: 16rem;
        font-weight: 700;
        color: var(--title-color);
        background-color: var(--background-primary);
        border-radius: 10rem;
    }

    li:first-of-type {
        margin-top: 0;
    }
</style>
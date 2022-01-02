<template>
    <main>
        <Header :back="{name: 'Списки', to: 'lists'}" :edit="true" />
        <CardCarousel :cards="cards" />
        <ActionButton :role="'add'" />
    </main>
</template>
<script>
    import Header from '../components/layout/Header.vue';
    import CardCarousel from '../components/CardCarousel.vue';
    import ActionButton from '../components/ActionButton.vue';

    export default {
        components: { Header, CardCarousel, ActionButton },
        data() {
            return {
                cards: [
                    {
                        ru: 'Машина',
                        en: 'Car',
                    },
                    {
                        ru: 'Собака',
                        en: 'Dog',
                    },
                    {
                        ru: 'Дом',
                        en: 'House',
                    },
                    {
                        ru: 'Жена',
                        en: 'Wife',
                    },
                    {
                        ru: 'Лампа',
                        en: 'Lamp',
                    },
                ],
            };
        },
    }
</script>
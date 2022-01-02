<template>
    <div class="page">
        <Header :title="'Списки'" :edit="true" />
        <MenuList :items="items" />
        <ActionButton :role="'add'" @add="toggleForm" v-if="!form" />
        <ActionForm v-if="form" :fields="fields" @close="toggleForm" @submit="submit" />
    </div>
</template>
<script>
    import Header from '../components/layout/Header.vue';
    import MenuList from '../components/lists/MenuList.vue';
    import ActionButton from '../components/ActionButton.vue';
    import ActionForm from '../components/ActionForm.vue';

    export default {
        components: { Header, MenuList, ActionButton, ActionForm },
        data() {
            return {
                items: [
                    {
                        name: '150 английских существительных',
                        id: 1,
                    },
                    {
                        name: '150 английских прилагательных',
                        id: 2,    
                    },
                    {
                        name: '150 английских глаголов',
                        id: 3,
                    },
                    {
                        name: 'Мои слова',
                        id: 4,
                    },
                ],
                fields: [
                    {
                        type: 'text',
                        name: 'en',
                        placeholder: 'Слово',
                        required: true,
                    },
                    {
                        type: 'text',
                        name: 'ru',
                        placeholder: 'Значение',
                        required: true,
                    }
                ],
                form: false,
            };
        },
        methods: {
            toggleForm() {
                this.form = !this.form;
            },
            submit(data) {
                console.log(data);
            }
        }
    }
</script>
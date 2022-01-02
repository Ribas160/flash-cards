<template>
    <section class="cardCarousel">
        <div class="container">
            <Slick 
                ref="slick" 
                :options="slickOptions"
                @swipe="onSwipe"
            >
                <div class="slide" v-for="(card, index) in cards" :key="index" @click="onClick">
                    <div class="front">{{ card.en }}</div>
                    <div class="back">{{ card.ru }}</div>
                </div>
            </Slick>
        </div>
    </section>
</template>
<script>
    import Slick from 'vue-slick';
    import 'slick-carousel/slick/slick.css';

    export default {
        components: { Slick },
        props: {
            cards: {
                type: Array,
                required: true,
            },
        },
        data() {
            return {
                slickOptions: {
                    slidesToShow: 1,
                    arrows : false,
                },
            };
        },
        methods: {
           onClick(e) {
               e.currentTarget.classList.toggle('rotate');
           },

           onSwipe() {
               this.$el.querySelectorAll('.slide').forEach(el => {
                   el.classList.remove('rotate');
               });
           }
        },
    }
</script>
<style scoped>
    .cardCarousel {
        position: absolute;
        top: 50%;
        transform: translateY(-50%);
        left: 0;
        right: 0;
    }

    .slide {
        position: relative;
        width: 100%;
        height: 392rem;
        perspective: 1000px;
    }

    .slide:hover {
        cursor: pointer;
    }

    .rotate .front {
        transform: rotateY(180deg);
    }

    .rotate .back {
        transform: rotateY(360deg);
    }

    .front, .back {
        position: absolute;
        top: 0;
        left: 0;
        width: calc(100% - 10rem);
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 30rem;
        font-weight: bold;
        border-radius: 10rem;
        box-shadow: 0 5rem 15rem 1rem rgba(0, 0, 0, 0.2);
        backface-visibility: hidden;
        transition: 1s;
    }

    .front {
        color: var(--title-color);
        background-color: var(--background-primary);
    }

    .back {
        color: #fff;
        background-color: var(--background-tertiary);
        transform: rotateY(180deg);
    }
</style>

<style>
    .slick-list {
        overflow: visible !important;
    }
</style>
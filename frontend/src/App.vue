<template>
    <main class="container py-4">
        <header class="pb-3 mb-4 border-bottom">
            <span class="d-flex align-items-center text-dark text-decoration-none fs-4">💡 QUIZ Light</span>
        </header>
        <div class="p-5 mb-4 bg-light rounded-3">
            <Home @startTheGame="startTheGame" v-if="showHome"/>
            <Question
                    v-if="gameStarted && quizData && !gameFinished"
                    :quizData="this.quizData"
                    @endGame="endGame"
            />
            <Final
                    v-if="gameFinished"
                    :correctAnswers="correctAnswers"
                    :numberOfQuestions="numberOfQuestions"
                    @resetGame="resetGame"
            />
        </div>
        <footer class="pt-3 mt-4 text-muted border-top">
            &copy; 2023
        </footer>
    </main>
</template>

<script>

import axios from "./services/axios";
import Home from "./components/Home.vue";
import Question from "./components/Question.vue";
import Final from "./components/Final.vue";

export default {
    name: "App",
    components: {
        Home,
        Question,
        Final,
    },
    data() {
        return {
            showHome: true,
            gameStarted: false,
            gameFinished: false,
            quizData: undefined,
            numberOfQuestions: 0,
        };
    },
    methods: {
        startTheGame({quizName}) {
            this.makeRequest(quizName);
            this.showHome = false;
            this.gameStarted = true;
        },
        async makeRequest(quizName) {
            try {
                const {data} = await axios.get(
                    `/quizzes/${quizName}`
                );
                this.quizData = data;
                console.log(data);
            } catch (error) {
                console.error(error);
            }
        },
        endGame({numberOfQuestions, correctAnswers}) {
            this.numberOfQuestions = numberOfQuestions;

            this.correctAnswers = correctAnswers;
            this.gameFinished = true;
        },
        resetGame() {
            this.showHome = true;
            this.gameStarted = false;
            this.gameFinished = false;
            this.quizData = undefined;
            this.numberOfQuestions = 0;
        },
    },
};
</script>
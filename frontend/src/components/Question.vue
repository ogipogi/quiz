<template>
    <section class="question">
        <h1>
            <span>{{ index + 1 }}</span>. <span v-html="questionText"></span>
        </h1>
        <div class="timer">
            <span class="clock"></span>
            <p>Timer 0:{{ secondsLeft }}</p>
        </div>

        <div class="d-grid gap-2 my-2" v-for="(item, index) in answers" :key="index">
            <button class="btn btn-outline-dark"
                    type="button"
                    @click="checkAnswer($event)"
                    :disabled="roundEnded">
                {{ item }}
            </button>
        </div>

        <div class="post-answer" :style="[roundEnded ? { display: 'block' } : { display: 'none' }]">
            <p :class="[isUserAnswerCorrect ? 'text-success' : 'text-danger']">
                {{ message }}
            </p>
            <button class="btn btn-primary"
                    @click="goToNextQuestion"
                    v-html="[index === numberOfQuestions - 1 ? 'See results!' : 'Next question',]">
            </button>
        </div>
    </section>
</template>

<script>

export default {
    name: "Question",
    props: {
        quizData: Object,
    },
    data() {
        return {
            gameFinished: false,
            correctAnswers: 0,
            secondsLeft: 0,
            index: 0,
            numberOfQuestions: 0,
            answers: [],
            correctAnswer: "",
            questionText: "",
            userAnswered: false,
            roundEnded: false,
            isUserAnswerCorrect: undefined,
            message: "",
        };
    },
    mounted() {
        this.prepareQuestion();
    },
    methods: {
        prepareQuestion() {
            const questions = this.quizData.questions;
            const index = this.index;

            this.numberOfQuestions = questions.length;
            this.questionText = questions[index].text;
            questions[index].answers.forEach(answer => {
                this.answers.push(answer.text);
                if (answer.correct) {
                    this.correctAnswer = answer.text
                }
            });

            this.countdown();
        },

        countdown() {
            this.secondsLeft = 45;
            this.timer = setInterval(() => {
                if (!this.userAnswered) {
                    if (this.secondsLeft > 0) this.secondsLeft--;
                    else if (this.secondsLeft <= 0 || this.roundEnded) this.stopTimer();
                    if (this.secondsLeft.toString().length < 2)
                        this.secondsLeft = `0${this.secondsLeft}`;
                }
                if (this.secondsLeft <= 0 && !this.userAnswered) {
                    this.message = "Time ran out! No points in this round.";
                    this.roundEnded = true;
                }
            }, 1000);
        },

        stopTimer() {
            clearInterval(this.timer);
        },

        checkAnswer(event) {
            if (!this.roundEnded) {
                this.userAnswered = true;
                this.roundEnded = true;

                this.isUserAnswerCorrect = event.target.innerText === this.correctAnswer;
                this.message = this.isUserAnswerCorrect ? "Good job!" : "Better luck next time.";

                if (this.isUserAnswerCorrect)
                    this.correctAnswers++;
            }
        },

        goToNextQuestion() {
            this.checkIfItIsTheLastRound();

            if (!this.gameFinished) {
                this.index++;
                this.roundEnded = false;
                this.userAnswered = false;
                this.isUserAnswerCorrect = undefined;
                this.answers = [];
                this.stopTimer();
                this.prepareQuestion();
            }
        },

        checkIfItIsTheLastRound() {
            if (this.index === this.numberOfQuestions - 1) {
                this.gameFinished = true;
                this.$emit("endGame", {
                    numberOfQuestions: this.numberOfQuestions,
                    correctAnswers: this.correctAnswers,
                });
            }
        },
    },
};
</script>

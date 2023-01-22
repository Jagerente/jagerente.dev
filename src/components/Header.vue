<template>
    <div class="header">
        <img
            class="header__icon"
            src="@/assets/logo.png"
        >
        <h1 class="header__title">Jagerente</h1>
        <p class="header__description">Quit? Don't quit? Noodles!</p>
    </div>
</template>

<style lang="scss" scoped>
.header {
    top: 0;
    left: 0;
    width: max-content;
    margin: 2em 0 1em 2em;

    display: grid;
    grid-template-areas:
        "a b"
        "a c";
    gap: 0 1em;

    @media (max-width: $sm) {
        position: absolute;
        margin: 4em auto 2em auto;
        position: relative;
        grid-template-areas:
            "a"
            "b"
            "c";
    }

    &__icon {
        grid-area: a;
        height: 5rem;
        margin: auto;
        filter: blur(.1px);
    }

    &__title {
        font-size: 2rem;
        border-bottom: $primary solid 3px;
        padding-bottom: 0.15rem;
        margin-bottom: 0.2rem em;
        grid-area: b;
    }

    &__description {
        grid-area: c;
        font-size: 1rem;
        margin-top: .3rem;
    }

}
</style>
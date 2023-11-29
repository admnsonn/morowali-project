<template>
    <div class="menu-item">
        <a href="#">
            {{ title }}
        </a>
        <svg viewBox="0 0 1030 638" width="10">
            <path d="M1017 68L541 626q-11 12-26 12t-26-12L13 68Q-3 49 6 24.5T39 0h952q24 0 33 24.5t-7 43.5z" fill="#FFF">
            </path>
        </svg>
    </div>
</template>

<script>
import { title } from 'process';

export default {
    name: 'dropdown'
}
</script>
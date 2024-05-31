<template>
    <div>
        <div class="objects"></div>
        <div class="content-browser"></div>
        <div class="inspector"></div>
        <canvas ref="scene" width="600" height="500"></canvas>
    </div>
</template>

<script>
export default {
    mounted() {
        const canvas = this.$refs.scene;
        const ctx = canvas.getContext('2d');

        ctx.fillStyle = 'black';
        ctx.fillRect(0, 0, canvas.width, canvas.height);

        ctx.strokeStyle = 'white';

        ctx.beginPath();
        ctx.arc(100, 100, 50, 0, 2 * Math.PI);
        ctx.stroke();
    }
}
</script>
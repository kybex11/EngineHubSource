<template>
    <div>
        <canvas ref="scene" width="200" height="200"></canvas>
    </div>
</template>

<script>
export default {
    mounted() {
        const canvas = this.$refs.scene;
        const ctx = canvas.getContext('2d');

        ctx.beginPath();
        ctx.arc(100, 100, 50, 0, 2 * Math.PI);
        ctx.stroke();
    }
}
</script>
<template>
    <div>
        <h1>{{sceneContent}}</h1>
        <div class="objects"></div>
        <div class="content-browser"></div>.
        <div class="inspector"></div>
        <canvas ref="scene" width="600" height="500"></canvas>
    </div>
</template>

<script>
import { ReadScenes } from '../../../wailsjs/go/main/App';

export default {
    data() {
        return {
            sceneContent: "",
        }
    },
    props: {
        projectName: {
            type: Object,
            required: true
        },
    },
    mounted() {
        ReadScenes(`${this.projectName}${this.projectName}`)
            .then(response => {
                console.log("reading scene file");
                if (response) {
                    this.sceneContent = response;
                } else {
                    console.error("Empty response received");
                }
            })
            .catch(error => {
                console.error("Error reading scene file:", error);
            });

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
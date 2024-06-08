<template>
    <div>
        <div class="objects"></div>
        <div class="content-browser"></div>.
        <div class="inspector"></div>
        <canvas ref="scene" width="600" height="500"></canvas>
    </div>
</template>

<script>
import { ReadScenes, CreateObject } from '../../../wailsjs/go/main/App';

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
                    console.log(response);

                    const lines = response.split("\n");
                    const action = lines[0].substring(0, 10).trim(); 
                    const name = lines[0].substring(10).trim();

                    const args = lines.slice(1).map(arg => {
                        const [key, value] = arg.split("(").map(item => item.trim().replace(/[()]/g, ""));
                        return { key, value };
                    });

                    if (action == "drawCircle") {
                        console.log("Action:", action);
                        console.log("Name:", name);
                        console.log("Arguments:");
                        args.forEach(arg => {
                            console.log(`const ${arg.key} = ${arg.value}`);
                        });
                    
                        const argumentsObject = {};
                        args.forEach(arg => {
                            argumentsObject[arg.key] = arg.value;
                        });
                    
                        const { color, posX, posY, radius } = argumentsObject;
                    
                        const canvas = this.$refs.scene;
                        const ctx = canvas.getContext('2d');
                        ctx.fillStyle = 'black';
                        ctx.fillRect(0, 0, canvas.width, canvas.height);
                    
                        ctx.strokeStyle = color;
                        ctx.beginPath();
                        ctx.arc(posX, posY, radius, 0, 2 * Math.PI);
                        ctx.stroke();
                    }

                    
                } else {
                    console.error("Empty response received");
                }
            })
            .catch(error => {
                console.error("Error reading scene file:", error);
            });
    },
    methods: {
        createObj(objName, objData) {
            CreateObject(objName, objData, this.projectName)
                .then(response => {
                    console.log(respnose);
                })
                .catch(e => {
                    console.error("Error creating object:",e);
                });
        }
    }
}
</script>
<template>
    <div class="">
        <div class="ScenePortal">
            <div  v-if="projectMode === '2d'">
                <two/>
            </div>
            <div v-if="projectMode === '3d'">
                <tree/>
            </div>
        </div>
    </div>
</template>

<script>
import two from './scenes/two.vue'
import tree from './scenes/tree.vue'
export default {
    components: { two, tree },
    props: {
        projectData: {
            type: Object,
            required: true
        },
        projectMode: {
            type: Object,
            required: true
        }
    }
}
</script>

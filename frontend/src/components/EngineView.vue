<template>
    <div class="upper">
        <button>File</button>
        <button>Edit</button>
        <button>Window</button>
        <button>Tools</button>
        <button>Build</button>
        <button>Select</button>
        <button>Actor</button>
    </div>
    <div class="">
        <h1>{{ projectData }}</h1>
    </div>
</template>

<script>
export default {
    props: {
        projectData: {
            type: Object,
            required: true
        }
    }
}
</script>

<style>
.upper {
    position: absolute;
    top: 0;
    left:0;
}
.upper button {
    background: none;
    border: none;
    color: white;
    cursor: pointer;
}
h1 {
    color: white;
}
</style>
<template>
    <div class="ProjectSelector">
        <h1>Select Project</h1>
        <br>
        <div class="ProjectsView" v-if="!toggleNew">
            <h3>Selected: {{ selectedProj }} <button @click="selectProject()">Unselect</button></h3>
            <ul>
                <li v-for="project in projects" :key="project">
                    <button class="ProjectContainButton" @click="selectProject(project)">{{ project }}</button>
                </li>
            </ul> 
        </div>
        <div class="ProjectsView" v-else>
            <h1>Create Project</h1>
        </div>
        <br>
        <div class="ProjectsViewEdit">
            <button @click="createProject">Create</button>
            <button @click="editProject">Edit</button>
            <button @click="openProject">Open</button>
            <button @click="getProjects">Update</button>
            <button @click="deleteProject">Delete</button>
        </div>
    </div>
</template>

<script>
import { ListProjects, DeleteProject, CreateProject } from '/wailsjs/go/main/App';

export default {
    data() {
        return {
            createProjectContainerView: false,
            editProjectContainerView: false,
            deleteProjectContainerView: false,
            openProjectContainerView: false,
            projects: [],
            toggleNew: false,
            selectedProj: "",
        }
    },
    mounted() {
        this.getProjects();
    },
    methods: {
        selectProject(project) {
            this.selectedProj = project;
        },
        openProject() {
            console.log("open project");
        },
        getProjects() {
            ListProjects("projects")
                .then(response => {
                  console.log("Received data:", response);
                  this.projects = response;
                })
                .catch(error => {
                  console.error('Error fetching projects:', error);
                })

            console.log("get projects");
        },
        createProject() {
            console.log("create project");
        },
        editProject() {
            console.log("edit project");
        },
        deleteProject() {
            if (this.selectedProj) {
                // Убедитесь, что selectedProj не содержит расширения .json
                const projectName = this.selectedProj.replace('.json', '');
                const filePath = `projects/${projectName}/${projectName}.json`;
                DeleteProject(filePath)
                    .then(response => {
                        console.log("Project deleted successfully:", response);
                        this.getProjects(); // Обновить список проектов после удаления
                    })
                    .catch(error => {
                        console.error('Error deleting project:', error);
                    })
            } else {
                console.log("No project selected");
            }
        }
    }
}
</script>

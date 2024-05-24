<template>
    <div class="ProjectSelector">
        <h1 v-if="!toggleNew && !engineUsed">Select Project</h1>
        <div class="ProjectsView" v-if="!toggleNew && !engineUsed">
            <h3>Selected: {{ selectedProj }} <button @click="selectProject()">Unselect</button></h3>
            <ul>
                <li v-for="project in projects" :key="project">
                    <button class="ProjectContainButton" @click="selectProject(project)">{{ project }}</button>
                </li>
            </ul> 
        </div>
        <div class="ProjectsView" v-if="toggleNew">
            <br>
            <h2 class="CP">Create Project</h2>
            <br>
            <input type="text" v-model="newProjName" class="project-input" placeholder="Name of proj">
            <br>
            <button class="ProjectsViewSubmitButtonContainerClass" @click="createProject">Ok</button>
            <button class="ProjectsViewSubmitButtonContainerClass" @click="cancelCreateProject">Cancel</button>
        </div>
        <br>
        <div class="ProjectsViewEdit" v-if="!toggleNew && !engineUsed">
            <button @click="createProject">Create</button>
            <button @click="editProject">Edit</button>
            <button @click="openProject">Open</button>
            <button @click="getProjects">Update</button>
            <button @click="deleteProject">Delete</button>
        </div>
    </div>
    <div class="EngineView" v-if="engineUsed">
        <EngineView :projectData="selectedProj"/>
    </div>
</template>

<script>
import EngineView from './components/EngineView.vue'
import { ListProjects, DeleteProject, CreateProject } from '/wailsjs/go/main/App';

export default {
	components: { EngineView },
    data() {
        return {
            newProjName: '',
            createProjectContainerView: false,
            editProjectContainerView: false,
            deleteProjectContainerView: false,
            openProjectContainerView: false,
            projects: [],
            toggleNew: false,
            selectedProj: "",
            engineUsed: false,
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
            if (this.selectedProj !== "") {
                this.engineUsed = !this.engineUsed;
                console.log("open project");
            } else {
                console.log("failing open project");
            }

        },
        cancelCreateProject() {
            this.toggleNew = !this.toggleNew;
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
            if (!this.toggleNew) {
                this.toggleNew = !this.toggleNew;
            } else {

                const newProjData = {
                    name: this.newProjName,
                };

                CreateProject(this.newProjName, newProjData)
                    .then(response => {
                        console.log("Project created successfully: ", response);
                        this.projects.push(this.newProjName);
                        this.selectedProj = "";
                    })
                    .catch(error => {
                        console.error('Error creating project: ', error);
                    })

                this.toggleNew = !this.toggleNew;
                this.getProjects();
            }
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
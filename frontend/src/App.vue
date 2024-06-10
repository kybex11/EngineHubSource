<template>
    <div class="ProjectSelector">
        <h1 v-if="!toggleNew && !engineUsed && !toggleEdit">Select Project</h1>
        <div class="ProjectsView" v-if="!toggleNew && !engineUsed && !toggleEdit">
            <h3>Selected: {{ selectedProj }} <button @click="selectProject()">Unselect</button></h3>
            <ul>
                <li v-for="project in projects" :key="project">
                    <button class="ProjectContainButton" @click="selectProject(project)">{{ project.substring(2)
                        }}</button>
                </li>
            </ul>
        </div>
        <div class="ProjectsView" v-if="toggleEdit">
            <br>
            <h2 class="CP">Edit Project: {{selectedProj}}</h2>
            <br>
            <input type="text" v-model="newEditName" class="project-input" placeholder="New name for project">
            <br><br>
            <button class="SelectModeButton" @click="editPrj">Edit</button>
            <button class="SelectModeButton" @click="editProject"></button>
        </div>
        <div class="ProjectsView" v-if="toggleNew">
            <br>
            <h2 class="CP">Create Project</h2>
            <br>
            <input type="text" v-model="newProjName" class="project-input" placeholder="Name of proj">
            <br><br>
            <button class="SelectModeButton" @click="selectMode('2d')">2D</button>
            <button class="SelectModeButton" @click="selectMode('3d')">3D</button>
            <br>
            <button class="ProjectsViewSubmitButtonContainerClass" @click="createProject">Ok</button>
            <button class="ProjectsViewSubmitButtonContainerClass" @click="cancelCreateProject">Cancel</button>
        </div>
        <br>
        <div class="ProjectsViewEdit" v-if="!toggleNew && !engineUsed && !toggleEdit">
            <button @click="createProject">Create</button>
            <button @click="editProject">Edit</button>
            <button @click="openProject">Open</button>
            <button @click="getProjects">Update</button>
            <button @click="deleteProject">Delete</button>
        </div>
    </div>
    <div class="EngineView" v-if="engineUsed">
        <EngineView :projectData="selectedProj" :projectMode="selectedProjMode" />
    </div>
</template>

<script>
import EngineView from './components/EngineView.vue'
import { ListProjects, DeleteProject, CreateProject, RenameProject } from '/wailsjs/go/main/App';

export default {
    components: { EngineView },
    data() {
        return {
            newProjName: '',
            createProjectContainerView: false,
            editProjectContainerView: false,
            deleteProjectContainerView: false,
            openProjectContainerView: false,
            newEditName: "",
            selectedProj: "",
            selectedProjMode: "",
            projectStokeName: "",
            toggleEdit: false,
            projects: [],

            toggleNew: false,

            engineUsed: false,
        }
    },
    mounted() {//
        this.getProjects();
    },
    methods: {
        selectMode(mode) {
            this.selectedProjMode = mode;
        },
        selectProject(project) {
            if (project) {
                this.projectStokeName = project;
                this.selectedProjMode = project.substring(0, 2); // Получить первые два символа
                this.selectedProj = project.substring(2); // Получить оставшуюся часть строки
            } else {
                this.projectStrokeName = "";
                this.selectedProjMode = "";
                this.selectedProj = "";
            }
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
        editPrj() {
            if (this.newEditName.length > 1) {
                RenameProject("projects", `${this.selectedProj}${this.selectedProj}`, `${this.newEditName}${this.newEditName}`)
                    .then(response => {
                        console.log("Received data:", response);
                    })
                    .catch(error => {
                        console.error("Error:",error);
                    })
            }
            
        },
        createProject() {
                if (!this.toggleNew) {
                    this.toggleNew = !this.toggleNew;
                } else {

                    const newProjData = {
                        name: `${this.selectedProjMode}${this.newProjName}`,
                    };

                    if (this.selectedProjMode || this.newProjName.length <= 1) {
                        CreateProject(`${this.newProjName}${this.newProjName}`, newProjData)
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
                    } else {
                        console.log("Error creating project");
                    }
                }

            console.log("create project");
        },
        editProject() {
            if (this.selectedProj) {
                this.toggleEdit = !this.toggleEdit;
            } else {
                console.log("Error. Project don't selected")
            }
            
            // coming soon;
        },
        deleteProject() {
            if (this.selectedProj) {
                const filePath = `projects/${this.selectedProj}${this.selectedProj}`;
                DeleteProject(filePath)
                    .then(response => {
                        console.log("Project deleted successfully:", response);
                        this.projectStrokeName = "";
                        this.selectedProjMode = "";
                        this.selectedProj = "";
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
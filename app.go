package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Project struct {
	Name string `json:"name"`
}
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ListProjects(basePath string) ([]string, error) {
	var projects []string
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			var project Project
			if err := json.Unmarshal(file, &project); err != nil {
				return err
			}
			projects = append(projects, project.Name)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("Projects:", projects)
	return projects, nil
}

func (a *App) CreateProject(projectName string, projectData map[string]interface{}) error {
	if len(projectName) < 1 {
		return errors.New("Project name must be at least 1 character long")
	}

	projectPath := filepath.Join("projects", projectName)
	err := os.MkdirAll(filepath.Join(projectPath, "resources"), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(projectPath, "scenes"), os.ModePerm)
	if err != nil {
		return err
	}

	filePath := filepath.Join(projectPath, projectName+".json")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(projectData); err != nil {
		return err
	}

	sceneFilePath := filepath.Join(projectPath, "scenes", "cube.scene")
	sceneFile, err := os.Create(sceneFilePath)
	if err != nil {
		return err
	}
	defer sceneFile.Close()

	content := "coming soon..."
	if _, err := sceneFile.WriteString(content); err != nil {
		return err
	}

	return nil
}

func (a *App) DeleteProject(basePath, projectName string) error {
	filePath := filepath.Join(basePath, projectName)
	err := os.RemoveAll(filePath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) ReadScenes(projectName1 string) ([]string, error) {
	projectName := projectName1 + projectName1
	var scenes []string
	scenesPath := filepath.Join("projects", projectName, "scenes")
	err := filepath.Walk(scenesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			sceneFile, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			scenes = append(scenes, string(sceneFile))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return scenes, nil
}

func (a *App) CreateObject(objectName string, objectData string, projectName string) (string, error) {
	fmt.Println("Creating object...")
	return "", nil
}

func (a *App) RenameProject(basePath, oldProjectName, newProjectName string) error {
	// Rename project folder
	oldPath := filepath.Join(basePath, oldProjectName)
	newPath := filepath.Join(basePath, newProjectName)
	if err := os.Rename(oldPath, newPath); err != nil {
		return err
	}

	// Rename project JSON file
	oldFilePath := filepath.Join(newPath, oldProjectName+".json")
	newFilePath := filepath.Join(newPath, newProjectName+".json")
	if err := os.Rename(oldFilePath, newFilePath); err != nil {
		return err
	}

	return nil
}

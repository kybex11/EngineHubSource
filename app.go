package main

import (
	"context"
	"encoding/json"
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

	sceneFilePath := filepath.Join(projectPath, "scenes", "mainScene.scene")
	sceneFile, err := os.Create(sceneFilePath)
	if err != nil {
		return err
	}
	defer sceneFile.Close()

	return nil
}

func (a *App) DeleteProject(basePath, projectName string) error {
	filePath := filepath.Join(basePath, projectName)
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	dirPath := filepath.Dir(filePath)
	return os.Remove(dirPath)
}

func (a *App) ReadScenes(projectName string) ([]string, error) {
	scenesPath := filepath.Join("projects", projectName, "scenes")
	var scenes []string

	err := filepath.Walk(scenesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			scenes = append(scenes, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return scenes, nil
}

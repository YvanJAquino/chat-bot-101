package env

import "os"


var (
	PRODUCTION string = "PRODUCTION"
	TESTING string = "TESTING"
)

type EnvironmentManager struct {
	ProjectID string
	Port string
	Environment string
} 

func (em *EnvironmentManager) Get() {
	
	project := os.Getenv("PROJECT_ID")
	if project == "" {
		project = "holy-diver-297719"
	}
	em.ProjectID = project

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	em.Port = ":" + port

	env := os.Getenv("ENVIRONMENT")
	switch env {
	case PRODUCTION:
		em.Environment = env
	default:
		em.Environment = TESTING
	}

}

func NewManager() *EnvironmentManager {
	return &EnvironmentManager{}
}

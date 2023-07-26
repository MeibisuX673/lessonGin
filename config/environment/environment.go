package environment

import (
	"github.com/joho/godotenv"
)

var Env *Environment

type Environment struct {
	envMap map[string]string
}

func (e *Environment) Init() error {

	if err := godotenv.Load(".env.local"); err != nil {
		return err
	}

	env, _ := godotenv.Read(".env.local")
	Env = &Environment{envMap: env}

	return nil

}

func (e *Environment) GetEnv(key string) string {

	return e.envMap[key]

}

package cmd

import (
	"log"

	"github.com/MarcelArt/ollama-query/scaffold"
)

func Scaffolder(modelName string) {
	modelCamel := scaffold.ToCamelCase(modelName)
	modelSnake := scaffold.ToSeparateByCharLowered(modelCamel, '_')
	handlerRoute := scaffold.ToSeparateByCharLowered(modelName, '-')
	scaffold.ScaffoldModel(modelName, modelCamel, modelSnake)
	scaffold.ScaffoldRepo(modelName, modelCamel)
	scaffold.ScaffoldHandler(modelName, handlerRoute)
	scaffold.ScaffoldRoute(modelName, handlerRoute)
	log.Println("Successfully scaffolded")
}

/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/Vause/Coffee-Rating-api/cmd"
	"github.com/Vause/Coffee-Rating-api/controllers"
	"github.com/Vause/Coffee-Rating-api/models"
	"github.com/Vause/Coffee-Rating-api/routes"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	controllers.InitLog(sugar)

	cmd.Execute()
	models.ConnectDb(sugar)
	r := routes.SetUpRouter()
	r.Run("localhost:9001")

}

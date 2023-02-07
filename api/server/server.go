package server

import "school-marks-app/api/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}

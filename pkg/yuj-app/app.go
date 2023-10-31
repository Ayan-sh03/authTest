package yujapp

import scripts "yuj/pkg/storage/db-scripts"

func Run() {
	scripts.ConnectDB()
}

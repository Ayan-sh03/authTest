package yujapp

import "yuj/pkg/storage"

func Run() {
	storage.ConnectDB()
}

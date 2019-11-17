package io

import "os"

func Mkdir(path string) error {
	return os.Mkdir(path, 0777)
}

func MkdirAll(path string) error {
	return os.MkdirAll(path, 0777)

}

func Remove(path string) error {
	return os.Remove(path)
}

func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

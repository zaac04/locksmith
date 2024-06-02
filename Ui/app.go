package ui

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	structs "locksmith/Structs"
	"locksmith/crypter"
	"locksmith/file"
	"locksmith/utilities"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetEnvs() (res []structs.Stages, err error) {
	matches, err := utilities.DetectFile()
	if err != nil {
		utilities.LogIfError(err)
		return nil, err
	}

	for _, filename := range matches {
		stageName, found := strings.CutSuffix(filepath.Base(filename), filepath.Ext(filename))
		if !found {
			continue
		}
		header, size := crypter.ReadCipherFile(filename)
		temp := structs.Stages{
			Algorithm:    header.Algorithm,
			LastModified: header.LastModified,
			CipherSize:   size,
			FileName:     filepath.Base(filename),
			StageName:    stageName,
			Id:           filename,
		}
		res = append(res, temp)
	}
	return res, nil
}

func (a *App) IsStagePresent(stage string) (res bool, err error) {
	stage = strings.ToLower(stage)
	matches, err := utilities.DetectFile()
	if err != nil {
		utilities.LogIfError(err)
		return false, err
	}
	for _, v := range matches {
		v = strings.ToLower(v)
		v, _ := strings.CutSuffix(filepath.Base(v), filepath.Ext(v))
		if v == stage {
			return true, nil
		}
	}
	return false, nil
}

func (a *App) ReadCipherFile(stage string, key string) (res string, err error) {
	matches, err := utilities.DetectFile()
	if err != nil {
		utilities.LogIfError(err)
		return
	}

	for _, v := range matches {
		if strings.Contains(v, stage) {
			stage = v
			break
		}
	}

	_, cipherByte, err := crypter.GetDecryptedValue(&stage, &key)
	if err != nil {
		utilities.LogIfError(err)
		return
	}
	return string(cipherByte), nil
}

func (a *App) WriteCipherFile(stage string, key string, val string) (newkey string, err error) {
	var lock crypter.Lock
	var cipherByte []byte

	fmt.Println(stage, key, string(val))
	if key == "" {
		lock, err = crypter.New(256)
		if err != nil {
			utilities.LogIfError(err)
			return
		}
		newkey = lock.GetKey()
		defer file.WriteFile(utilities.GetCipherName(stage)+".key", []byte(newkey))
		cipherByte = []byte{}
	} else {
		lock, err = crypter.LoadKey(key)
		if err != nil {
			utilities.LogIfError(err)
			return
		}
	}

	new_value := []byte(val)

	if key != "" {
		_, cipherByte, err = crypter.GetDecryptedValue(&stage, &key)
		if err != nil {
			utilities.LogIfError(err)
			return
		}
	}
	lock.LoadMessage(new_value)
	if crypter.FindByteDiff(&cipherByte, &new_value) {
		lock.Encrypt(stage)
	} else {
		fmt.Println("No Change Detected!")
	}

	return newkey, nil
}

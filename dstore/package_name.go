/**
 * Copyright (C) 2015 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package dstore

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type DQueryPkgNameTransaction struct {
	data map[string]string
}

// NewDQueryPkgNameTransaction returns package name of given desktop file.
func NewDQueryPkgNameTransaction(path string) (*DQueryPkgNameTransaction, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	t := new(DQueryPkgNameTransaction)
	decoder := json.NewDecoder(bufio.NewReader(f))
	var data map[string]string
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	t.data = convertDesktopPkgMap(data)
	return t, nil
}

func convertDesktopPkgMap(in map[string]string) map[string]string {
	out := make(map[string]string)
	for k, v := range in {
		if !filepath.IsAbs(k) {
			continue
		}

		id := getDesktopId(k)
		if id != "" {
			out[id] = v
		}
	}
	return out
}

var appDirs = []string{
	"/usr/share/applications",
	"/usr/local/share/applications",
}

func getDesktopId(file string) string {
	file = filepath.Clean(file)
	const desktopExt = ".desktop"
	if !strings.HasSuffix(file, desktopExt) {
		return ""
	}
	var desktopId string
	for _, dir := range appDirs {
		if strings.HasPrefix(file, dir) {
			desktopId, _ = filepath.Rel(dir, file)
			break
		}
	}
	if desktopId == "" {
		return ""
	}
	return strings.Replace(desktopId, "/", "-", -1)
}

func (t *DQueryPkgNameTransaction) Query(desktopID string) string {
	if t.data != nil {
		pkg := t.data[desktopID]
		return pkg
	}
	return ""
}

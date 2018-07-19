/*
 * This file is part of arduino-cli.
 *
 * Copyright 2018 ARDUINO AG (http://www.arduino.cc/)
 *
 * arduino-cli is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 */

package output

import (
	"fmt"
	"sort"

	"github.com/bcmi-labs/arduino-cli/arduino/libraries/librariesmanager"
	"github.com/gosuri/uitable"
)

// InstalledLibraries is a list of installed libraries
type InstalledLibraries struct {
	Libraries []*librariesmanager.LibraryAlternatives `json:"libraries"`
}

func (il InstalledLibraries) Len() int { return len(il.Libraries) }
func (il InstalledLibraries) Swap(i, j int) {
	il.Libraries[i], il.Libraries[j] = il.Libraries[j], il.Libraries[i]
}
func (il InstalledLibraries) Less(i, j int) bool {
	return il.Libraries[i].Alternatives[0].String() < il.Libraries[j].Alternatives[0].String()
}

func (il InstalledLibraries) String() string {
	table := uitable.New()
	table.MaxColWidth = 100
	table.Wrap = true

	table.AddRow("Name", "Installed", "Location")
	sort.Sort(il)
	lastName := ""
	for _, lib := range il.Libraries {
		for _, libAlt := range lib.Alternatives {
			name := libAlt.Name
			if name == lastName {
				name = ` "`
			} else {
				lastName = name
			}

			location := libAlt.Location.String()
			if libAlt.ContainerPlatform != nil {
				location = libAlt.ContainerPlatform.String()
			}
			table.AddRow(name, libAlt.Version, location)
		}
	}
	return fmt.Sprintln(table)
}
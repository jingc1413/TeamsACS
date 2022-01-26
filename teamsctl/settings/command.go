/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package settings

import (
	"fmt"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/teamsctl/apiclient"
	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	{
		Name:     "list",
		Usage:    "list settings",
		Category: "Settings",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "type", Aliases: []string{"t"}, Usage: "settings type: system | cwmp | cwmp ", Value: ""},
		},
		Action: func(c *cli.Context) error {
			result, err := apiclient.FindSettings(c.String("type"))
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(common.ToJson(result))
			return nil
		},
	},
	{
		Name:     "update",
		Usage:    "update settings",
		Category: "Settings",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "type", Aliases: []string{"t"}, Usage: "settings type: system | cwmp | cwmp ", Value: ""},
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Usage: "settings name", Value: ""},
			&cli.StringFlag{Name: "value", Aliases: []string{"v"}, Usage: "settings value", Value: ""},
			&cli.StringFlag{Name: "remark", Aliases: []string{"r"}, Usage: "settings remark", Value: ""},
		},
		Action: func(c *cli.Context) error {
			ctype := c.String("type")
			name := c.String("name")
			value := c.String("value")
			remark := c.String("remark")
			if ctype == "" || name == "" || value == "" {
				fmt.Println("type, name, value is required")
				return nil
			}
			result, err := apiclient.UpdateSettings(ctype, name, value, remark)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(common.ToJson(result))
			return nil
		},
	},
}
package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
)

// The output when unmarshalled looks like this:
// [
// {<nil> firewall my-asa-gw FPR9300 FTD 6.6 10.10.10.10   [ssh 22 admin Admin123] [https 443 admin Admin]}
// {<nil> router my-router-calo ISR4431 IOS-XE 6.6 10.10.10.20   [ssh 22 admin Admin123] [https 443 admin Admin]}
// ]

type MyLabFile []struct {
	Device   interface{} `yaml:"device"`
	Type     string      `yaml:"type"`
	Name     string      `yaml:"name"`
	Hardware string      `yaml:"hardware"`
	Software string      `yaml:"software"`
	Version  string      `yaml:"version"`
	IP       string      `yaml:"ip"`
	Notes    string      `yaml:"notes"`
	Cli      []string    `yaml:"cli"`
	Gui      []string    `yaml:"gui"`
}

func parseYAML(filename string) MyLabFile {

	// Read the filename in localhost
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	// Unmarshal the YAML data into a MyLabFile variable
	var lab MyLabFile

	err = yaml.Unmarshal(data, &lab)
	if err != nil {
		fmt.Println(err)
	}

	return lab

}

// labOrganizer organizes the parsed YAML by parseYAML function to be ready for labPrinter function.
func labOrganizer(lab MyLabFile) (majorSlice [][]string) {

	majorSlice = make([][]string, len(lab))

	for index, value := range lab {

		// Iterate through "lab" and obtain a slice of strings, "minorSlice":
		minorSlice := []string{value.Type, value.Name, value.Hardware, value.Software, value.Version, value.IP, value.Notes, value.Cli[0], value.Cli[1], value.Cli[2], value.Cli[3], value.Gui[0], value.Gui[1], value.Gui[2], value.Gui[3]}

		// Then copy the "minorSlice" into the "majorSlice":
		majorSlice[index] = minorSlice

		// Cleaning before next iteration
		minorSlice = nil

	}

	return majorSlice

}

func tableVisualizer(d [][]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"TYPE", "NAME", "HW", "SW", "VERSION", "IP", "NOTES", " CLI-PROTOCOL", "CLI-PORT", "CLI-USER", "CLI-PASSWORD", "GUI-PROTOCOL", "GUI-PORT", "GUI-USER", "GUI-PASSWORD"})
	// table.SetFooter([]string{"", "", "Footer3", "Footer4", "Footer5", "", "", "", "", "", "", "", "", "", ""})
	table.SetBorder(true)

	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
	)

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	)

	for _, row := range d {
		table.Append(row)
	}

	table.SetAutoMergeCells(false)
	table.Render()

}

func main() {

	// Read and parse the YAML file in localhost.
	lab := parseYAML("mylabfile.yaml")

	data := labOrganizer(lab)

	tableVisualizer(data)

}

/* This is how the current execution looks like:


go run yamlParser.go
+----------+--------------------+---------+--------+---------+-------------+------------------+---------------+----------+----------+--------------+--------------+----------+----------+--------------+
|   TYPE   |        NAME        |   HW    |   SW   | VERSION |     IP      |      NOTES       | CLI-PROTOCOL  | CLI-PORT | CLI-USER | CLI-PASSWORD | GUI-PROTOCOL | GUI-PORT | GUI-USER | GUI-PASSWORD |
+----------+--------------------+---------+--------+---------+-------------+------------------+---------------+----------+----------+--------------+--------------+----------+----------+--------------+
| firewall | my-asa-gw          | FPR9300 | FTD    | 6.6     | 10.10.10.10 |  wrong_password  | ssh           | 22       | admin    | Admin123     | https        | 443      | admin    | Admin        |
| firewall | JEAN-ASA           | FPR9300 | FTD    | 6.6     | 10.10.10.10 |  wrong_password  | ssh           | 22       | admin    | Admin123     | https        | 443      | admin    | Admin        |
| router   | my-router-calo     | ISR4431 | IOS-XE | 6.6     | 10.10.10.20 |  needs_a_reimage | ssh           | 22       | admin    | Admin123     | https        | 443      | admin    | Admin        |
| router   | my-router-brussels | ISR4431 | IOS-XE | 6.10    | 10.10.20.20 |  Booting loop    | ssh           | 22       | admin    | Admin123     | https        | 443      | admin    | Admin        |
| firewall | my-ftd             | ISR4432 | IOS-XE | 6.10    | 10.10.20.20 |  Booting loop    | ssh           | 22       | admin    | Admin123     | https        | 443      | admin    | Admin        |
| router   | my-router-krk      | ISR4432 | IOS-XE | 6.10    | 10.10.20.20 |  Booting loop    | ssh           | 22       | admin    | Admin123     | https        | 443      | admin    | Admin        |
+----------+--------------------+---------+--------+---------+-------------+------------------+---------------+----------+----------+--------------+--------------+----------+----------+--------------+


*/

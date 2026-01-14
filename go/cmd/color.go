/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// colorCmd represents the color command
var colorCmd = &cobra.Command{
	Use:   "color",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		recolor()
	},
}

func init() {
	rootCmd.AddCommand(colorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// colorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// colorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type ColorScheme struct {
	Name                string `json:"name"`
	Background          string `json:"background"`
	Black               string `json:"black"`
	Blue                string `json:"blue"`
	BrightBlack         string `json:"brightBlack"`
	BrightBlue          string `json:"brightBlue"`
	BrightCyan          string `json:"brightCyan"`
	BrightGreen         string `json:"brightGreen"`
	BrightPurple        string `json:"brightPurple"`
	BrightRed           string `json:"brightRed"`
	BrightWhite         string `json:"brightWhite"`
	BrightYellow        string `json:"brightYellow"`
	CursorColor         string `json:"cursorColor"`
	Cyan                string `json:"cyan"`
	Foreground          string `json:"foreground"`
	Green               string `json:"green"`
	Purple              string `json:"purple"`
	Red                 string `json:"red"`
	SelectionBackground string `json:"selectionBackground"`
	White               string `json:"white"`
	Yellow              string `json:"yellow"`
}

type TerminalProfile struct {
	AdjustIndistinguishableColors   *string  `json:"adjustIndistinguishableColors,omitempty"`
	BackgroundImage                 *string  `json:"backgroundImage,omitempty"`
	BackgroundImageOpacity          *float64 `json:"backgroundImageOpacity,omitempty"`
	BackgroundImageStretchMode      *string  `json:"backgroundImageStretchMode,omitempty"`
	ColorScheme                     *string  `json:"colorScheme,omitempty"`
	Commandline                     *string  `json:"commandline,omitempty"`
	CursorShape                     *string  `json:"cursorShape,omitempty"`
	Elevate                         *bool    `json:"elevate,omitempty"`
	ExperimentalRetroTerminalEffect *bool    `json:"experimental.retroTerminalEffect,omitempty"`
	Font                            *struct {
		Face *string `json:"face,omitempty"`
	} `json:"font,omitempty"`
	GUID              *string `json:"guid,omitempty"`
	Hidden            *bool   `json:"hidden,omitempty"`
	Icon              *string `json:"icon,omitempty"`
	Name              string  `json:"name,omitempty"`
	Opacity           *int    `json:"opacity,omitempty"`
	Source            *string `json:"source,omitempty"`
	StartingDirectory *string `json:"startingDirectory,omitempty"`
	UseAcrylic        *bool   `json:"useAcrylic,omitempty"`
	IntenseTextStyle  *string `json:"intenseTextStyle,omitempty"`
}

func recolor() {
	colorSchemes := getColorSchemes()
	names := make([]string, 0, len(colorSchemes))
	for i := 0; i < len(colorSchemes); i++ {
		names = append(names, colorSchemes[i].Name)
	}

	prompt := promptui.Select{
		Label: "Select ColorScheme",
		Items: names,
	}
	_, result, err := prompt.Run()

	if err != nil {
		log.Fatalf("themes select failed")
		return
	}

	fmt.Printf("You chose %q\n", result)

	readTerminalSettings()

}

func getColorSchemes() []ColorScheme {
	colorSchemesFilePath := "colorSchemes.json"
	jsonFile, err := os.Open(colorSchemesFilePath)
	if err != nil {
		log.Fatalf("could not read file")
	}

	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("could not read file 2")
	}

	var colorSchemes []ColorScheme

	err = json.Unmarshal(bytes, &colorSchemes)
	if err != nil {
		log.Fatalf("could not unmarshall json")
	}

	return colorSchemes
}

func readTerminalSettings() {
	localAppDataPath := os.Getenv("LOCALAPPDATA")
	if localAppDataPath == "" {
		log.Fatal("Could not read %LOCALAPPDATA%")
	}

	fileLocation := localAppDataPath + "/Packages/Microsoft.WindowsTerminal_8wekyb3d8bbwe/LocalState/settings.json"
	data, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		log.Fatal(err)
		log.Fatalf("Failed to read file")
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalf("Failed to unmarshall json")
	}

	var profileNames []string
	if profiles, ok := result["profiles"].(map[string]interface{}); ok {
		if profileList, ok := profiles["list"].([]interface{}); ok {
			for _, profile := range profileList {
				if profileMap, ok := profile.(map[string]interface{}); ok {
					if name, ok := profileMap["name"].(string); ok {
						profileNames = append(profileNames, name)
					} else {
						fmt.Println("Warning: Profile found without a 'name' field or 'name' is not a string.")
					}
				} else {
					fmt.Println("Warning: Item in 'profiles.list' is not a valid object.")
				}
			}
		} else {
			fmt.Println("Error: 'profiles.list' not found or is not an array.")
		}
	} else {
		fmt.Println("Error: 'profiles' section not found or is not an object.")
	}

	fmt.Println("Found Profile Names:")
	for _, name := range profileNames {
		fmt.Printf("- %s\n", name)
	}
}

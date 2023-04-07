package cmd

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

var cfgFile string
var numberOfHead int
var FileOfInput string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-head",
	Short: "head command by golang",
	Long: `Go言語で開発したheadコマンドです。
    example:
    go-head -n 10 yourfile.txt
    or
    cat yourfile.txt | go-head -n 100`,

	Run: func(cmd *cobra.Command, args []string) {
		optValidate()
		err := goHead()
		if err != nil {
			fmt.Println("[error]", err)
		}
	},
}

func optValidate() {
	// validete for FileOfInput
	stat, err := os.Stat(FileOfInput)
	if err != nil { // 適切なファイル指定なし and 標準出力なし ならエラー終了
		if terminal.IsTerminal(syscall.Stdin) {
			fmt.Println("指定されたファイル名=[", FileOfInput, "] は存在しません。\n",
				"go-head -f yourfile.txt と指定してください\n", stat)
			os.Exit(1)
		}
	}
}

func goHead() error {
	var scanner *bufio.Scanner
	if terminal.IsTerminal(syscall.Stdin) { // 標準入力がなければ、ファイルから
		fp, err := os.Open(FileOfInput)
		if err != nil {
			return err
		}
		scanner = bufio.NewScanner(fp)
	} else { // 標準入力から
		scanner = bufio.NewScanner(os.Stdin)
	}

	counter := 0
	for scanner.Scan() {
		if counter >= numberOfHead {
			return nil // 正常出力
		}
		fmt.Println(scanner.Text())
		counter += 1
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-head.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().IntVarP(&numberOfHead, "numberOfHead", "n", 10, "number of display lines from head")
	rootCmd.Flags().StringVarP(&FileOfInput, "FileOfInput", "f", "", "file of input")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-head" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-head")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

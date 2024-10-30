package cmd

import (
	"fmt"
	"github.com/P8api/cw-cli/cmd/retest"
	"github.com/P8api/cw-cli/cmd/ulxly"
	"os"

	"github.com/P8api/cw-cli/cmd/fork"
	"github.com/P8api/cw-cli/cmd/p2p"
	"github.com/P8api/cw-cli/cmd/parseethwallet"
	"github.com/P8api/cw-cli/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/P8api/cw-cli/cmd/abi"
	"github.com/P8api/cw-cli/cmd/dbbench"
	"github.com/P8api/cw-cli/cmd/dumpblocks"
	"github.com/P8api/cw-cli/cmd/ecrecover"
	"github.com/P8api/cw-cli/cmd/enr"
	"github.com/P8api/cw-cli/cmd/fund"
	"github.com/P8api/cw-cli/cmd/hash"
	"github.com/P8api/cw-cli/cmd/loadtest"
	"github.com/P8api/cw-cli/cmd/metricsToDash"
	"github.com/P8api/cw-cli/cmd/mnemonic"
	"github.com/P8api/cw-cli/cmd/monitor"
	"github.com/P8api/cw-cli/cmd/nodekey"
	"github.com/P8api/cw-cli/cmd/rpcfuzz"
	"github.com/P8api/cw-cli/cmd/signer"
	"github.com/P8api/cw-cli/cmd/version"
	"github.com/P8api/cw-cli/cmd/wallet"
)

var (
	cfgFile   string
	verbosity int
	pretty    bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd *cobra.Command

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd = NewPolycliCommand()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".polygon-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".polygon-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// NewPolycliCommand creates the `polycli` command.
func NewPolycliCommand() *cobra.Command {
	// Parent command to which all subcommands are added.
	cmd := &cobra.Command{
		Use:   "polycli",
		Short: "A Swiss Army knife of blockchain tools.",
		Long:  "Polycli is a collection of tools that are meant to be useful while building, testing, and running block chain applications.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			util.SetLogLevel(verbosity)
			logMode := util.JSON
			if pretty {
				logMode = util.Console
			}
			return util.SetLogMode(logMode)
		},
	}

	// Define flags and configuration settings.
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.polygon-cli.yaml)")
	cmd.PersistentFlags().IntVarP(&verbosity, "verbosity", "v", 500, "0 - Silent\n100 Panic\n200 Fatal\n300 Error\n400 Warning\n500 Info\n600 Debug\n700 Trace")
	cmd.PersistentFlags().BoolVar(&pretty, "pretty-logs", true, "Should logs be in pretty format or JSON")

	// Define local flags which will only run when this action is called directly.
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmd.SetOut(os.Stdout)

	// Define commands.
	cmd.AddCommand(
		abi.ABICmd,
		dumpblocks.DumpblocksCmd,
		ecrecover.EcRecoverCmd,
		fork.ForkCmd,
		fund.FundCmd,
		hash.HashCmd,
		enr.ENRCmd,
		dbbench.DBBenchCmd,
		loadtest.LoadtestCmd,
		metricsToDash.MetricsToDashCmd,
		mnemonic.MnemonicCmd,
		monitor.MonitorCmd,
		nodekey.NodekeyCmd,
		p2p.P2pCmd,
		parseethwallet.ParseETHWalletCmd,
		retest.RetestCmd,
		rpcfuzz.RPCFuzzCmd,
		signer.SignerCmd,
		ulxly.ULxLyCmd,
		version.VersionCmd,
		wallet.WalletCmd,
	)
	return cmd
}

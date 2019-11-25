package cmd

var interval uint
var ntimes uint
var prtVersion bool

func init() {
	rootCmd.PersistentFlags().UintVarP(&interval, "interval", "i", 1000, "interval of blinking led (msec)")
	rootCmd.PersistentFlags().UintVarP(&ntimes, "ntimes", "n", 3, "repeat n-times led blinking")
	rootCmd.PersistentFlags().BoolVarP(&prtVersion, "version", "v", false, "display gobc version")
}



// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/spf13/cobra"
// )

// func ping(domain string)(int, error){
// 	url := "http://" + domain
// 	req, err := http.NewRequest("HEAD", url, nil)
// 	if err != nil {
// 		return 0, err
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return 0, err
// 	}

// 	resp.Body.Close()
// 	return resp.StatusCode, nil
// }
// var(
// 	urlPath string
// 	 client = http.Client{
// 			// Transport : &http.Transport {
// 			// 	Dial: net.Dialer{	Timeout: 2 * time.Second}.Dial,
// 			// },
// 			Timeout: time.Second * 2,
// 		}
// )
// // pingCmd represents the ping command
// //! here ping is a private stuct
// var pingCmd = &cobra.Command{
// 	Use:   "ping",
// 	Short: "This pings a remote URL and return the response",
// 	Long: ``,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		//logic
// 			if resp, err := ping(urlPath); err != nil {
// 				fmt.Println(err)
// 			}else {
// 				fmt.Println(resp)
// 			}
// 		},
// }

// func init() {
	
// 	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping ")
	

// 	// if this floag doesnt exist when the suer runs the command it will throw an error
// 	if err:= pingCmd.MarkFlagRequired("url"); err != nil{
// 		fmt.Println(err)
// 	}
	

// 	NetCmd.AddCommand(pingCmd)
// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//
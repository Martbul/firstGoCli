package network

import (
	"fmt"
	"time"

	// "github.com/schollz/progressbar/v3"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/spf13/cobra"
)

// downloadSpeedCmd represents the downloadSpeed command
var downloadSpeedCmd = &cobra.Command{
	Use:   "sDownload",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {


		//! monitor network interfaces
	//  ifaces, err := net.Interfaces()
   //  if err != nil {
   //      panic(err)
   //  }
   //  for _, iface := range ifaces {
   //      fmt.Printf("Interface Name: %s\n", iface.Name)
   //      fmt.Printf("Interface Addrs: %v\n", iface.Addrs)
   //      fmt.Printf("Interface MTU: %d\n", iface.MTU)
   //  }
	// },





	//! track network trafic
// 	 netIOs, err := net.IOCounters(true)
//     if err != nil {
//         panic(err)
//     }
//     for _, netIO := range netIOs {
//         fmt.Printf("Interface: %s\n", netIO.Name)
//         fmt.Printf("Bytes Sent: %d\n", netIO.BytesSent)
//         fmt.Printf("Bytes Received: %d\n", netIO.BytesRecv)
//         fmt.Printf("Packets Sent: %d\n", netIO.PacketsSent)
//         fmt.Printf("Packets Received: %d\n", netIO.PacketsRecv)
//     }
// },







//! monitor network connections
// conns, err := net.Connections("all")
//     if err != nil {
//         panic(err)
//     }
//     for _, conn := range conns {
//         fmt.Printf("Type: %v\n", conn.Type)
//         fmt.Printf("Local Address: %s:%d\n", conn.Laddr.IP, conn.Laddr.Port)
//         fmt.Printf("Remote Address: %s:%d\n", conn.Raddr.IP, conn.Raddr.Port)
//         fmt.Printf("Status: %s\n", conn.Status)
//     }
// },



//! track download speed
var previousBytesRecv uint64
    interval := 1 * time.Second

    for {
        netIOs, err := net.IOCounters(true)
        if err != nil {
            panic(err)
        }

        var currentBytesRecv uint64
        for _, netIO := range netIOs {
            currentBytesRecv += netIO.BytesRecv
        }

        if previousBytesRecv > 0 {
            bytesPerSec := (currentBytesRecv - previousBytesRecv) / uint64(interval.Seconds())
            fmt.Printf("Download Speed: %d bytes/sec\n", bytesPerSec)
        }

        previousBytesRecv = currentBytesRecv
        time.Sleep(interval)
    }
},
}












//! NOT working download speed tracking with progress bar
//   var previousBytesRecv uint64
//     interval := 1 * time.Second

//     // Create a progress bar with a maximum value (e.g., 100 MB)
//     bar := progressbar.NewOptions(
//         100*1024*1024, // Example max value (100 MB)
//         progressbar.OptionSetDescription("Download Speed"),
//         progressbar.OptionShowBytes(true),
//         progressbar.OptionSetWidth(50), // Adjust width as needed
//         progressbar.OptionSetTheme(progressbar.Theme{
//             Saucer:        "▒",
//             SaucerPadding: "░",
//             BarStart:      "[",
//             BarEnd:        "]",
//         }),
//     )

//     for {
//         netIOs, err := net.IOCounters(true)
//         if err != nil {
//             fmt.Println("Error fetching network statistics:", err)
//             continue
//         }

//         var currentBytesRecv uint64
//         for _, netIO := range netIOs {
//             currentBytesRecv += netIO.BytesRecv
//         }

//         if previousBytesRecv > 0 {
//             bytesPerSec := (currentBytesRecv - previousBytesRecv) / uint64(interval.Seconds())
//             bar.Set(int(currentBytesRecv)) // Update progress bar
//             fmt.Printf("Download Speed: %d bytes/sec\n", bytesPerSec)
//         }

//         previousBytesRecv = currentBytesRecv
//         time.Sleep(interval)
//     }
// },
// }


var uploadSpeedCmd = &cobra.Command{
	Use:   "uploadSpeedCmd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("downloadSpeed called")
	},
}

func init() {
	NetworkCmd.AddCommand(downloadSpeedCmd)
	NetworkCmd.AddCommand(uploadSpeedCmd)
}

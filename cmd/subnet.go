/*
Copyright © 2024 j34sy j34sy@proton.me
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/j34sy/SubnetCalculator/pkg/subnetcalc"
	"github.com/spf13/cobra"
)

// subnetCmd represents the subnet command
var subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "Subnetting a given IPv4 address",
	Long: `This command will show all information about a given IPv4 address with CIDR.
	Example input: 10.0.0.1/24
	
	Information provided:
	- IPv4 address
	- CIDR
	- Subnet mask
	- Network address
	- Broadcast address
	- Usable host range
	- Total hosts
	- Usable hosts
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please provide an IP address with CIDR, e.g. 10.0.0.1/16")
			return
		}
		ipString := args[0]
		if ipString == "" {
			fmt.Println("Please provide an IP address")
			return
		}
		var ip [4]int
		var cidr uint8
		slashSplit := strings.Split(ipString, "/")
		if len(slashSplit) != 2 {
			fmt.Println("Invalid input. Please enter an IP address and CIDR.")
			return
		}
		ipSplit := strings.Split(slashSplit[0], ".")
		if len(ipSplit) != 4 {
			fmt.Println("Invalid input. Please enter a valid IP address.")
			return
		}
		cidrInt, err := strconv.Atoi(slashSplit[1])
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid CIDR.")
			return
		}
		if cidrInt > 0 || cidrInt < 33 {
			cidr = uint8(cidrInt)
		} else {
			fmt.Println("Invalid input. Please enter a valid CIDR.")
			return
		}

		for i, octet := range ipSplit {
			octetInt, err := strconv.Atoi(octet)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid IP address.")
				return
			}
			if octetInt > 0 || octetInt < 256 {
				ip[i] = octetInt
			} else {
				fmt.Println("Invalid input. Please enter a valid IP address.")
				return
			}
		}

		ipv4 := subnetcalc.NewIPv4Address(ip, cidr)

		ipv4.Calculate()

		fmt.Println("IPv4 address: ", ipv4.GetIPv4Address())
		fmt.Println("CIDR: ", ipv4.GetCIDR())
		fmt.Println("Subnet mask: ", ipv4.GetSubnetMask())
		fmt.Println("Network address: ", ipv4.GetNetworkAddress())
		fmt.Println("Broadcast address: ", ipv4.GetBroadcastAddress())
		fmt.Println("Usable host range: ", ipv4.GetUsableHostRange())
		fmt.Println("Total hosts: ", ipv4.GetTotalHosts())
		fmt.Println("Usable hosts: ", ipv4.GetUsableHosts())

	},
}

func init() {
	rootCmd.AddCommand(subnetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subnetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subnetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

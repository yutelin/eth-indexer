// Copyright © 2018 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package grpc

import (
	"fmt"
	"net"

	"github.com/getamis/sirius/log"
	"github.com/maichain/eth-indexer/cmd/flags"
	"github.com/maichain/mapi/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	host string
	port int

	dbDriver   string
	dbHost     string
	dbPort     int
	dbName     string
	dbUser     string
	dbPassword string
)

// ServerCmd represents the grpc-server command
var ServerCmd = &cobra.Command{
	Use:   "grpc",
	Short: "grpc runs a gRPC server for the indexer service",
	Long:  `grpc runs a gRPC server for the indexer service.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			log.Error("Failed to listen", "host", host, "port", port, "err", err)
			return err
		}

		db := MustNewDatabase()
		defer db.Close()

		s := api.NewServer(
		// indexer.New(tx.NewWithDB(db)),
		)

		// go func() {
		// 	sigs := make(chan os.Signal, 1)
		// 	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
		// 	defer signal.Stop(sigs)
		// 	log.Debug("Shutting down", "signal", <-sigs)
		// 	s.Shutdown()
		// }()

		// log.Info("Starting Regulator", "host", host, "port", port)

		if err := s.Serve(l); err != grpc.ErrServerStopped {
			log.Crit("Server stopped unexpectedly", "err", err)
		}

		return nil
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpc-ServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	ServerCmd.Flags().StringVar(&host, flags.Host, "", "The gRPC server listening host")
	ServerCmd.Flags().IntVar(&port, flags.Port, 5487, "The gRPC server listening port")

	// Database flags
	ServerCmd.Flags().StringVar(&dbDriver, flags.DbDriverFlag, "mysql", "The database driver")
	ServerCmd.Flags().StringVar(&dbHost, flags.DbHostFlag, "", "The database host")
	ServerCmd.Flags().IntVar(&dbPort, flags.DbPortFlag, 3306, "The database port")
	ServerCmd.Flags().StringVar(&dbName, flags.DbNameFlag, "eth-db", "The database name")
	ServerCmd.Flags().StringVar(&dbUser, flags.DbUserFlag, "root", "The database username to login")
	ServerCmd.Flags().StringVar(&dbPassword, flags.DbPasswordFlag, "my-secret-pw", "The database password to login")
}

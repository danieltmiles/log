package log

import (
	"fmt"
	"os"
	//"strconv"
	"sync"
	//"time"

	"github.com/kelseyhightower/envconfig"
)

/*
const (
	DEBUG              = 0
	INFO               = 1
	NOTICE             = 2
	WARNING            = 3
	ERROR              = 4
	FATAL              = 5
	defaultBufferLevel = WARNING
)
*/

type Configuration struct {
	//TODO: fill this in
}

var (
	config         Configuration
	tag            string
	mu             sync.Mutex
	hostname       string
	numNetWorkers  = 5
	numFileWorkers = 1
	numConnections = 1
)

/*
func init() {
	hostname, _ := os.Hostname()
}
*/

func configure() {
	tag = os.Args[0]
	err := envconfig.Process(tag, &config)
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
	}

	ok := true

	// Validate not blanks
	// if config.SqsQueues == "" {
	// 	logger.Error("APES_SQSQUEUES not set.")
	// 	ok = false
	// }
	// if config.StorageQueue == "" {
	// 	logger.Error("APES_STORAGEQUEUE not set.")
	// 	ok = false
	// }
	// if config.StorageHost == "" {
	// 	logger.Error("APES_STORAGEHOST not set.")
	// 	ok = false
	// }
	// if config.StoragePort == "" {
	// 	logger.Error("APES_STORAGEPORT not set.")
	// 	ok = false
	// }
	// if config.StorageUser == "" {
	// 	logger.Error("APES_STORAGEUSER not set.")
	// 	ok = false
	// }
	// if config.StoragePass == "" {
	// 	logger.Error("APES_STORAGEPASS not set.")
	// 	ok = false
	// }
	// if config.DistributeQueue == "" {
	// 	logger.Error("APES_DISTRIBUTEQUEUE not set.")
	// 	ok = false
	// }
	// if config.DistributeHost == "" {
	// 	logger.Error("APES_DISTRIBUTEHOST not set.")
	// 	ok = false
	// }
	// if config.DistributePort == "" {
	// 	logger.Error("APES_DISTRIBUTEPORT not set.")
	// 	ok = false
	// }
	// if config.DistributeUser == "" {
	// 	logger.Error("APES_DISTRIBUTEUSER not set.")
	// 	ok = false
	// }
	// if config.DistributePass == "" {
	// 	logger.Error("APES_DISTRIBUTEPASS not set.")
	// 	ok = false
	// }
	// if config.RedisHost == "" {
	// 	logger.Error("APES_REDISHOST not set.")
	// 	ok = false
	// }
	// if config.RedisPort == "" {
	// 	logger.Error("APES_REDISPORT not set.")
	// 	ok = false
	// }
	// if config.AWSAccessKey == "" {
	// 	logger.Error("APES_AWSACCESSKEY not set.")
	// 	ok = false
	// }
	// if config.AWSSecretKey == "" {
	// 	logger.Error("APES_AWSSECRETKEY not set.")
	// 	ok = false
	// }
	// if config.SrsHost == "" {
	// 	logger.Error("APES_SRSHOST not set.")
	// 	ok = false
	// }
	// if config.SrsPort == "" {
	// 	logger.Error("APES_SRSPORT not set.")
	// 	ok = false
	// }

	// if config.Host == "" {
	// 	logger.Warningf("APES_HOST not set, using default value: %v", defaultHost)
	// 	config.Host = defaultHost
	// }

	// if config.Port == "" {
	// 	logger.Warningf("APES_PORT not set, using default value: %v", defaultPort)
	// 	config.Port = defaultPort
	// }

	if !ok {
		fmt.Println("Error configuring application")
		os.Exit(1)
	}

	fmt.Printf("Configuration complete\n")
	fmt.Printf("Config: %v\n", config)
}

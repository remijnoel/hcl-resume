package cmd

import (

    "github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

func config(){

    // Read from environment variables (override previous)
    viper.SetEnvPrefix("HCL_RESUME")
	log.Debugf("Loading environment variables with prefix %s", viper.GetEnvPrefix())
    viper.AutomaticEnv() // Maps env vars to keys
}
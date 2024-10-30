package actions

import (
	"encoding/json"
	"log"

	"github.com/HemeraProtocol/avs/legacy/core/config"
	"github.com/HemeraProtocol/avs/legacy/operator"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/urfave/cli"
)

func DeregisterOperatorWithAvs(ctx *cli.Context) error {
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig := config.NodeConfig{}

	if configPath != "" {
		err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
		if err != nil {
			return err
		}
		configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Println("Config:", string(configJson))
	}

	operator, err := operator.NewOperatorFromConfig(nodeConfig, true)
	if err != nil {
		return err
	}

	err = operator.DeregisterOperatorWithAvs()
	if err != nil {
		return err
	}

	return nil
}

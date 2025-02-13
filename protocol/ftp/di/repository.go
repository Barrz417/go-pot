package di

import (
	"github.com/ryanolee/go-pot/config"
	"github.com/ryanolee/go-pot/core/stall"
	"github.com/ryanolee/go-pot/generator"
	"github.com/ryanolee/go-pot/protocol/ftp/logging"
	ftpStall "github.com/ryanolee/go-pot/protocol/ftp/stall"
	"github.com/ryanolee/go-pot/protocol/ftp/throttle"
	"github.com/ryanolee/go-pot/secrets"
)

// Repository used for passing Dependencies required by
// client elements of the FTP protocol
type (
	FtpRepository struct {
		config           *config.Config
		configGenerators *generator.ConfigGeneratorCollection
		secretGenerators *secrets.SecretGeneratorCollection
		throttle         *throttle.FtpThrottle
		stallPool        *stall.StallerPool
		ftpStallFactory  *ftpStall.FtpFileStallerFactory
		logger           *logging.FtpCommandLogger
	}
)

func NewFtpRepository(
	config *config.Config,
	configGenerators *generator.ConfigGeneratorCollection,
	secretGenerators *secrets.SecretGeneratorCollection,
	throttle *throttle.FtpThrottle,
	stallPool *stall.StallerPool,
	ftpStallFactory *ftpStall.FtpFileStallerFactory,
	ftpCommandLogger *logging.FtpCommandLogger,
) *FtpRepository {
	return &FtpRepository{
		config:           config,
		configGenerators: configGenerators,
		secretGenerators: secretGenerators,
		throttle:         throttle,
		stallPool:        stallPool,
		ftpStallFactory:  ftpStallFactory,
		logger:           ftpCommandLogger,
	}
}

func (r *FtpRepository) GetConfigGenerators() *generator.ConfigGeneratorCollection {
	return r.configGenerators
}

func (r *FtpRepository) GetSecretGenerators() *secrets.SecretGeneratorCollection {
	return r.secretGenerators
}

func (r *FtpRepository) GetThrottle() *throttle.FtpThrottle {
	return r.throttle
}

func (r *FtpRepository) GetStallPool() *stall.StallerPool {
	return r.stallPool
}

func (r *FtpRepository) GetFtpStallFactory() *ftpStall.FtpFileStallerFactory {
	return r.ftpStallFactory
}

func (r *FtpRepository) GetConfig() *config.Config {
	return r.config
}

func (r *FtpRepository) GetLogger() *logging.FtpCommandLogger {
	return r.logger
}

package worker

import (
	"context"
	"fmt"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/mail"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func Start(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config *utils.Config,
	redisOpt asynq.RedisClientOpt,
	gojo db.Gojo,
) {
	taskProcessor := NewRedisTaskProcessor(redisOpt, gojo, mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword))

	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start task processor")
	}

	waitGroup.Go(func() error {
		<-ctx.Done()
		fmt.Printf("\u001b[38;5;266mSHUTDOWN QUEUE SERVER ...\u001b[0m\n")

		taskProcessor.Shutdown()
		fmt.Printf("\u001b[38;5;196mQUEUE SERVER IS STOPPED.\u001b[0m\n")

		return nil
	})
}

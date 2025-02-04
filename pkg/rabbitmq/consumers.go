package rabbitmq

import (
	"github.com/nholuongut/hashbash-backend-go/pkg/dao"
	"github.com/nholuongut/hashbash-backend-go/pkg/rainbow"
	"github.com/nholuongut/rabbitmq-client-go/rabbitmq"
)

type HashbashMqConsumerWorkers struct {
	HashbashDeleteRainbowTableConsumer   *rabbitmq.Consumer
	HashbashGenerateRainbowTableConsumer *rabbitmq.Consumer
	HashbashSearchRainbowTableConsumer   *rabbitmq.Consumer
}

func CreateConsumerWorkers(
	connection *rabbitmq.ServerConnection,
	rainbowTableService dao.RainbowTableService,
	rainbowTableGenerateJobService *rainbow.TableGeneratorJobService,
	rainbowTableSearchJobService *rainbow.TableSearchJobService,
) (HashbashMqConsumerWorkers, error) {
	deleteRainbowTableConsumer, err0 := NewDeleteRainbowTableConsumer(connection, rainbowTableService)
	generateRainbowTableConsumer, err1 := NewGenerateRainbowTableConsumer(connection, rainbowTableService, rainbowTableGenerateJobService)
	searchRainbowTableConsumer, err2 := NewSearchRainbowTableConsumer(connection, rainbowTableSearchJobService)

	for _, e := range []error{err0, err1, err2} {
		if e != nil {
			return HashbashMqConsumerWorkers{}, e
		}
	}

	return HashbashMqConsumerWorkers{
		HashbashDeleteRainbowTableConsumer:   deleteRainbowTableConsumer,
		HashbashGenerateRainbowTableConsumer: generateRainbowTableConsumer,
		HashbashSearchRainbowTableConsumer:   searchRainbowTableConsumer,
	}, nil
}

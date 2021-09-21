package storage

import (
	"Tsystem/types"
	"context"
)

type Store struct {
	ctx    context.Context
	client *Client
}

func NewStore(ctx context.Context, client *Client) Store {
	return Store{
		ctx:    ctx,
		client: client,
	}

}

func (store Store) Create(data *types.Sea) error {
	sea := types.Sea{
		SenderName:       data.SenderName,
		Receiver_Name:    data.Receiver_Name,
		SenderLocation:   data.SenderLocation,
		ReceiverLocation: data.ReceiverLocation,
		Tracker:          data.Tracker,
		Bill:             data.Bill,
	}
	return store.client.db.Debug().Create(&sea).Error
}
func (store Store) Seas() ([]*types.Sea, error) {
	var sea []*types.Sea

	return sea, store.client.db.Debug().Find(&sea).Error
}

func (store Store) Sea(tracker string) (*types.Sea, error) {
	var res types.Sea

	return &res, store.client.db.Debug().Where("tracker = ?", tracker).Find(&res).Error
}

func (store Store) Update(data types.Sea) error {

	return store.client.db.Debug().Model(&types.Sea{}).Where("tracker = ?", data.Tracker).Update("receiver_location", data.ReceiverLocation).Error
}

func (store Store) Delete(data types.Sea) error {
	return store.client.db.Debug().Where("tracker = ?", data.Tracker).Unscoped().Delete(data).Error
}

func (store Store) SenderNames() ([]types.Sea, error) {
	var data []types.Sea
	return data, store.client.db.Debug().Select("sender_name").Find(&data).Error
}

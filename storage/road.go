package storage

import (
	"Tsystem/types"
	"context"
)

type DataStore struct {
	client *Client
	ctx    context.Context
}

func NewDataStore(ctx context.Context, c *Client) DataStore {
	return DataStore{
		client: c,
		ctx:    ctx,
	}
}

func (store DataStore) Create(data *types.Road) error {

	road := types.Road{
		SenderName:       data.SenderName,
		ReceiverName:     data.ReceiverName,
		SenderLocation:   data.SenderLocation,
		ReceiverLocation: data.ReceiverLocation,
		Tracker:          data.Tracker,
		Bill:             data.Bill,
	}

	return store.client.db.Debug().Create(&road).Error
}

func (store DataStore) Roads() ([]*types.Road, error) {
	var res []*types.Road
	return res, store.client.db.Debug().Find(&res).Error

}

func (store DataStore) Road(tracker string) (*types.Road, error) {

	var track types.Road
	return &track, store.client.db.Debug().Where("tracker = ?", tracker).Find(&track).Error
}

func (store DataStore) UpdateLocation(data types.Road) error {

	return store.client.db.Debug().Model(&types.Road{}).Where("tracker = ?", data.Tracker).Update("receiver_location", data.ReceiverLocation).Error

}

func (store DataStore) Delete(tracker string, data types.Road) error {
	return store.client.db.Debug().Where("tracker = ?", tracker).Unscoped().Delete(&data).Error
}

func (store DataStore) SenderNames() ([]types.Road, error) {
	var data []types.Road
	return data, store.client.db.Debug().Select("sender_name").Find(&data).Error
}

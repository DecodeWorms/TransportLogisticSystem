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

func (store DataStore) Create(data types.Road) error {
	road := types.Road{
		SenderName:       data.SenderName,
		ReceiverName:     data.ReceiverName,
		SenderLocation:   data.SenderLocation,
		ReceiverLocation: data.ReceiverLocation,
		Tracker:          data.Tracker,
		Bill:             data.Bill,
	}
	return store.client.db.Create(&road).Error
}

func (store DataStore) Roads() ([]*types.Road, error) {
	var res []*types.Road

	return res, store.client.db.Find(&res).Error

}

func (store DataStore) Road(tracker string) (types.Road, error) {
	var data types.Road

	// return data, store.client.db.Where("tracker = ?", tracker).Find(&data).Error
	return data, store.client.db.First(&data, "tracker = ?", tracker).Error
}

func (store DataStore) UpdateLocation(data types.Road) error {

	return store.client.db.Model(&types.Road{}).Where("tracker =?", data.Tracker).Update("receiver_location", data.ReceiverLocation).Error
}

func (store DataStore) GetSenderAndReceiverName(tracker string) (types.Road, error) {
	var data types.Road

	return data, store.client.db.Select("sender_name", "receiver_name").Where("tracker = ?", tracker).Find(&data).Error
}

func (store DataStore) Delete(tracker string, data types.Road) error {
	return store.client.db.Where("tracker = ?", tracker).Unscoped().Delete(&data).Error
}

func (store DataStore) SenderNames() ([]*types.Road, error) {
	var data []*types.Road

	return data, store.client.db.Select("sender_name").Find(&data).Error
}

func (store DataStore) GetLocation() ([]types.Road, error) {
	var data []types.Road

	return data, store.client.db.Select("tracker", "sender_location").Find(&data).Error
}

func (store DataStore) GetId(tracker string) (types.Road, error) {
	var data types.Road
	return data, store.client.db.First(&data, "tracker = ?", tracker).Error
}

func (store DataStore) UpdateSenderLocation(tracker string, data types.Road) error {

	return store.client.db.Model(&types.Road{}).Where("tracker = ?", tracker).Update("sender_location", data.SenderLocation).Error

}

func (store DataStore) CreateIdentifiedField(data types.Road) error {

	road := types.Road{
		SenderName: data.SenderName,
		Tracker:    data.Tracker,
		Bill:       data.Bill,
	}

	return store.client.db.Select("sender_name", "tracker", "bill").Create(&road).Error
}

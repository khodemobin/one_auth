package repository_test

// var d = &domain.Sample{}

// func TestFindDevice(t *testing.T) {
// 	db, mock := test_mock.NewMock()

// 	repo := repository.NewDeviceRepo(db)

// 	query := "^SELECT (.+) FROM devices*"

// 	rows := sqlmock.NewRows([]string{"id", "imei", "mac_address"}).
// 		AddRow(d.ID, d.Imei, d.Mac)

// 	mock.ExpectQuery(query).WithArgs(d.Imei, d.Mac).WillReturnRows(rows)

// 	deviceId, err := repo.FindByImei(context.TODO(), d.Imei, d.Mac)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, deviceId)
// }

// func TestFindDeviceNotFound(t *testing.T) {
// 	db, mock := test_mock.NewMock()

// 	repo := repository.NewDeviceRepo(db)

// 	query := "^SELECT (.+) FROM devices*"

// 	rows := sqlmock.NewRows([]string{"id", "imei", "mac_address"}).
// 		AddRow(d.ID, d.Imei, d.Mac)

// 	mock.ExpectQuery(query).WithArgs(d.Imei, d.Mac).WillReturnRows(rows)

// 	deviceId, err := repo.FindByImei(context.TODO(), d.Imei+"1", d.Mac)
// 	assert.Empty(t, deviceId)
// 	assert.Error(t, err)
// }

// func TestFindDeviceError(t *testing.T) {
// 	db, mock := test_mock.NewMock()

// 	repo := repository.NewDeviceRepo(db)

// 	query := "^SELECT (.+) FROM devices*"

// 	rows := sqlmock.NewRows([]string{"id", "imeii", "mac_address"}).
// 		AddRow(d.ID, d.Imei, d.Mac)

// 	mock.ExpectQuery(query).WithArgs(d.Imei, d.Mac).WillReturnRows(rows)

// 	deviceId, err := repo.FindByImei(context.TODO(), d.Imei, d.Mac)
// 	assert.Empty(t, deviceId)
// 	assert.Error(t, err)
// }

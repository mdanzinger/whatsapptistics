package report

import (
	"context"
)

type reportServer struct {
	store    ReportStore    // ex: db transactions
	notifier ReportNotifier // ex: sns notification
	storage  ReportStore    // ex: s3
}

func (rs *reportServer) Upload(ctx context.Context, r *Report) error {
	if err := rs.storage.Store(ctx, r); err != nil {
		return err
	}
	return nil
}

func (rs *reportServer) Notify(ctx context.Context, r *Report) error {
	if err := rs.notifier.Notify(ctx, r); err != nil {
		return err
	}
	return nil
}

func (rs *reportServer) Get(ctx context.Context, key string) (*Report, error) {
	r, err := rs.store.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func NewReportServer(store ReportStore, notifier ReportNotifier, storage ReportStore) *reportServer {
	return &reportServer{
		store:    store,
		notifier: notifier,
		storage:  storage,
	}
}

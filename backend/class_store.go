package main

import "context"

type ClassStore interface {
	List(ctx context.Context) ([]BoatClass, error)
	Add(ctx context.Context, bc BoatClass) error
	DeleteByName(ctx context.Context, name string) (BoatClass, bool, error)
}

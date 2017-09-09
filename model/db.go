package model

import "github.com/boltdb/bolt"

type DBface interface {
	Update(func(tx *bolt.Tx) error) error
	View(func(tx *bolt.Tx) error) error
}

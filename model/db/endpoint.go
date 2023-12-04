package db

import (
	"gorm.io/gorm"
	"sync"
)

type Endpoint func(tx *gorm.DB) error

func EndpointSync(tx *gorm.DB, fns ...Endpoint) error {
	cn := len(fns)
	var wg sync.WaitGroup
	errch := make(chan error, cn)
	wg.Add(cn)
	for _, fn := range fns {
		go func(fn Endpoint) {
			defer wg.Done()
			if err := fn(tx); err != nil {
				errch <- err
			}
			return
		}(fn)
	}
	wg.Wait()

	select {
	case err := <-errch:
		return err
	default:
		return nil
	}
}

func Exec(tx *gorm.DB, eps ...Endpoint) error {
	for _, ep := range eps {
		if err := ep(tx); err != nil {
			return err
		}
	}

	return nil
}

func ExecSync(tx *gorm.DB, eps ...Endpoint) error {
	if err := EndpointSync(tx, eps...); err != nil {
		return err
	}

	return nil
}

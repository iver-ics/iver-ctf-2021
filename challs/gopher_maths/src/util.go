package main

import (
	"bufio"
	crand "crypto/rand"
	"math/rand"
	"strings"

	"git.mills.io/prologic/go-gopher"
)

func cryptSeedRand() error {
	var seedBytes [8]byte
	if _, err := crand.Read(seedBytes[:]); err != nil {
		return err
	}
	var seed int64
	for _, byte := range seedBytes {
		seed <<= 8
		seed += int64(byte)
	}
	rand.Seed(seed)
	return nil
}

func writeInfoMultiline(w gopher.ResponseWriter, s string) error {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		if err := w.WriteInfo(scanner.Text()); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func writeDirectoryLink(w gopher.ResponseWriter, cfg Config, selector, description string) {
	w.WriteItem(&gopher.Item{
		Type:        gopher.DIRECTORY,
		Selector:    selector,
		Description: description,
		Host:        cfg.PublicHost,
		Port:        cfg.PublicPort,
	})
}

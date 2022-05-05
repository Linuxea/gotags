//go:build enter && pro
// +build enter,pro

package main

func init() {
	features = append(features, "enterpriseFeature#1", "enterpriseFeature#2")
}

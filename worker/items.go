// SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
//
// SPDX-License-Identifier: MIT
package main

import "fmt"

type ItemService struct{}

type SaleAction struct {
	ID          string  `json:"itemID"`
	NewPrice    float32 `json:"newPrice"`
	SalePercent int     `json:"salePercent"`
}

func (i ItemService) ApplySale(action SaleAction) {
	fmt.Printf("sale applied! item: %s, price: %f (-%d%% sale)",
		action.ID, action.NewPrice, action.SalePercent)
}

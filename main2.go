// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    test, err := UnmarshalTest(bytes)
//    bytes, err = test.Marshal()

package main

import "encoding/json"

func UnmarshalTest(data []byte) (Test, error) {
	var r Test
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Test) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Test struct {
	Header [][]string `json:"header"`
	Data   []Datum    `json:"data"`
}

type Datum struct {
	ID              int64  `json:"Id"`
	CreateAt        string `json:"CreateAt"`
	UserID          int64  `json:"UserId"`
	Gold            int64  `json:"Gold"`
	GoldDelta       int64  `json:"GoldDelta"`
	Type            int64  `json:"Type"`
	Remark          Remark `json:"Remark"`
	SenderID        int64  `json:"SenderId"`
	GiftID          int64  `json:"GiftId"`
	SendNum         int64  `json:"SendNum"`
	MaskInfo        string `json:"MaskInfo"`
	RoomID          int64  `json:"room_id"`
	WelfareReportID int64  `json:"welfare_report_id"`
	RiskID          int64  `json:"risk_id"`
	RiskStatus      int64  `json:"risk_status"`
	CostGoldTicket  int64  `json:"cost_gold_ticket"`
	CostNobleGold   int64  `json:"cost_noble_gold"`
	CostChargeGold  int64  `json:"cost_charge_gold"`
	PkRewardDeduct  int64  `json:"pk_reward_deduct"`
	PkReward        int64  `json:"pk_reward"`
	OpenFrom        string `json:"open_from"`
	App             App    `json:"app"`
}

type App string

const (
	Oxygen App = "oxygen"
)

type Remark string

const (
	一起玩结算 Remark = "一起玩结算"
)

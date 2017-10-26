package main

import "testing"

var lotteryKeys = []struct {
	Key string
	Valid bool
}{
	{
		Key: "powerball",
		Valid: true,
	},
	{
		Key: "mikes_made_up_lottery_ticket",
		Valid: false,
	},
}

var raffleKeys = []struct {
	Key string
	Valid bool
}{
	{
		Key: "mater_prize_home",
		Valid: true,
	},
	{
		Key: "mikes_made_up_raffle",
		Valid: false,
	},
}

func TestLotteryKey(t *testing.T) {
	fillCache()
	for _, i := range lotteryKeys {
		result := jsonCache.getLotteryKey(i.Key)
		if result == nil && i.Valid {
			t.Errorf("Unable to find lottery key %s in cache", i.Key)
		}

		if result != nil && !i.Valid {
			t.Errorf("Able to find invalid lottery key %s in cache", i.Key)
		}
	}
}

func TestRaffleKey(t *testing.T) {
	fillCache()
	for _, i := range raffleKeys {
		result := jsonCache.getRaffleKey(i.Key)
		if result == nil && i.Valid {
			t.Errorf("Unable to find raffle key %s in cache", i.Key)
		}

		if result != nil && !i.Valid {
			t.Errorf("Able to find invalid raffle key %s in cache", i.Key)
		}
	}
}

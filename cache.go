package main

import (
	"io/ioutil"
	"encoding/json"
)

// auto-generated with https://mholt.github.io/json-to-go/
//
// I utilised this online tool to help get a struct
// in place I can work with now, allowing me to produce
// at minimum an MVP. This can be broken into pieces further
// down the line and/or pushed into a cache or database
type PoorMansCache struct {
	Result []struct {
		Type         string `json:"type"`
		Key          string `json:"key"`
		Name         string `json:"name"`
		Autoplayable string `json:"autoplayable"`
		GameTypes    []struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Description string `json:"description"`
			GameOffers  []struct {
				Key         string `json:"key"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Price       struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"price"`
				MinGames      int  `json:"min_games"`
				MaxGames      int  `json:"max_games"`
				Multiple      int  `json:"multiple"`
				Ordered       bool `json:"ordered"`
				GameIncrement struct {
					Num4 int `json:"4"`
				} `json:"game_increment"`
				EquivalentGames int `json:"equivalent_games"`
				NumberSets      []struct {
					First int `json:"first"`
					Last  int `json:"last"`
					Sets  []struct {
						Name  string `json:"name"`
						Count int    `json:"count"`
					} `json:"sets"`
				} `json:"number_sets"`
				Combinations []interface{} `json:"combinations"`
				DisplayRange interface{}   `json:"display_range"`
			} `json:"game_offers"`
		} `json:"game_types,omitempty"`
		Draws []struct {
			Name      interface{} `json:"name"`
			Date      string      `json:"date"`
			Stop      string      `json:"stop"`
			DrawNo    int         `json:"draw_no"`
			PrizePool struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"prize_pool"`
			JackpotImage struct {
				ImageName          string `json:"image_name"`
				ImageURL           string `json:"image_url"`
				SvgURL             string `json:"svg_url"`
				ImageWidth         int    `json:"image_width"`
				ImageHeight        int    `json:"image_height"`
				ContentDescription string `json:"content_description"`
			} `json:"jackpot_image"`
		} `json:"draws,omitempty"`
		Days []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"days,omitempty"`
		Addons         []interface{} `json:"addons,omitempty"`
		QuickpickSizes []int         `json:"quickpick_sizes,omitempty"`
		Lottery        struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			Desc         string `json:"desc"`
			Multidraw    bool   `json:"multidraw"`
			Type         string `json:"type"`
			IconURL      string `json:"icon_url"`
			IconWhiteURL string `json:"icon_white_url"`
			PlayURL      string `json:"play_url"`
			LotteryID    int    `json:"lottery_id"`
		} `json:"lottery,omitempty"`
		Draw struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			DrawNumber  int    `json:"draw_number"`
			DrawStop    string `json:"draw_stop"`
			DrawDate    string `json:"draw_date"`
			Prize       struct {
				Type        string `json:"type"`
				CardTitle   string `json:"card_title"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Content     struct {
					SalesPitchHeading1    string   `json:"sales_pitch_heading_1"`
					SalesPitchSubHeading1 string   `json:"sales_pitch_sub_heading_1"`
					Paragraph1            string   `json:"paragraph_1"`
					Paragraph2            string   `json:"paragraph_2"`
					Paragraph3            string   `json:"paragraph_3"`
					Image                 string   `json:"image"`
					SalesPitchHeading2    string   `json:"sales_pitch_heading_2"`
					SalesPitchSubHeading2 string   `json:"sales_pitch_sub_heading_2"`
					Features              []string `json:"features"`
				} `json:"content"`
				Value struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"value"`
				ValueIsExact     bool        `json:"value_is_exact"`
				HeroImage        interface{} `json:"hero_image"`
				CarouselImages   []string    `json:"carousel_images"`
				FeatureDrawImage interface{} `json:"feature_draw_image"`
				EdmImage         string      `json:"edm_image"`
			} `json:"prize"`
			Offers []struct {
				Name       string `json:"name"`
				Key        string `json:"key"`
				NumTickets int    `json:"num_tickets"`
				Price      struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"price"`
				PricePerTicket struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"price_per_ticket"`
				Ribbon     string      `json:"ribbon"`
				BonusPrize interface{} `json:"bonus_prize"`
			} `json:"offers"`
			TermsAndConditionsURL string `json:"terms_and_conditions_url"`
		} `json:"draw,omitempty"`
	} `json:"result"`
	Messages []interface{} `json:"messages"`
}

func (c *PoorMansCache) getLotteryKey(key string) interface{} {
	for _, obj := range c.Result {
		if obj.Type == "lottery_ticket" && obj.Key == key {
			return obj
		}
	}

	return nil
}

func (c *PoorMansCache) getRaffleKey(key string) interface{} {
	for _, obj := range c.Result {
		if obj.Type == "raffle_ticket" && obj.Key == key {
			return obj
		}
	}

	return nil
}

// Allows me to cache the JSON document in a native
// format, allowing for much faster access and also
// passing through to the Go templating engine.
func fillCache() (err error) {
	fileContents, _ := ioutil.ReadFile("./json/response.json")
	err = json.Unmarshal(fileContents, jsonCache)
	if err != nil {
		return err
	}

	return nil
}

package lib

import (
	"time"
)

// FortuneResult は、おみくじを引いた結果
// - JSON返却値
type FortuneResult struct {
	Today         jsonTime `json:"today"`
	Fortune       string   `json:"fortune"`
	IsNewyearMode bool     `json:"is_newyear_mode"`
}

// 独自のjsonTimeを作成
// - https://qiita.com/taizo/items/2c3a338f1aeea86ce9e2
type jsonTime struct {
	time.Time
}

func (j jsonTime) format() string {
	return j.Time.Format("2006-01-02")
}
func (j jsonTime) parse(value string) (time.Time, error) {
	return time.Parse(`"2006-01-02"`, value)
}

// MarshalJSON() の実装
func (j jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + j.format() + `"`), nil
}

// UnmarshalJSON() の実装
func (j jsonTime) UnmarshalJSON(data []byte) error {
	tm, err := j.parse(string(data))
	j.Time = tm
	return err
}

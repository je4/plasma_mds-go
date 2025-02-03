package plasmaMDS

import (
	"emperror.dev/errors"
	"encoding/json"
	"strings"
)

type Quality string

const (
	QualityVerified  Quality = "verified"
	QualityPublished Quality = "published"
	QualityReviewed  Quality = "reviewed"
)

func (q Quality) MarshalJSON() ([]byte, error) {
	switch q {
	case QualityVerified, QualityPublished, QualityReviewed:
		return json.Marshal(string(q))
	default:
		return nil, errors.New("invalid quality value")
	}
}

func (q *Quality) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	s = strings.ToLower(s)
	switch Quality(s) {
	case QualityPublished, QualityReviewed, QualityVerified:
		*q = Quality(s)
	default:
		return errors.New("invalid quality value")
	}

	return nil
}

package grodata

import (
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/araddon/dateparse"
	"time"
)

type CustomDate struct {
	time.Time
}

const customDateFormat = "2006-01-02"

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return errors.Wrapf(err, "error unmarshalling date '%s'", string(b))
	}
	if str == "" {
		return nil
	}
	t, err := dateparse.ParseStrict(str)
	//t, err := time.Parse(customDateFormat, str)
	if err != nil {
		return fmt.Errorf("error parsing date: %w", err)
	}
	cd.Time = t
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(cd.Format(customDateFormat))
}

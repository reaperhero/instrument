package paginate

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	lis := []string{"1", "2", "a", "5"}
	count, err := Paginate(&lis, 1, 2)
	if err != nil {
		logrus.Errorf("%v",err)
	}

	assert.Equal(t, lis, []string{"1", "2"}, "")
	assert.Equal(t, count, 4, "")
}

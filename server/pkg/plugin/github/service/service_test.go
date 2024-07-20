package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatternReplaceAll(t *testing.T) {
	url := "http://localhost:3000/api{/user}"

	replaced := patternReplaceAll(url, map[string]string{"/user": ""})
	assert.Equal(t, "http://localhost:3000/api", replaced)

	replaced = patternReplaceAll(url, map[string]string{"/user": "/thenewboston"})
	assert.Equal(t, "http://localhost:3000/api/thenewboston", replaced)
}

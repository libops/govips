package vips

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadImageFromFileDirect(t *testing.T) {
	require.NoError(t, Startup(&Config{}))

	img, err := LoadImageFromFileDirect(resources+"tif.tif", nil)
	require.NoError(t, err)
	defer img.Close()

	assert.Greater(t, img.Width(), 0)
	assert.Greater(t, img.Height(), 0)
	assert.Equal(t, ImageTypeTIFF, img.Format())
	assert.Equal(t, ImageTypeTIFF, img.OriginalFormat())
	assert.Nil(t, img.buf)
}

func TestLoadImageFromFileDirectCopy(t *testing.T) {
	require.NoError(t, Startup(&Config{}))

	img, err := LoadImageFromFileDirect(resources+"tif.tif", nil)
	require.NoError(t, err)
	defer img.Close()

	copy, err := img.Copy()
	require.NoError(t, err)
	defer copy.Close()

	assert.Equal(t, img.Width(), copy.Width())
	assert.Equal(t, img.Height(), copy.Height())
	assert.Equal(t, img.Format(), copy.Format())
	assert.Equal(t, img.OriginalFormat(), copy.OriginalFormat())
	assert.Nil(t, copy.buf)
}

func TestLoadImageFromFileDirectRejectsNULFilename(t *testing.T) {
	require.NoError(t, Startup(&Config{}))

	img, err := LoadImageFromFileDirect(resources+"tif.tif\x00ignored", nil)
	require.Error(t, err)
	assert.Nil(t, img)
	assert.Contains(t, err.Error(), "filename contains NUL")
}

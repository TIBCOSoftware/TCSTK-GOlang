/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package urlshortener

import (
	"testing"

	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {

	settings := &Settings{
		Provider:      "Bitly",
		SecurityToken: "",
	}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	input := &Input{
		LongDynamicLink: "https://www.elmconfi.com/home/79c856a98d3e138add66239e3267184e151a39705178a88d9374b5a55b5a900039cbae9e86647be4cc5286102a13a626715703038f2fb94dd85e5bb44e12afa7fb7845e5eae5218ec8b630d73b5138cb992b10fe0c8d6a6b5a46b0f710bbe970cdea8b1695227594764750ad74004e16",
	}
	tc.SetInputObject(input)

	act.Eval(tc)

	output := &Output{}
	tc.GetOutputObject(output)
	assert.NotEmpty(t, output.ShortLink)
}

func TestEvalRebrandly(t *testing.T) {

	settings := &Settings{
		Provider:      "Rebrandly",
		SecurityToken: "",
	}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	input := &Input{
		LongDynamicLink: "https://www.elmconfi.com/home/79c856a98d3e138add66239e3267184e151a39705178a88d9374b5a55b5a900039cbae9e86647be4cc5286102a13a626715703038f2fb94dd85e5bb44e12afa7fb7845e5eae5218ec8b630d73b5138cb992b10fe0c8d6a6b5a46b0f710bbe970cdea8b1695227594764750ad74004e16",
	}
	tc.SetInputObject(input)

	act.Eval(tc)

	output := &Output{}
	tc.GetOutputObject(output)
	assert.NotEmpty(t, output.ShortLink)
}

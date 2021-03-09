/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package getcontext

import (
	"testing"

	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestEvalToken(t *testing.T) {

	settings := &Settings{}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	input := &Input{
		Region:      "US",
		AccessToken: "CIC~Lr4bJC8kYeF3UjMwMeeCNR2c",
	}
	tc.SetInputObject(input)

	act.Eval(tc)

	output := &Output{}
	tc.GetOutputObject(output)
	assert.NotEmpty(t, output.Context1)
}

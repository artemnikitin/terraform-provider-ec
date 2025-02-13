// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package acc

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDeployment_emptyconfig(t *testing.T) {
	resName := "ec_deployment.emptyconfig"
	randomName := prefix + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	startCfg := "testdata/deployment_emptyconfig.tf"

	cfgF := func(cfg string) string {
		return fixtureAccDeploymentResourceBasic(
			t, cfg, randomName, getRegion(), defaultTemplate,
		)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactory,
		CheckDestroy:      testAccDeploymentDestroy,
		Steps: []resource.TestStep{
			{
				Config: cfgF(startCfg),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "elasticsearch.#", "1"),
					resource.TestCheckResourceAttr(resName, "elasticsearch.0.config.#", "0"),
				),
				// Since the configuration specifies a `config {}` block but
				// the setting itself is `null`, the config {} block will be
				// set to empty and will cause the plan to always have a diff:
				// + config {}.
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

//
// Copyright 2023 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inmem

import (
	"context"
	"fmt"
	"strconv"

	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// Internal data: Jenkins
type jenkinsMap map[string][]*jenkinsLink
type jenkinsLink struct {
	id            uint32
	image         string
	jobName       string
	buildNumber   string
	gitUrl        string
	gitCommit     string
	gitCommitDiff string
	jobUrl        string
}

func (n *jenkinsLink) ID() uint32 { return n.id }

func (n *jenkinsLink) BuildModelNode(c *demoClient) (model.Node, error) {
	return c.buildJenkins(n, nil)
}

func (n *jenkinsLink) Neighbors(allowedEdges edgeMap) []uint32 {
	out := make([]uint32, 0, 2)
	return out
}

// Ingest Jenkins
func (c *demoClient) IngestJenkins(ctx context.Context, jenkins model.JenkinsInputSpec) (*model.Jenkins, error) {
	return c.ingestJenkins(ctx, jenkins, true)
}

func (c *demoClient) ingestJenkins(ctx context.Context, jenkins model.JenkinsInputSpec, readOnly bool) (*model.Jenkins, error) {
	funcName := "IngestJenkins"

	lock(&c.m, readOnly)
	defer unlock(&c.m, readOnly)

	if jenkinsData, hasJob := c.jenkins[jenkins.JobName]; hasJob {
		for _, savedJenkinsData := range jenkinsData {
			if savedJenkinsData.buildNumber == jenkins.BuildNumber {
				return nil, gqlerror.Errorf("%v ::  %s", funcName, "jenkins build number already exists for the given jobName cannot ingest duplicate data")
			}
		}
	}

	ingestJenkinsData := jenkinsLink{
		id:            c.getNextID(),
		image:         jenkins.Image,
		jobName:       jenkins.JobName,
		buildNumber:   jenkins.BuildNumber,
		gitUrl:        *jenkins.GitURL,
		gitCommit:     *jenkins.GitCommit,
		gitCommitDiff: *jenkins.GitCommitDiff,
		jobUrl:        *jenkins.JobURL,
	}

	c.jenkins[jenkins.JobName] = append(c.jenkins[jenkins.JobName], &ingestJenkinsData)

	c.index[ingestJenkinsData.id] = &ingestJenkinsData

	builtJenkins, err := c.buildJenkins(&ingestJenkinsData, nil)
	if err != nil {
		return nil, err
	}
	return builtJenkins, nil

}

func (c *demoClient) Jenkins(ctx context.Context, filter *model.JenkinsSpec) ([]*model.Jenkins, error) {
	c.m.RLock()
	defer c.m.RUnlock()
	funcName := "Jenkins"

	if filter != nil && filter.ID != nil {
		id64, err := strconv.ParseUint(*filter.ID, 10, 32)
		if err != nil {
			return nil, gqlerror.Errorf("%v :: invalid ID %s", funcName, err)
		}
		id := uint32(id64)
		link, err := byID[*jenkinsLink](id, c)
		if err != nil {
			// Not found
			return nil, nil
		}
		foundJenkins, err := c.buildJenkins(link, filter)
		if err != nil {
			return nil, gqlerror.Errorf("%v :: %v", funcName, err)
		}
		return []*model.Jenkins{foundJenkins}, nil
	}

	var returnJenkinsData []*model.Jenkins
	for jobName, jenkinsData := range c.jenkins {

		if filter != nil && filter.JobName != nil && filter.JobName != &jobName {
			continue
		}

		for _, jenkinsLink := range jenkinsData {

			if filter.BuildNumber != nil && jenkinsLink.buildNumber != *filter.BuildNumber {
				continue
			}

			if filter.Image != nil && jenkinsLink.image != *filter.Image {
				continue
			}

			appendJenkinsData, err := c.buildJenkins(jenkinsLink, filter)
			if err != nil {
				return nil, gqlerror.Errorf("%v :: %v", funcName, err)
			}
			returnJenkinsData = append(returnJenkinsData, appendJenkinsData)
		}
	}

	return returnJenkinsData, nil
}

func (c *demoClient) buildJenkins(jenkins *jenkinsLink, filter *model.JenkinsSpec) (*model.Jenkins, error) {

	returnJenkinsData := model.Jenkins{
		ID:            fmt.Sprint(jenkins.id),
		Image:         jenkins.image,
		JobName:       jenkins.jobName,
		BuildNumber:   jenkins.buildNumber,
		GitURL:        &jenkins.gitUrl,
		GitCommit:     &jenkins.gitCommit,
		GitCommitDiff: &jenkins.gitCommitDiff,
		JobURL:        &jenkins.jobUrl,
	}
	return &returnJenkinsData, nil
}

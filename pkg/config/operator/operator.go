// Copyright 2021 Allstar Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package operator contains config to be set by the GitHub App operator
package operator

import "time"

// AppID should be set to the application ID of the created GitHub App. See:
// https://docs.github.com/en/developers/apps/building-github-apps/authenticating-with-github-apps#authenticating-as-a-github-app
const AppID = 119816

// KeySecret should be set to the name of a secret containing a private key for
// the App. See:
// https://docs.github.com/en/developers/apps/building-github-apps/authenticating-with-github-apps#generating-a-private-key
// The secret is retrieved with gocloud.dev/runtimevar.
const KeySecret = "gcpsecretmanager://projects/allstar-ossf/secrets/allstar-private-key?decoder=bytes"

// OrgConfigRepo is the name of the expected org-level repo to contain config.
const OrgConfigRepo = ".allstar"

// RepoConfigDir is the name of the expected directory in each repo to contain
// repo-level config.
const RepoConfigDir = ".allstar"

// AppConfigFile is the name of the expected file in org or repo level config.
const AppConfigFile = "allstar.yaml"

// GitHubIssueLabel is the label used to tag, search, and identify GitHub
// Issues created by the bot.
const GitHubIssueLabel = "allstar"

// GitHubIssueFooter is added to the end of GitHub issues.
const GitHubIssueFooter = "Issue created by Allstar. https://github.com/ossf/allstar"

// NoticePingDuration is the duration to wait between pinging notice actions,
// such as updating a GitHub issue.
const NoticePingDuration = (24 * time.Hour)

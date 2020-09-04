package github

import (
	"time"

	"github.com/shurcooL/githubv4"
)

// RepositoryType returns some strategic
// $repo and $org need to be combined with
// this query type in order to perform a
// successful request against github's gql api
// query {
//     repository(name: "platform-internal", owner: "stoplightio") {
//       id
//       isPrivate
//       isFork
//       isArchived
//       isTemplate
//       url
//       issues(first: 1, orderBy: { field: CREATED_AT, direction: ASC}) {
//         totalCount
//         pageInfo {
//           startCursor
//           endCursor
//           hasNextPage
//           hasPreviousPage
//         }
//       }
//     }
// }
type RepositoryType struct {
	Repository struct {
		Name       string
		IsPrivate  bool
		IsFork     bool
		IsArchived bool
		IsTemplate bool
		URL        string
		CreatedAt  time.Time
		Owner      struct {
			Login string
		}
		Issues struct {
			TotalCount int
			PageInfo   pageInfo
		} `graphql:"issues(first: 1, orderBy: { field: CREATED_AT, direction: ASC})"`
	} `graphql:"repository(name: $repo, owner: $org)"`
}

// RepositoryQuery returns found repo info and issue count and cursors
func RepositoryQuery(organization string, repository string) (RepositoryType, map[string]interface{}) {
	variables := map[string]interface{}{
		"org":  githubv4.String(organization),
		"repo": githubv4.String(repository),
	}

	return RepositoryType{}, variables
}

// RepositoryIssuesQueryType is the type defining the query
// {
//   repository(name: "platform-internal", owner: "stoplightio") {
//     issues(last: 100, orderBy: {field: CREATED_AT, direction: ASC}) {
//       totalCount
//       pageInfo {
//         startCursor
//         endCursor
//         hasNextPage
//         hasPreviousPage
//       }
//       nodes {
//		   title
//         number
//		   url
//         createdAt
//         publishedAt
//         closedAt
//         milestone {
//           title
//         }
//         labels(first: 10) {
//           totalCount
//           nodes {
//             name
//           }
//         }
//         assignees(first: 5) {
//           totalCount
//           nodes {
//             login
//           }
//         }
//       }
//     }
//   }
// }
type RepositoryIssuesQueryType struct {
	Repository struct {
		Issues struct {
			Nodes      []Issue
			TotalCount int
			PageInfo   pageInfo
		} `graphql:"issues(first: 100, orderBy: { field: CREATED_AT, direction: DESC}, after: $cursor)"`
	} `graphql:"repository(name: $repo, owner: $org)"`
}

// IssuesQuery returns paginated issues and cursor info
func IssuesQuery(organization string, repository string) (RepositoryIssuesQueryType, map[string]interface{}) {
	variables := map[string]interface{}{
		"org":    githubv4.String(organization),
		"repo":   githubv4.String(repository),
		"cursor": (*githubv4.String)(nil),
	}

	return RepositoryIssuesQueryType{}, variables
}

type Issue struct {
	Title     string
	Number    int
	URL       string
	CreatedAt time.Time
	UpdatedAt time.Time
	ClosedAt  *time.Time
	Milestone struct {
		Title string
	}
	Assignees struct {
		TotalCount int
		Nodes      []assignee
	} `graphql:"assignees(first: 5)"`
	Labels struct {
		TotalCount int
		Nodes      []label
	} `graphql:"labels(first: 10)"`
}
type assignee struct {
	Login string
}
type label struct {
	Name string
}
type pageInfo struct {
	HasPreviousPage bool
	HasNextPage     bool
	StartCursor     string
	EndCursor       string
}

// RepoRepositoryPullRequestsQueryType
// {
//   repository(name: "platform-internal", owner: "stoplightio") {
//     pullRequests(last: 10, orderBy: {field: CREATED_AT, direction: ASC}) {
//       nodes {
//         title
//         number
//         url
//         createdAt
//         closedAt
//         mergedAt
//         isDraft
//         reviewDecision
//         publishedAt
//         state
//         updatedAt
//         labels(first: 10) {
//           totalCount
//           nodes {
//             name
//           }
//         }
//         author {
//           login
//         }
//         mergedBy {
//           login
//         }
//         commits(first: 1) {
//           nodes {
//             commit {
//               committedDate
//             }
//           }
//         }
//       }
//       totalCount
//		 pageInfo {
//         startCursor
//         endCursor
//         hasNextPage
//         hasPreviousPage
//       }
//     }
//   }
// }
type RepositoryPullRequestsQueryType struct {
	Repository struct {
		PullRequests struct {
			Nodes      []PullRequest
			TotalCount int
			PageInfo   pageInfo
		} `graphql:"pullRequests(first: 100, orderBy: { field: CREATED_AT, direction: DESC}, after: $cursor)"`
	} `graphql:"repository(name: $repo, owner: $org)"`
}

func PullRequestsQuery(organization string, repository string) (RepositoryPullRequestsQueryType, map[string]interface{}) {
	variables := map[string]interface{}{
		"org":    githubv4.String(organization),
		"repo":   githubv4.String(repository),
		"cursor": (*githubv4.String)(nil),
	}

	return RepositoryPullRequestsQueryType{}, variables
}

// PullRequest represents a gh pr
type PullRequest struct {
	Title          string
	Number         int
	URL            string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ClosedAt       *time.Time
	MergedAt       *time.Time
	PublishedAt    *time.Time
	IsDraft        bool
	State          string
	ReviewDecision string
	Milestone      struct {
		Title *string
	}
	MergedBy struct {
		Login *string
	}
	Author struct {
		Login string
	}
	Commits struct {
		Nodes []commit
	} `graphql:"commits(first: 1)"`
	Labels struct {
		TotalCount int
		Nodes      []label
	} `graphql:"labels(first: 10)"`
}

type commit struct {
	Commit struct {
		CommittedDate *time.Time
	}
}

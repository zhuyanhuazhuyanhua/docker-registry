package handlers

import (
	"context"
	"fmt"

	"github.com/distribution/distribution/v3"
	"github.com/opencontainers/go-digest"
	"github.com/zhuyanhuazhuyanhua/docker-registry/internal/dcontext"
	"github.com/zhuyanhuazhuyanhua/docker-registry/registry/api/errcode"
	v2 "github.com/zhuyanhuazhuyanhua/docker-registry/registry/api/v2"
)

// Context should contain the request specific context for use in across
// handlers. Resources that don't need to be shared across handlers should not
// be on this object.
type Context struct {
	// App points to the application structure that created this context.
	*App
	context.Context

	// Repository is the repository for the current request. All requests
	// should be scoped to a single repository. This field may be nil.
	Repository distribution.Repository

	// RepositoryRemover provides method to delete a repository
	RepositoryRemover distribution.RepositoryRemover

	// Errors is a collection of errors encountered during the request to be
	// returned to the client API. If errors are added to the collection, the
	// handler *must not* start the response via http.ResponseWriter.
	Errors errcode.Errors

	urlBuilder *v2.URLBuilder

	// TODO(stevvooe): The goal is too completely factor this context and
	// dispatching out of the web application. Ideally, we should lean on
	// context.Context for injection of these resources.
}

func getUploadUUID(ctx context.Context) (uuid string) {
	return dcontext.GetStringValue(ctx, "vars.uuid")
}

var errDigestNotAvailable = fmt.Errorf("digest not available in context")

func getDigest(ctx context.Context) (dgst digest.Digest, err error) {
	dgstStr := dcontext.GetStringValue(ctx, "vars.digest")

	if dgstStr == "" {
		dcontext.GetLogger(ctx).Errorf("digest not available")
		return "", errDigestNotAvailable
	}

	d, err := digest.Parse(dgstStr)
	if err != nil {
		dcontext.GetLogger(ctx).Errorf("error parsing digest=%q: %v", dgstStr, err)
		return "", err
	}

	return d, nil
}

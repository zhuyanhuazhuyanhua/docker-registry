package handlers

import (
	"context"
	"net/url"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/configuration"
	"github.com/distribution/distribution/v3/notifications"
	"github.com/distribution/distribution/v3/registry/auth"
	"github.com/docker/go-events"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	storagedriver "github.com/zhuyanhuazhuyanhua/docker-registry/registry/storage/driver"
)

// App is a global registry application object. Shared resources can be placed
// on this object that will be accessible from all requests. Any writable
// fields should be protected.
type App struct {
	context.Context

	Config *configuration.Configuration

	router           *mux.Router                    // main application router, configured with dispatchers
	driver           storagedriver.StorageDriver    // driver maintains the app global storage driver instance.
	registry         distribution.Namespace         // registry is the primary registry backend for the app instance.
	repoRemover      distribution.RepositoryRemover // repoRemover provides ability to delete repos
	accessController auth.AccessController          // main access controller for application

	// httpHost is a parsed representation of the http.host parameter from
	// the configuration. Only the Scheme and Host fields are used.
	httpHost url.URL

	// events contains notification related configuration.
	events struct {
		sink   events.Sink
		source notifications.SourceRecord
	}

	redis redis.UniversalClient

	// isCache is true if this registry is configured as a pull through cache
	isCache bool

	// readOnly is true if the registry is in a read-only maintenance mode
	readOnly bool
}

package ISession

type Session interface {

	// GetName get the name of the session.
	GetName() string

	// SetName Set the name of the session.
	SetName(name string)

	// GetId get the current session ID.
	GetId() string

	// SetId Set the session ID.
	SetId(id string)

	// Start  the session, reading the data from a handler.
	Start() bool

	// Save save the session data to storage.
	Save()

	// All get  session data.
	All() map[string]string

	// Exists Checks if a key exists.
	Exists(key string) bool

	// Has Checks if a key is present and not null.
	Has(key string) bool

	// Get  an item from the session.
	Get(key, defaultValue string) string

	// Pull  the value of a given key and then forget it.
	Pull(key, defaultValue string) string

	// Put  a key / value pair or array of key / value pairs in the session.
	Put(key, value string)

	// Token get the CSRF token value.
	Token() string

	// RegenerateToken regenerate the CSRF token value.
	RegenerateToken()

	// Remove  an item from the session, returning its value.
	Remove(key string) string

	// Forget remove one or many items from the session.
	Forget(keys ...string)

	// Flush remove all  items from the session.
	Flush()

	// Invalidate flush the session data and regenerate the ID.
	Invalidate() bool

	// Regenerate Generate a new session identifier.
	Regenerate(destroy bool) bool

	// Migrate Generate a new session ID for the session.
	Migrate(destroy bool) bool

	// IsStarted Determine if the session has been started.
	IsStarted() bool

	// PreviousUrl get the previous URL from the session.
	PreviousUrl() string

	// SetPreviousUrl Set the "previous" URL in the session.
	SetPreviousUrl(url string)
}

type SessionStore interface {

	// LoadSession Load the session data from the handler.
	LoadSession(id string) map[string]string

	// Save save the session data to storage.
	Save(id string, sessions map[string]string)
}

package commons

type Context struct {
	Dump    interface{}     `json:"dump,omitempty"`
	Error   []ErrorCMap     `json:"error,omitempty"`
	Query   *QueryContext   `json:"db,omitempty"`
	Request *RequestContext `json:"request,omitempty"`
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) AddDump(ctx interface{}) {
	c.Dump = &ctx
}

func (c *Context) AddErrorContext(ctx []ErrorCMap) {
	c.Error = ctx
}

func (c *Context) AddQueryContext(ctx QueryContext) {
	c.Query = &ctx
}

func (c *Context) AddRequestContext(ctx RequestContext) {
	c.Request = &ctx
}

// FromError this method convert and standard error to ErrorContext.
func (c *Context) FromError(err ErrorI) *Context {
	c.Error = Unwrap(err)

	return c
}

// QueryContext is used in gorm trace integration.
type QueryContext struct {
	Caller  string  `json:"caller"`
	Query   string  `json:"query"`
	Rows    int64   `json:"rows"`
	Errors  error   `json:"errors"`
	Latency float64 `json:"latency"`
}

// RequestContext is used in echo middleware integration.
type RequestContext struct {
	Host      string `json:"host"`
	Method    string `json:"method"`
	Uri       string `json:"uri"`
	RemoteIp  string `json:"remote_ip"`
	UserAgent string `json:"user_agent"`
	Latency   string `json:"latency"`
	LatencyH  string `json:"latency_human"`
	Status    int    `json:"status"`
	BytesIn   int64  `json:"bytes_in"`
	BytesOut  int64  `json:"bytes_out"`
}

package http

func (o *Request) ensureClient() {
	if o.Client == nil {
		o.WithDefaultClient()
	}
}

func (o *Request) ensureQuery() {
	if o.Query == nil {
		o.WithDefaultQuery()
	}
}

func (o *Request) ensureHeader() {
	if o.Header == nil {
		o.WithDefaultHeader()
	}
}

func (o *Request) ensure() {
	o.ensureClient()
	o.ensureQuery()
	o.ensureHeader()
}

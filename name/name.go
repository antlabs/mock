package name

func Name(opts ...Options) string {
	var options Options
	for _, opt := range opts {
		options = opt
	}

	if options.IsChinese {
		return CName()
	}

	return ""
}

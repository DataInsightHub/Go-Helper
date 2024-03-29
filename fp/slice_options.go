package fp

type SliceTransformOption interface {
	apply(*sliceOptions)
}


type (
	sliceOptions struct {
		limit *int
	}
)

type limitOption int

func (limitOption limitOption) apply(options *sliceOptions) {
	limit := int(limitOption)
	
	if limit <= 0 {
		return 
	}

	
	options.limit = &limit 
}

func WithLimit(limit int) SliceTransformOption {
	return limitOption(limit)
}
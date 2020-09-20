package method

// Std like methods
type (
	Print func(args ...interface{})

	Println func(args ...interface{})

	Printf func(fmt string, args ...interface{})

	Output func(depth int, msg string)
)

type (
	Debugf func(fmt string, args ...interface{})

	Infof func(fmt string, args ...interface{})

	Warnf func(fmt string, args ...interface{})

	Errorf func(fmt string, args ...interface{})
)

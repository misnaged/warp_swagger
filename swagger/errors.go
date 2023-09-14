package swagger

import "errors"

var ErrGenTypeUndefined = errors.New("generation type is undefined. Must be: server/models/client`")
var ErrNilGenerationPathOrOutput = errors.New("spec file path AND/OR generation output path are nil")
var ErrNilCfgServerOpts = errors.New(" value: `Server` in your config_swagger is nil")
var ErrNilCfgModelsOpts = errors.New(" value: `Models` in your config_swagger is nil")
var ErrClientNotImplemented = errors.New("client option is not implemented yet")

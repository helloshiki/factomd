// Copyright 2015 FactomProject Authors. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package anchor

import (
	log "github.com/sirupsen/logrus"
)


// setup subsystem loggers
var (
	anchorLog = log.WithFields(log.Fields{"package": "anchor"})
)

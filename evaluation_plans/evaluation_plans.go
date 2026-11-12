package evaluation_plans

import (
	"github.com/gemaraproj/go-gemara"

	"github.com/revanite-io/pvtr-aws-secmgmt/data"
	"github.com/revanite-io/pvtr-aws-secmgmt/evaluation_plans/reusable_steps"
)

// TypedStep is the signature every step in this plugin uses. It receives the
// plugin's own data.Payload rather than an untyped any; the SDK's typed
// registration helpers (pluginkit.AddEvaluationSuiteTyped and
// ...TypedForAllCatalogs) adapt it to gemara.AssessmentStep. Always register
// through those helpers: wrapping a step in a local closure erases its
// function name from the evaluation log.
type TypedStep func(data.Payload) (gemara.Result, string, gemara.ConfidenceLevel)

var (
	// Add more entries here if other catalogs or catalog versions are adopted by the plugin
	// Then use pluginkit.AddEvaluationSuiteTyped to make these available to the user
	// Exported so package main can reference it. snake_case preserves the
	// catalog id's casing, and most ids are lowercase, so the Suite_ prefix
	// is what guarantees the identifier is exported.
	Suite_CCC_SecMgmt_Controls = map[string][]TypedStep{
		// For best results, ensure every assessment id is represented in this map
		"CCC.SecMgmt.CN01.AR01": {
			/* When a secret is stored, a rotation schedule MUST be configured
with an interval of 90 days or less.
 */
			reusable_steps.NotImplemented,
		},
		"CCC.SecMgmt.CN01.AR02": {
			/* When a scheduled rotation has not produced a new secret version
within 24 hours of its due time, an alert MUST be published to a
monitored channel.
 */
			reusable_steps.NotImplemented,
		},
		"CCC.SecMgmt.CN03.AR01": {
			/* When a secret's access policy is configured, read permission on the
secret value MUST NOT be granted to a wildcard principal.
 */
			reusable_steps.NotImplemented,
		},
		"CCC.SecMgmt.CN03.AR02": {
			/* When read permission on a secret value is granted, the grant MUST
be scoped to the specific secret, vault, or tag-scoped set of
secrets that the identity consumes.
 */
			reusable_steps.NotImplemented,
		},
		"CCC.SecMgmt.CN03.AR03": {
			/* When an identity is granted read permission on a secret value for
consumption by a workload, that grant MUST NOT include permission
to create a new version of the secret.
 */
			reusable_steps.NotImplemented,
		},
		"CCC.SecMgmt.CN04.AR01": {
			/* When a secret or secret version is deleted, the service MUST retain
it in a recoverable state for at least 7 days before destruction.
 */
			reusable_steps.NotImplemented,
		},
		"CCC.SecMgmt.CN04.AR02": {
			/* When a secret version is disabled or scheduled for destruction, an
alert MUST be published to a monitored channel.
 */
			reusable_steps.NotImplemented,
		},
		// CCC.Core.Controls
		"CCC.Core.CN01": {
			reusable_steps.NotImplemented,
		},
		"CCC.Core.CN02": {
			reusable_steps.NotImplemented,
		},
		"CCC.Core.CN03": {
			reusable_steps.NotImplemented,
		},
		"CCC.Core.CN04": {
			reusable_steps.NotImplemented,
		},
		"CCC.Core.CN05": {
			reusable_steps.NotImplemented,
		},
		"CCC.Core.CN06": {
			reusable_steps.NotImplemented,
		},
		"CCC.Core.CN07": {
			reusable_steps.NotImplemented,
		},
		"CCC.Core.CN08": {
			reusable_steps.NotImplemented,
		},
		"CCC.Core.CN09": {
			reusable_steps.NotImplemented,
		},
		"CCC.Core.CN10": {
			reusable_steps.NotImplemented,
		},
		"CCC.Core.CN11": {
			reusable_steps.NotImplemented,
		},
		"CCC.Core.CN16": {
			reusable_steps.NotImplemented,
		},
	}
)

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Remote Cargo Repository Resource
//
// Provides an Artifactory remote `cargo` repository resource. This provides cargo specific fields and is the only way to get them
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Cargo+Registry)
//
// ## Example Usage
//
// Create a new Artifactory remote cargo repository called my-remote-cargo
// for brevity sake, only cargo specific fields are included; for other fields see documentation for
// generic repo.
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewRemoteCargoRepository(ctx, "my-remote-cargo", &artifactory.RemoteCargoRepositoryArgs{
// 			AnonymousAccess: pulumi.Bool(true),
// 			GitRegistryUrl:  pulumi.String("https://github.com/rust-lang/foo.index"),
// 			Key:             pulumi.String("my-remote-cargo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Note
//
// If you get a 400 error: `"Custom Base URL should be defined prior to creating a Cargo repository"`,
// you must set the base url at: `http://${host}/ui/admin/configuration/general`
type RemoteCargoRepository struct {
	pulumi.CustomResourceState

	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth pulumi.BoolOutput `pulumi:"allowAnyHostAuth"`
	// - Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option.
	AnonymousAccess pulumi.BoolPtrOutput `pulumi:"anonymousAccess"`
	// The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
	// an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
	// offline. Default to 300.
	AssumedOfflinePeriodSecs pulumi.IntPtrOutput `pulumi:"assumedOfflinePeriodSecs"`
	// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
	// resolution.
	BlackedOut pulumi.BoolOutput `pulumi:"blackedOut"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BlockMismatchingMimeTypes pulumi.BoolOutput `pulumi:"blockMismatchingMimeTypes"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests   pulumi.BoolOutput   `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate pulumi.StringOutput `pulumi:"clientTlsCertificate"`
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation RemoteCargoRepositoryContentSynchronisationOutput `pulumi:"contentSynchronisation"`
	Description            pulumi.StringOutput                               `pulumi:"description"`
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement pulumi.BoolOutput   `pulumi:"enableCookieManagement"`
	ExcludesPattern        pulumi.StringOutput `pulumi:"excludesPattern"`
	// Deprecated: This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
	FailedRetrievalCachePeriodSecs pulumi.IntOutput `pulumi:"failedRetrievalCachePeriodSecs"`
	// - This is the index url, expected to be a git repository. for remote artifactory use "arturl/git/repokey.git"
	GitRegistryUrl  pulumi.StringOutput `pulumi:"gitRegistryUrl"`
	HardFail        pulumi.BoolOutput   `pulumi:"hardFail"`
	IncludesPattern pulumi.StringOutput `pulumi:"includesPattern"`
	// The repository identifier. Must be unique system-wide
	Key          pulumi.StringOutput    `pulumi:"key"`
	LocalAddress pulumi.StringPtrOutput `pulumi:"localAddress"`
	// This is actually the missedRetrievalCachePeriodSecs in the API
	MissedCachePeriodSeconds pulumi.IntOutput       `pulumi:"missedCachePeriodSeconds"`
	Notes                    pulumi.StringPtrOutput `pulumi:"notes"`
	// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
	Offline     pulumi.BoolOutput      `pulumi:"offline"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	Password    pulumi.StringPtrOutput `pulumi:"password"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey           pulumi.StringPtrOutput   `pulumi:"projectKey"`
	PropagateQueryParams pulumi.BoolPtrOutput     `pulumi:"propagateQueryParams"`
	PropertySets         pulumi.StringArrayOutput `pulumi:"propertySets"`
	Proxy                pulumi.StringOutput      `pulumi:"proxy"`
	RemoteRepoLayoutRef  pulumi.StringOutput      `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef        pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds pulumi.IntOutput  `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration          pulumi.BoolOutput `pulumi:"shareConfiguration"`
	SocketTimeoutMillis         pulumi.IntOutput  `pulumi:"socketTimeoutMillis"`
	// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
	// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
	// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
	// servers.
	StoreArtifactsLocally pulumi.BoolOutput `pulumi:"storeArtifactsLocally"`
	// When set, remote artifacts are fetched along with their properties.
	SynchronizeProperties               pulumi.BoolOutput      `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodEnabled pulumi.BoolOutput      `pulumi:"unusedArtifactsCleanupPeriodEnabled"`
	UnusedArtifactsCleanupPeriodHours   pulumi.IntOutput       `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                                 pulumi.StringOutput    `pulumi:"url"`
	Username                            pulumi.StringPtrOutput `pulumi:"username"`
	XrayIndex                           pulumi.BoolOutput      `pulumi:"xrayIndex"`
}

// NewRemoteCargoRepository registers a new resource with the given unique name, arguments, and options.
func NewRemoteCargoRepository(ctx *pulumi.Context,
	name string, args *RemoteCargoRepositoryArgs, opts ...pulumi.ResourceOption) (*RemoteCargoRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GitRegistryUrl == nil {
		return nil, errors.New("invalid value for required argument 'GitRegistryUrl'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource RemoteCargoRepository
	err := ctx.RegisterResource("artifactory:index/remoteCargoRepository:RemoteCargoRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemoteCargoRepository gets an existing RemoteCargoRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemoteCargoRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemoteCargoRepositoryState, opts ...pulumi.ResourceOption) (*RemoteCargoRepository, error) {
	var resource RemoteCargoRepository
	err := ctx.ReadResource("artifactory:index/remoteCargoRepository:RemoteCargoRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemoteCargoRepository resources.
type remoteCargoRepositoryState struct {
	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth *bool `pulumi:"allowAnyHostAuth"`
	// - Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option.
	AnonymousAccess *bool `pulumi:"anonymousAccess"`
	// The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
	// an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
	// offline. Default to 300.
	AssumedOfflinePeriodSecs *int `pulumi:"assumedOfflinePeriodSecs"`
	// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
	// resolution.
	BlackedOut *bool `pulumi:"blackedOut"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BlockMismatchingMimeTypes *bool `pulumi:"blockMismatchingMimeTypes"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests   *bool   `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate *string `pulumi:"clientTlsCertificate"`
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation *RemoteCargoRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description            *string                                      `pulumi:"description"`
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement *bool   `pulumi:"enableCookieManagement"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// Deprecated: This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
	FailedRetrievalCachePeriodSecs *int `pulumi:"failedRetrievalCachePeriodSecs"`
	// - This is the index url, expected to be a git repository. for remote artifactory use "arturl/git/repokey.git"
	GitRegistryUrl  *string `pulumi:"gitRegistryUrl"`
	HardFail        *bool   `pulumi:"hardFail"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// The repository identifier. Must be unique system-wide
	Key          *string `pulumi:"key"`
	LocalAddress *string `pulumi:"localAddress"`
	// This is actually the missedRetrievalCachePeriodSecs in the API
	MissedCachePeriodSeconds *int    `pulumi:"missedCachePeriodSeconds"`
	Notes                    *string `pulumi:"notes"`
	// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
	Offline     *bool   `pulumi:"offline"`
	PackageType *string `pulumi:"packageType"`
	Password    *string `pulumi:"password"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey           *string  `pulumi:"projectKey"`
	PropagateQueryParams *bool    `pulumi:"propagateQueryParams"`
	PropertySets         []string `pulumi:"propertySets"`
	Proxy                *string  `pulumi:"proxy"`
	RemoteRepoLayoutRef  *string  `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef        *string  `pulumi:"repoLayoutRef"`
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds *int  `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration          *bool `pulumi:"shareConfiguration"`
	SocketTimeoutMillis         *int  `pulumi:"socketTimeoutMillis"`
	// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
	// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
	// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
	// servers.
	StoreArtifactsLocally *bool `pulumi:"storeArtifactsLocally"`
	// When set, remote artifacts are fetched along with their properties.
	SynchronizeProperties               *bool   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodEnabled *bool   `pulumi:"unusedArtifactsCleanupPeriodEnabled"`
	UnusedArtifactsCleanupPeriodHours   *int    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                                 *string `pulumi:"url"`
	Username                            *string `pulumi:"username"`
	XrayIndex                           *bool   `pulumi:"xrayIndex"`
}

type RemoteCargoRepositoryState struct {
	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth pulumi.BoolPtrInput
	// - Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option.
	AnonymousAccess pulumi.BoolPtrInput
	// The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
	// an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
	// offline. Default to 300.
	AssumedOfflinePeriodSecs pulumi.IntPtrInput
	// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
	// resolution.
	BlackedOut pulumi.BoolPtrInput
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BlockMismatchingMimeTypes pulumi.BoolPtrInput
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests   pulumi.BoolPtrInput
	ClientTlsCertificate pulumi.StringPtrInput
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation RemoteCargoRepositoryContentSynchronisationPtrInput
	Description            pulumi.StringPtrInput
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	// Deprecated: This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
	FailedRetrievalCachePeriodSecs pulumi.IntPtrInput
	// - This is the index url, expected to be a git repository. for remote artifactory use "arturl/git/repokey.git"
	GitRegistryUrl  pulumi.StringPtrInput
	HardFail        pulumi.BoolPtrInput
	IncludesPattern pulumi.StringPtrInput
	// The repository identifier. Must be unique system-wide
	Key          pulumi.StringPtrInput
	LocalAddress pulumi.StringPtrInput
	// This is actually the missedRetrievalCachePeriodSecs in the API
	MissedCachePeriodSeconds pulumi.IntPtrInput
	Notes                    pulumi.StringPtrInput
	// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
	Offline     pulumi.BoolPtrInput
	PackageType pulumi.StringPtrInput
	Password    pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey           pulumi.StringPtrInput
	PropagateQueryParams pulumi.BoolPtrInput
	PropertySets         pulumi.StringArrayInput
	Proxy                pulumi.StringPtrInput
	RemoteRepoLayoutRef  pulumi.StringPtrInput
	RepoLayoutRef        pulumi.StringPtrInput
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
	ShareConfiguration          pulumi.BoolPtrInput
	SocketTimeoutMillis         pulumi.IntPtrInput
	// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
	// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
	// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
	// servers.
	StoreArtifactsLocally pulumi.BoolPtrInput
	// When set, remote artifacts are fetched along with their properties.
	SynchronizeProperties               pulumi.BoolPtrInput
	UnusedArtifactsCleanupPeriodEnabled pulumi.BoolPtrInput
	UnusedArtifactsCleanupPeriodHours   pulumi.IntPtrInput
	Url                                 pulumi.StringPtrInput
	Username                            pulumi.StringPtrInput
	XrayIndex                           pulumi.BoolPtrInput
}

func (RemoteCargoRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteCargoRepositoryState)(nil)).Elem()
}

type remoteCargoRepositoryArgs struct {
	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth *bool `pulumi:"allowAnyHostAuth"`
	// - Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option.
	AnonymousAccess *bool `pulumi:"anonymousAccess"`
	// The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
	// an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
	// offline. Default to 300.
	AssumedOfflinePeriodSecs *int `pulumi:"assumedOfflinePeriodSecs"`
	// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
	// resolution.
	BlackedOut *bool `pulumi:"blackedOut"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BlockMismatchingMimeTypes *bool `pulumi:"blockMismatchingMimeTypes"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests   *bool   `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate *string `pulumi:"clientTlsCertificate"`
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation *RemoteCargoRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description            *string                                      `pulumi:"description"`
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement *bool   `pulumi:"enableCookieManagement"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// - This is the index url, expected to be a git repository. for remote artifactory use "arturl/git/repokey.git"
	GitRegistryUrl  string  `pulumi:"gitRegistryUrl"`
	HardFail        *bool   `pulumi:"hardFail"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// The repository identifier. Must be unique system-wide
	Key          string  `pulumi:"key"`
	LocalAddress *string `pulumi:"localAddress"`
	// This is actually the missedRetrievalCachePeriodSecs in the API
	MissedCachePeriodSeconds *int    `pulumi:"missedCachePeriodSeconds"`
	Notes                    *string `pulumi:"notes"`
	// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
	Offline  *bool   `pulumi:"offline"`
	Password *string `pulumi:"password"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey           *string  `pulumi:"projectKey"`
	PropagateQueryParams *bool    `pulumi:"propagateQueryParams"`
	PropertySets         []string `pulumi:"propertySets"`
	Proxy                *string  `pulumi:"proxy"`
	RemoteRepoLayoutRef  *string  `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef        *string  `pulumi:"repoLayoutRef"`
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds *int  `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration          *bool `pulumi:"shareConfiguration"`
	SocketTimeoutMillis         *int  `pulumi:"socketTimeoutMillis"`
	// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
	// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
	// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
	// servers.
	StoreArtifactsLocally *bool `pulumi:"storeArtifactsLocally"`
	// When set, remote artifacts are fetched along with their properties.
	SynchronizeProperties               *bool   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodEnabled *bool   `pulumi:"unusedArtifactsCleanupPeriodEnabled"`
	UnusedArtifactsCleanupPeriodHours   *int    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                                 string  `pulumi:"url"`
	Username                            *string `pulumi:"username"`
	XrayIndex                           *bool   `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a RemoteCargoRepository resource.
type RemoteCargoRepositoryArgs struct {
	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth pulumi.BoolPtrInput
	// - Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option.
	AnonymousAccess pulumi.BoolPtrInput
	// The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
	// an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
	// offline. Default to 300.
	AssumedOfflinePeriodSecs pulumi.IntPtrInput
	// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
	// resolution.
	BlackedOut pulumi.BoolPtrInput
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BlockMismatchingMimeTypes pulumi.BoolPtrInput
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests   pulumi.BoolPtrInput
	ClientTlsCertificate pulumi.StringPtrInput
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation RemoteCargoRepositoryContentSynchronisationPtrInput
	Description            pulumi.StringPtrInput
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	// - This is the index url, expected to be a git repository. for remote artifactory use "arturl/git/repokey.git"
	GitRegistryUrl  pulumi.StringInput
	HardFail        pulumi.BoolPtrInput
	IncludesPattern pulumi.StringPtrInput
	// The repository identifier. Must be unique system-wide
	Key          pulumi.StringInput
	LocalAddress pulumi.StringPtrInput
	// This is actually the missedRetrievalCachePeriodSecs in the API
	MissedCachePeriodSeconds pulumi.IntPtrInput
	Notes                    pulumi.StringPtrInput
	// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
	Offline  pulumi.BoolPtrInput
	Password pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey           pulumi.StringPtrInput
	PropagateQueryParams pulumi.BoolPtrInput
	PropertySets         pulumi.StringArrayInput
	Proxy                pulumi.StringPtrInput
	RemoteRepoLayoutRef  pulumi.StringPtrInput
	RepoLayoutRef        pulumi.StringPtrInput
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
	ShareConfiguration          pulumi.BoolPtrInput
	SocketTimeoutMillis         pulumi.IntPtrInput
	// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
	// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
	// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
	// servers.
	StoreArtifactsLocally pulumi.BoolPtrInput
	// When set, remote artifacts are fetched along with their properties.
	SynchronizeProperties               pulumi.BoolPtrInput
	UnusedArtifactsCleanupPeriodEnabled pulumi.BoolPtrInput
	UnusedArtifactsCleanupPeriodHours   pulumi.IntPtrInput
	Url                                 pulumi.StringInput
	Username                            pulumi.StringPtrInput
	XrayIndex                           pulumi.BoolPtrInput
}

func (RemoteCargoRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteCargoRepositoryArgs)(nil)).Elem()
}

type RemoteCargoRepositoryInput interface {
	pulumi.Input

	ToRemoteCargoRepositoryOutput() RemoteCargoRepositoryOutput
	ToRemoteCargoRepositoryOutputWithContext(ctx context.Context) RemoteCargoRepositoryOutput
}

func (*RemoteCargoRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteCargoRepository)(nil)).Elem()
}

func (i *RemoteCargoRepository) ToRemoteCargoRepositoryOutput() RemoteCargoRepositoryOutput {
	return i.ToRemoteCargoRepositoryOutputWithContext(context.Background())
}

func (i *RemoteCargoRepository) ToRemoteCargoRepositoryOutputWithContext(ctx context.Context) RemoteCargoRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteCargoRepositoryOutput)
}

// RemoteCargoRepositoryArrayInput is an input type that accepts RemoteCargoRepositoryArray and RemoteCargoRepositoryArrayOutput values.
// You can construct a concrete instance of `RemoteCargoRepositoryArrayInput` via:
//
//          RemoteCargoRepositoryArray{ RemoteCargoRepositoryArgs{...} }
type RemoteCargoRepositoryArrayInput interface {
	pulumi.Input

	ToRemoteCargoRepositoryArrayOutput() RemoteCargoRepositoryArrayOutput
	ToRemoteCargoRepositoryArrayOutputWithContext(context.Context) RemoteCargoRepositoryArrayOutput
}

type RemoteCargoRepositoryArray []RemoteCargoRepositoryInput

func (RemoteCargoRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemoteCargoRepository)(nil)).Elem()
}

func (i RemoteCargoRepositoryArray) ToRemoteCargoRepositoryArrayOutput() RemoteCargoRepositoryArrayOutput {
	return i.ToRemoteCargoRepositoryArrayOutputWithContext(context.Background())
}

func (i RemoteCargoRepositoryArray) ToRemoteCargoRepositoryArrayOutputWithContext(ctx context.Context) RemoteCargoRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteCargoRepositoryArrayOutput)
}

// RemoteCargoRepositoryMapInput is an input type that accepts RemoteCargoRepositoryMap and RemoteCargoRepositoryMapOutput values.
// You can construct a concrete instance of `RemoteCargoRepositoryMapInput` via:
//
//          RemoteCargoRepositoryMap{ "key": RemoteCargoRepositoryArgs{...} }
type RemoteCargoRepositoryMapInput interface {
	pulumi.Input

	ToRemoteCargoRepositoryMapOutput() RemoteCargoRepositoryMapOutput
	ToRemoteCargoRepositoryMapOutputWithContext(context.Context) RemoteCargoRepositoryMapOutput
}

type RemoteCargoRepositoryMap map[string]RemoteCargoRepositoryInput

func (RemoteCargoRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemoteCargoRepository)(nil)).Elem()
}

func (i RemoteCargoRepositoryMap) ToRemoteCargoRepositoryMapOutput() RemoteCargoRepositoryMapOutput {
	return i.ToRemoteCargoRepositoryMapOutputWithContext(context.Background())
}

func (i RemoteCargoRepositoryMap) ToRemoteCargoRepositoryMapOutputWithContext(ctx context.Context) RemoteCargoRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteCargoRepositoryMapOutput)
}

type RemoteCargoRepositoryOutput struct{ *pulumi.OutputState }

func (RemoteCargoRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteCargoRepository)(nil)).Elem()
}

func (o RemoteCargoRepositoryOutput) ToRemoteCargoRepositoryOutput() RemoteCargoRepositoryOutput {
	return o
}

func (o RemoteCargoRepositoryOutput) ToRemoteCargoRepositoryOutputWithContext(ctx context.Context) RemoteCargoRepositoryOutput {
	return o
}

type RemoteCargoRepositoryArrayOutput struct{ *pulumi.OutputState }

func (RemoteCargoRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemoteCargoRepository)(nil)).Elem()
}

func (o RemoteCargoRepositoryArrayOutput) ToRemoteCargoRepositoryArrayOutput() RemoteCargoRepositoryArrayOutput {
	return o
}

func (o RemoteCargoRepositoryArrayOutput) ToRemoteCargoRepositoryArrayOutputWithContext(ctx context.Context) RemoteCargoRepositoryArrayOutput {
	return o
}

func (o RemoteCargoRepositoryArrayOutput) Index(i pulumi.IntInput) RemoteCargoRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RemoteCargoRepository {
		return vs[0].([]*RemoteCargoRepository)[vs[1].(int)]
	}).(RemoteCargoRepositoryOutput)
}

type RemoteCargoRepositoryMapOutput struct{ *pulumi.OutputState }

func (RemoteCargoRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemoteCargoRepository)(nil)).Elem()
}

func (o RemoteCargoRepositoryMapOutput) ToRemoteCargoRepositoryMapOutput() RemoteCargoRepositoryMapOutput {
	return o
}

func (o RemoteCargoRepositoryMapOutput) ToRemoteCargoRepositoryMapOutputWithContext(ctx context.Context) RemoteCargoRepositoryMapOutput {
	return o
}

func (o RemoteCargoRepositoryMapOutput) MapIndex(k pulumi.StringInput) RemoteCargoRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RemoteCargoRepository {
		return vs[0].(map[string]*RemoteCargoRepository)[vs[1].(string)]
	}).(RemoteCargoRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteCargoRepositoryInput)(nil)).Elem(), &RemoteCargoRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteCargoRepositoryArrayInput)(nil)).Elem(), RemoteCargoRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteCargoRepositoryMapInput)(nil)).Elem(), RemoteCargoRepositoryMap{})
	pulumi.RegisterOutputType(RemoteCargoRepositoryOutput{})
	pulumi.RegisterOutputType(RemoteCargoRepositoryArrayOutput{})
	pulumi.RegisterOutputType(RemoteCargoRepositoryMapOutput{})
}

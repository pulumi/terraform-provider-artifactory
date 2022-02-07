// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Remote Repository Resource
//
// Provides an Artifactory remote `docker` repository resource. This provides docker specific fields and is the only way to
// get them
//
// ## Example Usage
//
// Includes only new and relevant fields
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
// 		_, err := artifactory.NewRemoteDockerRepository(ctx, "my_remote_docker", &artifactory.RemoteDockerRepositoryArgs{
// 			BlockPushingSchema1:         pulumi.Bool(true),
// 			EnableTokenAuthentication:   pulumi.Bool(true),
// 			ExternalDependenciesEnabled: pulumi.Bool(true),
// 			ExternalDependenciesPatterns: pulumi.StringArray{
// 				pulumi.String("**/hub.docker.io/**"),
// 				pulumi.String("**/bintray.jfrog.io/**"),
// 			},
// 			Key: pulumi.String("my-remote-docker"),
// 			Url: pulumi.String("https://hub.docker.io/"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RemoteDockerRepository struct {
	pulumi.CustomResourceState

	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth pulumi.BoolOutput `pulumi:"allowAnyHostAuth"`
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
	// When set, Artifactory will block the pulling of Docker images with manifest v2
	// schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1
	// that exist in the cache.
	BlockPushingSchema1 pulumi.BoolOutput `pulumi:"blockPushingSchema1"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests   pulumi.BoolOutput   `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate pulumi.StringOutput `pulumi:"clientTlsCertificate"`
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation RemoteDockerRepositoryContentSynchronisationOutput `pulumi:"contentSynchronisation"`
	Description            pulumi.StringOutput                                `pulumi:"description"`
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement pulumi.BoolOutput `pulumi:"enableCookieManagement"`
	// Enable token (Bearer) based authentication.
	EnableTokenAuthentication pulumi.BoolOutput   `pulumi:"enableTokenAuthentication"`
	ExcludesPattern           pulumi.StringOutput `pulumi:"excludesPattern"`
	// Also known as 'Foreign Layers Caching' on the UI
	ExternalDependenciesEnabled pulumi.BoolOutput `pulumi:"externalDependenciesEnabled"`
	// An allow list of Ant-style path patterns that determine which remote VCS
	ExternalDependenciesPatterns pulumi.StringArrayOutput `pulumi:"externalDependenciesPatterns"`
	// Deprecated: This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
	FailedRetrievalCachePeriodSecs pulumi.IntOutput    `pulumi:"failedRetrievalCachePeriodSecs"`
	HardFail                       pulumi.BoolOutput   `pulumi:"hardFail"`
	IncludesPattern                pulumi.StringOutput `pulumi:"includesPattern"`
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
	PriorityResolution   pulumi.BoolOutput        `pulumi:"priorityResolution"`
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
	SynchronizeProperties               pulumi.BoolOutput `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodEnabled pulumi.BoolOutput `pulumi:"unusedArtifactsCleanupPeriodEnabled"`
	UnusedArtifactsCleanupPeriodHours   pulumi.IntOutput  `pulumi:"unusedArtifactsCleanupPeriodHours"`
	// - the remote repo URL. You kinda don't have a remote repo without it
	Url       pulumi.StringOutput    `pulumi:"url"`
	Username  pulumi.StringPtrOutput `pulumi:"username"`
	XrayIndex pulumi.BoolOutput      `pulumi:"xrayIndex"`
}

// NewRemoteDockerRepository registers a new resource with the given unique name, arguments, and options.
func NewRemoteDockerRepository(ctx *pulumi.Context,
	name string, args *RemoteDockerRepositoryArgs, opts ...pulumi.ResourceOption) (*RemoteDockerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource RemoteDockerRepository
	err := ctx.RegisterResource("artifactory:index/remoteDockerRepository:RemoteDockerRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemoteDockerRepository gets an existing RemoteDockerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemoteDockerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemoteDockerRepositoryState, opts ...pulumi.ResourceOption) (*RemoteDockerRepository, error) {
	var resource RemoteDockerRepository
	err := ctx.ReadResource("artifactory:index/remoteDockerRepository:RemoteDockerRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemoteDockerRepository resources.
type remoteDockerRepositoryState struct {
	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth *bool `pulumi:"allowAnyHostAuth"`
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
	// When set, Artifactory will block the pulling of Docker images with manifest v2
	// schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1
	// that exist in the cache.
	BlockPushingSchema1 *bool `pulumi:"blockPushingSchema1"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests   *bool   `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate *string `pulumi:"clientTlsCertificate"`
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation *RemoteDockerRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description            *string                                       `pulumi:"description"`
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement *bool `pulumi:"enableCookieManagement"`
	// Enable token (Bearer) based authentication.
	EnableTokenAuthentication *bool   `pulumi:"enableTokenAuthentication"`
	ExcludesPattern           *string `pulumi:"excludesPattern"`
	// Also known as 'Foreign Layers Caching' on the UI
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// An allow list of Ant-style path patterns that determine which remote VCS
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
	// Deprecated: This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
	FailedRetrievalCachePeriodSecs *int    `pulumi:"failedRetrievalCachePeriodSecs"`
	HardFail                       *bool   `pulumi:"hardFail"`
	IncludesPattern                *string `pulumi:"includesPattern"`
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
	PriorityResolution   *bool    `pulumi:"priorityResolution"`
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
	SynchronizeProperties               *bool `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodEnabled *bool `pulumi:"unusedArtifactsCleanupPeriodEnabled"`
	UnusedArtifactsCleanupPeriodHours   *int  `pulumi:"unusedArtifactsCleanupPeriodHours"`
	// - the remote repo URL. You kinda don't have a remote repo without it
	Url       *string `pulumi:"url"`
	Username  *string `pulumi:"username"`
	XrayIndex *bool   `pulumi:"xrayIndex"`
}

type RemoteDockerRepositoryState struct {
	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth pulumi.BoolPtrInput
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
	// When set, Artifactory will block the pulling of Docker images with manifest v2
	// schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1
	// that exist in the cache.
	BlockPushingSchema1 pulumi.BoolPtrInput
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests   pulumi.BoolPtrInput
	ClientTlsCertificate pulumi.StringPtrInput
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation RemoteDockerRepositoryContentSynchronisationPtrInput
	Description            pulumi.StringPtrInput
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement pulumi.BoolPtrInput
	// Enable token (Bearer) based authentication.
	EnableTokenAuthentication pulumi.BoolPtrInput
	ExcludesPattern           pulumi.StringPtrInput
	// Also known as 'Foreign Layers Caching' on the UI
	ExternalDependenciesEnabled pulumi.BoolPtrInput
	// An allow list of Ant-style path patterns that determine which remote VCS
	ExternalDependenciesPatterns pulumi.StringArrayInput
	// Deprecated: This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
	FailedRetrievalCachePeriodSecs pulumi.IntPtrInput
	HardFail                       pulumi.BoolPtrInput
	IncludesPattern                pulumi.StringPtrInput
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
	PriorityResolution   pulumi.BoolPtrInput
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
	// - the remote repo URL. You kinda don't have a remote repo without it
	Url       pulumi.StringPtrInput
	Username  pulumi.StringPtrInput
	XrayIndex pulumi.BoolPtrInput
}

func (RemoteDockerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteDockerRepositoryState)(nil)).Elem()
}

type remoteDockerRepositoryArgs struct {
	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth *bool `pulumi:"allowAnyHostAuth"`
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
	// When set, Artifactory will block the pulling of Docker images with manifest v2
	// schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1
	// that exist in the cache.
	BlockPushingSchema1 *bool `pulumi:"blockPushingSchema1"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests   *bool   `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate *string `pulumi:"clientTlsCertificate"`
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation *RemoteDockerRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description            *string                                       `pulumi:"description"`
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement *bool `pulumi:"enableCookieManagement"`
	// Enable token (Bearer) based authentication.
	EnableTokenAuthentication *bool   `pulumi:"enableTokenAuthentication"`
	ExcludesPattern           *string `pulumi:"excludesPattern"`
	// Also known as 'Foreign Layers Caching' on the UI
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// An allow list of Ant-style path patterns that determine which remote VCS
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
	HardFail                     *bool    `pulumi:"hardFail"`
	IncludesPattern              *string  `pulumi:"includesPattern"`
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
	PriorityResolution   *bool    `pulumi:"priorityResolution"`
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
	SynchronizeProperties               *bool `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodEnabled *bool `pulumi:"unusedArtifactsCleanupPeriodEnabled"`
	UnusedArtifactsCleanupPeriodHours   *int  `pulumi:"unusedArtifactsCleanupPeriodHours"`
	// - the remote repo URL. You kinda don't have a remote repo without it
	Url       string  `pulumi:"url"`
	Username  *string `pulumi:"username"`
	XrayIndex *bool   `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a RemoteDockerRepository resource.
type RemoteDockerRepositoryArgs struct {
	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth pulumi.BoolPtrInput
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
	// When set, Artifactory will block the pulling of Docker images with manifest v2
	// schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1
	// that exist in the cache.
	BlockPushingSchema1 pulumi.BoolPtrInput
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests   pulumi.BoolPtrInput
	ClientTlsCertificate pulumi.StringPtrInput
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation RemoteDockerRepositoryContentSynchronisationPtrInput
	Description            pulumi.StringPtrInput
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement pulumi.BoolPtrInput
	// Enable token (Bearer) based authentication.
	EnableTokenAuthentication pulumi.BoolPtrInput
	ExcludesPattern           pulumi.StringPtrInput
	// Also known as 'Foreign Layers Caching' on the UI
	ExternalDependenciesEnabled pulumi.BoolPtrInput
	// An allow list of Ant-style path patterns that determine which remote VCS
	ExternalDependenciesPatterns pulumi.StringArrayInput
	HardFail                     pulumi.BoolPtrInput
	IncludesPattern              pulumi.StringPtrInput
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
	PriorityResolution   pulumi.BoolPtrInput
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
	// - the remote repo URL. You kinda don't have a remote repo without it
	Url       pulumi.StringInput
	Username  pulumi.StringPtrInput
	XrayIndex pulumi.BoolPtrInput
}

func (RemoteDockerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteDockerRepositoryArgs)(nil)).Elem()
}

type RemoteDockerRepositoryInput interface {
	pulumi.Input

	ToRemoteDockerRepositoryOutput() RemoteDockerRepositoryOutput
	ToRemoteDockerRepositoryOutputWithContext(ctx context.Context) RemoteDockerRepositoryOutput
}

func (*RemoteDockerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteDockerRepository)(nil))
}

func (i *RemoteDockerRepository) ToRemoteDockerRepositoryOutput() RemoteDockerRepositoryOutput {
	return i.ToRemoteDockerRepositoryOutputWithContext(context.Background())
}

func (i *RemoteDockerRepository) ToRemoteDockerRepositoryOutputWithContext(ctx context.Context) RemoteDockerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteDockerRepositoryOutput)
}

func (i *RemoteDockerRepository) ToRemoteDockerRepositoryPtrOutput() RemoteDockerRepositoryPtrOutput {
	return i.ToRemoteDockerRepositoryPtrOutputWithContext(context.Background())
}

func (i *RemoteDockerRepository) ToRemoteDockerRepositoryPtrOutputWithContext(ctx context.Context) RemoteDockerRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteDockerRepositoryPtrOutput)
}

type RemoteDockerRepositoryPtrInput interface {
	pulumi.Input

	ToRemoteDockerRepositoryPtrOutput() RemoteDockerRepositoryPtrOutput
	ToRemoteDockerRepositoryPtrOutputWithContext(ctx context.Context) RemoteDockerRepositoryPtrOutput
}

type remoteDockerRepositoryPtrType RemoteDockerRepositoryArgs

func (*remoteDockerRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteDockerRepository)(nil))
}

func (i *remoteDockerRepositoryPtrType) ToRemoteDockerRepositoryPtrOutput() RemoteDockerRepositoryPtrOutput {
	return i.ToRemoteDockerRepositoryPtrOutputWithContext(context.Background())
}

func (i *remoteDockerRepositoryPtrType) ToRemoteDockerRepositoryPtrOutputWithContext(ctx context.Context) RemoteDockerRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteDockerRepositoryPtrOutput)
}

// RemoteDockerRepositoryArrayInput is an input type that accepts RemoteDockerRepositoryArray and RemoteDockerRepositoryArrayOutput values.
// You can construct a concrete instance of `RemoteDockerRepositoryArrayInput` via:
//
//          RemoteDockerRepositoryArray{ RemoteDockerRepositoryArgs{...} }
type RemoteDockerRepositoryArrayInput interface {
	pulumi.Input

	ToRemoteDockerRepositoryArrayOutput() RemoteDockerRepositoryArrayOutput
	ToRemoteDockerRepositoryArrayOutputWithContext(context.Context) RemoteDockerRepositoryArrayOutput
}

type RemoteDockerRepositoryArray []RemoteDockerRepositoryInput

func (RemoteDockerRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemoteDockerRepository)(nil)).Elem()
}

func (i RemoteDockerRepositoryArray) ToRemoteDockerRepositoryArrayOutput() RemoteDockerRepositoryArrayOutput {
	return i.ToRemoteDockerRepositoryArrayOutputWithContext(context.Background())
}

func (i RemoteDockerRepositoryArray) ToRemoteDockerRepositoryArrayOutputWithContext(ctx context.Context) RemoteDockerRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteDockerRepositoryArrayOutput)
}

// RemoteDockerRepositoryMapInput is an input type that accepts RemoteDockerRepositoryMap and RemoteDockerRepositoryMapOutput values.
// You can construct a concrete instance of `RemoteDockerRepositoryMapInput` via:
//
//          RemoteDockerRepositoryMap{ "key": RemoteDockerRepositoryArgs{...} }
type RemoteDockerRepositoryMapInput interface {
	pulumi.Input

	ToRemoteDockerRepositoryMapOutput() RemoteDockerRepositoryMapOutput
	ToRemoteDockerRepositoryMapOutputWithContext(context.Context) RemoteDockerRepositoryMapOutput
}

type RemoteDockerRepositoryMap map[string]RemoteDockerRepositoryInput

func (RemoteDockerRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemoteDockerRepository)(nil)).Elem()
}

func (i RemoteDockerRepositoryMap) ToRemoteDockerRepositoryMapOutput() RemoteDockerRepositoryMapOutput {
	return i.ToRemoteDockerRepositoryMapOutputWithContext(context.Background())
}

func (i RemoteDockerRepositoryMap) ToRemoteDockerRepositoryMapOutputWithContext(ctx context.Context) RemoteDockerRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteDockerRepositoryMapOutput)
}

type RemoteDockerRepositoryOutput struct{ *pulumi.OutputState }

func (RemoteDockerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteDockerRepository)(nil))
}

func (o RemoteDockerRepositoryOutput) ToRemoteDockerRepositoryOutput() RemoteDockerRepositoryOutput {
	return o
}

func (o RemoteDockerRepositoryOutput) ToRemoteDockerRepositoryOutputWithContext(ctx context.Context) RemoteDockerRepositoryOutput {
	return o
}

func (o RemoteDockerRepositoryOutput) ToRemoteDockerRepositoryPtrOutput() RemoteDockerRepositoryPtrOutput {
	return o.ToRemoteDockerRepositoryPtrOutputWithContext(context.Background())
}

func (o RemoteDockerRepositoryOutput) ToRemoteDockerRepositoryPtrOutputWithContext(ctx context.Context) RemoteDockerRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemoteDockerRepository) *RemoteDockerRepository {
		return &v
	}).(RemoteDockerRepositoryPtrOutput)
}

type RemoteDockerRepositoryPtrOutput struct{ *pulumi.OutputState }

func (RemoteDockerRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteDockerRepository)(nil))
}

func (o RemoteDockerRepositoryPtrOutput) ToRemoteDockerRepositoryPtrOutput() RemoteDockerRepositoryPtrOutput {
	return o
}

func (o RemoteDockerRepositoryPtrOutput) ToRemoteDockerRepositoryPtrOutputWithContext(ctx context.Context) RemoteDockerRepositoryPtrOutput {
	return o
}

func (o RemoteDockerRepositoryPtrOutput) Elem() RemoteDockerRepositoryOutput {
	return o.ApplyT(func(v *RemoteDockerRepository) RemoteDockerRepository {
		if v != nil {
			return *v
		}
		var ret RemoteDockerRepository
		return ret
	}).(RemoteDockerRepositoryOutput)
}

type RemoteDockerRepositoryArrayOutput struct{ *pulumi.OutputState }

func (RemoteDockerRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RemoteDockerRepository)(nil))
}

func (o RemoteDockerRepositoryArrayOutput) ToRemoteDockerRepositoryArrayOutput() RemoteDockerRepositoryArrayOutput {
	return o
}

func (o RemoteDockerRepositoryArrayOutput) ToRemoteDockerRepositoryArrayOutputWithContext(ctx context.Context) RemoteDockerRepositoryArrayOutput {
	return o
}

func (o RemoteDockerRepositoryArrayOutput) Index(i pulumi.IntInput) RemoteDockerRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RemoteDockerRepository {
		return vs[0].([]RemoteDockerRepository)[vs[1].(int)]
	}).(RemoteDockerRepositoryOutput)
}

type RemoteDockerRepositoryMapOutput struct{ *pulumi.OutputState }

func (RemoteDockerRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RemoteDockerRepository)(nil))
}

func (o RemoteDockerRepositoryMapOutput) ToRemoteDockerRepositoryMapOutput() RemoteDockerRepositoryMapOutput {
	return o
}

func (o RemoteDockerRepositoryMapOutput) ToRemoteDockerRepositoryMapOutputWithContext(ctx context.Context) RemoteDockerRepositoryMapOutput {
	return o
}

func (o RemoteDockerRepositoryMapOutput) MapIndex(k pulumi.StringInput) RemoteDockerRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RemoteDockerRepository {
		return vs[0].(map[string]RemoteDockerRepository)[vs[1].(string)]
	}).(RemoteDockerRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteDockerRepositoryInput)(nil)).Elem(), &RemoteDockerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteDockerRepositoryPtrInput)(nil)).Elem(), &RemoteDockerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteDockerRepositoryArrayInput)(nil)).Elem(), RemoteDockerRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteDockerRepositoryMapInput)(nil)).Elem(), RemoteDockerRepositoryMap{})
	pulumi.RegisterOutputType(RemoteDockerRepositoryOutput{})
	pulumi.RegisterOutputType(RemoteDockerRepositoryPtrOutput{})
	pulumi.RegisterOutputType(RemoteDockerRepositoryArrayOutput{})
	pulumi.RegisterOutputType(RemoteDockerRepositoryMapOutput{})
}

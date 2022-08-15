// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.RemotePubRepositoryArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.RemotePubRepositoryState;
import com.pulumi.artifactory.outputs.RemotePubRepositoryContentSynchronisation;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a remote Pub repository.
 * Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Pub+Repositories).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.RemotePubRepository;
 * import com.pulumi.artifactory.RemotePubRepositoryArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var my_remote_pub = new RemotePubRepository(&#34;my-remote-pub&#34;, RemotePubRepositoryArgs.builder()        
 *             .key(&#34;my-remote-pub&#34;)
 *             .url(&#34;https://pub.dartlang.org&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Remote repositories can be imported using their name, e.g.
 * 
 * ```sh
 *  $ pulumi import artifactory:index/remotePubRepository:RemotePubRepository my-remote-pub my-remote-pub
 * ```
 * 
 */
@ResourceType(type="artifactory:index/remotePubRepository:RemotePubRepository")
public class RemotePubRepository extends com.pulumi.resources.CustomResource {
    /**
     * Also known as &#39;Lenient Host Authentication&#39;, Allow credentials of this repository to be used on requests redirected to
     * any other host.
     * 
     */
    @Export(name="allowAnyHostAuth", type=Boolean.class, parameters={})
    private Output<Boolean> allowAnyHostAuth;

    /**
     * @return Also known as &#39;Lenient Host Authentication&#39;, Allow credentials of this repository to be used on requests redirected to
     * any other host.
     * 
     */
    public Output<Boolean> allowAnyHostAuth() {
        return this.allowAnyHostAuth;
    }
    /**
     * The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
     * an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
     * offline. Default to 300.
     * 
     */
    @Export(name="assumedOfflinePeriodSecs", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> assumedOfflinePeriodSecs;

    /**
     * @return The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
     * an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
     * offline. Default to 300.
     * 
     */
    public Output<Optional<Integer>> assumedOfflinePeriodSecs() {
        return Codegen.optional(this.assumedOfflinePeriodSecs);
    }
    /**
     * (A.K.A &#39;Ignore Repository&#39; on the UI) When set, the repository or its local cache do not participate in artifact
     * resolution.
     * 
     */
    @Export(name="blackedOut", type=Boolean.class, parameters={})
    private Output<Boolean> blackedOut;

    /**
     * @return (A.K.A &#39;Ignore Repository&#39; on the UI) When set, the repository or its local cache do not participate in artifact
     * resolution.
     * 
     */
    public Output<Boolean> blackedOut() {
        return this.blackedOut;
    }
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     * 
     */
    @Export(name="blockMismatchingMimeTypes", type=Boolean.class, parameters={})
    private Output<Boolean> blockMismatchingMimeTypes;

    /**
     * @return Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     * 
     */
    public Output<Boolean> blockMismatchingMimeTypes() {
        return this.blockMismatchingMimeTypes;
    }
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     * 
     */
    @Export(name="bypassHeadRequests", type=Boolean.class, parameters={})
    private Output<Boolean> bypassHeadRequests;

    /**
     * @return Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     * 
     */
    public Output<Boolean> bypassHeadRequests() {
        return this.bypassHeadRequests;
    }
    @Export(name="clientTlsCertificate", type=String.class, parameters={})
    private Output<String> clientTlsCertificate;

    public Output<String> clientTlsCertificate() {
        return this.clientTlsCertificate;
    }
    @Export(name="contentSynchronisation", type=RemotePubRepositoryContentSynchronisation.class, parameters={})
    private Output<RemotePubRepositoryContentSynchronisation> contentSynchronisation;

    public Output<RemotePubRepositoryContentSynchronisation> contentSynchronisation() {
        return this.contentSynchronisation;
    }
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    public Output<String> description() {
        return this.description;
    }
    /**
     * Enables cookie management if the remote repository uses cookies to manage client state.
     * 
     */
    @Export(name="enableCookieManagement", type=Boolean.class, parameters={})
    private Output<Boolean> enableCookieManagement;

    /**
     * @return Enables cookie management if the remote repository uses cookies to manage client state.
     * 
     */
    public Output<Boolean> enableCookieManagement() {
        return this.enableCookieManagement;
    }
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
     * artifacts are excluded.
     * 
     */
    @Export(name="excludesPattern", type=String.class, parameters={})
    private Output<String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
     * artifacts are excluded.
     * 
     */
    public Output<String> excludesPattern() {
        return this.excludesPattern;
    }
    /**
     * @deprecated
     * This field is not returned in a get payload but is offered on the UI. It&#39;s inserted here for inclusive and informational reasons. It does not function
     * 
     */
    @Deprecated /* This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function */
    @Export(name="failedRetrievalCachePeriodSecs", type=Integer.class, parameters={})
    private Output<Integer> failedRetrievalCachePeriodSecs;

    public Output<Integer> failedRetrievalCachePeriodSecs() {
        return this.failedRetrievalCachePeriodSecs;
    }
    /**
     * When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to
     * communicate with this repository.
     * 
     */
    @Export(name="hardFail", type=Boolean.class, parameters={})
    private Output<Boolean> hardFail;

    /**
     * @return When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to
     * communicate with this repository.
     * 
     */
    public Output<Boolean> hardFail() {
        return this.hardFail;
    }
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Export(name="includesPattern", type=String.class, parameters={})
    private Output<String> includesPattern;

    /**
     * @return List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    public Output<String> includesPattern() {
        return this.includesPattern;
    }
    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of
     * the &#39;Retrieval Cache Period&#39;. Default value is &#39;false&#39;.
     * 
     */
    @Export(name="listRemoteFolderItems", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> listRemoteFolderItems;

    /**
     * @return Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of
     * the &#39;Retrieval Cache Period&#39;. Default value is &#39;false&#39;.
     * 
     */
    public Output<Optional<Boolean>> listRemoteFolderItems() {
        return Codegen.optional(this.listRemoteFolderItems);
    }
    /**
     * The local address to be used when creating connections. Useful for specifying the interface to use on systems with
     * multiple network interfaces.
     * 
     */
    @Export(name="localAddress", type=String.class, parameters={})
    private Output</* @Nullable */ String> localAddress;

    /**
     * @return The local address to be used when creating connections. Useful for specifying the interface to use on systems with
     * multiple network interfaces.
     * 
     */
    public Output<Optional<String>> localAddress() {
        return Codegen.optional(this.localAddress);
    }
    /**
     * The set of mime types that should override the block_mismatching_mime_types setting. Eg:
     * &#34;application/json,application/xml&#34;. Default value is empty.
     * 
     */
    @Export(name="mismatchingMimeTypesOverrideList", type=String.class, parameters={})
    private Output</* @Nullable */ String> mismatchingMimeTypesOverrideList;

    /**
     * @return The set of mime types that should override the block_mismatching_mime_types setting. Eg:
     * &#34;application/json,application/xml&#34;. Default value is empty.
     * 
     */
    public Output<Optional<String>> mismatchingMimeTypesOverrideList() {
        return Codegen.optional(this.mismatchingMimeTypesOverrideList);
    }
    /**
     * The number of seconds to cache artifact retrieval misses (artifact not found). A value of 0 indicates no caching.
     * 
     */
    @Export(name="missedCachePeriodSeconds", type=Integer.class, parameters={})
    private Output<Integer> missedCachePeriodSeconds;

    /**
     * @return The number of seconds to cache artifact retrieval misses (artifact not found). A value of 0 indicates no caching.
     * 
     */
    public Output<Integer> missedCachePeriodSeconds() {
        return this.missedCachePeriodSeconds;
    }
    @Export(name="notes", type=String.class, parameters={})
    private Output</* @Nullable */ String> notes;

    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    /**
     * If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     * 
     */
    @Export(name="offline", type=Boolean.class, parameters={})
    private Output<Boolean> offline;

    /**
     * @return If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     * 
     */
    public Output<Boolean> offline() {
        return this.offline;
    }
    @Export(name="packageType", type=String.class, parameters={})
    private Output<String> packageType;

    public Output<String> packageType() {
        return this.packageType;
    }
    @Export(name="password", type=String.class, parameters={})
    private Output</* @Nullable */ String> password;

    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    @Export(name="priorityResolution", type=Boolean.class, parameters={})
    private Output<Boolean> priorityResolution;

    /**
     * @return Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    public Output<Boolean> priorityResolution() {
        return this.priorityResolution;
    }
    /**
     * Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
     * 
     */
    @Export(name="projectEnvironments", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> projectEnvironments;

    /**
     * @return Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
     * 
     */
    public Output<Optional<List<String>>> projectEnvironments() {
        return Codegen.optional(this.projectEnvironments);
    }
    /**
     * Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    @Export(name="projectKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    public Output<Optional<String>> projectKey() {
        return Codegen.optional(this.projectKey);
    }
    /**
     * When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     * 
     */
    @Export(name="propagateQueryParams", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> propagateQueryParams;

    /**
     * @return When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     * 
     */
    public Output<Optional<Boolean>> propagateQueryParams() {
        return Codegen.optional(this.propagateQueryParams);
    }
    /**
     * List of property set names
     * 
     */
    @Export(name="propertySets", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> propertySets;

    /**
     * @return List of property set names
     * 
     */
    public Output<Optional<List<String>>> propertySets() {
        return Codegen.optional(this.propertySets);
    }
    /**
     * Proxy key from Artifactory Proxies settings
     * 
     */
    @Export(name="proxy", type=String.class, parameters={})
    private Output</* @Nullable */ String> proxy;

    /**
     * @return Proxy key from Artifactory Proxies settings
     * 
     */
    public Output<Optional<String>> proxy() {
        return Codegen.optional(this.proxy);
    }
    /**
     * Repository layout key for the remote layout mapping
     * 
     */
    @Export(name="remoteRepoLayoutRef", type=String.class, parameters={})
    private Output<String> remoteRepoLayoutRef;

    /**
     * @return Repository layout key for the remote layout mapping
     * 
     */
    public Output<String> remoteRepoLayoutRef() {
        return this.remoteRepoLayoutRef;
    }
    /**
     * Repository layout key for the local repository
     * 
     */
    @Export(name="repoLayoutRef", type=String.class, parameters={})
    private Output</* @Nullable */ String> repoLayoutRef;

    /**
     * @return Repository layout key for the local repository
     * 
     */
    public Output<Optional<String>> repoLayoutRef() {
        return Codegen.optional(this.repoLayoutRef);
    }
    /**
     * The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
     * 
     */
    @Export(name="retrievalCachePeriodSeconds", type=Integer.class, parameters={})
    private Output<Integer> retrievalCachePeriodSeconds;

    /**
     * @return The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
     * 
     */
    public Output<Integer> retrievalCachePeriodSeconds() {
        return this.retrievalCachePeriodSeconds;
    }
    @Export(name="shareConfiguration", type=Boolean.class, parameters={})
    private Output<Boolean> shareConfiguration;

    public Output<Boolean> shareConfiguration() {
        return this.shareConfiguration;
    }
    /**
     * Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network
     * operation is considered a retrieval failure.
     * 
     */
    @Export(name="socketTimeoutMillis", type=Integer.class, parameters={})
    private Output<Integer> socketTimeoutMillis;

    /**
     * @return Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network
     * operation is considered a retrieval failure.
     * 
     */
    public Output<Integer> socketTimeoutMillis() {
        return this.socketTimeoutMillis;
    }
    /**
     * When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
     * direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
     * one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
     * servers.
     * 
     */
    @Export(name="storeArtifactsLocally", type=Boolean.class, parameters={})
    private Output<Boolean> storeArtifactsLocally;

    /**
     * @return When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
     * direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
     * one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
     * servers.
     * 
     */
    public Output<Boolean> storeArtifactsLocally() {
        return this.storeArtifactsLocally;
    }
    /**
     * When set, remote artifacts are fetched along with their properties.
     * 
     */
    @Export(name="synchronizeProperties", type=Boolean.class, parameters={})
    private Output<Boolean> synchronizeProperties;

    /**
     * @return When set, remote artifacts are fetched along with their properties.
     * 
     */
    public Output<Boolean> synchronizeProperties() {
        return this.synchronizeProperties;
    }
    @Export(name="unusedArtifactsCleanupPeriodEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> unusedArtifactsCleanupPeriodEnabled;

    public Output<Boolean> unusedArtifactsCleanupPeriodEnabled() {
        return this.unusedArtifactsCleanupPeriodEnabled;
    }
    /**
     * The number of hours to wait before an artifact is deemed &#34;unused&#34; and eligible for cleanup from the repository. A value
     * of 0 means automatic cleanup of cached artifacts is disabled.
     * 
     */
    @Export(name="unusedArtifactsCleanupPeriodHours", type=Integer.class, parameters={})
    private Output<Integer> unusedArtifactsCleanupPeriodHours;

    /**
     * @return The number of hours to wait before an artifact is deemed &#34;unused&#34; and eligible for cleanup from the repository. A value
     * of 0 means automatic cleanup of cached artifacts is disabled.
     * 
     */
    public Output<Integer> unusedArtifactsCleanupPeriodHours() {
        return this.unusedArtifactsCleanupPeriodHours;
    }
    /**
     * The remote repository URL.
     * 
     */
    @Export(name="url", type=String.class, parameters={})
    private Output<String> url;

    /**
     * @return The remote repository URL.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    @Export(name="username", type=String.class, parameters={})
    private Output</* @Nullable */ String> username;

    public Output<Optional<String>> username() {
        return Codegen.optional(this.username);
    }
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    @Export(name="xrayIndex", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> xrayIndex;

    /**
     * @return Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    public Output<Optional<Boolean>> xrayIndex() {
        return Codegen.optional(this.xrayIndex);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RemotePubRepository(String name) {
        this(name, RemotePubRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RemotePubRepository(String name, RemotePubRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RemotePubRepository(String name, RemotePubRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/remotePubRepository:RemotePubRepository", name, args == null ? RemotePubRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RemotePubRepository(String name, Output<String> id, @Nullable RemotePubRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/remotePubRepository:RemotePubRepository", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static RemotePubRepository get(String name, Output<String> id, @Nullable RemotePubRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RemotePubRepository(name, id, state, options);
    }
}

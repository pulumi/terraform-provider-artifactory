// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.RemoteDebianRepositoryArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.RemoteDebianRepositoryState;
import com.pulumi.artifactory.outputs.RemoteDebianRepositoryContentSynchronisation;
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
 * Creates a remote Debian repository.
 * Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Debian+Repositories).
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.RemoteDebianRepository;
 * import com.pulumi.artifactory.RemoteDebianRepositoryArgs;
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
 *         var my_remote_debian = new RemoteDebianRepository("my-remote-debian", RemoteDebianRepositoryArgs.builder()
 *             .key("my-remote-Debian")
 *             .url("http://archive.ubuntu.com/ubuntu/")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Remote repositories can be imported using their name, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/remoteDebianRepository:RemoteDebianRepository my-remote-debian my-remote-debian
 * ```
 * 
 */
@ResourceType(type="artifactory:index/remoteDebianRepository:RemoteDebianRepository")
public class RemoteDebianRepository extends com.pulumi.resources.CustomResource {
    /**
     * &#39;Lenient Host Authentication&#39; in the UI. Allow credentials of this repository to be used on requests redirected to any
     * other host.
     * 
     */
    @Export(name="allowAnyHostAuth", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> allowAnyHostAuth;

    /**
     * @return &#39;Lenient Host Authentication&#39; in the UI. Allow credentials of this repository to be used on requests redirected to any
     * other host.
     * 
     */
    public Output<Boolean> allowAnyHostAuth() {
        return this.allowAnyHostAuth;
    }
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    @Export(name="archiveBrowsingEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> archiveBrowsingEnabled;

    /**
     * @return When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    public Output<Boolean> archiveBrowsingEnabled() {
        return this.archiveBrowsingEnabled;
    }
    /**
     * The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
     * an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
     * offline.
     * 
     */
    @Export(name="assumedOfflinePeriodSecs", refs={Integer.class}, tree="[0]")
    private Output<Integer> assumedOfflinePeriodSecs;

    /**
     * @return The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
     * an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
     * offline.
     * 
     */
    public Output<Integer> assumedOfflinePeriodSecs() {
        return this.assumedOfflinePeriodSecs;
    }
    /**
     * (A.K.A &#39;Ignore Repository&#39; on the UI) When set, the repository or its local cache do not participate in artifact
     * resolution.
     * 
     */
    @Export(name="blackedOut", refs={Boolean.class}, tree="[0]")
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
     * If set, artifacts will fail to download if a mismatch is detected between requested and received mimetype, according to
     * the list specified in the system properties file under blockedMismatchingMimeTypes. You can override by adding mimetypes
     * to the override list &#39;mismatching_mime_types_override_list&#39;.
     * 
     */
    @Export(name="blockMismatchingMimeTypes", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> blockMismatchingMimeTypes;

    /**
     * @return If set, artifacts will fail to download if a mismatch is detected between requested and received mimetype, according to
     * the list specified in the system properties file under blockedMismatchingMimeTypes. You can override by adding mimetypes
     * to the override list &#39;mismatching_mime_types_override_list&#39;.
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
    @Export(name="bypassHeadRequests", refs={Boolean.class}, tree="[0]")
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
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;
     * 
     */
    @Export(name="cdnRedirect", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> cdnRedirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;
     * 
     */
    public Output<Boolean> cdnRedirect() {
        return this.cdnRedirect;
    }
    /**
     * Client TLS certificate name.
     * 
     */
    @Export(name="clientTlsCertificate", refs={String.class}, tree="[0]")
    private Output<String> clientTlsCertificate;

    /**
     * @return Client TLS certificate name.
     * 
     */
    public Output<String> clientTlsCertificate() {
        return this.clientTlsCertificate;
    }
    @Export(name="contentSynchronisation", refs={RemoteDebianRepositoryContentSynchronisation.class}, tree="[0]")
    private Output</* @Nullable */ RemoteDebianRepositoryContentSynchronisation> contentSynchronisation;

    public Output<Optional<RemoteDebianRepositoryContentSynchronisation>> contentSynchronisation() {
        return Codegen.optional(this.contentSynchronisation);
    }
    /**
     * Public description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Public description.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set
     * for the Artifactory instance, it will be ignored, too. Introduced since Artifactory 7.41.7.
     * 
     */
    @Export(name="disableProxy", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disableProxy;

    /**
     * @return When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set
     * for the Artifactory instance, it will be ignored, too. Introduced since Artifactory 7.41.7.
     * 
     */
    public Output<Boolean> disableProxy() {
        return this.disableProxy;
    }
    /**
     * Whether to disable URL normalization. Default is `false`.
     * 
     */
    @Export(name="disableUrlNormalization", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disableUrlNormalization;

    /**
     * @return Whether to disable URL normalization. Default is `false`.
     * 
     */
    public Output<Boolean> disableUrlNormalization() {
        return this.disableUrlNormalization;
    }
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;.
     * 
     */
    @Export(name="downloadDirect", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> downloadDirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;.
     * 
     */
    public Output<Boolean> downloadDirect() {
        return this.downloadDirect;
    }
    /**
     * Enables cookie management if the remote repository uses cookies to manage client state.
     * 
     */
    @Export(name="enableCookieManagement", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enableCookieManagement;

    /**
     * @return Enables cookie management if the remote repository uses cookies to manage client state.
     * 
     */
    public Output<Boolean> enableCookieManagement() {
        return this.enableCookieManagement;
    }
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`.By default no
     * artifacts are excluded.
     * 
     */
    @Export(name="excludesPattern", refs={String.class}, tree="[0]")
    private Output<String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`.By default no
     * artifacts are excluded.
     * 
     */
    public Output<String> excludesPattern() {
        return this.excludesPattern;
    }
    /**
     * When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to
     * communicate with this repository.
     * 
     */
    @Export(name="hardFail", refs={Boolean.class}, tree="[0]")
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
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**&#47;*`).
     * 
     */
    @Export(name="includesPattern", refs={String.class}, tree="[0]")
    private Output<String> includesPattern;

    /**
     * @return List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**&#47;*`).
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
    @Export(name="key", refs={String.class}, tree="[0]")
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
     * the &#39;Retrieval Cache Period&#39;. Default value is &#39;false&#39;. This field exists in the API but not in the UI.
     * 
     */
    @Export(name="listRemoteFolderItems", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> listRemoteFolderItems;

    /**
     * @return Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of
     * the &#39;Retrieval Cache Period&#39;. Default value is &#39;false&#39;. This field exists in the API but not in the UI.
     * 
     */
    public Output<Boolean> listRemoteFolderItems() {
        return this.listRemoteFolderItems;
    }
    /**
     * The local address to be used when creating connections. Useful for specifying the interface to use on systems with
     * multiple network interfaces.
     * 
     */
    @Export(name="localAddress", refs={String.class}, tree="[0]")
    private Output<String> localAddress;

    /**
     * @return The local address to be used when creating connections. Useful for specifying the interface to use on systems with
     * multiple network interfaces.
     * 
     */
    public Output<String> localAddress() {
        return this.localAddress;
    }
    /**
     * Metadata Retrieval Cache Timeout (Sec) in the UI.This value refers to the number of seconds to wait for retrieval from
     * the remote before serving locally cached artifact or fail the request.
     * 
     */
    @Export(name="metadataRetrievalTimeoutSecs", refs={Integer.class}, tree="[0]")
    private Output<Integer> metadataRetrievalTimeoutSecs;

    /**
     * @return Metadata Retrieval Cache Timeout (Sec) in the UI.This value refers to the number of seconds to wait for retrieval from
     * the remote before serving locally cached artifact or fail the request.
     * 
     */
    public Output<Integer> metadataRetrievalTimeoutSecs() {
        return this.metadataRetrievalTimeoutSecs;
    }
    /**
     * The set of mime types that should override the block_mismatching_mime_types setting. Eg:
     * &#39;application/json,application/xml&#39;. Default value is empty.
     * 
     */
    @Export(name="mismatchingMimeTypesOverrideList", refs={String.class}, tree="[0]")
    private Output<String> mismatchingMimeTypesOverrideList;

    /**
     * @return The set of mime types that should override the block_mismatching_mime_types setting. Eg:
     * &#39;application/json,application/xml&#39;. Default value is empty.
     * 
     */
    public Output<String> mismatchingMimeTypesOverrideList() {
        return this.mismatchingMimeTypesOverrideList;
    }
    /**
     * Missed Retrieval Cache Period (Sec) in the UI. The number of seconds to cache artifact retrieval misses (artifact not
     * found). A value of 0 indicates no caching.
     * 
     */
    @Export(name="missedCachePeriodSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> missedCachePeriodSeconds;

    /**
     * @return Missed Retrieval Cache Period (Sec) in the UI. The number of seconds to cache artifact retrieval misses (artifact not
     * found). A value of 0 indicates no caching.
     * 
     */
    public Output<Integer> missedCachePeriodSeconds() {
        return this.missedCachePeriodSeconds;
    }
    /**
     * Internal description.
     * 
     */
    @Export(name="notes", refs={String.class}, tree="[0]")
    private Output<String> notes;

    /**
     * @return Internal description.
     * 
     */
    public Output<String> notes() {
        return this.notes;
    }
    /**
     * If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     * 
     */
    @Export(name="offline", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> offline;

    /**
     * @return If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     * 
     */
    public Output<Boolean> offline() {
        return this.offline;
    }
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    @Export(name="priorityResolution", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> priorityResolution;

    /**
     * @return Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    public Output<Boolean> priorityResolution() {
        return this.priorityResolution;
    }
    @Export(name="projectEnvironments", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> projectEnvironments;

    public Output<List<String>> projectEnvironments() {
        return this.projectEnvironments;
    }
    /**
     * Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    @Export(name="projectKey", refs={String.class}, tree="[0]")
    private Output<String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    public Output<String> projectKey() {
        return this.projectKey;
    }
    /**
     * List of property set name
     * 
     */
    @Export(name="propertySets", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> propertySets;

    /**
     * @return List of property set name
     * 
     */
    public Output<Optional<List<String>>> propertySets() {
        return Codegen.optional(this.propertySets);
    }
    /**
     * Proxy key from Artifactory Proxies settings. Can&#39;t be set if `disable_proxy = true`.
     * 
     */
    @Export(name="proxy", refs={String.class}, tree="[0]")
    private Output<String> proxy;

    /**
     * @return Proxy key from Artifactory Proxies settings. Can&#39;t be set if `disable_proxy = true`.
     * 
     */
    public Output<String> proxy() {
        return this.proxy;
    }
    /**
     * Custom HTTP query parameters that will be automatically included in all remote resource requests. For example:
     * `param1=val1&amp;param2=val2&amp;param3=val3`
     * 
     */
    @Export(name="queryParams", refs={String.class}, tree="[0]")
    private Output<String> queryParams;

    /**
     * @return Custom HTTP query parameters that will be automatically included in all remote resource requests. For example:
     * `param1=val1&amp;param2=val2&amp;param3=val3`
     * 
     */
    public Output<String> queryParams() {
        return this.queryParams;
    }
    /**
     * Repository layout key for the remote layout mapping. Repository can be created without this attribute (or set to an
     * empty string). Once it&#39;s set, it can&#39;t be removed by passing an empty string or removing the attribute, that will be
     * ignored by the Artifactory API. UI shows an error message, if the user tries to remove the value.
     * 
     */
    @Export(name="remoteRepoLayoutRef", refs={String.class}, tree="[0]")
    private Output<String> remoteRepoLayoutRef;

    /**
     * @return Repository layout key for the remote layout mapping. Repository can be created without this attribute (or set to an
     * empty string). Once it&#39;s set, it can&#39;t be removed by passing an empty string or removing the attribute, that will be
     * ignored by the Artifactory API. UI shows an error message, if the user tries to remove the value.
     * 
     */
    public Output<String> remoteRepoLayoutRef() {
        return this.remoteRepoLayoutRef;
    }
    /**
     * Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
     * corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
     * 
     */
    @Export(name="repoLayoutRef", refs={String.class}, tree="[0]")
    private Output<String> repoLayoutRef;

    /**
     * @return Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
     * corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
     * 
     */
    public Output<String> repoLayoutRef() {
        return this.repoLayoutRef;
    }
    /**
     * Metadata Retrieval Cache Period (Sec) in the UI. This value refers to the number of seconds to cache metadata files
     * before checking for newer versions on remote server. A value of 0 indicates no caching.
     * 
     */
    @Export(name="retrievalCachePeriodSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> retrievalCachePeriodSeconds;

    /**
     * @return Metadata Retrieval Cache Period (Sec) in the UI. This value refers to the number of seconds to cache metadata files
     * before checking for newer versions on remote server. A value of 0 indicates no caching.
     * 
     */
    public Output<Integer> retrievalCachePeriodSeconds() {
        return this.retrievalCachePeriodSeconds;
    }
    /**
     * @deprecated
     * No longer supported
     * 
     */
    @Deprecated /* No longer supported */
    @Export(name="shareConfiguration", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> shareConfiguration;

    public Output<Boolean> shareConfiguration() {
        return this.shareConfiguration;
    }
    /**
     * Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network
     * operation is considered a retrieval failure.
     * 
     */
    @Export(name="socketTimeoutMillis", refs={Integer.class}, tree="[0]")
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
    @Export(name="storeArtifactsLocally", refs={Boolean.class}, tree="[0]")
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
    @Export(name="synchronizeProperties", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> synchronizeProperties;

    /**
     * @return When set, remote artifacts are fetched along with their properties.
     * 
     */
    public Output<Boolean> synchronizeProperties() {
        return this.synchronizeProperties;
    }
    /**
     * Unused Artifacts Cleanup Period (Hr) in the UI. The number of hours to wait before an artifact is deemed &#39;unused&#39; and
     * eligible for cleanup from the repository. A value of 0 means automatic cleanup of cached artifacts is disabled.
     * 
     */
    @Export(name="unusedArtifactsCleanupPeriodHours", refs={Integer.class}, tree="[0]")
    private Output<Integer> unusedArtifactsCleanupPeriodHours;

    /**
     * @return Unused Artifacts Cleanup Period (Hr) in the UI. The number of hours to wait before an artifact is deemed &#39;unused&#39; and
     * eligible for cleanup from the repository. A value of 0 means automatic cleanup of cached artifacts is disabled.
     * 
     */
    public Output<Integer> unusedArtifactsCleanupPeriodHours() {
        return this.unusedArtifactsCleanupPeriodHours;
    }
    /**
     * The remote repo URL.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The remote repo URL.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    public Output<String> username() {
        return this.username;
    }
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    @Export(name="xrayIndex", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> xrayIndex;

    /**
     * @return Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    public Output<Boolean> xrayIndex() {
        return this.xrayIndex;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RemoteDebianRepository(java.lang.String name) {
        this(name, RemoteDebianRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RemoteDebianRepository(java.lang.String name, RemoteDebianRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RemoteDebianRepository(java.lang.String name, RemoteDebianRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/remoteDebianRepository:RemoteDebianRepository", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RemoteDebianRepository(java.lang.String name, Output<java.lang.String> id, @Nullable RemoteDebianRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/remoteDebianRepository:RemoteDebianRepository", name, state, makeResourceOptions(options, id), false);
    }

    private static RemoteDebianRepositoryArgs makeArgs(RemoteDebianRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RemoteDebianRepositoryArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
    public static RemoteDebianRepository get(java.lang.String name, Output<java.lang.String> id, @Nullable RemoteDebianRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RemoteDebianRepository(name, id, state, options);
    }
}

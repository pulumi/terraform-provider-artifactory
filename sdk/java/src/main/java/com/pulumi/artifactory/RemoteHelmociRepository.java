// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.RemoteHelmociRepositoryArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.RemoteHelmociRepositoryState;
import com.pulumi.artifactory.outputs.RemoteHelmociRepositoryContentSynchronisation;
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
 * Creates remote Helm OCI repository resource.
 * 
 * Official documentation can be found [here](https://jfrog.com/help/r/jfrog-artifactory-documentation/helm-oci-repositories)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.RemoteHelmociRepository;
 * import com.pulumi.artifactory.RemoteHelmociRepositoryArgs;
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
 *         var my_helmoci_remote = new RemoteHelmociRepository(&#34;my-helmoci-remote&#34;, RemoteHelmociRepositoryArgs.builder()        
 *             .enableTokenAuthentication(true)
 *             .externalDependenciesEnabled(true)
 *             .externalDependenciesPatterns(&#34;**{@literal /}registry-1.docker.io/**&#34;)
 *             .key(&#34;my-helmoci-remote&#34;)
 *             .url(&#34;https://registry-1.docker.io/&#34;)
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
 *  $ pulumi import artifactory:index/remoteHelmociRepository:RemoteHelmociRepository my-helmoci-remote my-helmoci-remote
 * ```
 * 
 */
@ResourceType(type="artifactory:index/remoteHelmociRepository:RemoteHelmociRepository")
public class RemoteHelmociRepository extends com.pulumi.resources.CustomResource {
    /**
     * &#39;Lenient Host Authentication&#39; in the UI. Allow credentials of this repository to be used on requests redirected to any
     * other host.
     * 
     */
    @Export(name="allowAnyHostAuth", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowAnyHostAuth;

    /**
     * @return &#39;Lenient Host Authentication&#39; in the UI. Allow credentials of this repository to be used on requests redirected to any
     * other host.
     * 
     */
    public Output<Optional<Boolean>> allowAnyHostAuth() {
        return Codegen.optional(this.allowAnyHostAuth);
    }
    /**
     * The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
     * an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
     * offline.
     * 
     */
    @Export(name="assumedOfflinePeriodSecs", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> assumedOfflinePeriodSecs;

    /**
     * @return The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
     * an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
     * offline.
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
    @Export(name="blackedOut", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> blackedOut;

    /**
     * @return (A.K.A &#39;Ignore Repository&#39; on the UI) When set, the repository or its local cache do not participate in artifact
     * resolution.
     * 
     */
    public Output<Optional<Boolean>> blackedOut() {
        return Codegen.optional(this.blackedOut);
    }
    /**
     * If set, artifacts will fail to download if a mismatch is detected between requested and received mimetype, according to
     * the list specified in the system properties file under blockedMismatchingMimeTypes. You can override by adding mimetypes
     * to the override list &#39;mismatching_mime_types_override_list&#39;.
     * 
     */
    @Export(name="blockMismatchingMimeTypes", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> blockMismatchingMimeTypes;

    /**
     * @return If set, artifacts will fail to download if a mismatch is detected between requested and received mimetype, according to
     * the list specified in the system properties file under blockedMismatchingMimeTypes. You can override by adding mimetypes
     * to the override list &#39;mismatching_mime_types_override_list&#39;.
     * 
     */
    public Output<Optional<Boolean>> blockMismatchingMimeTypes() {
        return Codegen.optional(this.blockMismatchingMimeTypes);
    }
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     * 
     */
    @Export(name="bypassHeadRequests", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> bypassHeadRequests;

    /**
     * @return Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     * 
     */
    public Output<Optional<Boolean>> bypassHeadRequests() {
        return Codegen.optional(this.bypassHeadRequests);
    }
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;
     * 
     */
    @Export(name="cdnRedirect", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> cdnRedirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;
     * 
     */
    public Output<Optional<Boolean>> cdnRedirect() {
        return Codegen.optional(this.cdnRedirect);
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
    @Export(name="contentSynchronisation", refs={RemoteHelmociRepositoryContentSynchronisation.class}, tree="[0]")
    private Output<RemoteHelmociRepositoryContentSynchronisation> contentSynchronisation;

    public Output<RemoteHelmociRepositoryContentSynchronisation> contentSynchronisation() {
        return this.contentSynchronisation;
    }
    /**
     * Public description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Public description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set
     * for the Artifactory instance, it will be ignored, too. Introduced since Artifactory 7.41.7.
     * 
     */
    @Export(name="disableProxy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableProxy;

    /**
     * @return When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set
     * for the Artifactory instance, it will be ignored, too. Introduced since Artifactory 7.41.7.
     * 
     */
    public Output<Optional<Boolean>> disableProxy() {
        return Codegen.optional(this.disableProxy);
    }
    /**
     * Whether to disable URL normalization, default is `false`.
     * 
     */
    @Export(name="disableUrlNormalization", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableUrlNormalization;

    /**
     * @return Whether to disable URL normalization, default is `false`.
     * 
     */
    public Output<Optional<Boolean>> disableUrlNormalization() {
        return Codegen.optional(this.disableUrlNormalization);
    }
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;.
     * 
     */
    @Export(name="downloadDirect", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> downloadDirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;.
     * 
     */
    public Output<Optional<Boolean>> downloadDirect() {
        return Codegen.optional(this.downloadDirect);
    }
    /**
     * Enables cookie management if the remote repository uses cookies to manage client state.
     * 
     */
    @Export(name="enableCookieManagement", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableCookieManagement;

    /**
     * @return Enables cookie management if the remote repository uses cookies to manage client state.
     * 
     */
    public Output<Optional<Boolean>> enableCookieManagement() {
        return Codegen.optional(this.enableCookieManagement);
    }
    /**
     * Enable token (Bearer) based authentication.
     * 
     */
    @Export(name="enableTokenAuthentication", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enableTokenAuthentication;

    /**
     * @return Enable token (Bearer) based authentication.
     * 
     */
    public Output<Boolean> enableTokenAuthentication() {
        return this.enableTokenAuthentication;
    }
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
     * artifacts are excluded.
     * 
     */
    @Export(name="excludesPattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
     * artifacts are excluded.
     * 
     */
    public Output<Optional<String>> excludesPattern() {
        return Codegen.optional(this.excludesPattern);
    }
    /**
     * Also known as &#39;Foreign Layers Caching&#39; on the UI.
     * 
     */
    @Export(name="externalDependenciesEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> externalDependenciesEnabled;

    /**
     * @return Also known as &#39;Foreign Layers Caching&#39; on the UI.
     * 
     */
    public Output<Optional<Boolean>> externalDependenciesEnabled() {
        return Codegen.optional(this.externalDependenciesEnabled);
    }
    /**
     * Optional include patterns to match external URLs. Ant-style path expressions are supported (*, **, ?). For example, specifying `**{@literal /}github.com/**` will only allow downloading foreign layers from github.com host. By default, this is set to `**` in the UI, which means that foreign layers may be downloaded from any external hosts. Due to SDKv2 limitations, we can&#39;t set the default value for the list. This value `**` must be assigned to the attribute manually, if user don&#39;t specify any other non-default values. We don&#39;t want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `**` on update if HCL doesn&#39;t have the attribute set or the list is empty.
     * 
     */
    @Export(name="externalDependenciesPatterns", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> externalDependenciesPatterns;

    /**
     * @return Optional include patterns to match external URLs. Ant-style path expressions are supported (*, **, ?). For example, specifying `**{@literal /}github.com/**` will only allow downloading foreign layers from github.com host. By default, this is set to `**` in the UI, which means that foreign layers may be downloaded from any external hosts. Due to SDKv2 limitations, we can&#39;t set the default value for the list. This value `**` must be assigned to the attribute manually, if user don&#39;t specify any other non-default values. We don&#39;t want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `**` on update if HCL doesn&#39;t have the attribute set or the list is empty.
     * 
     */
    public Output<Optional<List<String>>> externalDependenciesPatterns() {
        return Codegen.optional(this.externalDependenciesPatterns);
    }
    /**
     * When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to
     * communicate with this repository.
     * 
     */
    @Export(name="hardFail", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hardFail;

    /**
     * @return When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to
     * communicate with this repository.
     * 
     */
    public Output<Optional<Boolean>> hardFail() {
        return Codegen.optional(this.hardFail);
    }
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Export(name="includesPattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> includesPattern;

    /**
     * @return List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    public Output<Optional<String>> includesPattern() {
        return Codegen.optional(this.includesPattern);
    }
    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of
     * the &#39;Retrieval Cache Period&#39;. Default value is &#39;true&#39;.
     * 
     */
    @Export(name="listRemoteFolderItems", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> listRemoteFolderItems;

    /**
     * @return Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of
     * the &#39;Retrieval Cache Period&#39;. Default value is &#39;true&#39;.
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
    @Export(name="localAddress", refs={String.class}, tree="[0]")
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
     * Metadata Retrieval Cache Timeout (Sec) in the UI.This value refers to the number of seconds to wait for retrieval from
     * the remote before serving locally cached artifact or fail the request.
     * 
     */
    @Export(name="metadataRetrievalTimeoutSecs", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> metadataRetrievalTimeoutSecs;

    /**
     * @return Metadata Retrieval Cache Timeout (Sec) in the UI.This value refers to the number of seconds to wait for retrieval from
     * the remote before serving locally cached artifact or fail the request.
     * 
     */
    public Output<Optional<Integer>> metadataRetrievalTimeoutSecs() {
        return Codegen.optional(this.metadataRetrievalTimeoutSecs);
    }
    /**
     * The set of mime types that should override the block_mismatching_mime_types setting. Eg:
     * &#39;application/json,application/xml&#39;. Default value is empty.
     * 
     */
    @Export(name="mismatchingMimeTypesOverrideList", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mismatchingMimeTypesOverrideList;

    /**
     * @return The set of mime types that should override the block_mismatching_mime_types setting. Eg:
     * &#39;application/json,application/xml&#39;. Default value is empty.
     * 
     */
    public Output<Optional<String>> mismatchingMimeTypesOverrideList() {
        return Codegen.optional(this.mismatchingMimeTypesOverrideList);
    }
    /**
     * Missed Retrieval Cache Period (Sec) in the UI. The number of seconds to cache artifact retrieval misses (artifact not
     * found). A value of 0 indicates no caching.
     * 
     */
    @Export(name="missedCachePeriodSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> missedCachePeriodSeconds;

    /**
     * @return Missed Retrieval Cache Period (Sec) in the UI. The number of seconds to cache artifact retrieval misses (artifact not
     * found). A value of 0 indicates no caching.
     * 
     */
    public Output<Optional<Integer>> missedCachePeriodSeconds() {
        return Codegen.optional(this.missedCachePeriodSeconds);
    }
    /**
     * Internal description.
     * 
     */
    @Export(name="notes", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> notes;

    /**
     * @return Internal description.
     * 
     */
    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    /**
     * If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     * 
     */
    @Export(name="offline", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> offline;

    /**
     * @return If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     * 
     */
    public Output<Optional<Boolean>> offline() {
        return Codegen.optional(this.offline);
    }
    @Export(name="packageType", refs={String.class}, tree="[0]")
    private Output<String> packageType;

    public Output<String> packageType() {
        return this.packageType;
    }
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Setting Priority Resolution takes precedence over the resolution order when resolving virtual repositories. Setting
     * repositories with priority will cause metadata to be merged only from repositories set with a priority. If a package is
     * not found in those repositories, Artifactory will merge from repositories marked as non-priority.
     * 
     */
    @Export(name="priorityResolution", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> priorityResolution;

    /**
     * @return Setting Priority Resolution takes precedence over the resolution order when resolving virtual repositories. Setting
     * repositories with priority will cause metadata to be merged only from repositories set with a priority. If a package is
     * not found in those repositories, Artifactory will merge from repositories marked as non-priority.
     * 
     */
    public Output<Optional<Boolean>> priorityResolution() {
        return Codegen.optional(this.priorityResolution);
    }
    /**
     * Project environment for assigning this repository to. Allow values: &#34;DEV&#34;, &#34;PROD&#34;, or one of custom environment. Before
     * Artifactory 7.53.1, up to 2 values (&#34;DEV&#34; and &#34;PROD&#34;) are allowed. From 7.53.1 onward, only one value is allowed. The
     * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
     * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
     * 
     */
    @Export(name="projectEnvironments", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> projectEnvironments;

    /**
     * @return Project environment for assigning this repository to. Allow values: &#34;DEV&#34;, &#34;PROD&#34;, or one of custom environment. Before
     * Artifactory 7.53.1, up to 2 values (&#34;DEV&#34; and &#34;PROD&#34;) are allowed. From 7.53.1 onward, only one value is allowed. The
     * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
     * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
     * 
     */
    public Output<List<String>> projectEnvironments() {
        return this.projectEnvironments;
    }
    /**
     * Use this attribute to enter your GCR, GAR Project Id to limit the scope of this remote repo to a specific project in your third-party registry. When leaving this field blank or unset, remote repositories that support project id will default to their default project as you have set up in your account.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectId;

    /**
     * @return Use this attribute to enter your GCR, GAR Project Id to limit the scope of this remote repo to a specific project in your third-party registry. When leaving this field blank or unset, remote repositories that support project id will default to their default project as you have set up in your account.
     * 
     */
    public Output<Optional<String>> projectId() {
        return Codegen.optional(this.projectId);
    }
    /**
     * Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    @Export(name="projectKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    public Output<Optional<String>> projectKey() {
        return Codegen.optional(this.projectKey);
    }
    /**
     * List of property set names
     * 
     */
    @Export(name="propertySets", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> propertySets;

    /**
     * @return List of property set names
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
    private Output</* @Nullable */ String> proxy;

    /**
     * @return Proxy key from Artifactory Proxies settings. Can&#39;t be set if `disable_proxy = true`.
     * 
     */
    public Output<Optional<String>> proxy() {
        return Codegen.optional(this.proxy);
    }
    /**
     * Custom HTTP query parameters that will be automatically included in all remote resource requests. For example:
     * `param1=val1&amp;param2=val2&amp;param3=val3`
     * 
     */
    @Export(name="queryParams", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> queryParams;

    /**
     * @return Custom HTTP query parameters that will be automatically included in all remote resource requests. For example:
     * `param1=val1&amp;param2=val2&amp;param3=val3`
     * 
     */
    public Output<Optional<String>> queryParams() {
        return Codegen.optional(this.queryParams);
    }
    /**
     * Repository layout key for the remote layout mapping. Repository can be created without this attribute (or set to an
     * empty string). Once it&#39;s set, it can&#39;t be removed by passing an empty string or removing the attribute, that will be
     * ignored by the Artifactory API. UI shows an error message, if the user tries to remove the value.
     * 
     */
    @Export(name="remoteRepoLayoutRef", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> remoteRepoLayoutRef;

    /**
     * @return Repository layout key for the remote layout mapping. Repository can be created without this attribute (or set to an
     * empty string). Once it&#39;s set, it can&#39;t be removed by passing an empty string or removing the attribute, that will be
     * ignored by the Artifactory API. UI shows an error message, if the user tries to remove the value.
     * 
     */
    public Output<Optional<String>> remoteRepoLayoutRef() {
        return Codegen.optional(this.remoteRepoLayoutRef);
    }
    /**
     * Repository layout key for the remote repository
     * 
     */
    @Export(name="repoLayoutRef", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> repoLayoutRef;

    /**
     * @return Repository layout key for the remote repository
     * 
     */
    public Output<Optional<String>> repoLayoutRef() {
        return Codegen.optional(this.repoLayoutRef);
    }
    /**
     * Metadata Retrieval Cache Period (Sec) in the UI. This value refers to the number of seconds to cache metadata files
     * before checking for newer versions on remote server. A value of 0 indicates no caching.
     * 
     */
    @Export(name="retrievalCachePeriodSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> retrievalCachePeriodSeconds;

    /**
     * @return Metadata Retrieval Cache Period (Sec) in the UI. This value refers to the number of seconds to cache metadata files
     * before checking for newer versions on remote server. A value of 0 indicates no caching.
     * 
     */
    public Output<Optional<Integer>> retrievalCachePeriodSeconds() {
        return Codegen.optional(this.retrievalCachePeriodSeconds);
    }
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
    private Output</* @Nullable */ Integer> socketTimeoutMillis;

    /**
     * @return Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network
     * operation is considered a retrieval failure.
     * 
     */
    public Output<Optional<Integer>> socketTimeoutMillis() {
        return Codegen.optional(this.socketTimeoutMillis);
    }
    /**
     * When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
     * direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
     * one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
     * servers.
     * 
     */
    @Export(name="storeArtifactsLocally", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> storeArtifactsLocally;

    /**
     * @return When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
     * direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
     * one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
     * servers.
     * 
     */
    public Output<Optional<Boolean>> storeArtifactsLocally() {
        return Codegen.optional(this.storeArtifactsLocally);
    }
    /**
     * When set, remote artifacts are fetched along with their properties.
     * 
     */
    @Export(name="synchronizeProperties", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> synchronizeProperties;

    /**
     * @return When set, remote artifacts are fetched along with their properties.
     * 
     */
    public Output<Optional<Boolean>> synchronizeProperties() {
        return Codegen.optional(this.synchronizeProperties);
    }
    /**
     * Unused Artifacts Cleanup Period (Hr) in the UI. The number of hours to wait before an artifact is deemed &#39;unused&#39; and
     * eligible for cleanup from the repository. A value of 0 means automatic cleanup of cached artifacts is disabled.
     * 
     */
    @Export(name="unusedArtifactsCleanupPeriodHours", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> unusedArtifactsCleanupPeriodHours;

    /**
     * @return Unused Artifacts Cleanup Period (Hr) in the UI. The number of hours to wait before an artifact is deemed &#39;unused&#39; and
     * eligible for cleanup from the repository. A value of 0 means automatic cleanup of cached artifacts is disabled.
     * 
     */
    public Output<Optional<Integer>> unusedArtifactsCleanupPeriodHours() {
        return Codegen.optional(this.unusedArtifactsCleanupPeriodHours);
    }
    /**
     * This is a URL to the remote registry. Consider using HTTPS to ensure a secure connection.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return This is a URL to the remote registry. Consider using HTTPS to ensure a secure connection.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> username;

    public Output<Optional<String>> username() {
        return Codegen.optional(this.username);
    }
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    @Export(name="xrayIndex", refs={Boolean.class}, tree="[0]")
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
    public RemoteHelmociRepository(String name) {
        this(name, RemoteHelmociRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RemoteHelmociRepository(String name, RemoteHelmociRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RemoteHelmociRepository(String name, RemoteHelmociRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/remoteHelmociRepository:RemoteHelmociRepository", name, args == null ? RemoteHelmociRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RemoteHelmociRepository(String name, Output<String> id, @Nullable RemoteHelmociRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/remoteHelmociRepository:RemoteHelmociRepository", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static RemoteHelmociRepository get(String name, Output<String> id, @Nullable RemoteHelmociRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RemoteHelmociRepository(name, id, state, options);
    }
}

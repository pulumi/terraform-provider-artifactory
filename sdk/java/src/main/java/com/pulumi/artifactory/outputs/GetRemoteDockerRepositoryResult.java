// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.artifactory.outputs.GetRemoteDockerRepositoryContentSynchronisation;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRemoteDockerRepositoryResult {
    private @Nullable Boolean allowAnyHostAuth;
    private @Nullable Integer assumedOfflinePeriodSecs;
    private @Nullable Boolean blackedOut;
    private @Nullable Boolean blockMismatchingMimeTypes;
    /**
     * @return (Optional) When set, Artifactory will block the pulling of Docker images with manifest v2 schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1 that exist in the cache.
     * 
     */
    private Boolean blockPushingSchema1;
    private @Nullable Boolean bypassHeadRequests;
    private @Nullable Boolean cdnRedirect;
    private String clientTlsCertificate;
    private GetRemoteDockerRepositoryContentSynchronisation contentSynchronisation;
    private @Nullable Boolean curated;
    private @Nullable String description;
    private @Nullable Boolean disableProxy;
    /**
     * @return (Optional) Whether to disable URL normalization.
     * 
     */
    private @Nullable Boolean disableUrlNormalization;
    private @Nullable Boolean downloadDirect;
    private @Nullable Boolean enableCookieManagement;
    /**
     * @return (Optional) Enable token (Bearer) based authentication.
     * 
     */
    private Boolean enableTokenAuthentication;
    private @Nullable String excludesPattern;
    /**
     * @return (Optional) Also known as &#39;Foreign Layers Caching&#39; on the UI.
     * 
     */
    private @Nullable Boolean externalDependenciesEnabled;
    /**
     * @return (Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download remote modules from, when presented with &#39;go-import&#39; meta tags in the remote repository response. By default, this is set to `[**]` in the UI, which means that remote modules may be downloaded from any external VCS source. Due to SDKv2 limitations, we can&#39;t set the default value for the list. This value `[**]` must be assigned to the attribute manually, if user don&#39;t specify any other non-default values. We don&#39;t want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `[**]` on update if HCL doesn&#39;t have the attribute set or the list is empty.
     * 
     */
    private @Nullable List<String> externalDependenciesPatterns;
    private @Nullable Boolean hardFail;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String includesPattern;
    private String key;
    private @Nullable Boolean listRemoteFolderItems;
    private @Nullable String localAddress;
    private @Nullable Integer metadataRetrievalTimeoutSecs;
    private @Nullable String mismatchingMimeTypesOverrideList;
    private @Nullable Integer missedCachePeriodSeconds;
    private @Nullable String notes;
    private @Nullable Boolean offline;
    private String packageType;
    private @Nullable String password;
    private @Nullable Boolean priorityResolution;
    private List<String> projectEnvironments;
    private @Nullable String projectId;
    private @Nullable String projectKey;
    private @Nullable List<String> propertySets;
    private @Nullable String proxy;
    private @Nullable String queryParams;
    private @Nullable String remoteRepoLayoutRef;
    private @Nullable String repoLayoutRef;
    private @Nullable Integer retrievalCachePeriodSeconds;
    private Boolean shareConfiguration;
    private @Nullable Integer socketTimeoutMillis;
    private @Nullable Boolean storeArtifactsLocally;
    private @Nullable Boolean synchronizeProperties;
    private @Nullable Integer unusedArtifactsCleanupPeriodHours;
    private @Nullable String url;
    private @Nullable String username;
    private @Nullable Boolean xrayIndex;

    private GetRemoteDockerRepositoryResult() {}
    public Optional<Boolean> allowAnyHostAuth() {
        return Optional.ofNullable(this.allowAnyHostAuth);
    }
    public Optional<Integer> assumedOfflinePeriodSecs() {
        return Optional.ofNullable(this.assumedOfflinePeriodSecs);
    }
    public Optional<Boolean> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }
    public Optional<Boolean> blockMismatchingMimeTypes() {
        return Optional.ofNullable(this.blockMismatchingMimeTypes);
    }
    /**
     * @return (Optional) When set, Artifactory will block the pulling of Docker images with manifest v2 schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1 that exist in the cache.
     * 
     */
    public Boolean blockPushingSchema1() {
        return this.blockPushingSchema1;
    }
    public Optional<Boolean> bypassHeadRequests() {
        return Optional.ofNullable(this.bypassHeadRequests);
    }
    public Optional<Boolean> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }
    public String clientTlsCertificate() {
        return this.clientTlsCertificate;
    }
    public GetRemoteDockerRepositoryContentSynchronisation contentSynchronisation() {
        return this.contentSynchronisation;
    }
    public Optional<Boolean> curated() {
        return Optional.ofNullable(this.curated);
    }
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    public Optional<Boolean> disableProxy() {
        return Optional.ofNullable(this.disableProxy);
    }
    /**
     * @return (Optional) Whether to disable URL normalization.
     * 
     */
    public Optional<Boolean> disableUrlNormalization() {
        return Optional.ofNullable(this.disableUrlNormalization);
    }
    public Optional<Boolean> downloadDirect() {
        return Optional.ofNullable(this.downloadDirect);
    }
    public Optional<Boolean> enableCookieManagement() {
        return Optional.ofNullable(this.enableCookieManagement);
    }
    /**
     * @return (Optional) Enable token (Bearer) based authentication.
     * 
     */
    public Boolean enableTokenAuthentication() {
        return this.enableTokenAuthentication;
    }
    public Optional<String> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }
    /**
     * @return (Optional) Also known as &#39;Foreign Layers Caching&#39; on the UI.
     * 
     */
    public Optional<Boolean> externalDependenciesEnabled() {
        return Optional.ofNullable(this.externalDependenciesEnabled);
    }
    /**
     * @return (Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download remote modules from, when presented with &#39;go-import&#39; meta tags in the remote repository response. By default, this is set to `[**]` in the UI, which means that remote modules may be downloaded from any external VCS source. Due to SDKv2 limitations, we can&#39;t set the default value for the list. This value `[**]` must be assigned to the attribute manually, if user don&#39;t specify any other non-default values. We don&#39;t want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `[**]` on update if HCL doesn&#39;t have the attribute set or the list is empty.
     * 
     */
    public List<String> externalDependenciesPatterns() {
        return this.externalDependenciesPatterns == null ? List.of() : this.externalDependenciesPatterns;
    }
    public Optional<Boolean> hardFail() {
        return Optional.ofNullable(this.hardFail);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }
    public String key() {
        return this.key;
    }
    public Optional<Boolean> listRemoteFolderItems() {
        return Optional.ofNullable(this.listRemoteFolderItems);
    }
    public Optional<String> localAddress() {
        return Optional.ofNullable(this.localAddress);
    }
    public Optional<Integer> metadataRetrievalTimeoutSecs() {
        return Optional.ofNullable(this.metadataRetrievalTimeoutSecs);
    }
    public Optional<String> mismatchingMimeTypesOverrideList() {
        return Optional.ofNullable(this.mismatchingMimeTypesOverrideList);
    }
    public Optional<Integer> missedCachePeriodSeconds() {
        return Optional.ofNullable(this.missedCachePeriodSeconds);
    }
    public Optional<String> notes() {
        return Optional.ofNullable(this.notes);
    }
    public Optional<Boolean> offline() {
        return Optional.ofNullable(this.offline);
    }
    public String packageType() {
        return this.packageType;
    }
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    public Optional<Boolean> priorityResolution() {
        return Optional.ofNullable(this.priorityResolution);
    }
    public List<String> projectEnvironments() {
        return this.projectEnvironments;
    }
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }
    public Optional<String> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }
    public List<String> propertySets() {
        return this.propertySets == null ? List.of() : this.propertySets;
    }
    public Optional<String> proxy() {
        return Optional.ofNullable(this.proxy);
    }
    public Optional<String> queryParams() {
        return Optional.ofNullable(this.queryParams);
    }
    public Optional<String> remoteRepoLayoutRef() {
        return Optional.ofNullable(this.remoteRepoLayoutRef);
    }
    public Optional<String> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }
    public Optional<Integer> retrievalCachePeriodSeconds() {
        return Optional.ofNullable(this.retrievalCachePeriodSeconds);
    }
    public Boolean shareConfiguration() {
        return this.shareConfiguration;
    }
    public Optional<Integer> socketTimeoutMillis() {
        return Optional.ofNullable(this.socketTimeoutMillis);
    }
    public Optional<Boolean> storeArtifactsLocally() {
        return Optional.ofNullable(this.storeArtifactsLocally);
    }
    public Optional<Boolean> synchronizeProperties() {
        return Optional.ofNullable(this.synchronizeProperties);
    }
    public Optional<Integer> unusedArtifactsCleanupPeriodHours() {
        return Optional.ofNullable(this.unusedArtifactsCleanupPeriodHours);
    }
    public Optional<String> url() {
        return Optional.ofNullable(this.url);
    }
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }
    public Optional<Boolean> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRemoteDockerRepositoryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean allowAnyHostAuth;
        private @Nullable Integer assumedOfflinePeriodSecs;
        private @Nullable Boolean blackedOut;
        private @Nullable Boolean blockMismatchingMimeTypes;
        private Boolean blockPushingSchema1;
        private @Nullable Boolean bypassHeadRequests;
        private @Nullable Boolean cdnRedirect;
        private String clientTlsCertificate;
        private GetRemoteDockerRepositoryContentSynchronisation contentSynchronisation;
        private @Nullable Boolean curated;
        private @Nullable String description;
        private @Nullable Boolean disableProxy;
        private @Nullable Boolean disableUrlNormalization;
        private @Nullable Boolean downloadDirect;
        private @Nullable Boolean enableCookieManagement;
        private Boolean enableTokenAuthentication;
        private @Nullable String excludesPattern;
        private @Nullable Boolean externalDependenciesEnabled;
        private @Nullable List<String> externalDependenciesPatterns;
        private @Nullable Boolean hardFail;
        private String id;
        private @Nullable String includesPattern;
        private String key;
        private @Nullable Boolean listRemoteFolderItems;
        private @Nullable String localAddress;
        private @Nullable Integer metadataRetrievalTimeoutSecs;
        private @Nullable String mismatchingMimeTypesOverrideList;
        private @Nullable Integer missedCachePeriodSeconds;
        private @Nullable String notes;
        private @Nullable Boolean offline;
        private String packageType;
        private @Nullable String password;
        private @Nullable Boolean priorityResolution;
        private List<String> projectEnvironments;
        private @Nullable String projectId;
        private @Nullable String projectKey;
        private @Nullable List<String> propertySets;
        private @Nullable String proxy;
        private @Nullable String queryParams;
        private @Nullable String remoteRepoLayoutRef;
        private @Nullable String repoLayoutRef;
        private @Nullable Integer retrievalCachePeriodSeconds;
        private Boolean shareConfiguration;
        private @Nullable Integer socketTimeoutMillis;
        private @Nullable Boolean storeArtifactsLocally;
        private @Nullable Boolean synchronizeProperties;
        private @Nullable Integer unusedArtifactsCleanupPeriodHours;
        private @Nullable String url;
        private @Nullable String username;
        private @Nullable Boolean xrayIndex;
        public Builder() {}
        public Builder(GetRemoteDockerRepositoryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowAnyHostAuth = defaults.allowAnyHostAuth;
    	      this.assumedOfflinePeriodSecs = defaults.assumedOfflinePeriodSecs;
    	      this.blackedOut = defaults.blackedOut;
    	      this.blockMismatchingMimeTypes = defaults.blockMismatchingMimeTypes;
    	      this.blockPushingSchema1 = defaults.blockPushingSchema1;
    	      this.bypassHeadRequests = defaults.bypassHeadRequests;
    	      this.cdnRedirect = defaults.cdnRedirect;
    	      this.clientTlsCertificate = defaults.clientTlsCertificate;
    	      this.contentSynchronisation = defaults.contentSynchronisation;
    	      this.curated = defaults.curated;
    	      this.description = defaults.description;
    	      this.disableProxy = defaults.disableProxy;
    	      this.disableUrlNormalization = defaults.disableUrlNormalization;
    	      this.downloadDirect = defaults.downloadDirect;
    	      this.enableCookieManagement = defaults.enableCookieManagement;
    	      this.enableTokenAuthentication = defaults.enableTokenAuthentication;
    	      this.excludesPattern = defaults.excludesPattern;
    	      this.externalDependenciesEnabled = defaults.externalDependenciesEnabled;
    	      this.externalDependenciesPatterns = defaults.externalDependenciesPatterns;
    	      this.hardFail = defaults.hardFail;
    	      this.id = defaults.id;
    	      this.includesPattern = defaults.includesPattern;
    	      this.key = defaults.key;
    	      this.listRemoteFolderItems = defaults.listRemoteFolderItems;
    	      this.localAddress = defaults.localAddress;
    	      this.metadataRetrievalTimeoutSecs = defaults.metadataRetrievalTimeoutSecs;
    	      this.mismatchingMimeTypesOverrideList = defaults.mismatchingMimeTypesOverrideList;
    	      this.missedCachePeriodSeconds = defaults.missedCachePeriodSeconds;
    	      this.notes = defaults.notes;
    	      this.offline = defaults.offline;
    	      this.packageType = defaults.packageType;
    	      this.password = defaults.password;
    	      this.priorityResolution = defaults.priorityResolution;
    	      this.projectEnvironments = defaults.projectEnvironments;
    	      this.projectId = defaults.projectId;
    	      this.projectKey = defaults.projectKey;
    	      this.propertySets = defaults.propertySets;
    	      this.proxy = defaults.proxy;
    	      this.queryParams = defaults.queryParams;
    	      this.remoteRepoLayoutRef = defaults.remoteRepoLayoutRef;
    	      this.repoLayoutRef = defaults.repoLayoutRef;
    	      this.retrievalCachePeriodSeconds = defaults.retrievalCachePeriodSeconds;
    	      this.shareConfiguration = defaults.shareConfiguration;
    	      this.socketTimeoutMillis = defaults.socketTimeoutMillis;
    	      this.storeArtifactsLocally = defaults.storeArtifactsLocally;
    	      this.synchronizeProperties = defaults.synchronizeProperties;
    	      this.unusedArtifactsCleanupPeriodHours = defaults.unusedArtifactsCleanupPeriodHours;
    	      this.url = defaults.url;
    	      this.username = defaults.username;
    	      this.xrayIndex = defaults.xrayIndex;
        }

        @CustomType.Setter
        public Builder allowAnyHostAuth(@Nullable Boolean allowAnyHostAuth) {

            this.allowAnyHostAuth = allowAnyHostAuth;
            return this;
        }
        @CustomType.Setter
        public Builder assumedOfflinePeriodSecs(@Nullable Integer assumedOfflinePeriodSecs) {

            this.assumedOfflinePeriodSecs = assumedOfflinePeriodSecs;
            return this;
        }
        @CustomType.Setter
        public Builder blackedOut(@Nullable Boolean blackedOut) {

            this.blackedOut = blackedOut;
            return this;
        }
        @CustomType.Setter
        public Builder blockMismatchingMimeTypes(@Nullable Boolean blockMismatchingMimeTypes) {

            this.blockMismatchingMimeTypes = blockMismatchingMimeTypes;
            return this;
        }
        @CustomType.Setter
        public Builder blockPushingSchema1(Boolean blockPushingSchema1) {
            if (blockPushingSchema1 == null) {
              throw new MissingRequiredPropertyException("GetRemoteDockerRepositoryResult", "blockPushingSchema1");
            }
            this.blockPushingSchema1 = blockPushingSchema1;
            return this;
        }
        @CustomType.Setter
        public Builder bypassHeadRequests(@Nullable Boolean bypassHeadRequests) {

            this.bypassHeadRequests = bypassHeadRequests;
            return this;
        }
        @CustomType.Setter
        public Builder cdnRedirect(@Nullable Boolean cdnRedirect) {

            this.cdnRedirect = cdnRedirect;
            return this;
        }
        @CustomType.Setter
        public Builder clientTlsCertificate(String clientTlsCertificate) {
            if (clientTlsCertificate == null) {
              throw new MissingRequiredPropertyException("GetRemoteDockerRepositoryResult", "clientTlsCertificate");
            }
            this.clientTlsCertificate = clientTlsCertificate;
            return this;
        }
        @CustomType.Setter
        public Builder contentSynchronisation(GetRemoteDockerRepositoryContentSynchronisation contentSynchronisation) {
            if (contentSynchronisation == null) {
              throw new MissingRequiredPropertyException("GetRemoteDockerRepositoryResult", "contentSynchronisation");
            }
            this.contentSynchronisation = contentSynchronisation;
            return this;
        }
        @CustomType.Setter
        public Builder curated(@Nullable Boolean curated) {

            this.curated = curated;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder disableProxy(@Nullable Boolean disableProxy) {

            this.disableProxy = disableProxy;
            return this;
        }
        @CustomType.Setter
        public Builder disableUrlNormalization(@Nullable Boolean disableUrlNormalization) {

            this.disableUrlNormalization = disableUrlNormalization;
            return this;
        }
        @CustomType.Setter
        public Builder downloadDirect(@Nullable Boolean downloadDirect) {

            this.downloadDirect = downloadDirect;
            return this;
        }
        @CustomType.Setter
        public Builder enableCookieManagement(@Nullable Boolean enableCookieManagement) {

            this.enableCookieManagement = enableCookieManagement;
            return this;
        }
        @CustomType.Setter
        public Builder enableTokenAuthentication(Boolean enableTokenAuthentication) {
            if (enableTokenAuthentication == null) {
              throw new MissingRequiredPropertyException("GetRemoteDockerRepositoryResult", "enableTokenAuthentication");
            }
            this.enableTokenAuthentication = enableTokenAuthentication;
            return this;
        }
        @CustomType.Setter
        public Builder excludesPattern(@Nullable String excludesPattern) {

            this.excludesPattern = excludesPattern;
            return this;
        }
        @CustomType.Setter
        public Builder externalDependenciesEnabled(@Nullable Boolean externalDependenciesEnabled) {

            this.externalDependenciesEnabled = externalDependenciesEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder externalDependenciesPatterns(@Nullable List<String> externalDependenciesPatterns) {

            this.externalDependenciesPatterns = externalDependenciesPatterns;
            return this;
        }
        public Builder externalDependenciesPatterns(String... externalDependenciesPatterns) {
            return externalDependenciesPatterns(List.of(externalDependenciesPatterns));
        }
        @CustomType.Setter
        public Builder hardFail(@Nullable Boolean hardFail) {

            this.hardFail = hardFail;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetRemoteDockerRepositoryResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder includesPattern(@Nullable String includesPattern) {

            this.includesPattern = includesPattern;
            return this;
        }
        @CustomType.Setter
        public Builder key(String key) {
            if (key == null) {
              throw new MissingRequiredPropertyException("GetRemoteDockerRepositoryResult", "key");
            }
            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder listRemoteFolderItems(@Nullable Boolean listRemoteFolderItems) {

            this.listRemoteFolderItems = listRemoteFolderItems;
            return this;
        }
        @CustomType.Setter
        public Builder localAddress(@Nullable String localAddress) {

            this.localAddress = localAddress;
            return this;
        }
        @CustomType.Setter
        public Builder metadataRetrievalTimeoutSecs(@Nullable Integer metadataRetrievalTimeoutSecs) {

            this.metadataRetrievalTimeoutSecs = metadataRetrievalTimeoutSecs;
            return this;
        }
        @CustomType.Setter
        public Builder mismatchingMimeTypesOverrideList(@Nullable String mismatchingMimeTypesOverrideList) {

            this.mismatchingMimeTypesOverrideList = mismatchingMimeTypesOverrideList;
            return this;
        }
        @CustomType.Setter
        public Builder missedCachePeriodSeconds(@Nullable Integer missedCachePeriodSeconds) {

            this.missedCachePeriodSeconds = missedCachePeriodSeconds;
            return this;
        }
        @CustomType.Setter
        public Builder notes(@Nullable String notes) {

            this.notes = notes;
            return this;
        }
        @CustomType.Setter
        public Builder offline(@Nullable Boolean offline) {

            this.offline = offline;
            return this;
        }
        @CustomType.Setter
        public Builder packageType(String packageType) {
            if (packageType == null) {
              throw new MissingRequiredPropertyException("GetRemoteDockerRepositoryResult", "packageType");
            }
            this.packageType = packageType;
            return this;
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {

            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder priorityResolution(@Nullable Boolean priorityResolution) {

            this.priorityResolution = priorityResolution;
            return this;
        }
        @CustomType.Setter
        public Builder projectEnvironments(List<String> projectEnvironments) {
            if (projectEnvironments == null) {
              throw new MissingRequiredPropertyException("GetRemoteDockerRepositoryResult", "projectEnvironments");
            }
            this.projectEnvironments = projectEnvironments;
            return this;
        }
        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }
        @CustomType.Setter
        public Builder projectId(@Nullable String projectId) {

            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder projectKey(@Nullable String projectKey) {

            this.projectKey = projectKey;
            return this;
        }
        @CustomType.Setter
        public Builder propertySets(@Nullable List<String> propertySets) {

            this.propertySets = propertySets;
            return this;
        }
        public Builder propertySets(String... propertySets) {
            return propertySets(List.of(propertySets));
        }
        @CustomType.Setter
        public Builder proxy(@Nullable String proxy) {

            this.proxy = proxy;
            return this;
        }
        @CustomType.Setter
        public Builder queryParams(@Nullable String queryParams) {

            this.queryParams = queryParams;
            return this;
        }
        @CustomType.Setter
        public Builder remoteRepoLayoutRef(@Nullable String remoteRepoLayoutRef) {

            this.remoteRepoLayoutRef = remoteRepoLayoutRef;
            return this;
        }
        @CustomType.Setter
        public Builder repoLayoutRef(@Nullable String repoLayoutRef) {

            this.repoLayoutRef = repoLayoutRef;
            return this;
        }
        @CustomType.Setter
        public Builder retrievalCachePeriodSeconds(@Nullable Integer retrievalCachePeriodSeconds) {

            this.retrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
            return this;
        }
        @CustomType.Setter
        public Builder shareConfiguration(Boolean shareConfiguration) {
            if (shareConfiguration == null) {
              throw new MissingRequiredPropertyException("GetRemoteDockerRepositoryResult", "shareConfiguration");
            }
            this.shareConfiguration = shareConfiguration;
            return this;
        }
        @CustomType.Setter
        public Builder socketTimeoutMillis(@Nullable Integer socketTimeoutMillis) {

            this.socketTimeoutMillis = socketTimeoutMillis;
            return this;
        }
        @CustomType.Setter
        public Builder storeArtifactsLocally(@Nullable Boolean storeArtifactsLocally) {

            this.storeArtifactsLocally = storeArtifactsLocally;
            return this;
        }
        @CustomType.Setter
        public Builder synchronizeProperties(@Nullable Boolean synchronizeProperties) {

            this.synchronizeProperties = synchronizeProperties;
            return this;
        }
        @CustomType.Setter
        public Builder unusedArtifactsCleanupPeriodHours(@Nullable Integer unusedArtifactsCleanupPeriodHours) {

            this.unusedArtifactsCleanupPeriodHours = unusedArtifactsCleanupPeriodHours;
            return this;
        }
        @CustomType.Setter
        public Builder url(@Nullable String url) {

            this.url = url;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {

            this.username = username;
            return this;
        }
        @CustomType.Setter
        public Builder xrayIndex(@Nullable Boolean xrayIndex) {

            this.xrayIndex = xrayIndex;
            return this;
        }
        public GetRemoteDockerRepositoryResult build() {
            final var _resultValue = new GetRemoteDockerRepositoryResult();
            _resultValue.allowAnyHostAuth = allowAnyHostAuth;
            _resultValue.assumedOfflinePeriodSecs = assumedOfflinePeriodSecs;
            _resultValue.blackedOut = blackedOut;
            _resultValue.blockMismatchingMimeTypes = blockMismatchingMimeTypes;
            _resultValue.blockPushingSchema1 = blockPushingSchema1;
            _resultValue.bypassHeadRequests = bypassHeadRequests;
            _resultValue.cdnRedirect = cdnRedirect;
            _resultValue.clientTlsCertificate = clientTlsCertificate;
            _resultValue.contentSynchronisation = contentSynchronisation;
            _resultValue.curated = curated;
            _resultValue.description = description;
            _resultValue.disableProxy = disableProxy;
            _resultValue.disableUrlNormalization = disableUrlNormalization;
            _resultValue.downloadDirect = downloadDirect;
            _resultValue.enableCookieManagement = enableCookieManagement;
            _resultValue.enableTokenAuthentication = enableTokenAuthentication;
            _resultValue.excludesPattern = excludesPattern;
            _resultValue.externalDependenciesEnabled = externalDependenciesEnabled;
            _resultValue.externalDependenciesPatterns = externalDependenciesPatterns;
            _resultValue.hardFail = hardFail;
            _resultValue.id = id;
            _resultValue.includesPattern = includesPattern;
            _resultValue.key = key;
            _resultValue.listRemoteFolderItems = listRemoteFolderItems;
            _resultValue.localAddress = localAddress;
            _resultValue.metadataRetrievalTimeoutSecs = metadataRetrievalTimeoutSecs;
            _resultValue.mismatchingMimeTypesOverrideList = mismatchingMimeTypesOverrideList;
            _resultValue.missedCachePeriodSeconds = missedCachePeriodSeconds;
            _resultValue.notes = notes;
            _resultValue.offline = offline;
            _resultValue.packageType = packageType;
            _resultValue.password = password;
            _resultValue.priorityResolution = priorityResolution;
            _resultValue.projectEnvironments = projectEnvironments;
            _resultValue.projectId = projectId;
            _resultValue.projectKey = projectKey;
            _resultValue.propertySets = propertySets;
            _resultValue.proxy = proxy;
            _resultValue.queryParams = queryParams;
            _resultValue.remoteRepoLayoutRef = remoteRepoLayoutRef;
            _resultValue.repoLayoutRef = repoLayoutRef;
            _resultValue.retrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
            _resultValue.shareConfiguration = shareConfiguration;
            _resultValue.socketTimeoutMillis = socketTimeoutMillis;
            _resultValue.storeArtifactsLocally = storeArtifactsLocally;
            _resultValue.synchronizeProperties = synchronizeProperties;
            _resultValue.unusedArtifactsCleanupPeriodHours = unusedArtifactsCleanupPeriodHours;
            _resultValue.url = url;
            _resultValue.username = username;
            _resultValue.xrayIndex = xrayIndex;
            return _resultValue;
        }
    }
}

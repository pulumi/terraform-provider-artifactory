// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.artifactory.outputs.GetRemoteGenericRepositoryContentSynchronisation;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRemoteGenericRepositoryResult {
    private @Nullable Boolean allowAnyHostAuth;
    private @Nullable Integer assumedOfflinePeriodSecs;
    private @Nullable Boolean blackedOut;
    private @Nullable Boolean blockMismatchingMimeTypes;
    private @Nullable Boolean bypassHeadRequests;
    private @Nullable Boolean cdnRedirect;
    private String clientTlsCertificate;
    private GetRemoteGenericRepositoryContentSynchronisation contentSynchronisation;
    private @Nullable String description;
    private @Nullable Boolean downloadDirect;
    private @Nullable Boolean enableCookieManagement;
    private @Nullable String excludesPattern;
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
    private @Nullable String projectKey;
    /**
     * @return (Optional, Default: `false`) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     * 
     */
    private @Nullable Boolean propagateQueryParams;
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

    private GetRemoteGenericRepositoryResult() {}
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
    public Optional<Boolean> bypassHeadRequests() {
        return Optional.ofNullable(this.bypassHeadRequests);
    }
    public Optional<Boolean> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }
    public String clientTlsCertificate() {
        return this.clientTlsCertificate;
    }
    public GetRemoteGenericRepositoryContentSynchronisation contentSynchronisation() {
        return this.contentSynchronisation;
    }
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    public Optional<Boolean> downloadDirect() {
        return Optional.ofNullable(this.downloadDirect);
    }
    public Optional<Boolean> enableCookieManagement() {
        return Optional.ofNullable(this.enableCookieManagement);
    }
    public Optional<String> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
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
    public Optional<String> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }
    /**
     * @return (Optional, Default: `false`) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     * 
     */
    public Optional<Boolean> propagateQueryParams() {
        return Optional.ofNullable(this.propagateQueryParams);
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

    public static Builder builder(GetRemoteGenericRepositoryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean allowAnyHostAuth;
        private @Nullable Integer assumedOfflinePeriodSecs;
        private @Nullable Boolean blackedOut;
        private @Nullable Boolean blockMismatchingMimeTypes;
        private @Nullable Boolean bypassHeadRequests;
        private @Nullable Boolean cdnRedirect;
        private String clientTlsCertificate;
        private GetRemoteGenericRepositoryContentSynchronisation contentSynchronisation;
        private @Nullable String description;
        private @Nullable Boolean downloadDirect;
        private @Nullable Boolean enableCookieManagement;
        private @Nullable String excludesPattern;
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
        private @Nullable String projectKey;
        private @Nullable Boolean propagateQueryParams;
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
        public Builder(GetRemoteGenericRepositoryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowAnyHostAuth = defaults.allowAnyHostAuth;
    	      this.assumedOfflinePeriodSecs = defaults.assumedOfflinePeriodSecs;
    	      this.blackedOut = defaults.blackedOut;
    	      this.blockMismatchingMimeTypes = defaults.blockMismatchingMimeTypes;
    	      this.bypassHeadRequests = defaults.bypassHeadRequests;
    	      this.cdnRedirect = defaults.cdnRedirect;
    	      this.clientTlsCertificate = defaults.clientTlsCertificate;
    	      this.contentSynchronisation = defaults.contentSynchronisation;
    	      this.description = defaults.description;
    	      this.downloadDirect = defaults.downloadDirect;
    	      this.enableCookieManagement = defaults.enableCookieManagement;
    	      this.excludesPattern = defaults.excludesPattern;
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
    	      this.projectKey = defaults.projectKey;
    	      this.propagateQueryParams = defaults.propagateQueryParams;
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
            this.clientTlsCertificate = Objects.requireNonNull(clientTlsCertificate);
            return this;
        }
        @CustomType.Setter
        public Builder contentSynchronisation(GetRemoteGenericRepositoryContentSynchronisation contentSynchronisation) {
            this.contentSynchronisation = Objects.requireNonNull(contentSynchronisation);
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
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
        public Builder excludesPattern(@Nullable String excludesPattern) {
            this.excludesPattern = excludesPattern;
            return this;
        }
        @CustomType.Setter
        public Builder hardFail(@Nullable Boolean hardFail) {
            this.hardFail = hardFail;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder includesPattern(@Nullable String includesPattern) {
            this.includesPattern = includesPattern;
            return this;
        }
        @CustomType.Setter
        public Builder key(String key) {
            this.key = Objects.requireNonNull(key);
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
            this.packageType = Objects.requireNonNull(packageType);
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
            this.projectEnvironments = Objects.requireNonNull(projectEnvironments);
            return this;
        }
        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }
        @CustomType.Setter
        public Builder projectKey(@Nullable String projectKey) {
            this.projectKey = projectKey;
            return this;
        }
        @CustomType.Setter
        public Builder propagateQueryParams(@Nullable Boolean propagateQueryParams) {
            this.propagateQueryParams = propagateQueryParams;
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
            this.shareConfiguration = Objects.requireNonNull(shareConfiguration);
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
        public GetRemoteGenericRepositoryResult build() {
            final var o = new GetRemoteGenericRepositoryResult();
            o.allowAnyHostAuth = allowAnyHostAuth;
            o.assumedOfflinePeriodSecs = assumedOfflinePeriodSecs;
            o.blackedOut = blackedOut;
            o.blockMismatchingMimeTypes = blockMismatchingMimeTypes;
            o.bypassHeadRequests = bypassHeadRequests;
            o.cdnRedirect = cdnRedirect;
            o.clientTlsCertificate = clientTlsCertificate;
            o.contentSynchronisation = contentSynchronisation;
            o.description = description;
            o.downloadDirect = downloadDirect;
            o.enableCookieManagement = enableCookieManagement;
            o.excludesPattern = excludesPattern;
            o.hardFail = hardFail;
            o.id = id;
            o.includesPattern = includesPattern;
            o.key = key;
            o.listRemoteFolderItems = listRemoteFolderItems;
            o.localAddress = localAddress;
            o.metadataRetrievalTimeoutSecs = metadataRetrievalTimeoutSecs;
            o.mismatchingMimeTypesOverrideList = mismatchingMimeTypesOverrideList;
            o.missedCachePeriodSeconds = missedCachePeriodSeconds;
            o.notes = notes;
            o.offline = offline;
            o.packageType = packageType;
            o.password = password;
            o.priorityResolution = priorityResolution;
            o.projectEnvironments = projectEnvironments;
            o.projectKey = projectKey;
            o.propagateQueryParams = propagateQueryParams;
            o.propertySets = propertySets;
            o.proxy = proxy;
            o.queryParams = queryParams;
            o.remoteRepoLayoutRef = remoteRepoLayoutRef;
            o.repoLayoutRef = repoLayoutRef;
            o.retrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
            o.shareConfiguration = shareConfiguration;
            o.socketTimeoutMillis = socketTimeoutMillis;
            o.storeArtifactsLocally = storeArtifactsLocally;
            o.synchronizeProperties = synchronizeProperties;
            o.unusedArtifactsCleanupPeriodHours = unusedArtifactsCleanupPeriodHours;
            o.url = url;
            o.username = username;
            o.xrayIndex = xrayIndex;
            return o;
        }
    }
}

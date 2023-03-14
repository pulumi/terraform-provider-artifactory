// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.GetRemoteGenericRepositoryContentSynchronisation;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRemoteGenericRepositoryPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRemoteGenericRepositoryPlainArgs Empty = new GetRemoteGenericRepositoryPlainArgs();

    @Import(name="allowAnyHostAuth")
    private @Nullable Boolean allowAnyHostAuth;

    public Optional<Boolean> allowAnyHostAuth() {
        return Optional.ofNullable(this.allowAnyHostAuth);
    }

    @Import(name="assumedOfflinePeriodSecs")
    private @Nullable Integer assumedOfflinePeriodSecs;

    public Optional<Integer> assumedOfflinePeriodSecs() {
        return Optional.ofNullable(this.assumedOfflinePeriodSecs);
    }

    @Import(name="blackedOut")
    private @Nullable Boolean blackedOut;

    public Optional<Boolean> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }

    @Import(name="blockMismatchingMimeTypes")
    private @Nullable Boolean blockMismatchingMimeTypes;

    public Optional<Boolean> blockMismatchingMimeTypes() {
        return Optional.ofNullable(this.blockMismatchingMimeTypes);
    }

    @Import(name="bypassHeadRequests")
    private @Nullable Boolean bypassHeadRequests;

    public Optional<Boolean> bypassHeadRequests() {
        return Optional.ofNullable(this.bypassHeadRequests);
    }

    @Import(name="cdnRedirect")
    private @Nullable Boolean cdnRedirect;

    public Optional<Boolean> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }

    @Import(name="clientTlsCertificate")
    private @Nullable String clientTlsCertificate;

    public Optional<String> clientTlsCertificate() {
        return Optional.ofNullable(this.clientTlsCertificate);
    }

    @Import(name="contentSynchronisation")
    private @Nullable GetRemoteGenericRepositoryContentSynchronisation contentSynchronisation;

    public Optional<GetRemoteGenericRepositoryContentSynchronisation> contentSynchronisation() {
        return Optional.ofNullable(this.contentSynchronisation);
    }

    @Import(name="description")
    private @Nullable String description;

    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="downloadDirect")
    private @Nullable Boolean downloadDirect;

    public Optional<Boolean> downloadDirect() {
        return Optional.ofNullable(this.downloadDirect);
    }

    @Import(name="enableCookieManagement")
    private @Nullable Boolean enableCookieManagement;

    public Optional<Boolean> enableCookieManagement() {
        return Optional.ofNullable(this.enableCookieManagement);
    }

    @Import(name="excludesPattern")
    private @Nullable String excludesPattern;

    public Optional<String> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    @Import(name="hardFail")
    private @Nullable Boolean hardFail;

    public Optional<Boolean> hardFail() {
        return Optional.ofNullable(this.hardFail);
    }

    @Import(name="includesPattern")
    private @Nullable String includesPattern;

    public Optional<String> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }

    /**
     * the identity key of the repo.
     * 
     */
    @Import(name="key", required=true)
    private String key;

    /**
     * @return the identity key of the repo.
     * 
     */
    public String key() {
        return this.key;
    }

    @Import(name="listRemoteFolderItems")
    private @Nullable Boolean listRemoteFolderItems;

    public Optional<Boolean> listRemoteFolderItems() {
        return Optional.ofNullable(this.listRemoteFolderItems);
    }

    @Import(name="localAddress")
    private @Nullable String localAddress;

    public Optional<String> localAddress() {
        return Optional.ofNullable(this.localAddress);
    }

    @Import(name="metadataRetrievalTimeoutSecs")
    private @Nullable Integer metadataRetrievalTimeoutSecs;

    public Optional<Integer> metadataRetrievalTimeoutSecs() {
        return Optional.ofNullable(this.metadataRetrievalTimeoutSecs);
    }

    @Import(name="mismatchingMimeTypesOverrideList")
    private @Nullable String mismatchingMimeTypesOverrideList;

    public Optional<String> mismatchingMimeTypesOverrideList() {
        return Optional.ofNullable(this.mismatchingMimeTypesOverrideList);
    }

    @Import(name="missedCachePeriodSeconds")
    private @Nullable Integer missedCachePeriodSeconds;

    public Optional<Integer> missedCachePeriodSeconds() {
        return Optional.ofNullable(this.missedCachePeriodSeconds);
    }

    @Import(name="notes")
    private @Nullable String notes;

    public Optional<String> notes() {
        return Optional.ofNullable(this.notes);
    }

    @Import(name="offline")
    private @Nullable Boolean offline;

    public Optional<Boolean> offline() {
        return Optional.ofNullable(this.offline);
    }

    @Import(name="password")
    private @Nullable String password;

    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }

    @Import(name="priorityResolution")
    private @Nullable Boolean priorityResolution;

    public Optional<Boolean> priorityResolution() {
        return Optional.ofNullable(this.priorityResolution);
    }

    @Import(name="projectEnvironments")
    private @Nullable List<String> projectEnvironments;

    public Optional<List<String>> projectEnvironments() {
        return Optional.ofNullable(this.projectEnvironments);
    }

    @Import(name="projectKey")
    private @Nullable String projectKey;

    public Optional<String> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }

    /**
     * (Optional, Default: `false`) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     * 
     */
    @Import(name="propagateQueryParams")
    private @Nullable Boolean propagateQueryParams;

    /**
     * @return (Optional, Default: `false`) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     * 
     */
    public Optional<Boolean> propagateQueryParams() {
        return Optional.ofNullable(this.propagateQueryParams);
    }

    @Import(name="propertySets")
    private @Nullable List<String> propertySets;

    public Optional<List<String>> propertySets() {
        return Optional.ofNullable(this.propertySets);
    }

    @Import(name="proxy")
    private @Nullable String proxy;

    public Optional<String> proxy() {
        return Optional.ofNullable(this.proxy);
    }

    @Import(name="queryParams")
    private @Nullable String queryParams;

    public Optional<String> queryParams() {
        return Optional.ofNullable(this.queryParams);
    }

    @Import(name="remoteRepoLayoutRef")
    private @Nullable String remoteRepoLayoutRef;

    public Optional<String> remoteRepoLayoutRef() {
        return Optional.ofNullable(this.remoteRepoLayoutRef);
    }

    @Import(name="repoLayoutRef")
    private @Nullable String repoLayoutRef;

    public Optional<String> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    @Import(name="retrievalCachePeriodSeconds")
    private @Nullable Integer retrievalCachePeriodSeconds;

    public Optional<Integer> retrievalCachePeriodSeconds() {
        return Optional.ofNullable(this.retrievalCachePeriodSeconds);
    }

    @Import(name="shareConfiguration")
    private @Nullable Boolean shareConfiguration;

    public Optional<Boolean> shareConfiguration() {
        return Optional.ofNullable(this.shareConfiguration);
    }

    @Import(name="socketTimeoutMillis")
    private @Nullable Integer socketTimeoutMillis;

    public Optional<Integer> socketTimeoutMillis() {
        return Optional.ofNullable(this.socketTimeoutMillis);
    }

    @Import(name="storeArtifactsLocally")
    private @Nullable Boolean storeArtifactsLocally;

    public Optional<Boolean> storeArtifactsLocally() {
        return Optional.ofNullable(this.storeArtifactsLocally);
    }

    @Import(name="synchronizeProperties")
    private @Nullable Boolean synchronizeProperties;

    public Optional<Boolean> synchronizeProperties() {
        return Optional.ofNullable(this.synchronizeProperties);
    }

    @Import(name="unusedArtifactsCleanupPeriodHours")
    private @Nullable Integer unusedArtifactsCleanupPeriodHours;

    public Optional<Integer> unusedArtifactsCleanupPeriodHours() {
        return Optional.ofNullable(this.unusedArtifactsCleanupPeriodHours);
    }

    @Import(name="url")
    private @Nullable String url;

    public Optional<String> url() {
        return Optional.ofNullable(this.url);
    }

    @Import(name="username")
    private @Nullable String username;

    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }

    @Import(name="xrayIndex")
    private @Nullable Boolean xrayIndex;

    public Optional<Boolean> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    private GetRemoteGenericRepositoryPlainArgs() {}

    private GetRemoteGenericRepositoryPlainArgs(GetRemoteGenericRepositoryPlainArgs $) {
        this.allowAnyHostAuth = $.allowAnyHostAuth;
        this.assumedOfflinePeriodSecs = $.assumedOfflinePeriodSecs;
        this.blackedOut = $.blackedOut;
        this.blockMismatchingMimeTypes = $.blockMismatchingMimeTypes;
        this.bypassHeadRequests = $.bypassHeadRequests;
        this.cdnRedirect = $.cdnRedirect;
        this.clientTlsCertificate = $.clientTlsCertificate;
        this.contentSynchronisation = $.contentSynchronisation;
        this.description = $.description;
        this.downloadDirect = $.downloadDirect;
        this.enableCookieManagement = $.enableCookieManagement;
        this.excludesPattern = $.excludesPattern;
        this.hardFail = $.hardFail;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.listRemoteFolderItems = $.listRemoteFolderItems;
        this.localAddress = $.localAddress;
        this.metadataRetrievalTimeoutSecs = $.metadataRetrievalTimeoutSecs;
        this.mismatchingMimeTypesOverrideList = $.mismatchingMimeTypesOverrideList;
        this.missedCachePeriodSeconds = $.missedCachePeriodSeconds;
        this.notes = $.notes;
        this.offline = $.offline;
        this.password = $.password;
        this.priorityResolution = $.priorityResolution;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.propagateQueryParams = $.propagateQueryParams;
        this.propertySets = $.propertySets;
        this.proxy = $.proxy;
        this.queryParams = $.queryParams;
        this.remoteRepoLayoutRef = $.remoteRepoLayoutRef;
        this.repoLayoutRef = $.repoLayoutRef;
        this.retrievalCachePeriodSeconds = $.retrievalCachePeriodSeconds;
        this.shareConfiguration = $.shareConfiguration;
        this.socketTimeoutMillis = $.socketTimeoutMillis;
        this.storeArtifactsLocally = $.storeArtifactsLocally;
        this.synchronizeProperties = $.synchronizeProperties;
        this.unusedArtifactsCleanupPeriodHours = $.unusedArtifactsCleanupPeriodHours;
        this.url = $.url;
        this.username = $.username;
        this.xrayIndex = $.xrayIndex;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRemoteGenericRepositoryPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRemoteGenericRepositoryPlainArgs $;

        public Builder() {
            $ = new GetRemoteGenericRepositoryPlainArgs();
        }

        public Builder(GetRemoteGenericRepositoryPlainArgs defaults) {
            $ = new GetRemoteGenericRepositoryPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder allowAnyHostAuth(@Nullable Boolean allowAnyHostAuth) {
            $.allowAnyHostAuth = allowAnyHostAuth;
            return this;
        }

        public Builder assumedOfflinePeriodSecs(@Nullable Integer assumedOfflinePeriodSecs) {
            $.assumedOfflinePeriodSecs = assumedOfflinePeriodSecs;
            return this;
        }

        public Builder blackedOut(@Nullable Boolean blackedOut) {
            $.blackedOut = blackedOut;
            return this;
        }

        public Builder blockMismatchingMimeTypes(@Nullable Boolean blockMismatchingMimeTypes) {
            $.blockMismatchingMimeTypes = blockMismatchingMimeTypes;
            return this;
        }

        public Builder bypassHeadRequests(@Nullable Boolean bypassHeadRequests) {
            $.bypassHeadRequests = bypassHeadRequests;
            return this;
        }

        public Builder cdnRedirect(@Nullable Boolean cdnRedirect) {
            $.cdnRedirect = cdnRedirect;
            return this;
        }

        public Builder clientTlsCertificate(@Nullable String clientTlsCertificate) {
            $.clientTlsCertificate = clientTlsCertificate;
            return this;
        }

        public Builder contentSynchronisation(@Nullable GetRemoteGenericRepositoryContentSynchronisation contentSynchronisation) {
            $.contentSynchronisation = contentSynchronisation;
            return this;
        }

        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        public Builder downloadDirect(@Nullable Boolean downloadDirect) {
            $.downloadDirect = downloadDirect;
            return this;
        }

        public Builder enableCookieManagement(@Nullable Boolean enableCookieManagement) {
            $.enableCookieManagement = enableCookieManagement;
            return this;
        }

        public Builder excludesPattern(@Nullable String excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        public Builder hardFail(@Nullable Boolean hardFail) {
            $.hardFail = hardFail;
            return this;
        }

        public Builder includesPattern(@Nullable String includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        /**
         * @param key the identity key of the repo.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            $.key = key;
            return this;
        }

        public Builder listRemoteFolderItems(@Nullable Boolean listRemoteFolderItems) {
            $.listRemoteFolderItems = listRemoteFolderItems;
            return this;
        }

        public Builder localAddress(@Nullable String localAddress) {
            $.localAddress = localAddress;
            return this;
        }

        public Builder metadataRetrievalTimeoutSecs(@Nullable Integer metadataRetrievalTimeoutSecs) {
            $.metadataRetrievalTimeoutSecs = metadataRetrievalTimeoutSecs;
            return this;
        }

        public Builder mismatchingMimeTypesOverrideList(@Nullable String mismatchingMimeTypesOverrideList) {
            $.mismatchingMimeTypesOverrideList = mismatchingMimeTypesOverrideList;
            return this;
        }

        public Builder missedCachePeriodSeconds(@Nullable Integer missedCachePeriodSeconds) {
            $.missedCachePeriodSeconds = missedCachePeriodSeconds;
            return this;
        }

        public Builder notes(@Nullable String notes) {
            $.notes = notes;
            return this;
        }

        public Builder offline(@Nullable Boolean offline) {
            $.offline = offline;
            return this;
        }

        public Builder password(@Nullable String password) {
            $.password = password;
            return this;
        }

        public Builder priorityResolution(@Nullable Boolean priorityResolution) {
            $.priorityResolution = priorityResolution;
            return this;
        }

        public Builder projectEnvironments(@Nullable List<String> projectEnvironments) {
            $.projectEnvironments = projectEnvironments;
            return this;
        }

        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }

        public Builder projectKey(@Nullable String projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        /**
         * @param propagateQueryParams (Optional, Default: `false`) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
         * 
         * @return builder
         * 
         */
        public Builder propagateQueryParams(@Nullable Boolean propagateQueryParams) {
            $.propagateQueryParams = propagateQueryParams;
            return this;
        }

        public Builder propertySets(@Nullable List<String> propertySets) {
            $.propertySets = propertySets;
            return this;
        }

        public Builder propertySets(String... propertySets) {
            return propertySets(List.of(propertySets));
        }

        public Builder proxy(@Nullable String proxy) {
            $.proxy = proxy;
            return this;
        }

        public Builder queryParams(@Nullable String queryParams) {
            $.queryParams = queryParams;
            return this;
        }

        public Builder remoteRepoLayoutRef(@Nullable String remoteRepoLayoutRef) {
            $.remoteRepoLayoutRef = remoteRepoLayoutRef;
            return this;
        }

        public Builder repoLayoutRef(@Nullable String repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        public Builder retrievalCachePeriodSeconds(@Nullable Integer retrievalCachePeriodSeconds) {
            $.retrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
            return this;
        }

        public Builder shareConfiguration(@Nullable Boolean shareConfiguration) {
            $.shareConfiguration = shareConfiguration;
            return this;
        }

        public Builder socketTimeoutMillis(@Nullable Integer socketTimeoutMillis) {
            $.socketTimeoutMillis = socketTimeoutMillis;
            return this;
        }

        public Builder storeArtifactsLocally(@Nullable Boolean storeArtifactsLocally) {
            $.storeArtifactsLocally = storeArtifactsLocally;
            return this;
        }

        public Builder synchronizeProperties(@Nullable Boolean synchronizeProperties) {
            $.synchronizeProperties = synchronizeProperties;
            return this;
        }

        public Builder unusedArtifactsCleanupPeriodHours(@Nullable Integer unusedArtifactsCleanupPeriodHours) {
            $.unusedArtifactsCleanupPeriodHours = unusedArtifactsCleanupPeriodHours;
            return this;
        }

        public Builder url(@Nullable String url) {
            $.url = url;
            return this;
        }

        public Builder username(@Nullable String username) {
            $.username = username;
            return this;
        }

        public Builder xrayIndex(@Nullable Boolean xrayIndex) {
            $.xrayIndex = xrayIndex;
            return this;
        }

        public GetRemoteGenericRepositoryPlainArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}

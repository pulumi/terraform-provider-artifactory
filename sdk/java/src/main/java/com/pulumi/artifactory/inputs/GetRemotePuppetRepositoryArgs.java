// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.GetRemotePuppetRepositoryContentSynchronisationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRemotePuppetRepositoryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRemotePuppetRepositoryArgs Empty = new GetRemotePuppetRepositoryArgs();

    @Import(name="allowAnyHostAuth")
    private @Nullable Output<Boolean> allowAnyHostAuth;

    public Optional<Output<Boolean>> allowAnyHostAuth() {
        return Optional.ofNullable(this.allowAnyHostAuth);
    }

    @Import(name="assumedOfflinePeriodSecs")
    private @Nullable Output<Integer> assumedOfflinePeriodSecs;

    public Optional<Output<Integer>> assumedOfflinePeriodSecs() {
        return Optional.ofNullable(this.assumedOfflinePeriodSecs);
    }

    @Import(name="blackedOut")
    private @Nullable Output<Boolean> blackedOut;

    public Optional<Output<Boolean>> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }

    @Import(name="blockMismatchingMimeTypes")
    private @Nullable Output<Boolean> blockMismatchingMimeTypes;

    public Optional<Output<Boolean>> blockMismatchingMimeTypes() {
        return Optional.ofNullable(this.blockMismatchingMimeTypes);
    }

    @Import(name="bypassHeadRequests")
    private @Nullable Output<Boolean> bypassHeadRequests;

    public Optional<Output<Boolean>> bypassHeadRequests() {
        return Optional.ofNullable(this.bypassHeadRequests);
    }

    @Import(name="cdnRedirect")
    private @Nullable Output<Boolean> cdnRedirect;

    public Optional<Output<Boolean>> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }

    @Import(name="clientTlsCertificate")
    private @Nullable Output<String> clientTlsCertificate;

    public Optional<Output<String>> clientTlsCertificate() {
        return Optional.ofNullable(this.clientTlsCertificate);
    }

    @Import(name="contentSynchronisation")
    private @Nullable Output<GetRemotePuppetRepositoryContentSynchronisationArgs> contentSynchronisation;

    public Optional<Output<GetRemotePuppetRepositoryContentSynchronisationArgs>> contentSynchronisation() {
        return Optional.ofNullable(this.contentSynchronisation);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="downloadDirect")
    private @Nullable Output<Boolean> downloadDirect;

    public Optional<Output<Boolean>> downloadDirect() {
        return Optional.ofNullable(this.downloadDirect);
    }

    @Import(name="enableCookieManagement")
    private @Nullable Output<Boolean> enableCookieManagement;

    public Optional<Output<Boolean>> enableCookieManagement() {
        return Optional.ofNullable(this.enableCookieManagement);
    }

    @Import(name="excludesPattern")
    private @Nullable Output<String> excludesPattern;

    public Optional<Output<String>> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    @Import(name="hardFail")
    private @Nullable Output<Boolean> hardFail;

    public Optional<Output<Boolean>> hardFail() {
        return Optional.ofNullable(this.hardFail);
    }

    @Import(name="includesPattern")
    private @Nullable Output<String> includesPattern;

    public Optional<Output<String>> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }

    /**
     * the identity key of the repo.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return the identity key of the repo.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    @Import(name="listRemoteFolderItems")
    private @Nullable Output<Boolean> listRemoteFolderItems;

    public Optional<Output<Boolean>> listRemoteFolderItems() {
        return Optional.ofNullable(this.listRemoteFolderItems);
    }

    @Import(name="localAddress")
    private @Nullable Output<String> localAddress;

    public Optional<Output<String>> localAddress() {
        return Optional.ofNullable(this.localAddress);
    }

    @Import(name="metadataRetrievalTimeoutSecs")
    private @Nullable Output<Integer> metadataRetrievalTimeoutSecs;

    public Optional<Output<Integer>> metadataRetrievalTimeoutSecs() {
        return Optional.ofNullable(this.metadataRetrievalTimeoutSecs);
    }

    @Import(name="mismatchingMimeTypesOverrideList")
    private @Nullable Output<String> mismatchingMimeTypesOverrideList;

    public Optional<Output<String>> mismatchingMimeTypesOverrideList() {
        return Optional.ofNullable(this.mismatchingMimeTypesOverrideList);
    }

    @Import(name="missedCachePeriodSeconds")
    private @Nullable Output<Integer> missedCachePeriodSeconds;

    public Optional<Output<Integer>> missedCachePeriodSeconds() {
        return Optional.ofNullable(this.missedCachePeriodSeconds);
    }

    @Import(name="notes")
    private @Nullable Output<String> notes;

    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
    }

    @Import(name="offline")
    private @Nullable Output<Boolean> offline;

    public Optional<Output<Boolean>> offline() {
        return Optional.ofNullable(this.offline);
    }

    @Import(name="password")
    private @Nullable Output<String> password;

    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    @Import(name="priorityResolution")
    private @Nullable Output<Boolean> priorityResolution;

    public Optional<Output<Boolean>> priorityResolution() {
        return Optional.ofNullable(this.priorityResolution);
    }

    @Import(name="projectEnvironments")
    private @Nullable Output<List<String>> projectEnvironments;

    public Optional<Output<List<String>>> projectEnvironments() {
        return Optional.ofNullable(this.projectEnvironments);
    }

    @Import(name="projectKey")
    private @Nullable Output<String> projectKey;

    public Optional<Output<String>> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }

    @Import(name="propertySets")
    private @Nullable Output<List<String>> propertySets;

    public Optional<Output<List<String>>> propertySets() {
        return Optional.ofNullable(this.propertySets);
    }

    @Import(name="proxy")
    private @Nullable Output<String> proxy;

    public Optional<Output<String>> proxy() {
        return Optional.ofNullable(this.proxy);
    }

    @Import(name="queryParams")
    private @Nullable Output<String> queryParams;

    public Optional<Output<String>> queryParams() {
        return Optional.ofNullable(this.queryParams);
    }

    @Import(name="remoteRepoLayoutRef")
    private @Nullable Output<String> remoteRepoLayoutRef;

    public Optional<Output<String>> remoteRepoLayoutRef() {
        return Optional.ofNullable(this.remoteRepoLayoutRef);
    }

    @Import(name="repoLayoutRef")
    private @Nullable Output<String> repoLayoutRef;

    public Optional<Output<String>> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    @Import(name="retrievalCachePeriodSeconds")
    private @Nullable Output<Integer> retrievalCachePeriodSeconds;

    public Optional<Output<Integer>> retrievalCachePeriodSeconds() {
        return Optional.ofNullable(this.retrievalCachePeriodSeconds);
    }

    @Import(name="shareConfiguration")
    private @Nullable Output<Boolean> shareConfiguration;

    public Optional<Output<Boolean>> shareConfiguration() {
        return Optional.ofNullable(this.shareConfiguration);
    }

    @Import(name="socketTimeoutMillis")
    private @Nullable Output<Integer> socketTimeoutMillis;

    public Optional<Output<Integer>> socketTimeoutMillis() {
        return Optional.ofNullable(this.socketTimeoutMillis);
    }

    @Import(name="storeArtifactsLocally")
    private @Nullable Output<Boolean> storeArtifactsLocally;

    public Optional<Output<Boolean>> storeArtifactsLocally() {
        return Optional.ofNullable(this.storeArtifactsLocally);
    }

    @Import(name="synchronizeProperties")
    private @Nullable Output<Boolean> synchronizeProperties;

    public Optional<Output<Boolean>> synchronizeProperties() {
        return Optional.ofNullable(this.synchronizeProperties);
    }

    @Import(name="unusedArtifactsCleanupPeriodHours")
    private @Nullable Output<Integer> unusedArtifactsCleanupPeriodHours;

    public Optional<Output<Integer>> unusedArtifactsCleanupPeriodHours() {
        return Optional.ofNullable(this.unusedArtifactsCleanupPeriodHours);
    }

    @Import(name="url")
    private @Nullable Output<String> url;

    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    @Import(name="username")
    private @Nullable Output<String> username;

    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    @Import(name="xrayIndex")
    private @Nullable Output<Boolean> xrayIndex;

    public Optional<Output<Boolean>> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    private GetRemotePuppetRepositoryArgs() {}

    private GetRemotePuppetRepositoryArgs(GetRemotePuppetRepositoryArgs $) {
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
    public static Builder builder(GetRemotePuppetRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRemotePuppetRepositoryArgs $;

        public Builder() {
            $ = new GetRemotePuppetRepositoryArgs();
        }

        public Builder(GetRemotePuppetRepositoryArgs defaults) {
            $ = new GetRemotePuppetRepositoryArgs(Objects.requireNonNull(defaults));
        }

        public Builder allowAnyHostAuth(@Nullable Output<Boolean> allowAnyHostAuth) {
            $.allowAnyHostAuth = allowAnyHostAuth;
            return this;
        }

        public Builder allowAnyHostAuth(Boolean allowAnyHostAuth) {
            return allowAnyHostAuth(Output.of(allowAnyHostAuth));
        }

        public Builder assumedOfflinePeriodSecs(@Nullable Output<Integer> assumedOfflinePeriodSecs) {
            $.assumedOfflinePeriodSecs = assumedOfflinePeriodSecs;
            return this;
        }

        public Builder assumedOfflinePeriodSecs(Integer assumedOfflinePeriodSecs) {
            return assumedOfflinePeriodSecs(Output.of(assumedOfflinePeriodSecs));
        }

        public Builder blackedOut(@Nullable Output<Boolean> blackedOut) {
            $.blackedOut = blackedOut;
            return this;
        }

        public Builder blackedOut(Boolean blackedOut) {
            return blackedOut(Output.of(blackedOut));
        }

        public Builder blockMismatchingMimeTypes(@Nullable Output<Boolean> blockMismatchingMimeTypes) {
            $.blockMismatchingMimeTypes = blockMismatchingMimeTypes;
            return this;
        }

        public Builder blockMismatchingMimeTypes(Boolean blockMismatchingMimeTypes) {
            return blockMismatchingMimeTypes(Output.of(blockMismatchingMimeTypes));
        }

        public Builder bypassHeadRequests(@Nullable Output<Boolean> bypassHeadRequests) {
            $.bypassHeadRequests = bypassHeadRequests;
            return this;
        }

        public Builder bypassHeadRequests(Boolean bypassHeadRequests) {
            return bypassHeadRequests(Output.of(bypassHeadRequests));
        }

        public Builder cdnRedirect(@Nullable Output<Boolean> cdnRedirect) {
            $.cdnRedirect = cdnRedirect;
            return this;
        }

        public Builder cdnRedirect(Boolean cdnRedirect) {
            return cdnRedirect(Output.of(cdnRedirect));
        }

        public Builder clientTlsCertificate(@Nullable Output<String> clientTlsCertificate) {
            $.clientTlsCertificate = clientTlsCertificate;
            return this;
        }

        public Builder clientTlsCertificate(String clientTlsCertificate) {
            return clientTlsCertificate(Output.of(clientTlsCertificate));
        }

        public Builder contentSynchronisation(@Nullable Output<GetRemotePuppetRepositoryContentSynchronisationArgs> contentSynchronisation) {
            $.contentSynchronisation = contentSynchronisation;
            return this;
        }

        public Builder contentSynchronisation(GetRemotePuppetRepositoryContentSynchronisationArgs contentSynchronisation) {
            return contentSynchronisation(Output.of(contentSynchronisation));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder downloadDirect(@Nullable Output<Boolean> downloadDirect) {
            $.downloadDirect = downloadDirect;
            return this;
        }

        public Builder downloadDirect(Boolean downloadDirect) {
            return downloadDirect(Output.of(downloadDirect));
        }

        public Builder enableCookieManagement(@Nullable Output<Boolean> enableCookieManagement) {
            $.enableCookieManagement = enableCookieManagement;
            return this;
        }

        public Builder enableCookieManagement(Boolean enableCookieManagement) {
            return enableCookieManagement(Output.of(enableCookieManagement));
        }

        public Builder excludesPattern(@Nullable Output<String> excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        public Builder excludesPattern(String excludesPattern) {
            return excludesPattern(Output.of(excludesPattern));
        }

        public Builder hardFail(@Nullable Output<Boolean> hardFail) {
            $.hardFail = hardFail;
            return this;
        }

        public Builder hardFail(Boolean hardFail) {
            return hardFail(Output.of(hardFail));
        }

        public Builder includesPattern(@Nullable Output<String> includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        public Builder includesPattern(String includesPattern) {
            return includesPattern(Output.of(includesPattern));
        }

        /**
         * @param key the identity key of the repo.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key the identity key of the repo.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        public Builder listRemoteFolderItems(@Nullable Output<Boolean> listRemoteFolderItems) {
            $.listRemoteFolderItems = listRemoteFolderItems;
            return this;
        }

        public Builder listRemoteFolderItems(Boolean listRemoteFolderItems) {
            return listRemoteFolderItems(Output.of(listRemoteFolderItems));
        }

        public Builder localAddress(@Nullable Output<String> localAddress) {
            $.localAddress = localAddress;
            return this;
        }

        public Builder localAddress(String localAddress) {
            return localAddress(Output.of(localAddress));
        }

        public Builder metadataRetrievalTimeoutSecs(@Nullable Output<Integer> metadataRetrievalTimeoutSecs) {
            $.metadataRetrievalTimeoutSecs = metadataRetrievalTimeoutSecs;
            return this;
        }

        public Builder metadataRetrievalTimeoutSecs(Integer metadataRetrievalTimeoutSecs) {
            return metadataRetrievalTimeoutSecs(Output.of(metadataRetrievalTimeoutSecs));
        }

        public Builder mismatchingMimeTypesOverrideList(@Nullable Output<String> mismatchingMimeTypesOverrideList) {
            $.mismatchingMimeTypesOverrideList = mismatchingMimeTypesOverrideList;
            return this;
        }

        public Builder mismatchingMimeTypesOverrideList(String mismatchingMimeTypesOverrideList) {
            return mismatchingMimeTypesOverrideList(Output.of(mismatchingMimeTypesOverrideList));
        }

        public Builder missedCachePeriodSeconds(@Nullable Output<Integer> missedCachePeriodSeconds) {
            $.missedCachePeriodSeconds = missedCachePeriodSeconds;
            return this;
        }

        public Builder missedCachePeriodSeconds(Integer missedCachePeriodSeconds) {
            return missedCachePeriodSeconds(Output.of(missedCachePeriodSeconds));
        }

        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        public Builder notes(String notes) {
            return notes(Output.of(notes));
        }

        public Builder offline(@Nullable Output<Boolean> offline) {
            $.offline = offline;
            return this;
        }

        public Builder offline(Boolean offline) {
            return offline(Output.of(offline));
        }

        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        public Builder password(String password) {
            return password(Output.of(password));
        }

        public Builder priorityResolution(@Nullable Output<Boolean> priorityResolution) {
            $.priorityResolution = priorityResolution;
            return this;
        }

        public Builder priorityResolution(Boolean priorityResolution) {
            return priorityResolution(Output.of(priorityResolution));
        }

        public Builder projectEnvironments(@Nullable Output<List<String>> projectEnvironments) {
            $.projectEnvironments = projectEnvironments;
            return this;
        }

        public Builder projectEnvironments(List<String> projectEnvironments) {
            return projectEnvironments(Output.of(projectEnvironments));
        }

        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }

        public Builder projectKey(@Nullable Output<String> projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        public Builder projectKey(String projectKey) {
            return projectKey(Output.of(projectKey));
        }

        public Builder propertySets(@Nullable Output<List<String>> propertySets) {
            $.propertySets = propertySets;
            return this;
        }

        public Builder propertySets(List<String> propertySets) {
            return propertySets(Output.of(propertySets));
        }

        public Builder propertySets(String... propertySets) {
            return propertySets(List.of(propertySets));
        }

        public Builder proxy(@Nullable Output<String> proxy) {
            $.proxy = proxy;
            return this;
        }

        public Builder proxy(String proxy) {
            return proxy(Output.of(proxy));
        }

        public Builder queryParams(@Nullable Output<String> queryParams) {
            $.queryParams = queryParams;
            return this;
        }

        public Builder queryParams(String queryParams) {
            return queryParams(Output.of(queryParams));
        }

        public Builder remoteRepoLayoutRef(@Nullable Output<String> remoteRepoLayoutRef) {
            $.remoteRepoLayoutRef = remoteRepoLayoutRef;
            return this;
        }

        public Builder remoteRepoLayoutRef(String remoteRepoLayoutRef) {
            return remoteRepoLayoutRef(Output.of(remoteRepoLayoutRef));
        }

        public Builder repoLayoutRef(@Nullable Output<String> repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        public Builder repoLayoutRef(String repoLayoutRef) {
            return repoLayoutRef(Output.of(repoLayoutRef));
        }

        public Builder retrievalCachePeriodSeconds(@Nullable Output<Integer> retrievalCachePeriodSeconds) {
            $.retrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
            return this;
        }

        public Builder retrievalCachePeriodSeconds(Integer retrievalCachePeriodSeconds) {
            return retrievalCachePeriodSeconds(Output.of(retrievalCachePeriodSeconds));
        }

        public Builder shareConfiguration(@Nullable Output<Boolean> shareConfiguration) {
            $.shareConfiguration = shareConfiguration;
            return this;
        }

        public Builder shareConfiguration(Boolean shareConfiguration) {
            return shareConfiguration(Output.of(shareConfiguration));
        }

        public Builder socketTimeoutMillis(@Nullable Output<Integer> socketTimeoutMillis) {
            $.socketTimeoutMillis = socketTimeoutMillis;
            return this;
        }

        public Builder socketTimeoutMillis(Integer socketTimeoutMillis) {
            return socketTimeoutMillis(Output.of(socketTimeoutMillis));
        }

        public Builder storeArtifactsLocally(@Nullable Output<Boolean> storeArtifactsLocally) {
            $.storeArtifactsLocally = storeArtifactsLocally;
            return this;
        }

        public Builder storeArtifactsLocally(Boolean storeArtifactsLocally) {
            return storeArtifactsLocally(Output.of(storeArtifactsLocally));
        }

        public Builder synchronizeProperties(@Nullable Output<Boolean> synchronizeProperties) {
            $.synchronizeProperties = synchronizeProperties;
            return this;
        }

        public Builder synchronizeProperties(Boolean synchronizeProperties) {
            return synchronizeProperties(Output.of(synchronizeProperties));
        }

        public Builder unusedArtifactsCleanupPeriodHours(@Nullable Output<Integer> unusedArtifactsCleanupPeriodHours) {
            $.unusedArtifactsCleanupPeriodHours = unusedArtifactsCleanupPeriodHours;
            return this;
        }

        public Builder unusedArtifactsCleanupPeriodHours(Integer unusedArtifactsCleanupPeriodHours) {
            return unusedArtifactsCleanupPeriodHours(Output.of(unusedArtifactsCleanupPeriodHours));
        }

        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        public Builder url(String url) {
            return url(Output.of(url));
        }

        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        public Builder username(String username) {
            return username(Output.of(username));
        }

        public Builder xrayIndex(@Nullable Output<Boolean> xrayIndex) {
            $.xrayIndex = xrayIndex;
            return this;
        }

        public Builder xrayIndex(Boolean xrayIndex) {
            return xrayIndex(Output.of(xrayIndex));
        }

        public GetRemotePuppetRepositoryArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}

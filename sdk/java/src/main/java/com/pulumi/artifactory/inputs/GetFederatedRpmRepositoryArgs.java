// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.GetFederatedRpmRepositoryMemberArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFederatedRpmRepositoryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFederatedRpmRepositoryArgs Empty = new GetFederatedRpmRepositoryArgs();

    @Import(name="archiveBrowsingEnabled")
    private @Nullable Output<Boolean> archiveBrowsingEnabled;

    public Optional<Output<Boolean>> archiveBrowsingEnabled() {
        return Optional.ofNullable(this.archiveBrowsingEnabled);
    }

    @Import(name="blackedOut")
    private @Nullable Output<Boolean> blackedOut;

    public Optional<Output<Boolean>> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }

    @Import(name="calculateYumMetadata")
    private @Nullable Output<Boolean> calculateYumMetadata;

    public Optional<Output<Boolean>> calculateYumMetadata() {
        return Optional.ofNullable(this.calculateYumMetadata);
    }

    @Import(name="cdnRedirect")
    private @Nullable Output<Boolean> cdnRedirect;

    public Optional<Output<Boolean>> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }

    @Import(name="cleanupOnDelete")
    private @Nullable Output<Boolean> cleanupOnDelete;

    public Optional<Output<Boolean>> cleanupOnDelete() {
        return Optional.ofNullable(this.cleanupOnDelete);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
     * 
     */
    @Import(name="disableProxy")
    private @Nullable Output<Boolean> disableProxy;

    /**
     * @return When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
     * 
     */
    public Optional<Output<Boolean>> disableProxy() {
        return Optional.ofNullable(this.disableProxy);
    }

    @Import(name="downloadDirect")
    private @Nullable Output<Boolean> downloadDirect;

    public Optional<Output<Boolean>> downloadDirect() {
        return Optional.ofNullable(this.downloadDirect);
    }

    @Import(name="enableFileListsIndexing")
    private @Nullable Output<Boolean> enableFileListsIndexing;

    public Optional<Output<Boolean>> enableFileListsIndexing() {
        return Optional.ofNullable(this.enableFileListsIndexing);
    }

    @Import(name="excludesPattern")
    private @Nullable Output<String> excludesPattern;

    public Optional<Output<String>> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
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

    /**
     * The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     * 
     */
    @Import(name="members")
    private @Nullable Output<List<GetFederatedRpmRepositoryMemberArgs>> members;

    /**
     * @return The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     * 
     */
    public Optional<Output<List<GetFederatedRpmRepositoryMemberArgs>>> members() {
        return Optional.ofNullable(this.members);
    }

    @Import(name="notes")
    private @Nullable Output<String> notes;

    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
    }

    @Import(name="primaryKeypairRef")
    private @Nullable Output<String> primaryKeypairRef;

    public Optional<Output<String>> primaryKeypairRef() {
        return Optional.ofNullable(this.primaryKeypairRef);
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

    /**
     * Proxy key from Artifactory Proxies settings.
     * 
     */
    @Import(name="proxy")
    private @Nullable Output<String> proxy;

    /**
     * @return Proxy key from Artifactory Proxies settings.
     * 
     */
    public Optional<Output<String>> proxy() {
        return Optional.ofNullable(this.proxy);
    }

    @Import(name="repoLayoutRef")
    private @Nullable Output<String> repoLayoutRef;

    public Optional<Output<String>> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    @Import(name="secondaryKeypairRef")
    private @Nullable Output<String> secondaryKeypairRef;

    public Optional<Output<String>> secondaryKeypairRef() {
        return Optional.ofNullable(this.secondaryKeypairRef);
    }

    @Import(name="xrayIndex")
    private @Nullable Output<Boolean> xrayIndex;

    public Optional<Output<Boolean>> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    @Import(name="yumGroupFileNames")
    private @Nullable Output<String> yumGroupFileNames;

    public Optional<Output<String>> yumGroupFileNames() {
        return Optional.ofNullable(this.yumGroupFileNames);
    }

    @Import(name="yumRootDepth")
    private @Nullable Output<Integer> yumRootDepth;

    public Optional<Output<Integer>> yumRootDepth() {
        return Optional.ofNullable(this.yumRootDepth);
    }

    private GetFederatedRpmRepositoryArgs() {}

    private GetFederatedRpmRepositoryArgs(GetFederatedRpmRepositoryArgs $) {
        this.archiveBrowsingEnabled = $.archiveBrowsingEnabled;
        this.blackedOut = $.blackedOut;
        this.calculateYumMetadata = $.calculateYumMetadata;
        this.cdnRedirect = $.cdnRedirect;
        this.cleanupOnDelete = $.cleanupOnDelete;
        this.description = $.description;
        this.disableProxy = $.disableProxy;
        this.downloadDirect = $.downloadDirect;
        this.enableFileListsIndexing = $.enableFileListsIndexing;
        this.excludesPattern = $.excludesPattern;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.members = $.members;
        this.notes = $.notes;
        this.primaryKeypairRef = $.primaryKeypairRef;
        this.priorityResolution = $.priorityResolution;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.propertySets = $.propertySets;
        this.proxy = $.proxy;
        this.repoLayoutRef = $.repoLayoutRef;
        this.secondaryKeypairRef = $.secondaryKeypairRef;
        this.xrayIndex = $.xrayIndex;
        this.yumGroupFileNames = $.yumGroupFileNames;
        this.yumRootDepth = $.yumRootDepth;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFederatedRpmRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFederatedRpmRepositoryArgs $;

        public Builder() {
            $ = new GetFederatedRpmRepositoryArgs();
        }

        public Builder(GetFederatedRpmRepositoryArgs defaults) {
            $ = new GetFederatedRpmRepositoryArgs(Objects.requireNonNull(defaults));
        }

        public Builder archiveBrowsingEnabled(@Nullable Output<Boolean> archiveBrowsingEnabled) {
            $.archiveBrowsingEnabled = archiveBrowsingEnabled;
            return this;
        }

        public Builder archiveBrowsingEnabled(Boolean archiveBrowsingEnabled) {
            return archiveBrowsingEnabled(Output.of(archiveBrowsingEnabled));
        }

        public Builder blackedOut(@Nullable Output<Boolean> blackedOut) {
            $.blackedOut = blackedOut;
            return this;
        }

        public Builder blackedOut(Boolean blackedOut) {
            return blackedOut(Output.of(blackedOut));
        }

        public Builder calculateYumMetadata(@Nullable Output<Boolean> calculateYumMetadata) {
            $.calculateYumMetadata = calculateYumMetadata;
            return this;
        }

        public Builder calculateYumMetadata(Boolean calculateYumMetadata) {
            return calculateYumMetadata(Output.of(calculateYumMetadata));
        }

        public Builder cdnRedirect(@Nullable Output<Boolean> cdnRedirect) {
            $.cdnRedirect = cdnRedirect;
            return this;
        }

        public Builder cdnRedirect(Boolean cdnRedirect) {
            return cdnRedirect(Output.of(cdnRedirect));
        }

        public Builder cleanupOnDelete(@Nullable Output<Boolean> cleanupOnDelete) {
            $.cleanupOnDelete = cleanupOnDelete;
            return this;
        }

        public Builder cleanupOnDelete(Boolean cleanupOnDelete) {
            return cleanupOnDelete(Output.of(cleanupOnDelete));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param disableProxy When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
         * 
         * @return builder
         * 
         */
        public Builder disableProxy(@Nullable Output<Boolean> disableProxy) {
            $.disableProxy = disableProxy;
            return this;
        }

        /**
         * @param disableProxy When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
         * 
         * @return builder
         * 
         */
        public Builder disableProxy(Boolean disableProxy) {
            return disableProxy(Output.of(disableProxy));
        }

        public Builder downloadDirect(@Nullable Output<Boolean> downloadDirect) {
            $.downloadDirect = downloadDirect;
            return this;
        }

        public Builder downloadDirect(Boolean downloadDirect) {
            return downloadDirect(Output.of(downloadDirect));
        }

        public Builder enableFileListsIndexing(@Nullable Output<Boolean> enableFileListsIndexing) {
            $.enableFileListsIndexing = enableFileListsIndexing;
            return this;
        }

        public Builder enableFileListsIndexing(Boolean enableFileListsIndexing) {
            return enableFileListsIndexing(Output.of(enableFileListsIndexing));
        }

        public Builder excludesPattern(@Nullable Output<String> excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        public Builder excludesPattern(String excludesPattern) {
            return excludesPattern(Output.of(excludesPattern));
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

        /**
         * @param members The list of Federated members and must contain this repository URL (configured base URL
         * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
         * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
         * to set up Federated repositories correctly.
         * 
         * @return builder
         * 
         */
        public Builder members(@Nullable Output<List<GetFederatedRpmRepositoryMemberArgs>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members The list of Federated members and must contain this repository URL (configured base URL
         * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
         * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
         * to set up Federated repositories correctly.
         * 
         * @return builder
         * 
         */
        public Builder members(List<GetFederatedRpmRepositoryMemberArgs> members) {
            return members(Output.of(members));
        }

        /**
         * @param members The list of Federated members and must contain this repository URL (configured base URL
         * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
         * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
         * to set up Federated repositories correctly.
         * 
         * @return builder
         * 
         */
        public Builder members(GetFederatedRpmRepositoryMemberArgs... members) {
            return members(List.of(members));
        }

        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        public Builder notes(String notes) {
            return notes(Output.of(notes));
        }

        public Builder primaryKeypairRef(@Nullable Output<String> primaryKeypairRef) {
            $.primaryKeypairRef = primaryKeypairRef;
            return this;
        }

        public Builder primaryKeypairRef(String primaryKeypairRef) {
            return primaryKeypairRef(Output.of(primaryKeypairRef));
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

        /**
         * @param proxy Proxy key from Artifactory Proxies settings.
         * 
         * @return builder
         * 
         */
        public Builder proxy(@Nullable Output<String> proxy) {
            $.proxy = proxy;
            return this;
        }

        /**
         * @param proxy Proxy key from Artifactory Proxies settings.
         * 
         * @return builder
         * 
         */
        public Builder proxy(String proxy) {
            return proxy(Output.of(proxy));
        }

        public Builder repoLayoutRef(@Nullable Output<String> repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        public Builder repoLayoutRef(String repoLayoutRef) {
            return repoLayoutRef(Output.of(repoLayoutRef));
        }

        public Builder secondaryKeypairRef(@Nullable Output<String> secondaryKeypairRef) {
            $.secondaryKeypairRef = secondaryKeypairRef;
            return this;
        }

        public Builder secondaryKeypairRef(String secondaryKeypairRef) {
            return secondaryKeypairRef(Output.of(secondaryKeypairRef));
        }

        public Builder xrayIndex(@Nullable Output<Boolean> xrayIndex) {
            $.xrayIndex = xrayIndex;
            return this;
        }

        public Builder xrayIndex(Boolean xrayIndex) {
            return xrayIndex(Output.of(xrayIndex));
        }

        public Builder yumGroupFileNames(@Nullable Output<String> yumGroupFileNames) {
            $.yumGroupFileNames = yumGroupFileNames;
            return this;
        }

        public Builder yumGroupFileNames(String yumGroupFileNames) {
            return yumGroupFileNames(Output.of(yumGroupFileNames));
        }

        public Builder yumRootDepth(@Nullable Output<Integer> yumRootDepth) {
            $.yumRootDepth = yumRootDepth;
            return this;
        }

        public Builder yumRootDepth(Integer yumRootDepth) {
            return yumRootDepth(Output.of(yumRootDepth));
        }

        public GetFederatedRpmRepositoryArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("GetFederatedRpmRepositoryArgs", "key");
            }
            return $;
        }
    }

}

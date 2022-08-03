// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LocalRpmRepositoryState extends com.pulumi.resources.ResourceArgs {

    public static final LocalRpmRepositoryState Empty = new LocalRpmRepositoryState();

    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    @Import(name="archiveBrowsingEnabled")
    private @Nullable Output<Boolean> archiveBrowsingEnabled;

    /**
     * @return When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    public Optional<Output<Boolean>> archiveBrowsingEnabled() {
        return Optional.ofNullable(this.archiveBrowsingEnabled);
    }

    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     * 
     */
    @Import(name="blackedOut")
    private @Nullable Output<Boolean> blackedOut;

    /**
     * @return When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     * 
     */
    public Optional<Output<Boolean>> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }

    /**
     * Default: false.
     * 
     */
    @Import(name="calculateYumMetadata")
    private @Nullable Output<Boolean> calculateYumMetadata;

    /**
     * @return Default: false.
     * 
     */
    public Optional<Output<Boolean>> calculateYumMetadata() {
        return Optional.ofNullable(this.calculateYumMetadata);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     * 
     */
    @Import(name="downloadDirect")
    private @Nullable Output<Boolean> downloadDirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     * 
     */
    public Optional<Output<Boolean>> downloadDirect() {
        return Optional.ofNullable(this.downloadDirect);
    }

    /**
     * Default: false.
     * 
     */
    @Import(name="enableFileListsIndexing")
    private @Nullable Output<Boolean> enableFileListsIndexing;

    /**
     * @return Default: false.
     * 
     */
    public Optional<Output<Boolean>> enableFileListsIndexing() {
        return Optional.ofNullable(this.enableFileListsIndexing);
    }

    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
     * artifacts are excluded.
     * 
     */
    @Import(name="excludesPattern")
    private @Nullable Output<String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
     * artifacts are excluded.
     * 
     */
    public Optional<Output<String>> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Import(name="includesPattern")
    private @Nullable Output<String> includesPattern;

    /**
     * @return List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    public Optional<Output<String>> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }

    /**
     * the identity key of the repo.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return the identity key of the repo.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    @Import(name="notes")
    private @Nullable Output<String> notes;

    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
    }

    @Import(name="packageType")
    private @Nullable Output<String> packageType;

    public Optional<Output<String>> packageType() {
        return Optional.ofNullable(this.packageType);
    }

    /**
     * The primary GPG key to be used to sign packages.
     * 
     */
    @Import(name="primaryKeypairRef")
    private @Nullable Output<String> primaryKeypairRef;

    /**
     * @return The primary GPG key to be used to sign packages.
     * 
     */
    public Optional<Output<String>> primaryKeypairRef() {
        return Optional.ofNullable(this.primaryKeypairRef);
    }

    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    @Import(name="priorityResolution")
    private @Nullable Output<Boolean> priorityResolution;

    /**
     * @return Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    public Optional<Output<Boolean>> priorityResolution() {
        return Optional.ofNullable(this.priorityResolution);
    }

    /**
     * Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
     * 
     */
    @Import(name="projectEnvironments")
    private @Nullable Output<List<String>> projectEnvironments;

    /**
     * @return Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
     * 
     */
    public Optional<Output<List<String>>> projectEnvironments() {
        return Optional.ofNullable(this.projectEnvironments);
    }

    /**
     * Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
     * with project key, separated by a dash.
     * 
     */
    @Import(name="projectKey")
    private @Nullable Output<String> projectKey;

    /**
     * @return Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
     * with project key, separated by a dash.
     * 
     */
    public Optional<Output<String>> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }

    /**
     * List of property set name
     * 
     */
    @Import(name="propertySets")
    private @Nullable Output<List<String>> propertySets;

    /**
     * @return List of property set name
     * 
     */
    public Optional<Output<List<String>>> propertySets() {
        return Optional.ofNullable(this.propertySets);
    }

    /**
     * Repository layout key for the local repository
     * 
     */
    @Import(name="repoLayoutRef")
    private @Nullable Output<String> repoLayoutRef;

    /**
     * @return Repository layout key for the local repository
     * 
     */
    public Optional<Output<String>> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    /**
     * The secondary GPG key to be used to sign packages.
     * 
     */
    @Import(name="secondaryKeypairRef")
    private @Nullable Output<String> secondaryKeypairRef;

    /**
     * @return The secondary GPG key to be used to sign packages.
     * 
     */
    public Optional<Output<String>> secondaryKeypairRef() {
        return Optional.ofNullable(this.secondaryKeypairRef);
    }

    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    @Import(name="xrayIndex")
    private @Nullable Output<Boolean> xrayIndex;

    /**
     * @return Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    public Optional<Output<Boolean>> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    /**
     * A comma separated list of XML file names containing RPM group component definitions.
     * Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
     * generating a gzipped version of the group files, if required. Default is empty string.
     * 
     */
    @Import(name="yumGroupFileNames")
    private @Nullable Output<String> yumGroupFileNames;

    /**
     * @return A comma separated list of XML file names containing RPM group component definitions.
     * Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
     * generating a gzipped version of the group files, if required. Default is empty string.
     * 
     */
    public Optional<Output<String>> yumGroupFileNames() {
        return Optional.ofNullable(this.yumGroupFileNames);
    }

    /**
     * The depth, relative to the repository&#39;s root folder, where RPM metadata is created.
     * This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
     * your RPMs are stored under &#39;fedora/linux/$releasever/$basearch&#39;, specify a depth of 4. Once the number of snapshots
     * exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
     * snapshots are not cleaned up.
     * 
     */
    @Import(name="yumRootDepth")
    private @Nullable Output<Integer> yumRootDepth;

    /**
     * @return The depth, relative to the repository&#39;s root folder, where RPM metadata is created.
     * This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
     * your RPMs are stored under &#39;fedora/linux/$releasever/$basearch&#39;, specify a depth of 4. Once the number of snapshots
     * exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
     * snapshots are not cleaned up.
     * 
     */
    public Optional<Output<Integer>> yumRootDepth() {
        return Optional.ofNullable(this.yumRootDepth);
    }

    private LocalRpmRepositoryState() {}

    private LocalRpmRepositoryState(LocalRpmRepositoryState $) {
        this.archiveBrowsingEnabled = $.archiveBrowsingEnabled;
        this.blackedOut = $.blackedOut;
        this.calculateYumMetadata = $.calculateYumMetadata;
        this.description = $.description;
        this.downloadDirect = $.downloadDirect;
        this.enableFileListsIndexing = $.enableFileListsIndexing;
        this.excludesPattern = $.excludesPattern;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.notes = $.notes;
        this.packageType = $.packageType;
        this.primaryKeypairRef = $.primaryKeypairRef;
        this.priorityResolution = $.priorityResolution;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.propertySets = $.propertySets;
        this.repoLayoutRef = $.repoLayoutRef;
        this.secondaryKeypairRef = $.secondaryKeypairRef;
        this.xrayIndex = $.xrayIndex;
        this.yumGroupFileNames = $.yumGroupFileNames;
        this.yumRootDepth = $.yumRootDepth;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LocalRpmRepositoryState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LocalRpmRepositoryState $;

        public Builder() {
            $ = new LocalRpmRepositoryState();
        }

        public Builder(LocalRpmRepositoryState defaults) {
            $ = new LocalRpmRepositoryState(Objects.requireNonNull(defaults));
        }

        /**
         * @param archiveBrowsingEnabled When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
         * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
         * security (e.g., cross-site scripting attacks).
         * 
         * @return builder
         * 
         */
        public Builder archiveBrowsingEnabled(@Nullable Output<Boolean> archiveBrowsingEnabled) {
            $.archiveBrowsingEnabled = archiveBrowsingEnabled;
            return this;
        }

        /**
         * @param archiveBrowsingEnabled When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
         * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
         * security (e.g., cross-site scripting attacks).
         * 
         * @return builder
         * 
         */
        public Builder archiveBrowsingEnabled(Boolean archiveBrowsingEnabled) {
            return archiveBrowsingEnabled(Output.of(archiveBrowsingEnabled));
        }

        /**
         * @param blackedOut When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
         * 
         * @return builder
         * 
         */
        public Builder blackedOut(@Nullable Output<Boolean> blackedOut) {
            $.blackedOut = blackedOut;
            return this;
        }

        /**
         * @param blackedOut When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
         * 
         * @return builder
         * 
         */
        public Builder blackedOut(Boolean blackedOut) {
            return blackedOut(Output.of(blackedOut));
        }

        /**
         * @param calculateYumMetadata Default: false.
         * 
         * @return builder
         * 
         */
        public Builder calculateYumMetadata(@Nullable Output<Boolean> calculateYumMetadata) {
            $.calculateYumMetadata = calculateYumMetadata;
            return this;
        }

        /**
         * @param calculateYumMetadata Default: false.
         * 
         * @return builder
         * 
         */
        public Builder calculateYumMetadata(Boolean calculateYumMetadata) {
            return calculateYumMetadata(Output.of(calculateYumMetadata));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param downloadDirect When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
         * storage provider. Available in Enterprise+ and Edge licenses only.
         * 
         * @return builder
         * 
         */
        public Builder downloadDirect(@Nullable Output<Boolean> downloadDirect) {
            $.downloadDirect = downloadDirect;
            return this;
        }

        /**
         * @param downloadDirect When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
         * storage provider. Available in Enterprise+ and Edge licenses only.
         * 
         * @return builder
         * 
         */
        public Builder downloadDirect(Boolean downloadDirect) {
            return downloadDirect(Output.of(downloadDirect));
        }

        /**
         * @param enableFileListsIndexing Default: false.
         * 
         * @return builder
         * 
         */
        public Builder enableFileListsIndexing(@Nullable Output<Boolean> enableFileListsIndexing) {
            $.enableFileListsIndexing = enableFileListsIndexing;
            return this;
        }

        /**
         * @param enableFileListsIndexing Default: false.
         * 
         * @return builder
         * 
         */
        public Builder enableFileListsIndexing(Boolean enableFileListsIndexing) {
            return enableFileListsIndexing(Output.of(enableFileListsIndexing));
        }

        /**
         * @param excludesPattern List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
         * artifacts are excluded.
         * 
         * @return builder
         * 
         */
        public Builder excludesPattern(@Nullable Output<String> excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        /**
         * @param excludesPattern List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
         * artifacts are excluded.
         * 
         * @return builder
         * 
         */
        public Builder excludesPattern(String excludesPattern) {
            return excludesPattern(Output.of(excludesPattern));
        }

        /**
         * @param includesPattern List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
         * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
         * 
         * @return builder
         * 
         */
        public Builder includesPattern(@Nullable Output<String> includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        /**
         * @param includesPattern List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
         * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
         * 
         * @return builder
         * 
         */
        public Builder includesPattern(String includesPattern) {
            return includesPattern(Output.of(includesPattern));
        }

        /**
         * @param key the identity key of the repo.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
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

        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        public Builder notes(String notes) {
            return notes(Output.of(notes));
        }

        public Builder packageType(@Nullable Output<String> packageType) {
            $.packageType = packageType;
            return this;
        }

        public Builder packageType(String packageType) {
            return packageType(Output.of(packageType));
        }

        /**
         * @param primaryKeypairRef The primary GPG key to be used to sign packages.
         * 
         * @return builder
         * 
         */
        public Builder primaryKeypairRef(@Nullable Output<String> primaryKeypairRef) {
            $.primaryKeypairRef = primaryKeypairRef;
            return this;
        }

        /**
         * @param primaryKeypairRef The primary GPG key to be used to sign packages.
         * 
         * @return builder
         * 
         */
        public Builder primaryKeypairRef(String primaryKeypairRef) {
            return primaryKeypairRef(Output.of(primaryKeypairRef));
        }

        /**
         * @param priorityResolution Setting repositories with priority will cause metadata to be merged only from repositories set with this field
         * 
         * @return builder
         * 
         */
        public Builder priorityResolution(@Nullable Output<Boolean> priorityResolution) {
            $.priorityResolution = priorityResolution;
            return this;
        }

        /**
         * @param priorityResolution Setting repositories with priority will cause metadata to be merged only from repositories set with this field
         * 
         * @return builder
         * 
         */
        public Builder priorityResolution(Boolean priorityResolution) {
            return priorityResolution(Output.of(priorityResolution));
        }

        /**
         * @param projectEnvironments Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
         * 
         * @return builder
         * 
         */
        public Builder projectEnvironments(@Nullable Output<List<String>> projectEnvironments) {
            $.projectEnvironments = projectEnvironments;
            return this;
        }

        /**
         * @param projectEnvironments Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
         * 
         * @return builder
         * 
         */
        public Builder projectEnvironments(List<String> projectEnvironments) {
            return projectEnvironments(Output.of(projectEnvironments));
        }

        /**
         * @param projectEnvironments Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
         * 
         * @return builder
         * 
         */
        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }

        /**
         * @param projectKey Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
         * with project key, separated by a dash.
         * 
         * @return builder
         * 
         */
        public Builder projectKey(@Nullable Output<String> projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        /**
         * @param projectKey Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
         * with project key, separated by a dash.
         * 
         * @return builder
         * 
         */
        public Builder projectKey(String projectKey) {
            return projectKey(Output.of(projectKey));
        }

        /**
         * @param propertySets List of property set name
         * 
         * @return builder
         * 
         */
        public Builder propertySets(@Nullable Output<List<String>> propertySets) {
            $.propertySets = propertySets;
            return this;
        }

        /**
         * @param propertySets List of property set name
         * 
         * @return builder
         * 
         */
        public Builder propertySets(List<String> propertySets) {
            return propertySets(Output.of(propertySets));
        }

        /**
         * @param propertySets List of property set name
         * 
         * @return builder
         * 
         */
        public Builder propertySets(String... propertySets) {
            return propertySets(List.of(propertySets));
        }

        /**
         * @param repoLayoutRef Repository layout key for the local repository
         * 
         * @return builder
         * 
         */
        public Builder repoLayoutRef(@Nullable Output<String> repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        /**
         * @param repoLayoutRef Repository layout key for the local repository
         * 
         * @return builder
         * 
         */
        public Builder repoLayoutRef(String repoLayoutRef) {
            return repoLayoutRef(Output.of(repoLayoutRef));
        }

        /**
         * @param secondaryKeypairRef The secondary GPG key to be used to sign packages.
         * 
         * @return builder
         * 
         */
        public Builder secondaryKeypairRef(@Nullable Output<String> secondaryKeypairRef) {
            $.secondaryKeypairRef = secondaryKeypairRef;
            return this;
        }

        /**
         * @param secondaryKeypairRef The secondary GPG key to be used to sign packages.
         * 
         * @return builder
         * 
         */
        public Builder secondaryKeypairRef(String secondaryKeypairRef) {
            return secondaryKeypairRef(Output.of(secondaryKeypairRef));
        }

        /**
         * @param xrayIndex Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
         * Xray settings.
         * 
         * @return builder
         * 
         */
        public Builder xrayIndex(@Nullable Output<Boolean> xrayIndex) {
            $.xrayIndex = xrayIndex;
            return this;
        }

        /**
         * @param xrayIndex Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
         * Xray settings.
         * 
         * @return builder
         * 
         */
        public Builder xrayIndex(Boolean xrayIndex) {
            return xrayIndex(Output.of(xrayIndex));
        }

        /**
         * @param yumGroupFileNames A comma separated list of XML file names containing RPM group component definitions.
         * Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
         * generating a gzipped version of the group files, if required. Default is empty string.
         * 
         * @return builder
         * 
         */
        public Builder yumGroupFileNames(@Nullable Output<String> yumGroupFileNames) {
            $.yumGroupFileNames = yumGroupFileNames;
            return this;
        }

        /**
         * @param yumGroupFileNames A comma separated list of XML file names containing RPM group component definitions.
         * Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
         * generating a gzipped version of the group files, if required. Default is empty string.
         * 
         * @return builder
         * 
         */
        public Builder yumGroupFileNames(String yumGroupFileNames) {
            return yumGroupFileNames(Output.of(yumGroupFileNames));
        }

        /**
         * @param yumRootDepth The depth, relative to the repository&#39;s root folder, where RPM metadata is created.
         * This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
         * your RPMs are stored under &#39;fedora/linux/$releasever/$basearch&#39;, specify a depth of 4. Once the number of snapshots
         * exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
         * snapshots are not cleaned up.
         * 
         * @return builder
         * 
         */
        public Builder yumRootDepth(@Nullable Output<Integer> yumRootDepth) {
            $.yumRootDepth = yumRootDepth;
            return this;
        }

        /**
         * @param yumRootDepth The depth, relative to the repository&#39;s root folder, where RPM metadata is created.
         * This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
         * your RPMs are stored under &#39;fedora/linux/$releasever/$basearch&#39;, specify a depth of 4. Once the number of snapshots
         * exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
         * snapshots are not cleaned up.
         * 
         * @return builder
         * 
         */
        public Builder yumRootDepth(Integer yumRootDepth) {
            return yumRootDepth(Output.of(yumRootDepth));
        }

        public LocalRpmRepositoryState build() {
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

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


public final class LocalGradleRepositoryArgs extends com.pulumi.resources.ResourceArgs {

    public static final LocalGradleRepositoryArgs Empty = new LocalGradleRepositoryArgs();

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
     * When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;
     * 
     */
    @Import(name="cdnRedirect")
    private @Nullable Output<Boolean> cdnRedirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;
     * 
     */
    public Optional<Output<Boolean>> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }

    /**
     * Checksum policy determines how Artifactory behaves when a client checksum for a deployed
     * resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
     * `client-checksums` and `generated-checksums`. For more details,
     * please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
     * 
     */
    @Import(name="checksumPolicyType")
    private @Nullable Output<String> checksumPolicyType;

    /**
     * @return Checksum policy determines how Artifactory behaves when a client checksum for a deployed
     * resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
     * `client-checksums` and `generated-checksums`. For more details,
     * please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
     * 
     */
    public Optional<Output<String>> checksumPolicyType() {
        return Optional.ofNullable(this.checksumPolicyType);
    }

    /**
     * Public description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Public description.
     * 
     */
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
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
     * artifacts are excluded.
     * 
     */
    @Import(name="excludesPattern")
    private @Nullable Output<String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
     * artifacts are excluded.
     * 
     */
    public Optional<Output<String>> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    /**
     * If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
     * 
     */
    @Import(name="handleReleases")
    private @Nullable Output<Boolean> handleReleases;

    /**
     * @return If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
     * 
     */
    public Optional<Output<Boolean>> handleReleases() {
        return Optional.ofNullable(this.handleReleases);
    }

    /**
     * If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
     * 
     */
    @Import(name="handleSnapshots")
    private @Nullable Output<Boolean> handleSnapshots;

    /**
     * @return If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
     * 
     */
    public Optional<Output<Boolean>> handleSnapshots() {
        return Optional.ofNullable(this.handleSnapshots);
    }

    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Import(name="includesPattern")
    private @Nullable Output<String> includesPattern;

    /**
     * @return List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
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
     * The maximum number of unique snapshots of a single artifact to store.
     * Once the number of snapshots exceeds this setting, older versions are removed.
     * A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
     * 
     */
    @Import(name="maxUniqueSnapshots")
    private @Nullable Output<Integer> maxUniqueSnapshots;

    /**
     * @return The maximum number of unique snapshots of a single artifact to store.
     * Once the number of snapshots exceeds this setting, older versions are removed.
     * A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
     * 
     */
    public Optional<Output<Integer>> maxUniqueSnapshots() {
        return Optional.ofNullable(this.maxUniqueSnapshots);
    }

    /**
     * Internal description.
     * 
     */
    @Import(name="notes")
    private @Nullable Output<String> notes;

    /**
     * @return Internal description.
     * 
     */
    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
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
     * Project environment for assigning this repository to. Allow values: &#34;DEV&#34;, &#34;PROD&#34;, or one of custom environment. Before
     * Artifactory 7.53.1, up to 2 values (&#34;DEV&#34; and &#34;PROD&#34;) are allowed. From 7.53.1 onward, only one value is allowed. The
     * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
     * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
     * 
     */
    @Import(name="projectEnvironments")
    private @Nullable Output<List<String>> projectEnvironments;

    /**
     * @return Project environment for assigning this repository to. Allow values: &#34;DEV&#34;, &#34;PROD&#34;, or one of custom environment. Before
     * Artifactory 7.53.1, up to 2 values (&#34;DEV&#34; and &#34;PROD&#34;) are allowed. From 7.53.1 onward, only one value is allowed. The
     * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
     * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
     * 
     */
    public Optional<Output<List<String>>> projectEnvironments() {
        return Optional.ofNullable(this.projectEnvironments);
    }

    /**
     * Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    @Import(name="projectKey")
    private @Nullable Output<String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
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
     * Specifies the naming convention for Maven SNAPSHOT versions.
     * The options are -
     * 
     */
    @Import(name="snapshotVersionBehavior")
    private @Nullable Output<String> snapshotVersionBehavior;

    /**
     * @return Specifies the naming convention for Maven SNAPSHOT versions.
     * The options are -
     * 
     */
    public Optional<Output<String>> snapshotVersionBehavior() {
        return Optional.ofNullable(this.snapshotVersionBehavior);
    }

    /**
     * By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
     * If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error.
     * You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
     * 
     */
    @Import(name="suppressPomConsistencyChecks")
    private @Nullable Output<Boolean> suppressPomConsistencyChecks;

    /**
     * @return By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
     * If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error.
     * You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
     * 
     */
    public Optional<Output<Boolean>> suppressPomConsistencyChecks() {
        return Optional.ofNullable(this.suppressPomConsistencyChecks);
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

    private LocalGradleRepositoryArgs() {}

    private LocalGradleRepositoryArgs(LocalGradleRepositoryArgs $) {
        this.archiveBrowsingEnabled = $.archiveBrowsingEnabled;
        this.blackedOut = $.blackedOut;
        this.cdnRedirect = $.cdnRedirect;
        this.checksumPolicyType = $.checksumPolicyType;
        this.description = $.description;
        this.downloadDirect = $.downloadDirect;
        this.excludesPattern = $.excludesPattern;
        this.handleReleases = $.handleReleases;
        this.handleSnapshots = $.handleSnapshots;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.maxUniqueSnapshots = $.maxUniqueSnapshots;
        this.notes = $.notes;
        this.priorityResolution = $.priorityResolution;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.propertySets = $.propertySets;
        this.repoLayoutRef = $.repoLayoutRef;
        this.snapshotVersionBehavior = $.snapshotVersionBehavior;
        this.suppressPomConsistencyChecks = $.suppressPomConsistencyChecks;
        this.xrayIndex = $.xrayIndex;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LocalGradleRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LocalGradleRepositoryArgs $;

        public Builder() {
            $ = new LocalGradleRepositoryArgs();
        }

        public Builder(LocalGradleRepositoryArgs defaults) {
            $ = new LocalGradleRepositoryArgs(Objects.requireNonNull(defaults));
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
         * @param cdnRedirect When set, download requests to this repository will redirect the client to download the artifact directly from AWS
         * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder cdnRedirect(@Nullable Output<Boolean> cdnRedirect) {
            $.cdnRedirect = cdnRedirect;
            return this;
        }

        /**
         * @param cdnRedirect When set, download requests to this repository will redirect the client to download the artifact directly from AWS
         * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder cdnRedirect(Boolean cdnRedirect) {
            return cdnRedirect(Output.of(cdnRedirect));
        }

        /**
         * @param checksumPolicyType Checksum policy determines how Artifactory behaves when a client checksum for a deployed
         * resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
         * `client-checksums` and `generated-checksums`. For more details,
         * please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
         * 
         * @return builder
         * 
         */
        public Builder checksumPolicyType(@Nullable Output<String> checksumPolicyType) {
            $.checksumPolicyType = checksumPolicyType;
            return this;
        }

        /**
         * @param checksumPolicyType Checksum policy determines how Artifactory behaves when a client checksum for a deployed
         * resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
         * `client-checksums` and `generated-checksums`. For more details,
         * please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
         * 
         * @return builder
         * 
         */
        public Builder checksumPolicyType(String checksumPolicyType) {
            return checksumPolicyType(Output.of(checksumPolicyType));
        }

        /**
         * @param description Public description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Public description.
         * 
         * @return builder
         * 
         */
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
         * @param excludesPattern List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
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
         * @param excludesPattern List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
         * artifacts are excluded.
         * 
         * @return builder
         * 
         */
        public Builder excludesPattern(String excludesPattern) {
            return excludesPattern(Output.of(excludesPattern));
        }

        /**
         * @param handleReleases If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
         * 
         * @return builder
         * 
         */
        public Builder handleReleases(@Nullable Output<Boolean> handleReleases) {
            $.handleReleases = handleReleases;
            return this;
        }

        /**
         * @param handleReleases If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
         * 
         * @return builder
         * 
         */
        public Builder handleReleases(Boolean handleReleases) {
            return handleReleases(Output.of(handleReleases));
        }

        /**
         * @param handleSnapshots If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
         * 
         * @return builder
         * 
         */
        public Builder handleSnapshots(@Nullable Output<Boolean> handleSnapshots) {
            $.handleSnapshots = handleSnapshots;
            return this;
        }

        /**
         * @param handleSnapshots If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
         * 
         * @return builder
         * 
         */
        public Builder handleSnapshots(Boolean handleSnapshots) {
            return handleSnapshots(Output.of(handleSnapshots));
        }

        /**
         * @param includesPattern List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
         * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
         * 
         * @return builder
         * 
         */
        public Builder includesPattern(@Nullable Output<String> includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        /**
         * @param includesPattern List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
         * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
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
         * @param maxUniqueSnapshots The maximum number of unique snapshots of a single artifact to store.
         * Once the number of snapshots exceeds this setting, older versions are removed.
         * A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
         * 
         * @return builder
         * 
         */
        public Builder maxUniqueSnapshots(@Nullable Output<Integer> maxUniqueSnapshots) {
            $.maxUniqueSnapshots = maxUniqueSnapshots;
            return this;
        }

        /**
         * @param maxUniqueSnapshots The maximum number of unique snapshots of a single artifact to store.
         * Once the number of snapshots exceeds this setting, older versions are removed.
         * A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
         * 
         * @return builder
         * 
         */
        public Builder maxUniqueSnapshots(Integer maxUniqueSnapshots) {
            return maxUniqueSnapshots(Output.of(maxUniqueSnapshots));
        }

        /**
         * @param notes Internal description.
         * 
         * @return builder
         * 
         */
        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        /**
         * @param notes Internal description.
         * 
         * @return builder
         * 
         */
        public Builder notes(String notes) {
            return notes(Output.of(notes));
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
         * @param projectEnvironments Project environment for assigning this repository to. Allow values: &#34;DEV&#34;, &#34;PROD&#34;, or one of custom environment. Before
         * Artifactory 7.53.1, up to 2 values (&#34;DEV&#34; and &#34;PROD&#34;) are allowed. From 7.53.1 onward, only one value is allowed. The
         * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
         * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
         * 
         * @return builder
         * 
         */
        public Builder projectEnvironments(@Nullable Output<List<String>> projectEnvironments) {
            $.projectEnvironments = projectEnvironments;
            return this;
        }

        /**
         * @param projectEnvironments Project environment for assigning this repository to. Allow values: &#34;DEV&#34;, &#34;PROD&#34;, or one of custom environment. Before
         * Artifactory 7.53.1, up to 2 values (&#34;DEV&#34; and &#34;PROD&#34;) are allowed. From 7.53.1 onward, only one value is allowed. The
         * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
         * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
         * 
         * @return builder
         * 
         */
        public Builder projectEnvironments(List<String> projectEnvironments) {
            return projectEnvironments(Output.of(projectEnvironments));
        }

        /**
         * @param projectEnvironments Project environment for assigning this repository to. Allow values: &#34;DEV&#34;, &#34;PROD&#34;, or one of custom environment. Before
         * Artifactory 7.53.1, up to 2 values (&#34;DEV&#34; and &#34;PROD&#34;) are allowed. From 7.53.1 onward, only one value is allowed. The
         * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
         * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
         * 
         * @return builder
         * 
         */
        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }

        /**
         * @param projectKey Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
         * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
         * 
         * @return builder
         * 
         */
        public Builder projectKey(@Nullable Output<String> projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        /**
         * @param projectKey Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
         * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
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
         * @param snapshotVersionBehavior Specifies the naming convention for Maven SNAPSHOT versions.
         * The options are -
         * 
         * @return builder
         * 
         */
        public Builder snapshotVersionBehavior(@Nullable Output<String> snapshotVersionBehavior) {
            $.snapshotVersionBehavior = snapshotVersionBehavior;
            return this;
        }

        /**
         * @param snapshotVersionBehavior Specifies the naming convention for Maven SNAPSHOT versions.
         * The options are -
         * 
         * @return builder
         * 
         */
        public Builder snapshotVersionBehavior(String snapshotVersionBehavior) {
            return snapshotVersionBehavior(Output.of(snapshotVersionBehavior));
        }

        /**
         * @param suppressPomConsistencyChecks By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
         * If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error.
         * You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
         * 
         * @return builder
         * 
         */
        public Builder suppressPomConsistencyChecks(@Nullable Output<Boolean> suppressPomConsistencyChecks) {
            $.suppressPomConsistencyChecks = suppressPomConsistencyChecks;
            return this;
        }

        /**
         * @param suppressPomConsistencyChecks By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
         * If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error.
         * You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
         * 
         * @return builder
         * 
         */
        public Builder suppressPomConsistencyChecks(Boolean suppressPomConsistencyChecks) {
            return suppressPomConsistencyChecks(Output.of(suppressPomConsistencyChecks));
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

        public LocalGradleRepositoryArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("LocalGradleRepositoryArgs", "key");
            }
            return $;
        }
    }

}

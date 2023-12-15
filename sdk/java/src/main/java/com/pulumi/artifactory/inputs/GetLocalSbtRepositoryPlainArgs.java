// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLocalSbtRepositoryPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLocalSbtRepositoryPlainArgs Empty = new GetLocalSbtRepositoryPlainArgs();

    @Import(name="archiveBrowsingEnabled")
    private @Nullable Boolean archiveBrowsingEnabled;

    public Optional<Boolean> archiveBrowsingEnabled() {
        return Optional.ofNullable(this.archiveBrowsingEnabled);
    }

    @Import(name="blackedOut")
    private @Nullable Boolean blackedOut;

    public Optional<Boolean> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }

    @Import(name="cdnRedirect")
    private @Nullable Boolean cdnRedirect;

    public Optional<Boolean> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }

    /**
     * Checksum policy determines how Artifactory behaves when a client checksum for a deployed
     * resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
     * `client-checksums` and `generated-checksums`. For more details, please refer
     * to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
     * .
     * 
     */
    @Import(name="checksumPolicyType")
    private @Nullable String checksumPolicyType;

    /**
     * @return Checksum policy determines how Artifactory behaves when a client checksum for a deployed
     * resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
     * `client-checksums` and `generated-checksums`. For more details, please refer
     * to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
     * .
     * 
     */
    public Optional<String> checksumPolicyType() {
        return Optional.ofNullable(this.checksumPolicyType);
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

    @Import(name="excludesPattern")
    private @Nullable String excludesPattern;

    public Optional<String> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    /**
     * If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
     * .
     * 
     */
    @Import(name="handleReleases")
    private @Nullable Boolean handleReleases;

    /**
     * @return If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
     * .
     * 
     */
    public Optional<Boolean> handleReleases() {
        return Optional.ofNullable(this.handleReleases);
    }

    /**
     * If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
     * is `true`.
     * 
     */
    @Import(name="handleSnapshots")
    private @Nullable Boolean handleSnapshots;

    /**
     * @return If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
     * is `true`.
     * 
     */
    public Optional<Boolean> handleSnapshots() {
        return Optional.ofNullable(this.handleSnapshots);
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

    /**
     * The maximum number of unique snapshots of a single artifact to store. Once the number of
     * snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
     * unique snapshots are not cleaned up.
     * 
     */
    @Import(name="maxUniqueSnapshots")
    private @Nullable Integer maxUniqueSnapshots;

    /**
     * @return The maximum number of unique snapshots of a single artifact to store. Once the number of
     * snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
     * unique snapshots are not cleaned up.
     * 
     */
    public Optional<Integer> maxUniqueSnapshots() {
        return Optional.ofNullable(this.maxUniqueSnapshots);
    }

    @Import(name="notes")
    private @Nullable String notes;

    public Optional<String> notes() {
        return Optional.ofNullable(this.notes);
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

    @Import(name="propertySets")
    private @Nullable List<String> propertySets;

    public Optional<List<String>> propertySets() {
        return Optional.ofNullable(this.propertySets);
    }

    @Import(name="repoLayoutRef")
    private @Nullable String repoLayoutRef;

    public Optional<String> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    /**
     * Specifies the naming convention for Maven SNAPSHOT versions. The options are
     * ---
     * 
     */
    @Import(name="snapshotVersionBehavior")
    private @Nullable String snapshotVersionBehavior;

    /**
     * @return Specifies the naming convention for Maven SNAPSHOT versions. The options are
     * ---
     * 
     */
    public Optional<String> snapshotVersionBehavior() {
        return Optional.ofNullable(this.snapshotVersionBehavior);
    }

    /**
     * By default, Artifactory keeps your repositories healthy by refusing POMs with
     * incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
     * path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error. You can disable this behavior by setting the
     * Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
     * 
     */
    @Import(name="suppressPomConsistencyChecks")
    private @Nullable Boolean suppressPomConsistencyChecks;

    /**
     * @return By default, Artifactory keeps your repositories healthy by refusing POMs with
     * incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
     * path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error. You can disable this behavior by setting the
     * Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
     * 
     */
    public Optional<Boolean> suppressPomConsistencyChecks() {
        return Optional.ofNullable(this.suppressPomConsistencyChecks);
    }

    @Import(name="xrayIndex")
    private @Nullable Boolean xrayIndex;

    public Optional<Boolean> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    private GetLocalSbtRepositoryPlainArgs() {}

    private GetLocalSbtRepositoryPlainArgs(GetLocalSbtRepositoryPlainArgs $) {
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
    public static Builder builder(GetLocalSbtRepositoryPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLocalSbtRepositoryPlainArgs $;

        public Builder() {
            $ = new GetLocalSbtRepositoryPlainArgs();
        }

        public Builder(GetLocalSbtRepositoryPlainArgs defaults) {
            $ = new GetLocalSbtRepositoryPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder archiveBrowsingEnabled(@Nullable Boolean archiveBrowsingEnabled) {
            $.archiveBrowsingEnabled = archiveBrowsingEnabled;
            return this;
        }

        public Builder blackedOut(@Nullable Boolean blackedOut) {
            $.blackedOut = blackedOut;
            return this;
        }

        public Builder cdnRedirect(@Nullable Boolean cdnRedirect) {
            $.cdnRedirect = cdnRedirect;
            return this;
        }

        /**
         * @param checksumPolicyType Checksum policy determines how Artifactory behaves when a client checksum for a deployed
         * resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
         * `client-checksums` and `generated-checksums`. For more details, please refer
         * to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
         * .
         * 
         * @return builder
         * 
         */
        public Builder checksumPolicyType(@Nullable String checksumPolicyType) {
            $.checksumPolicyType = checksumPolicyType;
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

        public Builder excludesPattern(@Nullable String excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        /**
         * @param handleReleases If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
         * .
         * 
         * @return builder
         * 
         */
        public Builder handleReleases(@Nullable Boolean handleReleases) {
            $.handleReleases = handleReleases;
            return this;
        }

        /**
         * @param handleSnapshots If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
         * is `true`.
         * 
         * @return builder
         * 
         */
        public Builder handleSnapshots(@Nullable Boolean handleSnapshots) {
            $.handleSnapshots = handleSnapshots;
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

        /**
         * @param maxUniqueSnapshots The maximum number of unique snapshots of a single artifact to store. Once the number of
         * snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
         * unique snapshots are not cleaned up.
         * 
         * @return builder
         * 
         */
        public Builder maxUniqueSnapshots(@Nullable Integer maxUniqueSnapshots) {
            $.maxUniqueSnapshots = maxUniqueSnapshots;
            return this;
        }

        public Builder notes(@Nullable String notes) {
            $.notes = notes;
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

        public Builder propertySets(@Nullable List<String> propertySets) {
            $.propertySets = propertySets;
            return this;
        }

        public Builder propertySets(String... propertySets) {
            return propertySets(List.of(propertySets));
        }

        public Builder repoLayoutRef(@Nullable String repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        /**
         * @param snapshotVersionBehavior Specifies the naming convention for Maven SNAPSHOT versions. The options are
         * ---
         * 
         * @return builder
         * 
         */
        public Builder snapshotVersionBehavior(@Nullable String snapshotVersionBehavior) {
            $.snapshotVersionBehavior = snapshotVersionBehavior;
            return this;
        }

        /**
         * @param suppressPomConsistencyChecks By default, Artifactory keeps your repositories healthy by refusing POMs with
         * incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
         * path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error. You can disable this behavior by setting the
         * Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
         * 
         * @return builder
         * 
         */
        public Builder suppressPomConsistencyChecks(@Nullable Boolean suppressPomConsistencyChecks) {
            $.suppressPomConsistencyChecks = suppressPomConsistencyChecks;
            return this;
        }

        public Builder xrayIndex(@Nullable Boolean xrayIndex) {
            $.xrayIndex = xrayIndex;
            return this;
        }

        public GetLocalSbtRepositoryPlainArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}

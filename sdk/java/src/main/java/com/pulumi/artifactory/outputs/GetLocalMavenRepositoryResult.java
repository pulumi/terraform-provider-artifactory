// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

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
public final class GetLocalMavenRepositoryResult {
    private @Nullable Boolean archiveBrowsingEnabled;
    private @Nullable Boolean blackedOut;
    private @Nullable Boolean cdnRedirect;
    /**
     * @return Checksum policy determines how Artifactory behaves when a client checksum for a
     * deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are:
     * - `client-checksums`
     * - `server-generated-checksums`. For more details, please refer
     *   to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
     *   .
     * 
     */
    private @Nullable String checksumPolicyType;
    private @Nullable String description;
    private @Nullable Boolean downloadDirect;
    private @Nullable String excludesPattern;
    /**
     * @return If set, Artifactory allows you to deploy release artifacts into this repository.
     * Default is `true`.
     * 
     */
    private @Nullable Boolean handleReleases;
    /**
     * @return If set, Artifactory allows you to deploy snapshot artifacts into this repository.
     * Default is `true`.
     * 
     */
    private @Nullable Boolean handleSnapshots;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String includesPattern;
    private String key;
    /**
     * @return The maximum number of unique snapshots of a single artifact to store. Once the
     * number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no
     * limit, and unique snapshots are not cleaned up.
     * 
     */
    private @Nullable Integer maxUniqueSnapshots;
    private @Nullable String notes;
    private String packageType;
    private @Nullable Boolean priorityResolution;
    private List<String> projectEnvironments;
    private @Nullable String projectKey;
    private @Nullable List<String> propertySets;
    private @Nullable String repoLayoutRef;
    /**
     * @return Specifies the naming convention for Maven SNAPSHOT versions. The options are
     * ---
     * 
     */
    private @Nullable String snapshotVersionBehavior;
    /**
     * @return By default, Artifactory keeps your repositories healthy by refusing
     * POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match
     * the deployed path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error. You can disable this behavior by
     * setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository.
     * 
     */
    private @Nullable Boolean suppressPomConsistencyChecks;
    private @Nullable Boolean xrayIndex;

    private GetLocalMavenRepositoryResult() {}
    public Optional<Boolean> archiveBrowsingEnabled() {
        return Optional.ofNullable(this.archiveBrowsingEnabled);
    }
    public Optional<Boolean> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }
    public Optional<Boolean> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }
    /**
     * @return Checksum policy determines how Artifactory behaves when a client checksum for a
     * deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are:
     * - `client-checksums`
     * - `server-generated-checksums`. For more details, please refer
     *   to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
     *   .
     * 
     */
    public Optional<String> checksumPolicyType() {
        return Optional.ofNullable(this.checksumPolicyType);
    }
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    public Optional<Boolean> downloadDirect() {
        return Optional.ofNullable(this.downloadDirect);
    }
    public Optional<String> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }
    /**
     * @return If set, Artifactory allows you to deploy release artifacts into this repository.
     * Default is `true`.
     * 
     */
    public Optional<Boolean> handleReleases() {
        return Optional.ofNullable(this.handleReleases);
    }
    /**
     * @return If set, Artifactory allows you to deploy snapshot artifacts into this repository.
     * Default is `true`.
     * 
     */
    public Optional<Boolean> handleSnapshots() {
        return Optional.ofNullable(this.handleSnapshots);
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
    /**
     * @return The maximum number of unique snapshots of a single artifact to store. Once the
     * number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no
     * limit, and unique snapshots are not cleaned up.
     * 
     */
    public Optional<Integer> maxUniqueSnapshots() {
        return Optional.ofNullable(this.maxUniqueSnapshots);
    }
    public Optional<String> notes() {
        return Optional.ofNullable(this.notes);
    }
    public String packageType() {
        return this.packageType;
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
    public List<String> propertySets() {
        return this.propertySets == null ? List.of() : this.propertySets;
    }
    public Optional<String> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }
    /**
     * @return Specifies the naming convention for Maven SNAPSHOT versions. The options are
     * ---
     * 
     */
    public Optional<String> snapshotVersionBehavior() {
        return Optional.ofNullable(this.snapshotVersionBehavior);
    }
    /**
     * @return By default, Artifactory keeps your repositories healthy by refusing
     * POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match
     * the deployed path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error. You can disable this behavior by
     * setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository.
     * 
     */
    public Optional<Boolean> suppressPomConsistencyChecks() {
        return Optional.ofNullable(this.suppressPomConsistencyChecks);
    }
    public Optional<Boolean> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLocalMavenRepositoryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean archiveBrowsingEnabled;
        private @Nullable Boolean blackedOut;
        private @Nullable Boolean cdnRedirect;
        private @Nullable String checksumPolicyType;
        private @Nullable String description;
        private @Nullable Boolean downloadDirect;
        private @Nullable String excludesPattern;
        private @Nullable Boolean handleReleases;
        private @Nullable Boolean handleSnapshots;
        private String id;
        private @Nullable String includesPattern;
        private String key;
        private @Nullable Integer maxUniqueSnapshots;
        private @Nullable String notes;
        private String packageType;
        private @Nullable Boolean priorityResolution;
        private List<String> projectEnvironments;
        private @Nullable String projectKey;
        private @Nullable List<String> propertySets;
        private @Nullable String repoLayoutRef;
        private @Nullable String snapshotVersionBehavior;
        private @Nullable Boolean suppressPomConsistencyChecks;
        private @Nullable Boolean xrayIndex;
        public Builder() {}
        public Builder(GetLocalMavenRepositoryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.archiveBrowsingEnabled = defaults.archiveBrowsingEnabled;
    	      this.blackedOut = defaults.blackedOut;
    	      this.cdnRedirect = defaults.cdnRedirect;
    	      this.checksumPolicyType = defaults.checksumPolicyType;
    	      this.description = defaults.description;
    	      this.downloadDirect = defaults.downloadDirect;
    	      this.excludesPattern = defaults.excludesPattern;
    	      this.handleReleases = defaults.handleReleases;
    	      this.handleSnapshots = defaults.handleSnapshots;
    	      this.id = defaults.id;
    	      this.includesPattern = defaults.includesPattern;
    	      this.key = defaults.key;
    	      this.maxUniqueSnapshots = defaults.maxUniqueSnapshots;
    	      this.notes = defaults.notes;
    	      this.packageType = defaults.packageType;
    	      this.priorityResolution = defaults.priorityResolution;
    	      this.projectEnvironments = defaults.projectEnvironments;
    	      this.projectKey = defaults.projectKey;
    	      this.propertySets = defaults.propertySets;
    	      this.repoLayoutRef = defaults.repoLayoutRef;
    	      this.snapshotVersionBehavior = defaults.snapshotVersionBehavior;
    	      this.suppressPomConsistencyChecks = defaults.suppressPomConsistencyChecks;
    	      this.xrayIndex = defaults.xrayIndex;
        }

        @CustomType.Setter
        public Builder archiveBrowsingEnabled(@Nullable Boolean archiveBrowsingEnabled) {

            this.archiveBrowsingEnabled = archiveBrowsingEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder blackedOut(@Nullable Boolean blackedOut) {

            this.blackedOut = blackedOut;
            return this;
        }
        @CustomType.Setter
        public Builder cdnRedirect(@Nullable Boolean cdnRedirect) {

            this.cdnRedirect = cdnRedirect;
            return this;
        }
        @CustomType.Setter
        public Builder checksumPolicyType(@Nullable String checksumPolicyType) {

            this.checksumPolicyType = checksumPolicyType;
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
        public Builder excludesPattern(@Nullable String excludesPattern) {

            this.excludesPattern = excludesPattern;
            return this;
        }
        @CustomType.Setter
        public Builder handleReleases(@Nullable Boolean handleReleases) {

            this.handleReleases = handleReleases;
            return this;
        }
        @CustomType.Setter
        public Builder handleSnapshots(@Nullable Boolean handleSnapshots) {

            this.handleSnapshots = handleSnapshots;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetLocalMavenRepositoryResult", "id");
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
              throw new MissingRequiredPropertyException("GetLocalMavenRepositoryResult", "key");
            }
            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder maxUniqueSnapshots(@Nullable Integer maxUniqueSnapshots) {

            this.maxUniqueSnapshots = maxUniqueSnapshots;
            return this;
        }
        @CustomType.Setter
        public Builder notes(@Nullable String notes) {

            this.notes = notes;
            return this;
        }
        @CustomType.Setter
        public Builder packageType(String packageType) {
            if (packageType == null) {
              throw new MissingRequiredPropertyException("GetLocalMavenRepositoryResult", "packageType");
            }
            this.packageType = packageType;
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
              throw new MissingRequiredPropertyException("GetLocalMavenRepositoryResult", "projectEnvironments");
            }
            this.projectEnvironments = projectEnvironments;
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
        public Builder propertySets(@Nullable List<String> propertySets) {

            this.propertySets = propertySets;
            return this;
        }
        public Builder propertySets(String... propertySets) {
            return propertySets(List.of(propertySets));
        }
        @CustomType.Setter
        public Builder repoLayoutRef(@Nullable String repoLayoutRef) {

            this.repoLayoutRef = repoLayoutRef;
            return this;
        }
        @CustomType.Setter
        public Builder snapshotVersionBehavior(@Nullable String snapshotVersionBehavior) {

            this.snapshotVersionBehavior = snapshotVersionBehavior;
            return this;
        }
        @CustomType.Setter
        public Builder suppressPomConsistencyChecks(@Nullable Boolean suppressPomConsistencyChecks) {

            this.suppressPomConsistencyChecks = suppressPomConsistencyChecks;
            return this;
        }
        @CustomType.Setter
        public Builder xrayIndex(@Nullable Boolean xrayIndex) {

            this.xrayIndex = xrayIndex;
            return this;
        }
        public GetLocalMavenRepositoryResult build() {
            final var _resultValue = new GetLocalMavenRepositoryResult();
            _resultValue.archiveBrowsingEnabled = archiveBrowsingEnabled;
            _resultValue.blackedOut = blackedOut;
            _resultValue.cdnRedirect = cdnRedirect;
            _resultValue.checksumPolicyType = checksumPolicyType;
            _resultValue.description = description;
            _resultValue.downloadDirect = downloadDirect;
            _resultValue.excludesPattern = excludesPattern;
            _resultValue.handleReleases = handleReleases;
            _resultValue.handleSnapshots = handleSnapshots;
            _resultValue.id = id;
            _resultValue.includesPattern = includesPattern;
            _resultValue.key = key;
            _resultValue.maxUniqueSnapshots = maxUniqueSnapshots;
            _resultValue.notes = notes;
            _resultValue.packageType = packageType;
            _resultValue.priorityResolution = priorityResolution;
            _resultValue.projectEnvironments = projectEnvironments;
            _resultValue.projectKey = projectKey;
            _resultValue.propertySets = propertySets;
            _resultValue.repoLayoutRef = repoLayoutRef;
            _resultValue.snapshotVersionBehavior = snapshotVersionBehavior;
            _resultValue.suppressPomConsistencyChecks = suppressPomConsistencyChecks;
            _resultValue.xrayIndex = xrayIndex;
            return _resultValue;
        }
    }
}

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
public final class GetLocalDockerV2RepositoryResult {
    /**
     * @return &#34;The Docker API version to use. This cannot be set&#34;
     * 
     */
    private String apiVersion;
    private @Nullable Boolean archiveBrowsingEnabled;
    private @Nullable Boolean blackedOut;
    /**
     * @return When set, Artifactory will block the pushing of Docker images with manifest v2
     * schema 1 to this repository.
     * 
     */
    private Boolean blockPushingSchema1;
    private @Nullable Boolean cdnRedirect;
    private @Nullable String description;
    private @Nullable Boolean downloadDirect;
    private @Nullable String excludesPattern;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String includesPattern;
    private String key;
    /**
     * @return The maximum number of unique tags of a single Docker image to store in this repository.
     * Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there
     * is no limit. This only applies to manifest v2.
     * 
     */
    private @Nullable Integer maxUniqueTags;
    private @Nullable String notes;
    private String packageType;
    private @Nullable Boolean priorityResolution;
    private List<String> projectEnvironments;
    private @Nullable String projectKey;
    private @Nullable List<String> propertySets;
    private @Nullable String repoLayoutRef;
    /**
     * @return If greater than 1, overwritten tags will be saved by their digest, up to the set up
     * number. This only applies to manifest V2.
     * 
     */
    private @Nullable Integer tagRetention;
    private @Nullable Boolean xrayIndex;

    private GetLocalDockerV2RepositoryResult() {}
    /**
     * @return &#34;The Docker API version to use. This cannot be set&#34;
     * 
     */
    public String apiVersion() {
        return this.apiVersion;
    }
    public Optional<Boolean> archiveBrowsingEnabled() {
        return Optional.ofNullable(this.archiveBrowsingEnabled);
    }
    public Optional<Boolean> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }
    /**
     * @return When set, Artifactory will block the pushing of Docker images with manifest v2
     * schema 1 to this repository.
     * 
     */
    public Boolean blockPushingSchema1() {
        return this.blockPushingSchema1;
    }
    public Optional<Boolean> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
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
     * @return The maximum number of unique tags of a single Docker image to store in this repository.
     * Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there
     * is no limit. This only applies to manifest v2.
     * 
     */
    public Optional<Integer> maxUniqueTags() {
        return Optional.ofNullable(this.maxUniqueTags);
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
     * @return If greater than 1, overwritten tags will be saved by their digest, up to the set up
     * number. This only applies to manifest V2.
     * 
     */
    public Optional<Integer> tagRetention() {
        return Optional.ofNullable(this.tagRetention);
    }
    public Optional<Boolean> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLocalDockerV2RepositoryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String apiVersion;
        private @Nullable Boolean archiveBrowsingEnabled;
        private @Nullable Boolean blackedOut;
        private Boolean blockPushingSchema1;
        private @Nullable Boolean cdnRedirect;
        private @Nullable String description;
        private @Nullable Boolean downloadDirect;
        private @Nullable String excludesPattern;
        private String id;
        private @Nullable String includesPattern;
        private String key;
        private @Nullable Integer maxUniqueTags;
        private @Nullable String notes;
        private String packageType;
        private @Nullable Boolean priorityResolution;
        private List<String> projectEnvironments;
        private @Nullable String projectKey;
        private @Nullable List<String> propertySets;
        private @Nullable String repoLayoutRef;
        private @Nullable Integer tagRetention;
        private @Nullable Boolean xrayIndex;
        public Builder() {}
        public Builder(GetLocalDockerV2RepositoryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiVersion = defaults.apiVersion;
    	      this.archiveBrowsingEnabled = defaults.archiveBrowsingEnabled;
    	      this.blackedOut = defaults.blackedOut;
    	      this.blockPushingSchema1 = defaults.blockPushingSchema1;
    	      this.cdnRedirect = defaults.cdnRedirect;
    	      this.description = defaults.description;
    	      this.downloadDirect = defaults.downloadDirect;
    	      this.excludesPattern = defaults.excludesPattern;
    	      this.id = defaults.id;
    	      this.includesPattern = defaults.includesPattern;
    	      this.key = defaults.key;
    	      this.maxUniqueTags = defaults.maxUniqueTags;
    	      this.notes = defaults.notes;
    	      this.packageType = defaults.packageType;
    	      this.priorityResolution = defaults.priorityResolution;
    	      this.projectEnvironments = defaults.projectEnvironments;
    	      this.projectKey = defaults.projectKey;
    	      this.propertySets = defaults.propertySets;
    	      this.repoLayoutRef = defaults.repoLayoutRef;
    	      this.tagRetention = defaults.tagRetention;
    	      this.xrayIndex = defaults.xrayIndex;
        }

        @CustomType.Setter
        public Builder apiVersion(String apiVersion) {
            if (apiVersion == null) {
              throw new MissingRequiredPropertyException("GetLocalDockerV2RepositoryResult", "apiVersion");
            }
            this.apiVersion = apiVersion;
            return this;
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
        public Builder blockPushingSchema1(Boolean blockPushingSchema1) {
            if (blockPushingSchema1 == null) {
              throw new MissingRequiredPropertyException("GetLocalDockerV2RepositoryResult", "blockPushingSchema1");
            }
            this.blockPushingSchema1 = blockPushingSchema1;
            return this;
        }
        @CustomType.Setter
        public Builder cdnRedirect(@Nullable Boolean cdnRedirect) {

            this.cdnRedirect = cdnRedirect;
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
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetLocalDockerV2RepositoryResult", "id");
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
              throw new MissingRequiredPropertyException("GetLocalDockerV2RepositoryResult", "key");
            }
            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder maxUniqueTags(@Nullable Integer maxUniqueTags) {

            this.maxUniqueTags = maxUniqueTags;
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
              throw new MissingRequiredPropertyException("GetLocalDockerV2RepositoryResult", "packageType");
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
              throw new MissingRequiredPropertyException("GetLocalDockerV2RepositoryResult", "projectEnvironments");
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
        public Builder tagRetention(@Nullable Integer tagRetention) {

            this.tagRetention = tagRetention;
            return this;
        }
        @CustomType.Setter
        public Builder xrayIndex(@Nullable Boolean xrayIndex) {

            this.xrayIndex = xrayIndex;
            return this;
        }
        public GetLocalDockerV2RepositoryResult build() {
            final var _resultValue = new GetLocalDockerV2RepositoryResult();
            _resultValue.apiVersion = apiVersion;
            _resultValue.archiveBrowsingEnabled = archiveBrowsingEnabled;
            _resultValue.blackedOut = blackedOut;
            _resultValue.blockPushingSchema1 = blockPushingSchema1;
            _resultValue.cdnRedirect = cdnRedirect;
            _resultValue.description = description;
            _resultValue.downloadDirect = downloadDirect;
            _resultValue.excludesPattern = excludesPattern;
            _resultValue.id = id;
            _resultValue.includesPattern = includesPattern;
            _resultValue.key = key;
            _resultValue.maxUniqueTags = maxUniqueTags;
            _resultValue.notes = notes;
            _resultValue.packageType = packageType;
            _resultValue.priorityResolution = priorityResolution;
            _resultValue.projectEnvironments = projectEnvironments;
            _resultValue.projectKey = projectKey;
            _resultValue.propertySets = propertySets;
            _resultValue.repoLayoutRef = repoLayoutRef;
            _resultValue.tagRetention = tagRetention;
            _resultValue.xrayIndex = xrayIndex;
            return _resultValue;
        }
    }
}

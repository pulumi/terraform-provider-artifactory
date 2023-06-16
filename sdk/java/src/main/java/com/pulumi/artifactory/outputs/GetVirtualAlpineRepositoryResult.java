// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVirtualAlpineRepositoryResult {
    private @Nullable Boolean artifactoryRequestsCanRetrieveRemoteArtifacts;
    private @Nullable String defaultDeploymentRepo;
    private @Nullable String description;
    private @Nullable String excludesPattern;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String includesPattern;
    private String key;
    private @Nullable String notes;
    private String packageType;
    /**
     * @return (Optional) Primary keypair used to sign artifacts. Default value is empty.
     * 
     */
    private @Nullable String primaryKeypairRef;
    private List<String> projectEnvironments;
    private @Nullable String projectKey;
    private @Nullable String repoLayoutRef;
    private @Nullable List<String> repositories;
    /**
     * @return (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
     * 
     */
    private @Nullable Integer retrievalCachePeriodSeconds;

    private GetVirtualAlpineRepositoryResult() {}
    public Optional<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts() {
        return Optional.ofNullable(this.artifactoryRequestsCanRetrieveRemoteArtifacts);
    }
    public Optional<String> defaultDeploymentRepo() {
        return Optional.ofNullable(this.defaultDeploymentRepo);
    }
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
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
    public Optional<String> notes() {
        return Optional.ofNullable(this.notes);
    }
    public String packageType() {
        return this.packageType;
    }
    /**
     * @return (Optional) Primary keypair used to sign artifacts. Default value is empty.
     * 
     */
    public Optional<String> primaryKeypairRef() {
        return Optional.ofNullable(this.primaryKeypairRef);
    }
    public List<String> projectEnvironments() {
        return this.projectEnvironments;
    }
    public Optional<String> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }
    public Optional<String> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }
    public List<String> repositories() {
        return this.repositories == null ? List.of() : this.repositories;
    }
    /**
     * @return (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
     * 
     */
    public Optional<Integer> retrievalCachePeriodSeconds() {
        return Optional.ofNullable(this.retrievalCachePeriodSeconds);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVirtualAlpineRepositoryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean artifactoryRequestsCanRetrieveRemoteArtifacts;
        private @Nullable String defaultDeploymentRepo;
        private @Nullable String description;
        private @Nullable String excludesPattern;
        private String id;
        private @Nullable String includesPattern;
        private String key;
        private @Nullable String notes;
        private String packageType;
        private @Nullable String primaryKeypairRef;
        private List<String> projectEnvironments;
        private @Nullable String projectKey;
        private @Nullable String repoLayoutRef;
        private @Nullable List<String> repositories;
        private @Nullable Integer retrievalCachePeriodSeconds;
        public Builder() {}
        public Builder(GetVirtualAlpineRepositoryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.artifactoryRequestsCanRetrieveRemoteArtifacts = defaults.artifactoryRequestsCanRetrieveRemoteArtifacts;
    	      this.defaultDeploymentRepo = defaults.defaultDeploymentRepo;
    	      this.description = defaults.description;
    	      this.excludesPattern = defaults.excludesPattern;
    	      this.id = defaults.id;
    	      this.includesPattern = defaults.includesPattern;
    	      this.key = defaults.key;
    	      this.notes = defaults.notes;
    	      this.packageType = defaults.packageType;
    	      this.primaryKeypairRef = defaults.primaryKeypairRef;
    	      this.projectEnvironments = defaults.projectEnvironments;
    	      this.projectKey = defaults.projectKey;
    	      this.repoLayoutRef = defaults.repoLayoutRef;
    	      this.repositories = defaults.repositories;
    	      this.retrievalCachePeriodSeconds = defaults.retrievalCachePeriodSeconds;
        }

        @CustomType.Setter
        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(@Nullable Boolean artifactoryRequestsCanRetrieveRemoteArtifacts) {
            this.artifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            return this;
        }
        @CustomType.Setter
        public Builder defaultDeploymentRepo(@Nullable String defaultDeploymentRepo) {
            this.defaultDeploymentRepo = defaultDeploymentRepo;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder excludesPattern(@Nullable String excludesPattern) {
            this.excludesPattern = excludesPattern;
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
        public Builder notes(@Nullable String notes) {
            this.notes = notes;
            return this;
        }
        @CustomType.Setter
        public Builder packageType(String packageType) {
            this.packageType = Objects.requireNonNull(packageType);
            return this;
        }
        @CustomType.Setter
        public Builder primaryKeypairRef(@Nullable String primaryKeypairRef) {
            this.primaryKeypairRef = primaryKeypairRef;
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
        public Builder repoLayoutRef(@Nullable String repoLayoutRef) {
            this.repoLayoutRef = repoLayoutRef;
            return this;
        }
        @CustomType.Setter
        public Builder repositories(@Nullable List<String> repositories) {
            this.repositories = repositories;
            return this;
        }
        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }
        @CustomType.Setter
        public Builder retrievalCachePeriodSeconds(@Nullable Integer retrievalCachePeriodSeconds) {
            this.retrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
            return this;
        }
        public GetVirtualAlpineRepositoryResult build() {
            final var o = new GetVirtualAlpineRepositoryResult();
            o.artifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            o.defaultDeploymentRepo = defaultDeploymentRepo;
            o.description = description;
            o.excludesPattern = excludesPattern;
            o.id = id;
            o.includesPattern = includesPattern;
            o.key = key;
            o.notes = notes;
            o.packageType = packageType;
            o.primaryKeypairRef = primaryKeypairRef;
            o.projectEnvironments = projectEnvironments;
            o.projectKey = projectKey;
            o.repoLayoutRef = repoLayoutRef;
            o.repositories = repositories;
            o.retrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
            return o;
        }
    }
}

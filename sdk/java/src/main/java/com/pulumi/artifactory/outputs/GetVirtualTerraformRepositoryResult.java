// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVirtualTerraformRepositoryResult {
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
    private List<String> projectEnvironments;
    private @Nullable String projectKey;
    private @Nullable String repoLayoutRef;
    private @Nullable List<String> repositories;

    private GetVirtualTerraformRepositoryResult() {}
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVirtualTerraformRepositoryResult defaults) {
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
        private List<String> projectEnvironments;
        private @Nullable String projectKey;
        private @Nullable String repoLayoutRef;
        private @Nullable List<String> repositories;
        public Builder() {}
        public Builder(GetVirtualTerraformRepositoryResult defaults) {
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
    	      this.projectEnvironments = defaults.projectEnvironments;
    	      this.projectKey = defaults.projectKey;
    	      this.repoLayoutRef = defaults.repoLayoutRef;
    	      this.repositories = defaults.repositories;
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
        public GetVirtualTerraformRepositoryResult build() {
            final var _resultValue = new GetVirtualTerraformRepositoryResult();
            _resultValue.artifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            _resultValue.defaultDeploymentRepo = defaultDeploymentRepo;
            _resultValue.description = description;
            _resultValue.excludesPattern = excludesPattern;
            _resultValue.id = id;
            _resultValue.includesPattern = includesPattern;
            _resultValue.key = key;
            _resultValue.notes = notes;
            _resultValue.packageType = packageType;
            _resultValue.projectEnvironments = projectEnvironments;
            _resultValue.projectKey = projectKey;
            _resultValue.repoLayoutRef = repoLayoutRef;
            _resultValue.repositories = repositories;
            return _resultValue;
        }
    }
}

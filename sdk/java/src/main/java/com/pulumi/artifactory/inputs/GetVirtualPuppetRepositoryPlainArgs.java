// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVirtualPuppetRepositoryPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVirtualPuppetRepositoryPlainArgs Empty = new GetVirtualPuppetRepositoryPlainArgs();

    @Import(name="artifactoryRequestsCanRetrieveRemoteArtifacts")
    private @Nullable Boolean artifactoryRequestsCanRetrieveRemoteArtifacts;

    public Optional<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts() {
        return Optional.ofNullable(this.artifactoryRequestsCanRetrieveRemoteArtifacts);
    }

    @Import(name="defaultDeploymentRepo")
    private @Nullable String defaultDeploymentRepo;

    public Optional<String> defaultDeploymentRepo() {
        return Optional.ofNullable(this.defaultDeploymentRepo);
    }

    @Import(name="description")
    private @Nullable String description;

    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="excludesPattern")
    private @Nullable String excludesPattern;

    public Optional<String> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    @Import(name="includesPattern")
    private @Nullable String includesPattern;

    public Optional<String> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }

    @Import(name="key", required=true)
    private String key;

    public String key() {
        return this.key;
    }

    @Import(name="notes")
    private @Nullable String notes;

    public Optional<String> notes() {
        return Optional.ofNullable(this.notes);
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

    @Import(name="repoLayoutRef")
    private @Nullable String repoLayoutRef;

    public Optional<String> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    @Import(name="repositories")
    private @Nullable List<String> repositories;

    public Optional<List<String>> repositories() {
        return Optional.ofNullable(this.repositories);
    }

    private GetVirtualPuppetRepositoryPlainArgs() {}

    private GetVirtualPuppetRepositoryPlainArgs(GetVirtualPuppetRepositoryPlainArgs $) {
        this.artifactoryRequestsCanRetrieveRemoteArtifacts = $.artifactoryRequestsCanRetrieveRemoteArtifacts;
        this.defaultDeploymentRepo = $.defaultDeploymentRepo;
        this.description = $.description;
        this.excludesPattern = $.excludesPattern;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.notes = $.notes;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.repoLayoutRef = $.repoLayoutRef;
        this.repositories = $.repositories;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVirtualPuppetRepositoryPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVirtualPuppetRepositoryPlainArgs $;

        public Builder() {
            $ = new GetVirtualPuppetRepositoryPlainArgs();
        }

        public Builder(GetVirtualPuppetRepositoryPlainArgs defaults) {
            $ = new GetVirtualPuppetRepositoryPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(@Nullable Boolean artifactoryRequestsCanRetrieveRemoteArtifacts) {
            $.artifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            return this;
        }

        public Builder defaultDeploymentRepo(@Nullable String defaultDeploymentRepo) {
            $.defaultDeploymentRepo = defaultDeploymentRepo;
            return this;
        }

        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        public Builder excludesPattern(@Nullable String excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        public Builder includesPattern(@Nullable String includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        public Builder key(String key) {
            $.key = key;
            return this;
        }

        public Builder notes(@Nullable String notes) {
            $.notes = notes;
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

        public Builder repoLayoutRef(@Nullable String repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        public Builder repositories(@Nullable List<String> repositories) {
            $.repositories = repositories;
            return this;
        }

        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }

        public GetVirtualPuppetRepositoryPlainArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}

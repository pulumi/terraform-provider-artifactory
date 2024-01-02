// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVirtualSbtRepositoryPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVirtualSbtRepositoryPlainArgs Empty = new GetVirtualSbtRepositoryPlainArgs();

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

    @Import(name="forceMavenAuthentication")
    private @Nullable Boolean forceMavenAuthentication;

    public Optional<Boolean> forceMavenAuthentication() {
        return Optional.ofNullable(this.forceMavenAuthentication);
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
     * (Optional) The keypair used to sign artifacts.
     * 
     */
    @Import(name="keyPair")
    private @Nullable String keyPair;

    /**
     * @return (Optional) The keypair used to sign artifacts.
     * 
     */
    public Optional<String> keyPair() {
        return Optional.ofNullable(this.keyPair);
    }

    @Import(name="notes")
    private @Nullable String notes;

    public Optional<String> notes() {
        return Optional.ofNullable(this.notes);
    }

    /**
     * (Optional)
     * - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
     * - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
     * - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
     * 
     */
    @Import(name="pomRepositoryReferencesCleanupPolicy")
    private @Nullable String pomRepositoryReferencesCleanupPolicy;

    /**
     * @return (Optional)
     * - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
     * - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
     * - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
     * 
     */
    public Optional<String> pomRepositoryReferencesCleanupPolicy() {
        return Optional.ofNullable(this.pomRepositoryReferencesCleanupPolicy);
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

    private GetVirtualSbtRepositoryPlainArgs() {}

    private GetVirtualSbtRepositoryPlainArgs(GetVirtualSbtRepositoryPlainArgs $) {
        this.artifactoryRequestsCanRetrieveRemoteArtifacts = $.artifactoryRequestsCanRetrieveRemoteArtifacts;
        this.defaultDeploymentRepo = $.defaultDeploymentRepo;
        this.description = $.description;
        this.excludesPattern = $.excludesPattern;
        this.forceMavenAuthentication = $.forceMavenAuthentication;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.keyPair = $.keyPair;
        this.notes = $.notes;
        this.pomRepositoryReferencesCleanupPolicy = $.pomRepositoryReferencesCleanupPolicy;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.repoLayoutRef = $.repoLayoutRef;
        this.repositories = $.repositories;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVirtualSbtRepositoryPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVirtualSbtRepositoryPlainArgs $;

        public Builder() {
            $ = new GetVirtualSbtRepositoryPlainArgs();
        }

        public Builder(GetVirtualSbtRepositoryPlainArgs defaults) {
            $ = new GetVirtualSbtRepositoryPlainArgs(Objects.requireNonNull(defaults));
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

        public Builder forceMavenAuthentication(@Nullable Boolean forceMavenAuthentication) {
            $.forceMavenAuthentication = forceMavenAuthentication;
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
         * @param keyPair (Optional) The keypair used to sign artifacts.
         * 
         * @return builder
         * 
         */
        public Builder keyPair(@Nullable String keyPair) {
            $.keyPair = keyPair;
            return this;
        }

        public Builder notes(@Nullable String notes) {
            $.notes = notes;
            return this;
        }

        /**
         * @param pomRepositoryReferencesCleanupPolicy (Optional)
         * - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
         * - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
         * - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
         * 
         * @return builder
         * 
         */
        public Builder pomRepositoryReferencesCleanupPolicy(@Nullable String pomRepositoryReferencesCleanupPolicy) {
            $.pomRepositoryReferencesCleanupPolicy = pomRepositoryReferencesCleanupPolicy;
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

        public GetVirtualSbtRepositoryPlainArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("GetVirtualSbtRepositoryPlainArgs", "key");
            }
            return $;
        }
    }

}

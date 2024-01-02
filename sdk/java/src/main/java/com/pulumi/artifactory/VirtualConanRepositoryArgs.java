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


public final class VirtualConanRepositoryArgs extends com.pulumi.resources.ResourceArgs {

    public static final VirtualConanRepositoryArgs Empty = new VirtualConanRepositoryArgs();

    /**
     * Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    @Import(name="artifactoryRequestsCanRetrieveRemoteArtifacts")
    private @Nullable Output<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts;

    /**
     * @return Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    public Optional<Output<Boolean>> artifactoryRequestsCanRetrieveRemoteArtifacts() {
        return Optional.ofNullable(this.artifactoryRequestsCanRetrieveRemoteArtifacts);
    }

    /**
     * Default repository to deploy artifacts.
     * 
     */
    @Import(name="defaultDeploymentRepo")
    private @Nullable Output<String> defaultDeploymentRepo;

    /**
     * @return Default repository to deploy artifacts.
     * 
     */
    public Optional<Output<String>> defaultDeploymentRepo() {
        return Optional.ofNullable(this.defaultDeploymentRepo);
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
     * Force basic authentication credentials in order to use this repository. Default value is `false`.
     * 
     */
    @Import(name="forceConanAuthentication")
    private @Nullable Output<Boolean> forceConanAuthentication;

    /**
     * @return Force basic authentication credentials in order to use this repository. Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> forceConanAuthentication() {
        return Optional.ofNullable(this.forceConanAuthentication);
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
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     * 
     */
    public Output<String> key() {
        return this.key;
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
     * Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    @Import(name="projectKey")
    private @Nullable Output<String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    public Optional<Output<String>> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }

    /**
     * Repository layout key for the virtual repository
     * 
     */
    @Import(name="repoLayoutRef")
    private @Nullable Output<String> repoLayoutRef;

    /**
     * @return Repository layout key for the virtual repository
     * 
     */
    public Optional<Output<String>> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    /**
     * The effective list of actual repositories included in this virtual repository.
     * 
     */
    @Import(name="repositories")
    private @Nullable Output<List<String>> repositories;

    /**
     * @return The effective list of actual repositories included in this virtual repository.
     * 
     */
    public Optional<Output<List<String>>> repositories() {
        return Optional.ofNullable(this.repositories);
    }

    /**
     * This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
     * 
     */
    @Import(name="retrievalCachePeriodSeconds")
    private @Nullable Output<Integer> retrievalCachePeriodSeconds;

    /**
     * @return This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
     * 
     */
    public Optional<Output<Integer>> retrievalCachePeriodSeconds() {
        return Optional.ofNullable(this.retrievalCachePeriodSeconds);
    }

    private VirtualConanRepositoryArgs() {}

    private VirtualConanRepositoryArgs(VirtualConanRepositoryArgs $) {
        this.artifactoryRequestsCanRetrieveRemoteArtifacts = $.artifactoryRequestsCanRetrieveRemoteArtifacts;
        this.defaultDeploymentRepo = $.defaultDeploymentRepo;
        this.description = $.description;
        this.excludesPattern = $.excludesPattern;
        this.forceConanAuthentication = $.forceConanAuthentication;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.notes = $.notes;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.repoLayoutRef = $.repoLayoutRef;
        this.repositories = $.repositories;
        this.retrievalCachePeriodSeconds = $.retrievalCachePeriodSeconds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualConanRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualConanRepositoryArgs $;

        public Builder() {
            $ = new VirtualConanRepositoryArgs();
        }

        public Builder(VirtualConanRepositoryArgs defaults) {
            $ = new VirtualConanRepositoryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param artifactoryRequestsCanRetrieveRemoteArtifacts Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
         * another Artifactory instance.
         * 
         * @return builder
         * 
         */
        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(@Nullable Output<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts) {
            $.artifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            return this;
        }

        /**
         * @param artifactoryRequestsCanRetrieveRemoteArtifacts Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
         * another Artifactory instance.
         * 
         * @return builder
         * 
         */
        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(Boolean artifactoryRequestsCanRetrieveRemoteArtifacts) {
            return artifactoryRequestsCanRetrieveRemoteArtifacts(Output.of(artifactoryRequestsCanRetrieveRemoteArtifacts));
        }

        /**
         * @param defaultDeploymentRepo Default repository to deploy artifacts.
         * 
         * @return builder
         * 
         */
        public Builder defaultDeploymentRepo(@Nullable Output<String> defaultDeploymentRepo) {
            $.defaultDeploymentRepo = defaultDeploymentRepo;
            return this;
        }

        /**
         * @param defaultDeploymentRepo Default repository to deploy artifacts.
         * 
         * @return builder
         * 
         */
        public Builder defaultDeploymentRepo(String defaultDeploymentRepo) {
            return defaultDeploymentRepo(Output.of(defaultDeploymentRepo));
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
         * @param forceConanAuthentication Force basic authentication credentials in order to use this repository. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder forceConanAuthentication(@Nullable Output<Boolean> forceConanAuthentication) {
            $.forceConanAuthentication = forceConanAuthentication;
            return this;
        }

        /**
         * @param forceConanAuthentication Force basic authentication credentials in order to use this repository. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder forceConanAuthentication(Boolean forceConanAuthentication) {
            return forceConanAuthentication(Output.of(forceConanAuthentication));
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
         * @param key A mandatory identifier for the repository that must be unique. It cannot begin with a number or
         * contain spaces or special characters.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key A mandatory identifier for the repository that must be unique. It cannot begin with a number or
         * contain spaces or special characters.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
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
         * @param projectKey Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
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
         * @param projectKey Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
         * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
         * 
         * @return builder
         * 
         */
        public Builder projectKey(String projectKey) {
            return projectKey(Output.of(projectKey));
        }

        /**
         * @param repoLayoutRef Repository layout key for the virtual repository
         * 
         * @return builder
         * 
         */
        public Builder repoLayoutRef(@Nullable Output<String> repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        /**
         * @param repoLayoutRef Repository layout key for the virtual repository
         * 
         * @return builder
         * 
         */
        public Builder repoLayoutRef(String repoLayoutRef) {
            return repoLayoutRef(Output.of(repoLayoutRef));
        }

        /**
         * @param repositories The effective list of actual repositories included in this virtual repository.
         * 
         * @return builder
         * 
         */
        public Builder repositories(@Nullable Output<List<String>> repositories) {
            $.repositories = repositories;
            return this;
        }

        /**
         * @param repositories The effective list of actual repositories included in this virtual repository.
         * 
         * @return builder
         * 
         */
        public Builder repositories(List<String> repositories) {
            return repositories(Output.of(repositories));
        }

        /**
         * @param repositories The effective list of actual repositories included in this virtual repository.
         * 
         * @return builder
         * 
         */
        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }

        /**
         * @param retrievalCachePeriodSeconds This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
         * 
         * @return builder
         * 
         */
        public Builder retrievalCachePeriodSeconds(@Nullable Output<Integer> retrievalCachePeriodSeconds) {
            $.retrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
            return this;
        }

        /**
         * @param retrievalCachePeriodSeconds This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
         * 
         * @return builder
         * 
         */
        public Builder retrievalCachePeriodSeconds(Integer retrievalCachePeriodSeconds) {
            return retrievalCachePeriodSeconds(Output.of(retrievalCachePeriodSeconds));
        }

        public VirtualConanRepositoryArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("VirtualConanRepositoryArgs", "key");
            }
            return $;
        }
    }

}

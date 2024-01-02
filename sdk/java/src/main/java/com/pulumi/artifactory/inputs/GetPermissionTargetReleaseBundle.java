// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.GetPermissionTargetReleaseBundleActions;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPermissionTargetReleaseBundle extends com.pulumi.resources.InvokeArgs {

    public static final GetPermissionTargetReleaseBundle Empty = new GetPermissionTargetReleaseBundle();

    @Import(name="actions")
    private @Nullable GetPermissionTargetReleaseBundleActions actions;

    public Optional<GetPermissionTargetReleaseBundleActions> actions() {
        return Optional.ofNullable(this.actions);
    }

    /**
     * Pattern of artifacts to exclude.
     * 
     */
    @Import(name="excludesPatterns")
    private @Nullable List<String> excludesPatterns;

    /**
     * @return Pattern of artifacts to exclude.
     * 
     */
    public Optional<List<String>> excludesPatterns() {
        return Optional.ofNullable(this.excludesPatterns);
    }

    /**
     * Pattern of artifacts to include.
     * 
     */
    @Import(name="includesPatterns")
    private @Nullable List<String> includesPatterns;

    /**
     * @return Pattern of artifacts to include.
     * 
     */
    public Optional<List<String>> includesPatterns() {
        return Optional.ofNullable(this.includesPatterns);
    }

    /**
     * List of repositories this permission target is applicable for. You can specify the
     * name `ANY` in the repositories section in order to apply to all repositories, `ANY REMOTE` for all remote
     * repositories and `ANY LOCAL` for all local repositories. The default value will be `[]` if nothing is specified.
     * 
     */
    @Import(name="repositories", required=true)
    private List<String> repositories;

    /**
     * @return List of repositories this permission target is applicable for. You can specify the
     * name `ANY` in the repositories section in order to apply to all repositories, `ANY REMOTE` for all remote
     * repositories and `ANY LOCAL` for all local repositories. The default value will be `[]` if nothing is specified.
     * 
     */
    public List<String> repositories() {
        return this.repositories;
    }

    private GetPermissionTargetReleaseBundle() {}

    private GetPermissionTargetReleaseBundle(GetPermissionTargetReleaseBundle $) {
        this.actions = $.actions;
        this.excludesPatterns = $.excludesPatterns;
        this.includesPatterns = $.includesPatterns;
        this.repositories = $.repositories;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPermissionTargetReleaseBundle defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPermissionTargetReleaseBundle $;

        public Builder() {
            $ = new GetPermissionTargetReleaseBundle();
        }

        public Builder(GetPermissionTargetReleaseBundle defaults) {
            $ = new GetPermissionTargetReleaseBundle(Objects.requireNonNull(defaults));
        }

        public Builder actions(@Nullable GetPermissionTargetReleaseBundleActions actions) {
            $.actions = actions;
            return this;
        }

        /**
         * @param excludesPatterns Pattern of artifacts to exclude.
         * 
         * @return builder
         * 
         */
        public Builder excludesPatterns(@Nullable List<String> excludesPatterns) {
            $.excludesPatterns = excludesPatterns;
            return this;
        }

        /**
         * @param excludesPatterns Pattern of artifacts to exclude.
         * 
         * @return builder
         * 
         */
        public Builder excludesPatterns(String... excludesPatterns) {
            return excludesPatterns(List.of(excludesPatterns));
        }

        /**
         * @param includesPatterns Pattern of artifacts to include.
         * 
         * @return builder
         * 
         */
        public Builder includesPatterns(@Nullable List<String> includesPatterns) {
            $.includesPatterns = includesPatterns;
            return this;
        }

        /**
         * @param includesPatterns Pattern of artifacts to include.
         * 
         * @return builder
         * 
         */
        public Builder includesPatterns(String... includesPatterns) {
            return includesPatterns(List.of(includesPatterns));
        }

        /**
         * @param repositories List of repositories this permission target is applicable for. You can specify the
         * name `ANY` in the repositories section in order to apply to all repositories, `ANY REMOTE` for all remote
         * repositories and `ANY LOCAL` for all local repositories. The default value will be `[]` if nothing is specified.
         * 
         * @return builder
         * 
         */
        public Builder repositories(List<String> repositories) {
            $.repositories = repositories;
            return this;
        }

        /**
         * @param repositories List of repositories this permission target is applicable for. You can specify the
         * name `ANY` in the repositories section in order to apply to all repositories, `ANY REMOTE` for all remote
         * repositories and `ANY LOCAL` for all local repositories. The default value will be `[]` if nothing is specified.
         * 
         * @return builder
         * 
         */
        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }

        public GetPermissionTargetReleaseBundle build() {
            if ($.repositories == null) {
                throw new MissingRequiredPropertyException("GetPermissionTargetReleaseBundle", "repositories");
            }
            return $;
        }
    }

}

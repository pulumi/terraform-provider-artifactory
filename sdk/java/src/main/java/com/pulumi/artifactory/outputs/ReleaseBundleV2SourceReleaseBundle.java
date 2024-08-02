// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ReleaseBundleV2SourceReleaseBundle {
    /**
     * @return The name of the release bundle.
     * 
     */
    private String name;
    /**
     * @return Project key of the release bundle.
     * 
     */
    private @Nullable String projectKey;
    /**
     * @return The key of the release bundle repository.
     * 
     */
    private @Nullable String repositoryKey;
    /**
     * @return The version of the release bundle.
     * 
     */
    private String version;

    private ReleaseBundleV2SourceReleaseBundle() {}
    /**
     * @return The name of the release bundle.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Project key of the release bundle.
     * 
     */
    public Optional<String> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }
    /**
     * @return The key of the release bundle repository.
     * 
     */
    public Optional<String> repositoryKey() {
        return Optional.ofNullable(this.repositoryKey);
    }
    /**
     * @return The version of the release bundle.
     * 
     */
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ReleaseBundleV2SourceReleaseBundle defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private @Nullable String projectKey;
        private @Nullable String repositoryKey;
        private String version;
        public Builder() {}
        public Builder(ReleaseBundleV2SourceReleaseBundle defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.projectKey = defaults.projectKey;
    	      this.repositoryKey = defaults.repositoryKey;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("ReleaseBundleV2SourceReleaseBundle", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder projectKey(@Nullable String projectKey) {

            this.projectKey = projectKey;
            return this;
        }
        @CustomType.Setter
        public Builder repositoryKey(@Nullable String repositoryKey) {

            this.repositoryKey = repositoryKey;
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            if (version == null) {
              throw new MissingRequiredPropertyException("ReleaseBundleV2SourceReleaseBundle", "version");
            }
            this.version = version;
            return this;
        }
        public ReleaseBundleV2SourceReleaseBundle build() {
            final var _resultValue = new ReleaseBundleV2SourceReleaseBundle();
            _resultValue.name = name;
            _resultValue.projectKey = projectKey;
            _resultValue.repositoryKey = repositoryKey;
            _resultValue.version = version;
            return _resultValue;
        }
    }
}

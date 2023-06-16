// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.artifactory.outputs.PermissionTargetReleaseBundleActions;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PermissionTargetReleaseBundle {
    private @Nullable PermissionTargetReleaseBundleActions actions;
    private @Nullable List<String> excludesPatterns;
    private @Nullable List<String> includesPatterns;
    private List<String> repositories;

    private PermissionTargetReleaseBundle() {}
    public Optional<PermissionTargetReleaseBundleActions> actions() {
        return Optional.ofNullable(this.actions);
    }
    public List<String> excludesPatterns() {
        return this.excludesPatterns == null ? List.of() : this.excludesPatterns;
    }
    public List<String> includesPatterns() {
        return this.includesPatterns == null ? List.of() : this.includesPatterns;
    }
    public List<String> repositories() {
        return this.repositories;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PermissionTargetReleaseBundle defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable PermissionTargetReleaseBundleActions actions;
        private @Nullable List<String> excludesPatterns;
        private @Nullable List<String> includesPatterns;
        private List<String> repositories;
        public Builder() {}
        public Builder(PermissionTargetReleaseBundle defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actions = defaults.actions;
    	      this.excludesPatterns = defaults.excludesPatterns;
    	      this.includesPatterns = defaults.includesPatterns;
    	      this.repositories = defaults.repositories;
        }

        @CustomType.Setter
        public Builder actions(@Nullable PermissionTargetReleaseBundleActions actions) {
            this.actions = actions;
            return this;
        }
        @CustomType.Setter
        public Builder excludesPatterns(@Nullable List<String> excludesPatterns) {
            this.excludesPatterns = excludesPatterns;
            return this;
        }
        public Builder excludesPatterns(String... excludesPatterns) {
            return excludesPatterns(List.of(excludesPatterns));
        }
        @CustomType.Setter
        public Builder includesPatterns(@Nullable List<String> includesPatterns) {
            this.includesPatterns = includesPatterns;
            return this;
        }
        public Builder includesPatterns(String... includesPatterns) {
            return includesPatterns(List.of(includesPatterns));
        }
        @CustomType.Setter
        public Builder repositories(List<String> repositories) {
            this.repositories = Objects.requireNonNull(repositories);
            return this;
        }
        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }
        public PermissionTargetReleaseBundle build() {
            final var o = new PermissionTargetReleaseBundle();
            o.actions = actions;
            o.excludesPatterns = excludesPatterns;
            o.includesPatterns = includesPatterns;
            o.repositories = repositories;
            return o;
        }
    }
}

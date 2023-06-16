// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ArtifactoryReleaseBundleWebhookCriteria {
    private Boolean anyReleaseBundle;
    private @Nullable List<String> excludePatterns;
    private @Nullable List<String> includePatterns;
    private List<String> registeredReleaseBundleNames;

    private ArtifactoryReleaseBundleWebhookCriteria() {}
    public Boolean anyReleaseBundle() {
        return this.anyReleaseBundle;
    }
    public List<String> excludePatterns() {
        return this.excludePatterns == null ? List.of() : this.excludePatterns;
    }
    public List<String> includePatterns() {
        return this.includePatterns == null ? List.of() : this.includePatterns;
    }
    public List<String> registeredReleaseBundleNames() {
        return this.registeredReleaseBundleNames;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ArtifactoryReleaseBundleWebhookCriteria defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean anyReleaseBundle;
        private @Nullable List<String> excludePatterns;
        private @Nullable List<String> includePatterns;
        private List<String> registeredReleaseBundleNames;
        public Builder() {}
        public Builder(ArtifactoryReleaseBundleWebhookCriteria defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.anyReleaseBundle = defaults.anyReleaseBundle;
    	      this.excludePatterns = defaults.excludePatterns;
    	      this.includePatterns = defaults.includePatterns;
    	      this.registeredReleaseBundleNames = defaults.registeredReleaseBundleNames;
        }

        @CustomType.Setter
        public Builder anyReleaseBundle(Boolean anyReleaseBundle) {
            this.anyReleaseBundle = Objects.requireNonNull(anyReleaseBundle);
            return this;
        }
        @CustomType.Setter
        public Builder excludePatterns(@Nullable List<String> excludePatterns) {
            this.excludePatterns = excludePatterns;
            return this;
        }
        public Builder excludePatterns(String... excludePatterns) {
            return excludePatterns(List.of(excludePatterns));
        }
        @CustomType.Setter
        public Builder includePatterns(@Nullable List<String> includePatterns) {
            this.includePatterns = includePatterns;
            return this;
        }
        public Builder includePatterns(String... includePatterns) {
            return includePatterns(List.of(includePatterns));
        }
        @CustomType.Setter
        public Builder registeredReleaseBundleNames(List<String> registeredReleaseBundleNames) {
            this.registeredReleaseBundleNames = Objects.requireNonNull(registeredReleaseBundleNames);
            return this;
        }
        public Builder registeredReleaseBundleNames(String... registeredReleaseBundleNames) {
            return registeredReleaseBundleNames(List.of(registeredReleaseBundleNames));
        }
        public ArtifactoryReleaseBundleWebhookCriteria build() {
            final var o = new ArtifactoryReleaseBundleWebhookCriteria();
            o.anyReleaseBundle = anyReleaseBundle;
            o.excludePatterns = excludePatterns;
            o.includePatterns = includePatterns;
            o.registeredReleaseBundleNames = registeredReleaseBundleNames;
            return o;
        }
    }
}

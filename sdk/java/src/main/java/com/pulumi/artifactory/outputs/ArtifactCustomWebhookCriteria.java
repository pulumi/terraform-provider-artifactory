// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ArtifactCustomWebhookCriteria {
    /**
     * @return Trigger on any local repo.
     * 
     */
    private Boolean anyLocal;
    /**
     * @return Trigger on any remote repo.
     * 
     */
    private Boolean anyRemote;
    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
     * 
     */
    private @Nullable List<String> excludePatterns;
    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
     * 
     */
    private @Nullable List<String> includePatterns;
    /**
     * @return Trigger on this list of repo keys.
     * 
     */
    private List<String> repoKeys;

    private ArtifactCustomWebhookCriteria() {}
    /**
     * @return Trigger on any local repo.
     * 
     */
    public Boolean anyLocal() {
        return this.anyLocal;
    }
    /**
     * @return Trigger on any remote repo.
     * 
     */
    public Boolean anyRemote() {
        return this.anyRemote;
    }
    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
     * 
     */
    public List<String> excludePatterns() {
        return this.excludePatterns == null ? List.of() : this.excludePatterns;
    }
    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
     * 
     */
    public List<String> includePatterns() {
        return this.includePatterns == null ? List.of() : this.includePatterns;
    }
    /**
     * @return Trigger on this list of repo keys.
     * 
     */
    public List<String> repoKeys() {
        return this.repoKeys;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ArtifactCustomWebhookCriteria defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean anyLocal;
        private Boolean anyRemote;
        private @Nullable List<String> excludePatterns;
        private @Nullable List<String> includePatterns;
        private List<String> repoKeys;
        public Builder() {}
        public Builder(ArtifactCustomWebhookCriteria defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.anyLocal = defaults.anyLocal;
    	      this.anyRemote = defaults.anyRemote;
    	      this.excludePatterns = defaults.excludePatterns;
    	      this.includePatterns = defaults.includePatterns;
    	      this.repoKeys = defaults.repoKeys;
        }

        @CustomType.Setter
        public Builder anyLocal(Boolean anyLocal) {
            if (anyLocal == null) {
              throw new MissingRequiredPropertyException("ArtifactCustomWebhookCriteria", "anyLocal");
            }
            this.anyLocal = anyLocal;
            return this;
        }
        @CustomType.Setter
        public Builder anyRemote(Boolean anyRemote) {
            if (anyRemote == null) {
              throw new MissingRequiredPropertyException("ArtifactCustomWebhookCriteria", "anyRemote");
            }
            this.anyRemote = anyRemote;
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
        public Builder repoKeys(List<String> repoKeys) {
            if (repoKeys == null) {
              throw new MissingRequiredPropertyException("ArtifactCustomWebhookCriteria", "repoKeys");
            }
            this.repoKeys = repoKeys;
            return this;
        }
        public Builder repoKeys(String... repoKeys) {
            return repoKeys(List.of(repoKeys));
        }
        public ArtifactCustomWebhookCriteria build() {
            final var _resultValue = new ArtifactCustomWebhookCriteria();
            _resultValue.anyLocal = anyLocal;
            _resultValue.anyRemote = anyRemote;
            _resultValue.excludePatterns = excludePatterns;
            _resultValue.includePatterns = includePatterns;
            _resultValue.repoKeys = repoKeys;
            return _resultValue;
        }
    }
}

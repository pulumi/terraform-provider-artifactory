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
public final class ArtifactPropertyWebhookCriteria {
    /**
     * @return Trigger on any local repo.
     * 
     */
    private final Boolean anyLocal;
    /**
     * @return Trigger on any remote repo.
     * 
     */
    private final Boolean anyRemote;
    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, *\*, ?).\nFor example: &#34;org/apache/**&#34;.
     * 
     */
    private final @Nullable List<String> excludePatterns;
    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, *\*, ?).\nFor example: &#34;org/apache/**&#34;.
     * 
     */
    private final @Nullable List<String> includePatterns;
    /**
     * @return Trigger on this list of repo keys.
     * 
     */
    private final List<String> repoKeys;

    @CustomType.Constructor
    private ArtifactPropertyWebhookCriteria(
        @CustomType.Parameter("anyLocal") Boolean anyLocal,
        @CustomType.Parameter("anyRemote") Boolean anyRemote,
        @CustomType.Parameter("excludePatterns") @Nullable List<String> excludePatterns,
        @CustomType.Parameter("includePatterns") @Nullable List<String> includePatterns,
        @CustomType.Parameter("repoKeys") List<String> repoKeys) {
        this.anyLocal = anyLocal;
        this.anyRemote = anyRemote;
        this.excludePatterns = excludePatterns;
        this.includePatterns = includePatterns;
        this.repoKeys = repoKeys;
    }

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
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, *\*, ?).\nFor example: &#34;org/apache/**&#34;.
     * 
     */
    public List<String> excludePatterns() {
        return this.excludePatterns == null ? List.of() : this.excludePatterns;
    }
    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, *\*, ?).\nFor example: &#34;org/apache/**&#34;.
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

    public static Builder builder(ArtifactPropertyWebhookCriteria defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Boolean anyLocal;
        private Boolean anyRemote;
        private @Nullable List<String> excludePatterns;
        private @Nullable List<String> includePatterns;
        private List<String> repoKeys;

        public Builder() {
    	      // Empty
        }

        public Builder(ArtifactPropertyWebhookCriteria defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.anyLocal = defaults.anyLocal;
    	      this.anyRemote = defaults.anyRemote;
    	      this.excludePatterns = defaults.excludePatterns;
    	      this.includePatterns = defaults.includePatterns;
    	      this.repoKeys = defaults.repoKeys;
        }

        public Builder anyLocal(Boolean anyLocal) {
            this.anyLocal = Objects.requireNonNull(anyLocal);
            return this;
        }
        public Builder anyRemote(Boolean anyRemote) {
            this.anyRemote = Objects.requireNonNull(anyRemote);
            return this;
        }
        public Builder excludePatterns(@Nullable List<String> excludePatterns) {
            this.excludePatterns = excludePatterns;
            return this;
        }
        public Builder excludePatterns(String... excludePatterns) {
            return excludePatterns(List.of(excludePatterns));
        }
        public Builder includePatterns(@Nullable List<String> includePatterns) {
            this.includePatterns = includePatterns;
            return this;
        }
        public Builder includePatterns(String... includePatterns) {
            return includePatterns(List.of(includePatterns));
        }
        public Builder repoKeys(List<String> repoKeys) {
            this.repoKeys = Objects.requireNonNull(repoKeys);
            return this;
        }
        public Builder repoKeys(String... repoKeys) {
            return repoKeys(List.of(repoKeys));
        }        public ArtifactPropertyWebhookCriteria build() {
            return new ArtifactPropertyWebhookCriteria(anyLocal, anyRemote, excludePatterns, includePatterns, repoKeys);
        }
    }
}

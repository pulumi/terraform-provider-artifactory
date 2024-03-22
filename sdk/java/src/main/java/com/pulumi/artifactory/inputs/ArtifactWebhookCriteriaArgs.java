// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ArtifactWebhookCriteriaArgs extends com.pulumi.resources.ResourceArgs {

    public static final ArtifactWebhookCriteriaArgs Empty = new ArtifactWebhookCriteriaArgs();

    /**
     * Trigger on any federated repo.
     * 
     */
    @Import(name="anyFederated", required=true)
    private Output<Boolean> anyFederated;

    /**
     * @return Trigger on any federated repo.
     * 
     */
    public Output<Boolean> anyFederated() {
        return this.anyFederated;
    }

    /**
     * Trigger on any local repo.
     * 
     */
    @Import(name="anyLocal", required=true)
    private Output<Boolean> anyLocal;

    /**
     * @return Trigger on any local repo.
     * 
     */
    public Output<Boolean> anyLocal() {
        return this.anyLocal;
    }

    /**
     * Trigger on any remote repo.
     * 
     */
    @Import(name="anyRemote", required=true)
    private Output<Boolean> anyRemote;

    /**
     * @return Trigger on any remote repo.
     * 
     */
    public Output<Boolean> anyRemote() {
        return this.anyRemote;
    }

    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
     * 
     */
    @Import(name="excludePatterns")
    private @Nullable Output<List<String>> excludePatterns;

    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
     * 
     */
    public Optional<Output<List<String>>> excludePatterns() {
        return Optional.ofNullable(this.excludePatterns);
    }

    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
     * 
     */
    @Import(name="includePatterns")
    private @Nullable Output<List<String>> includePatterns;

    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
     * 
     */
    public Optional<Output<List<String>>> includePatterns() {
        return Optional.ofNullable(this.includePatterns);
    }

    /**
     * Trigger on this list of repo keys.
     * 
     */
    @Import(name="repoKeys", required=true)
    private Output<List<String>> repoKeys;

    /**
     * @return Trigger on this list of repo keys.
     * 
     */
    public Output<List<String>> repoKeys() {
        return this.repoKeys;
    }

    private ArtifactWebhookCriteriaArgs() {}

    private ArtifactWebhookCriteriaArgs(ArtifactWebhookCriteriaArgs $) {
        this.anyFederated = $.anyFederated;
        this.anyLocal = $.anyLocal;
        this.anyRemote = $.anyRemote;
        this.excludePatterns = $.excludePatterns;
        this.includePatterns = $.includePatterns;
        this.repoKeys = $.repoKeys;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ArtifactWebhookCriteriaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ArtifactWebhookCriteriaArgs $;

        public Builder() {
            $ = new ArtifactWebhookCriteriaArgs();
        }

        public Builder(ArtifactWebhookCriteriaArgs defaults) {
            $ = new ArtifactWebhookCriteriaArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param anyFederated Trigger on any federated repo.
         * 
         * @return builder
         * 
         */
        public Builder anyFederated(Output<Boolean> anyFederated) {
            $.anyFederated = anyFederated;
            return this;
        }

        /**
         * @param anyFederated Trigger on any federated repo.
         * 
         * @return builder
         * 
         */
        public Builder anyFederated(Boolean anyFederated) {
            return anyFederated(Output.of(anyFederated));
        }

        /**
         * @param anyLocal Trigger on any local repo.
         * 
         * @return builder
         * 
         */
        public Builder anyLocal(Output<Boolean> anyLocal) {
            $.anyLocal = anyLocal;
            return this;
        }

        /**
         * @param anyLocal Trigger on any local repo.
         * 
         * @return builder
         * 
         */
        public Builder anyLocal(Boolean anyLocal) {
            return anyLocal(Output.of(anyLocal));
        }

        /**
         * @param anyRemote Trigger on any remote repo.
         * 
         * @return builder
         * 
         */
        public Builder anyRemote(Output<Boolean> anyRemote) {
            $.anyRemote = anyRemote;
            return this;
        }

        /**
         * @param anyRemote Trigger on any remote repo.
         * 
         * @return builder
         * 
         */
        public Builder anyRemote(Boolean anyRemote) {
            return anyRemote(Output.of(anyRemote));
        }

        /**
         * @param excludePatterns Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
         * 
         * @return builder
         * 
         */
        public Builder excludePatterns(@Nullable Output<List<String>> excludePatterns) {
            $.excludePatterns = excludePatterns;
            return this;
        }

        /**
         * @param excludePatterns Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
         * 
         * @return builder
         * 
         */
        public Builder excludePatterns(List<String> excludePatterns) {
            return excludePatterns(Output.of(excludePatterns));
        }

        /**
         * @param excludePatterns Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
         * 
         * @return builder
         * 
         */
        public Builder excludePatterns(String... excludePatterns) {
            return excludePatterns(List.of(excludePatterns));
        }

        /**
         * @param includePatterns Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
         * 
         * @return builder
         * 
         */
        public Builder includePatterns(@Nullable Output<List<String>> includePatterns) {
            $.includePatterns = includePatterns;
            return this;
        }

        /**
         * @param includePatterns Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
         * 
         * @return builder
         * 
         */
        public Builder includePatterns(List<String> includePatterns) {
            return includePatterns(Output.of(includePatterns));
        }

        /**
         * @param includePatterns Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
         * 
         * @return builder
         * 
         */
        public Builder includePatterns(String... includePatterns) {
            return includePatterns(List.of(includePatterns));
        }

        /**
         * @param repoKeys Trigger on this list of repo keys.
         * 
         * @return builder
         * 
         */
        public Builder repoKeys(Output<List<String>> repoKeys) {
            $.repoKeys = repoKeys;
            return this;
        }

        /**
         * @param repoKeys Trigger on this list of repo keys.
         * 
         * @return builder
         * 
         */
        public Builder repoKeys(List<String> repoKeys) {
            return repoKeys(Output.of(repoKeys));
        }

        /**
         * @param repoKeys Trigger on this list of repo keys.
         * 
         * @return builder
         * 
         */
        public Builder repoKeys(String... repoKeys) {
            return repoKeys(List.of(repoKeys));
        }

        public ArtifactWebhookCriteriaArgs build() {
            if ($.anyFederated == null) {
                throw new MissingRequiredPropertyException("ArtifactWebhookCriteriaArgs", "anyFederated");
            }
            if ($.anyLocal == null) {
                throw new MissingRequiredPropertyException("ArtifactWebhookCriteriaArgs", "anyLocal");
            }
            if ($.anyRemote == null) {
                throw new MissingRequiredPropertyException("ArtifactWebhookCriteriaArgs", "anyRemote");
            }
            if ($.repoKeys == null) {
                throw new MissingRequiredPropertyException("ArtifactWebhookCriteriaArgs", "repoKeys");
            }
            return $;
        }
    }

}

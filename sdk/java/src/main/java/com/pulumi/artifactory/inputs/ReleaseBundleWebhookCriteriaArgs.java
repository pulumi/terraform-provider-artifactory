// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReleaseBundleWebhookCriteriaArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReleaseBundleWebhookCriteriaArgs Empty = new ReleaseBundleWebhookCriteriaArgs();

    @Import(name="anyReleaseBundle", required=true)
    private Output<Boolean> anyReleaseBundle;

    public Output<Boolean> anyReleaseBundle() {
        return this.anyReleaseBundle;
    }

    @Import(name="excludePatterns")
    private @Nullable Output<List<String>> excludePatterns;

    public Optional<Output<List<String>>> excludePatterns() {
        return Optional.ofNullable(this.excludePatterns);
    }

    @Import(name="includePatterns")
    private @Nullable Output<List<String>> includePatterns;

    public Optional<Output<List<String>>> includePatterns() {
        return Optional.ofNullable(this.includePatterns);
    }

    @Import(name="registeredReleaseBundleNames", required=true)
    private Output<List<String>> registeredReleaseBundleNames;

    public Output<List<String>> registeredReleaseBundleNames() {
        return this.registeredReleaseBundleNames;
    }

    private ReleaseBundleWebhookCriteriaArgs() {}

    private ReleaseBundleWebhookCriteriaArgs(ReleaseBundleWebhookCriteriaArgs $) {
        this.anyReleaseBundle = $.anyReleaseBundle;
        this.excludePatterns = $.excludePatterns;
        this.includePatterns = $.includePatterns;
        this.registeredReleaseBundleNames = $.registeredReleaseBundleNames;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReleaseBundleWebhookCriteriaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReleaseBundleWebhookCriteriaArgs $;

        public Builder() {
            $ = new ReleaseBundleWebhookCriteriaArgs();
        }

        public Builder(ReleaseBundleWebhookCriteriaArgs defaults) {
            $ = new ReleaseBundleWebhookCriteriaArgs(Objects.requireNonNull(defaults));
        }

        public Builder anyReleaseBundle(Output<Boolean> anyReleaseBundle) {
            $.anyReleaseBundle = anyReleaseBundle;
            return this;
        }

        public Builder anyReleaseBundle(Boolean anyReleaseBundle) {
            return anyReleaseBundle(Output.of(anyReleaseBundle));
        }

        public Builder excludePatterns(@Nullable Output<List<String>> excludePatterns) {
            $.excludePatterns = excludePatterns;
            return this;
        }

        public Builder excludePatterns(List<String> excludePatterns) {
            return excludePatterns(Output.of(excludePatterns));
        }

        public Builder excludePatterns(String... excludePatterns) {
            return excludePatterns(List.of(excludePatterns));
        }

        public Builder includePatterns(@Nullable Output<List<String>> includePatterns) {
            $.includePatterns = includePatterns;
            return this;
        }

        public Builder includePatterns(List<String> includePatterns) {
            return includePatterns(Output.of(includePatterns));
        }

        public Builder includePatterns(String... includePatterns) {
            return includePatterns(List.of(includePatterns));
        }

        public Builder registeredReleaseBundleNames(Output<List<String>> registeredReleaseBundleNames) {
            $.registeredReleaseBundleNames = registeredReleaseBundleNames;
            return this;
        }

        public Builder registeredReleaseBundleNames(List<String> registeredReleaseBundleNames) {
            return registeredReleaseBundleNames(Output.of(registeredReleaseBundleNames));
        }

        public Builder registeredReleaseBundleNames(String... registeredReleaseBundleNames) {
            return registeredReleaseBundleNames(List.of(registeredReleaseBundleNames));
        }

        public ReleaseBundleWebhookCriteriaArgs build() {
            $.anyReleaseBundle = Objects.requireNonNull($.anyReleaseBundle, "expected parameter 'anyReleaseBundle' to be non-null");
            $.registeredReleaseBundleNames = Objects.requireNonNull($.registeredReleaseBundleNames, "expected parameter 'registeredReleaseBundleNames' to be non-null");
            return $;
        }
    }

}

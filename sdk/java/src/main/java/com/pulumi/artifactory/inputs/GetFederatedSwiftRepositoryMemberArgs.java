// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class GetFederatedSwiftRepositoryMemberArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetFederatedSwiftRepositoryMemberArgs Empty = new GetFederatedSwiftRepositoryMemberArgs();

    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    public Output<Boolean> enabled() {
        return this.enabled;
    }

    @Import(name="url", required=true)
    private Output<String> url;

    public Output<String> url() {
        return this.url;
    }

    private GetFederatedSwiftRepositoryMemberArgs() {}

    private GetFederatedSwiftRepositoryMemberArgs(GetFederatedSwiftRepositoryMemberArgs $) {
        this.enabled = $.enabled;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFederatedSwiftRepositoryMemberArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFederatedSwiftRepositoryMemberArgs $;

        public Builder() {
            $ = new GetFederatedSwiftRepositoryMemberArgs();
        }

        public Builder(GetFederatedSwiftRepositoryMemberArgs defaults) {
            $ = new GetFederatedSwiftRepositoryMemberArgs(Objects.requireNonNull(defaults));
        }

        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        public Builder url(String url) {
            return url(Output.of(url));
        }

        public GetFederatedSwiftRepositoryMemberArgs build() {
            $.enabled = Objects.requireNonNull($.enabled, "expected parameter 'enabled' to be non-null");
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}

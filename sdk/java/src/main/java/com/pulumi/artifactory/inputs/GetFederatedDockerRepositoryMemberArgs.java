// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class GetFederatedDockerRepositoryMemberArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetFederatedDockerRepositoryMemberArgs Empty = new GetFederatedDockerRepositoryMemberArgs();

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

    private GetFederatedDockerRepositoryMemberArgs() {}

    private GetFederatedDockerRepositoryMemberArgs(GetFederatedDockerRepositoryMemberArgs $) {
        this.enabled = $.enabled;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFederatedDockerRepositoryMemberArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFederatedDockerRepositoryMemberArgs $;

        public Builder() {
            $ = new GetFederatedDockerRepositoryMemberArgs();
        }

        public Builder(GetFederatedDockerRepositoryMemberArgs defaults) {
            $ = new GetFederatedDockerRepositoryMemberArgs(Objects.requireNonNull(defaults));
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

        public GetFederatedDockerRepositoryMemberArgs build() {
            if ($.enabled == null) {
                throw new MissingRequiredPropertyException("GetFederatedDockerRepositoryMemberArgs", "enabled");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("GetFederatedDockerRepositoryMemberArgs", "url");
            }
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class GetFederatedPuppetRepositoryMember extends com.pulumi.resources.InvokeArgs {

    public static final GetFederatedPuppetRepositoryMember Empty = new GetFederatedPuppetRepositoryMember();

    @Import(name="enabled", required=true)
    private Boolean enabled;

    public Boolean enabled() {
        return this.enabled;
    }

    @Import(name="url", required=true)
    private String url;

    public String url() {
        return this.url;
    }

    private GetFederatedPuppetRepositoryMember() {}

    private GetFederatedPuppetRepositoryMember(GetFederatedPuppetRepositoryMember $) {
        this.enabled = $.enabled;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFederatedPuppetRepositoryMember defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFederatedPuppetRepositoryMember $;

        public Builder() {
            $ = new GetFederatedPuppetRepositoryMember();
        }

        public Builder(GetFederatedPuppetRepositoryMember defaults) {
            $ = new GetFederatedPuppetRepositoryMember(Objects.requireNonNull(defaults));
        }

        public Builder enabled(Boolean enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder url(String url) {
            $.url = url;
            return this;
        }

        public GetFederatedPuppetRepositoryMember build() {
            $.enabled = Objects.requireNonNull($.enabled, "expected parameter 'enabled' to be non-null");
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}

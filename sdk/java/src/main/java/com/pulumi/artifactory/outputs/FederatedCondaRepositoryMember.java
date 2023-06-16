// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class FederatedCondaRepositoryMember {
    private Boolean enabled;
    private String url;

    private FederatedCondaRepositoryMember() {}
    public Boolean enabled() {
        return this.enabled;
    }
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FederatedCondaRepositoryMember defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean enabled;
        private String url;
        public Builder() {}
        public Builder(FederatedCondaRepositoryMember defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        public FederatedCondaRepositoryMember build() {
            final var o = new FederatedCondaRepositoryMember();
            o.enabled = enabled;
            o.url = url;
            return o;
        }
    }
}

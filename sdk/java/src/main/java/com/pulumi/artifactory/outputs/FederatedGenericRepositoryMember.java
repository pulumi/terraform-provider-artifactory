// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class FederatedGenericRepositoryMember {
    /**
     * @return Represents the active state of the federated member. It is supported to change the enabled
     * status of my own member. The config will be updated on the other federated members automatically.
     * 
     */
    private Boolean enabled;
    /**
     * @return Full URL to ending with the repository name.
     * 
     */
    private String url;

    private FederatedGenericRepositoryMember() {}
    /**
     * @return Represents the active state of the federated member. It is supported to change the enabled
     * status of my own member. The config will be updated on the other federated members automatically.
     * 
     */
    public Boolean enabled() {
        return this.enabled;
    }
    /**
     * @return Full URL to ending with the repository name.
     * 
     */
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FederatedGenericRepositoryMember defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean enabled;
        private String url;
        public Builder() {}
        public Builder(FederatedGenericRepositoryMember defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            if (enabled == null) {
              throw new MissingRequiredPropertyException("FederatedGenericRepositoryMember", "enabled");
            }
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("FederatedGenericRepositoryMember", "url");
            }
            this.url = url;
            return this;
        }
        public FederatedGenericRepositoryMember build() {
            final var _resultValue = new FederatedGenericRepositoryMember();
            _resultValue.enabled = enabled;
            _resultValue.url = url;
            return _resultValue;
        }
    }
}

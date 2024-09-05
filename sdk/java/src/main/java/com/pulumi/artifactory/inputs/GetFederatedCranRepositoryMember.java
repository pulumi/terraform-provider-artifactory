// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFederatedCranRepositoryMember extends com.pulumi.resources.InvokeArgs {

    public static final GetFederatedCranRepositoryMember Empty = new GetFederatedCranRepositoryMember();

    /**
     * Admin access token for this member Artifactory instance. Used in conjunction with `cleanup_on_delete` attribute when Access Federation for access tokens is not enabled.
     * 
     */
    @Import(name="accessToken")
    private @Nullable String accessToken;

    /**
     * @return Admin access token for this member Artifactory instance. Used in conjunction with `cleanup_on_delete` attribute when Access Federation for access tokens is not enabled.
     * 
     */
    public Optional<String> accessToken() {
        return Optional.ofNullable(this.accessToken);
    }

    /**
     * Represents the active state of the federated member. It is supported to change the enabled
     * status of my own member. The config will be updated on the other federated members automatically.
     * 
     */
    @Import(name="enabled", required=true)
    private Boolean enabled;

    /**
     * @return Represents the active state of the federated member. It is supported to change the enabled
     * status of my own member. The config will be updated on the other federated members automatically.
     * 
     */
    public Boolean enabled() {
        return this.enabled;
    }

    /**
     * Full URL to ending with the repository name.
     * 
     */
    @Import(name="url", required=true)
    private String url;

    /**
     * @return Full URL to ending with the repository name.
     * 
     */
    public String url() {
        return this.url;
    }

    private GetFederatedCranRepositoryMember() {}

    private GetFederatedCranRepositoryMember(GetFederatedCranRepositoryMember $) {
        this.accessToken = $.accessToken;
        this.enabled = $.enabled;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFederatedCranRepositoryMember defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFederatedCranRepositoryMember $;

        public Builder() {
            $ = new GetFederatedCranRepositoryMember();
        }

        public Builder(GetFederatedCranRepositoryMember defaults) {
            $ = new GetFederatedCranRepositoryMember(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessToken Admin access token for this member Artifactory instance. Used in conjunction with `cleanup_on_delete` attribute when Access Federation for access tokens is not enabled.
         * 
         * @return builder
         * 
         */
        public Builder accessToken(@Nullable String accessToken) {
            $.accessToken = accessToken;
            return this;
        }

        /**
         * @param enabled Represents the active state of the federated member. It is supported to change the enabled
         * status of my own member. The config will be updated on the other federated members automatically.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param url Full URL to ending with the repository name.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            $.url = url;
            return this;
        }

        public GetFederatedCranRepositoryMember build() {
            if ($.enabled == null) {
                throw new MissingRequiredPropertyException("GetFederatedCranRepositoryMember", "enabled");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("GetFederatedCranRepositoryMember", "url");
            }
            return $;
        }
    }

}

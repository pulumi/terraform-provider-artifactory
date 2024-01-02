// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class GetFederatedDockerV1RepositoryMember extends com.pulumi.resources.InvokeArgs {

    public static final GetFederatedDockerV1RepositoryMember Empty = new GetFederatedDockerV1RepositoryMember();

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

    private GetFederatedDockerV1RepositoryMember() {}

    private GetFederatedDockerV1RepositoryMember(GetFederatedDockerV1RepositoryMember $) {
        this.enabled = $.enabled;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFederatedDockerV1RepositoryMember defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFederatedDockerV1RepositoryMember $;

        public Builder() {
            $ = new GetFederatedDockerV1RepositoryMember();
        }

        public Builder(GetFederatedDockerV1RepositoryMember defaults) {
            $ = new GetFederatedDockerV1RepositoryMember(Objects.requireNonNull(defaults));
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

        public GetFederatedDockerV1RepositoryMember build() {
            if ($.enabled == null) {
                throw new MissingRequiredPropertyException("GetFederatedDockerV1RepositoryMember", "enabled");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("GetFederatedDockerV1RepositoryMember", "url");
            }
            return $;
        }
    }

}

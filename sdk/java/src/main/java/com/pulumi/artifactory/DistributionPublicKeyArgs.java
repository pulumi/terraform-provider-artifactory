// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class DistributionPublicKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final DistributionPublicKeyArgs Empty = new DistributionPublicKeyArgs();

    /**
     * Will be used as an identifier when uploading/retrieving the public key via REST API.
     * 
     */
    @Import(name="alias", required=true)
    private Output<String> alias;

    /**
     * @return Will be used as an identifier when uploading/retrieving the public key via REST API.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }

    /**
     * The Public key to add as a trusted distribution GPG key.
     * 
     * The following additional attributes are exported:
     * 
     */
    @Import(name="publicKey", required=true)
    private Output<String> publicKey;

    /**
     * @return The Public key to add as a trusted distribution GPG key.
     * 
     * The following additional attributes are exported:
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }

    private DistributionPublicKeyArgs() {}

    private DistributionPublicKeyArgs(DistributionPublicKeyArgs $) {
        this.alias = $.alias;
        this.publicKey = $.publicKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DistributionPublicKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DistributionPublicKeyArgs $;

        public Builder() {
            $ = new DistributionPublicKeyArgs();
        }

        public Builder(DistributionPublicKeyArgs defaults) {
            $ = new DistributionPublicKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alias Will be used as an identifier when uploading/retrieving the public key via REST API.
         * 
         * @return builder
         * 
         */
        public Builder alias(Output<String> alias) {
            $.alias = alias;
            return this;
        }

        /**
         * @param alias Will be used as an identifier when uploading/retrieving the public key via REST API.
         * 
         * @return builder
         * 
         */
        public Builder alias(String alias) {
            return alias(Output.of(alias));
        }

        /**
         * @param publicKey The Public key to add as a trusted distribution GPG key.
         * 
         * The following additional attributes are exported:
         * 
         * @return builder
         * 
         */
        public Builder publicKey(Output<String> publicKey) {
            $.publicKey = publicKey;
            return this;
        }

        /**
         * @param publicKey The Public key to add as a trusted distribution GPG key.
         * 
         * The following additional attributes are exported:
         * 
         * @return builder
         * 
         */
        public Builder publicKey(String publicKey) {
            return publicKey(Output.of(publicKey));
        }

        public DistributionPublicKeyArgs build() {
            $.alias = Objects.requireNonNull($.alias, "expected parameter 'alias' to be non-null");
            $.publicKey = Objects.requireNonNull($.publicKey, "expected parameter 'publicKey' to be non-null");
            return $;
        }
    }

}

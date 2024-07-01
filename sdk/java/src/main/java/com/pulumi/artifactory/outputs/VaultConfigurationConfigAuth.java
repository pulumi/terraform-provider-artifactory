// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VaultConfigurationConfigAuth {
    /**
     * @return Client certificate (in PEM format) for `Certificate` type.
     * 
     */
    private @Nullable String certificate;
    /**
     * @return Private key (in PEM format) for `Certificate` type.
     * 
     */
    private @Nullable String certificateKey;
    /**
     * @return Role ID for `AppRole` type
     * 
     */
    private @Nullable String roleId;
    /**
     * @return Secret ID for `AppRole` type
     * 
     */
    private @Nullable String secretId;
    private String type;

    private VaultConfigurationConfigAuth() {}
    /**
     * @return Client certificate (in PEM format) for `Certificate` type.
     * 
     */
    public Optional<String> certificate() {
        return Optional.ofNullable(this.certificate);
    }
    /**
     * @return Private key (in PEM format) for `Certificate` type.
     * 
     */
    public Optional<String> certificateKey() {
        return Optional.ofNullable(this.certificateKey);
    }
    /**
     * @return Role ID for `AppRole` type
     * 
     */
    public Optional<String> roleId() {
        return Optional.ofNullable(this.roleId);
    }
    /**
     * @return Secret ID for `AppRole` type
     * 
     */
    public Optional<String> secretId() {
        return Optional.ofNullable(this.secretId);
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VaultConfigurationConfigAuth defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String certificate;
        private @Nullable String certificateKey;
        private @Nullable String roleId;
        private @Nullable String secretId;
        private String type;
        public Builder() {}
        public Builder(VaultConfigurationConfigAuth defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.certificate = defaults.certificate;
    	      this.certificateKey = defaults.certificateKey;
    	      this.roleId = defaults.roleId;
    	      this.secretId = defaults.secretId;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder certificate(@Nullable String certificate) {

            this.certificate = certificate;
            return this;
        }
        @CustomType.Setter
        public Builder certificateKey(@Nullable String certificateKey) {

            this.certificateKey = certificateKey;
            return this;
        }
        @CustomType.Setter
        public Builder roleId(@Nullable String roleId) {

            this.roleId = roleId;
            return this;
        }
        @CustomType.Setter
        public Builder secretId(@Nullable String secretId) {

            this.secretId = secretId;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("VaultConfigurationConfigAuth", "type");
            }
            this.type = type;
            return this;
        }
        public VaultConfigurationConfigAuth build() {
            final var _resultValue = new VaultConfigurationConfigAuth();
            _resultValue.certificate = certificate;
            _resultValue.certificateKey = certificateKey;
            _resultValue.roleId = roleId;
            _resultValue.secretId = secretId;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}

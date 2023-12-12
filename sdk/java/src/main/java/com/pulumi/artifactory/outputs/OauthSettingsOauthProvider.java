// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OauthSettingsOauthProvider {
    /**
     * @return OAuth user info endpoint for the IdP.
     * 
     */
    private String apiUrl;
    /**
     * @return OAuth authorization endpoint for the IdP.
     * 
     */
    private String authUrl;
    /**
     * @return OAuth client ID configured on the IdP.
     * 
     */
    private String clientId;
    /**
     * @return OAuth client secret configured on the IdP.
     * 
     */
    private String clientSecret;
    /**
     * @return Enable the Artifactory OAuth provider.  Default value is `true`.
     * 
     */
    private @Nullable Boolean enabled;
    /**
     * @return Name of the Artifactory OAuth provider.
     * 
     */
    private String name;
    /**
     * @return OAuth token endpoint for the IdP.
     * 
     */
    private String tokenUrl;
    /**
     * @return Type of OAuth provider. (e.g., `github`, `google`, `cloudfoundry`, or `openId`)
     * 
     */
    private String type;

    private OauthSettingsOauthProvider() {}
    /**
     * @return OAuth user info endpoint for the IdP.
     * 
     */
    public String apiUrl() {
        return this.apiUrl;
    }
    /**
     * @return OAuth authorization endpoint for the IdP.
     * 
     */
    public String authUrl() {
        return this.authUrl;
    }
    /**
     * @return OAuth client ID configured on the IdP.
     * 
     */
    public String clientId() {
        return this.clientId;
    }
    /**
     * @return OAuth client secret configured on the IdP.
     * 
     */
    public String clientSecret() {
        return this.clientSecret;
    }
    /**
     * @return Enable the Artifactory OAuth provider.  Default value is `true`.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    /**
     * @return Name of the Artifactory OAuth provider.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return OAuth token endpoint for the IdP.
     * 
     */
    public String tokenUrl() {
        return this.tokenUrl;
    }
    /**
     * @return Type of OAuth provider. (e.g., `github`, `google`, `cloudfoundry`, or `openId`)
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OauthSettingsOauthProvider defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String apiUrl;
        private String authUrl;
        private String clientId;
        private String clientSecret;
        private @Nullable Boolean enabled;
        private String name;
        private String tokenUrl;
        private String type;
        public Builder() {}
        public Builder(OauthSettingsOauthProvider defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiUrl = defaults.apiUrl;
    	      this.authUrl = defaults.authUrl;
    	      this.clientId = defaults.clientId;
    	      this.clientSecret = defaults.clientSecret;
    	      this.enabled = defaults.enabled;
    	      this.name = defaults.name;
    	      this.tokenUrl = defaults.tokenUrl;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder apiUrl(String apiUrl) {
            this.apiUrl = Objects.requireNonNull(apiUrl);
            return this;
        }
        @CustomType.Setter
        public Builder authUrl(String authUrl) {
            this.authUrl = Objects.requireNonNull(authUrl);
            return this;
        }
        @CustomType.Setter
        public Builder clientId(String clientId) {
            this.clientId = Objects.requireNonNull(clientId);
            return this;
        }
        @CustomType.Setter
        public Builder clientSecret(String clientSecret) {
            this.clientSecret = Objects.requireNonNull(clientSecret);
            return this;
        }
        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder tokenUrl(String tokenUrl) {
            this.tokenUrl = Objects.requireNonNull(tokenUrl);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public OauthSettingsOauthProvider build() {
            final var _resultValue = new OauthSettingsOauthProvider();
            _resultValue.apiUrl = apiUrl;
            _resultValue.authUrl = authUrl;
            _resultValue.clientId = clientId;
            _resultValue.clientSecret = clientSecret;
            _resultValue.enabled = enabled;
            _resultValue.name = name;
            _resultValue.tokenUrl = tokenUrl;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}

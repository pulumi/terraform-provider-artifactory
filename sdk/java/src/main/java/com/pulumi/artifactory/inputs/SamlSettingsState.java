// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SamlSettingsState extends com.pulumi.resources.ResourceArgs {

    public static final SamlSettingsState Empty = new SamlSettingsState();

    /**
     * Allow persisted users to access their profile. Default value is &#34;true&#34;.
     * 
     */
    @Import(name="allowUserToAccessProfile")
    private @Nullable Output<Boolean> allowUserToAccessProfile;

    /**
     * @return Allow persisted users to access their profile. Default value is &#34;true&#34;.
     * 
     */
    public Optional<Output<Boolean>> allowUserToAccessProfile() {
        return Optional.ofNullable(this.allowUserToAccessProfile);
    }

    /**
     * Auto redirect to login through the IdP when clicking on Artifactory&#39;s login link. Default value is &#34;false&#34;.
     * 
     */
    @Import(name="autoRedirect")
    private @Nullable Output<Boolean> autoRedirect;

    /**
     * @return Auto redirect to login through the IdP when clicking on Artifactory&#39;s login link. Default value is &#34;false&#34;.
     * 
     */
    public Optional<Output<Boolean>> autoRedirect() {
        return Optional.ofNullable(this.autoRedirect);
    }

    /**
     * SAML certificate that contains the public key for the IdP service provider. Used by Artifactory to verify sign-in
     * requests. Default value is &#34;&#34;.
     * 
     */
    @Import(name="certificate")
    private @Nullable Output<String> certificate;

    /**
     * @return SAML certificate that contains the public key for the IdP service provider. Used by Artifactory to verify sign-in
     * requests. Default value is &#34;&#34;.
     * 
     */
    public Optional<Output<String>> certificate() {
        return Optional.ofNullable(this.certificate);
    }

    /**
     * Name of the attribute in the SAML response from the IdP that contains the user&#39;s email. Default value is &#34;&#34;.
     * 
     */
    @Import(name="emailAttribute")
    private @Nullable Output<String> emailAttribute;

    /**
     * @return Name of the attribute in the SAML response from the IdP that contains the user&#39;s email. Default value is &#34;&#34;.
     * 
     */
    public Optional<Output<String>> emailAttribute() {
        return Optional.ofNullable(this.emailAttribute);
    }

    /**
     * Enable SAML SSO. Default value is &#34;true&#34;.
     * 
     */
    @Import(name="enable")
    private @Nullable Output<Boolean> enable;

    /**
     * @return Enable SAML SSO. Default value is &#34;true&#34;.
     * 
     */
    public Optional<Output<Boolean>> enable() {
        return Optional.ofNullable(this.enable);
    }

    /**
     * Name of the attribute in the SAML response from the IdP that contains the user&#39;s group memberships. Default value is &#34;&#34;.
     * 
     */
    @Import(name="groupAttribute")
    private @Nullable Output<String> groupAttribute;

    /**
     * @return Name of the attribute in the SAML response from the IdP that contains the user&#39;s group memberships. Default value is &#34;&#34;.
     * 
     */
    public Optional<Output<String>> groupAttribute() {
        return Optional.ofNullable(this.groupAttribute);
    }

    /**
     * Service provider login url configured on the IdP.
     * 
     */
    @Import(name="loginUrl")
    private @Nullable Output<String> loginUrl;

    /**
     * @return Service provider login url configured on the IdP.
     * 
     */
    public Optional<Output<String>> loginUrl() {
        return Optional.ofNullable(this.loginUrl);
    }

    /**
     * Service provider logout url, or where to redirect after user logs out.
     * 
     */
    @Import(name="logoutUrl")
    private @Nullable Output<String> logoutUrl;

    /**
     * @return Service provider logout url, or where to redirect after user logs out.
     * 
     */
    public Optional<Output<String>> logoutUrl() {
        return Optional.ofNullable(this.logoutUrl);
    }

    /**
     * When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for
     * every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and
     * the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory
     * to manage user permissions not attached to their default groups. Default value is &#34;false&#34;.
     * 
     */
    @Import(name="noAutoUserCreation")
    private @Nullable Output<Boolean> noAutoUserCreation;

    /**
     * @return When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for
     * every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and
     * the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory
     * to manage user permissions not attached to their default groups. Default value is &#34;false&#34;.
     * 
     */
    public Optional<Output<Boolean>> noAutoUserCreation() {
        return Optional.ofNullable(this.noAutoUserCreation);
    }

    /**
     * The SAML service provider name. This should be a URI that is also known as the entityID, providerID, or entity identity.
     * 
     */
    @Import(name="serviceProviderName")
    private @Nullable Output<String> serviceProviderName;

    /**
     * @return The SAML service provider name. This should be a URI that is also known as the entityID, providerID, or entity identity.
     * 
     */
    public Optional<Output<String>> serviceProviderName() {
        return Optional.ofNullable(this.serviceProviderName);
    }

    /**
     * Associate user with Artifactory groups based on the &#34;group_attribute&#34; provided in the SAML response from the identity
     * provider. Default value is &#34;false&#34;.
     * 
     */
    @Import(name="syncGroups")
    private @Nullable Output<Boolean> syncGroups;

    /**
     * @return Associate user with Artifactory groups based on the &#34;group_attribute&#34; provided in the SAML response from the identity
     * provider. Default value is &#34;false&#34;.
     * 
     */
    public Optional<Output<Boolean>> syncGroups() {
        return Optional.ofNullable(this.syncGroups);
    }

    /**
     * When set, an X.509 public certificate will be created by Artifactory. Download this certificate and upload it to your
     * IDP and choose your own encryption algorithm. This process will let you encrypt the assertion section in your SAML
     * response. Default value is &#34;false&#34;.
     * 
     */
    @Import(name="useEncryptedAssertion")
    private @Nullable Output<Boolean> useEncryptedAssertion;

    /**
     * @return When set, an X.509 public certificate will be created by Artifactory. Download this certificate and upload it to your
     * IDP and choose your own encryption algorithm. This process will let you encrypt the assertion section in your SAML
     * response. Default value is &#34;false&#34;.
     * 
     */
    public Optional<Output<Boolean>> useEncryptedAssertion() {
        return Optional.ofNullable(this.useEncryptedAssertion);
    }

    /**
     * Enable &#34;audience&#34;, or who the SAML assertion is intended for. Ensures that the correct service provider intended for
     * Artifactory is used on the IdP. Default value is &#34;true&#34;.
     * 
     */
    @Import(name="verifyAudienceRestriction")
    private @Nullable Output<Boolean> verifyAudienceRestriction;

    /**
     * @return Enable &#34;audience&#34;, or who the SAML assertion is intended for. Ensures that the correct service provider intended for
     * Artifactory is used on the IdP. Default value is &#34;true&#34;.
     * 
     */
    public Optional<Output<Boolean>> verifyAudienceRestriction() {
        return Optional.ofNullable(this.verifyAudienceRestriction);
    }

    private SamlSettingsState() {}

    private SamlSettingsState(SamlSettingsState $) {
        this.allowUserToAccessProfile = $.allowUserToAccessProfile;
        this.autoRedirect = $.autoRedirect;
        this.certificate = $.certificate;
        this.emailAttribute = $.emailAttribute;
        this.enable = $.enable;
        this.groupAttribute = $.groupAttribute;
        this.loginUrl = $.loginUrl;
        this.logoutUrl = $.logoutUrl;
        this.noAutoUserCreation = $.noAutoUserCreation;
        this.serviceProviderName = $.serviceProviderName;
        this.syncGroups = $.syncGroups;
        this.useEncryptedAssertion = $.useEncryptedAssertion;
        this.verifyAudienceRestriction = $.verifyAudienceRestriction;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SamlSettingsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SamlSettingsState $;

        public Builder() {
            $ = new SamlSettingsState();
        }

        public Builder(SamlSettingsState defaults) {
            $ = new SamlSettingsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowUserToAccessProfile Allow persisted users to access their profile. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder allowUserToAccessProfile(@Nullable Output<Boolean> allowUserToAccessProfile) {
            $.allowUserToAccessProfile = allowUserToAccessProfile;
            return this;
        }

        /**
         * @param allowUserToAccessProfile Allow persisted users to access their profile. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder allowUserToAccessProfile(Boolean allowUserToAccessProfile) {
            return allowUserToAccessProfile(Output.of(allowUserToAccessProfile));
        }

        /**
         * @param autoRedirect Auto redirect to login through the IdP when clicking on Artifactory&#39;s login link. Default value is &#34;false&#34;.
         * 
         * @return builder
         * 
         */
        public Builder autoRedirect(@Nullable Output<Boolean> autoRedirect) {
            $.autoRedirect = autoRedirect;
            return this;
        }

        /**
         * @param autoRedirect Auto redirect to login through the IdP when clicking on Artifactory&#39;s login link. Default value is &#34;false&#34;.
         * 
         * @return builder
         * 
         */
        public Builder autoRedirect(Boolean autoRedirect) {
            return autoRedirect(Output.of(autoRedirect));
        }

        /**
         * @param certificate SAML certificate that contains the public key for the IdP service provider. Used by Artifactory to verify sign-in
         * requests. Default value is &#34;&#34;.
         * 
         * @return builder
         * 
         */
        public Builder certificate(@Nullable Output<String> certificate) {
            $.certificate = certificate;
            return this;
        }

        /**
         * @param certificate SAML certificate that contains the public key for the IdP service provider. Used by Artifactory to verify sign-in
         * requests. Default value is &#34;&#34;.
         * 
         * @return builder
         * 
         */
        public Builder certificate(String certificate) {
            return certificate(Output.of(certificate));
        }

        /**
         * @param emailAttribute Name of the attribute in the SAML response from the IdP that contains the user&#39;s email. Default value is &#34;&#34;.
         * 
         * @return builder
         * 
         */
        public Builder emailAttribute(@Nullable Output<String> emailAttribute) {
            $.emailAttribute = emailAttribute;
            return this;
        }

        /**
         * @param emailAttribute Name of the attribute in the SAML response from the IdP that contains the user&#39;s email. Default value is &#34;&#34;.
         * 
         * @return builder
         * 
         */
        public Builder emailAttribute(String emailAttribute) {
            return emailAttribute(Output.of(emailAttribute));
        }

        /**
         * @param enable Enable SAML SSO. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder enable(@Nullable Output<Boolean> enable) {
            $.enable = enable;
            return this;
        }

        /**
         * @param enable Enable SAML SSO. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder enable(Boolean enable) {
            return enable(Output.of(enable));
        }

        /**
         * @param groupAttribute Name of the attribute in the SAML response from the IdP that contains the user&#39;s group memberships. Default value is &#34;&#34;.
         * 
         * @return builder
         * 
         */
        public Builder groupAttribute(@Nullable Output<String> groupAttribute) {
            $.groupAttribute = groupAttribute;
            return this;
        }

        /**
         * @param groupAttribute Name of the attribute in the SAML response from the IdP that contains the user&#39;s group memberships. Default value is &#34;&#34;.
         * 
         * @return builder
         * 
         */
        public Builder groupAttribute(String groupAttribute) {
            return groupAttribute(Output.of(groupAttribute));
        }

        /**
         * @param loginUrl Service provider login url configured on the IdP.
         * 
         * @return builder
         * 
         */
        public Builder loginUrl(@Nullable Output<String> loginUrl) {
            $.loginUrl = loginUrl;
            return this;
        }

        /**
         * @param loginUrl Service provider login url configured on the IdP.
         * 
         * @return builder
         * 
         */
        public Builder loginUrl(String loginUrl) {
            return loginUrl(Output.of(loginUrl));
        }

        /**
         * @param logoutUrl Service provider logout url, or where to redirect after user logs out.
         * 
         * @return builder
         * 
         */
        public Builder logoutUrl(@Nullable Output<String> logoutUrl) {
            $.logoutUrl = logoutUrl;
            return this;
        }

        /**
         * @param logoutUrl Service provider logout url, or where to redirect after user logs out.
         * 
         * @return builder
         * 
         */
        public Builder logoutUrl(String logoutUrl) {
            return logoutUrl(Output.of(logoutUrl));
        }

        /**
         * @param noAutoUserCreation When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for
         * every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and
         * the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory
         * to manage user permissions not attached to their default groups. Default value is &#34;false&#34;.
         * 
         * @return builder
         * 
         */
        public Builder noAutoUserCreation(@Nullable Output<Boolean> noAutoUserCreation) {
            $.noAutoUserCreation = noAutoUserCreation;
            return this;
        }

        /**
         * @param noAutoUserCreation When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for
         * every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and
         * the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory
         * to manage user permissions not attached to their default groups. Default value is &#34;false&#34;.
         * 
         * @return builder
         * 
         */
        public Builder noAutoUserCreation(Boolean noAutoUserCreation) {
            return noAutoUserCreation(Output.of(noAutoUserCreation));
        }

        /**
         * @param serviceProviderName The SAML service provider name. This should be a URI that is also known as the entityID, providerID, or entity identity.
         * 
         * @return builder
         * 
         */
        public Builder serviceProviderName(@Nullable Output<String> serviceProviderName) {
            $.serviceProviderName = serviceProviderName;
            return this;
        }

        /**
         * @param serviceProviderName The SAML service provider name. This should be a URI that is also known as the entityID, providerID, or entity identity.
         * 
         * @return builder
         * 
         */
        public Builder serviceProviderName(String serviceProviderName) {
            return serviceProviderName(Output.of(serviceProviderName));
        }

        /**
         * @param syncGroups Associate user with Artifactory groups based on the &#34;group_attribute&#34; provided in the SAML response from the identity
         * provider. Default value is &#34;false&#34;.
         * 
         * @return builder
         * 
         */
        public Builder syncGroups(@Nullable Output<Boolean> syncGroups) {
            $.syncGroups = syncGroups;
            return this;
        }

        /**
         * @param syncGroups Associate user with Artifactory groups based on the &#34;group_attribute&#34; provided in the SAML response from the identity
         * provider. Default value is &#34;false&#34;.
         * 
         * @return builder
         * 
         */
        public Builder syncGroups(Boolean syncGroups) {
            return syncGroups(Output.of(syncGroups));
        }

        /**
         * @param useEncryptedAssertion When set, an X.509 public certificate will be created by Artifactory. Download this certificate and upload it to your
         * IDP and choose your own encryption algorithm. This process will let you encrypt the assertion section in your SAML
         * response. Default value is &#34;false&#34;.
         * 
         * @return builder
         * 
         */
        public Builder useEncryptedAssertion(@Nullable Output<Boolean> useEncryptedAssertion) {
            $.useEncryptedAssertion = useEncryptedAssertion;
            return this;
        }

        /**
         * @param useEncryptedAssertion When set, an X.509 public certificate will be created by Artifactory. Download this certificate and upload it to your
         * IDP and choose your own encryption algorithm. This process will let you encrypt the assertion section in your SAML
         * response. Default value is &#34;false&#34;.
         * 
         * @return builder
         * 
         */
        public Builder useEncryptedAssertion(Boolean useEncryptedAssertion) {
            return useEncryptedAssertion(Output.of(useEncryptedAssertion));
        }

        /**
         * @param verifyAudienceRestriction Enable &#34;audience&#34;, or who the SAML assertion is intended for. Ensures that the correct service provider intended for
         * Artifactory is used on the IdP. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder verifyAudienceRestriction(@Nullable Output<Boolean> verifyAudienceRestriction) {
            $.verifyAudienceRestriction = verifyAudienceRestriction;
            return this;
        }

        /**
         * @param verifyAudienceRestriction Enable &#34;audience&#34;, or who the SAML assertion is intended for. Ensures that the correct service provider intended for
         * Artifactory is used on the IdP. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder verifyAudienceRestriction(Boolean verifyAudienceRestriction) {
            return verifyAudienceRestriction(Output.of(verifyAudienceRestriction));
        }

        public SamlSettingsState build() {
            return $;
        }
    }

}

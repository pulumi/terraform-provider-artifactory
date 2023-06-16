// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LdapSettingArgs extends com.pulumi.resources.ResourceArgs {

    public static final LdapSettingArgs Empty = new LdapSettingArgs();

    /**
     * Auto created users will have access to their profile page and will be able to perform actions such as generating an API
     * key. Default value is &#34;false&#34;.
     * 
     */
    @Import(name="allowUserToAccessProfile")
    private @Nullable Output<Boolean> allowUserToAccessProfile;

    /**
     * @return Auto created users will have access to their profile page and will be able to perform actions such as generating an API
     * key. Default value is &#34;false&#34;.
     * 
     */
    public Optional<Output<Boolean>> allowUserToAccessProfile() {
        return Optional.ofNullable(this.allowUserToAccessProfile);
    }

    /**
     * When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join
     * groups defined in Artifactory. Default value is &#34;true&#34;.
     * 
     */
    @Import(name="autoCreateUser")
    private @Nullable Output<Boolean> autoCreateUser;

    /**
     * @return When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join
     * groups defined in Artifactory. Default value is &#34;true&#34;.
     * 
     */
    public Optional<Output<Boolean>> autoCreateUser() {
        return Optional.ofNullable(this.autoCreateUser);
    }

    /**
     * An attribute that can be used to map a user&#39;s email address to a user created automatically in Artifactory. Default
     * value is &#34;mail&#34;.
     * 
     */
    @Import(name="emailAttribute")
    private @Nullable Output<String> emailAttribute;

    /**
     * @return An attribute that can be used to map a user&#39;s email address to a user created automatically in Artifactory. Default
     * value is &#34;mail&#34;.
     * 
     */
    public Optional<Output<String>> emailAttribute() {
        return Optional.ofNullable(this.emailAttribute);
    }

    /**
     * Flag to enable or disable the ldap setting. Default value is &#34;true&#34;.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Flag to enable or disable the ldap setting. Default value is &#34;true&#34;.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Ldap setting name.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return Ldap setting name.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     * Protects against LDAP poisoning by filtering out users exposed to vulnerabilities. Default value is &#34;true&#34;.
     * 
     */
    @Import(name="ldapPoisoningProtection")
    private @Nullable Output<Boolean> ldapPoisoningProtection;

    /**
     * @return Protects against LDAP poisoning by filtering out users exposed to vulnerabilities. Default value is &#34;true&#34;.
     * 
     */
    public Optional<Output<Boolean>> ldapPoisoningProtection() {
        return Optional.ofNullable(this.ldapPoisoningProtection);
    }

    /**
     * Location of the LDAP server in the following format: ldap://myldapserver/dc=sampledomain,dc=com
     * 
     */
    @Import(name="ldapUrl", required=true)
    private Output<String> ldapUrl;

    /**
     * @return Location of the LDAP server in the following format: ldap://myldapserver/dc=sampledomain,dc=com
     * 
     */
    public Output<String> ldapUrl() {
        return this.ldapUrl;
    }

    /**
     * The full DN of the user that binds to the LDAP server to perform user searches. Only used with &#34;search&#34; authentication.
     * 
     */
    @Import(name="managerDn")
    private @Nullable Output<String> managerDn;

    /**
     * @return The full DN of the user that binds to the LDAP server to perform user searches. Only used with &#34;search&#34; authentication.
     * 
     */
    public Optional<Output<String>> managerDn() {
        return Optional.ofNullable(this.managerDn);
    }

    /**
     * The password of the user that binds to the LDAP server to perform the search. Only used with &#34;search&#34; authentication.
     * 
     */
    @Import(name="managerPassword")
    private @Nullable Output<String> managerPassword;

    /**
     * @return The password of the user that binds to the LDAP server to perform the search. Only used with &#34;search&#34; authentication.
     * 
     */
    public Optional<Output<String>> managerPassword() {
        return Optional.ofNullable(this.managerPassword);
    }

    /**
     * When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a
     * PagedResultsControl configuration. Default value is &#34;true&#34;.
     * 
     */
    @Import(name="pagingSupportEnabled")
    private @Nullable Output<Boolean> pagingSupportEnabled;

    /**
     * @return When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a
     * PagedResultsControl configuration. Default value is &#34;true&#34;.
     * 
     */
    public Optional<Output<Boolean>> pagingSupportEnabled() {
        return Optional.ofNullable(this.pagingSupportEnabled);
    }

    /**
     * A context name to search in relative to the base DN of the LDAP URL. For example, &#39;ou=users&#39; With the LDAP Group Add-on
     * enabled, it is possible to enter multiple search base entries separated by a pipe (&#39;|&#39;) character.
     * 
     */
    @Import(name="searchBase")
    private @Nullable Output<String> searchBase;

    /**
     * @return A context name to search in relative to the base DN of the LDAP URL. For example, &#39;ou=users&#39; With the LDAP Group Add-on
     * enabled, it is possible to enter multiple search base entries separated by a pipe (&#39;|&#39;) character.
     * 
     */
    public Optional<Output<String>> searchBase() {
        return Optional.ofNullable(this.searchBase);
    }

    /**
     * A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as
     * defined in &#39;RFC 2254&#39;) with optional arguments. In this case, the username is the only argument, and is denoted by
     * &#39;{0}&#39;. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is
     * performed from the DN found if successful.
     * 
     */
    @Import(name="searchFilter")
    private @Nullable Output<String> searchFilter;

    /**
     * @return A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as
     * defined in &#39;RFC 2254&#39;) with optional arguments. In this case, the username is the only argument, and is denoted by
     * &#39;{0}&#39;. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is
     * performed from the DN found if successful.
     * 
     */
    public Optional<Output<String>> searchFilter() {
        return Optional.ofNullable(this.searchFilter);
    }

    /**
     * When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is &#34;true&#34;.
     * 
     */
    @Import(name="searchSubTree")
    private @Nullable Output<Boolean> searchSubTree;

    /**
     * @return When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is &#34;true&#34;.
     * 
     */
    public Optional<Output<Boolean>> searchSubTree() {
        return Optional.ofNullable(this.searchSubTree);
    }

    /**
     * A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for &#39;direct&#39;
     * user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced
     * with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the
     * default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is
     * blank/empty.
     * 
     */
    @Import(name="userDnPattern")
    private @Nullable Output<String> userDnPattern;

    /**
     * @return A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for &#39;direct&#39;
     * user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced
     * with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the
     * default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is
     * blank/empty.
     * 
     */
    public Optional<Output<String>> userDnPattern() {
        return Optional.ofNullable(this.userDnPattern);
    }

    private LdapSettingArgs() {}

    private LdapSettingArgs(LdapSettingArgs $) {
        this.allowUserToAccessProfile = $.allowUserToAccessProfile;
        this.autoCreateUser = $.autoCreateUser;
        this.emailAttribute = $.emailAttribute;
        this.enabled = $.enabled;
        this.key = $.key;
        this.ldapPoisoningProtection = $.ldapPoisoningProtection;
        this.ldapUrl = $.ldapUrl;
        this.managerDn = $.managerDn;
        this.managerPassword = $.managerPassword;
        this.pagingSupportEnabled = $.pagingSupportEnabled;
        this.searchBase = $.searchBase;
        this.searchFilter = $.searchFilter;
        this.searchSubTree = $.searchSubTree;
        this.userDnPattern = $.userDnPattern;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LdapSettingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LdapSettingArgs $;

        public Builder() {
            $ = new LdapSettingArgs();
        }

        public Builder(LdapSettingArgs defaults) {
            $ = new LdapSettingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowUserToAccessProfile Auto created users will have access to their profile page and will be able to perform actions such as generating an API
         * key. Default value is &#34;false&#34;.
         * 
         * @return builder
         * 
         */
        public Builder allowUserToAccessProfile(@Nullable Output<Boolean> allowUserToAccessProfile) {
            $.allowUserToAccessProfile = allowUserToAccessProfile;
            return this;
        }

        /**
         * @param allowUserToAccessProfile Auto created users will have access to their profile page and will be able to perform actions such as generating an API
         * key. Default value is &#34;false&#34;.
         * 
         * @return builder
         * 
         */
        public Builder allowUserToAccessProfile(Boolean allowUserToAccessProfile) {
            return allowUserToAccessProfile(Output.of(allowUserToAccessProfile));
        }

        /**
         * @param autoCreateUser When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join
         * groups defined in Artifactory. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder autoCreateUser(@Nullable Output<Boolean> autoCreateUser) {
            $.autoCreateUser = autoCreateUser;
            return this;
        }

        /**
         * @param autoCreateUser When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join
         * groups defined in Artifactory. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder autoCreateUser(Boolean autoCreateUser) {
            return autoCreateUser(Output.of(autoCreateUser));
        }

        /**
         * @param emailAttribute An attribute that can be used to map a user&#39;s email address to a user created automatically in Artifactory. Default
         * value is &#34;mail&#34;.
         * 
         * @return builder
         * 
         */
        public Builder emailAttribute(@Nullable Output<String> emailAttribute) {
            $.emailAttribute = emailAttribute;
            return this;
        }

        /**
         * @param emailAttribute An attribute that can be used to map a user&#39;s email address to a user created automatically in Artifactory. Default
         * value is &#34;mail&#34;.
         * 
         * @return builder
         * 
         */
        public Builder emailAttribute(String emailAttribute) {
            return emailAttribute(Output.of(emailAttribute));
        }

        /**
         * @param enabled Flag to enable or disable the ldap setting. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Flag to enable or disable the ldap setting. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param key Ldap setting name.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key Ldap setting name.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param ldapPoisoningProtection Protects against LDAP poisoning by filtering out users exposed to vulnerabilities. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder ldapPoisoningProtection(@Nullable Output<Boolean> ldapPoisoningProtection) {
            $.ldapPoisoningProtection = ldapPoisoningProtection;
            return this;
        }

        /**
         * @param ldapPoisoningProtection Protects against LDAP poisoning by filtering out users exposed to vulnerabilities. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder ldapPoisoningProtection(Boolean ldapPoisoningProtection) {
            return ldapPoisoningProtection(Output.of(ldapPoisoningProtection));
        }

        /**
         * @param ldapUrl Location of the LDAP server in the following format: ldap://myldapserver/dc=sampledomain,dc=com
         * 
         * @return builder
         * 
         */
        public Builder ldapUrl(Output<String> ldapUrl) {
            $.ldapUrl = ldapUrl;
            return this;
        }

        /**
         * @param ldapUrl Location of the LDAP server in the following format: ldap://myldapserver/dc=sampledomain,dc=com
         * 
         * @return builder
         * 
         */
        public Builder ldapUrl(String ldapUrl) {
            return ldapUrl(Output.of(ldapUrl));
        }

        /**
         * @param managerDn The full DN of the user that binds to the LDAP server to perform user searches. Only used with &#34;search&#34; authentication.
         * 
         * @return builder
         * 
         */
        public Builder managerDn(@Nullable Output<String> managerDn) {
            $.managerDn = managerDn;
            return this;
        }

        /**
         * @param managerDn The full DN of the user that binds to the LDAP server to perform user searches. Only used with &#34;search&#34; authentication.
         * 
         * @return builder
         * 
         */
        public Builder managerDn(String managerDn) {
            return managerDn(Output.of(managerDn));
        }

        /**
         * @param managerPassword The password of the user that binds to the LDAP server to perform the search. Only used with &#34;search&#34; authentication.
         * 
         * @return builder
         * 
         */
        public Builder managerPassword(@Nullable Output<String> managerPassword) {
            $.managerPassword = managerPassword;
            return this;
        }

        /**
         * @param managerPassword The password of the user that binds to the LDAP server to perform the search. Only used with &#34;search&#34; authentication.
         * 
         * @return builder
         * 
         */
        public Builder managerPassword(String managerPassword) {
            return managerPassword(Output.of(managerPassword));
        }

        /**
         * @param pagingSupportEnabled When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a
         * PagedResultsControl configuration. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder pagingSupportEnabled(@Nullable Output<Boolean> pagingSupportEnabled) {
            $.pagingSupportEnabled = pagingSupportEnabled;
            return this;
        }

        /**
         * @param pagingSupportEnabled When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a
         * PagedResultsControl configuration. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder pagingSupportEnabled(Boolean pagingSupportEnabled) {
            return pagingSupportEnabled(Output.of(pagingSupportEnabled));
        }

        /**
         * @param searchBase A context name to search in relative to the base DN of the LDAP URL. For example, &#39;ou=users&#39; With the LDAP Group Add-on
         * enabled, it is possible to enter multiple search base entries separated by a pipe (&#39;|&#39;) character.
         * 
         * @return builder
         * 
         */
        public Builder searchBase(@Nullable Output<String> searchBase) {
            $.searchBase = searchBase;
            return this;
        }

        /**
         * @param searchBase A context name to search in relative to the base DN of the LDAP URL. For example, &#39;ou=users&#39; With the LDAP Group Add-on
         * enabled, it is possible to enter multiple search base entries separated by a pipe (&#39;|&#39;) character.
         * 
         * @return builder
         * 
         */
        public Builder searchBase(String searchBase) {
            return searchBase(Output.of(searchBase));
        }

        /**
         * @param searchFilter A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as
         * defined in &#39;RFC 2254&#39;) with optional arguments. In this case, the username is the only argument, and is denoted by
         * &#39;{0}&#39;. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is
         * performed from the DN found if successful.
         * 
         * @return builder
         * 
         */
        public Builder searchFilter(@Nullable Output<String> searchFilter) {
            $.searchFilter = searchFilter;
            return this;
        }

        /**
         * @param searchFilter A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as
         * defined in &#39;RFC 2254&#39;) with optional arguments. In this case, the username is the only argument, and is denoted by
         * &#39;{0}&#39;. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is
         * performed from the DN found if successful.
         * 
         * @return builder
         * 
         */
        public Builder searchFilter(String searchFilter) {
            return searchFilter(Output.of(searchFilter));
        }

        /**
         * @param searchSubTree When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder searchSubTree(@Nullable Output<Boolean> searchSubTree) {
            $.searchSubTree = searchSubTree;
            return this;
        }

        /**
         * @param searchSubTree When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is &#34;true&#34;.
         * 
         * @return builder
         * 
         */
        public Builder searchSubTree(Boolean searchSubTree) {
            return searchSubTree(Output.of(searchSubTree));
        }

        /**
         * @param userDnPattern A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for &#39;direct&#39;
         * user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced
         * with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the
         * default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is
         * blank/empty.
         * 
         * @return builder
         * 
         */
        public Builder userDnPattern(@Nullable Output<String> userDnPattern) {
            $.userDnPattern = userDnPattern;
            return this;
        }

        /**
         * @param userDnPattern A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for &#39;direct&#39;
         * user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced
         * with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the
         * default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is
         * blank/empty.
         * 
         * @return builder
         * 
         */
        public Builder userDnPattern(String userDnPattern) {
            return userDnPattern(Output.of(userDnPattern));
        }

        public LdapSettingArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            $.ldapUrl = Objects.requireNonNull($.ldapUrl, "expected parameter 'ldapUrl' to be non-null");
            return $;
        }
    }

}

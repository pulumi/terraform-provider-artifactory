// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetUserPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserPlainArgs Empty = new GetUserPlainArgs();

    /**
     * When enabled, this user is an administrator with all the ensuing privileges. Default value
     * is `false`.
     * 
     */
    @Import(name="admin")
    private @Nullable Boolean admin;

    /**
     * @return When enabled, this user is an administrator with all the ensuing privileges. Default value
     * is `false`.
     * 
     */
    public Optional<Boolean> admin() {
        return Optional.ofNullable(this.admin);
    }

    /**
     * When set, this user can only access Artifactory through the REST API. This option
     * cannot be set if the user has Admin privileges. Default value is `true`.
     * 
     */
    @Import(name="disableUiAccess")
    private @Nullable Boolean disableUiAccess;

    /**
     * @return When set, this user can only access Artifactory through the REST API. This option
     * cannot be set if the user has Admin privileges. Default value is `true`.
     * 
     */
    public Optional<Boolean> disableUiAccess() {
        return Optional.ofNullable(this.disableUiAccess);
    }

    /**
     * Email for user.
     * 
     */
    @Import(name="email")
    private @Nullable String email;

    /**
     * @return Email for user.
     * 
     */
    public Optional<String> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * List of groups this user is a part of.
     * 
     */
    @Import(name="groups")
    private @Nullable List<String> groups;

    /**
     * @return List of groups this user is a part of.
     * 
     */
    public Optional<List<String>> groups() {
        return Optional.ofNullable(this.groups);
    }

    /**
     * When set, disables the fallback of using an internal password when external
     * authentication (such as LDAP) is enabled.
     * 
     */
    @Import(name="internalPasswordDisabled")
    private @Nullable Boolean internalPasswordDisabled;

    /**
     * @return When set, disables the fallback of using an internal password when external
     * authentication (such as LDAP) is enabled.
     * 
     */
    public Optional<Boolean> internalPasswordDisabled() {
        return Optional.ofNullable(this.internalPasswordDisabled);
    }

    /**
     * Name of the user.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Name of the user.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * When set, this user can update his profile details (except for the password. Only an
     * administrator can update the password). Default value is `true`.
     * 
     */
    @Import(name="profileUpdatable")
    private @Nullable Boolean profileUpdatable;

    /**
     * @return When set, this user can update his profile details (except for the password. Only an
     * administrator can update the password). Default value is `true`.
     * 
     */
    public Optional<Boolean> profileUpdatable() {
        return Optional.ofNullable(this.profileUpdatable);
    }

    private GetUserPlainArgs() {}

    private GetUserPlainArgs(GetUserPlainArgs $) {
        this.admin = $.admin;
        this.disableUiAccess = $.disableUiAccess;
        this.email = $.email;
        this.groups = $.groups;
        this.internalPasswordDisabled = $.internalPasswordDisabled;
        this.name = $.name;
        this.profileUpdatable = $.profileUpdatable;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserPlainArgs $;

        public Builder() {
            $ = new GetUserPlainArgs();
        }

        public Builder(GetUserPlainArgs defaults) {
            $ = new GetUserPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param admin When enabled, this user is an administrator with all the ensuing privileges. Default value
         * is `false`.
         * 
         * @return builder
         * 
         */
        public Builder admin(@Nullable Boolean admin) {
            $.admin = admin;
            return this;
        }

        /**
         * @param disableUiAccess When set, this user can only access Artifactory through the REST API. This option
         * cannot be set if the user has Admin privileges. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder disableUiAccess(@Nullable Boolean disableUiAccess) {
            $.disableUiAccess = disableUiAccess;
            return this;
        }

        /**
         * @param email Email for user.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable String email) {
            $.email = email;
            return this;
        }

        /**
         * @param groups List of groups this user is a part of.
         * 
         * @return builder
         * 
         */
        public Builder groups(@Nullable List<String> groups) {
            $.groups = groups;
            return this;
        }

        /**
         * @param groups List of groups this user is a part of.
         * 
         * @return builder
         * 
         */
        public Builder groups(String... groups) {
            return groups(List.of(groups));
        }

        /**
         * @param internalPasswordDisabled When set, disables the fallback of using an internal password when external
         * authentication (such as LDAP) is enabled.
         * 
         * @return builder
         * 
         */
        public Builder internalPasswordDisabled(@Nullable Boolean internalPasswordDisabled) {
            $.internalPasswordDisabled = internalPasswordDisabled;
            return this;
        }

        /**
         * @param name Name of the user.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param profileUpdatable When set, this user can update his profile details (except for the password. Only an
         * administrator can update the password). Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder profileUpdatable(@Nullable Boolean profileUpdatable) {
            $.profileUpdatable = profileUpdatable;
            return this;
        }

        public GetUserPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}

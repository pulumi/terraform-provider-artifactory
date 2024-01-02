// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserArgs Empty = new UserArgs();

    /**
     * (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
     * 
     */
    @Import(name="admin")
    private @Nullable Output<Boolean> admin;

    /**
     * @return (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
     * 
     */
    public Optional<Output<Boolean>> admin() {
        return Optional.ofNullable(this.admin);
    }

    /**
     * (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
     * 
     */
    @Import(name="disableUiAccess")
    private @Nullable Output<Boolean> disableUiAccess;

    /**
     * @return (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
     * 
     */
    public Optional<Output<Boolean>> disableUiAccess() {
        return Optional.ofNullable(this.disableUiAccess);
    }

    /**
     * Email for user.
     * 
     */
    @Import(name="email", required=true)
    private Output<String> email;

    /**
     * @return Email for user.
     * 
     */
    public Output<String> email() {
        return this.email;
    }

    /**
     * List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
     * 
     */
    @Import(name="groups")
    private @Nullable Output<List<String>> groups;

    /**
     * @return List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
     * 
     */
    public Optional<Output<List<String>>> groups() {
        return Optional.ofNullable(this.groups);
    }

    /**
     * (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
     * 
     */
    @Import(name="internalPasswordDisabled")
    private @Nullable Output<Boolean> internalPasswordDisabled;

    /**
     * @return (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
     * 
     */
    public Optional<Output<Boolean>> internalPasswordDisabled() {
        return Optional.ofNullable(this.internalPasswordDisabled);
    }

    /**
     * Username for user. May contain lowercase letters, numbers and symbols: &#39;.-_@&#39;
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Username for user. May contain lowercase letters, numbers and symbols: &#39;.-_@&#39;
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
     * 
     */
    @Import(name="profileUpdatable")
    private @Nullable Output<Boolean> profileUpdatable;

    /**
     * @return (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
     * 
     */
    public Optional<Output<Boolean>> profileUpdatable() {
        return Optional.ofNullable(this.profileUpdatable);
    }

    private UserArgs() {}

    private UserArgs(UserArgs $) {
        this.admin = $.admin;
        this.disableUiAccess = $.disableUiAccess;
        this.email = $.email;
        this.groups = $.groups;
        this.internalPasswordDisabled = $.internalPasswordDisabled;
        this.name = $.name;
        this.password = $.password;
        this.profileUpdatable = $.profileUpdatable;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserArgs $;

        public Builder() {
            $ = new UserArgs();
        }

        public Builder(UserArgs defaults) {
            $ = new UserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param admin (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
         * 
         * @return builder
         * 
         */
        public Builder admin(@Nullable Output<Boolean> admin) {
            $.admin = admin;
            return this;
        }

        /**
         * @param admin (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
         * 
         * @return builder
         * 
         */
        public Builder admin(Boolean admin) {
            return admin(Output.of(admin));
        }

        /**
         * @param disableUiAccess (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
         * 
         * @return builder
         * 
         */
        public Builder disableUiAccess(@Nullable Output<Boolean> disableUiAccess) {
            $.disableUiAccess = disableUiAccess;
            return this;
        }

        /**
         * @param disableUiAccess (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
         * 
         * @return builder
         * 
         */
        public Builder disableUiAccess(Boolean disableUiAccess) {
            return disableUiAccess(Output.of(disableUiAccess));
        }

        /**
         * @param email Email for user.
         * 
         * @return builder
         * 
         */
        public Builder email(Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email Email for user.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param groups List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
         * 
         * @return builder
         * 
         */
        public Builder groups(@Nullable Output<List<String>> groups) {
            $.groups = groups;
            return this;
        }

        /**
         * @param groups List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
         * 
         * @return builder
         * 
         */
        public Builder groups(List<String> groups) {
            return groups(Output.of(groups));
        }

        /**
         * @param groups List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
         * 
         * @return builder
         * 
         */
        public Builder groups(String... groups) {
            return groups(List.of(groups));
        }

        /**
         * @param internalPasswordDisabled (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
         * 
         * @return builder
         * 
         */
        public Builder internalPasswordDisabled(@Nullable Output<Boolean> internalPasswordDisabled) {
            $.internalPasswordDisabled = internalPasswordDisabled;
            return this;
        }

        /**
         * @param internalPasswordDisabled (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
         * 
         * @return builder
         * 
         */
        public Builder internalPasswordDisabled(Boolean internalPasswordDisabled) {
            return internalPasswordDisabled(Output.of(internalPasswordDisabled));
        }

        /**
         * @param name Username for user. May contain lowercase letters, numbers and symbols: &#39;.-_@&#39;
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Username for user. May contain lowercase letters, numbers and symbols: &#39;.-_@&#39;
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param profileUpdatable (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
         * 
         * @return builder
         * 
         */
        public Builder profileUpdatable(@Nullable Output<Boolean> profileUpdatable) {
            $.profileUpdatable = profileUpdatable;
            return this;
        }

        /**
         * @param profileUpdatable (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
         * 
         * @return builder
         * 
         */
        public Builder profileUpdatable(Boolean profileUpdatable) {
            return profileUpdatable(Output.of(profileUpdatable));
        }

        public UserArgs build() {
            if ($.email == null) {
                throw new MissingRequiredPropertyException("UserArgs", "email");
            }
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetUserResult {
    /**
     * @return When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
     * 
     */
    private @Nullable Boolean admin;
    /**
     * @return When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
     * 
     */
    private @Nullable Boolean disableUiAccess;
    /**
     * @return Email for user.
     * 
     */
    private @Nullable String email;
    /**
     * @return List of groups this user is a part of.
     * 
     */
    private @Nullable List<String> groups;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
     * 
     */
    private @Nullable Boolean internalPasswordDisabled;
    private String name;
    /**
     * @return When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
     * 
     */
    private @Nullable Boolean profileUpdatable;

    private GetUserResult() {}
    /**
     * @return When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
     * 
     */
    public Optional<Boolean> admin() {
        return Optional.ofNullable(this.admin);
    }
    /**
     * @return When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
     * 
     */
    public Optional<Boolean> disableUiAccess() {
        return Optional.ofNullable(this.disableUiAccess);
    }
    /**
     * @return Email for user.
     * 
     */
    public Optional<String> email() {
        return Optional.ofNullable(this.email);
    }
    /**
     * @return List of groups this user is a part of.
     * 
     */
    public List<String> groups() {
        return this.groups == null ? List.of() : this.groups;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
     * 
     */
    public Optional<Boolean> internalPasswordDisabled() {
        return Optional.ofNullable(this.internalPasswordDisabled);
    }
    public String name() {
        return this.name;
    }
    /**
     * @return When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
     * 
     */
    public Optional<Boolean> profileUpdatable() {
        return Optional.ofNullable(this.profileUpdatable);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUserResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean admin;
        private @Nullable Boolean disableUiAccess;
        private @Nullable String email;
        private @Nullable List<String> groups;
        private String id;
        private @Nullable Boolean internalPasswordDisabled;
        private String name;
        private @Nullable Boolean profileUpdatable;
        public Builder() {}
        public Builder(GetUserResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.admin = defaults.admin;
    	      this.disableUiAccess = defaults.disableUiAccess;
    	      this.email = defaults.email;
    	      this.groups = defaults.groups;
    	      this.id = defaults.id;
    	      this.internalPasswordDisabled = defaults.internalPasswordDisabled;
    	      this.name = defaults.name;
    	      this.profileUpdatable = defaults.profileUpdatable;
        }

        @CustomType.Setter
        public Builder admin(@Nullable Boolean admin) {
            this.admin = admin;
            return this;
        }
        @CustomType.Setter
        public Builder disableUiAccess(@Nullable Boolean disableUiAccess) {
            this.disableUiAccess = disableUiAccess;
            return this;
        }
        @CustomType.Setter
        public Builder email(@Nullable String email) {
            this.email = email;
            return this;
        }
        @CustomType.Setter
        public Builder groups(@Nullable List<String> groups) {
            this.groups = groups;
            return this;
        }
        public Builder groups(String... groups) {
            return groups(List.of(groups));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder internalPasswordDisabled(@Nullable Boolean internalPasswordDisabled) {
            this.internalPasswordDisabled = internalPasswordDisabled;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder profileUpdatable(@Nullable Boolean profileUpdatable) {
            this.profileUpdatable = profileUpdatable;
            return this;
        }
        public GetUserResult build() {
            final var o = new GetUserResult();
            o.admin = admin;
            o.disableUiAccess = disableUiAccess;
            o.email = email;
            o.groups = groups;
            o.id = id;
            o.internalPasswordDisabled = internalPasswordDisabled;
            o.name = name;
            o.profileUpdatable = profileUpdatable;
            return o;
        }
    }
}

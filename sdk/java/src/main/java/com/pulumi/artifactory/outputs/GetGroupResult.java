// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetGroupResult {
    /**
     * @return Any users added to this group will automatically be assigned with admin privileges in the system.
     * 
     */
    private Boolean adminPrivileges;
    /**
     * @return When this parameter is set, any new users defined in the system are automatically assigned to this group.
     * 
     */
    private Boolean autoJoin;
    /**
     * @return A description for the group
     * 
     */
    private @Nullable String description;
    /**
     * @return New external group ID used to configure the corresponding group in Azure AD.
     * 
     */
    private @Nullable String externalId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String includeUsers;
    private String name;
    /**
     * @return When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
     * 
     */
    private @Nullable Boolean policyManager;
    /**
     * @return The realm for the group.
     * 
     */
    private String realm;
    /**
     * @return The realm attributes for the group.
     * 
     */
    private @Nullable String realmAttributes;
    /**
     * @return When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
     * 
     */
    private @Nullable Boolean reportsManager;
    /**
     * @return List of users assigned to the group. Set include_users to `true` to retrieve this list.
     * 
     */
    private @Nullable List<String> usersNames;
    /**
     * @return When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
     * 
     */
    private @Nullable Boolean watchManager;

    private GetGroupResult() {}
    /**
     * @return Any users added to this group will automatically be assigned with admin privileges in the system.
     * 
     */
    public Boolean adminPrivileges() {
        return this.adminPrivileges;
    }
    /**
     * @return When this parameter is set, any new users defined in the system are automatically assigned to this group.
     * 
     */
    public Boolean autoJoin() {
        return this.autoJoin;
    }
    /**
     * @return A description for the group
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return New external group ID used to configure the corresponding group in Azure AD.
     * 
     */
    public Optional<String> externalId() {
        return Optional.ofNullable(this.externalId);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> includeUsers() {
        return Optional.ofNullable(this.includeUsers);
    }
    public String name() {
        return this.name;
    }
    /**
     * @return When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
     * 
     */
    public Optional<Boolean> policyManager() {
        return Optional.ofNullable(this.policyManager);
    }
    /**
     * @return The realm for the group.
     * 
     */
    public String realm() {
        return this.realm;
    }
    /**
     * @return The realm attributes for the group.
     * 
     */
    public Optional<String> realmAttributes() {
        return Optional.ofNullable(this.realmAttributes);
    }
    /**
     * @return When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
     * 
     */
    public Optional<Boolean> reportsManager() {
        return Optional.ofNullable(this.reportsManager);
    }
    /**
     * @return List of users assigned to the group. Set include_users to `true` to retrieve this list.
     * 
     */
    public List<String> usersNames() {
        return this.usersNames == null ? List.of() : this.usersNames;
    }
    /**
     * @return When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
     * 
     */
    public Optional<Boolean> watchManager() {
        return Optional.ofNullable(this.watchManager);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean adminPrivileges;
        private Boolean autoJoin;
        private @Nullable String description;
        private @Nullable String externalId;
        private String id;
        private @Nullable String includeUsers;
        private String name;
        private @Nullable Boolean policyManager;
        private String realm;
        private @Nullable String realmAttributes;
        private @Nullable Boolean reportsManager;
        private @Nullable List<String> usersNames;
        private @Nullable Boolean watchManager;
        public Builder() {}
        public Builder(GetGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminPrivileges = defaults.adminPrivileges;
    	      this.autoJoin = defaults.autoJoin;
    	      this.description = defaults.description;
    	      this.externalId = defaults.externalId;
    	      this.id = defaults.id;
    	      this.includeUsers = defaults.includeUsers;
    	      this.name = defaults.name;
    	      this.policyManager = defaults.policyManager;
    	      this.realm = defaults.realm;
    	      this.realmAttributes = defaults.realmAttributes;
    	      this.reportsManager = defaults.reportsManager;
    	      this.usersNames = defaults.usersNames;
    	      this.watchManager = defaults.watchManager;
        }

        @CustomType.Setter
        public Builder adminPrivileges(Boolean adminPrivileges) {
            if (adminPrivileges == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "adminPrivileges");
            }
            this.adminPrivileges = adminPrivileges;
            return this;
        }
        @CustomType.Setter
        public Builder autoJoin(Boolean autoJoin) {
            if (autoJoin == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "autoJoin");
            }
            this.autoJoin = autoJoin;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder externalId(@Nullable String externalId) {

            this.externalId = externalId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder includeUsers(@Nullable String includeUsers) {

            this.includeUsers = includeUsers;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder policyManager(@Nullable Boolean policyManager) {

            this.policyManager = policyManager;
            return this;
        }
        @CustomType.Setter
        public Builder realm(String realm) {
            if (realm == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "realm");
            }
            this.realm = realm;
            return this;
        }
        @CustomType.Setter
        public Builder realmAttributes(@Nullable String realmAttributes) {

            this.realmAttributes = realmAttributes;
            return this;
        }
        @CustomType.Setter
        public Builder reportsManager(@Nullable Boolean reportsManager) {

            this.reportsManager = reportsManager;
            return this;
        }
        @CustomType.Setter
        public Builder usersNames(@Nullable List<String> usersNames) {

            this.usersNames = usersNames;
            return this;
        }
        public Builder usersNames(String... usersNames) {
            return usersNames(List.of(usersNames));
        }
        @CustomType.Setter
        public Builder watchManager(@Nullable Boolean watchManager) {

            this.watchManager = watchManager;
            return this;
        }
        public GetGroupResult build() {
            final var _resultValue = new GetGroupResult();
            _resultValue.adminPrivileges = adminPrivileges;
            _resultValue.autoJoin = autoJoin;
            _resultValue.description = description;
            _resultValue.externalId = externalId;
            _resultValue.id = id;
            _resultValue.includeUsers = includeUsers;
            _resultValue.name = name;
            _resultValue.policyManager = policyManager;
            _resultValue.realm = realm;
            _resultValue.realmAttributes = realmAttributes;
            _resultValue.reportsManager = reportsManager;
            _resultValue.usersNames = usersNames;
            _resultValue.watchManager = watchManager;
            return _resultValue;
        }
    }
}

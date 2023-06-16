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
public final class GetGroupResult {
    private Boolean adminPrivileges;
    private Boolean autoJoin;
    private @Nullable String description;
    private @Nullable String externalId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String includeUsers;
    private String name;
    private @Nullable Boolean policyManager;
    private String realm;
    private @Nullable String realmAttributes;
    private @Nullable Boolean reportsManager;
    private @Nullable List<String> usersNames;
    private @Nullable Boolean watchManager;

    private GetGroupResult() {}
    public Boolean adminPrivileges() {
        return this.adminPrivileges;
    }
    public Boolean autoJoin() {
        return this.autoJoin;
    }
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
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
    public Optional<Boolean> policyManager() {
        return Optional.ofNullable(this.policyManager);
    }
    public String realm() {
        return this.realm;
    }
    public Optional<String> realmAttributes() {
        return Optional.ofNullable(this.realmAttributes);
    }
    public Optional<Boolean> reportsManager() {
        return Optional.ofNullable(this.reportsManager);
    }
    public List<String> usersNames() {
        return this.usersNames == null ? List.of() : this.usersNames;
    }
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
            this.adminPrivileges = Objects.requireNonNull(adminPrivileges);
            return this;
        }
        @CustomType.Setter
        public Builder autoJoin(Boolean autoJoin) {
            this.autoJoin = Objects.requireNonNull(autoJoin);
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
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder includeUsers(@Nullable String includeUsers) {
            this.includeUsers = includeUsers;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder policyManager(@Nullable Boolean policyManager) {
            this.policyManager = policyManager;
            return this;
        }
        @CustomType.Setter
        public Builder realm(String realm) {
            this.realm = Objects.requireNonNull(realm);
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
            final var o = new GetGroupResult();
            o.adminPrivileges = adminPrivileges;
            o.autoJoin = autoJoin;
            o.description = description;
            o.externalId = externalId;
            o.id = id;
            o.includeUsers = includeUsers;
            o.name = name;
            o.policyManager = policyManager;
            o.realm = realm;
            o.realmAttributes = realmAttributes;
            o.reportsManager = reportsManager;
            o.usersNames = usersNames;
            o.watchManager = watchManager;
            return o;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetPermissionTargetRepoActionsUser {
    /**
     * @return Name of the permission target.
     * 
     */
    private String name;
    private List<String> permissions;

    private GetPermissionTargetRepoActionsUser() {}
    /**
     * @return Name of the permission target.
     * 
     */
    public String name() {
        return this.name;
    }
    public List<String> permissions() {
        return this.permissions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPermissionTargetRepoActionsUser defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private List<String> permissions;
        public Builder() {}
        public Builder(GetPermissionTargetRepoActionsUser defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.permissions = defaults.permissions;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetPermissionTargetRepoActionsUser", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder permissions(List<String> permissions) {
            if (permissions == null) {
              throw new MissingRequiredPropertyException("GetPermissionTargetRepoActionsUser", "permissions");
            }
            this.permissions = permissions;
            return this;
        }
        public Builder permissions(String... permissions) {
            return permissions(List.of(permissions));
        }
        public GetPermissionTargetRepoActionsUser build() {
            final var _resultValue = new GetPermissionTargetRepoActionsUser();
            _resultValue.name = name;
            _resultValue.permissions = permissions;
            return _resultValue;
        }
    }
}

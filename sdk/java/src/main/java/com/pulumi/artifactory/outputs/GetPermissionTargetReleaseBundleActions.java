// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.artifactory.outputs.GetPermissionTargetReleaseBundleActionsGroup;
import com.pulumi.artifactory.outputs.GetPermissionTargetReleaseBundleActionsUser;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetPermissionTargetReleaseBundleActions {
    /**
     * @return Groups this permission applies for.
     * 
     */
    private @Nullable List<GetPermissionTargetReleaseBundleActionsGroup> groups;
    /**
     * @return Users this permission target applies for.
     * 
     */
    private @Nullable List<GetPermissionTargetReleaseBundleActionsUser> users;

    private GetPermissionTargetReleaseBundleActions() {}
    /**
     * @return Groups this permission applies for.
     * 
     */
    public List<GetPermissionTargetReleaseBundleActionsGroup> groups() {
        return this.groups == null ? List.of() : this.groups;
    }
    /**
     * @return Users this permission target applies for.
     * 
     */
    public List<GetPermissionTargetReleaseBundleActionsUser> users() {
        return this.users == null ? List.of() : this.users;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPermissionTargetReleaseBundleActions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<GetPermissionTargetReleaseBundleActionsGroup> groups;
        private @Nullable List<GetPermissionTargetReleaseBundleActionsUser> users;
        public Builder() {}
        public Builder(GetPermissionTargetReleaseBundleActions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groups = defaults.groups;
    	      this.users = defaults.users;
        }

        @CustomType.Setter
        public Builder groups(@Nullable List<GetPermissionTargetReleaseBundleActionsGroup> groups) {
            this.groups = groups;
            return this;
        }
        public Builder groups(GetPermissionTargetReleaseBundleActionsGroup... groups) {
            return groups(List.of(groups));
        }
        @CustomType.Setter
        public Builder users(@Nullable List<GetPermissionTargetReleaseBundleActionsUser> users) {
            this.users = users;
            return this;
        }
        public Builder users(GetPermissionTargetReleaseBundleActionsUser... users) {
            return users(List.of(users));
        }
        public GetPermissionTargetReleaseBundleActions build() {
            final var _resultValue = new GetPermissionTargetReleaseBundleActions();
            _resultValue.groups = groups;
            _resultValue.users = users;
            return _resultValue;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.artifactory.outputs.GetPermissionTargetRepoActionsGroup;
import com.pulumi.artifactory.outputs.GetPermissionTargetRepoActionsUser;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetPermissionTargetRepoActions {
    /**
     * @return Groups this permission applies for.
     * 
     */
    private @Nullable List<GetPermissionTargetRepoActionsGroup> groups;
    /**
     * @return Users this permission target applies for.
     * 
     */
    private @Nullable List<GetPermissionTargetRepoActionsUser> users;

    private GetPermissionTargetRepoActions() {}
    /**
     * @return Groups this permission applies for.
     * 
     */
    public List<GetPermissionTargetRepoActionsGroup> groups() {
        return this.groups == null ? List.of() : this.groups;
    }
    /**
     * @return Users this permission target applies for.
     * 
     */
    public List<GetPermissionTargetRepoActionsUser> users() {
        return this.users == null ? List.of() : this.users;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPermissionTargetRepoActions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<GetPermissionTargetRepoActionsGroup> groups;
        private @Nullable List<GetPermissionTargetRepoActionsUser> users;
        public Builder() {}
        public Builder(GetPermissionTargetRepoActions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groups = defaults.groups;
    	      this.users = defaults.users;
        }

        @CustomType.Setter
        public Builder groups(@Nullable List<GetPermissionTargetRepoActionsGroup> groups) {

            this.groups = groups;
            return this;
        }
        public Builder groups(GetPermissionTargetRepoActionsGroup... groups) {
            return groups(List.of(groups));
        }
        @CustomType.Setter
        public Builder users(@Nullable List<GetPermissionTargetRepoActionsUser> users) {

            this.users = users;
            return this;
        }
        public Builder users(GetPermissionTargetRepoActionsUser... users) {
            return users(List.of(users));
        }
        public GetPermissionTargetRepoActions build() {
            final var _resultValue = new GetPermissionTargetRepoActions();
            _resultValue.groups = groups;
            _resultValue.users = users;
            return _resultValue;
        }
    }
}

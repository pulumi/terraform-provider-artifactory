// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.artifactory.outputs.PermissionTargetRepoActionsGroup;
import com.pulumi.artifactory.outputs.PermissionTargetRepoActionsUser;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class PermissionTargetRepoActions {
    private @Nullable List<PermissionTargetRepoActionsGroup> groups;
    private @Nullable List<PermissionTargetRepoActionsUser> users;

    private PermissionTargetRepoActions() {}
    public List<PermissionTargetRepoActionsGroup> groups() {
        return this.groups == null ? List.of() : this.groups;
    }
    public List<PermissionTargetRepoActionsUser> users() {
        return this.users == null ? List.of() : this.users;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PermissionTargetRepoActions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<PermissionTargetRepoActionsGroup> groups;
        private @Nullable List<PermissionTargetRepoActionsUser> users;
        public Builder() {}
        public Builder(PermissionTargetRepoActions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groups = defaults.groups;
    	      this.users = defaults.users;
        }

        @CustomType.Setter
        public Builder groups(@Nullable List<PermissionTargetRepoActionsGroup> groups) {

            this.groups = groups;
            return this;
        }
        public Builder groups(PermissionTargetRepoActionsGroup... groups) {
            return groups(List.of(groups));
        }
        @CustomType.Setter
        public Builder users(@Nullable List<PermissionTargetRepoActionsUser> users) {

            this.users = users;
            return this;
        }
        public Builder users(PermissionTargetRepoActionsUser... users) {
            return users(List.of(users));
        }
        public PermissionTargetRepoActions build() {
            final var _resultValue = new PermissionTargetRepoActions();
            _resultValue.groups = groups;
            _resultValue.users = users;
            return _resultValue;
        }
    }
}

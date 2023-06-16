// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetPermissionTargetRepoActionsGroup extends com.pulumi.resources.InvokeArgs {

    public static final GetPermissionTargetRepoActionsGroup Empty = new GetPermissionTargetRepoActionsGroup();

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    @Import(name="permissions", required=true)
    private List<String> permissions;

    public List<String> permissions() {
        return this.permissions;
    }

    private GetPermissionTargetRepoActionsGroup() {}

    private GetPermissionTargetRepoActionsGroup(GetPermissionTargetRepoActionsGroup $) {
        this.name = $.name;
        this.permissions = $.permissions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPermissionTargetRepoActionsGroup defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPermissionTargetRepoActionsGroup $;

        public Builder() {
            $ = new GetPermissionTargetRepoActionsGroup();
        }

        public Builder(GetPermissionTargetRepoActionsGroup defaults) {
            $ = new GetPermissionTargetRepoActionsGroup(Objects.requireNonNull(defaults));
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public Builder permissions(List<String> permissions) {
            $.permissions = permissions;
            return this;
        }

        public Builder permissions(String... permissions) {
            return permissions(List.of(permissions));
        }

        public GetPermissionTargetRepoActionsGroup build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.permissions = Objects.requireNonNull($.permissions, "expected parameter 'permissions' to be non-null");
            return $;
        }
    }

}

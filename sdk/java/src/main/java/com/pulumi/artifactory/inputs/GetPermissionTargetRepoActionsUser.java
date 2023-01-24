// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetPermissionTargetRepoActionsUser extends com.pulumi.resources.InvokeArgs {

    public static final GetPermissionTargetRepoActionsUser Empty = new GetPermissionTargetRepoActionsUser();

    /**
     * Name of the permission target.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Name of the permission target.
     * 
     */
    public String name() {
        return this.name;
    }

    @Import(name="permissions", required=true)
    private List<String> permissions;

    public List<String> permissions() {
        return this.permissions;
    }

    private GetPermissionTargetRepoActionsUser() {}

    private GetPermissionTargetRepoActionsUser(GetPermissionTargetRepoActionsUser $) {
        this.name = $.name;
        this.permissions = $.permissions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPermissionTargetRepoActionsUser defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPermissionTargetRepoActionsUser $;

        public Builder() {
            $ = new GetPermissionTargetRepoActionsUser();
        }

        public Builder(GetPermissionTargetRepoActionsUser defaults) {
            $ = new GetPermissionTargetRepoActionsUser(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the permission target.
         * 
         * @return builder
         * 
         */
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

        public GetPermissionTargetRepoActionsUser build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.permissions = Objects.requireNonNull($.permissions, "expected parameter 'permissions' to be non-null");
            return $;
        }
    }

}

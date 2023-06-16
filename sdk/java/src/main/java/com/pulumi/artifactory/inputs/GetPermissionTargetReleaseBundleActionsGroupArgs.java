// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetPermissionTargetReleaseBundleActionsGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetPermissionTargetReleaseBundleActionsGroupArgs Empty = new GetPermissionTargetReleaseBundleActionsGroupArgs();

    @Import(name="name", required=true)
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }

    @Import(name="permissions", required=true)
    private Output<List<String>> permissions;

    public Output<List<String>> permissions() {
        return this.permissions;
    }

    private GetPermissionTargetReleaseBundleActionsGroupArgs() {}

    private GetPermissionTargetReleaseBundleActionsGroupArgs(GetPermissionTargetReleaseBundleActionsGroupArgs $) {
        this.name = $.name;
        this.permissions = $.permissions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPermissionTargetReleaseBundleActionsGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPermissionTargetReleaseBundleActionsGroupArgs $;

        public Builder() {
            $ = new GetPermissionTargetReleaseBundleActionsGroupArgs();
        }

        public Builder(GetPermissionTargetReleaseBundleActionsGroupArgs defaults) {
            $ = new GetPermissionTargetReleaseBundleActionsGroupArgs(Objects.requireNonNull(defaults));
        }

        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder permissions(Output<List<String>> permissions) {
            $.permissions = permissions;
            return this;
        }

        public Builder permissions(List<String> permissions) {
            return permissions(Output.of(permissions));
        }

        public Builder permissions(String... permissions) {
            return permissions(List.of(permissions));
        }

        public GetPermissionTargetReleaseBundleActionsGroupArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.permissions = Objects.requireNonNull($.permissions, "expected parameter 'permissions' to be non-null");
            return $;
        }
    }

}

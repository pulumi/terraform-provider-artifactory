// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.GetPermissionTargetBuildActionsGroupArgs;
import com.pulumi.artifactory.inputs.GetPermissionTargetBuildActionsUserArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPermissionTargetBuildActionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetPermissionTargetBuildActionsArgs Empty = new GetPermissionTargetBuildActionsArgs();

    @Import(name="groups")
    private @Nullable Output<List<GetPermissionTargetBuildActionsGroupArgs>> groups;

    public Optional<Output<List<GetPermissionTargetBuildActionsGroupArgs>>> groups() {
        return Optional.ofNullable(this.groups);
    }

    @Import(name="users")
    private @Nullable Output<List<GetPermissionTargetBuildActionsUserArgs>> users;

    public Optional<Output<List<GetPermissionTargetBuildActionsUserArgs>>> users() {
        return Optional.ofNullable(this.users);
    }

    private GetPermissionTargetBuildActionsArgs() {}

    private GetPermissionTargetBuildActionsArgs(GetPermissionTargetBuildActionsArgs $) {
        this.groups = $.groups;
        this.users = $.users;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPermissionTargetBuildActionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPermissionTargetBuildActionsArgs $;

        public Builder() {
            $ = new GetPermissionTargetBuildActionsArgs();
        }

        public Builder(GetPermissionTargetBuildActionsArgs defaults) {
            $ = new GetPermissionTargetBuildActionsArgs(Objects.requireNonNull(defaults));
        }

        public Builder groups(@Nullable Output<List<GetPermissionTargetBuildActionsGroupArgs>> groups) {
            $.groups = groups;
            return this;
        }

        public Builder groups(List<GetPermissionTargetBuildActionsGroupArgs> groups) {
            return groups(Output.of(groups));
        }

        public Builder groups(GetPermissionTargetBuildActionsGroupArgs... groups) {
            return groups(List.of(groups));
        }

        public Builder users(@Nullable Output<List<GetPermissionTargetBuildActionsUserArgs>> users) {
            $.users = users;
            return this;
        }

        public Builder users(List<GetPermissionTargetBuildActionsUserArgs> users) {
            return users(Output.of(users));
        }

        public Builder users(GetPermissionTargetBuildActionsUserArgs... users) {
            return users(List.of(users));
        }

        public GetPermissionTargetBuildActionsArgs build() {
            return $;
        }
    }

}

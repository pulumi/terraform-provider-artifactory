// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.PermissionTargetBuildActionsGroupArgs;
import com.pulumi.artifactory.inputs.PermissionTargetBuildActionsUserArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PermissionTargetBuildActionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final PermissionTargetBuildActionsArgs Empty = new PermissionTargetBuildActionsArgs();

    /**
     * Groups this permission applies for.
     * 
     */
    @Import(name="groups")
    private @Nullable Output<List<PermissionTargetBuildActionsGroupArgs>> groups;

    /**
     * @return Groups this permission applies for.
     * 
     */
    public Optional<Output<List<PermissionTargetBuildActionsGroupArgs>>> groups() {
        return Optional.ofNullable(this.groups);
    }

    /**
     * Users this permission target applies for.
     * 
     */
    @Import(name="users")
    private @Nullable Output<List<PermissionTargetBuildActionsUserArgs>> users;

    /**
     * @return Users this permission target applies for.
     * 
     */
    public Optional<Output<List<PermissionTargetBuildActionsUserArgs>>> users() {
        return Optional.ofNullable(this.users);
    }

    private PermissionTargetBuildActionsArgs() {}

    private PermissionTargetBuildActionsArgs(PermissionTargetBuildActionsArgs $) {
        this.groups = $.groups;
        this.users = $.users;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PermissionTargetBuildActionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PermissionTargetBuildActionsArgs $;

        public Builder() {
            $ = new PermissionTargetBuildActionsArgs();
        }

        public Builder(PermissionTargetBuildActionsArgs defaults) {
            $ = new PermissionTargetBuildActionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groups Groups this permission applies for.
         * 
         * @return builder
         * 
         */
        public Builder groups(@Nullable Output<List<PermissionTargetBuildActionsGroupArgs>> groups) {
            $.groups = groups;
            return this;
        }

        /**
         * @param groups Groups this permission applies for.
         * 
         * @return builder
         * 
         */
        public Builder groups(List<PermissionTargetBuildActionsGroupArgs> groups) {
            return groups(Output.of(groups));
        }

        /**
         * @param groups Groups this permission applies for.
         * 
         * @return builder
         * 
         */
        public Builder groups(PermissionTargetBuildActionsGroupArgs... groups) {
            return groups(List.of(groups));
        }

        /**
         * @param users Users this permission target applies for.
         * 
         * @return builder
         * 
         */
        public Builder users(@Nullable Output<List<PermissionTargetBuildActionsUserArgs>> users) {
            $.users = users;
            return this;
        }

        /**
         * @param users Users this permission target applies for.
         * 
         * @return builder
         * 
         */
        public Builder users(List<PermissionTargetBuildActionsUserArgs> users) {
            return users(Output.of(users));
        }

        /**
         * @param users Users this permission target applies for.
         * 
         * @return builder
         * 
         */
        public Builder users(PermissionTargetBuildActionsUserArgs... users) {
            return users(List.of(users));
        }

        public PermissionTargetBuildActionsArgs build() {
            return $;
        }
    }

}

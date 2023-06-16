// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.GetPermissionTargetBuild;
import com.pulumi.artifactory.inputs.GetPermissionTargetReleaseBundle;
import com.pulumi.artifactory.inputs.GetPermissionTargetRepo;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPermissionTargetPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPermissionTargetPlainArgs Empty = new GetPermissionTargetPlainArgs();

    @Import(name="build")
    private @Nullable GetPermissionTargetBuild build;

    public Optional<GetPermissionTargetBuild> build() {
        return Optional.ofNullable(this.build);
    }

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    @Import(name="releaseBundle")
    private @Nullable GetPermissionTargetReleaseBundle releaseBundle;

    public Optional<GetPermissionTargetReleaseBundle> releaseBundle() {
        return Optional.ofNullable(this.releaseBundle);
    }

    @Import(name="repo")
    private @Nullable GetPermissionTargetRepo repo;

    public Optional<GetPermissionTargetRepo> repo() {
        return Optional.ofNullable(this.repo);
    }

    private GetPermissionTargetPlainArgs() {}

    private GetPermissionTargetPlainArgs(GetPermissionTargetPlainArgs $) {
        this.build = $.build;
        this.name = $.name;
        this.releaseBundle = $.releaseBundle;
        this.repo = $.repo;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPermissionTargetPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPermissionTargetPlainArgs $;

        public Builder() {
            $ = new GetPermissionTargetPlainArgs();
        }

        public Builder(GetPermissionTargetPlainArgs defaults) {
            $ = new GetPermissionTargetPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder build(@Nullable GetPermissionTargetBuild build) {
            $.build = build;
            return this;
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public Builder releaseBundle(@Nullable GetPermissionTargetReleaseBundle releaseBundle) {
            $.releaseBundle = releaseBundle;
            return this;
        }

        public Builder repo(@Nullable GetPermissionTargetRepo repo) {
            $.repo = repo;
            return this;
        }

        public GetPermissionTargetPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}

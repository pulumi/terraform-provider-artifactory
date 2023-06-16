// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class PermissionTargetBuildActionsGroup {
    private String name;
    private List<String> permissions;

    private PermissionTargetBuildActionsGroup() {}
    public String name() {
        return this.name;
    }
    public List<String> permissions() {
        return this.permissions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PermissionTargetBuildActionsGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private List<String> permissions;
        public Builder() {}
        public Builder(PermissionTargetBuildActionsGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.permissions = defaults.permissions;
        }

        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder permissions(List<String> permissions) {
            this.permissions = Objects.requireNonNull(permissions);
            return this;
        }
        public Builder permissions(String... permissions) {
            return permissions(List.of(permissions));
        }
        public PermissionTargetBuildActionsGroup build() {
            final var o = new PermissionTargetBuildActionsGroup();
            o.name = name;
            o.permissions = permissions;
            return o;
        }
    }
}

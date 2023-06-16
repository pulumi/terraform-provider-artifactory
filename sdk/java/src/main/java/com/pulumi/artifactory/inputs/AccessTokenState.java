// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.AccessTokenAdminTokenArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccessTokenState extends com.pulumi.resources.ResourceArgs {

    public static final AccessTokenState Empty = new AccessTokenState();

    @Import(name="accessToken")
    private @Nullable Output<String> accessToken;

    public Optional<Output<String>> accessToken() {
        return Optional.ofNullable(this.accessToken);
    }

    @Import(name="adminToken")
    private @Nullable Output<AccessTokenAdminTokenArgs> adminToken;

    public Optional<Output<AccessTokenAdminTokenArgs>> adminToken() {
        return Optional.ofNullable(this.adminToken);
    }

    @Import(name="audience")
    private @Nullable Output<String> audience;

    public Optional<Output<String>> audience() {
        return Optional.ofNullable(this.audience);
    }

    @Import(name="endDate")
    private @Nullable Output<String> endDate;

    public Optional<Output<String>> endDate() {
        return Optional.ofNullable(this.endDate);
    }

    @Import(name="endDateRelative")
    private @Nullable Output<String> endDateRelative;

    public Optional<Output<String>> endDateRelative() {
        return Optional.ofNullable(this.endDateRelative);
    }

    @Import(name="groups")
    private @Nullable Output<List<String>> groups;

    public Optional<Output<List<String>>> groups() {
        return Optional.ofNullable(this.groups);
    }

    @Import(name="refreshToken")
    private @Nullable Output<String> refreshToken;

    public Optional<Output<String>> refreshToken() {
        return Optional.ofNullable(this.refreshToken);
    }

    @Import(name="refreshable")
    private @Nullable Output<Boolean> refreshable;

    public Optional<Output<Boolean>> refreshable() {
        return Optional.ofNullable(this.refreshable);
    }

    @Import(name="username")
    private @Nullable Output<String> username;

    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private AccessTokenState() {}

    private AccessTokenState(AccessTokenState $) {
        this.accessToken = $.accessToken;
        this.adminToken = $.adminToken;
        this.audience = $.audience;
        this.endDate = $.endDate;
        this.endDateRelative = $.endDateRelative;
        this.groups = $.groups;
        this.refreshToken = $.refreshToken;
        this.refreshable = $.refreshable;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessTokenState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessTokenState $;

        public Builder() {
            $ = new AccessTokenState();
        }

        public Builder(AccessTokenState defaults) {
            $ = new AccessTokenState(Objects.requireNonNull(defaults));
        }

        public Builder accessToken(@Nullable Output<String> accessToken) {
            $.accessToken = accessToken;
            return this;
        }

        public Builder accessToken(String accessToken) {
            return accessToken(Output.of(accessToken));
        }

        public Builder adminToken(@Nullable Output<AccessTokenAdminTokenArgs> adminToken) {
            $.adminToken = adminToken;
            return this;
        }

        public Builder adminToken(AccessTokenAdminTokenArgs adminToken) {
            return adminToken(Output.of(adminToken));
        }

        public Builder audience(@Nullable Output<String> audience) {
            $.audience = audience;
            return this;
        }

        public Builder audience(String audience) {
            return audience(Output.of(audience));
        }

        public Builder endDate(@Nullable Output<String> endDate) {
            $.endDate = endDate;
            return this;
        }

        public Builder endDate(String endDate) {
            return endDate(Output.of(endDate));
        }

        public Builder endDateRelative(@Nullable Output<String> endDateRelative) {
            $.endDateRelative = endDateRelative;
            return this;
        }

        public Builder endDateRelative(String endDateRelative) {
            return endDateRelative(Output.of(endDateRelative));
        }

        public Builder groups(@Nullable Output<List<String>> groups) {
            $.groups = groups;
            return this;
        }

        public Builder groups(List<String> groups) {
            return groups(Output.of(groups));
        }

        public Builder groups(String... groups) {
            return groups(List.of(groups));
        }

        public Builder refreshToken(@Nullable Output<String> refreshToken) {
            $.refreshToken = refreshToken;
            return this;
        }

        public Builder refreshToken(String refreshToken) {
            return refreshToken(Output.of(refreshToken));
        }

        public Builder refreshable(@Nullable Output<Boolean> refreshable) {
            $.refreshable = refreshable;
            return this;
        }

        public Builder refreshable(Boolean refreshable) {
            return refreshable(Output.of(refreshable));
        }

        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        public Builder username(String username) {
            return username(Output.of(username));
        }

        public AccessTokenState build() {
            return $;
        }
    }

}

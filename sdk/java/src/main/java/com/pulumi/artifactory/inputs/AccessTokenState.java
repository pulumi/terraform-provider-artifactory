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

    /**
     * Returns the access token to authenciate to Artifactory
     * 
     */
    @Import(name="accessToken")
    private @Nullable Output<String> accessToken;

    /**
     * @return Returns the access token to authenciate to Artifactory
     * 
     */
    public Optional<Output<String>> accessToken() {
        return Optional.ofNullable(this.accessToken);
    }

    /**
     * (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
     * 
     */
    @Import(name="adminToken")
    private @Nullable Output<AccessTokenAdminTokenArgs> adminToken;

    /**
     * @return (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
     * 
     */
    public Optional<Output<AccessTokenAdminTokenArgs>> adminToken() {
        return Optional.ofNullable(this.adminToken);
    }

    /**
     * (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `&#34;jfrt{@literal @}*&#34;` so the token to be accepted by all Artifactory instances.
     * 
     */
    @Import(name="audience")
    private @Nullable Output<String> audience;

    /**
     * @return (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `&#34;jfrt{@literal @}*&#34;` so the token to be accepted by all Artifactory instances.
     * 
     */
    public Optional<Output<String>> audience() {
        return Optional.ofNullable(this.audience);
    }

    /**
     * (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * 
     */
    @Import(name="endDate")
    private @Nullable Output<String> endDate;

    /**
     * @return (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * 
     */
    public Optional<Output<String>> endDate() {
        return Optional.ofNullable(this.endDate);
    }

    /**
     * (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are &#34;s&#34;, &#34;m&#34;, &#34;h&#34;.
     * 
     */
    @Import(name="endDateRelative")
    private @Nullable Output<String> endDateRelative;

    /**
     * @return (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are &#34;s&#34;, &#34;m&#34;, &#34;h&#34;.
     * 
     */
    public Optional<Output<String>> endDateRelative() {
        return Optional.ofNullable(this.endDateRelative);
    }

    /**
     * (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `[&#34;*&#34;]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
     * 
     */
    @Import(name="groups")
    private @Nullable Output<List<String>> groups;

    /**
     * @return (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `[&#34;*&#34;]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
     * 
     */
    public Optional<Output<List<String>>> groups() {
        return Optional.ofNullable(this.groups);
    }

    /**
     * Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
     * 
     */
    @Import(name="refreshToken")
    private @Nullable Output<String> refreshToken;

    /**
     * @return Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
     * 
     */
    public Optional<Output<String>> refreshToken() {
        return Optional.ofNullable(this.refreshToken);
    }

    /**
     * (Optional) Is this token refreshable? Defaults to `false`
     * 
     */
    @Import(name="refreshable")
    private @Nullable Output<Boolean> refreshable;

    /**
     * @return (Optional) Is this token refreshable? Defaults to `false`
     * 
     */
    public Optional<Output<Boolean>> refreshable() {
        return Optional.ofNullable(this.refreshable);
    }

    /**
     * (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
     * 
     */
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

        /**
         * @param accessToken Returns the access token to authenciate to Artifactory
         * 
         * @return builder
         * 
         */
        public Builder accessToken(@Nullable Output<String> accessToken) {
            $.accessToken = accessToken;
            return this;
        }

        /**
         * @param accessToken Returns the access token to authenciate to Artifactory
         * 
         * @return builder
         * 
         */
        public Builder accessToken(String accessToken) {
            return accessToken(Output.of(accessToken));
        }

        /**
         * @param adminToken (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
         * 
         * @return builder
         * 
         */
        public Builder adminToken(@Nullable Output<AccessTokenAdminTokenArgs> adminToken) {
            $.adminToken = adminToken;
            return this;
        }

        /**
         * @param adminToken (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
         * 
         * @return builder
         * 
         */
        public Builder adminToken(AccessTokenAdminTokenArgs adminToken) {
            return adminToken(Output.of(adminToken));
        }

        /**
         * @param audience (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `&#34;jfrt{@literal @}*&#34;` so the token to be accepted by all Artifactory instances.
         * 
         * @return builder
         * 
         */
        public Builder audience(@Nullable Output<String> audience) {
            $.audience = audience;
            return this;
        }

        /**
         * @param audience (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `&#34;jfrt{@literal @}*&#34;` so the token to be accepted by all Artifactory instances.
         * 
         * @return builder
         * 
         */
        public Builder audience(String audience) {
            return audience(Output.of(audience));
        }

        /**
         * @param endDate (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
         * 
         * @return builder
         * 
         */
        public Builder endDate(@Nullable Output<String> endDate) {
            $.endDate = endDate;
            return this;
        }

        /**
         * @param endDate (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
         * 
         * @return builder
         * 
         */
        public Builder endDate(String endDate) {
            return endDate(Output.of(endDate));
        }

        /**
         * @param endDateRelative (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are &#34;s&#34;, &#34;m&#34;, &#34;h&#34;.
         * 
         * @return builder
         * 
         */
        public Builder endDateRelative(@Nullable Output<String> endDateRelative) {
            $.endDateRelative = endDateRelative;
            return this;
        }

        /**
         * @param endDateRelative (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are &#34;s&#34;, &#34;m&#34;, &#34;h&#34;.
         * 
         * @return builder
         * 
         */
        public Builder endDateRelative(String endDateRelative) {
            return endDateRelative(Output.of(endDateRelative));
        }

        /**
         * @param groups (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `[&#34;*&#34;]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
         * 
         * @return builder
         * 
         */
        public Builder groups(@Nullable Output<List<String>> groups) {
            $.groups = groups;
            return this;
        }

        /**
         * @param groups (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `[&#34;*&#34;]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
         * 
         * @return builder
         * 
         */
        public Builder groups(List<String> groups) {
            return groups(Output.of(groups));
        }

        /**
         * @param groups (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `[&#34;*&#34;]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
         * 
         * @return builder
         * 
         */
        public Builder groups(String... groups) {
            return groups(List.of(groups));
        }

        /**
         * @param refreshToken Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
         * 
         * @return builder
         * 
         */
        public Builder refreshToken(@Nullable Output<String> refreshToken) {
            $.refreshToken = refreshToken;
            return this;
        }

        /**
         * @param refreshToken Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
         * 
         * @return builder
         * 
         */
        public Builder refreshToken(String refreshToken) {
            return refreshToken(Output.of(refreshToken));
        }

        /**
         * @param refreshable (Optional) Is this token refreshable? Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder refreshable(@Nullable Output<Boolean> refreshable) {
            $.refreshable = refreshable;
            return this;
        }

        /**
         * @param refreshable (Optional) Is this token refreshable? Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder refreshable(Boolean refreshable) {
            return refreshable(Output.of(refreshable));
        }

        /**
         * @param username (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public AccessTokenState build() {
            return $;
        }
    }

}

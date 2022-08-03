// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.AccessTokenArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.AccessTokenState;
import com.pulumi.artifactory.outputs.AccessTokenAdminToken;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Artifactory **does not** retain access tokens and cannot be imported into state.
 * 
 */
@ResourceType(type="artifactory:index/accessToken:AccessToken")
public class AccessToken extends com.pulumi.resources.CustomResource {
    /**
     * Returns the access token to authenciate to Artifactory
     * 
     */
    @Export(name="accessToken", type=String.class, parameters={})
    private Output<String> accessToken;

    /**
     * @return Returns the access token to authenciate to Artifactory
     * 
     */
    public Output<String> accessToken() {
        return this.accessToken;
    }
    /**
     * (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
     * 
     */
    @Export(name="adminToken", type=AccessTokenAdminToken.class, parameters={})
    private Output</* @Nullable */ AccessTokenAdminToken> adminToken;

    /**
     * @return (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
     * 
     */
    public Output<Optional<AccessTokenAdminToken>> adminToken() {
        return Codegen.optional(this.adminToken);
    }
    /**
     * (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `&#34;jfrt@*&#34;` so the token to be accepted by all Artifactory instances.
     * 
     */
    @Export(name="audience", type=String.class, parameters={})
    private Output</* @Nullable */ String> audience;

    /**
     * @return (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `&#34;jfrt@*&#34;` so the token to be accepted by all Artifactory instances.
     * 
     */
    public Output<Optional<String>> audience() {
        return Codegen.optional(this.audience);
    }
    /**
     * (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * 
     */
    @Export(name="endDate", type=String.class, parameters={})
    private Output<String> endDate;

    /**
     * @return (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * 
     */
    public Output<String> endDate() {
        return this.endDate;
    }
    /**
     * (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are &#34;s&#34;, &#34;m&#34;, &#34;h&#34;.
     * 
     */
    @Export(name="endDateRelative", type=String.class, parameters={})
    private Output</* @Nullable */ String> endDateRelative;

    /**
     * @return (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are &#34;s&#34;, &#34;m&#34;, &#34;h&#34;.
     * 
     */
    public Output<Optional<String>> endDateRelative() {
        return Codegen.optional(this.endDateRelative);
    }
    /**
     * (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `[&#34;*&#34;]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
     * 
     */
    @Export(name="groups", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> groups;

    /**
     * @return (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `[&#34;*&#34;]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
     * 
     */
    public Output<Optional<List<String>>> groups() {
        return Codegen.optional(this.groups);
    }
    /**
     * Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
     * 
     */
    @Export(name="refreshToken", type=String.class, parameters={})
    private Output<String> refreshToken;

    /**
     * @return Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
     * 
     */
    public Output<String> refreshToken() {
        return this.refreshToken;
    }
    /**
     * (Optional) Is this token refreshable? Defaults to `false`
     * 
     */
    @Export(name="refreshable", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> refreshable;

    /**
     * @return (Optional) Is this token refreshable? Defaults to `false`
     * 
     */
    public Output<Optional<Boolean>> refreshable() {
        return Codegen.optional(this.refreshable);
    }
    /**
     * (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
     * 
     */
    @Export(name="username", type=String.class, parameters={})
    private Output<String> username;

    /**
     * @return (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessToken(String name) {
        this(name, AccessTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessToken(String name, AccessTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessToken(String name, AccessTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/accessToken:AccessToken", name, args == null ? AccessTokenArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccessToken(String name, Output<String> id, @Nullable AccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/accessToken:AccessToken", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AccessToken get(String name, Output<String> id, @Nullable AccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessToken(name, id, state, options);
    }
}

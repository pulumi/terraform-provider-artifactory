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

@ResourceType(type="artifactory:index/accessToken:AccessToken")
public class AccessToken extends com.pulumi.resources.CustomResource {
    @Export(name="accessToken", type=String.class, parameters={})
    private Output<String> accessToken;

    public Output<String> accessToken() {
        return this.accessToken;
    }
    @Export(name="adminToken", type=AccessTokenAdminToken.class, parameters={})
    private Output</* @Nullable */ AccessTokenAdminToken> adminToken;

    public Output<Optional<AccessTokenAdminToken>> adminToken() {
        return Codegen.optional(this.adminToken);
    }
    @Export(name="audience", type=String.class, parameters={})
    private Output</* @Nullable */ String> audience;

    public Output<Optional<String>> audience() {
        return Codegen.optional(this.audience);
    }
    @Export(name="endDate", type=String.class, parameters={})
    private Output<String> endDate;

    public Output<String> endDate() {
        return this.endDate;
    }
    @Export(name="endDateRelative", type=String.class, parameters={})
    private Output</* @Nullable */ String> endDateRelative;

    public Output<Optional<String>> endDateRelative() {
        return Codegen.optional(this.endDateRelative);
    }
    @Export(name="groups", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> groups;

    public Output<Optional<List<String>>> groups() {
        return Codegen.optional(this.groups);
    }
    @Export(name="refreshToken", type=String.class, parameters={})
    private Output<String> refreshToken;

    public Output<String> refreshToken() {
        return this.refreshToken;
    }
    @Export(name="refreshable", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> refreshable;

    public Output<Optional<Boolean>> refreshable() {
        return Codegen.optional(this.refreshable);
    }
    @Export(name="username", type=String.class, parameters={})
    private Output<String> username;

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
            .additionalSecretOutputs(List.of(
                "accessToken",
                "refreshToken"
            ))
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

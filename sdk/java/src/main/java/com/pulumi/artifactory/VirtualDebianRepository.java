// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.VirtualDebianRepositoryArgs;
import com.pulumi.artifactory.inputs.VirtualDebianRepositoryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="artifactory:index/virtualDebianRepository:VirtualDebianRepository")
public class VirtualDebianRepository extends com.pulumi.resources.CustomResource {
    /**
     * Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    @Export(name="artifactoryRequestsCanRetrieveRemoteArtifacts", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts;

    /**
     * @return Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    public Output<Optional<Boolean>> artifactoryRequestsCanRetrieveRemoteArtifacts() {
        return Codegen.optional(this.artifactoryRequestsCanRetrieveRemoteArtifacts);
    }
    /**
     * Specifying architectures will speed up Artifactory&#39;s initial metadata indexing process. The default architecture values
     * are amd64 and i386.
     * 
     */
    @Export(name="debianDefaultArchitectures", type=String.class, parameters={})
    private Output</* @Nullable */ String> debianDefaultArchitectures;

    /**
     * @return Specifying architectures will speed up Artifactory&#39;s initial metadata indexing process. The default architecture values
     * are amd64 and i386.
     * 
     */
    public Output<Optional<String>> debianDefaultArchitectures() {
        return Codegen.optional(this.debianDefaultArchitectures);
    }
    /**
     * Default repository to deploy artifacts.
     * 
     */
    @Export(name="defaultDeploymentRepo", type=String.class, parameters={})
    private Output</* @Nullable */ String> defaultDeploymentRepo;

    /**
     * @return Default repository to deploy artifacts.
     * 
     */
    public Output<Optional<String>> defaultDeploymentRepo() {
        return Codegen.optional(this.defaultDeploymentRepo);
    }
    /**
     * Public description.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Public description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
     * artifacts are excluded.
     * 
     */
    @Export(name="excludesPattern", type=String.class, parameters={})
    private Output</* @Nullable */ String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
     * artifacts are excluded.
     * 
     */
    public Output<Optional<String>> excludesPattern() {
        return Codegen.optional(this.excludesPattern);
    }
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Export(name="includesPattern", type=String.class, parameters={})
    private Output</* @Nullable */ String> includesPattern;

    /**
     * @return List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    public Output<Optional<String>> includesPattern() {
        return Codegen.optional(this.includesPattern);
    }
    /**
     * A mandatory identifier for the repository that must be unique. Must be 3 - 10 lowercase alphanumeric and hyphen
     * characters. It cannot begin with a number or contain spaces or special characters.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return A mandatory identifier for the repository that must be unique. Must be 3 - 10 lowercase alphanumeric and hyphen
     * characters. It cannot begin with a number or contain spaces or special characters.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * Internal description.
     * 
     */
    @Export(name="notes", type=String.class, parameters={})
    private Output</* @Nullable */ String> notes;

    /**
     * @return Internal description.
     * 
     */
    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    /**
     * Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are
     * &#39;bz2&#39;,&#39;lzma&#39; and &#39;xz&#39;. Default value is &#39;bz2&#39;.
     * 
     */
    @Export(name="optionalIndexCompressionFormats", type=List.class, parameters={String.class})
    private Output<List<String>> optionalIndexCompressionFormats;

    /**
     * @return Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are
     * &#39;bz2&#39;,&#39;lzma&#39; and &#39;xz&#39;. Default value is &#39;bz2&#39;.
     * 
     */
    public Output<List<String>> optionalIndexCompressionFormats() {
        return this.optionalIndexCompressionFormats;
    }
    @Export(name="packageType", type=String.class, parameters={})
    private Output<String> packageType;

    public Output<String> packageType() {
        return this.packageType;
    }
    /**
     * Primary keypair used to sign artifacts. Default is empty.
     * 
     */
    @Export(name="primaryKeypairRef", type=String.class, parameters={})
    private Output</* @Nullable */ String> primaryKeypairRef;

    /**
     * @return Primary keypair used to sign artifacts. Default is empty.
     * 
     */
    public Output<Optional<String>> primaryKeypairRef() {
        return Codegen.optional(this.primaryKeypairRef);
    }
    /**
     * Project environment for assigning this repository to. Allow values: &#34;DEV&#34;, &#34;PROD&#34;, or one of custom environment. Before
     * Artifactory 7.53.1, up to 2 values (&#34;DEV&#34; and &#34;PROD&#34;) are allowed. From 7.53.1 onward, only one value is allowed. The
     * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
     * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
     * 
     */
    @Export(name="projectEnvironments", type=List.class, parameters={String.class})
    private Output<List<String>> projectEnvironments;

    /**
     * @return Project environment for assigning this repository to. Allow values: &#34;DEV&#34;, &#34;PROD&#34;, or one of custom environment. Before
     * Artifactory 7.53.1, up to 2 values (&#34;DEV&#34; and &#34;PROD&#34;) are allowed. From 7.53.1 onward, only one value is allowed. The
     * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
     * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
     * 
     */
    public Output<List<String>> projectEnvironments() {
        return this.projectEnvironments;
    }
    /**
     * Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    @Export(name="projectKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    public Output<Optional<String>> projectKey() {
        return Codegen.optional(this.projectKey);
    }
    /**
     * Repository layout key for the local repository
     * 
     */
    @Export(name="repoLayoutRef", type=String.class, parameters={})
    private Output</* @Nullable */ String> repoLayoutRef;

    /**
     * @return Repository layout key for the local repository
     * 
     */
    public Output<Optional<String>> repoLayoutRef() {
        return Codegen.optional(this.repoLayoutRef);
    }
    /**
     * The effective list of actual repositories included in this virtual repository.
     * 
     */
    @Export(name="repositories", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> repositories;

    /**
     * @return The effective list of actual repositories included in this virtual repository.
     * 
     */
    public Output<Optional<List<String>>> repositories() {
        return Codegen.optional(this.repositories);
    }
    /**
     * This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
     * repositories. A value of 0 indicates no caching.
     * 
     */
    @Export(name="retrievalCachePeriodSeconds", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> retrievalCachePeriodSeconds;

    /**
     * @return This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
     * repositories. A value of 0 indicates no caching.
     * 
     */
    public Output<Optional<Integer>> retrievalCachePeriodSeconds() {
        return Codegen.optional(this.retrievalCachePeriodSeconds);
    }
    /**
     * Secondary keypair used to sign artifacts. Default is empty.
     * 
     */
    @Export(name="secondaryKeypairRef", type=String.class, parameters={})
    private Output</* @Nullable */ String> secondaryKeypairRef;

    /**
     * @return Secondary keypair used to sign artifacts. Default is empty.
     * 
     */
    public Output<Optional<String>> secondaryKeypairRef() {
        return Codegen.optional(this.secondaryKeypairRef);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VirtualDebianRepository(String name) {
        this(name, VirtualDebianRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VirtualDebianRepository(String name, VirtualDebianRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VirtualDebianRepository(String name, VirtualDebianRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/virtualDebianRepository:VirtualDebianRepository", name, args == null ? VirtualDebianRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VirtualDebianRepository(String name, Output<String> id, @Nullable VirtualDebianRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/virtualDebianRepository:VirtualDebianRepository", name, state, makeResourceOptions(options, id));
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
    public static VirtualDebianRepository get(String name, Output<String> id, @Nullable VirtualDebianRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VirtualDebianRepository(name, id, state, options);
    }
}

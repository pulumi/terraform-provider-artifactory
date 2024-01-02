// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFilePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFilePlainArgs Empty = new GetFilePlainArgs();

    /**
     * If set to true, an existing file in the output_path will be overwritten. Default: `false`
     * 
     */
    @Import(name="forceOverwrite")
    private @Nullable Boolean forceOverwrite;

    /**
     * @return If set to true, an existing file in the output_path will be overwritten. Default: `false`
     * 
     */
    public Optional<Boolean> forceOverwrite() {
        return Optional.ofNullable(this.forceOverwrite);
    }

    /**
     * The local path the file should be downloaded to.
     * 
     */
    @Import(name="outputPath", required=true)
    private String outputPath;

    /**
     * @return The local path the file should be downloaded to.
     * 
     */
    public String outputPath() {
        return this.outputPath;
    }

    /**
     * The path to the file within the repository.
     * 
     */
    @Import(name="path", required=true)
    private String path;

    /**
     * @return The path to the file within the repository.
     * 
     */
    public String path() {
        return this.path;
    }

    /**
     * If set to `true`, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory
     * if the file exists. More details in the [official documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RetrieveLatestArtifact)
     * 
     */
    @Import(name="pathIsAliased")
    private @Nullable Boolean pathIsAliased;

    /**
     * @return If set to `true`, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory
     * if the file exists. More details in the [official documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RetrieveLatestArtifact)
     * 
     */
    public Optional<Boolean> pathIsAliased() {
        return Optional.ofNullable(this.pathIsAliased);
    }

    /**
     * Name of the repository where the file is stored.
     * 
     */
    @Import(name="repository", required=true)
    private String repository;

    /**
     * @return Name of the repository where the file is stored.
     * 
     */
    public String repository() {
        return this.repository;
    }

    private GetFilePlainArgs() {}

    private GetFilePlainArgs(GetFilePlainArgs $) {
        this.forceOverwrite = $.forceOverwrite;
        this.outputPath = $.outputPath;
        this.path = $.path;
        this.pathIsAliased = $.pathIsAliased;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFilePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFilePlainArgs $;

        public Builder() {
            $ = new GetFilePlainArgs();
        }

        public Builder(GetFilePlainArgs defaults) {
            $ = new GetFilePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param forceOverwrite If set to true, an existing file in the output_path will be overwritten. Default: `false`
         * 
         * @return builder
         * 
         */
        public Builder forceOverwrite(@Nullable Boolean forceOverwrite) {
            $.forceOverwrite = forceOverwrite;
            return this;
        }

        /**
         * @param outputPath The local path the file should be downloaded to.
         * 
         * @return builder
         * 
         */
        public Builder outputPath(String outputPath) {
            $.outputPath = outputPath;
            return this;
        }

        /**
         * @param path The path to the file within the repository.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            $.path = path;
            return this;
        }

        /**
         * @param pathIsAliased If set to `true`, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory
         * if the file exists. More details in the [official documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RetrieveLatestArtifact)
         * 
         * @return builder
         * 
         */
        public Builder pathIsAliased(@Nullable Boolean pathIsAliased) {
            $.pathIsAliased = pathIsAliased;
            return this;
        }

        /**
         * @param repository Name of the repository where the file is stored.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            $.repository = repository;
            return this;
        }

        public GetFilePlainArgs build() {
            if ($.outputPath == null) {
                throw new MissingRequiredPropertyException("GetFilePlainArgs", "outputPath");
            }
            if ($.path == null) {
                throw new MissingRequiredPropertyException("GetFilePlainArgs", "path");
            }
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("GetFilePlainArgs", "repository");
            }
            return $;
        }
    }

}

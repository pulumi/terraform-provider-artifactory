// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ArtifactState extends com.pulumi.resources.ResourceArgs {

    public static final ArtifactState Empty = new ArtifactState();

    /**
     * MD5 checksum of the artifact.
     * 
     */
    @Import(name="checksumMd5")
    private @Nullable Output<String> checksumMd5;

    /**
     * @return MD5 checksum of the artifact.
     * 
     */
    public Optional<Output<String>> checksumMd5() {
        return Optional.ofNullable(this.checksumMd5);
    }

    /**
     * SHA1 checksum of the artifact.
     * 
     */
    @Import(name="checksumSha1")
    private @Nullable Output<String> checksumSha1;

    /**
     * @return SHA1 checksum of the artifact.
     * 
     */
    public Optional<Output<String>> checksumSha1() {
        return Optional.ofNullable(this.checksumSha1);
    }

    /**
     * SHA256 checksum of the artifact.
     * 
     */
    @Import(name="checksumSha256")
    private @Nullable Output<String> checksumSha256;

    /**
     * @return SHA256 checksum of the artifact.
     * 
     */
    public Optional<Output<String>> checksumSha256() {
        return Optional.ofNullable(this.checksumSha256);
    }

    /**
     * Base64 content of the source file. Conflicts with `file_path`. Either one of these attribute must be set.
     * 
     */
    @Import(name="contentBase64")
    private @Nullable Output<String> contentBase64;

    /**
     * @return Base64 content of the source file. Conflicts with `file_path`. Either one of these attribute must be set.
     * 
     */
    public Optional<Output<String>> contentBase64() {
        return Optional.ofNullable(this.contentBase64);
    }

    /**
     * Timestamp when artifact is created.
     * 
     */
    @Import(name="created")
    private @Nullable Output<String> created;

    /**
     * @return Timestamp when artifact is created.
     * 
     */
    public Optional<Output<String>> created() {
        return Optional.ofNullable(this.created);
    }

    /**
     * User who deploys the artifact.
     * 
     */
    @Import(name="createdBy")
    private @Nullable Output<String> createdBy;

    /**
     * @return User who deploys the artifact.
     * 
     */
    public Optional<Output<String>> createdBy() {
        return Optional.ofNullable(this.createdBy);
    }

    /**
     * Download URI of the artifact.
     * 
     */
    @Import(name="downloadUri")
    private @Nullable Output<String> downloadUri;

    /**
     * @return Download URI of the artifact.
     * 
     */
    public Optional<Output<String>> downloadUri() {
        return Optional.ofNullable(this.downloadUri);
    }

    /**
     * Path to the source file. Conflicts with `content_base64`. Either one of these attribute must be set.
     * 
     */
    @Import(name="filePath")
    private @Nullable Output<String> filePath;

    /**
     * @return Path to the source file. Conflicts with `content_base64`. Either one of these attribute must be set.
     * 
     */
    public Optional<Output<String>> filePath() {
        return Optional.ofNullable(this.filePath);
    }

    /**
     * MIME type of the artifact.
     * 
     */
    @Import(name="mimeType")
    private @Nullable Output<String> mimeType;

    /**
     * @return MIME type of the artifact.
     * 
     */
    public Optional<Output<String>> mimeType() {
        return Optional.ofNullable(this.mimeType);
    }

    /**
     * The relative path in the target repository. Must begin with a &#39;/&#39;. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The relative path in the target repository. Must begin with a &#39;/&#39;. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * Name of the respository.
     * 
     */
    @Import(name="repository")
    private @Nullable Output<String> repository;

    /**
     * @return Name of the respository.
     * 
     */
    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    /**
     * Size of the artifact, in bytes.
     * 
     */
    @Import(name="size")
    private @Nullable Output<Integer> size;

    /**
     * @return Size of the artifact, in bytes.
     * 
     */
    public Optional<Output<Integer>> size() {
        return Optional.ofNullable(this.size);
    }

    /**
     * URI of the artifact.
     * 
     */
    @Import(name="uri")
    private @Nullable Output<String> uri;

    /**
     * @return URI of the artifact.
     * 
     */
    public Optional<Output<String>> uri() {
        return Optional.ofNullable(this.uri);
    }

    private ArtifactState() {}

    private ArtifactState(ArtifactState $) {
        this.checksumMd5 = $.checksumMd5;
        this.checksumSha1 = $.checksumSha1;
        this.checksumSha256 = $.checksumSha256;
        this.contentBase64 = $.contentBase64;
        this.created = $.created;
        this.createdBy = $.createdBy;
        this.downloadUri = $.downloadUri;
        this.filePath = $.filePath;
        this.mimeType = $.mimeType;
        this.path = $.path;
        this.repository = $.repository;
        this.size = $.size;
        this.uri = $.uri;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ArtifactState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ArtifactState $;

        public Builder() {
            $ = new ArtifactState();
        }

        public Builder(ArtifactState defaults) {
            $ = new ArtifactState(Objects.requireNonNull(defaults));
        }

        /**
         * @param checksumMd5 MD5 checksum of the artifact.
         * 
         * @return builder
         * 
         */
        public Builder checksumMd5(@Nullable Output<String> checksumMd5) {
            $.checksumMd5 = checksumMd5;
            return this;
        }

        /**
         * @param checksumMd5 MD5 checksum of the artifact.
         * 
         * @return builder
         * 
         */
        public Builder checksumMd5(String checksumMd5) {
            return checksumMd5(Output.of(checksumMd5));
        }

        /**
         * @param checksumSha1 SHA1 checksum of the artifact.
         * 
         * @return builder
         * 
         */
        public Builder checksumSha1(@Nullable Output<String> checksumSha1) {
            $.checksumSha1 = checksumSha1;
            return this;
        }

        /**
         * @param checksumSha1 SHA1 checksum of the artifact.
         * 
         * @return builder
         * 
         */
        public Builder checksumSha1(String checksumSha1) {
            return checksumSha1(Output.of(checksumSha1));
        }

        /**
         * @param checksumSha256 SHA256 checksum of the artifact.
         * 
         * @return builder
         * 
         */
        public Builder checksumSha256(@Nullable Output<String> checksumSha256) {
            $.checksumSha256 = checksumSha256;
            return this;
        }

        /**
         * @param checksumSha256 SHA256 checksum of the artifact.
         * 
         * @return builder
         * 
         */
        public Builder checksumSha256(String checksumSha256) {
            return checksumSha256(Output.of(checksumSha256));
        }

        /**
         * @param contentBase64 Base64 content of the source file. Conflicts with `file_path`. Either one of these attribute must be set.
         * 
         * @return builder
         * 
         */
        public Builder contentBase64(@Nullable Output<String> contentBase64) {
            $.contentBase64 = contentBase64;
            return this;
        }

        /**
         * @param contentBase64 Base64 content of the source file. Conflicts with `file_path`. Either one of these attribute must be set.
         * 
         * @return builder
         * 
         */
        public Builder contentBase64(String contentBase64) {
            return contentBase64(Output.of(contentBase64));
        }

        /**
         * @param created Timestamp when artifact is created.
         * 
         * @return builder
         * 
         */
        public Builder created(@Nullable Output<String> created) {
            $.created = created;
            return this;
        }

        /**
         * @param created Timestamp when artifact is created.
         * 
         * @return builder
         * 
         */
        public Builder created(String created) {
            return created(Output.of(created));
        }

        /**
         * @param createdBy User who deploys the artifact.
         * 
         * @return builder
         * 
         */
        public Builder createdBy(@Nullable Output<String> createdBy) {
            $.createdBy = createdBy;
            return this;
        }

        /**
         * @param createdBy User who deploys the artifact.
         * 
         * @return builder
         * 
         */
        public Builder createdBy(String createdBy) {
            return createdBy(Output.of(createdBy));
        }

        /**
         * @param downloadUri Download URI of the artifact.
         * 
         * @return builder
         * 
         */
        public Builder downloadUri(@Nullable Output<String> downloadUri) {
            $.downloadUri = downloadUri;
            return this;
        }

        /**
         * @param downloadUri Download URI of the artifact.
         * 
         * @return builder
         * 
         */
        public Builder downloadUri(String downloadUri) {
            return downloadUri(Output.of(downloadUri));
        }

        /**
         * @param filePath Path to the source file. Conflicts with `content_base64`. Either one of these attribute must be set.
         * 
         * @return builder
         * 
         */
        public Builder filePath(@Nullable Output<String> filePath) {
            $.filePath = filePath;
            return this;
        }

        /**
         * @param filePath Path to the source file. Conflicts with `content_base64`. Either one of these attribute must be set.
         * 
         * @return builder
         * 
         */
        public Builder filePath(String filePath) {
            return filePath(Output.of(filePath));
        }

        /**
         * @param mimeType MIME type of the artifact.
         * 
         * @return builder
         * 
         */
        public Builder mimeType(@Nullable Output<String> mimeType) {
            $.mimeType = mimeType;
            return this;
        }

        /**
         * @param mimeType MIME type of the artifact.
         * 
         * @return builder
         * 
         */
        public Builder mimeType(String mimeType) {
            return mimeType(Output.of(mimeType));
        }

        /**
         * @param path The relative path in the target repository. Must begin with a &#39;/&#39;. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The relative path in the target repository. Must begin with a &#39;/&#39;. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param repository Name of the respository.
         * 
         * @return builder
         * 
         */
        public Builder repository(@Nullable Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository Name of the respository.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param size Size of the artifact, in bytes.
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size Size of the artifact, in bytes.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param uri URI of the artifact.
         * 
         * @return builder
         * 
         */
        public Builder uri(@Nullable Output<String> uri) {
            $.uri = uri;
            return this;
        }

        /**
         * @param uri URI of the artifact.
         * 
         * @return builder
         * 
         */
        public Builder uri(String uri) {
            return uri(Output.of(uri));
        }

        public ArtifactState build() {
            return $;
        }
    }

}

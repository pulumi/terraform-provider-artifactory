// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CertificateArgs extends com.pulumi.resources.ResourceArgs {

    public static final CertificateArgs Empty = new CertificateArgs();

    /**
     * Name of certificate.
     * 
     */
    @Import(name="alias", required=true)
    private Output<String> alias;

    /**
     * @return Name of certificate.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }

    /**
     * PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
     * 
     */
    @Import(name="content")
    private @Nullable Output<String> content;

    /**
     * @return PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
     * 
     */
    public Optional<Output<String>> content() {
        return Optional.ofNullable(this.content);
    }

    /**
     * Path to the PEM file. Cannot be set with `content` attribute simultaneously.
     * 
     */
    @Import(name="file")
    private @Nullable Output<String> file;

    /**
     * @return Path to the PEM file. Cannot be set with `content` attribute simultaneously.
     * 
     */
    public Optional<Output<String>> file() {
        return Optional.ofNullable(this.file);
    }

    private CertificateArgs() {}

    private CertificateArgs(CertificateArgs $) {
        this.alias = $.alias;
        this.content = $.content;
        this.file = $.file;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CertificateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CertificateArgs $;

        public Builder() {
            $ = new CertificateArgs();
        }

        public Builder(CertificateArgs defaults) {
            $ = new CertificateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alias Name of certificate.
         * 
         * @return builder
         * 
         */
        public Builder alias(Output<String> alias) {
            $.alias = alias;
            return this;
        }

        /**
         * @param alias Name of certificate.
         * 
         * @return builder
         * 
         */
        public Builder alias(String alias) {
            return alias(Output.of(alias));
        }

        /**
         * @param content PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
         * 
         * @return builder
         * 
         */
        public Builder content(@Nullable Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        /**
         * @param file Path to the PEM file. Cannot be set with `content` attribute simultaneously.
         * 
         * @return builder
         * 
         */
        public Builder file(@Nullable Output<String> file) {
            $.file = file;
            return this;
        }

        /**
         * @param file Path to the PEM file. Cannot be set with `content` attribute simultaneously.
         * 
         * @return builder
         * 
         */
        public Builder file(String file) {
            return file(Output.of(file));
        }

        public CertificateArgs build() {
            if ($.alias == null) {
                throw new MissingRequiredPropertyException("CertificateArgs", "alias");
            }
            return $;
        }
    }

}

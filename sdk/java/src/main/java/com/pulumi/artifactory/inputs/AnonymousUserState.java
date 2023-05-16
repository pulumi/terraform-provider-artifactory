// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AnonymousUserState extends com.pulumi.resources.ResourceArgs {

    public static final AnonymousUserState Empty = new AnonymousUserState();

    /**
     * Username for anonymous user. This is only for ensuring resource schema is valid for Terraform. This is not meant to be
     * set or updated in the HCL.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Username for anonymous user. This is only for ensuring resource schema is valid for Terraform. This is not meant to be
     * set or updated in the HCL.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private AnonymousUserState() {}

    private AnonymousUserState(AnonymousUserState $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AnonymousUserState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AnonymousUserState $;

        public Builder() {
            $ = new AnonymousUserState();
        }

        public Builder(AnonymousUserState defaults) {
            $ = new AnonymousUserState(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Username for anonymous user. This is only for ensuring resource schema is valid for Terraform. This is not meant to be
         * set or updated in the HCL.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Username for anonymous user. This is only for ensuring resource schema is valid for Terraform. This is not meant to be
         * set or updated in the HCL.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public AnonymousUserState build() {
            return $;
        }
    }

}

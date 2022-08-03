// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ApiKeyState extends com.pulumi.resources.ResourceArgs {

    public static final ApiKeyState Empty = new ApiKeyState();

    /**
     * The API key.
     * 
     */
    @Import(name="apiKey")
    private @Nullable Output<String> apiKey;

    /**
     * @return The API key.
     * 
     */
    public Optional<Output<String>> apiKey() {
        return Optional.ofNullable(this.apiKey);
    }

    private ApiKeyState() {}

    private ApiKeyState(ApiKeyState $) {
        this.apiKey = $.apiKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApiKeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApiKeyState $;

        public Builder() {
            $ = new ApiKeyState();
        }

        public Builder(ApiKeyState defaults) {
            $ = new ApiKeyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiKey The API key.
         * 
         * @return builder
         * 
         */
        public Builder apiKey(@Nullable Output<String> apiKey) {
            $.apiKey = apiKey;
            return this;
        }

        /**
         * @param apiKey The API key.
         * 
         * @return builder
         * 
         */
        public Builder apiKey(String apiKey) {
            return apiKey(Output.of(apiKey));
        }

        public ApiKeyState build() {
            return $;
        }
    }

}

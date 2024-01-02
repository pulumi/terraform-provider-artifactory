// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.inputs.ArtifactCustomWebhookCriteriaArgs;
import com.pulumi.artifactory.inputs.ArtifactCustomWebhookHandlerArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ArtifactCustomWebhookArgs extends com.pulumi.resources.ResourceArgs {

    public static final ArtifactCustomWebhookArgs Empty = new ArtifactCustomWebhookArgs();

    /**
     * Specifies where the webhook will be applied on which repositories.
     * 
     */
    @Import(name="criteria", required=true)
    private Output<ArtifactCustomWebhookCriteriaArgs> criteria;

    /**
     * @return Specifies where the webhook will be applied on which repositories.
     * 
     */
    public Output<ArtifactCustomWebhookCriteriaArgs> criteria() {
        return this.criteria;
    }

    /**
     * Webhook description. Max length 1000 characters.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Webhook description. Max length 1000 characters.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Status of webhook. Default to `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Status of webhook. Default to `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
     * 
     */
    @Import(name="eventTypes", required=true)
    private Output<List<String>> eventTypes;

    /**
     * @return List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
     * 
     */
    public Output<List<String>> eventTypes() {
        return this.eventTypes;
    }

    /**
     * At least one is required.
     * 
     */
    @Import(name="handlers", required=true)
    private Output<List<ArtifactCustomWebhookHandlerArgs>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Output<List<ArtifactCustomWebhookHandlerArgs>> handlers() {
        return this.handlers;
    }

    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    private ArtifactCustomWebhookArgs() {}

    private ArtifactCustomWebhookArgs(ArtifactCustomWebhookArgs $) {
        this.criteria = $.criteria;
        this.description = $.description;
        this.enabled = $.enabled;
        this.eventTypes = $.eventTypes;
        this.handlers = $.handlers;
        this.key = $.key;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ArtifactCustomWebhookArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ArtifactCustomWebhookArgs $;

        public Builder() {
            $ = new ArtifactCustomWebhookArgs();
        }

        public Builder(ArtifactCustomWebhookArgs defaults) {
            $ = new ArtifactCustomWebhookArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param criteria Specifies where the webhook will be applied on which repositories.
         * 
         * @return builder
         * 
         */
        public Builder criteria(Output<ArtifactCustomWebhookCriteriaArgs> criteria) {
            $.criteria = criteria;
            return this;
        }

        /**
         * @param criteria Specifies where the webhook will be applied on which repositories.
         * 
         * @return builder
         * 
         */
        public Builder criteria(ArtifactCustomWebhookCriteriaArgs criteria) {
            return criteria(Output.of(criteria));
        }

        /**
         * @param description Webhook description. Max length 1000 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Webhook description. Max length 1000 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param enabled Status of webhook. Default to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Status of webhook. Default to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param eventTypes List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(Output<List<String>> eventTypes) {
            $.eventTypes = eventTypes;
            return this;
        }

        /**
         * @param eventTypes List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(List<String> eventTypes) {
            return eventTypes(Output.of(eventTypes));
        }

        /**
         * @param eventTypes List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(String... eventTypes) {
            return eventTypes(List.of(eventTypes));
        }

        /**
         * @param handlers At least one is required.
         * 
         * @return builder
         * 
         */
        public Builder handlers(Output<List<ArtifactCustomWebhookHandlerArgs>> handlers) {
            $.handlers = handlers;
            return this;
        }

        /**
         * @param handlers At least one is required.
         * 
         * @return builder
         * 
         */
        public Builder handlers(List<ArtifactCustomWebhookHandlerArgs> handlers) {
            return handlers(Output.of(handlers));
        }

        /**
         * @param handlers At least one is required.
         * 
         * @return builder
         * 
         */
        public Builder handlers(ArtifactCustomWebhookHandlerArgs... handlers) {
            return handlers(List.of(handlers));
        }

        /**
         * @param key The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        public ArtifactCustomWebhookArgs build() {
            if ($.criteria == null) {
                throw new MissingRequiredPropertyException("ArtifactCustomWebhookArgs", "criteria");
            }
            if ($.eventTypes == null) {
                throw new MissingRequiredPropertyException("ArtifactCustomWebhookArgs", "eventTypes");
            }
            if ($.handlers == null) {
                throw new MissingRequiredPropertyException("ArtifactCustomWebhookArgs", "handlers");
            }
            if ($.key == null) {
                throw new MissingRequiredPropertyException("ArtifactCustomWebhookArgs", "key");
            }
            return $;
        }
    }

}

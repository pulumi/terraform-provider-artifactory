// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.inputs.ReleaseBundleCustomWebhookCriteriaArgs;
import com.pulumi.artifactory.inputs.ReleaseBundleCustomWebhookHandlerArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReleaseBundleCustomWebhookArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReleaseBundleCustomWebhookArgs Empty = new ReleaseBundleCustomWebhookArgs();

    /**
     * Specifies where the webhook will be applied on which repositories.
     * 
     */
    @Import(name="criteria", required=true)
    private Output<ReleaseBundleCustomWebhookCriteriaArgs> criteria;

    /**
     * @return Specifies where the webhook will be applied on which repositories.
     * 
     */
    public Output<ReleaseBundleCustomWebhookCriteriaArgs> criteria() {
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
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
     * 
     */
    @Import(name="eventTypes", required=true)
    private Output<List<String>> eventTypes;

    /**
     * @return List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
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
    private Output<List<ReleaseBundleCustomWebhookHandlerArgs>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Output<List<ReleaseBundleCustomWebhookHandlerArgs>> handlers() {
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

    private ReleaseBundleCustomWebhookArgs() {}

    private ReleaseBundleCustomWebhookArgs(ReleaseBundleCustomWebhookArgs $) {
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
    public static Builder builder(ReleaseBundleCustomWebhookArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReleaseBundleCustomWebhookArgs $;

        public Builder() {
            $ = new ReleaseBundleCustomWebhookArgs();
        }

        public Builder(ReleaseBundleCustomWebhookArgs defaults) {
            $ = new ReleaseBundleCustomWebhookArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param criteria Specifies where the webhook will be applied on which repositories.
         * 
         * @return builder
         * 
         */
        public Builder criteria(Output<ReleaseBundleCustomWebhookCriteriaArgs> criteria) {
            $.criteria = criteria;
            return this;
        }

        /**
         * @param criteria Specifies where the webhook will be applied on which repositories.
         * 
         * @return builder
         * 
         */
        public Builder criteria(ReleaseBundleCustomWebhookCriteriaArgs criteria) {
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
         * @param eventTypes List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(Output<List<String>> eventTypes) {
            $.eventTypes = eventTypes;
            return this;
        }

        /**
         * @param eventTypes List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(List<String> eventTypes) {
            return eventTypes(Output.of(eventTypes));
        }

        /**
         * @param eventTypes List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
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
        public Builder handlers(Output<List<ReleaseBundleCustomWebhookHandlerArgs>> handlers) {
            $.handlers = handlers;
            return this;
        }

        /**
         * @param handlers At least one is required.
         * 
         * @return builder
         * 
         */
        public Builder handlers(List<ReleaseBundleCustomWebhookHandlerArgs> handlers) {
            return handlers(Output.of(handlers));
        }

        /**
         * @param handlers At least one is required.
         * 
         * @return builder
         * 
         */
        public Builder handlers(ReleaseBundleCustomWebhookHandlerArgs... handlers) {
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

        public ReleaseBundleCustomWebhookArgs build() {
            $.criteria = Objects.requireNonNull($.criteria, "expected parameter 'criteria' to be non-null");
            $.eventTypes = Objects.requireNonNull($.eventTypes, "expected parameter 'eventTypes' to be non-null");
            $.handlers = Objects.requireNonNull($.handlers, "expected parameter 'handlers' to be non-null");
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.inputs.DockerWebhookCriteriaArgs;
import com.pulumi.artifactory.inputs.DockerWebhookHandlerArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DockerWebhookArgs extends com.pulumi.resources.ResourceArgs {

    public static final DockerWebhookArgs Empty = new DockerWebhookArgs();

    /**
     * Specifies where the webhook will be applied on which repositories.
     * 
     */
    @Import(name="criteria", required=true)
    private Output<DockerWebhookCriteriaArgs> criteria;

    /**
     * @return Specifies where the webhook will be applied on which repositories.
     * 
     */
    public Output<DockerWebhookCriteriaArgs> criteria() {
        return this.criteria;
    }

    /**
     * Description of webhook. Max length 1000 characters.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of webhook. Max length 1000 characters.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Status of webhook. Default to &#39;true&#39;
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Status of webhook. Default to &#39;true&#39;
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
     * values: pushed, deleted, promoted
     * 
     */
    @Import(name="eventTypes", required=true)
    private Output<List<String>> eventTypes;

    /**
     * @return List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
     * values: pushed, deleted, promoted
     * 
     */
    public Output<List<String>> eventTypes() {
        return this.eventTypes;
    }

    @Import(name="handlers", required=true)
    private Output<List<DockerWebhookHandlerArgs>> handlers;

    public Output<List<DockerWebhookHandlerArgs>> handlers() {
        return this.handlers;
    }

    /**
     * Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    private DockerWebhookArgs() {}

    private DockerWebhookArgs(DockerWebhookArgs $) {
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
    public static Builder builder(DockerWebhookArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DockerWebhookArgs $;

        public Builder() {
            $ = new DockerWebhookArgs();
        }

        public Builder(DockerWebhookArgs defaults) {
            $ = new DockerWebhookArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param criteria Specifies where the webhook will be applied on which repositories.
         * 
         * @return builder
         * 
         */
        public Builder criteria(Output<DockerWebhookCriteriaArgs> criteria) {
            $.criteria = criteria;
            return this;
        }

        /**
         * @param criteria Specifies where the webhook will be applied on which repositories.
         * 
         * @return builder
         * 
         */
        public Builder criteria(DockerWebhookCriteriaArgs criteria) {
            return criteria(Output.of(criteria));
        }

        /**
         * @param description Description of webhook. Max length 1000 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of webhook. Max length 1000 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param enabled Status of webhook. Default to &#39;true&#39;
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Status of webhook. Default to &#39;true&#39;
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param eventTypes List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
         * values: pushed, deleted, promoted
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(Output<List<String>> eventTypes) {
            $.eventTypes = eventTypes;
            return this;
        }

        /**
         * @param eventTypes List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
         * values: pushed, deleted, promoted
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(List<String> eventTypes) {
            return eventTypes(Output.of(eventTypes));
        }

        /**
         * @param eventTypes List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
         * values: pushed, deleted, promoted
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(String... eventTypes) {
            return eventTypes(List.of(eventTypes));
        }

        public Builder handlers(Output<List<DockerWebhookHandlerArgs>> handlers) {
            $.handlers = handlers;
            return this;
        }

        public Builder handlers(List<DockerWebhookHandlerArgs> handlers) {
            return handlers(Output.of(handlers));
        }

        public Builder handlers(DockerWebhookHandlerArgs... handlers) {
            return handlers(List.of(handlers));
        }

        /**
         * @param key Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        public DockerWebhookArgs build() {
            $.criteria = Objects.requireNonNull($.criteria, "expected parameter 'criteria' to be non-null");
            $.eventTypes = Objects.requireNonNull($.eventTypes, "expected parameter 'eventTypes' to be non-null");
            $.handlers = Objects.requireNonNull($.handlers, "expected parameter 'handlers' to be non-null");
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}

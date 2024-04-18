// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.PropertySetPropertyPredefinedValueArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PropertySetPropertyArgs extends com.pulumi.resources.ResourceArgs {

    public static final PropertySetPropertyArgs Empty = new PropertySetPropertyArgs();

    /**
     * Disables `multiple_choice` if set to `false` at the same time with multiple_choice set to `true`. Default value is `false`
     * 
     */
    @Import(name="closedPredefinedValues")
    private @Nullable Output<Boolean> closedPredefinedValues;

    /**
     * @return Disables `multiple_choice` if set to `false` at the same time with multiple_choice set to `true`. Default value is `false`
     * 
     */
    public Optional<Output<Boolean>> closedPredefinedValues() {
        return Optional.ofNullable(this.closedPredefinedValues);
    }

    /**
     * Defines if user can select multiple values. `closed_predefined_values` should be set to `true`. Default value is `false`.
     * 
     */
    @Import(name="multipleChoice")
    private @Nullable Output<Boolean> multipleChoice;

    /**
     * @return Defines if user can select multiple values. `closed_predefined_values` should be set to `true`. Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> multipleChoice() {
        return Optional.ofNullable(this.multipleChoice);
    }

    /**
     * Predefined property name.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Predefined property name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Properties in the property set.
     * 
     */
    @Import(name="predefinedValues")
    private @Nullable Output<List<PropertySetPropertyPredefinedValueArgs>> predefinedValues;

    /**
     * @return Properties in the property set.
     * 
     */
    public Optional<Output<List<PropertySetPropertyPredefinedValueArgs>>> predefinedValues() {
        return Optional.ofNullable(this.predefinedValues);
    }

    private PropertySetPropertyArgs() {}

    private PropertySetPropertyArgs(PropertySetPropertyArgs $) {
        this.closedPredefinedValues = $.closedPredefinedValues;
        this.multipleChoice = $.multipleChoice;
        this.name = $.name;
        this.predefinedValues = $.predefinedValues;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PropertySetPropertyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PropertySetPropertyArgs $;

        public Builder() {
            $ = new PropertySetPropertyArgs();
        }

        public Builder(PropertySetPropertyArgs defaults) {
            $ = new PropertySetPropertyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param closedPredefinedValues Disables `multiple_choice` if set to `false` at the same time with multiple_choice set to `true`. Default value is `false`
         * 
         * @return builder
         * 
         */
        public Builder closedPredefinedValues(@Nullable Output<Boolean> closedPredefinedValues) {
            $.closedPredefinedValues = closedPredefinedValues;
            return this;
        }

        /**
         * @param closedPredefinedValues Disables `multiple_choice` if set to `false` at the same time with multiple_choice set to `true`. Default value is `false`
         * 
         * @return builder
         * 
         */
        public Builder closedPredefinedValues(Boolean closedPredefinedValues) {
            return closedPredefinedValues(Output.of(closedPredefinedValues));
        }

        /**
         * @param multipleChoice Defines if user can select multiple values. `closed_predefined_values` should be set to `true`. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder multipleChoice(@Nullable Output<Boolean> multipleChoice) {
            $.multipleChoice = multipleChoice;
            return this;
        }

        /**
         * @param multipleChoice Defines if user can select multiple values. `closed_predefined_values` should be set to `true`. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder multipleChoice(Boolean multipleChoice) {
            return multipleChoice(Output.of(multipleChoice));
        }

        /**
         * @param name Predefined property name.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Predefined property name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param predefinedValues Properties in the property set.
         * 
         * @return builder
         * 
         */
        public Builder predefinedValues(@Nullable Output<List<PropertySetPropertyPredefinedValueArgs>> predefinedValues) {
            $.predefinedValues = predefinedValues;
            return this;
        }

        /**
         * @param predefinedValues Properties in the property set.
         * 
         * @return builder
         * 
         */
        public Builder predefinedValues(List<PropertySetPropertyPredefinedValueArgs> predefinedValues) {
            return predefinedValues(Output.of(predefinedValues));
        }

        /**
         * @param predefinedValues Properties in the property set.
         * 
         * @return builder
         * 
         */
        public Builder predefinedValues(PropertySetPropertyPredefinedValueArgs... predefinedValues) {
            return predefinedValues(List.of(predefinedValues));
        }

        public PropertySetPropertyArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("PropertySetPropertyArgs", "name");
            }
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class PropertySetPropertyPredefinedValueArgs extends com.pulumi.resources.ResourceArgs {

    public static final PropertySetPropertyPredefinedValueArgs Empty = new PropertySetPropertyPredefinedValueArgs();

    /**
     * Whether the value is selected by default in the UI.
     * 
     */
    @Import(name="defaultValue", required=true)
    private Output<Boolean> defaultValue;

    /**
     * @return Whether the value is selected by default in the UI.
     * 
     */
    public Output<Boolean> defaultValue() {
        return this.defaultValue;
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

    private PropertySetPropertyPredefinedValueArgs() {}

    private PropertySetPropertyPredefinedValueArgs(PropertySetPropertyPredefinedValueArgs $) {
        this.defaultValue = $.defaultValue;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PropertySetPropertyPredefinedValueArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PropertySetPropertyPredefinedValueArgs $;

        public Builder() {
            $ = new PropertySetPropertyPredefinedValueArgs();
        }

        public Builder(PropertySetPropertyPredefinedValueArgs defaults) {
            $ = new PropertySetPropertyPredefinedValueArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultValue Whether the value is selected by default in the UI.
         * 
         * @return builder
         * 
         */
        public Builder defaultValue(Output<Boolean> defaultValue) {
            $.defaultValue = defaultValue;
            return this;
        }

        /**
         * @param defaultValue Whether the value is selected by default in the UI.
         * 
         * @return builder
         * 
         */
        public Builder defaultValue(Boolean defaultValue) {
            return defaultValue(Output.of(defaultValue));
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

        public PropertySetPropertyPredefinedValueArgs build() {
            if ($.defaultValue == null) {
                throw new MissingRequiredPropertyException("PropertySetPropertyPredefinedValueArgs", "defaultValue");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("PropertySetPropertyPredefinedValueArgs", "name");
            }
            return $;
        }
    }

}

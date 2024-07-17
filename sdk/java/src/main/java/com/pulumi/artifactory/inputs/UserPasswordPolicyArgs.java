// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserPasswordPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserPasswordPolicyArgs Empty = new UserPasswordPolicyArgs();

    /**
     * Minimum number of digits that the password must contain
     * 
     */
    @Import(name="digit")
    private @Nullable Output<Integer> digit;

    /**
     * @return Minimum number of digits that the password must contain
     * 
     */
    public Optional<Output<Integer>> digit() {
        return Optional.ofNullable(this.digit);
    }

    /**
     * Minimum length of the password
     * 
     */
    @Import(name="length")
    private @Nullable Output<Integer> length;

    /**
     * @return Minimum length of the password
     * 
     */
    public Optional<Output<Integer>> length() {
        return Optional.ofNullable(this.length);
    }

    /**
     * Minimum number of lowercase letters that the password must contain
     * 
     */
    @Import(name="lowercase")
    private @Nullable Output<Integer> lowercase;

    /**
     * @return Minimum number of lowercase letters that the password must contain
     * 
     */
    public Optional<Output<Integer>> lowercase() {
        return Optional.ofNullable(this.lowercase);
    }

    /**
     * Minimum number of special char that the password must contain. Special chars list: `!&#34;#$%&amp;&#39;()*+,-./:;&lt;=&gt;?{@literal @}[\]^_``{|}~`
     * 
     */
    @Import(name="specialChar")
    private @Nullable Output<Integer> specialChar;

    /**
     * @return Minimum number of special char that the password must contain. Special chars list: `!&#34;#$%&amp;&#39;()*+,-./:;&lt;=&gt;?{@literal @}[\]^_``{|}~`
     * 
     */
    public Optional<Output<Integer>> specialChar() {
        return Optional.ofNullable(this.specialChar);
    }

    /**
     * Minimum number of uppercase letters that the password must contain
     * 
     */
    @Import(name="uppercase")
    private @Nullable Output<Integer> uppercase;

    /**
     * @return Minimum number of uppercase letters that the password must contain
     * 
     */
    public Optional<Output<Integer>> uppercase() {
        return Optional.ofNullable(this.uppercase);
    }

    private UserPasswordPolicyArgs() {}

    private UserPasswordPolicyArgs(UserPasswordPolicyArgs $) {
        this.digit = $.digit;
        this.length = $.length;
        this.lowercase = $.lowercase;
        this.specialChar = $.specialChar;
        this.uppercase = $.uppercase;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserPasswordPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserPasswordPolicyArgs $;

        public Builder() {
            $ = new UserPasswordPolicyArgs();
        }

        public Builder(UserPasswordPolicyArgs defaults) {
            $ = new UserPasswordPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param digit Minimum number of digits that the password must contain
         * 
         * @return builder
         * 
         */
        public Builder digit(@Nullable Output<Integer> digit) {
            $.digit = digit;
            return this;
        }

        /**
         * @param digit Minimum number of digits that the password must contain
         * 
         * @return builder
         * 
         */
        public Builder digit(Integer digit) {
            return digit(Output.of(digit));
        }

        /**
         * @param length Minimum length of the password
         * 
         * @return builder
         * 
         */
        public Builder length(@Nullable Output<Integer> length) {
            $.length = length;
            return this;
        }

        /**
         * @param length Minimum length of the password
         * 
         * @return builder
         * 
         */
        public Builder length(Integer length) {
            return length(Output.of(length));
        }

        /**
         * @param lowercase Minimum number of lowercase letters that the password must contain
         * 
         * @return builder
         * 
         */
        public Builder lowercase(@Nullable Output<Integer> lowercase) {
            $.lowercase = lowercase;
            return this;
        }

        /**
         * @param lowercase Minimum number of lowercase letters that the password must contain
         * 
         * @return builder
         * 
         */
        public Builder lowercase(Integer lowercase) {
            return lowercase(Output.of(lowercase));
        }

        /**
         * @param specialChar Minimum number of special char that the password must contain. Special chars list: `!&#34;#$%&amp;&#39;()*+,-./:;&lt;=&gt;?{@literal @}[\]^_``{|}~`
         * 
         * @return builder
         * 
         */
        public Builder specialChar(@Nullable Output<Integer> specialChar) {
            $.specialChar = specialChar;
            return this;
        }

        /**
         * @param specialChar Minimum number of special char that the password must contain. Special chars list: `!&#34;#$%&amp;&#39;()*+,-./:;&lt;=&gt;?{@literal @}[\]^_``{|}~`
         * 
         * @return builder
         * 
         */
        public Builder specialChar(Integer specialChar) {
            return specialChar(Output.of(specialChar));
        }

        /**
         * @param uppercase Minimum number of uppercase letters that the password must contain
         * 
         * @return builder
         * 
         */
        public Builder uppercase(@Nullable Output<Integer> uppercase) {
            $.uppercase = uppercase;
            return this;
        }

        /**
         * @param uppercase Minimum number of uppercase letters that the password must contain
         * 
         * @return builder
         * 
         */
        public Builder uppercase(Integer uppercase) {
            return uppercase(Output.of(uppercase));
        }

        public UserPasswordPolicyArgs build() {
            return $;
        }
    }

}

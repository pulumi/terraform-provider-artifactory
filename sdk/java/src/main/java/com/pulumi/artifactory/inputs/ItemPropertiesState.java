// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ItemPropertiesState extends com.pulumi.resources.ResourceArgs {

    public static final ItemPropertiesState Empty = new ItemPropertiesState();

    /**
     * Add this property to the selected folder and to all of artifacts and folders under this folder. Default to `false`
     * 
     */
    @Import(name="isRecursive")
    private @Nullable Output<Boolean> isRecursive;

    /**
     * @return Add this property to the selected folder and to all of artifacts and folders under this folder. Default to `false`
     * 
     */
    public Optional<Output<Boolean>> isRecursive() {
        return Optional.ofNullable(this.isRecursive);
    }

    /**
     * The relative path of the item (file/folder/repository). Leave unset for repository.
     * 
     */
    @Import(name="itemPath")
    private @Nullable Output<String> itemPath;

    /**
     * @return The relative path of the item (file/folder/repository). Leave unset for repository.
     * 
     */
    public Optional<Output<String>> itemPath() {
        return Optional.ofNullable(this.itemPath);
    }

    /**
     * Map of key and list of values.
     * 
     */
    @Import(name="properties")
    private @Nullable Output<Map<String,List<String>>> properties;

    /**
     * @return Map of key and list of values.
     * 
     */
    public Optional<Output<Map<String,List<String>>>> properties() {
        return Optional.ofNullable(this.properties);
    }

    /**
     * Respository key.
     * 
     */
    @Import(name="repoKey")
    private @Nullable Output<String> repoKey;

    /**
     * @return Respository key.
     * 
     */
    public Optional<Output<String>> repoKey() {
        return Optional.ofNullable(this.repoKey);
    }

    private ItemPropertiesState() {}

    private ItemPropertiesState(ItemPropertiesState $) {
        this.isRecursive = $.isRecursive;
        this.itemPath = $.itemPath;
        this.properties = $.properties;
        this.repoKey = $.repoKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ItemPropertiesState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ItemPropertiesState $;

        public Builder() {
            $ = new ItemPropertiesState();
        }

        public Builder(ItemPropertiesState defaults) {
            $ = new ItemPropertiesState(Objects.requireNonNull(defaults));
        }

        /**
         * @param isRecursive Add this property to the selected folder and to all of artifacts and folders under this folder. Default to `false`
         * 
         * @return builder
         * 
         */
        public Builder isRecursive(@Nullable Output<Boolean> isRecursive) {
            $.isRecursive = isRecursive;
            return this;
        }

        /**
         * @param isRecursive Add this property to the selected folder and to all of artifacts and folders under this folder. Default to `false`
         * 
         * @return builder
         * 
         */
        public Builder isRecursive(Boolean isRecursive) {
            return isRecursive(Output.of(isRecursive));
        }

        /**
         * @param itemPath The relative path of the item (file/folder/repository). Leave unset for repository.
         * 
         * @return builder
         * 
         */
        public Builder itemPath(@Nullable Output<String> itemPath) {
            $.itemPath = itemPath;
            return this;
        }

        /**
         * @param itemPath The relative path of the item (file/folder/repository). Leave unset for repository.
         * 
         * @return builder
         * 
         */
        public Builder itemPath(String itemPath) {
            return itemPath(Output.of(itemPath));
        }

        /**
         * @param properties Map of key and list of values.
         * 
         * @return builder
         * 
         */
        public Builder properties(@Nullable Output<Map<String,List<String>>> properties) {
            $.properties = properties;
            return this;
        }

        /**
         * @param properties Map of key and list of values.
         * 
         * @return builder
         * 
         */
        public Builder properties(Map<String,List<String>> properties) {
            return properties(Output.of(properties));
        }

        /**
         * @param repoKey Respository key.
         * 
         * @return builder
         * 
         */
        public Builder repoKey(@Nullable Output<String> repoKey) {
            $.repoKey = repoKey;
            return this;
        }

        /**
         * @param repoKey Respository key.
         * 
         * @return builder
         * 
         */
        public Builder repoKey(String repoKey) {
            return repoKey(Output.of(repoKey));
        }

        public ItemPropertiesState build() {
            return $;
        }
    }

}

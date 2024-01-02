// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLocalCargoRepositoryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLocalCargoRepositoryArgs Empty = new GetLocalCargoRepositoryArgs();

    /**
     * Cargo client does not send credentials when performing download and search for crates.
     * Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous
     * access option. Default value is `false`.
     * 
     */
    @Import(name="anonymousAccess")
    private @Nullable Output<Boolean> anonymousAccess;

    /**
     * @return Cargo client does not send credentials when performing download and search for crates.
     * Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous
     * access option. Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> anonymousAccess() {
        return Optional.ofNullable(this.anonymousAccess);
    }

    @Import(name="archiveBrowsingEnabled")
    private @Nullable Output<Boolean> archiveBrowsingEnabled;

    public Optional<Output<Boolean>> archiveBrowsingEnabled() {
        return Optional.ofNullable(this.archiveBrowsingEnabled);
    }

    @Import(name="blackedOut")
    private @Nullable Output<Boolean> blackedOut;

    public Optional<Output<Boolean>> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }

    @Import(name="cdnRedirect")
    private @Nullable Output<Boolean> cdnRedirect;

    public Optional<Output<Boolean>> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="downloadDirect")
    private @Nullable Output<Boolean> downloadDirect;

    public Optional<Output<Boolean>> downloadDirect() {
        return Optional.ofNullable(this.downloadDirect);
    }

    /**
     * Enable internal index support based on Cargo sparse index specifications, instead
     * of the default git index. Default value is `false`.
     * 
     */
    @Import(name="enableSparseIndex")
    private @Nullable Output<Boolean> enableSparseIndex;

    /**
     * @return Enable internal index support based on Cargo sparse index specifications, instead
     * of the default git index. Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> enableSparseIndex() {
        return Optional.ofNullable(this.enableSparseIndex);
    }

    @Import(name="excludesPattern")
    private @Nullable Output<String> excludesPattern;

    public Optional<Output<String>> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    @Import(name="includesPattern")
    private @Nullable Output<String> includesPattern;

    public Optional<Output<String>> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }

    @Import(name="indexCompressionFormats")
    private @Nullable Output<List<String>> indexCompressionFormats;

    public Optional<Output<List<String>>> indexCompressionFormats() {
        return Optional.ofNullable(this.indexCompressionFormats);
    }

    /**
     * the identity key of the repo.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return the identity key of the repo.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    @Import(name="notes")
    private @Nullable Output<String> notes;

    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
    }

    @Import(name="priorityResolution")
    private @Nullable Output<Boolean> priorityResolution;

    public Optional<Output<Boolean>> priorityResolution() {
        return Optional.ofNullable(this.priorityResolution);
    }

    @Import(name="projectEnvironments")
    private @Nullable Output<List<String>> projectEnvironments;

    public Optional<Output<List<String>>> projectEnvironments() {
        return Optional.ofNullable(this.projectEnvironments);
    }

    @Import(name="projectKey")
    private @Nullable Output<String> projectKey;

    public Optional<Output<String>> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }

    @Import(name="propertySets")
    private @Nullable Output<List<String>> propertySets;

    public Optional<Output<List<String>>> propertySets() {
        return Optional.ofNullable(this.propertySets);
    }

    @Import(name="repoLayoutRef")
    private @Nullable Output<String> repoLayoutRef;

    public Optional<Output<String>> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    @Import(name="xrayIndex")
    private @Nullable Output<Boolean> xrayIndex;

    public Optional<Output<Boolean>> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    private GetLocalCargoRepositoryArgs() {}

    private GetLocalCargoRepositoryArgs(GetLocalCargoRepositoryArgs $) {
        this.anonymousAccess = $.anonymousAccess;
        this.archiveBrowsingEnabled = $.archiveBrowsingEnabled;
        this.blackedOut = $.blackedOut;
        this.cdnRedirect = $.cdnRedirect;
        this.description = $.description;
        this.downloadDirect = $.downloadDirect;
        this.enableSparseIndex = $.enableSparseIndex;
        this.excludesPattern = $.excludesPattern;
        this.includesPattern = $.includesPattern;
        this.indexCompressionFormats = $.indexCompressionFormats;
        this.key = $.key;
        this.notes = $.notes;
        this.priorityResolution = $.priorityResolution;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.propertySets = $.propertySets;
        this.repoLayoutRef = $.repoLayoutRef;
        this.xrayIndex = $.xrayIndex;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLocalCargoRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLocalCargoRepositoryArgs $;

        public Builder() {
            $ = new GetLocalCargoRepositoryArgs();
        }

        public Builder(GetLocalCargoRepositoryArgs defaults) {
            $ = new GetLocalCargoRepositoryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param anonymousAccess Cargo client does not send credentials when performing download and search for crates.
         * Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous
         * access option. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder anonymousAccess(@Nullable Output<Boolean> anonymousAccess) {
            $.anonymousAccess = anonymousAccess;
            return this;
        }

        /**
         * @param anonymousAccess Cargo client does not send credentials when performing download and search for crates.
         * Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous
         * access option. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder anonymousAccess(Boolean anonymousAccess) {
            return anonymousAccess(Output.of(anonymousAccess));
        }

        public Builder archiveBrowsingEnabled(@Nullable Output<Boolean> archiveBrowsingEnabled) {
            $.archiveBrowsingEnabled = archiveBrowsingEnabled;
            return this;
        }

        public Builder archiveBrowsingEnabled(Boolean archiveBrowsingEnabled) {
            return archiveBrowsingEnabled(Output.of(archiveBrowsingEnabled));
        }

        public Builder blackedOut(@Nullable Output<Boolean> blackedOut) {
            $.blackedOut = blackedOut;
            return this;
        }

        public Builder blackedOut(Boolean blackedOut) {
            return blackedOut(Output.of(blackedOut));
        }

        public Builder cdnRedirect(@Nullable Output<Boolean> cdnRedirect) {
            $.cdnRedirect = cdnRedirect;
            return this;
        }

        public Builder cdnRedirect(Boolean cdnRedirect) {
            return cdnRedirect(Output.of(cdnRedirect));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder downloadDirect(@Nullable Output<Boolean> downloadDirect) {
            $.downloadDirect = downloadDirect;
            return this;
        }

        public Builder downloadDirect(Boolean downloadDirect) {
            return downloadDirect(Output.of(downloadDirect));
        }

        /**
         * @param enableSparseIndex Enable internal index support based on Cargo sparse index specifications, instead
         * of the default git index. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableSparseIndex(@Nullable Output<Boolean> enableSparseIndex) {
            $.enableSparseIndex = enableSparseIndex;
            return this;
        }

        /**
         * @param enableSparseIndex Enable internal index support based on Cargo sparse index specifications, instead
         * of the default git index. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableSparseIndex(Boolean enableSparseIndex) {
            return enableSparseIndex(Output.of(enableSparseIndex));
        }

        public Builder excludesPattern(@Nullable Output<String> excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        public Builder excludesPattern(String excludesPattern) {
            return excludesPattern(Output.of(excludesPattern));
        }

        public Builder includesPattern(@Nullable Output<String> includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        public Builder includesPattern(String includesPattern) {
            return includesPattern(Output.of(includesPattern));
        }

        public Builder indexCompressionFormats(@Nullable Output<List<String>> indexCompressionFormats) {
            $.indexCompressionFormats = indexCompressionFormats;
            return this;
        }

        public Builder indexCompressionFormats(List<String> indexCompressionFormats) {
            return indexCompressionFormats(Output.of(indexCompressionFormats));
        }

        public Builder indexCompressionFormats(String... indexCompressionFormats) {
            return indexCompressionFormats(List.of(indexCompressionFormats));
        }

        /**
         * @param key the identity key of the repo.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key the identity key of the repo.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        public Builder notes(String notes) {
            return notes(Output.of(notes));
        }

        public Builder priorityResolution(@Nullable Output<Boolean> priorityResolution) {
            $.priorityResolution = priorityResolution;
            return this;
        }

        public Builder priorityResolution(Boolean priorityResolution) {
            return priorityResolution(Output.of(priorityResolution));
        }

        public Builder projectEnvironments(@Nullable Output<List<String>> projectEnvironments) {
            $.projectEnvironments = projectEnvironments;
            return this;
        }

        public Builder projectEnvironments(List<String> projectEnvironments) {
            return projectEnvironments(Output.of(projectEnvironments));
        }

        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }

        public Builder projectKey(@Nullable Output<String> projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        public Builder projectKey(String projectKey) {
            return projectKey(Output.of(projectKey));
        }

        public Builder propertySets(@Nullable Output<List<String>> propertySets) {
            $.propertySets = propertySets;
            return this;
        }

        public Builder propertySets(List<String> propertySets) {
            return propertySets(Output.of(propertySets));
        }

        public Builder propertySets(String... propertySets) {
            return propertySets(List.of(propertySets));
        }

        public Builder repoLayoutRef(@Nullable Output<String> repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        public Builder repoLayoutRef(String repoLayoutRef) {
            return repoLayoutRef(Output.of(repoLayoutRef));
        }

        public Builder xrayIndex(@Nullable Output<Boolean> xrayIndex) {
            $.xrayIndex = xrayIndex;
            return this;
        }

        public Builder xrayIndex(Boolean xrayIndex) {
            return xrayIndex(Output.of(xrayIndex));
        }

        public GetLocalCargoRepositoryArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("GetLocalCargoRepositoryArgs", "key");
            }
            return $;
        }
    }

}

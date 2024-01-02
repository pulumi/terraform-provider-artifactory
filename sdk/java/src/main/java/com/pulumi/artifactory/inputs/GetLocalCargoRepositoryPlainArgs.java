// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLocalCargoRepositoryPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLocalCargoRepositoryPlainArgs Empty = new GetLocalCargoRepositoryPlainArgs();

    /**
     * Cargo client does not send credentials when performing download and search for crates.
     * Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous
     * access option. Default value is `false`.
     * 
     */
    @Import(name="anonymousAccess")
    private @Nullable Boolean anonymousAccess;

    /**
     * @return Cargo client does not send credentials when performing download and search for crates.
     * Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous
     * access option. Default value is `false`.
     * 
     */
    public Optional<Boolean> anonymousAccess() {
        return Optional.ofNullable(this.anonymousAccess);
    }

    @Import(name="archiveBrowsingEnabled")
    private @Nullable Boolean archiveBrowsingEnabled;

    public Optional<Boolean> archiveBrowsingEnabled() {
        return Optional.ofNullable(this.archiveBrowsingEnabled);
    }

    @Import(name="blackedOut")
    private @Nullable Boolean blackedOut;

    public Optional<Boolean> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }

    @Import(name="cdnRedirect")
    private @Nullable Boolean cdnRedirect;

    public Optional<Boolean> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }

    @Import(name="description")
    private @Nullable String description;

    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="downloadDirect")
    private @Nullable Boolean downloadDirect;

    public Optional<Boolean> downloadDirect() {
        return Optional.ofNullable(this.downloadDirect);
    }

    /**
     * Enable internal index support based on Cargo sparse index specifications, instead
     * of the default git index. Default value is `false`.
     * 
     */
    @Import(name="enableSparseIndex")
    private @Nullable Boolean enableSparseIndex;

    /**
     * @return Enable internal index support based on Cargo sparse index specifications, instead
     * of the default git index. Default value is `false`.
     * 
     */
    public Optional<Boolean> enableSparseIndex() {
        return Optional.ofNullable(this.enableSparseIndex);
    }

    @Import(name="excludesPattern")
    private @Nullable String excludesPattern;

    public Optional<String> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    @Import(name="includesPattern")
    private @Nullable String includesPattern;

    public Optional<String> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }

    @Import(name="indexCompressionFormats")
    private @Nullable List<String> indexCompressionFormats;

    public Optional<List<String>> indexCompressionFormats() {
        return Optional.ofNullable(this.indexCompressionFormats);
    }

    /**
     * the identity key of the repo.
     * 
     */
    @Import(name="key", required=true)
    private String key;

    /**
     * @return the identity key of the repo.
     * 
     */
    public String key() {
        return this.key;
    }

    @Import(name="notes")
    private @Nullable String notes;

    public Optional<String> notes() {
        return Optional.ofNullable(this.notes);
    }

    @Import(name="priorityResolution")
    private @Nullable Boolean priorityResolution;

    public Optional<Boolean> priorityResolution() {
        return Optional.ofNullable(this.priorityResolution);
    }

    @Import(name="projectEnvironments")
    private @Nullable List<String> projectEnvironments;

    public Optional<List<String>> projectEnvironments() {
        return Optional.ofNullable(this.projectEnvironments);
    }

    @Import(name="projectKey")
    private @Nullable String projectKey;

    public Optional<String> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }

    @Import(name="propertySets")
    private @Nullable List<String> propertySets;

    public Optional<List<String>> propertySets() {
        return Optional.ofNullable(this.propertySets);
    }

    @Import(name="repoLayoutRef")
    private @Nullable String repoLayoutRef;

    public Optional<String> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    @Import(name="xrayIndex")
    private @Nullable Boolean xrayIndex;

    public Optional<Boolean> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    private GetLocalCargoRepositoryPlainArgs() {}

    private GetLocalCargoRepositoryPlainArgs(GetLocalCargoRepositoryPlainArgs $) {
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
    public static Builder builder(GetLocalCargoRepositoryPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLocalCargoRepositoryPlainArgs $;

        public Builder() {
            $ = new GetLocalCargoRepositoryPlainArgs();
        }

        public Builder(GetLocalCargoRepositoryPlainArgs defaults) {
            $ = new GetLocalCargoRepositoryPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param anonymousAccess Cargo client does not send credentials when performing download and search for crates.
         * Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous
         * access option. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder anonymousAccess(@Nullable Boolean anonymousAccess) {
            $.anonymousAccess = anonymousAccess;
            return this;
        }

        public Builder archiveBrowsingEnabled(@Nullable Boolean archiveBrowsingEnabled) {
            $.archiveBrowsingEnabled = archiveBrowsingEnabled;
            return this;
        }

        public Builder blackedOut(@Nullable Boolean blackedOut) {
            $.blackedOut = blackedOut;
            return this;
        }

        public Builder cdnRedirect(@Nullable Boolean cdnRedirect) {
            $.cdnRedirect = cdnRedirect;
            return this;
        }

        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        public Builder downloadDirect(@Nullable Boolean downloadDirect) {
            $.downloadDirect = downloadDirect;
            return this;
        }

        /**
         * @param enableSparseIndex Enable internal index support based on Cargo sparse index specifications, instead
         * of the default git index. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableSparseIndex(@Nullable Boolean enableSparseIndex) {
            $.enableSparseIndex = enableSparseIndex;
            return this;
        }

        public Builder excludesPattern(@Nullable String excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        public Builder includesPattern(@Nullable String includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        public Builder indexCompressionFormats(@Nullable List<String> indexCompressionFormats) {
            $.indexCompressionFormats = indexCompressionFormats;
            return this;
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
        public Builder key(String key) {
            $.key = key;
            return this;
        }

        public Builder notes(@Nullable String notes) {
            $.notes = notes;
            return this;
        }

        public Builder priorityResolution(@Nullable Boolean priorityResolution) {
            $.priorityResolution = priorityResolution;
            return this;
        }

        public Builder projectEnvironments(@Nullable List<String> projectEnvironments) {
            $.projectEnvironments = projectEnvironments;
            return this;
        }

        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }

        public Builder projectKey(@Nullable String projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        public Builder propertySets(@Nullable List<String> propertySets) {
            $.propertySets = propertySets;
            return this;
        }

        public Builder propertySets(String... propertySets) {
            return propertySets(List.of(propertySets));
        }

        public Builder repoLayoutRef(@Nullable String repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        public Builder xrayIndex(@Nullable Boolean xrayIndex) {
            $.xrayIndex = xrayIndex;
            return this;
        }

        public GetLocalCargoRepositoryPlainArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("GetLocalCargoRepositoryPlainArgs", "key");
            }
            return $;
        }
    }

}

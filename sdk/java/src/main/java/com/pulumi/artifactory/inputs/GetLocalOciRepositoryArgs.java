// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLocalOciRepositoryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLocalOciRepositoryArgs Empty = new GetLocalOciRepositoryArgs();

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

    /**
     * The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
     * 
     */
    @Import(name="maxUniqueTags")
    private @Nullable Output<Integer> maxUniqueTags;

    /**
     * @return The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
     * 
     */
    public Optional<Output<Integer>> maxUniqueTags() {
        return Optional.ofNullable(this.maxUniqueTags);
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

    /**
     * If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
     * 
     */
    @Import(name="tagRetention")
    private @Nullable Output<Integer> tagRetention;

    /**
     * @return If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
     * 
     */
    public Optional<Output<Integer>> tagRetention() {
        return Optional.ofNullable(this.tagRetention);
    }

    @Import(name="xrayIndex")
    private @Nullable Output<Boolean> xrayIndex;

    public Optional<Output<Boolean>> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    private GetLocalOciRepositoryArgs() {}

    private GetLocalOciRepositoryArgs(GetLocalOciRepositoryArgs $) {
        this.archiveBrowsingEnabled = $.archiveBrowsingEnabled;
        this.blackedOut = $.blackedOut;
        this.cdnRedirect = $.cdnRedirect;
        this.description = $.description;
        this.downloadDirect = $.downloadDirect;
        this.excludesPattern = $.excludesPattern;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.maxUniqueTags = $.maxUniqueTags;
        this.notes = $.notes;
        this.priorityResolution = $.priorityResolution;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.propertySets = $.propertySets;
        this.repoLayoutRef = $.repoLayoutRef;
        this.tagRetention = $.tagRetention;
        this.xrayIndex = $.xrayIndex;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLocalOciRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLocalOciRepositoryArgs $;

        public Builder() {
            $ = new GetLocalOciRepositoryArgs();
        }

        public Builder(GetLocalOciRepositoryArgs defaults) {
            $ = new GetLocalOciRepositoryArgs(Objects.requireNonNull(defaults));
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

        /**
         * @param maxUniqueTags The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
         * 
         * @return builder
         * 
         */
        public Builder maxUniqueTags(@Nullable Output<Integer> maxUniqueTags) {
            $.maxUniqueTags = maxUniqueTags;
            return this;
        }

        /**
         * @param maxUniqueTags The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
         * 
         * @return builder
         * 
         */
        public Builder maxUniqueTags(Integer maxUniqueTags) {
            return maxUniqueTags(Output.of(maxUniqueTags));
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

        /**
         * @param tagRetention If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
         * 
         * @return builder
         * 
         */
        public Builder tagRetention(@Nullable Output<Integer> tagRetention) {
            $.tagRetention = tagRetention;
            return this;
        }

        /**
         * @param tagRetention If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
         * 
         * @return builder
         * 
         */
        public Builder tagRetention(Integer tagRetention) {
            return tagRetention(Output.of(tagRetention));
        }

        public Builder xrayIndex(@Nullable Output<Boolean> xrayIndex) {
            $.xrayIndex = xrayIndex;
            return this;
        }

        public Builder xrayIndex(Boolean xrayIndex) {
            return xrayIndex(Output.of(xrayIndex));
        }

        public GetLocalOciRepositoryArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("GetLocalOciRepositoryArgs", "key");
            }
            return $;
        }
    }

}

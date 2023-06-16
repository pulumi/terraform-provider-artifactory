// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.GetFederatedRpmRepositoryMember;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFederatedRpmRepositoryPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFederatedRpmRepositoryPlainArgs Empty = new GetFederatedRpmRepositoryPlainArgs();

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

    @Import(name="calculateYumMetadata")
    private @Nullable Boolean calculateYumMetadata;

    public Optional<Boolean> calculateYumMetadata() {
        return Optional.ofNullable(this.calculateYumMetadata);
    }

    @Import(name="cdnRedirect")
    private @Nullable Boolean cdnRedirect;

    public Optional<Boolean> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }

    @Import(name="cleanupOnDelete")
    private @Nullable Boolean cleanupOnDelete;

    public Optional<Boolean> cleanupOnDelete() {
        return Optional.ofNullable(this.cleanupOnDelete);
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

    @Import(name="enableFileListsIndexing")
    private @Nullable Boolean enableFileListsIndexing;

    public Optional<Boolean> enableFileListsIndexing() {
        return Optional.ofNullable(this.enableFileListsIndexing);
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

    @Import(name="key", required=true)
    private String key;

    public String key() {
        return this.key;
    }

    @Import(name="members")
    private @Nullable List<GetFederatedRpmRepositoryMember> members;

    public Optional<List<GetFederatedRpmRepositoryMember>> members() {
        return Optional.ofNullable(this.members);
    }

    @Import(name="notes")
    private @Nullable String notes;

    public Optional<String> notes() {
        return Optional.ofNullable(this.notes);
    }

    @Import(name="primaryKeypairRef")
    private @Nullable String primaryKeypairRef;

    public Optional<String> primaryKeypairRef() {
        return Optional.ofNullable(this.primaryKeypairRef);
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

    @Import(name="secondaryKeypairRef")
    private @Nullable String secondaryKeypairRef;

    public Optional<String> secondaryKeypairRef() {
        return Optional.ofNullable(this.secondaryKeypairRef);
    }

    @Import(name="xrayIndex")
    private @Nullable Boolean xrayIndex;

    public Optional<Boolean> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    @Import(name="yumGroupFileNames")
    private @Nullable String yumGroupFileNames;

    public Optional<String> yumGroupFileNames() {
        return Optional.ofNullable(this.yumGroupFileNames);
    }

    @Import(name="yumRootDepth")
    private @Nullable Integer yumRootDepth;

    public Optional<Integer> yumRootDepth() {
        return Optional.ofNullable(this.yumRootDepth);
    }

    private GetFederatedRpmRepositoryPlainArgs() {}

    private GetFederatedRpmRepositoryPlainArgs(GetFederatedRpmRepositoryPlainArgs $) {
        this.archiveBrowsingEnabled = $.archiveBrowsingEnabled;
        this.blackedOut = $.blackedOut;
        this.calculateYumMetadata = $.calculateYumMetadata;
        this.cdnRedirect = $.cdnRedirect;
        this.cleanupOnDelete = $.cleanupOnDelete;
        this.description = $.description;
        this.downloadDirect = $.downloadDirect;
        this.enableFileListsIndexing = $.enableFileListsIndexing;
        this.excludesPattern = $.excludesPattern;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.members = $.members;
        this.notes = $.notes;
        this.primaryKeypairRef = $.primaryKeypairRef;
        this.priorityResolution = $.priorityResolution;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.propertySets = $.propertySets;
        this.repoLayoutRef = $.repoLayoutRef;
        this.secondaryKeypairRef = $.secondaryKeypairRef;
        this.xrayIndex = $.xrayIndex;
        this.yumGroupFileNames = $.yumGroupFileNames;
        this.yumRootDepth = $.yumRootDepth;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFederatedRpmRepositoryPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFederatedRpmRepositoryPlainArgs $;

        public Builder() {
            $ = new GetFederatedRpmRepositoryPlainArgs();
        }

        public Builder(GetFederatedRpmRepositoryPlainArgs defaults) {
            $ = new GetFederatedRpmRepositoryPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder archiveBrowsingEnabled(@Nullable Boolean archiveBrowsingEnabled) {
            $.archiveBrowsingEnabled = archiveBrowsingEnabled;
            return this;
        }

        public Builder blackedOut(@Nullable Boolean blackedOut) {
            $.blackedOut = blackedOut;
            return this;
        }

        public Builder calculateYumMetadata(@Nullable Boolean calculateYumMetadata) {
            $.calculateYumMetadata = calculateYumMetadata;
            return this;
        }

        public Builder cdnRedirect(@Nullable Boolean cdnRedirect) {
            $.cdnRedirect = cdnRedirect;
            return this;
        }

        public Builder cleanupOnDelete(@Nullable Boolean cleanupOnDelete) {
            $.cleanupOnDelete = cleanupOnDelete;
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

        public Builder enableFileListsIndexing(@Nullable Boolean enableFileListsIndexing) {
            $.enableFileListsIndexing = enableFileListsIndexing;
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

        public Builder key(String key) {
            $.key = key;
            return this;
        }

        public Builder members(@Nullable List<GetFederatedRpmRepositoryMember> members) {
            $.members = members;
            return this;
        }

        public Builder members(GetFederatedRpmRepositoryMember... members) {
            return members(List.of(members));
        }

        public Builder notes(@Nullable String notes) {
            $.notes = notes;
            return this;
        }

        public Builder primaryKeypairRef(@Nullable String primaryKeypairRef) {
            $.primaryKeypairRef = primaryKeypairRef;
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

        public Builder secondaryKeypairRef(@Nullable String secondaryKeypairRef) {
            $.secondaryKeypairRef = secondaryKeypairRef;
            return this;
        }

        public Builder xrayIndex(@Nullable Boolean xrayIndex) {
            $.xrayIndex = xrayIndex;
            return this;
        }

        public Builder yumGroupFileNames(@Nullable String yumGroupFileNames) {
            $.yumGroupFileNames = yumGroupFileNames;
            return this;
        }

        public Builder yumRootDepth(@Nullable Integer yumRootDepth) {
            $.yumRootDepth = yumRootDepth;
            return this;
        }

        public GetFederatedRpmRepositoryPlainArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}

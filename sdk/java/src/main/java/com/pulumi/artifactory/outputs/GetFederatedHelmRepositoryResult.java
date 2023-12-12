// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.artifactory.outputs.GetFederatedHelmRepositoryMember;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetFederatedHelmRepositoryResult {
    private @Nullable Boolean archiveBrowsingEnabled;
    private @Nullable Boolean blackedOut;
    private @Nullable Boolean cdnRedirect;
    private @Nullable Boolean cleanupOnDelete;
    private @Nullable String description;
    /**
     * @return When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
     * 
     */
    private @Nullable Boolean disableProxy;
    private @Nullable Boolean downloadDirect;
    private @Nullable String excludesPattern;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String includesPattern;
    private String key;
    /**
     * @return The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     * 
     */
    private @Nullable List<GetFederatedHelmRepositoryMember> members;
    private @Nullable String notes;
    private String packageType;
    private @Nullable Boolean priorityResolution;
    private List<String> projectEnvironments;
    private @Nullable String projectKey;
    private @Nullable List<String> propertySets;
    /**
     * @return Proxy key from Artifactory Proxies settings.
     * 
     */
    private @Nullable String proxy;
    private @Nullable String repoLayoutRef;
    private @Nullable Boolean xrayIndex;

    private GetFederatedHelmRepositoryResult() {}
    public Optional<Boolean> archiveBrowsingEnabled() {
        return Optional.ofNullable(this.archiveBrowsingEnabled);
    }
    public Optional<Boolean> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }
    public Optional<Boolean> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }
    public Optional<Boolean> cleanupOnDelete() {
        return Optional.ofNullable(this.cleanupOnDelete);
    }
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
     * 
     */
    public Optional<Boolean> disableProxy() {
        return Optional.ofNullable(this.disableProxy);
    }
    public Optional<Boolean> downloadDirect() {
        return Optional.ofNullable(this.downloadDirect);
    }
    public Optional<String> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }
    public String key() {
        return this.key;
    }
    /**
     * @return The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     * 
     */
    public List<GetFederatedHelmRepositoryMember> members() {
        return this.members == null ? List.of() : this.members;
    }
    public Optional<String> notes() {
        return Optional.ofNullable(this.notes);
    }
    public String packageType() {
        return this.packageType;
    }
    public Optional<Boolean> priorityResolution() {
        return Optional.ofNullable(this.priorityResolution);
    }
    public List<String> projectEnvironments() {
        return this.projectEnvironments;
    }
    public Optional<String> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }
    public List<String> propertySets() {
        return this.propertySets == null ? List.of() : this.propertySets;
    }
    /**
     * @return Proxy key from Artifactory Proxies settings.
     * 
     */
    public Optional<String> proxy() {
        return Optional.ofNullable(this.proxy);
    }
    public Optional<String> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }
    public Optional<Boolean> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetFederatedHelmRepositoryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean archiveBrowsingEnabled;
        private @Nullable Boolean blackedOut;
        private @Nullable Boolean cdnRedirect;
        private @Nullable Boolean cleanupOnDelete;
        private @Nullable String description;
        private @Nullable Boolean disableProxy;
        private @Nullable Boolean downloadDirect;
        private @Nullable String excludesPattern;
        private String id;
        private @Nullable String includesPattern;
        private String key;
        private @Nullable List<GetFederatedHelmRepositoryMember> members;
        private @Nullable String notes;
        private String packageType;
        private @Nullable Boolean priorityResolution;
        private List<String> projectEnvironments;
        private @Nullable String projectKey;
        private @Nullable List<String> propertySets;
        private @Nullable String proxy;
        private @Nullable String repoLayoutRef;
        private @Nullable Boolean xrayIndex;
        public Builder() {}
        public Builder(GetFederatedHelmRepositoryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.archiveBrowsingEnabled = defaults.archiveBrowsingEnabled;
    	      this.blackedOut = defaults.blackedOut;
    	      this.cdnRedirect = defaults.cdnRedirect;
    	      this.cleanupOnDelete = defaults.cleanupOnDelete;
    	      this.description = defaults.description;
    	      this.disableProxy = defaults.disableProxy;
    	      this.downloadDirect = defaults.downloadDirect;
    	      this.excludesPattern = defaults.excludesPattern;
    	      this.id = defaults.id;
    	      this.includesPattern = defaults.includesPattern;
    	      this.key = defaults.key;
    	      this.members = defaults.members;
    	      this.notes = defaults.notes;
    	      this.packageType = defaults.packageType;
    	      this.priorityResolution = defaults.priorityResolution;
    	      this.projectEnvironments = defaults.projectEnvironments;
    	      this.projectKey = defaults.projectKey;
    	      this.propertySets = defaults.propertySets;
    	      this.proxy = defaults.proxy;
    	      this.repoLayoutRef = defaults.repoLayoutRef;
    	      this.xrayIndex = defaults.xrayIndex;
        }

        @CustomType.Setter
        public Builder archiveBrowsingEnabled(@Nullable Boolean archiveBrowsingEnabled) {
            this.archiveBrowsingEnabled = archiveBrowsingEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder blackedOut(@Nullable Boolean blackedOut) {
            this.blackedOut = blackedOut;
            return this;
        }
        @CustomType.Setter
        public Builder cdnRedirect(@Nullable Boolean cdnRedirect) {
            this.cdnRedirect = cdnRedirect;
            return this;
        }
        @CustomType.Setter
        public Builder cleanupOnDelete(@Nullable Boolean cleanupOnDelete) {
            this.cleanupOnDelete = cleanupOnDelete;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder disableProxy(@Nullable Boolean disableProxy) {
            this.disableProxy = disableProxy;
            return this;
        }
        @CustomType.Setter
        public Builder downloadDirect(@Nullable Boolean downloadDirect) {
            this.downloadDirect = downloadDirect;
            return this;
        }
        @CustomType.Setter
        public Builder excludesPattern(@Nullable String excludesPattern) {
            this.excludesPattern = excludesPattern;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder includesPattern(@Nullable String includesPattern) {
            this.includesPattern = includesPattern;
            return this;
        }
        @CustomType.Setter
        public Builder key(String key) {
            this.key = Objects.requireNonNull(key);
            return this;
        }
        @CustomType.Setter
        public Builder members(@Nullable List<GetFederatedHelmRepositoryMember> members) {
            this.members = members;
            return this;
        }
        public Builder members(GetFederatedHelmRepositoryMember... members) {
            return members(List.of(members));
        }
        @CustomType.Setter
        public Builder notes(@Nullable String notes) {
            this.notes = notes;
            return this;
        }
        @CustomType.Setter
        public Builder packageType(String packageType) {
            this.packageType = Objects.requireNonNull(packageType);
            return this;
        }
        @CustomType.Setter
        public Builder priorityResolution(@Nullable Boolean priorityResolution) {
            this.priorityResolution = priorityResolution;
            return this;
        }
        @CustomType.Setter
        public Builder projectEnvironments(List<String> projectEnvironments) {
            this.projectEnvironments = Objects.requireNonNull(projectEnvironments);
            return this;
        }
        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }
        @CustomType.Setter
        public Builder projectKey(@Nullable String projectKey) {
            this.projectKey = projectKey;
            return this;
        }
        @CustomType.Setter
        public Builder propertySets(@Nullable List<String> propertySets) {
            this.propertySets = propertySets;
            return this;
        }
        public Builder propertySets(String... propertySets) {
            return propertySets(List.of(propertySets));
        }
        @CustomType.Setter
        public Builder proxy(@Nullable String proxy) {
            this.proxy = proxy;
            return this;
        }
        @CustomType.Setter
        public Builder repoLayoutRef(@Nullable String repoLayoutRef) {
            this.repoLayoutRef = repoLayoutRef;
            return this;
        }
        @CustomType.Setter
        public Builder xrayIndex(@Nullable Boolean xrayIndex) {
            this.xrayIndex = xrayIndex;
            return this;
        }
        public GetFederatedHelmRepositoryResult build() {
            final var _resultValue = new GetFederatedHelmRepositoryResult();
            _resultValue.archiveBrowsingEnabled = archiveBrowsingEnabled;
            _resultValue.blackedOut = blackedOut;
            _resultValue.cdnRedirect = cdnRedirect;
            _resultValue.cleanupOnDelete = cleanupOnDelete;
            _resultValue.description = description;
            _resultValue.disableProxy = disableProxy;
            _resultValue.downloadDirect = downloadDirect;
            _resultValue.excludesPattern = excludesPattern;
            _resultValue.id = id;
            _resultValue.includesPattern = includesPattern;
            _resultValue.key = key;
            _resultValue.members = members;
            _resultValue.notes = notes;
            _resultValue.packageType = packageType;
            _resultValue.priorityResolution = priorityResolution;
            _resultValue.projectEnvironments = projectEnvironments;
            _resultValue.projectKey = projectKey;
            _resultValue.propertySets = propertySets;
            _resultValue.proxy = proxy;
            _resultValue.repoLayoutRef = repoLayoutRef;
            _resultValue.xrayIndex = xrayIndex;
            return _resultValue;
        }
    }
}

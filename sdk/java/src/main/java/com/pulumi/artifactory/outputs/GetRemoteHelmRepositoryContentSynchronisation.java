// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRemoteHelmRepositoryContentSynchronisation {
    /**
     * @return If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is &#39;false&#39;.
     * 
     */
    private @Nullable Boolean enabled;
    /**
     * @return If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is &#39;false&#39;.
     * 
     */
    private @Nullable Boolean propertiesEnabled;
    /**
     * @return If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is &#39;false&#39;
     * 
     */
    private @Nullable Boolean sourceOriginAbsenceDetection;
    /**
     * @return If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is &#39;false&#39;.
     * 
     */
    private @Nullable Boolean statisticsEnabled;

    private GetRemoteHelmRepositoryContentSynchronisation() {}
    /**
     * @return If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is &#39;false&#39;.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    /**
     * @return If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is &#39;false&#39;.
     * 
     */
    public Optional<Boolean> propertiesEnabled() {
        return Optional.ofNullable(this.propertiesEnabled);
    }
    /**
     * @return If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is &#39;false&#39;
     * 
     */
    public Optional<Boolean> sourceOriginAbsenceDetection() {
        return Optional.ofNullable(this.sourceOriginAbsenceDetection);
    }
    /**
     * @return If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is &#39;false&#39;.
     * 
     */
    public Optional<Boolean> statisticsEnabled() {
        return Optional.ofNullable(this.statisticsEnabled);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRemoteHelmRepositoryContentSynchronisation defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enabled;
        private @Nullable Boolean propertiesEnabled;
        private @Nullable Boolean sourceOriginAbsenceDetection;
        private @Nullable Boolean statisticsEnabled;
        public Builder() {}
        public Builder(GetRemoteHelmRepositoryContentSynchronisation defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.propertiesEnabled = defaults.propertiesEnabled;
    	      this.sourceOriginAbsenceDetection = defaults.sourceOriginAbsenceDetection;
    	      this.statisticsEnabled = defaults.statisticsEnabled;
        }

        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {

            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder propertiesEnabled(@Nullable Boolean propertiesEnabled) {

            this.propertiesEnabled = propertiesEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder sourceOriginAbsenceDetection(@Nullable Boolean sourceOriginAbsenceDetection) {

            this.sourceOriginAbsenceDetection = sourceOriginAbsenceDetection;
            return this;
        }
        @CustomType.Setter
        public Builder statisticsEnabled(@Nullable Boolean statisticsEnabled) {

            this.statisticsEnabled = statisticsEnabled;
            return this;
        }
        public GetRemoteHelmRepositoryContentSynchronisation build() {
            final var _resultValue = new GetRemoteHelmRepositoryContentSynchronisation();
            _resultValue.enabled = enabled;
            _resultValue.propertiesEnabled = propertiesEnabled;
            _resultValue.sourceOriginAbsenceDetection = sourceOriginAbsenceDetection;
            _resultValue.statisticsEnabled = statisticsEnabled;
            return _resultValue;
        }
    }
}

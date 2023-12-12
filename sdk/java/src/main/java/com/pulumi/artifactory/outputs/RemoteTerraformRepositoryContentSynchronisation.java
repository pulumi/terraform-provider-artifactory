// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RemoteTerraformRepositoryContentSynchronisation {
    private @Nullable Boolean enabled;
    private @Nullable Boolean propertiesEnabled;
    private @Nullable Boolean sourceOriginAbsenceDetection;
    private @Nullable Boolean statisticsEnabled;

    private RemoteTerraformRepositoryContentSynchronisation() {}
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    public Optional<Boolean> propertiesEnabled() {
        return Optional.ofNullable(this.propertiesEnabled);
    }
    public Optional<Boolean> sourceOriginAbsenceDetection() {
        return Optional.ofNullable(this.sourceOriginAbsenceDetection);
    }
    public Optional<Boolean> statisticsEnabled() {
        return Optional.ofNullable(this.statisticsEnabled);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RemoteTerraformRepositoryContentSynchronisation defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enabled;
        private @Nullable Boolean propertiesEnabled;
        private @Nullable Boolean sourceOriginAbsenceDetection;
        private @Nullable Boolean statisticsEnabled;
        public Builder() {}
        public Builder(RemoteTerraformRepositoryContentSynchronisation defaults) {
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
        public RemoteTerraformRepositoryContentSynchronisation build() {
            final var _resultValue = new RemoteTerraformRepositoryContentSynchronisation();
            _resultValue.enabled = enabled;
            _resultValue.propertiesEnabled = propertiesEnabled;
            _resultValue.sourceOriginAbsenceDetection = sourceOriginAbsenceDetection;
            _resultValue.statisticsEnabled = statisticsEnabled;
            return _resultValue;
        }
    }
}

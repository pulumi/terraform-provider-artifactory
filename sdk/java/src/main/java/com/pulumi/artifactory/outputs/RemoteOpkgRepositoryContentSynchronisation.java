// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RemoteOpkgRepositoryContentSynchronisation {
    private final @Nullable Boolean enabled;
    private final @Nullable Boolean propertiesEnabled;
    private final @Nullable Boolean sourceOriginAbsenceDetection;
    private final @Nullable Boolean statisticsEnabled;

    @CustomType.Constructor
    private RemoteOpkgRepositoryContentSynchronisation(
        @CustomType.Parameter("enabled") @Nullable Boolean enabled,
        @CustomType.Parameter("propertiesEnabled") @Nullable Boolean propertiesEnabled,
        @CustomType.Parameter("sourceOriginAbsenceDetection") @Nullable Boolean sourceOriginAbsenceDetection,
        @CustomType.Parameter("statisticsEnabled") @Nullable Boolean statisticsEnabled) {
        this.enabled = enabled;
        this.propertiesEnabled = propertiesEnabled;
        this.sourceOriginAbsenceDetection = sourceOriginAbsenceDetection;
        this.statisticsEnabled = statisticsEnabled;
    }

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

    public static Builder builder(RemoteOpkgRepositoryContentSynchronisation defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Boolean enabled;
        private @Nullable Boolean propertiesEnabled;
        private @Nullable Boolean sourceOriginAbsenceDetection;
        private @Nullable Boolean statisticsEnabled;

        public Builder() {
    	      // Empty
        }

        public Builder(RemoteOpkgRepositoryContentSynchronisation defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.propertiesEnabled = defaults.propertiesEnabled;
    	      this.sourceOriginAbsenceDetection = defaults.sourceOriginAbsenceDetection;
    	      this.statisticsEnabled = defaults.statisticsEnabled;
        }

        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        public Builder propertiesEnabled(@Nullable Boolean propertiesEnabled) {
            this.propertiesEnabled = propertiesEnabled;
            return this;
        }
        public Builder sourceOriginAbsenceDetection(@Nullable Boolean sourceOriginAbsenceDetection) {
            this.sourceOriginAbsenceDetection = sourceOriginAbsenceDetection;
            return this;
        }
        public Builder statisticsEnabled(@Nullable Boolean statisticsEnabled) {
            this.statisticsEnabled = statisticsEnabled;
            return this;
        }        public RemoteOpkgRepositoryContentSynchronisation build() {
            return new RemoteOpkgRepositoryContentSynchronisation(enabled, propertiesEnabled, sourceOriginAbsenceDetection, statisticsEnabled);
        }
    }
}

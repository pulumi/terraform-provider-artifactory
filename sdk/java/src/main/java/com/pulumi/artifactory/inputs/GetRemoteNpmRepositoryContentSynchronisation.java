// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRemoteNpmRepositoryContentSynchronisation extends com.pulumi.resources.InvokeArgs {

    public static final GetRemoteNpmRepositoryContentSynchronisation Empty = new GetRemoteNpmRepositoryContentSynchronisation();

    @Import(name="enabled")
    private @Nullable Boolean enabled;

    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="propertiesEnabled")
    private @Nullable Boolean propertiesEnabled;

    public Optional<Boolean> propertiesEnabled() {
        return Optional.ofNullable(this.propertiesEnabled);
    }

    @Import(name="sourceOriginAbsenceDetection")
    private @Nullable Boolean sourceOriginAbsenceDetection;

    public Optional<Boolean> sourceOriginAbsenceDetection() {
        return Optional.ofNullable(this.sourceOriginAbsenceDetection);
    }

    @Import(name="statisticsEnabled")
    private @Nullable Boolean statisticsEnabled;

    public Optional<Boolean> statisticsEnabled() {
        return Optional.ofNullable(this.statisticsEnabled);
    }

    private GetRemoteNpmRepositoryContentSynchronisation() {}

    private GetRemoteNpmRepositoryContentSynchronisation(GetRemoteNpmRepositoryContentSynchronisation $) {
        this.enabled = $.enabled;
        this.propertiesEnabled = $.propertiesEnabled;
        this.sourceOriginAbsenceDetection = $.sourceOriginAbsenceDetection;
        this.statisticsEnabled = $.statisticsEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRemoteNpmRepositoryContentSynchronisation defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRemoteNpmRepositoryContentSynchronisation $;

        public Builder() {
            $ = new GetRemoteNpmRepositoryContentSynchronisation();
        }

        public Builder(GetRemoteNpmRepositoryContentSynchronisation defaults) {
            $ = new GetRemoteNpmRepositoryContentSynchronisation(Objects.requireNonNull(defaults));
        }

        public Builder enabled(@Nullable Boolean enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder propertiesEnabled(@Nullable Boolean propertiesEnabled) {
            $.propertiesEnabled = propertiesEnabled;
            return this;
        }

        public Builder sourceOriginAbsenceDetection(@Nullable Boolean sourceOriginAbsenceDetection) {
            $.sourceOriginAbsenceDetection = sourceOriginAbsenceDetection;
            return this;
        }

        public Builder statisticsEnabled(@Nullable Boolean statisticsEnabled) {
            $.statisticsEnabled = statisticsEnabled;
            return this;
        }

        public GetRemoteNpmRepositoryContentSynchronisation build() {
            return $;
        }
    }

}

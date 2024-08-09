// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;





import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.Optional;
import java.util.stream.Collectors;
import javax.annotation.Nullable;
import com.pulumi.core.internal.Environment;
import com.pulumi.deployment.InvokeOptions;

public class Utilities {

	public static Optional<java.lang.String> getEnv(java.lang.String... names) {
        for (var n : names) {
            var value = Environment.getEnvironmentVariable(n);
            if (value.isValue()) {
                return Optional.of(value.value());
            }
        }
        return Optional.empty();
    }

	public static Optional<java.lang.Boolean> getEnvBoolean(java.lang.String... names) {
        for (var n : names) {
            var value = Environment.getBooleanEnvironmentVariable(n);
            if (value.isValue()) {
                return Optional.of(value.value());
            }
        }
        return Optional.empty();
	}

	public static Optional<java.lang.Integer> getEnvInteger(java.lang.String... names) {
        for (var n : names) {
            var value = Environment.getIntegerEnvironmentVariable(n);
            if (value.isValue()) {
                return Optional.of(value.value());
            }
        }
        return Optional.empty();
	}

	public static Optional<java.lang.Double> getEnvDouble(java.lang.String... names) {
        for (var n : names) {
            var value = Environment.getDoubleEnvironmentVariable(n);
            if (value.isValue()) {
                return Optional.of(value.value());
            }
        }
        return Optional.empty();
	}

	public static InvokeOptions withVersion(@Nullable InvokeOptions options) {
            if (options != null && options.getVersion().isPresent()) {
                return options;
            }
            return new InvokeOptions(
                options == null ? null : options.getParent().orElse(null),
                options == null ? null : options.getProvider().orElse(null),
                getVersion()
            );
        }

    private static final java.lang.String version;
    public static java.lang.String getVersion() {
        return version;
    }

    static {
        var resourceName = "com/pulumi/artifactory/version.txt";
        var versionFile = Utilities.class.getClassLoader().getResourceAsStream(resourceName);
        if (versionFile == null) {
            throw new IllegalStateException(
                    java.lang.String.format("expected resource '%s' on Classpath, not found", resourceName)
            );
        }
        version = new BufferedReader(new InputStreamReader(versionFile))
                .lines()
                .collect(Collectors.joining("\n"))
                .trim();
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ArtifactCustomWebhookHandlerArgs extends com.pulumi.resources.ResourceArgs {

    public static final ArtifactCustomWebhookHandlerArgs Empty = new ArtifactCustomWebhookHandlerArgs();

    /**
     * HTTP headers you wish to use to invoke the Webhook, comprise key/value pair.
     * 
     */
    @Import(name="httpHeaders")
    private @Nullable Output<Map<String,String>> httpHeaders;

    /**
     * @return HTTP headers you wish to use to invoke the Webhook, comprise key/value pair.
     * 
     */
    public Optional<Output<Map<String,String>>> httpHeaders() {
        return Optional.ofNullable(this.httpHeaders);
    }

    @Import(name="payload")
    private @Nullable Output<String> payload;

    public Optional<Output<String>> payload() {
        return Optional.ofNullable(this.payload);
    }

    /**
     * Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
     * 
     */
    @Import(name="proxy")
    private @Nullable Output<String> proxy;

    /**
     * @return Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
     * 
     */
    public Optional<Output<String>> proxy() {
        return Optional.ofNullable(this.proxy);
    }

    /**
     * Defines a set of sensitive values (such as, tokens and passwords) that can be injected in the headers and/or payload.Secrets’ values are encrypted. In the header/payload, the value can be invoked using the `{{.secrets.token}}` format, where token is the name provided for the secret value. Comprise key/value pair. **Note:** if multiple handlers are used, same secret name and different secret value for the same url won&#39;t work. Example:
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *     }
     * }
     * ```
     * 
     */
    @Import(name="secrets")
    private @Nullable Output<Map<String,String>> secrets;

    /**
     * @return Defines a set of sensitive values (such as, tokens and passwords) that can be injected in the headers and/or payload.Secrets’ values are encrypted. In the header/payload, the value can be invoked using the `{{.secrets.token}}` format, where token is the name provided for the secret value. Comprise key/value pair. **Note:** if multiple handlers are used, same secret name and different secret value for the same url won&#39;t work. Example:
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *     }
     * }
     * ```
     * 
     */
    public Optional<Output<Map<String,String>>> secrets() {
        return Optional.ofNullable(this.secrets);
    }

    /**
     * Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    private ArtifactCustomWebhookHandlerArgs() {}

    private ArtifactCustomWebhookHandlerArgs(ArtifactCustomWebhookHandlerArgs $) {
        this.httpHeaders = $.httpHeaders;
        this.payload = $.payload;
        this.proxy = $.proxy;
        this.secrets = $.secrets;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ArtifactCustomWebhookHandlerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ArtifactCustomWebhookHandlerArgs $;

        public Builder() {
            $ = new ArtifactCustomWebhookHandlerArgs();
        }

        public Builder(ArtifactCustomWebhookHandlerArgs defaults) {
            $ = new ArtifactCustomWebhookHandlerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param httpHeaders HTTP headers you wish to use to invoke the Webhook, comprise key/value pair.
         * 
         * @return builder
         * 
         */
        public Builder httpHeaders(@Nullable Output<Map<String,String>> httpHeaders) {
            $.httpHeaders = httpHeaders;
            return this;
        }

        /**
         * @param httpHeaders HTTP headers you wish to use to invoke the Webhook, comprise key/value pair.
         * 
         * @return builder
         * 
         */
        public Builder httpHeaders(Map<String,String> httpHeaders) {
            return httpHeaders(Output.of(httpHeaders));
        }

        public Builder payload(@Nullable Output<String> payload) {
            $.payload = payload;
            return this;
        }

        public Builder payload(String payload) {
            return payload(Output.of(payload));
        }

        /**
         * @param proxy Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
         * 
         * @return builder
         * 
         */
        public Builder proxy(@Nullable Output<String> proxy) {
            $.proxy = proxy;
            return this;
        }

        /**
         * @param proxy Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
         * 
         * @return builder
         * 
         */
        public Builder proxy(String proxy) {
            return proxy(Output.of(proxy));
        }

        /**
         * @param secrets Defines a set of sensitive values (such as, tokens and passwords) that can be injected in the headers and/or payload.Secrets’ values are encrypted. In the header/payload, the value can be invoked using the `{{.secrets.token}}` format, where token is the name provided for the secret value. Comprise key/value pair. **Note:** if multiple handlers are used, same secret name and different secret value for the same url won&#39;t work. Example:
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
         * import java.util.List;
         * import java.util.ArrayList;
         * import java.util.Map;
         * import java.io.File;
         * import java.nio.file.Files;
         * import java.nio.file.Paths;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder secrets(@Nullable Output<Map<String,String>> secrets) {
            $.secrets = secrets;
            return this;
        }

        /**
         * @param secrets Defines a set of sensitive values (such as, tokens and passwords) that can be injected in the headers and/or payload.Secrets’ values are encrypted. In the header/payload, the value can be invoked using the `{{.secrets.token}}` format, where token is the name provided for the secret value. Comprise key/value pair. **Note:** if multiple handlers are used, same secret name and different secret value for the same url won&#39;t work. Example:
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
         * import java.util.List;
         * import java.util.ArrayList;
         * import java.util.Map;
         * import java.io.File;
         * import java.nio.file.Files;
         * import java.nio.file.Paths;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder secrets(Map<String,String> secrets) {
            return secrets(Output.of(secrets));
        }

        /**
         * @param url Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public ArtifactCustomWebhookHandlerArgs build() {
            if ($.url == null) {
                throw new MissingRequiredPropertyException("ArtifactCustomWebhookHandlerArgs", "url");
            }
            return $;
        }
    }

}

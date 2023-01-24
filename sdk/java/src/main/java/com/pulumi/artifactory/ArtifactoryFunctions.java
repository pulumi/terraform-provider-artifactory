// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.GetFileArgs;
import com.pulumi.artifactory.inputs.GetFilePlainArgs;
import com.pulumi.artifactory.inputs.GetFileinfoArgs;
import com.pulumi.artifactory.inputs.GetFileinfoPlainArgs;
import com.pulumi.artifactory.inputs.GetGroupArgs;
import com.pulumi.artifactory.inputs.GetGroupPlainArgs;
import com.pulumi.artifactory.inputs.GetPermissionTargetArgs;
import com.pulumi.artifactory.inputs.GetPermissionTargetPlainArgs;
import com.pulumi.artifactory.inputs.GetUserArgs;
import com.pulumi.artifactory.inputs.GetUserPlainArgs;
import com.pulumi.artifactory.outputs.GetFileResult;
import com.pulumi.artifactory.outputs.GetFileinfoResult;
import com.pulumi.artifactory.outputs.GetGroupResult;
import com.pulumi.artifactory.outputs.GetPermissionTargetResult;
import com.pulumi.artifactory.outputs.GetUserResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class ArtifactoryFunctions {
    /**
     * ## # Artifactory File Data Source
     * 
     * Provides an Artifactory file datasource. This can be used to download a file from a given Artifactory repository.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetFileArgs;
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
     *         final var my-file = ArtifactoryFunctions.getFile(GetFileArgs.builder()
     *             .outputPath(&#34;tmp/artifact.zip&#34;)
     *             .path(&#34;/path/to/the/artifact.zip&#34;)
     *             .repository(&#34;repo-key&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetFileResult> getFile(GetFileArgs args) {
        return getFile(args, InvokeOptions.Empty);
    }
    /**
     * ## # Artifactory File Data Source
     * 
     * Provides an Artifactory file datasource. This can be used to download a file from a given Artifactory repository.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetFileArgs;
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
     *         final var my-file = ArtifactoryFunctions.getFile(GetFileArgs.builder()
     *             .outputPath(&#34;tmp/artifact.zip&#34;)
     *             .path(&#34;/path/to/the/artifact.zip&#34;)
     *             .repository(&#34;repo-key&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetFileResult> getFilePlain(GetFilePlainArgs args) {
        return getFilePlain(args, InvokeOptions.Empty);
    }
    /**
     * ## # Artifactory File Data Source
     * 
     * Provides an Artifactory file datasource. This can be used to download a file from a given Artifactory repository.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetFileArgs;
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
     *         final var my-file = ArtifactoryFunctions.getFile(GetFileArgs.builder()
     *             .outputPath(&#34;tmp/artifact.zip&#34;)
     *             .path(&#34;/path/to/the/artifact.zip&#34;)
     *             .repository(&#34;repo-key&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetFileResult> getFile(GetFileArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("artifactory:index/getFile:getFile", TypeShape.of(GetFileResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## # Artifactory File Data Source
     * 
     * Provides an Artifactory file datasource. This can be used to download a file from a given Artifactory repository.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetFileArgs;
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
     *         final var my-file = ArtifactoryFunctions.getFile(GetFileArgs.builder()
     *             .outputPath(&#34;tmp/artifact.zip&#34;)
     *             .path(&#34;/path/to/the/artifact.zip&#34;)
     *             .repository(&#34;repo-key&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetFileResult> getFilePlain(GetFilePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("artifactory:index/getFile:getFile", TypeShape.of(GetFileResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## # Artifactory File Info Data Source
     * 
     * Provides an Artifactory fileinfo datasource. This can be used to read metadata of files stored in Artifactory repositories.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetFileinfoArgs;
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
     *         final var my-file = ArtifactoryFunctions.getFileinfo(GetFileinfoArgs.builder()
     *             .path(&#34;/path/to/the/artifact.zip&#34;)
     *             .repository(&#34;repo-key&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetFileinfoResult> getFileinfo(GetFileinfoArgs args) {
        return getFileinfo(args, InvokeOptions.Empty);
    }
    /**
     * ## # Artifactory File Info Data Source
     * 
     * Provides an Artifactory fileinfo datasource. This can be used to read metadata of files stored in Artifactory repositories.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetFileinfoArgs;
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
     *         final var my-file = ArtifactoryFunctions.getFileinfo(GetFileinfoArgs.builder()
     *             .path(&#34;/path/to/the/artifact.zip&#34;)
     *             .repository(&#34;repo-key&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetFileinfoResult> getFileinfoPlain(GetFileinfoPlainArgs args) {
        return getFileinfoPlain(args, InvokeOptions.Empty);
    }
    /**
     * ## # Artifactory File Info Data Source
     * 
     * Provides an Artifactory fileinfo datasource. This can be used to read metadata of files stored in Artifactory repositories.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetFileinfoArgs;
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
     *         final var my-file = ArtifactoryFunctions.getFileinfo(GetFileinfoArgs.builder()
     *             .path(&#34;/path/to/the/artifact.zip&#34;)
     *             .repository(&#34;repo-key&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetFileinfoResult> getFileinfo(GetFileinfoArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("artifactory:index/getFileinfo:getFileinfo", TypeShape.of(GetFileinfoResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## # Artifactory File Info Data Source
     * 
     * Provides an Artifactory fileinfo datasource. This can be used to read metadata of files stored in Artifactory repositories.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetFileinfoArgs;
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
     *         final var my-file = ArtifactoryFunctions.getFileinfo(GetFileinfoArgs.builder()
     *             .path(&#34;/path/to/the/artifact.zip&#34;)
     *             .repository(&#34;repo-key&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetFileinfoResult> getFileinfoPlain(GetFileinfoPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("artifactory:index/getFileinfo:getFileinfo", TypeShape.of(GetFileinfoResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## # Artifactory Group Data Source
     * 
     * Provides an Artifactory group datasource. This can be used to read the configuration of groups in artifactory.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetGroupArgs;
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
     *         final var myGroup = ArtifactoryFunctions.getGroup(GetGroupArgs.builder()
     *             .includeUsers(true)
     *             .name(&#34;my_group&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetGroupResult> getGroup(GetGroupArgs args) {
        return getGroup(args, InvokeOptions.Empty);
    }
    /**
     * ## # Artifactory Group Data Source
     * 
     * Provides an Artifactory group datasource. This can be used to read the configuration of groups in artifactory.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetGroupArgs;
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
     *         final var myGroup = ArtifactoryFunctions.getGroup(GetGroupArgs.builder()
     *             .includeUsers(true)
     *             .name(&#34;my_group&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetGroupResult> getGroupPlain(GetGroupPlainArgs args) {
        return getGroupPlain(args, InvokeOptions.Empty);
    }
    /**
     * ## # Artifactory Group Data Source
     * 
     * Provides an Artifactory group datasource. This can be used to read the configuration of groups in artifactory.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetGroupArgs;
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
     *         final var myGroup = ArtifactoryFunctions.getGroup(GetGroupArgs.builder()
     *             .includeUsers(true)
     *             .name(&#34;my_group&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetGroupResult> getGroup(GetGroupArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("artifactory:index/getGroup:getGroup", TypeShape.of(GetGroupResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## # Artifactory Group Data Source
     * 
     * Provides an Artifactory group datasource. This can be used to read the configuration of groups in artifactory.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetGroupArgs;
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
     *         final var myGroup = ArtifactoryFunctions.getGroup(GetGroupArgs.builder()
     *             .includeUsers(true)
     *             .name(&#34;my_group&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetGroupResult> getGroupPlain(GetGroupPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("artifactory:index/getGroup:getGroup", TypeShape.of(GetGroupResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## # Artifactory Permission Target Data Source
     * 
     * Provides an Artifactory permission target data source. This can be used to read the configuration of permission targets in artifactory.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetPermissionTargetArgs;
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
     *         final var target1 = ArtifactoryFunctions.getPermissionTarget(GetPermissionTargetArgs.builder()
     *             .name(&#34;my_permission&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetPermissionTargetResult> getPermissionTarget(GetPermissionTargetArgs args) {
        return getPermissionTarget(args, InvokeOptions.Empty);
    }
    /**
     * ## # Artifactory Permission Target Data Source
     * 
     * Provides an Artifactory permission target data source. This can be used to read the configuration of permission targets in artifactory.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetPermissionTargetArgs;
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
     *         final var target1 = ArtifactoryFunctions.getPermissionTarget(GetPermissionTargetArgs.builder()
     *             .name(&#34;my_permission&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetPermissionTargetResult> getPermissionTargetPlain(GetPermissionTargetPlainArgs args) {
        return getPermissionTargetPlain(args, InvokeOptions.Empty);
    }
    /**
     * ## # Artifactory Permission Target Data Source
     * 
     * Provides an Artifactory permission target data source. This can be used to read the configuration of permission targets in artifactory.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetPermissionTargetArgs;
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
     *         final var target1 = ArtifactoryFunctions.getPermissionTarget(GetPermissionTargetArgs.builder()
     *             .name(&#34;my_permission&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetPermissionTargetResult> getPermissionTarget(GetPermissionTargetArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("artifactory:index/getPermissionTarget:getPermissionTarget", TypeShape.of(GetPermissionTargetResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## # Artifactory Permission Target Data Source
     * 
     * Provides an Artifactory permission target data source. This can be used to read the configuration of permission targets in artifactory.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetPermissionTargetArgs;
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
     *         final var target1 = ArtifactoryFunctions.getPermissionTarget(GetPermissionTargetArgs.builder()
     *             .name(&#34;my_permission&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetPermissionTargetResult> getPermissionTargetPlain(GetPermissionTargetPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("artifactory:index/getPermissionTarget:getPermissionTarget", TypeShape.of(GetPermissionTargetResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## # Artifactory User Data Source
     * 
     * Provides an Artifactory user data source. This can be used to read the configuration of users in artifactory.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetUserArgs;
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
     *         final var user1 = ArtifactoryFunctions.getUser(GetUserArgs.builder()
     *             .name(&#34;user1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetUserResult> getUser(GetUserArgs args) {
        return getUser(args, InvokeOptions.Empty);
    }
    /**
     * ## # Artifactory User Data Source
     * 
     * Provides an Artifactory user data source. This can be used to read the configuration of users in artifactory.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetUserArgs;
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
     *         final var user1 = ArtifactoryFunctions.getUser(GetUserArgs.builder()
     *             .name(&#34;user1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetUserResult> getUserPlain(GetUserPlainArgs args) {
        return getUserPlain(args, InvokeOptions.Empty);
    }
    /**
     * ## # Artifactory User Data Source
     * 
     * Provides an Artifactory user data source. This can be used to read the configuration of users in artifactory.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetUserArgs;
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
     *         final var user1 = ArtifactoryFunctions.getUser(GetUserArgs.builder()
     *             .name(&#34;user1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetUserResult> getUser(GetUserArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("artifactory:index/getUser:getUser", TypeShape.of(GetUserResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## # Artifactory User Data Source
     * 
     * Provides an Artifactory user data source. This can be used to read the configuration of users in artifactory.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.artifactory.ArtifactoryFunctions;
     * import com.pulumi.artifactory.inputs.GetUserArgs;
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
     *         final var user1 = ArtifactoryFunctions.getUser(GetUserArgs.builder()
     *             .name(&#34;user1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetUserResult> getUserPlain(GetUserPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("artifactory:index/getUser:getUser", TypeShape.of(GetUserResult.class), args, Utilities.withVersion(options));
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.BackupArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.BackupState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can be used to manage the automatic and periodic backups of the entire Artifactory instance.
 * 
 * When an `artifactory.Backup` resource is configured and enabled to true, backup of the entire Artifactory system will be done automatically and periodically.
 * The backup process creates a time-stamped directory in the target backup directory.
 * 
 * ~&gt;The `artifactory.Backup` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.Backup;
 * import com.pulumi.artifactory.BackupArgs;
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
 *         var backupConfigName = new Backup(&#34;backupConfigName&#34;, BackupArgs.builder()        
 *             .createArchive(false)
 *             .cronExp(&#34;0 0 12 * * ? *&#34;)
 *             .enabled(true)
 *             .excludeNewRepositories(true)
 *             .excludedRepositories()
 *             .exportMissionControl(true)
 *             .key(&#34;backup_config_name&#34;)
 *             .retentionPeriodHours(1000)
 *             .sendMailOnError(true)
 *             .verifyDiskSpace(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * Note: `Key` argument has to match to the resource name.
 * Reference Link: [JFrog Artifactory Backup](https://www.jfrog.com/confluence/display/JFROG/Backups)
 * 
 * ## Import
 * 
 * Backup config can be imported using the key, e.g.
 * 
 * ```sh
 *  $ pulumi import artifactory:index/backup:Backup backup_name backup_name
 * ```
 * 
 */
@ResourceType(type="artifactory:index/backup:Backup")
public class Backup extends com.pulumi.resources.CustomResource {
    /**
     * If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
     * 
     */
    @Export(name="createArchive", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> createArchive;

    /**
     * @return If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> createArchive() {
        return Codegen.optional(this.createArchive);
    }
    /**
     * A valid CRON expression that you can use to control backup frequency. Eg: &#34;0 0 12 * * ? *&#34;, &#34;0 0 2 ? * MON-SAT *&#34;. Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
     * 
     */
    @Export(name="cronExp", type=String.class, parameters={})
    private Output<String> cronExp;

    /**
     * @return A valid CRON expression that you can use to control backup frequency. Eg: &#34;0 0 12 * * ? *&#34;, &#34;0 0 2 ? * MON-SAT *&#34;. Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
     * 
     */
    public Output<String> cronExp() {
        return this.cronExp;
    }
    /**
     * Flag to enable or disable the backup config. Default value is `true`.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Flag to enable or disable the backup config. Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * When set, new repositories will not be automatically added to the backup. Default value is `false`.
     * 
     */
    @Export(name="excludeNewRepositories", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> excludeNewRepositories;

    /**
     * @return When set, new repositories will not be automatically added to the backup. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> excludeNewRepositories() {
        return Codegen.optional(this.excludeNewRepositories);
    }
    /**
     * A list of excluded repositories from the backup. Default is empty list.
     * 
     */
    @Export(name="excludedRepositories", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> excludedRepositories;

    /**
     * @return A list of excluded repositories from the backup. Default is empty list.
     * 
     */
    public Output<Optional<List<String>>> excludedRepositories() {
        return Codegen.optional(this.excludedRepositories);
    }
    /**
     * When set to true, mission control will not be automatically added to the backup. Default value is &#39;false&#39;.
     * 
     */
    @Export(name="exportMissionControl", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> exportMissionControl;

    /**
     * @return When set to true, mission control will not be automatically added to the backup. Default value is &#39;false&#39;.
     * 
     */
    public Output<Optional<Boolean>> exportMissionControl() {
        return Codegen.optional(this.exportMissionControl);
    }
    /**
     * The unique ID of the artifactory backup config.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return The unique ID of the artifactory backup config.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
     * 
     */
    @Export(name="retentionPeriodHours", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> retentionPeriodHours;

    /**
     * @return The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
     * 
     */
    public Output<Optional<Integer>> retentionPeriodHours() {
        return Codegen.optional(this.retentionPeriodHours);
    }
    /**
     * If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
     * 
     */
    @Export(name="sendMailOnError", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> sendMailOnError;

    /**
     * @return If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> sendMailOnError() {
        return Codegen.optional(this.sendMailOnError);
    }
    /**
     * If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
     * 
     */
    @Export(name="verifyDiskSpace", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> verifyDiskSpace;

    /**
     * @return If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
     * 
     */
    public Output<Optional<Boolean>> verifyDiskSpace() {
        return Codegen.optional(this.verifyDiskSpace);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Backup(String name) {
        this(name, BackupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Backup(String name, BackupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Backup(String name, BackupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/backup:Backup", name, args == null ? BackupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Backup(String name, Output<String> id, @Nullable BackupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/backup:Backup", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Backup get(String name, Output<String> id, @Nullable BackupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Backup(name, id, state, options);
    }
}

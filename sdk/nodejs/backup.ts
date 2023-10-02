// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can be used to manage the automatic and periodic backups of the entire Artifactory instance.
 *
 * When an `artifactory.Backup` resource is configured and enabled to true, backup of the entire Artifactory system will be done automatically and periodically.
 * The backup process creates a time-stamped directory in the target backup directory.
 *
 * ~>The `artifactory.Backup` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * // Configure Artifactory Backup system config
 * const backupConfigName = new artifactory.Backup("backupConfigName", {
 *     createArchive: false,
 *     cronExp: "0 0 12 * * ? *",
 *     enabled: true,
 *     excludeNewRepositories: true,
 *     excludedRepositories: [],
 *     exportMissionControl: true,
 *     key: "backup_config_name",
 *     retentionPeriodHours: 1000,
 *     sendMailOnError: true,
 *     verifyDiskSpace: true,
 * });
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
 */
export class Backup extends pulumi.CustomResource {
    /**
     * Get an existing Backup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupState, opts?: pulumi.CustomResourceOptions): Backup {
        return new Backup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/backup:Backup';

    /**
     * Returns true if the given object is an instance of Backup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Backup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Backup.__pulumiType;
    }

    /**
     * If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
     */
    public readonly createArchive!: pulumi.Output<boolean>;
    /**
     * A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
     */
    public readonly cronExp!: pulumi.Output<string>;
    /**
     * Flag to enable or disable the backup config. Default value is `true`.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * When set, new repositories will not be automatically added to the backup. Default value is `false`.
     */
    public readonly excludeNewRepositories!: pulumi.Output<boolean>;
    /**
     * A list of excluded repositories from the backup. Default is empty list.
     */
    public readonly excludedRepositories!: pulumi.Output<string[] | undefined>;
    /**
     * When set to true, mission control will not be automatically added to the backup. Default value is `false`.
     */
    public readonly exportMissionControl!: pulumi.Output<boolean>;
    /**
     * The unique ID of the artifactory backup config.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
     */
    public readonly retentionPeriodHours!: pulumi.Output<number>;
    /**
     * If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
     */
    public readonly sendMailOnError!: pulumi.Output<boolean>;
    /**
     * If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
     */
    public readonly verifyDiskSpace!: pulumi.Output<boolean>;

    /**
     * Create a Backup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupArgs | BackupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackupState | undefined;
            resourceInputs["createArchive"] = state ? state.createArchive : undefined;
            resourceInputs["cronExp"] = state ? state.cronExp : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["excludeNewRepositories"] = state ? state.excludeNewRepositories : undefined;
            resourceInputs["excludedRepositories"] = state ? state.excludedRepositories : undefined;
            resourceInputs["exportMissionControl"] = state ? state.exportMissionControl : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["retentionPeriodHours"] = state ? state.retentionPeriodHours : undefined;
            resourceInputs["sendMailOnError"] = state ? state.sendMailOnError : undefined;
            resourceInputs["verifyDiskSpace"] = state ? state.verifyDiskSpace : undefined;
        } else {
            const args = argsOrState as BackupArgs | undefined;
            if ((!args || args.cronExp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cronExp'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["createArchive"] = args ? args.createArchive : undefined;
            resourceInputs["cronExp"] = args ? args.cronExp : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["excludeNewRepositories"] = args ? args.excludeNewRepositories : undefined;
            resourceInputs["excludedRepositories"] = args ? args.excludedRepositories : undefined;
            resourceInputs["exportMissionControl"] = args ? args.exportMissionControl : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["retentionPeriodHours"] = args ? args.retentionPeriodHours : undefined;
            resourceInputs["sendMailOnError"] = args ? args.sendMailOnError : undefined;
            resourceInputs["verifyDiskSpace"] = args ? args.verifyDiskSpace : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Backup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Backup resources.
 */
export interface BackupState {
    /**
     * If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
     */
    createArchive?: pulumi.Input<boolean>;
    /**
     * A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
     */
    cronExp?: pulumi.Input<string>;
    /**
     * Flag to enable or disable the backup config. Default value is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * When set, new repositories will not be automatically added to the backup. Default value is `false`.
     */
    excludeNewRepositories?: pulumi.Input<boolean>;
    /**
     * A list of excluded repositories from the backup. Default is empty list.
     */
    excludedRepositories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When set to true, mission control will not be automatically added to the backup. Default value is `false`.
     */
    exportMissionControl?: pulumi.Input<boolean>;
    /**
     * The unique ID of the artifactory backup config.
     */
    key?: pulumi.Input<string>;
    /**
     * The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
     */
    retentionPeriodHours?: pulumi.Input<number>;
    /**
     * If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
     */
    sendMailOnError?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
     */
    verifyDiskSpace?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Backup resource.
 */
export interface BackupArgs {
    /**
     * If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
     */
    createArchive?: pulumi.Input<boolean>;
    /**
     * A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
     */
    cronExp: pulumi.Input<string>;
    /**
     * Flag to enable or disable the backup config. Default value is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * When set, new repositories will not be automatically added to the backup. Default value is `false`.
     */
    excludeNewRepositories?: pulumi.Input<boolean>;
    /**
     * A list of excluded repositories from the backup. Default is empty list.
     */
    excludedRepositories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When set to true, mission control will not be automatically added to the backup. Default value is `false`.
     */
    exportMissionControl?: pulumi.Input<boolean>;
    /**
     * The unique ID of the artifactory backup config.
     */
    key: pulumi.Input<string>;
    /**
     * The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
     */
    retentionPeriodHours?: pulumi.Input<number>;
    /**
     * If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
     */
    sendMailOnError?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
     */
    verifyDiskSpace?: pulumi.Input<boolean>;
}

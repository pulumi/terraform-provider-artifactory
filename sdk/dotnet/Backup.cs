// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// This resource can be used to manage the automatic and periodic backups of the entire Artifactory instance.
    /// 
    /// When an `artifactory.Backup` resource is configured and enabled to true, backup of the entire Artifactory system will be done automatically and periodically.
    /// The backup process creates a time-stamped directory in the target backup directory.
    /// 
    /// ~&gt;The `artifactory.Backup` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Configure Artifactory Backup system config
    ///     var backupConfigName = new Artifactory.Index.Backup.Backup("backupConfigName", new()
    ///     {
    ///         CreateArchive = false,
    ///         CronExp = "0 0 12 * * ? *",
    ///         Enabled = true,
    ///         ExcludeNewRepositories = true,
    ///         ExcludedRepositories = new[] {},
    ///         ExportMissionControl = true,
    ///         Key = "backup_config_name",
    ///         RetentionPeriodHours = 1000,
    ///         SendMailOnError = true,
    ///         VerifyDiskSpace = true,
    ///     });
    /// 
    /// });
    /// ```
    /// Note: `Key` argument has to match to the resource name.
    /// Reference Link: [JFrog Artifactory Backup](https://www.jfrog.com/confluence/display/JFROG/Backups)
    /// 
    /// ## Import
    /// 
    /// Backup config can be imported using the key, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/backup:Backup backup_name backup_name
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/backup:Backup")]
    public partial class Backup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        /// </summary>
        [Output("createArchive")]
        public Output<bool?> CreateArchive { get; private set; } = null!;

        /// <summary>
        /// A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
        /// </summary>
        [Output("cronExp")]
        public Output<string> CronExp { get; private set; } = null!;

        /// <summary>
        /// Flag to enable or disable the backup config. Default value is `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// When set, new repositories will not be automatically added to the backup. Default value is `false`.
        /// </summary>
        [Output("excludeNewRepositories")]
        public Output<bool?> ExcludeNewRepositories { get; private set; } = null!;

        /// <summary>
        /// A list of excluded repositories from the backup. Default is empty list.
        /// </summary>
        [Output("excludedRepositories")]
        public Output<ImmutableArray<string>> ExcludedRepositories { get; private set; } = null!;

        /// <summary>
        /// When set to true, mission control will not be automatically added to the backup. Default value is `false`.
        /// </summary>
        [Output("exportMissionControl")]
        public Output<bool?> ExportMissionControl { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the artifactory backup config.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        /// </summary>
        [Output("retentionPeriodHours")]
        public Output<int?> RetentionPeriodHours { get; private set; } = null!;

        /// <summary>
        /// If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        /// </summary>
        [Output("sendMailOnError")]
        public Output<bool?> SendMailOnError { get; private set; } = null!;

        /// <summary>
        /// If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
        /// </summary>
        [Output("verifyDiskSpace")]
        public Output<bool?> VerifyDiskSpace { get; private set; } = null!;


        /// <summary>
        /// Create a Backup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backup(string name, BackupArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/backup:Backup", name, args ?? new BackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Backup(string name, Input<string> id, BackupState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/backup:Backup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Backup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Backup Get(string name, Input<string> id, BackupState? state = null, CustomResourceOptions? options = null)
        {
            return new Backup(name, id, state, options);
        }
    }

    public sealed class BackupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        /// </summary>
        [Input("createArchive")]
        public Input<bool>? CreateArchive { get; set; }

        /// <summary>
        /// A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
        /// </summary>
        [Input("cronExp", required: true)]
        public Input<string> CronExp { get; set; } = null!;

        /// <summary>
        /// Flag to enable or disable the backup config. Default value is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When set, new repositories will not be automatically added to the backup. Default value is `false`.
        /// </summary>
        [Input("excludeNewRepositories")]
        public Input<bool>? ExcludeNewRepositories { get; set; }

        [Input("excludedRepositories")]
        private InputList<string>? _excludedRepositories;

        /// <summary>
        /// A list of excluded repositories from the backup. Default is empty list.
        /// </summary>
        public InputList<string> ExcludedRepositories
        {
            get => _excludedRepositories ?? (_excludedRepositories = new InputList<string>());
            set => _excludedRepositories = value;
        }

        /// <summary>
        /// When set to true, mission control will not be automatically added to the backup. Default value is `false`.
        /// </summary>
        [Input("exportMissionControl")]
        public Input<bool>? ExportMissionControl { get; set; }

        /// <summary>
        /// The unique ID of the artifactory backup config.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        /// </summary>
        [Input("retentionPeriodHours")]
        public Input<int>? RetentionPeriodHours { get; set; }

        /// <summary>
        /// If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        /// </summary>
        [Input("sendMailOnError")]
        public Input<bool>? SendMailOnError { get; set; }

        /// <summary>
        /// If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
        /// </summary>
        [Input("verifyDiskSpace")]
        public Input<bool>? VerifyDiskSpace { get; set; }

        public BackupArgs()
        {
        }
        public static new BackupArgs Empty => new BackupArgs();
    }

    public sealed class BackupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        /// </summary>
        [Input("createArchive")]
        public Input<bool>? CreateArchive { get; set; }

        /// <summary>
        /// A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
        /// </summary>
        [Input("cronExp")]
        public Input<string>? CronExp { get; set; }

        /// <summary>
        /// Flag to enable or disable the backup config. Default value is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When set, new repositories will not be automatically added to the backup. Default value is `false`.
        /// </summary>
        [Input("excludeNewRepositories")]
        public Input<bool>? ExcludeNewRepositories { get; set; }

        [Input("excludedRepositories")]
        private InputList<string>? _excludedRepositories;

        /// <summary>
        /// A list of excluded repositories from the backup. Default is empty list.
        /// </summary>
        public InputList<string> ExcludedRepositories
        {
            get => _excludedRepositories ?? (_excludedRepositories = new InputList<string>());
            set => _excludedRepositories = value;
        }

        /// <summary>
        /// When set to true, mission control will not be automatically added to the backup. Default value is `false`.
        /// </summary>
        [Input("exportMissionControl")]
        public Input<bool>? ExportMissionControl { get; set; }

        /// <summary>
        /// The unique ID of the artifactory backup config.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        /// </summary>
        [Input("retentionPeriodHours")]
        public Input<int>? RetentionPeriodHours { get; set; }

        /// <summary>
        /// If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        /// </summary>
        [Input("sendMailOnError")]
        public Input<bool>? SendMailOnError { get; set; }

        /// <summary>
        /// If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
        /// </summary>
        [Input("verifyDiskSpace")]
        public Input<bool>? VerifyDiskSpace { get; set; }

        public BackupState()
        {
        }
        public static new BackupState Empty => new BackupState();
    }
}

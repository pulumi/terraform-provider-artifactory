// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This resource can be used to manage the automatic and periodic backups of the entire Artifactory instance.
//
// When an `Backup` resource is configured and enabled to true, backup of the entire Artifactory system will be done automatically and periodically.
// The backup process creates a time-stamped directory in the target backup directory.
//
// ~>The `Backup` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewBackup(ctx, "backupConfigName", &artifactory.BackupArgs{
//				CreateArchive:          pulumi.Bool(false),
//				CronExp:                pulumi.String("0 0 12 * * ? *"),
//				Enabled:                pulumi.Bool(true),
//				ExcludeNewRepositories: pulumi.Bool(true),
//				ExcludedRepositories:   pulumi.StringArray{},
//				ExportMissionControl:   pulumi.Bool(true),
//				Key:                    pulumi.String("backup_config_name"),
//				RetentionPeriodHours:   pulumi.Int(1000),
//				SendMailOnError:        pulumi.Bool(true),
//				VerifyDiskSpace:        pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// Note: `Key` argument has to match to the resource name.
// Reference Link: [JFrog Artifactory Backup](https://www.jfrog.com/confluence/display/JFROG/Backups)
//
// ## Import
//
// Backup config can be imported using the key, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/backup:Backup backup_name backup_name
//
// ```
type Backup struct {
	pulumi.CustomResourceState

	// If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
	CreateArchive pulumi.BoolPtrOutput `pulumi:"createArchive"`
	// A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
	CronExp pulumi.StringOutput `pulumi:"cronExp"`
	// Flag to enable or disable the backup config. Default value is `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// When set, new repositories will not be automatically added to the backup. Default value is `false`.
	ExcludeNewRepositories pulumi.BoolPtrOutput `pulumi:"excludeNewRepositories"`
	// A list of excluded repositories from the backup. Default is empty list.
	ExcludedRepositories pulumi.StringArrayOutput `pulumi:"excludedRepositories"`
	// When set to true, mission control will not be automatically added to the backup. Default value is `false`.
	ExportMissionControl pulumi.BoolPtrOutput `pulumi:"exportMissionControl"`
	// The unique ID of the artifactory backup config.
	Key pulumi.StringOutput `pulumi:"key"`
	// The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
	RetentionPeriodHours pulumi.IntPtrOutput `pulumi:"retentionPeriodHours"`
	// If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
	SendMailOnError pulumi.BoolPtrOutput `pulumi:"sendMailOnError"`
	// If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
	VerifyDiskSpace pulumi.BoolPtrOutput `pulumi:"verifyDiskSpace"`
}

// NewBackup registers a new resource with the given unique name, arguments, and options.
func NewBackup(ctx *pulumi.Context,
	name string, args *BackupArgs, opts ...pulumi.ResourceOption) (*Backup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CronExp == nil {
		return nil, errors.New("invalid value for required argument 'CronExp'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Backup
	err := ctx.RegisterResource("artifactory:index/backup:Backup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackup gets an existing Backup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupState, opts ...pulumi.ResourceOption) (*Backup, error) {
	var resource Backup
	err := ctx.ReadResource("artifactory:index/backup:Backup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backup resources.
type backupState struct {
	// If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
	CreateArchive *bool `pulumi:"createArchive"`
	// A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
	CronExp *string `pulumi:"cronExp"`
	// Flag to enable or disable the backup config. Default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// When set, new repositories will not be automatically added to the backup. Default value is `false`.
	ExcludeNewRepositories *bool `pulumi:"excludeNewRepositories"`
	// A list of excluded repositories from the backup. Default is empty list.
	ExcludedRepositories []string `pulumi:"excludedRepositories"`
	// When set to true, mission control will not be automatically added to the backup. Default value is `false`.
	ExportMissionControl *bool `pulumi:"exportMissionControl"`
	// The unique ID of the artifactory backup config.
	Key *string `pulumi:"key"`
	// The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
	RetentionPeriodHours *int `pulumi:"retentionPeriodHours"`
	// If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
	SendMailOnError *bool `pulumi:"sendMailOnError"`
	// If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
	VerifyDiskSpace *bool `pulumi:"verifyDiskSpace"`
}

type BackupState struct {
	// If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
	CreateArchive pulumi.BoolPtrInput
	// A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
	CronExp pulumi.StringPtrInput
	// Flag to enable or disable the backup config. Default value is `true`.
	Enabled pulumi.BoolPtrInput
	// When set, new repositories will not be automatically added to the backup. Default value is `false`.
	ExcludeNewRepositories pulumi.BoolPtrInput
	// A list of excluded repositories from the backup. Default is empty list.
	ExcludedRepositories pulumi.StringArrayInput
	// When set to true, mission control will not be automatically added to the backup. Default value is `false`.
	ExportMissionControl pulumi.BoolPtrInput
	// The unique ID of the artifactory backup config.
	Key pulumi.StringPtrInput
	// The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
	RetentionPeriodHours pulumi.IntPtrInput
	// If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
	SendMailOnError pulumi.BoolPtrInput
	// If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
	VerifyDiskSpace pulumi.BoolPtrInput
}

func (BackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupState)(nil)).Elem()
}

type backupArgs struct {
	// If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
	CreateArchive *bool `pulumi:"createArchive"`
	// A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
	CronExp string `pulumi:"cronExp"`
	// Flag to enable or disable the backup config. Default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// When set, new repositories will not be automatically added to the backup. Default value is `false`.
	ExcludeNewRepositories *bool `pulumi:"excludeNewRepositories"`
	// A list of excluded repositories from the backup. Default is empty list.
	ExcludedRepositories []string `pulumi:"excludedRepositories"`
	// When set to true, mission control will not be automatically added to the backup. Default value is `false`.
	ExportMissionControl *bool `pulumi:"exportMissionControl"`
	// The unique ID of the artifactory backup config.
	Key string `pulumi:"key"`
	// The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
	RetentionPeriodHours *int `pulumi:"retentionPeriodHours"`
	// If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
	SendMailOnError *bool `pulumi:"sendMailOnError"`
	// If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
	VerifyDiskSpace *bool `pulumi:"verifyDiskSpace"`
}

// The set of arguments for constructing a Backup resource.
type BackupArgs struct {
	// If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
	CreateArchive pulumi.BoolPtrInput
	// A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
	CronExp pulumi.StringInput
	// Flag to enable or disable the backup config. Default value is `true`.
	Enabled pulumi.BoolPtrInput
	// When set, new repositories will not be automatically added to the backup. Default value is `false`.
	ExcludeNewRepositories pulumi.BoolPtrInput
	// A list of excluded repositories from the backup. Default is empty list.
	ExcludedRepositories pulumi.StringArrayInput
	// When set to true, mission control will not be automatically added to the backup. Default value is `false`.
	ExportMissionControl pulumi.BoolPtrInput
	// The unique ID of the artifactory backup config.
	Key pulumi.StringInput
	// The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
	RetentionPeriodHours pulumi.IntPtrInput
	// If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
	SendMailOnError pulumi.BoolPtrInput
	// If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
	VerifyDiskSpace pulumi.BoolPtrInput
}

func (BackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupArgs)(nil)).Elem()
}

type BackupInput interface {
	pulumi.Input

	ToBackupOutput() BackupOutput
	ToBackupOutputWithContext(ctx context.Context) BackupOutput
}

func (*Backup) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (i *Backup) ToBackupOutput() BackupOutput {
	return i.ToBackupOutputWithContext(context.Background())
}

func (i *Backup) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput)
}

func (i *Backup) ToOutput(ctx context.Context) pulumix.Output[*Backup] {
	return pulumix.Output[*Backup]{
		OutputState: i.ToBackupOutputWithContext(ctx).OutputState,
	}
}

// BackupArrayInput is an input type that accepts BackupArray and BackupArrayOutput values.
// You can construct a concrete instance of `BackupArrayInput` via:
//
//	BackupArray{ BackupArgs{...} }
type BackupArrayInput interface {
	pulumi.Input

	ToBackupArrayOutput() BackupArrayOutput
	ToBackupArrayOutputWithContext(context.Context) BackupArrayOutput
}

type BackupArray []BackupInput

func (BackupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Backup)(nil)).Elem()
}

func (i BackupArray) ToBackupArrayOutput() BackupArrayOutput {
	return i.ToBackupArrayOutputWithContext(context.Background())
}

func (i BackupArray) ToBackupArrayOutputWithContext(ctx context.Context) BackupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupArrayOutput)
}

func (i BackupArray) ToOutput(ctx context.Context) pulumix.Output[[]*Backup] {
	return pulumix.Output[[]*Backup]{
		OutputState: i.ToBackupArrayOutputWithContext(ctx).OutputState,
	}
}

// BackupMapInput is an input type that accepts BackupMap and BackupMapOutput values.
// You can construct a concrete instance of `BackupMapInput` via:
//
//	BackupMap{ "key": BackupArgs{...} }
type BackupMapInput interface {
	pulumi.Input

	ToBackupMapOutput() BackupMapOutput
	ToBackupMapOutputWithContext(context.Context) BackupMapOutput
}

type BackupMap map[string]BackupInput

func (BackupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Backup)(nil)).Elem()
}

func (i BackupMap) ToBackupMapOutput() BackupMapOutput {
	return i.ToBackupMapOutputWithContext(context.Background())
}

func (i BackupMap) ToBackupMapOutputWithContext(ctx context.Context) BackupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupMapOutput)
}

func (i BackupMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Backup] {
	return pulumix.Output[map[string]*Backup]{
		OutputState: i.ToBackupMapOutputWithContext(ctx).OutputState,
	}
}

type BackupOutput struct{ *pulumi.OutputState }

func (BackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (o BackupOutput) ToBackupOutput() BackupOutput {
	return o
}

func (o BackupOutput) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return o
}

func (o BackupOutput) ToOutput(ctx context.Context) pulumix.Output[*Backup] {
	return pulumix.Output[*Backup]{
		OutputState: o.OutputState,
	}
}

// If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
func (o BackupOutput) CreateArchive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Backup) pulumi.BoolPtrOutput { return v.CreateArchive }).(pulumi.BoolPtrOutput)
}

// A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
func (o BackupOutput) CronExp() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.CronExp }).(pulumi.StringOutput)
}

// Flag to enable or disable the backup config. Default value is `true`.
func (o BackupOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Backup) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// When set, new repositories will not be automatically added to the backup. Default value is `false`.
func (o BackupOutput) ExcludeNewRepositories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Backup) pulumi.BoolPtrOutput { return v.ExcludeNewRepositories }).(pulumi.BoolPtrOutput)
}

// A list of excluded repositories from the backup. Default is empty list.
func (o BackupOutput) ExcludedRepositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringArrayOutput { return v.ExcludedRepositories }).(pulumi.StringArrayOutput)
}

// When set to true, mission control will not be automatically added to the backup. Default value is `false`.
func (o BackupOutput) ExportMissionControl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Backup) pulumi.BoolPtrOutput { return v.ExportMissionControl }).(pulumi.BoolPtrOutput)
}

// The unique ID of the artifactory backup config.
func (o BackupOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
func (o BackupOutput) RetentionPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Backup) pulumi.IntPtrOutput { return v.RetentionPeriodHours }).(pulumi.IntPtrOutput)
}

// If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
func (o BackupOutput) SendMailOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Backup) pulumi.BoolPtrOutput { return v.SendMailOnError }).(pulumi.BoolPtrOutput)
}

// If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
func (o BackupOutput) VerifyDiskSpace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Backup) pulumi.BoolPtrOutput { return v.VerifyDiskSpace }).(pulumi.BoolPtrOutput)
}

type BackupArrayOutput struct{ *pulumi.OutputState }

func (BackupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Backup)(nil)).Elem()
}

func (o BackupArrayOutput) ToBackupArrayOutput() BackupArrayOutput {
	return o
}

func (o BackupArrayOutput) ToBackupArrayOutputWithContext(ctx context.Context) BackupArrayOutput {
	return o
}

func (o BackupArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Backup] {
	return pulumix.Output[[]*Backup]{
		OutputState: o.OutputState,
	}
}

func (o BackupArrayOutput) Index(i pulumi.IntInput) BackupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Backup {
		return vs[0].([]*Backup)[vs[1].(int)]
	}).(BackupOutput)
}

type BackupMapOutput struct{ *pulumi.OutputState }

func (BackupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Backup)(nil)).Elem()
}

func (o BackupMapOutput) ToBackupMapOutput() BackupMapOutput {
	return o
}

func (o BackupMapOutput) ToBackupMapOutputWithContext(ctx context.Context) BackupMapOutput {
	return o
}

func (o BackupMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Backup] {
	return pulumix.Output[map[string]*Backup]{
		OutputState: o.OutputState,
	}
}

func (o BackupMapOutput) MapIndex(k pulumi.StringInput) BackupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Backup {
		return vs[0].(map[string]*Backup)[vs[1].(string)]
	}).(BackupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupInput)(nil)).Elem(), &Backup{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupArrayInput)(nil)).Elem(), BackupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupMapInput)(nil)).Elem(), BackupMap{})
	pulumi.RegisterOutputType(BackupOutput{})
	pulumi.RegisterOutputType(BackupArrayOutput{})
	pulumi.RegisterOutputType(BackupMapOutput{})
}

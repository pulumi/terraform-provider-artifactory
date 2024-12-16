# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['BackupArgs', 'Backup']

@pulumi.input_type
class BackupArgs:
    def __init__(__self__, *,
                 cron_exp: pulumi.Input[str],
                 key: pulumi.Input[str],
                 create_archive: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 exclude_new_repositories: Optional[pulumi.Input[bool]] = None,
                 excluded_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 export_mission_control: Optional[pulumi.Input[bool]] = None,
                 retention_period_hours: Optional[pulumi.Input[int]] = None,
                 send_mail_on_error: Optional[pulumi.Input[bool]] = None,
                 verify_disk_space: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Backup resource.
        :param pulumi.Input[str] cron_exp: A valid CRON expression that you can use to control backup frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. **Note:** please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
        :param pulumi.Input[bool] create_archive: If set to true, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`
        :param pulumi.Input[bool] enabled: Flag to enable or disable the backup config. Default value is `true`.
        :param pulumi.Input[bool] exclude_new_repositories: When set to true, new repositories will not be automatically added to the backup. Default value is `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_repositories: List of excluded repositories from the backup.
        :param pulumi.Input[bool] export_mission_control: When set to true, mission control will not be automatically added to the backup. Default value is `false`.
        :param pulumi.Input[int] retention_period_hours: The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours i.e. 7 days.
        :param pulumi.Input[bool] send_mail_on_error: If set to true, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        :param pulumi.Input[bool] verify_disk_space: If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups. Default value is `false`.
        """
        pulumi.set(__self__, "cron_exp", cron_exp)
        pulumi.set(__self__, "key", key)
        if create_archive is not None:
            pulumi.set(__self__, "create_archive", create_archive)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if exclude_new_repositories is not None:
            pulumi.set(__self__, "exclude_new_repositories", exclude_new_repositories)
        if excluded_repositories is not None:
            pulumi.set(__self__, "excluded_repositories", excluded_repositories)
        if export_mission_control is not None:
            pulumi.set(__self__, "export_mission_control", export_mission_control)
        if retention_period_hours is not None:
            pulumi.set(__self__, "retention_period_hours", retention_period_hours)
        if send_mail_on_error is not None:
            pulumi.set(__self__, "send_mail_on_error", send_mail_on_error)
        if verify_disk_space is not None:
            pulumi.set(__self__, "verify_disk_space", verify_disk_space)

    @property
    @pulumi.getter(name="cronExp")
    def cron_exp(self) -> pulumi.Input[str]:
        """
        A valid CRON expression that you can use to control backup frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. **Note:** please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
        """
        return pulumi.get(self, "cron_exp")

    @cron_exp.setter
    def cron_exp(self, value: pulumi.Input[str]):
        pulumi.set(self, "cron_exp", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="createArchive")
    def create_archive(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`
        """
        return pulumi.get(self, "create_archive")

    @create_archive.setter
    def create_archive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_archive", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to enable or disable the backup config. Default value is `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="excludeNewRepositories")
    def exclude_new_repositories(self) -> Optional[pulumi.Input[bool]]:
        """
        When set to true, new repositories will not be automatically added to the backup. Default value is `false`.
        """
        return pulumi.get(self, "exclude_new_repositories")

    @exclude_new_repositories.setter
    def exclude_new_repositories(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_new_repositories", value)

    @property
    @pulumi.getter(name="excludedRepositories")
    def excluded_repositories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of excluded repositories from the backup.
        """
        return pulumi.get(self, "excluded_repositories")

    @excluded_repositories.setter
    def excluded_repositories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_repositories", value)

    @property
    @pulumi.getter(name="exportMissionControl")
    def export_mission_control(self) -> Optional[pulumi.Input[bool]]:
        """
        When set to true, mission control will not be automatically added to the backup. Default value is `false`.
        """
        return pulumi.get(self, "export_mission_control")

    @export_mission_control.setter
    def export_mission_control(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "export_mission_control", value)

    @property
    @pulumi.getter(name="retentionPeriodHours")
    def retention_period_hours(self) -> Optional[pulumi.Input[int]]:
        """
        The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours i.e. 7 days.
        """
        return pulumi.get(self, "retention_period_hours")

    @retention_period_hours.setter
    def retention_period_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_period_hours", value)

    @property
    @pulumi.getter(name="sendMailOnError")
    def send_mail_on_error(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        """
        return pulumi.get(self, "send_mail_on_error")

    @send_mail_on_error.setter
    def send_mail_on_error(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_mail_on_error", value)

    @property
    @pulumi.getter(name="verifyDiskSpace")
    def verify_disk_space(self) -> Optional[pulumi.Input[bool]]:
        """
        If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups. Default value is `false`.
        """
        return pulumi.get(self, "verify_disk_space")

    @verify_disk_space.setter
    def verify_disk_space(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verify_disk_space", value)


@pulumi.input_type
class _BackupState:
    def __init__(__self__, *,
                 create_archive: Optional[pulumi.Input[bool]] = None,
                 cron_exp: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 exclude_new_repositories: Optional[pulumi.Input[bool]] = None,
                 excluded_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 export_mission_control: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 retention_period_hours: Optional[pulumi.Input[int]] = None,
                 send_mail_on_error: Optional[pulumi.Input[bool]] = None,
                 verify_disk_space: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering Backup resources.
        :param pulumi.Input[bool] create_archive: If set to true, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`
        :param pulumi.Input[str] cron_exp: A valid CRON expression that you can use to control backup frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. **Note:** please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
        :param pulumi.Input[bool] enabled: Flag to enable or disable the backup config. Default value is `true`.
        :param pulumi.Input[bool] exclude_new_repositories: When set to true, new repositories will not be automatically added to the backup. Default value is `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_repositories: List of excluded repositories from the backup.
        :param pulumi.Input[bool] export_mission_control: When set to true, mission control will not be automatically added to the backup. Default value is `false`.
        :param pulumi.Input[int] retention_period_hours: The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours i.e. 7 days.
        :param pulumi.Input[bool] send_mail_on_error: If set to true, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        :param pulumi.Input[bool] verify_disk_space: If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups. Default value is `false`.
        """
        if create_archive is not None:
            pulumi.set(__self__, "create_archive", create_archive)
        if cron_exp is not None:
            pulumi.set(__self__, "cron_exp", cron_exp)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if exclude_new_repositories is not None:
            pulumi.set(__self__, "exclude_new_repositories", exclude_new_repositories)
        if excluded_repositories is not None:
            pulumi.set(__self__, "excluded_repositories", excluded_repositories)
        if export_mission_control is not None:
            pulumi.set(__self__, "export_mission_control", export_mission_control)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if retention_period_hours is not None:
            pulumi.set(__self__, "retention_period_hours", retention_period_hours)
        if send_mail_on_error is not None:
            pulumi.set(__self__, "send_mail_on_error", send_mail_on_error)
        if verify_disk_space is not None:
            pulumi.set(__self__, "verify_disk_space", verify_disk_space)

    @property
    @pulumi.getter(name="createArchive")
    def create_archive(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`
        """
        return pulumi.get(self, "create_archive")

    @create_archive.setter
    def create_archive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_archive", value)

    @property
    @pulumi.getter(name="cronExp")
    def cron_exp(self) -> Optional[pulumi.Input[str]]:
        """
        A valid CRON expression that you can use to control backup frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. **Note:** please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
        """
        return pulumi.get(self, "cron_exp")

    @cron_exp.setter
    def cron_exp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cron_exp", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to enable or disable the backup config. Default value is `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="excludeNewRepositories")
    def exclude_new_repositories(self) -> Optional[pulumi.Input[bool]]:
        """
        When set to true, new repositories will not be automatically added to the backup. Default value is `false`.
        """
        return pulumi.get(self, "exclude_new_repositories")

    @exclude_new_repositories.setter
    def exclude_new_repositories(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_new_repositories", value)

    @property
    @pulumi.getter(name="excludedRepositories")
    def excluded_repositories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of excluded repositories from the backup.
        """
        return pulumi.get(self, "excluded_repositories")

    @excluded_repositories.setter
    def excluded_repositories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_repositories", value)

    @property
    @pulumi.getter(name="exportMissionControl")
    def export_mission_control(self) -> Optional[pulumi.Input[bool]]:
        """
        When set to true, mission control will not be automatically added to the backup. Default value is `false`.
        """
        return pulumi.get(self, "export_mission_control")

    @export_mission_control.setter
    def export_mission_control(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "export_mission_control", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="retentionPeriodHours")
    def retention_period_hours(self) -> Optional[pulumi.Input[int]]:
        """
        The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours i.e. 7 days.
        """
        return pulumi.get(self, "retention_period_hours")

    @retention_period_hours.setter
    def retention_period_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_period_hours", value)

    @property
    @pulumi.getter(name="sendMailOnError")
    def send_mail_on_error(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        """
        return pulumi.get(self, "send_mail_on_error")

    @send_mail_on_error.setter
    def send_mail_on_error(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_mail_on_error", value)

    @property
    @pulumi.getter(name="verifyDiskSpace")
    def verify_disk_space(self) -> Optional[pulumi.Input[bool]]:
        """
        If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups. Default value is `false`.
        """
        return pulumi.get(self, "verify_disk_space")

    @verify_disk_space.setter
    def verify_disk_space(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verify_disk_space", value)


class Backup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_archive: Optional[pulumi.Input[bool]] = None,
                 cron_exp: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 exclude_new_repositories: Optional[pulumi.Input[bool]] = None,
                 excluded_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 export_mission_control: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 retention_period_hours: Optional[pulumi.Input[int]] = None,
                 send_mail_on_error: Optional[pulumi.Input[bool]] = None,
                 verify_disk_space: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        This resource can be used to manage the automatic and periodic backups of the entire Artifactory instance.

        When an `Backup` resource is configured and enabled to true, backup of the entire Artifactory system will be done automatically and periodically.

        The backup process creates a time-stamped directory in the target backup directory.

        See [JFrog Artifactory Backup](https://www.jfrog.com/confluence/display/JFROG/Backups) for more details.

        ~>Only supported in self-hosted environment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        backup_config_name = artifactory.Backup("backup_config_name",
            key="backup_config_name",
            enabled=True,
            cron_exp="0 0 12 * * ? *",
            retention_period_hours=1000,
            excluded_repositories=["my-docker-local"],
            create_archive=False,
            exclude_new_repositories=True,
            send_mail_on_error=True,
            verify_disk_space=True,
            export_mission_control=True)
        ```

        ## Import

        ```sh
        $ pulumi import artifactory:index/backup:Backup backup_name backup_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] create_archive: If set to true, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`
        :param pulumi.Input[str] cron_exp: A valid CRON expression that you can use to control backup frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. **Note:** please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
        :param pulumi.Input[bool] enabled: Flag to enable or disable the backup config. Default value is `true`.
        :param pulumi.Input[bool] exclude_new_repositories: When set to true, new repositories will not be automatically added to the backup. Default value is `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_repositories: List of excluded repositories from the backup.
        :param pulumi.Input[bool] export_mission_control: When set to true, mission control will not be automatically added to the backup. Default value is `false`.
        :param pulumi.Input[int] retention_period_hours: The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours i.e. 7 days.
        :param pulumi.Input[bool] send_mail_on_error: If set to true, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        :param pulumi.Input[bool] verify_disk_space: If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups. Default value is `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can be used to manage the automatic and periodic backups of the entire Artifactory instance.

        When an `Backup` resource is configured and enabled to true, backup of the entire Artifactory system will be done automatically and periodically.

        The backup process creates a time-stamped directory in the target backup directory.

        See [JFrog Artifactory Backup](https://www.jfrog.com/confluence/display/JFROG/Backups) for more details.

        ~>Only supported in self-hosted environment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        backup_config_name = artifactory.Backup("backup_config_name",
            key="backup_config_name",
            enabled=True,
            cron_exp="0 0 12 * * ? *",
            retention_period_hours=1000,
            excluded_repositories=["my-docker-local"],
            create_archive=False,
            exclude_new_repositories=True,
            send_mail_on_error=True,
            verify_disk_space=True,
            export_mission_control=True)
        ```

        ## Import

        ```sh
        $ pulumi import artifactory:index/backup:Backup backup_name backup_name
        ```

        :param str resource_name: The name of the resource.
        :param BackupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_archive: Optional[pulumi.Input[bool]] = None,
                 cron_exp: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 exclude_new_repositories: Optional[pulumi.Input[bool]] = None,
                 excluded_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 export_mission_control: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 retention_period_hours: Optional[pulumi.Input[int]] = None,
                 send_mail_on_error: Optional[pulumi.Input[bool]] = None,
                 verify_disk_space: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackupArgs.__new__(BackupArgs)

            __props__.__dict__["create_archive"] = create_archive
            if cron_exp is None and not opts.urn:
                raise TypeError("Missing required property 'cron_exp'")
            __props__.__dict__["cron_exp"] = cron_exp
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["exclude_new_repositories"] = exclude_new_repositories
            __props__.__dict__["excluded_repositories"] = excluded_repositories
            __props__.__dict__["export_mission_control"] = export_mission_control
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["retention_period_hours"] = retention_period_hours
            __props__.__dict__["send_mail_on_error"] = send_mail_on_error
            __props__.__dict__["verify_disk_space"] = verify_disk_space
        super(Backup, __self__).__init__(
            'artifactory:index/backup:Backup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_archive: Optional[pulumi.Input[bool]] = None,
            cron_exp: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            exclude_new_repositories: Optional[pulumi.Input[bool]] = None,
            excluded_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            export_mission_control: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            retention_period_hours: Optional[pulumi.Input[int]] = None,
            send_mail_on_error: Optional[pulumi.Input[bool]] = None,
            verify_disk_space: Optional[pulumi.Input[bool]] = None) -> 'Backup':
        """
        Get an existing Backup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] create_archive: If set to true, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`
        :param pulumi.Input[str] cron_exp: A valid CRON expression that you can use to control backup frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. **Note:** please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
        :param pulumi.Input[bool] enabled: Flag to enable or disable the backup config. Default value is `true`.
        :param pulumi.Input[bool] exclude_new_repositories: When set to true, new repositories will not be automatically added to the backup. Default value is `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_repositories: List of excluded repositories from the backup.
        :param pulumi.Input[bool] export_mission_control: When set to true, mission control will not be automatically added to the backup. Default value is `false`.
        :param pulumi.Input[int] retention_period_hours: The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours i.e. 7 days.
        :param pulumi.Input[bool] send_mail_on_error: If set to true, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        :param pulumi.Input[bool] verify_disk_space: If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups. Default value is `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BackupState.__new__(_BackupState)

        __props__.__dict__["create_archive"] = create_archive
        __props__.__dict__["cron_exp"] = cron_exp
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["exclude_new_repositories"] = exclude_new_repositories
        __props__.__dict__["excluded_repositories"] = excluded_repositories
        __props__.__dict__["export_mission_control"] = export_mission_control
        __props__.__dict__["key"] = key
        __props__.__dict__["retention_period_hours"] = retention_period_hours
        __props__.__dict__["send_mail_on_error"] = send_mail_on_error
        __props__.__dict__["verify_disk_space"] = verify_disk_space
        return Backup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createArchive")
    def create_archive(self) -> pulumi.Output[bool]:
        """
        If set to true, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`
        """
        return pulumi.get(self, "create_archive")

    @property
    @pulumi.getter(name="cronExp")
    def cron_exp(self) -> pulumi.Output[str]:
        """
        A valid CRON expression that you can use to control backup frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. **Note:** please use 7 character format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year. Also, specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) and in [Cronexp package readme](https://github.com/gorhill/cronexpr#other-details).
        """
        return pulumi.get(self, "cron_exp")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Flag to enable or disable the backup config. Default value is `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="excludeNewRepositories")
    def exclude_new_repositories(self) -> pulumi.Output[bool]:
        """
        When set to true, new repositories will not be automatically added to the backup. Default value is `false`.
        """
        return pulumi.get(self, "exclude_new_repositories")

    @property
    @pulumi.getter(name="excludedRepositories")
    def excluded_repositories(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of excluded repositories from the backup.
        """
        return pulumi.get(self, "excluded_repositories")

    @property
    @pulumi.getter(name="exportMissionControl")
    def export_mission_control(self) -> pulumi.Output[bool]:
        """
        When set to true, mission control will not be automatically added to the backup. Default value is `false`.
        """
        return pulumi.get(self, "export_mission_control")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="retentionPeriodHours")
    def retention_period_hours(self) -> pulumi.Output[int]:
        """
        The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours i.e. 7 days.
        """
        return pulumi.get(self, "retention_period_hours")

    @property
    @pulumi.getter(name="sendMailOnError")
    def send_mail_on_error(self) -> pulumi.Output[bool]:
        """
        If set to true, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        """
        return pulumi.get(self, "send_mail_on_error")

    @property
    @pulumi.getter(name="verifyDiskSpace")
    def verify_disk_space(self) -> pulumi.Output[bool]:
        """
        If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups. Default value is `false`.
        """
        return pulumi.get(self, "verify_disk_space")


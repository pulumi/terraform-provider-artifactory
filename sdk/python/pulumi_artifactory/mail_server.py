# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MailServerArgs', 'MailServer']

@pulumi.input_type
class MailServerArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 host: pulumi.Input[str],
                 port: pulumi.Input[int],
                 artifactory_url: Optional[pulumi.Input[str]] = None,
                 from_: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 subject_prefix: Optional[pulumi.Input[str]] = None,
                 use_ssl: Optional[pulumi.Input[bool]] = None,
                 use_tls: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MailServer resource.
        :param pulumi.Input[bool] enabled: When set, mail notifications are enabled.
        :param pulumi.Input[str] host: The mail server IP address / DNS.
        :param pulumi.Input[int] port: The port number of the mail server.
        :param pulumi.Input[str] artifactory_url: The Artifactory URL to to link to in all outgoing messages.
        :param pulumi.Input[str] from_: The 'from' address header to use in all outgoing messages.
        :param pulumi.Input[str] password: The password for authentication with the mail server.
        :param pulumi.Input[str] subject_prefix: A prefix to use for the subject of all outgoing mails.
        :param pulumi.Input[bool] use_ssl: When set to 'true', uses a secure connection to the mail server.
        :param pulumi.Input[bool] use_tls: When set to 'true', uses Transport Layer Security when connecting to the mail server.
        :param pulumi.Input[str] username: The username for authentication with the mail server.
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "host", host)
        pulumi.set(__self__, "port", port)
        if artifactory_url is not None:
            pulumi.set(__self__, "artifactory_url", artifactory_url)
        if from_ is not None:
            pulumi.set(__self__, "from_", from_)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if subject_prefix is not None:
            pulumi.set(__self__, "subject_prefix", subject_prefix)
        if use_ssl is not None:
            pulumi.set(__self__, "use_ssl", use_ssl)
        if use_tls is not None:
            pulumi.set(__self__, "use_tls", use_tls)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        When set, mail notifications are enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Input[str]:
        """
        The mail server IP address / DNS.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: pulumi.Input[str]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        The port number of the mail server.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="artifactoryUrl")
    def artifactory_url(self) -> Optional[pulumi.Input[str]]:
        """
        The Artifactory URL to to link to in all outgoing messages.
        """
        return pulumi.get(self, "artifactory_url")

    @artifactory_url.setter
    def artifactory_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "artifactory_url", value)

    @property
    @pulumi.getter(name="from")
    def from_(self) -> Optional[pulumi.Input[str]]:
        """
        The 'from' address header to use in all outgoing messages.
        """
        return pulumi.get(self, "from_")

    @from_.setter
    def from_(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "from_", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password for authentication with the mail server.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="subjectPrefix")
    def subject_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        A prefix to use for the subject of all outgoing mails.
        """
        return pulumi.get(self, "subject_prefix")

    @subject_prefix.setter
    def subject_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subject_prefix", value)

    @property
    @pulumi.getter(name="useSsl")
    def use_ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        When set to 'true', uses a secure connection to the mail server.
        """
        return pulumi.get(self, "use_ssl")

    @use_ssl.setter
    def use_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_ssl", value)

    @property
    @pulumi.getter(name="useTls")
    def use_tls(self) -> Optional[pulumi.Input[bool]]:
        """
        When set to 'true', uses Transport Layer Security when connecting to the mail server.
        """
        return pulumi.get(self, "use_tls")

    @use_tls.setter
    def use_tls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_tls", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username for authentication with the mail server.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _MailServerState:
    def __init__(__self__, *,
                 artifactory_url: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 from_: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 subject_prefix: Optional[pulumi.Input[str]] = None,
                 use_ssl: Optional[pulumi.Input[bool]] = None,
                 use_tls: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MailServer resources.
        :param pulumi.Input[str] artifactory_url: The Artifactory URL to to link to in all outgoing messages.
        :param pulumi.Input[bool] enabled: When set, mail notifications are enabled.
        :param pulumi.Input[str] from_: The 'from' address header to use in all outgoing messages.
        :param pulumi.Input[str] host: The mail server IP address / DNS.
        :param pulumi.Input[str] password: The password for authentication with the mail server.
        :param pulumi.Input[int] port: The port number of the mail server.
        :param pulumi.Input[str] subject_prefix: A prefix to use for the subject of all outgoing mails.
        :param pulumi.Input[bool] use_ssl: When set to 'true', uses a secure connection to the mail server.
        :param pulumi.Input[bool] use_tls: When set to 'true', uses Transport Layer Security when connecting to the mail server.
        :param pulumi.Input[str] username: The username for authentication with the mail server.
        """
        if artifactory_url is not None:
            pulumi.set(__self__, "artifactory_url", artifactory_url)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if from_ is not None:
            pulumi.set(__self__, "from_", from_)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if subject_prefix is not None:
            pulumi.set(__self__, "subject_prefix", subject_prefix)
        if use_ssl is not None:
            pulumi.set(__self__, "use_ssl", use_ssl)
        if use_tls is not None:
            pulumi.set(__self__, "use_tls", use_tls)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="artifactoryUrl")
    def artifactory_url(self) -> Optional[pulumi.Input[str]]:
        """
        The Artifactory URL to to link to in all outgoing messages.
        """
        return pulumi.get(self, "artifactory_url")

    @artifactory_url.setter
    def artifactory_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "artifactory_url", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, mail notifications are enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="from")
    def from_(self) -> Optional[pulumi.Input[str]]:
        """
        The 'from' address header to use in all outgoing messages.
        """
        return pulumi.get(self, "from_")

    @from_.setter
    def from_(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "from_", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The mail server IP address / DNS.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password for authentication with the mail server.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port number of the mail server.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="subjectPrefix")
    def subject_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        A prefix to use for the subject of all outgoing mails.
        """
        return pulumi.get(self, "subject_prefix")

    @subject_prefix.setter
    def subject_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subject_prefix", value)

    @property
    @pulumi.getter(name="useSsl")
    def use_ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        When set to 'true', uses a secure connection to the mail server.
        """
        return pulumi.get(self, "use_ssl")

    @use_ssl.setter
    def use_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_ssl", value)

    @property
    @pulumi.getter(name="useTls")
    def use_tls(self) -> Optional[pulumi.Input[bool]]:
        """
        When set to 'true', uses Transport Layer Security when connecting to the mail server.
        """
        return pulumi.get(self, "use_tls")

    @use_tls.setter
    def use_tls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_tls", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username for authentication with the mail server.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class MailServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 artifactory_url: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 from_: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 subject_prefix: Optional[pulumi.Input[str]] = None,
                 use_ssl: Optional[pulumi.Input[bool]] = None,
                 use_tls: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an Artifactory Mail Server resource. This can be used to create and manage Artifactory mail server configuration.

        ## Example Usage

        ### S

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        mymailserver = artifactory.MailServer("mymailserver",
            artifactory_url="http://tempurl.org",
            enabled=True,
            from_="test@jfrog.com",
            host="http://tempurl.org",
            password="test-password",
            port=25,
            subject_prefix="[Test]",
            use_ssl=True,
            use_tls=True,
            username="test-user")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ```sh
        $ pulumi import artifactory:index/mailServer:MailServer my-mail-server mymailserver
        ```

        ~>The `password` attribute is not retrievable from Artifactory thus there will be state drift after importing this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] artifactory_url: The Artifactory URL to to link to in all outgoing messages.
        :param pulumi.Input[bool] enabled: When set, mail notifications are enabled.
        :param pulumi.Input[str] from_: The 'from' address header to use in all outgoing messages.
        :param pulumi.Input[str] host: The mail server IP address / DNS.
        :param pulumi.Input[str] password: The password for authentication with the mail server.
        :param pulumi.Input[int] port: The port number of the mail server.
        :param pulumi.Input[str] subject_prefix: A prefix to use for the subject of all outgoing mails.
        :param pulumi.Input[bool] use_ssl: When set to 'true', uses a secure connection to the mail server.
        :param pulumi.Input[bool] use_tls: When set to 'true', uses Transport Layer Security when connecting to the mail server.
        :param pulumi.Input[str] username: The username for authentication with the mail server.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MailServerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory Mail Server resource. This can be used to create and manage Artifactory mail server configuration.

        ## Example Usage

        ### S

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        mymailserver = artifactory.MailServer("mymailserver",
            artifactory_url="http://tempurl.org",
            enabled=True,
            from_="test@jfrog.com",
            host="http://tempurl.org",
            password="test-password",
            port=25,
            subject_prefix="[Test]",
            use_ssl=True,
            use_tls=True,
            username="test-user")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ```sh
        $ pulumi import artifactory:index/mailServer:MailServer my-mail-server mymailserver
        ```

        ~>The `password` attribute is not retrievable from Artifactory thus there will be state drift after importing this resource.

        :param str resource_name: The name of the resource.
        :param MailServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MailServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 artifactory_url: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 from_: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 subject_prefix: Optional[pulumi.Input[str]] = None,
                 use_ssl: Optional[pulumi.Input[bool]] = None,
                 use_tls: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MailServerArgs.__new__(MailServerArgs)

            __props__.__dict__["artifactory_url"] = artifactory_url
            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["from_"] = from_
            if host is None and not opts.urn:
                raise TypeError("Missing required property 'host'")
            __props__.__dict__["host"] = host
            __props__.__dict__["password"] = password
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            __props__.__dict__["subject_prefix"] = subject_prefix
            __props__.__dict__["use_ssl"] = use_ssl
            __props__.__dict__["use_tls"] = use_tls
            __props__.__dict__["username"] = username
        super(MailServer, __self__).__init__(
            'artifactory:index/mailServer:MailServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            artifactory_url: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            from_: Optional[pulumi.Input[str]] = None,
            host: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            subject_prefix: Optional[pulumi.Input[str]] = None,
            use_ssl: Optional[pulumi.Input[bool]] = None,
            use_tls: Optional[pulumi.Input[bool]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'MailServer':
        """
        Get an existing MailServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] artifactory_url: The Artifactory URL to to link to in all outgoing messages.
        :param pulumi.Input[bool] enabled: When set, mail notifications are enabled.
        :param pulumi.Input[str] from_: The 'from' address header to use in all outgoing messages.
        :param pulumi.Input[str] host: The mail server IP address / DNS.
        :param pulumi.Input[str] password: The password for authentication with the mail server.
        :param pulumi.Input[int] port: The port number of the mail server.
        :param pulumi.Input[str] subject_prefix: A prefix to use for the subject of all outgoing mails.
        :param pulumi.Input[bool] use_ssl: When set to 'true', uses a secure connection to the mail server.
        :param pulumi.Input[bool] use_tls: When set to 'true', uses Transport Layer Security when connecting to the mail server.
        :param pulumi.Input[str] username: The username for authentication with the mail server.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MailServerState.__new__(_MailServerState)

        __props__.__dict__["artifactory_url"] = artifactory_url
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["from_"] = from_
        __props__.__dict__["host"] = host
        __props__.__dict__["password"] = password
        __props__.__dict__["port"] = port
        __props__.__dict__["subject_prefix"] = subject_prefix
        __props__.__dict__["use_ssl"] = use_ssl
        __props__.__dict__["use_tls"] = use_tls
        __props__.__dict__["username"] = username
        return MailServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="artifactoryUrl")
    def artifactory_url(self) -> pulumi.Output[Optional[str]]:
        """
        The Artifactory URL to to link to in all outgoing messages.
        """
        return pulumi.get(self, "artifactory_url")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        When set, mail notifications are enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="from")
    def from_(self) -> pulumi.Output[Optional[str]]:
        """
        The 'from' address header to use in all outgoing messages.
        """
        return pulumi.get(self, "from_")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        The mail server IP address / DNS.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password for authentication with the mail server.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        The port number of the mail server.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="subjectPrefix")
    def subject_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        A prefix to use for the subject of all outgoing mails.
        """
        return pulumi.get(self, "subject_prefix")

    @property
    @pulumi.getter(name="useSsl")
    def use_ssl(self) -> pulumi.Output[bool]:
        """
        When set to 'true', uses a secure connection to the mail server.
        """
        return pulumi.get(self, "use_ssl")

    @property
    @pulumi.getter(name="useTls")
    def use_tls(self) -> pulumi.Output[bool]:
        """
        When set to 'true', uses Transport Layer Security when connecting to the mail server.
        """
        return pulumi.get(self, "use_tls")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        The username for authentication with the mail server.
        """
        return pulumi.get(self, "username")


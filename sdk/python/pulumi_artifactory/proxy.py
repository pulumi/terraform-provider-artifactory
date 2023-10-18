# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProxyArgs', 'Proxy']

@pulumi.input_type
class ProxyArgs:
    def __init__(__self__, *,
                 host: pulumi.Input[str],
                 key: pulumi.Input[str],
                 port: pulumi.Input[int],
                 nt_domain: Optional[pulumi.Input[str]] = None,
                 nt_host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 platform_default: Optional[pulumi.Input[bool]] = None,
                 redirect_to_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Proxy resource.
        :param pulumi.Input[str] host: The name of the proxy host.
        :param pulumi.Input[str] key: The unique ID of the proxy.
        :param pulumi.Input[int] port: The proxy port number.
        :param pulumi.Input[str] nt_domain: The proxy domain/realm name.
        :param pulumi.Input[str] nt_host: The computer name of the machine (the machine connecting to the NTLM proxy).
        :param pulumi.Input[str] password: The proxy password when authentication credentials are required.
        :param pulumi.Input[bool] platform_default: When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] redirect_to_hosts: An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
        :param pulumi.Input[str] username: The proxy username when authentication credentials are required.
        """
        ProxyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            host=host,
            key=key,
            port=port,
            nt_domain=nt_domain,
            nt_host=nt_host,
            password=password,
            platform_default=platform_default,
            redirect_to_hosts=redirect_to_hosts,
            services=services,
            username=username,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             host: pulumi.Input[str],
             key: pulumi.Input[str],
             port: pulumi.Input[int],
             nt_domain: Optional[pulumi.Input[str]] = None,
             nt_host: Optional[pulumi.Input[str]] = None,
             password: Optional[pulumi.Input[str]] = None,
             platform_default: Optional[pulumi.Input[bool]] = None,
             redirect_to_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             username: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'ntDomain' in kwargs:
            nt_domain = kwargs['ntDomain']
        if 'ntHost' in kwargs:
            nt_host = kwargs['ntHost']
        if 'platformDefault' in kwargs:
            platform_default = kwargs['platformDefault']
        if 'redirectToHosts' in kwargs:
            redirect_to_hosts = kwargs['redirectToHosts']

        _setter("host", host)
        _setter("key", key)
        _setter("port", port)
        if nt_domain is not None:
            _setter("nt_domain", nt_domain)
        if nt_host is not None:
            _setter("nt_host", nt_host)
        if password is not None:
            _setter("password", password)
        if platform_default is not None:
            _setter("platform_default", platform_default)
        if redirect_to_hosts is not None:
            _setter("redirect_to_hosts", redirect_to_hosts)
        if services is not None:
            _setter("services", services)
        if username is not None:
            _setter("username", username)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Input[str]:
        """
        The name of the proxy host.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: pulumi.Input[str]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The unique ID of the proxy.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        The proxy port number.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="ntDomain")
    def nt_domain(self) -> Optional[pulumi.Input[str]]:
        """
        The proxy domain/realm name.
        """
        return pulumi.get(self, "nt_domain")

    @nt_domain.setter
    def nt_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nt_domain", value)

    @property
    @pulumi.getter(name="ntHost")
    def nt_host(self) -> Optional[pulumi.Input[str]]:
        """
        The computer name of the machine (the machine connecting to the NTLM proxy).
        """
        return pulumi.get(self, "nt_host")

    @nt_host.setter
    def nt_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nt_host", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The proxy password when authentication credentials are required.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="platformDefault")
    def platform_default(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
        """
        return pulumi.get(self, "platform_default")

    @platform_default.setter
    def platform_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "platform_default", value)

    @property
    @pulumi.getter(name="redirectToHosts")
    def redirect_to_hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
        """
        return pulumi.get(self, "redirect_to_hosts")

    @redirect_to_hosts.setter
    def redirect_to_hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "redirect_to_hosts", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The proxy username when authentication credentials are required.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _ProxyState:
    def __init__(__self__, *,
                 host: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 nt_domain: Optional[pulumi.Input[str]] = None,
                 nt_host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 platform_default: Optional[pulumi.Input[bool]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 redirect_to_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Proxy resources.
        :param pulumi.Input[str] host: The name of the proxy host.
        :param pulumi.Input[str] key: The unique ID of the proxy.
        :param pulumi.Input[str] nt_domain: The proxy domain/realm name.
        :param pulumi.Input[str] nt_host: The computer name of the machine (the machine connecting to the NTLM proxy).
        :param pulumi.Input[str] password: The proxy password when authentication credentials are required.
        :param pulumi.Input[bool] platform_default: When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
        :param pulumi.Input[int] port: The proxy port number.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] redirect_to_hosts: An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
        :param pulumi.Input[str] username: The proxy username when authentication credentials are required.
        """
        _ProxyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            host=host,
            key=key,
            nt_domain=nt_domain,
            nt_host=nt_host,
            password=password,
            platform_default=platform_default,
            port=port,
            redirect_to_hosts=redirect_to_hosts,
            services=services,
            username=username,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             host: Optional[pulumi.Input[str]] = None,
             key: Optional[pulumi.Input[str]] = None,
             nt_domain: Optional[pulumi.Input[str]] = None,
             nt_host: Optional[pulumi.Input[str]] = None,
             password: Optional[pulumi.Input[str]] = None,
             platform_default: Optional[pulumi.Input[bool]] = None,
             port: Optional[pulumi.Input[int]] = None,
             redirect_to_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             username: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'ntDomain' in kwargs:
            nt_domain = kwargs['ntDomain']
        if 'ntHost' in kwargs:
            nt_host = kwargs['ntHost']
        if 'platformDefault' in kwargs:
            platform_default = kwargs['platformDefault']
        if 'redirectToHosts' in kwargs:
            redirect_to_hosts = kwargs['redirectToHosts']

        if host is not None:
            _setter("host", host)
        if key is not None:
            _setter("key", key)
        if nt_domain is not None:
            _setter("nt_domain", nt_domain)
        if nt_host is not None:
            _setter("nt_host", nt_host)
        if password is not None:
            _setter("password", password)
        if platform_default is not None:
            _setter("platform_default", platform_default)
        if port is not None:
            _setter("port", port)
        if redirect_to_hosts is not None:
            _setter("redirect_to_hosts", redirect_to_hosts)
        if services is not None:
            _setter("services", services)
        if username is not None:
            _setter("username", username)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the proxy host.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ID of the proxy.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="ntDomain")
    def nt_domain(self) -> Optional[pulumi.Input[str]]:
        """
        The proxy domain/realm name.
        """
        return pulumi.get(self, "nt_domain")

    @nt_domain.setter
    def nt_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nt_domain", value)

    @property
    @pulumi.getter(name="ntHost")
    def nt_host(self) -> Optional[pulumi.Input[str]]:
        """
        The computer name of the machine (the machine connecting to the NTLM proxy).
        """
        return pulumi.get(self, "nt_host")

    @nt_host.setter
    def nt_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nt_host", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The proxy password when authentication credentials are required.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="platformDefault")
    def platform_default(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
        """
        return pulumi.get(self, "platform_default")

    @platform_default.setter
    def platform_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "platform_default", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The proxy port number.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="redirectToHosts")
    def redirect_to_hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
        """
        return pulumi.get(self, "redirect_to_hosts")

    @redirect_to_hosts.setter
    def redirect_to_hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "redirect_to_hosts", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The proxy username when authentication credentials are required.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class Proxy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 nt_domain: Optional[pulumi.Input[str]] = None,
                 nt_host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 platform_default: Optional[pulumi.Input[bool]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 redirect_to_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an Artifactory Proxy resource.

        This resource configuration corresponds to 'proxies' config block in system configuration XML
        (REST endpoint: [artifactory/api/system/configuration](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GeneralConfiguration)).

        ~>The `Proxy` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        my_proxy = artifactory.Proxy("my-proxy",
            host="my-proxy.mycompany.com",
            key="my-proxy",
            nt_domain="MYCOMPANY",
            nt_host="MYCOMPANY.COM",
            password="password",
            platform_default=False,
            port=8888,
            redirect_to_hosts=["redirec-host.mycompany.com"],
            services=[
                "jfrt",
                "jfxr",
            ],
            username="user1")
        ```

        ## Import

        Current Proxy can be imported using `proxy-key` from Artifactory as the `ID`, e.g.

        ```sh
         $ pulumi import artifactory:index/proxy:Proxy my-proxy proxy-key
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: The name of the proxy host.
        :param pulumi.Input[str] key: The unique ID of the proxy.
        :param pulumi.Input[str] nt_domain: The proxy domain/realm name.
        :param pulumi.Input[str] nt_host: The computer name of the machine (the machine connecting to the NTLM proxy).
        :param pulumi.Input[str] password: The proxy password when authentication credentials are required.
        :param pulumi.Input[bool] platform_default: When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
        :param pulumi.Input[int] port: The proxy port number.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] redirect_to_hosts: An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
        :param pulumi.Input[str] username: The proxy username when authentication credentials are required.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProxyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory Proxy resource.

        This resource configuration corresponds to 'proxies' config block in system configuration XML
        (REST endpoint: [artifactory/api/system/configuration](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GeneralConfiguration)).

        ~>The `Proxy` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        my_proxy = artifactory.Proxy("my-proxy",
            host="my-proxy.mycompany.com",
            key="my-proxy",
            nt_domain="MYCOMPANY",
            nt_host="MYCOMPANY.COM",
            password="password",
            platform_default=False,
            port=8888,
            redirect_to_hosts=["redirec-host.mycompany.com"],
            services=[
                "jfrt",
                "jfxr",
            ],
            username="user1")
        ```

        ## Import

        Current Proxy can be imported using `proxy-key` from Artifactory as the `ID`, e.g.

        ```sh
         $ pulumi import artifactory:index/proxy:Proxy my-proxy proxy-key
        ```

        :param str resource_name: The name of the resource.
        :param ProxyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProxyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ProxyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 nt_domain: Optional[pulumi.Input[str]] = None,
                 nt_host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 platform_default: Optional[pulumi.Input[bool]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 redirect_to_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProxyArgs.__new__(ProxyArgs)

            if host is None and not opts.urn:
                raise TypeError("Missing required property 'host'")
            __props__.__dict__["host"] = host
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["nt_domain"] = nt_domain
            __props__.__dict__["nt_host"] = nt_host
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["platform_default"] = platform_default
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            __props__.__dict__["redirect_to_hosts"] = redirect_to_hosts
            __props__.__dict__["services"] = services
            __props__.__dict__["username"] = username
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Proxy, __self__).__init__(
            'artifactory:index/proxy:Proxy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            host: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            nt_domain: Optional[pulumi.Input[str]] = None,
            nt_host: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            platform_default: Optional[pulumi.Input[bool]] = None,
            port: Optional[pulumi.Input[int]] = None,
            redirect_to_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'Proxy':
        """
        Get an existing Proxy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: The name of the proxy host.
        :param pulumi.Input[str] key: The unique ID of the proxy.
        :param pulumi.Input[str] nt_domain: The proxy domain/realm name.
        :param pulumi.Input[str] nt_host: The computer name of the machine (the machine connecting to the NTLM proxy).
        :param pulumi.Input[str] password: The proxy password when authentication credentials are required.
        :param pulumi.Input[bool] platform_default: When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
        :param pulumi.Input[int] port: The proxy port number.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] redirect_to_hosts: An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
        :param pulumi.Input[str] username: The proxy username when authentication credentials are required.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProxyState.__new__(_ProxyState)

        __props__.__dict__["host"] = host
        __props__.__dict__["key"] = key
        __props__.__dict__["nt_domain"] = nt_domain
        __props__.__dict__["nt_host"] = nt_host
        __props__.__dict__["password"] = password
        __props__.__dict__["platform_default"] = platform_default
        __props__.__dict__["port"] = port
        __props__.__dict__["redirect_to_hosts"] = redirect_to_hosts
        __props__.__dict__["services"] = services
        __props__.__dict__["username"] = username
        return Proxy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        The name of the proxy host.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The unique ID of the proxy.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="ntDomain")
    def nt_domain(self) -> pulumi.Output[Optional[str]]:
        """
        The proxy domain/realm name.
        """
        return pulumi.get(self, "nt_domain")

    @property
    @pulumi.getter(name="ntHost")
    def nt_host(self) -> pulumi.Output[Optional[str]]:
        """
        The computer name of the machine (the machine connecting to the NTLM proxy).
        """
        return pulumi.get(self, "nt_host")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The proxy password when authentication credentials are required.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="platformDefault")
    def platform_default(self) -> pulumi.Output[Optional[bool]]:
        """
        When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
        """
        return pulumi.get(self, "platform_default")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        The proxy port number.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="redirectToHosts")
    def redirect_to_hosts(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
        """
        return pulumi.get(self, "redirect_to_hosts")

    @property
    @pulumi.getter
    def services(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        The proxy username when authentication credentials are required.
        """
        return pulumi.get(self, "username")


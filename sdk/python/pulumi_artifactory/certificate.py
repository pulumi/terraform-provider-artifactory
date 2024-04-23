# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CertificateArgs', 'Certificate']

@pulumi.input_type
class CertificateArgs:
    def __init__(__self__, *,
                 alias: pulumi.Input[str],
                 content: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Certificate resource.
        :param pulumi.Input[str] alias: Name of certificate.
        :param pulumi.Input[str] content: PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
        :param pulumi.Input[str] file: Path to the PEM file. Cannot be set with `content` attribute simultaneously.
        """
        pulumi.set(__self__, "alias", alias)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if file is not None:
            pulumi.set(__self__, "file", file)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Input[str]:
        """
        Name of certificate.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input[str]]:
        """
        Path to the PEM file. Cannot be set with `content` attribute simultaneously.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file", value)


@pulumi.input_type
class _CertificateState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 issued_by: Optional[pulumi.Input[str]] = None,
                 issued_on: Optional[pulumi.Input[str]] = None,
                 issued_to: Optional[pulumi.Input[str]] = None,
                 valid_until: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Certificate resources.
        :param pulumi.Input[str] alias: Name of certificate.
        :param pulumi.Input[str] content: PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
        :param pulumi.Input[str] file: Path to the PEM file. Cannot be set with `content` attribute simultaneously.
        :param pulumi.Input[str] fingerprint: SHA256 fingerprint of the certificate.
        :param pulumi.Input[str] issued_by: Name of the certificate authority that issued the certificate.
        :param pulumi.Input[str] issued_on: The time & date when the certificate is valid from.
        :param pulumi.Input[str] issued_to: Name of whom the certificate has been issued to.
        :param pulumi.Input[str] valid_until: The time & date when the certificate expires.
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if file is not None:
            pulumi.set(__self__, "file", file)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if issued_by is not None:
            pulumi.set(__self__, "issued_by", issued_by)
        if issued_on is not None:
            pulumi.set(__self__, "issued_on", issued_on)
        if issued_to is not None:
            pulumi.set(__self__, "issued_to", issued_to)
        if valid_until is not None:
            pulumi.set(__self__, "valid_until", valid_until)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        Name of certificate.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input[str]]:
        """
        Path to the PEM file. Cannot be set with `content` attribute simultaneously.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        SHA256 fingerprint of the certificate.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter(name="issuedBy")
    def issued_by(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the certificate authority that issued the certificate.
        """
        return pulumi.get(self, "issued_by")

    @issued_by.setter
    def issued_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issued_by", value)

    @property
    @pulumi.getter(name="issuedOn")
    def issued_on(self) -> Optional[pulumi.Input[str]]:
        """
        The time & date when the certificate is valid from.
        """
        return pulumi.get(self, "issued_on")

    @issued_on.setter
    def issued_on(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issued_on", value)

    @property
    @pulumi.getter(name="issuedTo")
    def issued_to(self) -> Optional[pulumi.Input[str]]:
        """
        Name of whom the certificate has been issued to.
        """
        return pulumi.get(self, "issued_to")

    @issued_to.setter
    def issued_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issued_to", value)

    @property
    @pulumi.getter(name="validUntil")
    def valid_until(self) -> Optional[pulumi.Input[str]]:
        """
        The time & date when the certificate expires.
        """
        return pulumi.get(self, "valid_until")

    @valid_until.setter
    def valid_until(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "valid_until", value)


class Certificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an Artifactory certificate resource. This can be used to create and manage Artifactory certificates which can be used as client authentication against remote repositories.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory
        import pulumi_std as std

        # Create a new Artifactory certificate called my-cert
        my_cert = artifactory.Certificate("my-cert",
            alias="my-cert",
            content=std.file(input="/path/to/bundle.pem").result)
        # This can then be used by a remote repository
        my_remote = artifactory.RemoteMavenRepository("my-remote", client_tls_certificate=my_cert.alias)
        ```

        ## Import

        Certificates can be imported using their alias, e.g.

        ```sh
        $ pulumi import artifactory:index/certificate:Certificate my-cert my-cert
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: Name of certificate.
        :param pulumi.Input[str] content: PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
        :param pulumi.Input[str] file: Path to the PEM file. Cannot be set with `content` attribute simultaneously.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory certificate resource. This can be used to create and manage Artifactory certificates which can be used as client authentication against remote repositories.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory
        import pulumi_std as std

        # Create a new Artifactory certificate called my-cert
        my_cert = artifactory.Certificate("my-cert",
            alias="my-cert",
            content=std.file(input="/path/to/bundle.pem").result)
        # This can then be used by a remote repository
        my_remote = artifactory.RemoteMavenRepository("my-remote", client_tls_certificate=my_cert.alias)
        ```

        ## Import

        Certificates can be imported using their alias, e.g.

        ```sh
        $ pulumi import artifactory:index/certificate:Certificate my-cert my-cert
        ```

        :param str resource_name: The name of the resource.
        :param CertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CertificateArgs.__new__(CertificateArgs)

            if alias is None and not opts.urn:
                raise TypeError("Missing required property 'alias'")
            __props__.__dict__["alias"] = alias
            __props__.__dict__["content"] = None if content is None else pulumi.Output.secret(content)
            __props__.__dict__["file"] = None if file is None else pulumi.Output.secret(file)
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["issued_by"] = None
            __props__.__dict__["issued_on"] = None
            __props__.__dict__["issued_to"] = None
            __props__.__dict__["valid_until"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["content", "file"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Certificate, __self__).__init__(
            'artifactory:index/certificate:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[str]] = None,
            content: Optional[pulumi.Input[str]] = None,
            file: Optional[pulumi.Input[str]] = None,
            fingerprint: Optional[pulumi.Input[str]] = None,
            issued_by: Optional[pulumi.Input[str]] = None,
            issued_on: Optional[pulumi.Input[str]] = None,
            issued_to: Optional[pulumi.Input[str]] = None,
            valid_until: Optional[pulumi.Input[str]] = None) -> 'Certificate':
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: Name of certificate.
        :param pulumi.Input[str] content: PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
        :param pulumi.Input[str] file: Path to the PEM file. Cannot be set with `content` attribute simultaneously.
        :param pulumi.Input[str] fingerprint: SHA256 fingerprint of the certificate.
        :param pulumi.Input[str] issued_by: Name of the certificate authority that issued the certificate.
        :param pulumi.Input[str] issued_on: The time & date when the certificate is valid from.
        :param pulumi.Input[str] issued_to: Name of whom the certificate has been issued to.
        :param pulumi.Input[str] valid_until: The time & date when the certificate expires.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CertificateState.__new__(_CertificateState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["content"] = content
        __props__.__dict__["file"] = file
        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["issued_by"] = issued_by
        __props__.__dict__["issued_on"] = issued_on
        __props__.__dict__["issued_to"] = issued_to
        __props__.__dict__["valid_until"] = valid_until
        return Certificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        """
        Name of certificate.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[Optional[str]]:
        """
        PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def file(self) -> pulumi.Output[Optional[str]]:
        """
        Path to the PEM file. Cannot be set with `content` attribute simultaneously.
        """
        return pulumi.get(self, "file")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        SHA256 fingerprint of the certificate.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="issuedBy")
    def issued_by(self) -> pulumi.Output[str]:
        """
        Name of the certificate authority that issued the certificate.
        """
        return pulumi.get(self, "issued_by")

    @property
    @pulumi.getter(name="issuedOn")
    def issued_on(self) -> pulumi.Output[str]:
        """
        The time & date when the certificate is valid from.
        """
        return pulumi.get(self, "issued_on")

    @property
    @pulumi.getter(name="issuedTo")
    def issued_to(self) -> pulumi.Output[str]:
        """
        Name of whom the certificate has been issued to.
        """
        return pulumi.get(self, "issued_to")

    @property
    @pulumi.getter(name="validUntil")
    def valid_until(self) -> pulumi.Output[str]:
        """
        The time & date when the certificate expires.
        """
        return pulumi.get(self, "valid_until")


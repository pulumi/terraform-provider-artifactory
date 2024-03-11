# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['KeypairArgs', 'Keypair']

@pulumi.input_type
class KeypairArgs:
    def __init__(__self__, *,
                 alias: pulumi.Input[str],
                 pair_name: pulumi.Input[str],
                 pair_type: pulumi.Input[str],
                 private_key: pulumi.Input[str],
                 public_key: pulumi.Input[str],
                 passphrase: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Keypair resource.
        :param pulumi.Input[str] alias: Will be used as a filename when retrieving the public key via REST API.
        :param pulumi.Input[str] pair_name: A unique identifier for the Key Pair record.
        :param pulumi.Input[str] pair_type: Key Pair type. Supported types - GPG and RSA.
        :param pulumi.Input[str] private_key: Private key. PEM format will be validated. Must not include extranous spaces or tabs.
        :param pulumi.Input[str] public_key: Public key. PEM format will be validated. Must not include extranous spaces or tabs.
               
               Artifactory REST API call 'Get Key Pair' doesn't return attributes `private_key` and `passphrase`, but consumes these keys in the POST call.
        :param pulumi.Input[str] passphrase: Passphrase will be used to decrypt the private key. Validated server side.
        """
        pulumi.set(__self__, "alias", alias)
        pulumi.set(__self__, "pair_name", pair_name)
        pulumi.set(__self__, "pair_type", pair_type)
        pulumi.set(__self__, "private_key", private_key)
        pulumi.set(__self__, "public_key", public_key)
        if passphrase is not None:
            pulumi.set(__self__, "passphrase", passphrase)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Input[str]:
        """
        Will be used as a filename when retrieving the public key via REST API.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="pairName")
    def pair_name(self) -> pulumi.Input[str]:
        """
        A unique identifier for the Key Pair record.
        """
        return pulumi.get(self, "pair_name")

    @pair_name.setter
    def pair_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "pair_name", value)

    @property
    @pulumi.getter(name="pairType")
    def pair_type(self) -> pulumi.Input[str]:
        """
        Key Pair type. Supported types - GPG and RSA.
        """
        return pulumi.get(self, "pair_type")

    @pair_type.setter
    def pair_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "pair_type", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Input[str]:
        """
        Private key. PEM format will be validated. Must not include extranous spaces or tabs.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Input[str]:
        """
        Public key. PEM format will be validated. Must not include extranous spaces or tabs.

        Artifactory REST API call 'Get Key Pair' doesn't return attributes `private_key` and `passphrase`, but consumes these keys in the POST call.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def passphrase(self) -> Optional[pulumi.Input[str]]:
        """
        Passphrase will be used to decrypt the private key. Validated server side.
        """
        return pulumi.get(self, "passphrase")

    @passphrase.setter
    def passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passphrase", value)


@pulumi.input_type
class _KeypairState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 pair_name: Optional[pulumi.Input[str]] = None,
                 pair_type: Optional[pulumi.Input[str]] = None,
                 passphrase: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Keypair resources.
        :param pulumi.Input[str] alias: Will be used as a filename when retrieving the public key via REST API.
        :param pulumi.Input[str] pair_name: A unique identifier for the Key Pair record.
        :param pulumi.Input[str] pair_type: Key Pair type. Supported types - GPG and RSA.
        :param pulumi.Input[str] passphrase: Passphrase will be used to decrypt the private key. Validated server side.
        :param pulumi.Input[str] private_key: Private key. PEM format will be validated. Must not include extranous spaces or tabs.
        :param pulumi.Input[str] public_key: Public key. PEM format will be validated. Must not include extranous spaces or tabs.
               
               Artifactory REST API call 'Get Key Pair' doesn't return attributes `private_key` and `passphrase`, but consumes these keys in the POST call.
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if pair_name is not None:
            pulumi.set(__self__, "pair_name", pair_name)
        if pair_type is not None:
            pulumi.set(__self__, "pair_type", pair_type)
        if passphrase is not None:
            pulumi.set(__self__, "passphrase", passphrase)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        Will be used as a filename when retrieving the public key via REST API.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="pairName")
    def pair_name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique identifier for the Key Pair record.
        """
        return pulumi.get(self, "pair_name")

    @pair_name.setter
    def pair_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pair_name", value)

    @property
    @pulumi.getter(name="pairType")
    def pair_type(self) -> Optional[pulumi.Input[str]]:
        """
        Key Pair type. Supported types - GPG and RSA.
        """
        return pulumi.get(self, "pair_type")

    @pair_type.setter
    def pair_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pair_type", value)

    @property
    @pulumi.getter
    def passphrase(self) -> Optional[pulumi.Input[str]]:
        """
        Passphrase will be used to decrypt the private key. Validated server side.
        """
        return pulumi.get(self, "passphrase")

    @passphrase.setter
    def passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passphrase", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        Private key. PEM format will be validated. Must not include extranous spaces or tabs.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        Public key. PEM format will be validated. Must not include extranous spaces or tabs.

        Artifactory REST API call 'Get Key Pair' doesn't return attributes `private_key` and `passphrase`, but consumes these keys in the POST call.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)


class Keypair(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 pair_name: Optional[pulumi.Input[str]] = None,
                 pair_type: Optional[pulumi.Input[str]] = None,
                 passphrase: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        RSA key pairs are used to sign and verify the Alpine Linux index files in JFrog Artifactory, while GPG key pairs are
        used to sign and validate packages integrity in JFrog Distribution. The JFrog Platform enables you to manage multiple RSA and GPG signing keys through the Keys Management UI and REST API. The JFrog Platform supports managing multiple pairs of GPG signing keys to sign packages for authentication of several package types such as Debian, Opkg, and RPM through the Keys Management UI and REST API.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        some_keypair_6543461672124900137 = artifactory.Keypair("some-keypair-6543461672124900137",
            pair_name="some-keypair-6543461672124900137",
            pair_type="RSA",
            alias="some-alias-6543461672124900137",
            private_key=(lambda path: open(path).read())("samples/rsa.priv"),
            public_key=(lambda path: open(path).read())("samples/rsa.pub"),
            passphrase="PASSPHRASE")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Keypair can be imported using the pair name, e.g.

        ```sh
        $ pulumi import artifactory:index/keypair:Keypair my-keypair my-keypair-name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: Will be used as a filename when retrieving the public key via REST API.
        :param pulumi.Input[str] pair_name: A unique identifier for the Key Pair record.
        :param pulumi.Input[str] pair_type: Key Pair type. Supported types - GPG and RSA.
        :param pulumi.Input[str] passphrase: Passphrase will be used to decrypt the private key. Validated server side.
        :param pulumi.Input[str] private_key: Private key. PEM format will be validated. Must not include extranous spaces or tabs.
        :param pulumi.Input[str] public_key: Public key. PEM format will be validated. Must not include extranous spaces or tabs.
               
               Artifactory REST API call 'Get Key Pair' doesn't return attributes `private_key` and `passphrase`, but consumes these keys in the POST call.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeypairArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        RSA key pairs are used to sign and verify the Alpine Linux index files in JFrog Artifactory, while GPG key pairs are
        used to sign and validate packages integrity in JFrog Distribution. The JFrog Platform enables you to manage multiple RSA and GPG signing keys through the Keys Management UI and REST API. The JFrog Platform supports managing multiple pairs of GPG signing keys to sign packages for authentication of several package types such as Debian, Opkg, and RPM through the Keys Management UI and REST API.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        some_keypair_6543461672124900137 = artifactory.Keypair("some-keypair-6543461672124900137",
            pair_name="some-keypair-6543461672124900137",
            pair_type="RSA",
            alias="some-alias-6543461672124900137",
            private_key=(lambda path: open(path).read())("samples/rsa.priv"),
            public_key=(lambda path: open(path).read())("samples/rsa.pub"),
            passphrase="PASSPHRASE")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Keypair can be imported using the pair name, e.g.

        ```sh
        $ pulumi import artifactory:index/keypair:Keypair my-keypair my-keypair-name
        ```

        :param str resource_name: The name of the resource.
        :param KeypairArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeypairArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 pair_name: Optional[pulumi.Input[str]] = None,
                 pair_type: Optional[pulumi.Input[str]] = None,
                 passphrase: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KeypairArgs.__new__(KeypairArgs)

            if alias is None and not opts.urn:
                raise TypeError("Missing required property 'alias'")
            __props__.__dict__["alias"] = alias
            if pair_name is None and not opts.urn:
                raise TypeError("Missing required property 'pair_name'")
            __props__.__dict__["pair_name"] = pair_name
            if pair_type is None and not opts.urn:
                raise TypeError("Missing required property 'pair_type'")
            __props__.__dict__["pair_type"] = pair_type
            __props__.__dict__["passphrase"] = None if passphrase is None else pulumi.Output.secret(passphrase)
            if private_key is None and not opts.urn:
                raise TypeError("Missing required property 'private_key'")
            __props__.__dict__["private_key"] = None if private_key is None else pulumi.Output.secret(private_key)
            if public_key is None and not opts.urn:
                raise TypeError("Missing required property 'public_key'")
            __props__.__dict__["public_key"] = public_key
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["passphrase", "privateKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Keypair, __self__).__init__(
            'artifactory:index/keypair:Keypair',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[str]] = None,
            pair_name: Optional[pulumi.Input[str]] = None,
            pair_type: Optional[pulumi.Input[str]] = None,
            passphrase: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            public_key: Optional[pulumi.Input[str]] = None) -> 'Keypair':
        """
        Get an existing Keypair resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: Will be used as a filename when retrieving the public key via REST API.
        :param pulumi.Input[str] pair_name: A unique identifier for the Key Pair record.
        :param pulumi.Input[str] pair_type: Key Pair type. Supported types - GPG and RSA.
        :param pulumi.Input[str] passphrase: Passphrase will be used to decrypt the private key. Validated server side.
        :param pulumi.Input[str] private_key: Private key. PEM format will be validated. Must not include extranous spaces or tabs.
        :param pulumi.Input[str] public_key: Public key. PEM format will be validated. Must not include extranous spaces or tabs.
               
               Artifactory REST API call 'Get Key Pair' doesn't return attributes `private_key` and `passphrase`, but consumes these keys in the POST call.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KeypairState.__new__(_KeypairState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["pair_name"] = pair_name
        __props__.__dict__["pair_type"] = pair_type
        __props__.__dict__["passphrase"] = passphrase
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["public_key"] = public_key
        return Keypair(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        """
        Will be used as a filename when retrieving the public key via REST API.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter(name="pairName")
    def pair_name(self) -> pulumi.Output[str]:
        """
        A unique identifier for the Key Pair record.
        """
        return pulumi.get(self, "pair_name")

    @property
    @pulumi.getter(name="pairType")
    def pair_type(self) -> pulumi.Output[str]:
        """
        Key Pair type. Supported types - GPG and RSA.
        """
        return pulumi.get(self, "pair_type")

    @property
    @pulumi.getter
    def passphrase(self) -> pulumi.Output[Optional[str]]:
        """
        Passphrase will be used to decrypt the private key. Validated server side.
        """
        return pulumi.get(self, "passphrase")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        """
        Private key. PEM format will be validated. Must not include extranous spaces or tabs.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[str]:
        """
        Public key. PEM format will be validated. Must not include extranous spaces or tabs.

        Artifactory REST API call 'Get Key Pair' doesn't return attributes `private_key` and `passphrase`, but consumes these keys in the POST call.
        """
        return pulumi.get(self, "public_key")


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetFileResult',
    'AwaitableGetFileResult',
    'get_file',
    'get_file_output',
]

@pulumi.output_type
class GetFileResult:
    """
    A collection of values returned by getFile.
    """
    def __init__(__self__, created=None, created_by=None, download_uri=None, force_overwrite=None, id=None, last_modified=None, last_updated=None, md5=None, mimetype=None, modified_by=None, output_path=None, path=None, path_is_aliased=None, repository=None, sha1=None, sha256=None, size=None):
        if created and not isinstance(created, str):
            raise TypeError("Expected argument 'created' to be a str")
        pulumi.set(__self__, "created", created)
        if created_by and not isinstance(created_by, str):
            raise TypeError("Expected argument 'created_by' to be a str")
        pulumi.set(__self__, "created_by", created_by)
        if download_uri and not isinstance(download_uri, str):
            raise TypeError("Expected argument 'download_uri' to be a str")
        pulumi.set(__self__, "download_uri", download_uri)
        if force_overwrite and not isinstance(force_overwrite, bool):
            raise TypeError("Expected argument 'force_overwrite' to be a bool")
        pulumi.set(__self__, "force_overwrite", force_overwrite)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_modified and not isinstance(last_modified, str):
            raise TypeError("Expected argument 'last_modified' to be a str")
        pulumi.set(__self__, "last_modified", last_modified)
        if last_updated and not isinstance(last_updated, str):
            raise TypeError("Expected argument 'last_updated' to be a str")
        pulumi.set(__self__, "last_updated", last_updated)
        if md5 and not isinstance(md5, str):
            raise TypeError("Expected argument 'md5' to be a str")
        pulumi.set(__self__, "md5", md5)
        if mimetype and not isinstance(mimetype, str):
            raise TypeError("Expected argument 'mimetype' to be a str")
        pulumi.set(__self__, "mimetype", mimetype)
        if modified_by and not isinstance(modified_by, str):
            raise TypeError("Expected argument 'modified_by' to be a str")
        pulumi.set(__self__, "modified_by", modified_by)
        if output_path and not isinstance(output_path, str):
            raise TypeError("Expected argument 'output_path' to be a str")
        pulumi.set(__self__, "output_path", output_path)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if path_is_aliased and not isinstance(path_is_aliased, bool):
            raise TypeError("Expected argument 'path_is_aliased' to be a bool")
        pulumi.set(__self__, "path_is_aliased", path_is_aliased)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)
        if sha1 and not isinstance(sha1, str):
            raise TypeError("Expected argument 'sha1' to be a str")
        pulumi.set(__self__, "sha1", sha1)
        if sha256 and not isinstance(sha256, str):
            raise TypeError("Expected argument 'sha256' to be a str")
        pulumi.set(__self__, "sha256", sha256)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)

    @property
    @pulumi.getter
    def created(self) -> str:
        """
        The time & date when the file was created.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> str:
        """
        The user who created the file.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="downloadUri")
    def download_uri(self) -> str:
        """
        The URI that can be used to download the file.
        """
        return pulumi.get(self, "download_uri")

    @property
    @pulumi.getter(name="forceOverwrite")
    def force_overwrite(self) -> Optional[bool]:
        return pulumi.get(self, "force_overwrite")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> str:
        """
        The time & date when the file was last modified.
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> str:
        """
        The time & date when the file was last updated.
        """
        return pulumi.get(self, "last_updated")

    @property
    @pulumi.getter
    def md5(self) -> str:
        """
        MD5 checksum of the file.
        """
        return pulumi.get(self, "md5")

    @property
    @pulumi.getter
    def mimetype(self) -> str:
        """
        The mimetype of the file.
        """
        return pulumi.get(self, "mimetype")

    @property
    @pulumi.getter(name="modifiedBy")
    def modified_by(self) -> str:
        """
        The user who last modified the file.
        """
        return pulumi.get(self, "modified_by")

    @property
    @pulumi.getter(name="outputPath")
    def output_path(self) -> str:
        return pulumi.get(self, "output_path")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="pathIsAliased")
    def path_is_aliased(self) -> Optional[bool]:
        return pulumi.get(self, "path_is_aliased")

    @property
    @pulumi.getter
    def repository(self) -> str:
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def sha1(self) -> str:
        """
        SHA1 checksum of the file.
        """
        return pulumi.get(self, "sha1")

    @property
    @pulumi.getter
    def sha256(self) -> str:
        """
        SHA256 checksum of the file.
        """
        return pulumi.get(self, "sha256")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The size of the file.
        """
        return pulumi.get(self, "size")


class AwaitableGetFileResult(GetFileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileResult(
            created=self.created,
            created_by=self.created_by,
            download_uri=self.download_uri,
            force_overwrite=self.force_overwrite,
            id=self.id,
            last_modified=self.last_modified,
            last_updated=self.last_updated,
            md5=self.md5,
            mimetype=self.mimetype,
            modified_by=self.modified_by,
            output_path=self.output_path,
            path=self.path,
            path_is_aliased=self.path_is_aliased,
            repository=self.repository,
            sha1=self.sha1,
            sha256=self.sha256,
            size=self.size)


def get_file(force_overwrite: Optional[bool] = None,
             output_path: Optional[str] = None,
             path: Optional[str] = None,
             path_is_aliased: Optional[bool] = None,
             repository: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFileResult:
    """
    ## # Artifactory File Data Source

    Provides an Artifactory file datasource. This can be used to download a file from a given Artifactory repository.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    my_file = artifactory.get_file(repository="repo-key",
        path="/path/to/the/artifact.zip",
        output_path="tmp/artifact.zip")
    ```
    <!--End PulumiCodeChooser -->


    :param bool force_overwrite: If set to true, an existing file in the output_path will be overwritten. Default: `false`
    :param str output_path: The local path the file should be downloaded to.
    :param str path: The path to the file within the repository.
    :param bool path_is_aliased: If set to `true`, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory if the file exists. When using a smart remote repository, it is recommended to set this attribute to `true`. This is necessary to ensure that the provider fetches the artifact directly from Artifactory. If this attribute is not set or is set to `false`, there is a risk of fetching the `-cache` directory in Artifactory, potentially resulting in resource expiration and a 404 error.
    :param str repository: Name of the repository where the file is stored.
    """
    __args__ = dict()
    __args__['forceOverwrite'] = force_overwrite
    __args__['outputPath'] = output_path
    __args__['path'] = path
    __args__['pathIsAliased'] = path_is_aliased
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('artifactory:index/getFile:getFile', __args__, opts=opts, typ=GetFileResult).value

    return AwaitableGetFileResult(
        created=pulumi.get(__ret__, 'created'),
        created_by=pulumi.get(__ret__, 'created_by'),
        download_uri=pulumi.get(__ret__, 'download_uri'),
        force_overwrite=pulumi.get(__ret__, 'force_overwrite'),
        id=pulumi.get(__ret__, 'id'),
        last_modified=pulumi.get(__ret__, 'last_modified'),
        last_updated=pulumi.get(__ret__, 'last_updated'),
        md5=pulumi.get(__ret__, 'md5'),
        mimetype=pulumi.get(__ret__, 'mimetype'),
        modified_by=pulumi.get(__ret__, 'modified_by'),
        output_path=pulumi.get(__ret__, 'output_path'),
        path=pulumi.get(__ret__, 'path'),
        path_is_aliased=pulumi.get(__ret__, 'path_is_aliased'),
        repository=pulumi.get(__ret__, 'repository'),
        sha1=pulumi.get(__ret__, 'sha1'),
        sha256=pulumi.get(__ret__, 'sha256'),
        size=pulumi.get(__ret__, 'size'))


@_utilities.lift_output_func(get_file)
def get_file_output(force_overwrite: Optional[pulumi.Input[Optional[bool]]] = None,
                    output_path: Optional[pulumi.Input[str]] = None,
                    path: Optional[pulumi.Input[str]] = None,
                    path_is_aliased: Optional[pulumi.Input[Optional[bool]]] = None,
                    repository: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFileResult]:
    """
    ## # Artifactory File Data Source

    Provides an Artifactory file datasource. This can be used to download a file from a given Artifactory repository.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    my_file = artifactory.get_file(repository="repo-key",
        path="/path/to/the/artifact.zip",
        output_path="tmp/artifact.zip")
    ```
    <!--End PulumiCodeChooser -->


    :param bool force_overwrite: If set to true, an existing file in the output_path will be overwritten. Default: `false`
    :param str output_path: The local path the file should be downloaded to.
    :param str path: The path to the file within the repository.
    :param bool path_is_aliased: If set to `true`, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory if the file exists. When using a smart remote repository, it is recommended to set this attribute to `true`. This is necessary to ensure that the provider fetches the artifact directly from Artifactory. If this attribute is not set or is set to `false`, there is a risk of fetching the `-cache` directory in Artifactory, potentially resulting in resource expiration and a 404 error.
    :param str repository: Name of the repository where the file is stored.
    """
    ...

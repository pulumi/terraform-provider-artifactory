# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetFileinfoResult',
    'AwaitableGetFileinfoResult',
    'get_fileinfo',
    'get_fileinfo_output',
]

@pulumi.output_type
class GetFileinfoResult:
    """
    A collection of values returned by getFileinfo.
    """
    def __init__(__self__, created=None, created_by=None, download_uri=None, id=None, last_modified=None, last_updated=None, md5=None, mimetype=None, modified_by=None, path=None, repository=None, sha1=None, sha256=None, size=None):
        if created and not isinstance(created, str):
            raise TypeError("Expected argument 'created' to be a str")
        pulumi.set(__self__, "created", created)
        if created_by and not isinstance(created_by, str):
            raise TypeError("Expected argument 'created_by' to be a str")
        pulumi.set(__self__, "created_by", created_by)
        if download_uri and not isinstance(download_uri, str):
            raise TypeError("Expected argument 'download_uri' to be a str")
        pulumi.set(__self__, "download_uri", download_uri)
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
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
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
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

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


class AwaitableGetFileinfoResult(GetFileinfoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileinfoResult(
            created=self.created,
            created_by=self.created_by,
            download_uri=self.download_uri,
            id=self.id,
            last_modified=self.last_modified,
            last_updated=self.last_updated,
            md5=self.md5,
            mimetype=self.mimetype,
            modified_by=self.modified_by,
            path=self.path,
            repository=self.repository,
            sha1=self.sha1,
            sha256=self.sha256,
            size=self.size)


def get_fileinfo(path: Optional[str] = None,
                 repository: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFileinfoResult:
    """
    ## # Artifactory File Info Data Source

    Provides an Artifactory fileinfo datasource. This can be used to read metadata of files stored in Artifactory repositories.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    my_file = artifactory.get_fileinfo(path="/path/to/the/artifact.zip",
        repository="repo-key")
    ```


    :param str path: The path to the file within the repository.
    :param str repository: Name of the repository where the file is stored.
    """
    __args__ = dict()
    __args__['path'] = path
    __args__['repository'] = repository
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('artifactory:index/getFileinfo:getFileinfo', __args__, opts=opts, typ=GetFileinfoResult).value

    return AwaitableGetFileinfoResult(
        created=__ret__.created,
        created_by=__ret__.created_by,
        download_uri=__ret__.download_uri,
        id=__ret__.id,
        last_modified=__ret__.last_modified,
        last_updated=__ret__.last_updated,
        md5=__ret__.md5,
        mimetype=__ret__.mimetype,
        modified_by=__ret__.modified_by,
        path=__ret__.path,
        repository=__ret__.repository,
        sha1=__ret__.sha1,
        sha256=__ret__.sha256,
        size=__ret__.size)


@_utilities.lift_output_func(get_fileinfo)
def get_fileinfo_output(path: Optional[pulumi.Input[str]] = None,
                        repository: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFileinfoResult]:
    """
    ## # Artifactory File Info Data Source

    Provides an Artifactory fileinfo datasource. This can be used to read metadata of files stored in Artifactory repositories.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    my_file = artifactory.get_fileinfo(path="/path/to/the/artifact.zip",
        repository="repo-key")
    ```


    :param str path: The path to the file within the repository.
    :param str repository: Name of the repository where the file is stored.
    """
    ...

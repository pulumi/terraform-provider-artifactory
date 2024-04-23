# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetFileListResult',
    'AwaitableGetFileListResult',
    'get_file_list',
    'get_file_list_output',
]

@pulumi.output_type
class GetFileListResult:
    """
    A collection of values returned by getFileList.
    """
    def __init__(__self__, created=None, deep_listing=None, depth=None, files=None, folder_path=None, id=None, include_root_path=None, list_folders=None, metadata_timestamps=None, repository_key=None, uri=None):
        if created and not isinstance(created, str):
            raise TypeError("Expected argument 'created' to be a str")
        pulumi.set(__self__, "created", created)
        if deep_listing and not isinstance(deep_listing, bool):
            raise TypeError("Expected argument 'deep_listing' to be a bool")
        pulumi.set(__self__, "deep_listing", deep_listing)
        if depth and not isinstance(depth, int):
            raise TypeError("Expected argument 'depth' to be a int")
        pulumi.set(__self__, "depth", depth)
        if files and not isinstance(files, list):
            raise TypeError("Expected argument 'files' to be a list")
        pulumi.set(__self__, "files", files)
        if folder_path and not isinstance(folder_path, str):
            raise TypeError("Expected argument 'folder_path' to be a str")
        pulumi.set(__self__, "folder_path", folder_path)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_root_path and not isinstance(include_root_path, bool):
            raise TypeError("Expected argument 'include_root_path' to be a bool")
        pulumi.set(__self__, "include_root_path", include_root_path)
        if list_folders and not isinstance(list_folders, bool):
            raise TypeError("Expected argument 'list_folders' to be a bool")
        pulumi.set(__self__, "list_folders", list_folders)
        if metadata_timestamps and not isinstance(metadata_timestamps, bool):
            raise TypeError("Expected argument 'metadata_timestamps' to be a bool")
        pulumi.set(__self__, "metadata_timestamps", metadata_timestamps)
        if repository_key and not isinstance(repository_key, str):
            raise TypeError("Expected argument 'repository_key' to be a str")
        pulumi.set(__self__, "repository_key", repository_key)
        if uri and not isinstance(uri, str):
            raise TypeError("Expected argument 'uri' to be a str")
        pulumi.set(__self__, "uri", uri)

    @property
    @pulumi.getter
    def created(self) -> str:
        """
        Creation time
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="deepListing")
    def deep_listing(self) -> Optional[bool]:
        """
        Get deep listing
        """
        return pulumi.get(self, "deep_listing")

    @property
    @pulumi.getter
    def depth(self) -> Optional[int]:
        """
        Depth of the deep listing
        """
        return pulumi.get(self, "depth")

    @property
    @pulumi.getter
    def files(self) -> Sequence['outputs.GetFileListFileResult']:
        """
        A list of files.
        """
        return pulumi.get(self, "files")

    @property
    @pulumi.getter(name="folderPath")
    def folder_path(self) -> str:
        """
        Path of the folder
        """
        return pulumi.get(self, "folder_path")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includeRootPath")
    def include_root_path(self) -> Optional[bool]:
        """
        Include root path
        """
        return pulumi.get(self, "include_root_path")

    @property
    @pulumi.getter(name="listFolders")
    def list_folders(self) -> Optional[bool]:
        """
        Include folders
        """
        return pulumi.get(self, "list_folders")

    @property
    @pulumi.getter(name="metadataTimestamps")
    def metadata_timestamps(self) -> Optional[bool]:
        """
        Include metadata timestamps
        """
        return pulumi.get(self, "metadata_timestamps")

    @property
    @pulumi.getter(name="repositoryKey")
    def repository_key(self) -> str:
        """
        Repository key
        """
        return pulumi.get(self, "repository_key")

    @property
    @pulumi.getter
    def uri(self) -> str:
        """
        URL to file/path
        """
        return pulumi.get(self, "uri")


class AwaitableGetFileListResult(GetFileListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileListResult(
            created=self.created,
            deep_listing=self.deep_listing,
            depth=self.depth,
            files=self.files,
            folder_path=self.folder_path,
            id=self.id,
            include_root_path=self.include_root_path,
            list_folders=self.list_folders,
            metadata_timestamps=self.metadata_timestamps,
            repository_key=self.repository_key,
            uri=self.uri)


def get_file_list(deep_listing: Optional[bool] = None,
                  depth: Optional[int] = None,
                  folder_path: Optional[str] = None,
                  include_root_path: Optional[bool] = None,
                  list_folders: Optional[bool] = None,
                  metadata_timestamps: Optional[bool] = None,
                  repository_key: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFileListResult:
    """
    Get a flat (the default) or deep listing of the files and folders (not included by default) within a folder. For deep listing you can specify an optional depth to limit the results. Optionally include a map of metadata timestamp values as part of the result.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    my_repo_file_list = artifactory.get_file_list(repository_key="my-generic-local",
        folder_path="path/to/artifact")
    ```


    :param bool deep_listing: Get deep listing
    :param int depth: Depth of the deep listing
    :param str folder_path: Path of the folder
    :param bool include_root_path: Include root path
    :param bool list_folders: Include folders
    :param bool metadata_timestamps: Include metadata timestamps
    :param str repository_key: Repository key
    """
    __args__ = dict()
    __args__['deepListing'] = deep_listing
    __args__['depth'] = depth
    __args__['folderPath'] = folder_path
    __args__['includeRootPath'] = include_root_path
    __args__['listFolders'] = list_folders
    __args__['metadataTimestamps'] = metadata_timestamps
    __args__['repositoryKey'] = repository_key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('artifactory:index/getFileList:getFileList', __args__, opts=opts, typ=GetFileListResult).value

    return AwaitableGetFileListResult(
        created=pulumi.get(__ret__, 'created'),
        deep_listing=pulumi.get(__ret__, 'deep_listing'),
        depth=pulumi.get(__ret__, 'depth'),
        files=pulumi.get(__ret__, 'files'),
        folder_path=pulumi.get(__ret__, 'folder_path'),
        id=pulumi.get(__ret__, 'id'),
        include_root_path=pulumi.get(__ret__, 'include_root_path'),
        list_folders=pulumi.get(__ret__, 'list_folders'),
        metadata_timestamps=pulumi.get(__ret__, 'metadata_timestamps'),
        repository_key=pulumi.get(__ret__, 'repository_key'),
        uri=pulumi.get(__ret__, 'uri'))


@_utilities.lift_output_func(get_file_list)
def get_file_list_output(deep_listing: Optional[pulumi.Input[Optional[bool]]] = None,
                         depth: Optional[pulumi.Input[Optional[int]]] = None,
                         folder_path: Optional[pulumi.Input[str]] = None,
                         include_root_path: Optional[pulumi.Input[Optional[bool]]] = None,
                         list_folders: Optional[pulumi.Input[Optional[bool]]] = None,
                         metadata_timestamps: Optional[pulumi.Input[Optional[bool]]] = None,
                         repository_key: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFileListResult]:
    """
    Get a flat (the default) or deep listing of the files and folders (not included by default) within a folder. For deep listing you can specify an optional depth to limit the results. Optionally include a map of metadata timestamp values as part of the result.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    my_repo_file_list = artifactory.get_file_list(repository_key="my-generic-local",
        folder_path="path/to/artifact")
    ```


    :param bool deep_listing: Get deep listing
    :param int depth: Depth of the deep listing
    :param str folder_path: Path of the folder
    :param bool include_root_path: Include root path
    :param bool list_folders: Include folders
    :param bool metadata_timestamps: Include metadata timestamps
    :param str repository_key: Repository key
    """
    ...

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetFileinfo
    {
        public static Task<GetFileinfoResult> InvokeAsync(GetFileinfoArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFileinfoResult>("artifactory:index/getFileinfo:getFileinfo", args ?? new GetFileinfoArgs(), options.WithDefaults());

        public static Output<GetFileinfoResult> Invoke(GetFileinfoInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFileinfoResult>("artifactory:index/getFileinfo:getFileinfo", args ?? new GetFileinfoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFileinfoArgs : Pulumi.InvokeArgs
    {
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetFileinfoArgs()
        {
        }
    }

    public sealed class GetFileinfoInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetFileinfoInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFileinfoResult
    {
        public readonly string Created;
        public readonly string CreatedBy;
        public readonly string DownloadUri;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LastModified;
        public readonly string LastUpdated;
        public readonly string Md5;
        public readonly string Mimetype;
        public readonly string ModifiedBy;
        public readonly string Path;
        public readonly string Repository;
        public readonly string Sha1;
        public readonly string Sha256;
        public readonly int Size;

        [OutputConstructor]
        private GetFileinfoResult(
            string created,

            string createdBy,

            string downloadUri,

            string id,

            string lastModified,

            string lastUpdated,

            string md5,

            string mimetype,

            string modifiedBy,

            string path,

            string repository,

            string sha1,

            string sha256,

            int size)
        {
            Created = created;
            CreatedBy = createdBy;
            DownloadUri = downloadUri;
            Id = id;
            LastModified = lastModified;
            LastUpdated = lastUpdated;
            Md5 = md5;
            Mimetype = mimetype;
            ModifiedBy = modifiedBy;
            Path = path;
            Repository = repository;
            Sha1 = sha1;
            Sha256 = sha256;
            Size = size;
        }
    }
}

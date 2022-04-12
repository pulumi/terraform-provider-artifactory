// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetFile
    {
        public static Task<GetFileResult> InvokeAsync(GetFileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFileResult>("artifactory:index/getFile:getFile", args ?? new GetFileArgs(), options.WithDefaults());

        public static Output<GetFileResult> Invoke(GetFileInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFileResult>("artifactory:index/getFile:getFile", args ?? new GetFileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFileArgs : Pulumi.InvokeArgs
    {
        [Input("forceOverwrite")]
        public bool? ForceOverwrite { get; set; }

        [Input("outputPath", required: true)]
        public string OutputPath { get; set; } = null!;

        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        [Input("pathIsAliased")]
        public bool? PathIsAliased { get; set; }

        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetFileArgs()
        {
        }
    }

    public sealed class GetFileInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("forceOverwrite")]
        public Input<bool>? ForceOverwrite { get; set; }

        [Input("outputPath", required: true)]
        public Input<string> OutputPath { get; set; } = null!;

        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("pathIsAliased")]
        public Input<bool>? PathIsAliased { get; set; }

        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetFileInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFileResult
    {
        public readonly string Created;
        public readonly string CreatedBy;
        public readonly string DownloadUri;
        public readonly bool? ForceOverwrite;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LastModified;
        public readonly string LastUpdated;
        public readonly string Md5;
        public readonly string Mimetype;
        public readonly string ModifiedBy;
        public readonly string OutputPath;
        public readonly string Path;
        public readonly bool? PathIsAliased;
        public readonly string Repository;
        public readonly string Sha1;
        public readonly string Sha256;
        public readonly int Size;

        [OutputConstructor]
        private GetFileResult(
            string created,

            string createdBy,

            string downloadUri,

            bool? forceOverwrite,

            string id,

            string lastModified,

            string lastUpdated,

            string md5,

            string mimetype,

            string modifiedBy,

            string outputPath,

            string path,

            bool? pathIsAliased,

            string repository,

            string sha1,

            string sha256,

            int size)
        {
            Created = created;
            CreatedBy = createdBy;
            DownloadUri = downloadUri;
            ForceOverwrite = forceOverwrite;
            Id = id;
            LastModified = lastModified;
            LastUpdated = lastUpdated;
            Md5 = md5;
            Mimetype = mimetype;
            ModifiedBy = modifiedBy;
            OutputPath = outputPath;
            Path = path;
            PathIsAliased = pathIsAliased;
            Repository = repository;
            Sha1 = sha1;
            Sha256 = sha256;
            Size = size;
        }
    }
}

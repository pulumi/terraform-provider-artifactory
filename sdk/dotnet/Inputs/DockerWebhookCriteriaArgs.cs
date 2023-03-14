// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class DockerWebhookCriteriaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Trigger on any local repo.
        /// </summary>
        [Input("anyLocal", required: true)]
        public Input<bool> AnyLocal { get; set; } = null!;

        /// <summary>
        /// Trigger on any remote repo.
        /// </summary>
        [Input("anyRemote", required: true)]
        public Input<bool> AnyRemote { get; set; } = null!;

        [Input("excludePatterns")]
        private InputList<string>? _excludePatterns;

        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
        /// </summary>
        public InputList<string> ExcludePatterns
        {
            get => _excludePatterns ?? (_excludePatterns = new InputList<string>());
            set => _excludePatterns = value;
        }

        [Input("includePatterns")]
        private InputList<string>? _includePatterns;

        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
        /// </summary>
        public InputList<string> IncludePatterns
        {
            get => _includePatterns ?? (_includePatterns = new InputList<string>());
            set => _includePatterns = value;
        }

        [Input("repoKeys", required: true)]
        private InputList<string>? _repoKeys;

        /// <summary>
        /// Trigger on this list of repo keys.
        /// </summary>
        public InputList<string> RepoKeys
        {
            get => _repoKeys ?? (_repoKeys = new InputList<string>());
            set => _repoKeys = value;
        }

        public DockerWebhookCriteriaArgs()
        {
        }
        public static new DockerWebhookCriteriaArgs Empty => new DockerWebhookCriteriaArgs();
    }
}

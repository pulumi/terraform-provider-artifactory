// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Artifactory
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly Pulumi.Config __config = new Pulumi.Config("artifactory");

        private static readonly __Value<string?> _accessToken = new __Value<string?>(() => __config.Get("accessToken"));
        /// <summary>
        /// This is a access token that can be given to you by your admin under `Identity and Access`
        /// </summary>
        public static string? AccessToken
        {
            get => _accessToken.Get();
            set => _accessToken.Set(value);
        }

        private static readonly __Value<string?> _apiKey = new __Value<string?>(() => __config.Get("apiKey"));
        /// <summary>
        /// API token. Projects functionality will not work with any auth method other than access tokens
        /// </summary>
        public static string? ApiKey
        {
            get => _apiKey.Get();
            set => _apiKey.Set(value);
        }

        private static readonly __Value<bool?> _checkLicense = new __Value<bool?>(() => __config.GetBoolean("checkLicense") ?? false);
        /// <summary>
        /// Toggle for pre-flight checking of Artifactory Pro and Enterprise license. Default to `true`.
        /// </summary>
        public static bool? CheckLicense
        {
            get => _checkLicense.Get();
            set => _checkLicense.Set(value);
        }

        private static readonly __Value<string?> _url = new __Value<string?>(() => __config.Get("url"));
        public static string? Url
        {
            get => _url.Get();
            set => _url.Set(value);
        }

    }
}

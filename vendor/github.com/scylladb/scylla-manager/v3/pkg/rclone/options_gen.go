// Code generated by generate_options.go. DO NOT EDIT.

package rclone

// AzureOptions is a clone rclone file system Options designed for inclusion
// in Scylla Manager Agent config, and YAML parsing.
type AzureOptions struct {
	// Storage Account Name (leave blank to use SAS URL or Emulator)
	Account string `yaml:"account"`
	// Path to file containing credentials for use with a service principal.
	//
	// Leave blank normally. Needed only if you want to use a service principal instead of interactive login.
	//
	//	$ az sp create-for-rbac --name "<name>" \
	//	  --role "Storage Blob Data Owner" \
	//	  --scopes "/subscriptions/<subscription>/resourceGroups/<resource-group>/providers/Microsoft.Storage/storageAccounts/<storage-account>/blobServices/default/containers/<container>" \
	//	  > azure-principal.json
	//
	// See [Use Azure CLI to assign an Azure role for access to blob and queue data](https://docs.microsoft.com/en-us/azure/storage/common/storage-auth-aad-rbac-cli)
	// for more details.
	ServicePrincipalFile string `yaml:"service_principal_file"`
	// Storage Account Key (leave blank to use SAS URL or Emulator)
	Key string `yaml:"key"`
	// SAS URL for container level access only
	// (leave blank if using account/key or Emulator)
	SasUrl string `yaml:"sas_url"`
	// Use a managed service identity to authenticate (only works in Azure)
	//
	// When true, use a [managed service identity](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/)
	// to authenticate to Azure Storage instead of a SAS token or account key.
	//
	// If the VM(SS) on which this program is running has a system-assigned identity, it will
	// be used by default. If the resource has no system-assigned but exactly one user-assigned identity,
	// the user-assigned identity will be used by default. If the resource has multiple user-assigned
	// identities, the identity to use must be explicitly specified using exactly one of the msi_object_id,
	// msi_client_id, or msi_mi_res_id parameters.
	UseMsi string `yaml:"use_msi"`
	// Object ID of the user-assigned MSI to use, if any. Leave blank if msi_client_id or msi_mi_res_id specified.
	MsiObjectID string `yaml:"msi_object_id"`
	// Object ID of the user-assigned MSI to use, if any. Leave blank if msi_object_id or msi_mi_res_id specified.
	MsiClientID string `yaml:"msi_client_id"`
	// Azure resource ID of the user-assigned MSI to use, if any. Leave blank if msi_client_id or msi_object_id specified.
	MsiMiResID string `yaml:"msi_mi_res_id"`
	// Uses local storage emulator if provided as 'true' (leave blank if using real azure storage endpoint)
	UseEmulator string `yaml:"use_emulator"`
	// Endpoint for the service
	// Leave blank normally.
	Endpoint string `yaml:"endpoint"`
	// Cutoff for switching to chunked upload (<= 256MB). (Deprecated)
	UploadCutoff string `yaml:"upload_cutoff"`
	// Upload chunk size (<= 100MB).
	//
	// Note that this is stored in memory and there may be up to
	// "--transfers" chunks stored at once in memory.
	ChunkSize string `yaml:"chunk_size"`
	// Size of blob list.
	//
	// This sets the number of blobs requested in each listing chunk. Default
	// is the maximum, 5000. "List blobs" requests are permitted 2 minutes
	// per megabyte to complete. If an operation is taking longer than 2
	// minutes per megabyte on average, it will time out (
	// [source](https://docs.microsoft.com/en-us/rest/api/storageservices/setting-timeouts-for-blob-service-operations#exceptions-to-default-timeout-interval)
	// ). This can be used to limit the number of blobs items to return, to
	// avoid the time out.
	ListChunk string `yaml:"list_chunk"`
	// Access tier of blob: hot, cool or archive.
	//
	// Archived blobs can be restored by setting access tier to hot or
	// cool. Leave blank if you intend to use default access tier, which is
	// set at account level
	//
	// If there is no "access tier" specified, rclone doesn't apply any tier.
	// rclone performs "Set Tier" operation on blobs while uploading, if objects
	// are not modified, specifying "access tier" to new one will have no effect.
	// If blobs are in "archive tier" at remote, trying to perform data transfer
	// operations from remote will not be allowed. User should first restore by
	// tiering blob to "Hot" or "Cool".
	AccessTier string `yaml:"access_tier"`
	// Delete archive tier blobs before overwriting.
	//
	// Archive tier blobs cannot be updated. So without this flag, if you
	// attempt to update an archive tier blob, then rclone will produce the
	// error:
	//
	//	can't update archive tier blob without --azureblob-archive-tier-delete
	//
	// With this flag set then before rclone attempts to overwrite an archive
	// tier blob, it will delete the existing blob before uploading its
	// replacement.  This has the potential for data loss if the upload fails
	// (unlike updating a normal blob) and also may cost more since deleting
	// archive tier blobs early may be chargable.
	ArchiveTierDelete string `yaml:"archive_tier_delete"`
	// Don't store MD5 checksum with object metadata.
	//
	// Normally rclone will calculate the MD5 checksum of the input before
	// uploading it so it can add it to metadata on the object. This is great
	// for data integrity checking but can cause long delays for large files
	// to start uploading.
	DisableChecksum string `yaml:"disable_checksum"`
	// How often internal memory buffer pools will be flushed.
	// Uploads which requires additional buffers (f.e multipart) will use memory pool for allocations.
	// This option controls how often unused buffers will be removed from the pool.
	MemoryPoolFlushTime string `yaml:"memory_pool_flush_time"`
	// Whether to use mmap buffers in internal memory pool.
	MemoryPoolUseMmap string `yaml:"memory_pool_use_mmap"`
	// This sets the encoding for the backend.
	//
	// See: the [encoding section in the overview](/overview/#encoding) for more info.
	Encoding string `yaml:"encoding"`
}

// GCSOptions is a clone rclone file system Options designed for inclusion
// in Scylla Manager Agent config, and YAML parsing.
type GCSOptions struct {
	// OAuth Client Id
	// Leave blank normally.
	ClientID string `yaml:"client_id"`
	// OAuth Client Secret
	// Leave blank normally.
	ClientSecret string `yaml:"client_secret"`
	// OAuth Access Token as a JSON blob.
	Token string `yaml:"token"`
	// Auth server URL.
	// Leave blank to use the provider defaults.
	AuthUrl string `yaml:"auth_url"`
	// Token server url.
	// Leave blank to use the provider defaults.
	TokenUrl string `yaml:"token_url"`
	// Project number.
	// Optional - needed only for list/create/delete buckets - see your developer console.
	ProjectNumber string `yaml:"project_number"`
	// Service Account Credentials JSON file path
	// Leave blank normally.
	// Needed only if you want use SA instead of interactive login.
	//
	// Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.
	ServiceAccountFile string `yaml:"service_account_file"`
	// Service Account Credentials JSON blob
	// Leave blank normally.
	// Needed only if you want use SA instead of interactive login.
	ServiceAccountCredentials string `yaml:"service_account_credentials"`
	// Access public buckets and objects without credentials
	// Set to 'true' if you just want to download files and don't configure credentials.
	Anonymous string `yaml:"anonymous"`
	// Access Control List for new objects.
	ObjectAcl string `yaml:"object_acl"`
	// Access Control List for new buckets.
	BucketAcl string `yaml:"bucket_acl"`
	// Access checks should use bucket-level IAM policies.
	//
	// If you want to upload objects to a bucket with Bucket Policy Only set
	// then you will need to set this.
	//
	// When it is set, rclone:
	//
	// - ignores ACLs set on buckets
	// - ignores ACLs set on objects
	// - creates buckets with Bucket Policy Only set
	//
	// Docs: https://cloud.google.com/storage/docs/bucket-policy-only
	BucketPolicyOnly string `yaml:"bucket_policy_only"`
	// Location for the newly created buckets.
	Location string `yaml:"location"`
	// The storage class to use when storing objects in Google Cloud Storage.
	StorageClass string `yaml:"storage_class"`
	// This sets the encoding for the backend.
	//
	// See: the [encoding section in the overview](/overview/#encoding) for more info.
	Encoding string `yaml:"encoding"`
	// How often internal memory buffer pools will be flushed.
	// Uploads which requires additional buffers (f.e multipart) will use memory pool for allocations.
	// This option controls how often unused buffers will be removed from the pool.
	MemoryPoolFlushTime string `yaml:"memory_pool_flush_time"`
	// Whether to use mmap buffers in internal memory pool.
	MemoryPoolUseMmap string `yaml:"memory_pool_use_mmap"`
	// Chunk size to use for uploading.
	//
	// When uploading large files or files with unknown
	// size (eg from "rclone rcat" or uploaded with "rclone mount" or google
	// photos or google docs) they will be uploaded as multi chunk uploads
	// using this chunk size.
	//
	// Files which contains fewer than size bytes will be uploaded in a single request.
	// Files which contains size bytes or more will be uploaded in separate chunks.
	// If size is zero, media will be uploaded in a single request.
	ChunkSize string `yaml:"chunk_size"`
	// How many items are returned in one chunk during directory listing
	ListChunk string `yaml:"list_chunk"`
	// Whether to create bucket if it doesn't exists.
	// If bucket doesn't exists, error will be returned.'
	AllowCreateBucket string `yaml:"allow_create_bucket"`
}

// LocalOptions is a clone rclone file system Options designed for inclusion
// in Scylla Manager Agent config, and YAML parsing.
type LocalOptions struct {
	// Disable UNC (long path names) conversion on Windows
	Nounc string `yaml:"nounc"`
	// Follow symlinks and copy the pointed to item.
	CopyLinks string `yaml:"copy_links"`
	// Translate symlinks to/from regular files with a '.rclonelink' extension
	Links string `yaml:"links"`
	// Don't warn about skipped symlinks.
	// This flag disables warning messages on skipped symlinks or junction
	// points, as you explicitly acknowledge that they should be skipped.
	SkipLinks string `yaml:"skip_links"`
	// Assume the Stat size of links is zero (and read them instead)
	//
	// On some virtual filesystems (such ash LucidLink), reading a link size via a Stat call always returns 0.
	// However, on unix it reads as the length of the text in the link. This may cause errors like this when
	// syncing:
	//
	//	Failed to copy: corrupted on transfer: sizes differ 0 vs 13
	//
	// Setting this flag causes rclone to read the link and use that as the size of the link
	// instead of 0 which in most cases fixes the problem.
	ZeroSizeLinks string `yaml:"zero_size_links"`
	// Don't apply unicode normalization to paths and filenames (Deprecated)
	//
	// This flag is deprecated now.  Rclone no longer normalizes unicode file
	// names, but it compares them with unicode normalization in the sync
	// routine instead.
	NoUnicodeNormalization string `yaml:"no_unicode_normalization"`
	// Don't check to see if the files change during upload
	//
	// Normally rclone checks the size and modification time of files as they
	// are being uploaded and aborts with a message which starts "can't copy
	// - source file is being updated" if the file changes during upload.
	//
	// However on some file systems this modification time check may fail (e.g.
	// [Glusterfs #2206](https://github.com/rclone/rclone/issues/2206)) so this
	// check can be disabled with this flag.
	//
	// If this flag is set, rclone will use its best efforts to transfer a
	// file which is being updated. If the file is only having things
	// appended to it (e.g. a log) then rclone will transfer the log file with
	// the size it had the first time rclone saw it.
	//
	// If the file is being modified throughout (not just appended to) then
	// the transfer may fail with a hash check failure.
	//
	// In detail, once the file has had stat() called on it for the first
	// time we:
	//
	// - Only transfer the size that stat gave
	// - Only checksum the size that stat gave
	// - Don't update the stat info for the file
	NoCheckUpdated string `yaml:"no_check_updated"`
	// Don't cross filesystem boundaries (unix/macOS only).
	OneFileSystem string `yaml:"one_file_system"`
	// Force the filesystem to report itself as case sensitive.
	//
	// Normally the local backend declares itself as case insensitive on
	// Windows/macOS and case sensitive for everything else.  Use this flag
	// to override the default choice.
	CaseSensitive string `yaml:"case_sensitive"`
	// Force the filesystem to report itself as case insensitive
	//
	// Normally the local backend declares itself as case insensitive on
	// Windows/macOS and case sensitive for everything else.  Use this flag
	// to override the default choice.
	CaseInsensitive string `yaml:"case_insensitive"`
	// Disable sparse files for multi-thread downloads
	//
	// On Windows platforms rclone will make sparse files when doing
	// multi-thread downloads. This avoids long pauses on large files where
	// the OS zeros the file. However sparse files may be undesirable as they
	// cause disk fragmentation and can be slow to work with.
	NoSparse string `yaml:"no_sparse"`
	// Disable setting modtime
	//
	// Normally rclone updates modification time of files after they are done
	// uploading. This can cause permissions issues on Linux platforms when
	// the user rclone is running as does not own the file uploaded, such as
	// when copying to a CIFS mount owned by another user. If this option is
	// enabled, rclone will no longer update the modtime after copying a file.
	NoSetModtime string `yaml:"no_set_modtime"`
	// This sets the encoding for the backend.
	//
	// See: the [encoding section in the overview](/overview/#encoding) for more info.
	Encoding string `yaml:"encoding"`
}

// S3Options is a clone rclone file system Options designed for inclusion
// in Scylla Manager Agent config, and YAML parsing.
type S3Options struct {
	// Choose your S3 provider.
	Provider string `yaml:"provider"`
	// Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	// Only applies if access_key_id and secret_access_key is blank.
	EnvAuth string `yaml:"env_auth"`
	// AWS Access Key ID.
	// Leave blank for anonymous access or runtime credentials.
	AccessKeyID string `yaml:"access_key_id"`
	// AWS Secret Access Key (password)
	// Leave blank for anonymous access or runtime credentials.
	SecretAccessKey string `yaml:"secret_access_key"`
	// Region to connect to.
	Region string `yaml:"region"`
	// Endpoint for S3 API.
	// Leave blank if using AWS to use the default endpoint for the region.
	Endpoint string `yaml:"endpoint"`
	// Location constraint - must be set to match the Region.
	// Used when creating buckets only.
	LocationConstraint string `yaml:"location_constraint"`
	// Canned ACL used when creating buckets and storing or copying objects.
	//
	// This ACL is used for creating objects and if bucket_acl isn't set, for creating buckets too.
	//
	// For more info visit https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
	//
	// Note that this ACL is applied when server-side copying objects as S3
	// doesn't copy the ACL from the source but rather writes a fresh one.
	Acl string `yaml:"acl"`
	// Canned ACL used when creating buckets.
	//
	// For more info visit https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
	//
	// Note that this ACL is applied when only when creating buckets.  If it
	// isn't set then "acl" is used instead.
	BucketAcl string `yaml:"bucket_acl"`
	// Enables requester pays option when interacting with S3 bucket.
	RequesterPays string `yaml:"requester_pays"`
	// The server-side encryption algorithm used when storing this object in S3.
	ServerSideEncryption string `yaml:"server_side_encryption"`
	// If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	SseCustomerAlgorithm string `yaml:"sse_customer_algorithm"`
	// If using KMS ID you must provide the ARN of Key.
	SseKmsKeyID string `yaml:"sse_kms_key_id"`
	// If using SSE-C you must provide the secret encryption key used to encrypt/decrypt your data.
	SseCustomerKey string `yaml:"sse_customer_key"`
	// If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	//
	// If you leave it blank, this is calculated automatically from the sse_customer_key provided.
	SseCustomerKeyMd5 string `yaml:"sse_customer_key_md5"`
	// The storage class to use when storing new objects in S3.
	StorageClass string `yaml:"storage_class"`
	// Cutoff for switching to chunked upload
	//
	// Any files larger than this will be uploaded in chunks of chunk_size.
	// The minimum is 0 and the maximum is 5GB.
	UploadCutoff string `yaml:"upload_cutoff"`
	// Chunk size to use for uploading.
	//
	// When uploading files larger than upload_cutoff or files with unknown
	// size (e.g. from "rclone rcat" or uploaded with "rclone mount" or google
	// photos or google docs) they will be uploaded as multipart uploads
	// using this chunk size.
	//
	// Note that "--s3-upload-concurrency" chunks of this size are buffered
	// in memory per transfer.
	//
	// If you are transferring large files over high-speed links and you have
	// enough memory, then increasing this will speed up the transfers.
	//
	// Rclone will automatically increase the chunk size when uploading a
	// large file of known size to stay below the 10,000 chunks limit.
	//
	// Files of unknown size are uploaded with the configured
	// chunk_size. Since the default chunk size is 5MB and there can be at
	// most 10,000 chunks, this means that by default the maximum size of
	// a file you can stream upload is 48GB.  If you wish to stream upload
	// larger files then you will need to increase chunk_size.
	ChunkSize string `yaml:"chunk_size"`
	// Maximum number of parts in a multipart upload.
	//
	// This option defines the maximum number of multipart chunks to use
	// when doing a multipart upload.
	//
	// This can be useful if a service does not support the AWS S3
	// specification of 10,000 chunks.
	//
	// Rclone will automatically increase the chunk size when uploading a
	// large file of a known size to stay below this number of chunks limit.
	MaxUploadParts string `yaml:"max_upload_parts"`
	// Cutoff for switching to multipart copy
	//
	// Any files larger than this that need to be server-side copied will be
	// copied in chunks of this size.
	//
	// The minimum is 0 and the maximum is 5GB.
	CopyCutoff string `yaml:"copy_cutoff"`
	// Don't store MD5 checksum with object metadata
	//
	// Normally rclone will calculate the MD5 checksum of the input before
	// uploading it so it can add it to metadata on the object. This is great
	// for data integrity checking but can cause long delays for large files
	// to start uploading.
	DisableChecksum string `yaml:"disable_checksum"`
	// Path to the shared credentials file
	//
	// If env_auth = true then rclone can use a shared credentials file.
	//
	// If this variable is empty rclone will look for the
	// "AWS_SHARED_CREDENTIALS_FILE" env variable. If the env value is empty
	// it will default to the current user's home directory.
	//
	//	Linux/OSX: "$HOME/.aws/credentials"
	//	Windows:   "%USERPROFILE%\.aws\credentials"
	SharedCredentialsFile string `yaml:"shared_credentials_file"`
	// Profile to use in the shared credentials file
	//
	// If env_auth = true then rclone can use a shared credentials file. This
	// variable controls which profile is used in that file.
	//
	// If empty it will default to the environment variable "AWS_PROFILE" or
	// "default" if that environment variable is also not set.
	Profile string `yaml:"profile"`
	// An AWS session token
	SessionToken string `yaml:"session_token"`
	// Concurrency for multipart uploads.
	//
	// This is the number of chunks of the same file that are uploaded
	// concurrently.
	//
	// If you are uploading small numbers of large files over high-speed links
	// and these uploads do not fully utilize your bandwidth, then increasing
	// this may help to speed up the transfers.
	UploadConcurrency string `yaml:"upload_concurrency"`
	// If true use path style access if false use virtual hosted style.
	//
	// If this is true (the default) then rclone will use path style access,
	// if false then rclone will use virtual path style. See [the AWS S3
	// docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
	// for more info.
	//
	// Some providers (e.g. AWS, Aliyun OSS, Netease COS, or Tencent COS) require this set to
	// false - rclone will do this automatically based on the provider
	// setting.
	ForcePathStyle string `yaml:"force_path_style"`
	// If true use v2 authentication.
	//
	// If this is false (the default) then rclone will use v4 authentication.
	// If it is set then rclone will use v2 authentication.
	//
	// Use this only if v4 signatures don't work, e.g. pre Jewel/v10 CEPH.
	V2Auth string `yaml:"v2_auth"`
	// If true use the AWS S3 accelerated endpoint.
	//
	// See: [AWS S3 Transfer acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration-examples.html)
	UseAccelerateEndpoint string `yaml:"use_accelerate_endpoint"`
	// If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	//
	// It should be set to true for resuming uploads across different sessions.
	//
	// WARNING: Storing parts of an incomplete multipart upload counts towards space usage on S3 and will add additional costs if not cleaned up.
	LeavePartsOnError string `yaml:"leave_parts_on_error"`
	// Size of listing chunk (response list for each ListObject S3 request).
	//
	// This option is also known as "MaxKeys", "max-items", or "page-size" from the AWS S3 specification.
	// Most services truncate the response list to 1000 objects even if requested more than that.
	// In AWS S3 this is a global maximum and cannot be changed, see [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html).
	// In Ceph, this can be increased with the "rgw list buckets max chunk" option.
	ListChunk string `yaml:"list_chunk"`
	// If set, don't attempt to check the bucket exists or create it
	//
	// This can be useful when trying to minimise the number of transactions
	// rclone does if you know the bucket exists already.
	//
	// It can also be needed if the user you are using does not have bucket
	// creation permissions. Before v1.52.0 this would have passed silently
	// due to a bug.
	NoCheckBucket string `yaml:"no_check_bucket"`
	// If set, don't HEAD uploaded objects to check integrity
	//
	// This can be useful when trying to minimise the number of transactions
	// rclone does.
	//
	// Setting it means that if rclone receives a 200 OK message after
	// uploading an object with PUT then it will assume that it got uploaded
	// properly.
	//
	// In particular it will assume:
	//
	// - the metadata, including modtime, storage class and content type was as uploaded
	// - the size was as uploaded
	//
	// It reads the following items from the response for a single part PUT:
	//
	// - the MD5SUM
	// - The uploaded date
	//
	// For multipart uploads these items aren't read.
	//
	// If an source object of unknown length is uploaded then rclone **will** do a
	// HEAD request.
	//
	// Setting this flag increases the chance for undetected upload failures,
	// in particular an incorrect size, so it isn't recommended for normal
	// operation. In practice the chance of an undetected upload failure is
	// very small even with this flag.
	NoHead string `yaml:"no_head"`
	// This sets the encoding for the backend.
	//
	// See: the [encoding section in the overview](/overview/#encoding) for more info.
	Encoding string `yaml:"encoding"`
	// How often internal memory buffer pools will be flushed.
	// Uploads which requires additional buffers (f.e multipart) will use memory pool for allocations.
	// This option controls how often unused buffers will be removed from the pool.
	MemoryPoolFlushTime string `yaml:"memory_pool_flush_time"`
	// Whether to use mmap buffers in internal memory pool.
	MemoryPoolUseMmap string `yaml:"memory_pool_use_mmap"`
	// Disable usage of http2 for S3 backends
	//
	// There is currently an unsolved issue with the s3 (specifically minio) backend
	// and HTTP/2.  HTTP/2 is enabled by default for the s3 backend but can be
	// disabled here.  When the issue is solved this flag will be removed.
	//
	// See: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
	DisableHttp2 string `yaml:"disable_http2"`
}

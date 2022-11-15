package Filesystem

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Filesystem/adapters"
	"github.com/kmsar/laravel-go/Framework/Logs"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"
	"github.com/kmsar/laravel-go/Framework/Support/Field"

	"github.com/kmsar/laravel-go/Framework/Contracts/IFilesystem"
	"io/fs"
	"time"
)

var (
	UndefinedDefineErr = errors.New("unsupported file system")
)

func New(config Config) IFilesystem.FileSystemFactory {
	return &Factory{
		config: config,
		disks:  make(map[string]IFilesystem.FileSystem),
		drivers: map[string]IFilesystem.FileSystemProvider{
			"local": adapters.LocalAdapter,
			//"qiniu": adapters.QiniuAdapter,
		},
	}

}

type Factory struct {
	config  Config
	disks   map[string]IFilesystem.FileSystem
	drivers map[string]IFilesystem.FileSystemProvider
}

func (this *Factory) Disk(name string) IFilesystem.FileSystem {
	if disk, existsStore := this.disks[name]; existsStore {
		return disk
	}

	this.disks[name] = this.get(name)

	return this.disks[name]
}

func (this *Factory) Extend(driver string, provider IFilesystem.FileSystemProvider) {
	this.drivers[driver] = provider
}

func (this *Factory) get(name string) IFilesystem.FileSystem {
	var (
		config = this.config.Disks[name]
		driver = Field.GetStringField(config, "driver", this.config.Default)
	)
	var driveProvider, existsProvider = this.drivers[driver]
	if !existsProvider {
		Logs.WithError(UndefinedDefineErr).Error(fmt.Sprintf("filesystem.Factory: unsupported file system %s", driver))
		panic(Exception{Exceptions.WithError(UndefinedDefineErr, config)})
	}
	return driveProvider(name, config)
}

func (this *Factory) Name() string {
	return this.Disk(this.config.Default).Name()
}

func (this *Factory) Exists(path string) bool {
	return this.Disk(this.config.Default).Exists(path)
}

func (this *Factory) Get(path string) (string, error) {
	return this.Disk(this.config.Default).Get(path)
}

func (this *Factory) Read(path string) ([]byte, error) {
	return this.Disk(this.config.Default).Read(path)
}

func (this *Factory) ReadStream(path string) (*bufio.Reader, error) {
	return this.Disk(this.config.Default).ReadStream(path)
}

func (this *Factory) Put(path, contents string) error {
	return this.Disk(this.config.Default).Put(path, contents)
}

func (this *Factory) WriteStream(path string, contents string) error {
	return this.Disk(this.config.Default).WriteStream(path, contents)
}

func (this *Factory) GetVisibility(path string) IFilesystem.FileVisibility {
	return this.Disk(this.config.Default).GetVisibility(path)
}

func (this *Factory) SetVisibility(path string, perm fs.FileMode) error {
	return this.Disk(this.config.Default).SetVisibility(path, perm)
}

func (this *Factory) Prepend(path, contents string) error {
	return this.Disk(this.config.Default).Prepend(path, contents)
}

func (this *Factory) Append(path, contents string) error {
	return this.Disk(this.config.Default).Append(path, contents)
}

func (this *Factory) Delete(path string) error {
	return this.Disk(this.config.Default).Delete(path)
}

func (this *Factory) Copy(from, to string) error {
	return this.Disk(this.config.Default).Copy(from, to)
}

func (this *Factory) Move(from, to string) error {
	return this.Disk(this.config.Default).Move(from, to)
}

func (this *Factory) Size(path string) (int64, error) {
	return this.Disk(this.config.Default).Size(path)
}

func (this *Factory) LastModified(path string) (time.Time, error) {
	return this.Disk(this.config.Default).LastModified(path)
}

func (this *Factory) Files(directory string) []IFilesystem.File {
	return this.Disk(this.config.Default).Files(directory)
}

func (this *Factory) AllFiles(directory string) []IFilesystem.File {
	return this.Disk(this.config.Default).AllFiles(directory)
}

func (this *Factory) Directories(directory string) []string {
	return this.Disk(this.config.Default).Directories(directory)
}

func (this *Factory) AllDirectories(directory string) []string {
	return this.Disk(this.config.Default).AllDirectories(directory)
}

func (this *Factory) MakeDirectory(path string) error {
	return this.Disk(this.config.Default).MakeDirectory(path)
}

func (this *Factory) DeleteDirectory(directory string) error {
	return this.Disk(this.config.Default).DeleteDirectory(directory)
}

package Container

//
//https://laravel.com/docs/6.x/container
//$api = new \HelpSpot\API(new HttpClient);
//
//$this->app->instance('HelpSpot\API', $api)
//
//$this->app->when(PhotoController::class)
//->needs(Filesystem::class)
//->give(function () {
//return Storage::disk('local');
//});
//
//$this->app->when([VideoController::class, UploadController::class])
//->needs(Filesystem::class)
//->give(function () {
//return Storage::disk('s3');
//});
//
//$this->app->resolving(function ($object, $app) {
//// Called when container resolves object of any type...
//});

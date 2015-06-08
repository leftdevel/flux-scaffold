# flux-scaffold
A Facebok Flux scaffold generator in Golang

#### TL;DR
Yes, I know, I should have done it in an npm package instead ;) But my main goal was to learn a bit of Go.
I  also didn't include tests for now, please don't hate me.

Although Go is pretty easy to pick, I struggled to get nice composition and avoid repeating code. Gladly I think
I achieved it :)

Files api.go, component.go, constant.go, action.go and store.go implements resource-interface. resource-generator depends on this interface, so it stays abstracted away from each resource implementation details.

### Usage
Afer you have downloaded and installed the packaged you will be able to invoke `fluxscaffold` command.

A full example of how to use it from a terminal:

```
$ cd /into/your/flux/root/dir/
$ fluxscaffold -domain=product -const=PRODUCTS_FETCH_ALL -action=fetchAll -api
```

The available flags are:
* domain: A domain object name, e.g. "user", "product", "pizza". This value will be used to infer the constant,
action, store, component and webapi file name. For instance, if you choose "user" you will get "constants/user-constants.js",
"actions/user-actions.js", "stores/user-store.js", "webapi/user-api.js", "components/user.js"

* action: A single action method, e.g. "fetchAll". This will create an action method with that name inside the
action creator file.

* const: A single constant name, e.g. "USERS_FETCH_ALL". This will be referenced in the constant, action and store
files.

* api: Optional. No value required, you can invoke it just like "-api". This will tell the tool to also scaffold a
webapi file.

### 'api' Flag
Besides telling the tool to generate a webapi file, it will also modify the content of the constant, action, component
and store files.

If you passed `-const=USERS_FETCH_ALL`, the constant file will looks like this:

```
  UserConstants = keymirror({
    USERS_FETCH_ALL: null,
    USERS_FETCH_ALL_SUCCESS: null
  });
```

This is a little magic to help with fetching collections. If the api flag is present we will assume the default
action is to fetch a domain object collection. You will notice a different signature in the rest of the files as well.

For example, with no api flag, the store  will listen to "USERS_FETCH_ALL". If the api flag is present,
the store will listen to "USERS_FETCH_ALL_SUCCESS" instead, since we assume the `users` collection is only available
until the webapi request has "succeeded".

By the way, when passing the api flag, you don't need to change the domain value to "users". You should keep
using "user". The tool will append an "s" to the domain value where necessary. Of course, this won't be grammatically accurate
100% of the time, but it should to work for at least 90% of the time.

Keep in mind you will need to manually update the files. This tool is not meant to fire & forget, but to avoid
the tedious work of creating and bootstrapping each flux file.

### Is this vanilla flux?
Yes, it is. I've been working with Flux since almost a year from now, and I feel more confortable with vanilla flux
than with other solutions out there.

### Is this a serious project?
Not at all. But if you think this could be useful to you then drop me a line. If you prefer a npm package I'll be
happy to create one, but only if it's of any interest.

Feel free to request any feature, or to refactor anything as long as it goes with the default flux conventions.

### Gotchas
If you generated an api file, you will noticed it was created under webapi/ directory. All directories are hardcoded by now
but I'm planning to make them flexible. I wasn't sure about what would be the default web api utils directory becuase
in the flux chatapp demo, they use a utils/ directory and a WebApiUtil file. They also mix some other non api utils under
that directory, so I prefered to use a proper directory for api calls only.

In the generated api file, you will notice a reference to `ApiResponseHandler`. I found myself repeating the same response
handler code over and over, so I decided to move that common code to a separate file. The file could read something like this:

```
module.exports = {
    handle: function(err, res, successCallback) {
        if (err) {
            // @TODO add proper error handling
            alert('ERROR!!!!!');
            console.log(err);
            return;
        }

        successCallback(JSON.parse(res.text));
    }
};
```

It's not a full featured handler ("alerting" is pretty ugly btw), but you can extend it to your needs.

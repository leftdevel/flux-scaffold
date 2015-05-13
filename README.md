# flux-scaffold
A Facebok Flux scaffold generator in Golang

#### TL;DR
Yes, I know, I should have done it in an npm package instead ;) But my main goal was to learn a bit of Go.
I haven't also included tests, please don't hate me.

### Usage
Afer you have downloaded and installed the packaged you will be able to invoke `fluxscaffold` command.
The available flags are:
* domain: A domain object name, e.g. "user", "product", "pizza". This value will be used to infer the constant,
action, store, and webapi file name. For instance, if you choose "user" you will get "constants/user-constants.js",
"actions/user-actions.js", "stores/user-store.js", "webapi/user-api.js"

* action: A single action method, e.g. "fetchAll". This will create an action method with that name inside the
action creator file.

* constant: A single constant name, e.g. "USERS_FETCH_ALL". This will be referenced in the constant, action and store
files.

* api: Optional. No value required, you can invoke it just like "-api". This will tell the tool to also scaffold a
webapi file.

### 'api' Flag
Besides telling the tool to generate a webapi file, it will also modify the content of the constant, action and store
files.

If you passed `-constant=USERS_FETCH_ALL`, the constant file will looks like this:

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

Keep in mind you will need to manually update the files. This tool is not meant to fire & forget, but to avoid
the tedious work of creating and boostraping each flux file.

### Is this vanilla flux?
Yes, it is. I've been working with Flux since almost a year from now, and I feel more confortable with vanilla flux
than with other solutions out there.

### Is this a serious project?
Not at all. But if you think this could be useful to you then drop me a line. If you prefer a npm package I'll be
happy to create one, but only if it's of any interest.

Feel free to request any feature, or to refactor anything as long as it goes with the default flux conventions.
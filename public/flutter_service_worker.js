'use strict';
const MANIFEST = 'flutter-app-manifest';
const TEMP = 'flutter-temp-cache';
const CACHE_NAME = 'flutter-app-cache';

const RESOURCES = {"flutter.js": "f393d3c16b631f36852323de8e583132",
"favicon.png": "5dcef449791fa27946b3d35ad8803796",
"index.html": "d2df2c48d404bf59875fbcb75e3f0969",
"/": "d2df2c48d404bf59875fbcb75e3f0969",
"version.json": "ac1a8d00bd01e2290045fb8f7d0419ec",
"icons/Icon-192.png": "ac9a721a12bbc803b44f645561ecb1e1",
"icons/Icon-maskable-512.png": "301a7604d45b3e739efc881eb04896ea",
"icons/Icon-512.png": "96e752610906ba2a93c65f8abe1645f1",
"icons/Icon-maskable-192.png": "c457ef57daa1d16f64b27b786ec2ea3c",
"main.dart.js": "7d6894743bda37ac0d37283dd32199bc",
"flutter_bootstrap.js": "7ab5aff830f48c0c2598733a613d986b",
"assets/NOTICES": "7f99b18380cb64e7702c05acc9130e48",
"assets/shaders/ink_sparkle.frag": "ecc85a2e95f5e9f53123dcaf8cb9b6ce",
"assets/AssetManifest.bin": "3805e479aa635c503dd15c4ffb921eee",
"assets/packages/bego_ui/assets/icons/BegoIcons.ttf": "ea9c0193cd166c5bad4f9374b7523018",
"assets/packages/bego_ui/assets/fonts/ReadexPro-Regular.ttf": "7b7b051bc5a9e8afed7313e794550ab9",
"assets/packages/bego_ui/assets/fonts/ReadexPro-ExtraLight.ttf": "06b70785d73c1cc252ba8c75b7e74106",
"assets/packages/bego_ui/assets/fonts/ReadexPro-Medium.ttf": "ae2b7f6d8d1792897e0270e122e77b7a",
"assets/packages/bego_ui/assets/fonts/ReadexPro-Light.ttf": "a822bf9824dbb3ea2c550027d90a2052",
"assets/packages/bego_ui/assets/fonts/ReadexPro-Bold.ttf": "37da4efe33b3e774e7585dd4f285058b",
"assets/packages/bego_ui/assets/fonts/ReadexPro-SemiBold.ttf": "88de789da14c49cd8e33611dfeb0636e",
"assets/packages/bego_app/assets/images/working_illustration.png": "81045f5b6507e78b8fa5c53df7598422",
"assets/AssetManifest.json": "74495ffae6d2569cfcd623deb867d505",
"assets/assets/svg/business_ocean.svg": "908f8816912a03e0c4e64ff2f790c29e",
"assets/assets/svg/blob_3.svg": "8588571132399a6db9638b5c33a714d3",
"assets/assets/svg/bego_icon.svg": "4059b8f59a471186e2b916714e33f1a0",
"assets/assets/svg/blob_1.svg": "5c51083d8a14679ddbbdc06c5686cc37",
"assets/assets/svg/stacked_waves.svg": "9ec831d980ab3d8b51d6c2f97aa0b638",
"assets/assets/svg/blob_2.svg": "14c5f35b6426a7d7fba40f99f1a56141",
"assets/assets/svg/google_logo.svg": "64ac7e5288152732be55ea53a0d62d45",
"assets/assets/image/bg_abstract.jpeg": "fc767b445747905594214fd4d7ae4768",
"assets/assets/image/header_bg_1.jpg": "9e4ea2dcbf4e88d05a41ad875fed31c8",
"assets/assets/image/header_bg_3.jpg": "b8cb1502cb8c0591407f712342d9ab93",
"assets/assets/image/header_bg_2.jpg": "61888f867b9f09bd8b4631bbcf3dc971",
"assets/assets/logo/bego_logo.png": "286a226fa32d3e7a996aaae90f9613c3",
"assets/fonts/MaterialIcons-Regular.otf": "54ad14caae9a181a141f5831d7f71cd0",
"assets/AssetManifest.bin.json": "386c81e04164352271c9ff7ef33999ce",
"assets/FontManifest.json": "b2d53c5d6805104ca5bd7c73b212b59c",
"canvaskit/skwasm.worker.js": "89990e8c92bcb123999aa81f7e203b1c",
"canvaskit/skwasm.wasm": "9f0c0c02b82a910d12ce0543ec130e60",
"canvaskit/skwasm.js.symbols": "262f4827a1317abb59d71d6c587a93e2",
"canvaskit/skwasm.js": "694fda5704053957c2594de355805228",
"canvaskit/canvaskit.js.symbols": "48c83a2ce573d9692e8d970e288d75f7",
"canvaskit/chromium/canvaskit.js.symbols": "a012ed99ccba193cf96bb2643003f6fc",
"canvaskit/chromium/canvaskit.js": "671c6b4f8fcc199dcc551c7bb125f239",
"canvaskit/chromium/canvaskit.wasm": "b1ac05b29c127d86df4bcfbf50dd902a",
"canvaskit/canvaskit.js": "66177750aff65a66cb07bb44b8c6422b",
"canvaskit/canvaskit.wasm": "1f237a213d7370cf95f443d896176460",
"manifest.json": "f64a4be80754c3e6eb0b6cc940906f20"};
// The application shell files that are downloaded before a service worker can
// start.
const CORE = ["main.dart.js",
"index.html",
"flutter_bootstrap.js",
"assets/AssetManifest.bin.json",
"assets/FontManifest.json"];

// During install, the TEMP cache is populated with the application shell files.
self.addEventListener("install", (event) => {
  self.skipWaiting();
  return event.waitUntil(
    caches.open(TEMP).then((cache) => {
      return cache.addAll(
        CORE.map((value) => new Request(value, {'cache': 'reload'})));
    })
  );
});
// During activate, the cache is populated with the temp files downloaded in
// install. If this service worker is upgrading from one with a saved
// MANIFEST, then use this to retain unchanged resource files.
self.addEventListener("activate", function(event) {
  return event.waitUntil(async function() {
    try {
      var contentCache = await caches.open(CACHE_NAME);
      var tempCache = await caches.open(TEMP);
      var manifestCache = await caches.open(MANIFEST);
      var manifest = await manifestCache.match('manifest');
      // When there is no prior manifest, clear the entire cache.
      if (!manifest) {
        await caches.delete(CACHE_NAME);
        contentCache = await caches.open(CACHE_NAME);
        for (var request of await tempCache.keys()) {
          var response = await tempCache.match(request);
          await contentCache.put(request, response);
        }
        await caches.delete(TEMP);
        // Save the manifest to make future upgrades efficient.
        await manifestCache.put('manifest', new Response(JSON.stringify(RESOURCES)));
        // Claim client to enable caching on first launch
        self.clients.claim();
        return;
      }
      var oldManifest = await manifest.json();
      var origin = self.location.origin;
      for (var request of await contentCache.keys()) {
        var key = request.url.substring(origin.length + 1);
        if (key == "") {
          key = "/";
        }
        // If a resource from the old manifest is not in the new cache, or if
        // the MD5 sum has changed, delete it. Otherwise the resource is left
        // in the cache and can be reused by the new service worker.
        if (!RESOURCES[key] || RESOURCES[key] != oldManifest[key]) {
          await contentCache.delete(request);
        }
      }
      // Populate the cache with the app shell TEMP files, potentially overwriting
      // cache files preserved above.
      for (var request of await tempCache.keys()) {
        var response = await tempCache.match(request);
        await contentCache.put(request, response);
      }
      await caches.delete(TEMP);
      // Save the manifest to make future upgrades efficient.
      await manifestCache.put('manifest', new Response(JSON.stringify(RESOURCES)));
      // Claim client to enable caching on first launch
      self.clients.claim();
      return;
    } catch (err) {
      // On an unhandled exception the state of the cache cannot be guaranteed.
      console.error('Failed to upgrade service worker: ' + err);
      await caches.delete(CACHE_NAME);
      await caches.delete(TEMP);
      await caches.delete(MANIFEST);
    }
  }());
});
// The fetch handler redirects requests for RESOURCE files to the service
// worker cache.
self.addEventListener("fetch", (event) => {
  if (event.request.method !== 'GET') {
    return;
  }
  var origin = self.location.origin;
  var key = event.request.url.substring(origin.length + 1);
  // Redirect URLs to the index.html
  if (key.indexOf('?v=') != -1) {
    key = key.split('?v=')[0];
  }
  if (event.request.url == origin || event.request.url.startsWith(origin + '/#') || key == '') {
    key = '/';
  }
  // If the URL is not the RESOURCE list then return to signal that the
  // browser should take over.
  if (!RESOURCES[key]) {
    return;
  }
  // If the URL is the index.html, perform an online-first request.
  if (key == '/') {
    return onlineFirst(event);
  }
  event.respondWith(caches.open(CACHE_NAME)
    .then((cache) =>  {
      return cache.match(event.request).then((response) => {
        // Either respond with the cached resource, or perform a fetch and
        // lazily populate the cache only if the resource was successfully fetched.
        return response || fetch(event.request).then((response) => {
          if (response && Boolean(response.ok)) {
            cache.put(event.request, response.clone());
          }
          return response;
        });
      })
    })
  );
});
self.addEventListener('message', (event) => {
  // SkipWaiting can be used to immediately activate a waiting service worker.
  // This will also require a page refresh triggered by the main worker.
  if (event.data === 'skipWaiting') {
    self.skipWaiting();
    return;
  }
  if (event.data === 'downloadOffline') {
    downloadOffline();
    return;
  }
});
// Download offline will check the RESOURCES for all files not in the cache
// and populate them.
async function downloadOffline() {
  var resources = [];
  var contentCache = await caches.open(CACHE_NAME);
  var currentContent = {};
  for (var request of await contentCache.keys()) {
    var key = request.url.substring(origin.length + 1);
    if (key == "") {
      key = "/";
    }
    currentContent[key] = true;
  }
  for (var resourceKey of Object.keys(RESOURCES)) {
    if (!currentContent[resourceKey]) {
      resources.push(resourceKey);
    }
  }
  return contentCache.addAll(resources);
}
// Attempt to download the resource online before falling back to
// the offline cache.
function onlineFirst(event) {
  return event.respondWith(
    fetch(event.request).then((response) => {
      return caches.open(CACHE_NAME).then((cache) => {
        cache.put(event.request, response.clone());
        return response;
      });
    }).catch((error) => {
      return caches.open(CACHE_NAME).then((cache) => {
        return cache.match(event.request).then((response) => {
          if (response != null) {
            return response;
          }
          throw error;
        });
      });
    })
  );
}

const CACHE_NAME = 'my-pwa-cache-v1';
const CACHE_FILES = [
  '/auth/login',
  '/auth/parent/login',
  '/auth/signup',
  '/dashboard',
  '/css/styles.css',
  '/js/dist_htmx.min.js',
  '/img/site/logo.png',
];

self.addEventListener('install', function(event) {
    event.waitUntil(
        caches.open(CACHE_NAME).then(function(cache) {
            return cache.addAll(CACHE_FILES).catch(function(error) {
            console.error('Caching failed:', error);
            });
        })
    )
});

self.addEventListener('fetch', function(event) {
  event.respondWith(
    caches.match(event.request).then(function(response) {
      return response || fetch(event.request);
    })
  );
});


// self.addEventListener('fetch', function (event) {
//     const request = event.request;
//     if (request.url === '/') {
//       event.respondWith(fetch('/auth/login'));
//     }
//   });

// self.addEventListener('fetch', function(event) {
//   const requestUrl = new URL(event.request.url);

//   if (
//     requestUrl === '/' &&
//     requestUrl === '/about'
//   ) {
    
//     event.respondWith(caches.match('/auth/login').then(function(response) {
//       return response || fetch('/auth/login');
//     })); // Serve an error page
//   } else {
//     caches.match(event.request).then(function(response) {
//       return response || fetch(event.request);
//     }) // Serve the requested page
//   }
// });


// self.addEventListener('fetch', (event) => {
//   const requestUrl = new URL(event.request.url);

//   // Check if the PWA is installable and the requested URL is the homepage or about page
//   if (isPwaInstallable && (requestUrl.pathname === '/' || requestUrl.pathname === '/about')) {
//     // Redirect to the login page
//     event.respondWith(fetch('/auth/login'));
//   } else {
//     // Serve the requested page or resource
//     event.respondWith(fetch(event.request));
//   }
// });




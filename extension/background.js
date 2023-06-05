browser.action.onClicked.addListener(function (tab, _data) {
  fetch('https://link.giodamelio.com', {
    method: 'POST',
    body: tab.url,
    headers: {
      'Content-Type': 'text/plain'
    }
  })
    .then(response => {
      if (response.ok) {
        browser.notifications.create('link-notification', {
          type: 'basic',
          iconUrl: browser.runtime.getURL('icons/icon-dark-64.png'),
          title: 'Link set',
          message: tab.url,
        });
      } else {
        browser.notifications.create('link-notification', {
          type: 'basic',
          iconUrl: browser.runtime.getURL('icons/icon-dark-64.png'),
          title: 'Link could not be set',
          message: tab.url,
        });
      }
    })
    .catch(error => console.log(error));
});

import {makeFramesHover} from './frames.script.js?static'

const loader = document.createElement('div');
loader.style.cssText = 'position:fixed;inset:0;background:#e0e0e0;z-index:9999;';
document.body.appendChild(loader);

import('/.meta/layout.js?static')
  .then(() => import('/.meta/cards.js?static'))
  .then((cards) => cards.ready)
  .then(() => loader.remove());

makeFramesHover('#grid-artworks-for-sale')

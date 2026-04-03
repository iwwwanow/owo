// Fix image quality on span-N cards (span-2, span-4)
document.querySelectorAll('.card__wrapper').forEach((card) => {
  const picture = card.querySelector('picture.card__cover-container');
  if (!picture) return;

  const w = Math.round(card.getBoundingClientRect().width);
  const h = Math.round(card.getBoundingClientRect().height);

  if (w <= 190) return;

  const source = picture.querySelector('source:not([media])');
  if (source) {
    const base = source.srcset.split(',')[0].trim().split(/\s+/)[0];
    const make = (scale) => {
      const url = new URL(base, location.origin);
      url.searchParams.set('width', w * scale);
      url.searchParams.set('height', h * scale);
      return `${url.pathname}${url.search} ${scale}x`;
    };
    source.srcset = `${make(1)}, ${make(2)}`;
  }

  const img = picture.querySelector('img');
  if (img) {
    const url = new URL(img.src, location.origin);
    url.searchParams.set('width', w);
    url.searchParams.set('height', h);
    img.src = `${url.pathname}${url.search}`;
  }
});

// xtc-toaster: image 1/6 card width, full width in tile mode
const xtcCard = document.getElementById('card-_featured-projects-xtc-toaster');
if (xtcCard) {
  const picture = xtcCard.querySelector('picture.card__cover-container');
  if (picture) {
    const img = picture.querySelector('img');
    if (img) {
      const tileW = Math.round(xtcCard.getBoundingClientRect().width / 6);

      const url = new URL(img.src, location.origin);
      url.searchParams.set('width', tileW);
      url.searchParams.delete('height');

      img.style.display = 'none';
      picture.style.cssText = 'display: block; width: 100%; height: 100%;';
      picture.style.backgroundImage = `url('${url.pathname}${url.search}')`;
      picture.style.backgroundRepeat = 'repeat';
      picture.style.backgroundSize = `${tileW}px auto`;

      let x = 0, y = 0;
      const speed = 0.4;
      const animate = () => {
        x = (x + speed) % tileW;
        y = (y + speed) % tileW;
        picture.style.backgroundPosition = `${x}px ${y}px`;
        requestAnimationFrame(animate);
      };
      requestAnimationFrame(animate);
    }
  }
}

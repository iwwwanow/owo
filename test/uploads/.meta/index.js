document.addEventListener('DOMContentLoaded', () => {
  // Fix image quality for span-2/span-4 cards
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

  // xtc-toaster: thumbnail 1/6 card width, full width in tile mode
  const xtcCard = document.getElementById('card-_featured-projects-xtc-toaster');
  if (xtcCard) {
    const picture = xtcCard.querySelector('picture.card__cover-container');
    if (picture) {
      picture.style.cssText = 'width: calc(100% / 6); flex-shrink: 0;';

      const grid = xtcCard.closest('.grid');
      if (grid) {
        const update = () => {
          const isTile = grid.classList.contains('grid--hidden');
          picture.style.width = isTile ? '100%' : 'calc(100% / 6)';
        };
        new MutationObserver(update).observe(grid, {
          attributes: true,
          attributeFilter: ['class'],
        });
      }
    }
  }
});

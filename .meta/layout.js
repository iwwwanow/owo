const rightContent = document.querySelector('.grid__right-content');
if (rightContent) {
  const grid = rightContent.closest('.grid');
  [...rightContent.children].forEach((el) => grid.appendChild(el));
  rightContent.remove();
}

const grids = document.querySelectorAll('.grid');
if (grids.length >= 2) {
  const first = grids[1];
  const second = grids[2];
  first.id = 'grid-common';
  second.querySelectorAll('.card__wrapper').forEach((card) => first.appendChild(card));
  second.remove();
}

function placeCard(cardId, gridId, position) {
  const grid = document.getElementById(gridId);
  const card = document.getElementById(cardId);
  if (!grid || !card) return;

  const cards = [...grid.querySelectorAll('.card__wrapper')];
  card.remove();

  const target = cards.filter((c) => c !== card)[position];
  target ? grid.insertBefore(card, target) : grid.appendChild(card);
}

function placeGrid(gridId, position) {
  const wrapper = document.querySelector('.wrapper');
  const grid = document.getElementById(gridId);
  if (!wrapper || !grid) return;

  const grids = [...wrapper.querySelectorAll('.grid')];
  grid.remove();

  const target = grids.filter((g) => g !== grid)[position];
  target ? wrapper.insertBefore(grid, target) : wrapper.appendChild(grid);
}

placeCard('card-Grafika%2C%20Zhivopis', 'grid-common', 3);
placeCard('card-Code', 'grid-common', 4);
placeCard('card-Digital%20art', 'grid-common', 5);
placeCard('card-Fotografiia', 'grid-common', 6);
placeCard('card-About', 'grid-common', 7);

placeGrid('grid-featured-projects', 2);
placeCard("card-_featured-projects-Put'-k-khramu-Bogini-Tsvetov", 'grid-featured-projects', 0);
placeCard('card-_featured-projects-owo', 'grid-featured-projects', 1);
placeCard('card-_featured-projects-xtc-toaster', 'grid-featured-projects', 2);

placeGrid('grid-artworks-for-sale', 3);
placeCard('card-_artworks-for-sale-Shaman', 'grid-artworks-for-sale', 0);
placeCard('card-_artworks-for-sale-Faust', 'grid-artworks-for-sale', 1);
placeCard('card-_artworks-for-sale-Shaman-2', 'grid-artworks-for-sale', 2);
placeCard("card-_artworks-for-sale-Put'-k-khramu-Bogini-Tsvetov", 'grid-artworks-for-sale', 3);

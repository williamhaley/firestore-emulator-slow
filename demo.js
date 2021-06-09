const admin = require('firebase-admin');

admin.initializeApp();

const db = admin.firestore();

(async () => {
  const docRef = db.collection('users').doc('alovelace');

  const result = await docRef.set({
    first: 'Ada',
    last: 'Lovelace',
    born: 1815
  });

  console.log('write done', result);

  const snap = await docRef.get();
  console.log('read done', snap.data());
})();
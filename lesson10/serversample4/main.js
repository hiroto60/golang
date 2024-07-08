// 全ての本を取得して表示する関数
function getBooks() {
    fetch('http://localhost:8080/books')
      .then(response => response.json())
      .then(books => {
        const tbody = document.querySelector('#bookTable tbody');
        tbody.innerHTML = '';
        books.forEach(book => {
          const row = document.createElement('tr');
          row.innerHTML = `
                    <td>${book.id}</td>
                    <td>${book.title}</td>
                    <td>${book.price}</td>
                    <td>${new Date(book.created_at).toLocaleString()}</td>
                `;
          tbody.appendChild(row);
        });
      });
  }
  
  // 新しい本を追加する関数
  function addBook(title, price) {
    const book = {
      title,
      price: parseInt(price, 10)
    };
    fetch('http://localhost:8080/books', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(book)
    })
      .then(() => {
        getBooks();
        document.querySelector('#bookForm').reset();
      });
  }
  
  // ページ読み込み時に全ての本を取得して表示
  document.addEventListener('DOMContentLoaded', getBooks);
  
  // 新しい本を追加するフォームの送信イベント
  document.querySelector('#bookForm').addEventListener('submit', event => {
    event.preventDefault();
    const title = document.querySelector('#title').value;
    const price = document.querySelector('#price').value;
    addBook(title, price);
  });
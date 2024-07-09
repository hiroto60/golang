
// 全ての記事を取得して表示する関数
function getPosts() {
  fetch('http://localhost:8080/posts')
    .then(response => response.json())
    .then(posts => {
      const tbody = document.querySelector('#postTable tbody');
      tbody.innerHTML = '';
      posts.forEach(post => {
        const row = document.createElement('tr');
        row.innerHTML = `
                  <td>${post.id}</td>
                  <td>${post.title}</td>
                  <td>${post.content}</td>
                  <td>${post.likes}</td>
                  <td>${new Date(post.created_at).toLocaleString()}</td>
                  <td><button class="like-btn" data-id="${post.id}">Like</button></td>`;
        tbody.appendChild(row);
      });

      // いいねボタンのクリックイベントを追加
      const likeButtons = document.querySelectorAll('.like-btn');
      likeButtons.forEach(button => {
        button.addEventListener('click', () => {
          const postId = button.dataset.id;
          likePost(postId);
        });
      });
    });
}

// 新しい記事を追加する関数
function addPost(title, content) {
  const post = {
    title,
    content
  };

  fetch('http://localhost:8080/posts', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(post)
  })
    .then(() => {
      getPosts();
      document.querySelector('#postForm').reset();
    });
}

// 記事にいいねを追加する関数
function likePost(postId) {
  fetch(`http://localhost:8080/posts/${postId}`, {
    method: 'POST'
  })
    .then(() => {
      getPosts();
    });
}

// ページ読み込み時に全ての記事を取得して表示
document.addEventListener('DOMContentLoaded', getPosts);

// 新しい記事を追加するフォームの送信イベント
document.querySelector('#postForm').addEventListener('submit', event => {
  event.preventDefault();
  const title = document.querySelector('#title').value;
  const content = document.querySelector('#content').value;
  addPost(title, content);
});  
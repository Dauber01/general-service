// 获取表单元素
const username = document.getElementById('username');
const password = document.getElementById('password');
const form = document.querySelector('form');

// 添加表单提交事件
form.addEventListener('submit', function(event) {
  // 阻止表单默认提交行为
  event.preventDefault();

  // 获取表单数据
  const usernameValue = username.value.trim();
  const passwordValue = password.value.trim();

  // 进行简单的表单验证
  if (usernameValue === '' || passwordValue === '') {
    alert('用户名和密码不能为空！');
    return;
  }

  // 发送登录请求
  // 这里使用了一个示例 URL，你需要将其替换为真实的登录接口地址
  fetch('https://localhost:8080/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      username: usernameValue,
      password: passwordValue
    })
  })
  .then(response => {
    if (response.ok) {
      // 登录成功，跳转到其他页面
      window.location.href = 'https://localhost:8080/dashboard';
    } else {
      // 登录失败，显示错误信息
      alert('登录失败，请检查用户名和密码是否正确！');
    }
  })
  .catch(error => {
    // 处理网络错误
    console.error('登录请求失败：', error);
    alert('登录请求失败，请检查网络连接！');
  });
});

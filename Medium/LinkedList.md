# Linked List in Golang
Vì Go không có lớp (class) như Java hay C++, chúng ta sử dụng struct kết hợp với con trỏ (Pointer) để tự định nghĩa một Node.

Một Node cơ bản của danh sách liên kết đơn (Singly Linked List) thường có cấu trúc như sau:

```go
type Node struct {
    Val  int       // Chứa giá trị của nút hiện tại (ở đây ví dụ là kiểu số nguyên)
    Next *Node // Con trỏ lưu địa chỉ của nút tiếp theo trong danh sách
}
```

#### 1. Phân biệt Node và Linked List
Để dễ hình dung, bạn hãy tưởng tượng Linked List giống như một đoàn tàu hỏa, còn Node chính là từng toa tàu.

**Node (Toa tàu):** Là một đơn vị/phần tử cấu thành nên danh sách. Mỗi toa tàu (Node) sẽ chở theo hai thứ:

-	value: Hàng hóa nó mang theo (dữ liệu).
-	next: Móc nối dùng để liên kết với toa tàu ngay phía sau nó.

**Linked List (Đoàn tàu):** Là một khái niệm tổng thể để chỉ toàn bộ chuỗi các toa tàu đó cộng lại. Để quản lý cả một đoàn tàu, nhà ga (chương trình của bạn) thường chỉ cần nắm được vị trí của toa đầu tiên (gọi là head).


#### 2. Thể hiện trong Golang
Trong Golang, chúng ta sử dụng struct để định nghĩa các cấu trúc này. Sự tách biệt giữa Node và Linked List được thể hiện rất rõ ràng qua code:
```go
package main

import "fmt"

// 1. Định nghĩa Node (Toa tàu)
type Node struct {
    value int    // Dữ liệu của node (có thể là int, string, struct khác...)
    next  *Node  // Con trỏ trỏ tới Node tiếp theo
}

// 2. Định nghĩa Linked List (Đoàn tàu)
type LinkedList struct {
    head *Node   // Chỉ cần lưu trữ con trỏ trỏ tới Node ĐẦU TIÊN
    size int     // (Tùy chọn) Lưu kích thước của danh sách để dễ quản lý
}
```




#### 4. Tại sao Next phải là một con trỏ (*NodeNode)?
Trong Go, nếu bạn cố tình định nghĩa kiểu dữ liệu lồng chính nó mà không dùng con trỏ như thế này:
```go
type Node struct {
    Val  int
    Next Node // ❌ LỖI BIÊN DỊCH: invalid recursive type Node
}
```


Trình biên dịch của Go sẽ báo lỗi ngay lập tức. Lý do là vì Go cần biết chính xác kích thước vùng nhớ (bao nhiêu bytes) của một struct khi biên dịch. Nếu Node chứa một Node bên trong, và Node bên trong lại chứa một Node khác nữa... thì kích thước của nó sẽ trở thành vô hạn và máy tính không thể cấp phát bộ nhớ.

Khi ta đổi thành con trỏ *Node, Go sẽ hiểu: "À, trường Next này không chứa toàn bộ cái nhà, nó chỉ chứa một mảnh giấy ghi địa chỉ số nhà (địa chỉ vùng nhớ trên RAM) của cái nhà tiếp theo thôi". Kích thước của một con trỏ trên hệ điều hành 64-bit luôn cố định là 8 bytes, giúp Go quản lý bộ nhớ một cách hoàn hảo.


#### Dùng LinkedList để giải quyết các bài toán kinh điển như:

- Đảo ngược danh sách liên kết (Reverse Linked List).

- Trộn hai danh sách liên kết đã sắp xếp (Merge Two Sorted Lists).

- lPhát hiện vòng lặp trong danh sách (Linked List Cycle).
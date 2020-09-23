// Bình có một danh bạ điện thoại dùng để lưu thông tin về các người bạn. Các số điện thoại trong danh bạ của Bình đều là số có 9 chữ số và không được bắt đầu bởi chữ số 0.

// Với hệ điều hành mà Bình đang sử dụng, khi bấm tìm kiếm một số điện thoại nào đó trong danh bạ thì màn hình sẽ hiển thị ra tất cả các số điện thoại mà có các chữ số liên tiếp trùng với số được tìm kiếm. Ví dụ trong danh bạ của Bình có 3 số điện thoại là 100000007, 123456789, 100012345 thì:

// Khi bấm 00 thì màn hình sẽ hiển thị ra 2 số là 100000007 và 100012345.
// Khi bấm 345 thì màn hình sẽ hiển thị ra 2 số là 123456789 và 100012345.
// Khi bấm 07 thì chỉ có 1 số điện thoại được hiển thị ra màn hình là 100000007
// Biết thông tin về các số điện thoại trong danh bạ được lưu dưới dạng mảng các xâu arr, ứng với mỗi số điện thoại trong danh bạ, bạn hãy xác định xâu có độ dài ngắn nhất mà khi Bình tìm kiếm thì màn hình sẽ chỉ hiển thị đúng số điện thoại này. Nếu có nhiều hơn một xâu thỏa mãn thì lấy xâu có giá trị số nguyên nhỏ nhất.

// Kết quả trả về sẽ là mảng các xâu tương ứng với mỗi số điện thoại.

// Ví dụ:

// Với arr = ["100000007", "123456789", "100012345"] thì output là findPhoneNumbers(arr) = ["07", "6", "01"].
// Giải thích:
// Khi bấm "07" thì chỉ có số điện thoại thứ nhất hiện ra.
// Lưu ý: Bạn cũng có thể bấm các số khác như "007", "10000", ... nhưng các số này không có độ dài nhỏ nhất.
// Khi bấm "6" thì chỉ có số điện thoại thứ hai hiện ra.
// Khi bấm "01" thì chỉ có số điện thoại thứ 3 hiện ra.
// Với arr = ["101122334", "230415555"] thì output là findPhoneNumbers(arr) = ["01", "5"].
// Lưu ý: với số điện thoại đầu tiên bạn có thể bấm "10", "11", "22", ... nhưng các số này đều có giá trị nguyên lớn hơn "01" nên kết quả của số đầu tiên sẽ là "01".
// Đầu vào/Đầu ra

// [Giới hạn thời gian chạy] 0.5 giây với C++, 3 giây với Java và C#, 4s với Python, GO và Js.
// [Đầu vào] Array of strings arr
// 1 <= arr.size <= 500
// "100000000" <= arr[i] <= "999999999"
// Đầu vào luôn đảm bảo các số điện thoại trong danh bạ là khác nhau.
// [Đầu ra] Array of strings

package main

func findPhoneNumbers(arr []string) []string {
	newArr := []string{}
	n := len(arr)

	for i := 0; i < n; i++ {
		current := arr[i]
		lenchar := len(current)
		diffs := []string{}

		for j := 1; j <= lenchar; j++ {
			ec := lenchar - j

			for s := 0; s <= ec; s++ {
				e := s + j
				ch := current[s:e]
				isExists := false

				for id := 0; id < n; id++ {
					if id == i {
						continue
					}

					sd := arr[id]
					for si := 0; si <= ec; si++ {
						ed := si + j
						sdc := sd[si:ed]
						if sdc == ch {
							isExists = true
						}
					}

				}
				if isExists == false {
					diffs = append(diffs, ch)
				}
			}
		}

		if len(diffs) == 0 {
			continue
		}

		minLength := 999
		for _, v := range diffs {
			n := len(v)
			if n < minLength {
				minLength = n
			}
		}
		newdiffs := []string{}
		for _, v := range diffs {
			n := len(v)
			if n == minLength {
				newdiffs = append(newdiffs, v)
			}
		}

		min := "9999999"
		for _, v := range newdiffs {
			if v < min {
				min = v
			}
		}

		newArr = append(newArr, min)
	}

	return newArr
}
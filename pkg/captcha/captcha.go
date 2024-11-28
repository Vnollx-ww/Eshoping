package captcha

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"time"
)

const (
	WIDTH          = 60
	HEIGHT         = 50
	CAPTCHA_LENGTH = 4
)

func generateRandomString(length int) string {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	captcha := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		index := rand.Intn(len(characters))
		captcha[i] = characters[index]
	}
	return string(captcha)
}
func GenerateCaptcha(ctx *app.RequestContext) string {
	// 创建图像并绘制背景
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// 生成随机验证码
	captcha := generateRandomString(CAPTCHA_LENGTH)

	// 设置绘制位置和字体
	point := fixed.Point26_6{X: fixed.Int26_6(10 << 6), Y: fixed.Int26_6(35 << 6)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(captcha)

	// 设置响应头
	ctx.Header("Content-Type", "image/png")

	// 将图片编码为 PNG 格式并写入响应
	if err := png.Encode(ctx, img); err != nil {
		// 错误处理
		fmt.Println("Error encoding captcha image:", err)
	}
	// 返回生成的验证码文本
	return captcha
}

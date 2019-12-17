package endeaesrsaecc

import (
	"encoding/base64"
	"fmt"
	"testing"
)

const  (
	eccpri=`
-----BEGIN ECC PRIVATE KEY -----
MHcCAQEEIFcWCOLVJ2VoDzKSZiNXUcVOqhlt2i9/k9urCCgCJ4TaoAoGCCqGSM49
AwEHoUQDQgAEeRcuzZ6fPlixH02gJG5c3laWMxWySeD/JBPL6fbSgj3YPl8x3AYm
bHDrAnpe1BAMZPbAARuojZTAkCDhp7TTkA==
-----END ECC PRIVATE KEY -----
`
	eccpub=`
-----BEGIN ECC PUBLIC KEY -----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEeRcuzZ6fPlixH02gJG5c3laWMxWy
SeD/JBPL6fbSgj3YPl8x3AYmbHDrAnpe1BAMZPbAARuojZTAkCDhp7TTkA==
-----END ECC PUBLIC KEY -----
`
	rsapri=`
-----BEGIN RSA PRIVATE KEY -----
MIIEpQIBAAKCAQEArlyVbgvuddsymapX0kWO+8krzMwukijyhRU/VEahf1K3JqEP
+gMitXjUtkIzLdhRGop34808y0cc1hvVJMsEXmNR3YhNPjOVCyJXjCTqDrG/+IRW
91edyEM1tKmkYJ/w9aHUoVLGHm0kOwSXSqScDdWj+0NVbsWzAQm6wwI/L0TRSH1c
pzcBo0zXLwfBF4LV2OPPasuTUzLai4ACKJ537RxEelijHdGFs/LT5NpsrcfWKkMl
582pW4pqv6POa/uack/mKdgdoDCApfAzks9BI7N8LbWqStVSWFUWxq1DqodDlWDI
iUZMW/OM1d82uhpmvOj30SgWcoHYOHpeqotNlQIDAQABAoIBAQCLdrp/X0PJOR7s
EnhUVBbeBjbmhJrrhZ0WHbyd6DDc6ohceY+R5lgo1xEtBx5wmQmmNQNYTp1F6weB
qpl96HUCGmcszw0Zp6CbW0izbANa5YoreY8mIAqwWDHo45f7QXM2xc6Riue1Bo9o
NW/d4HSCDFQxcdYv2CoptmKQAIgRgaLcOwurZ7FE+lp6Xws1EQC6dDhDcpfatPNX
YGN0yc/enXq67RDezuNszJNt4BOk6cFN01UuQPdHUcllFSBdyAhtwcszeqYuE4zW
jwaBax/kAwbkDjdRAMpeOTwO+objTCSd2HIyeJ5CY96fIVKsWX3whiLaTLd54BZ0
HbYRN1sBAoGBAOUhRPatXOlkrtx4Zfn2kBXb+QNgCEMCac7iXjvkNpL+d7gWWJ1B
5zXWYCmv9xVxVxC9XqBCuvQgvJy0iN/zRmsanAPvyXbxFkRXIAOf685UscCdYSyy
VOD/929GAtvFmUTPZzZjqfC2MaZURqlktl+no+6WV3mT86Dkz2irqiGlAoGBAMLP
HDrj7/NXMYG7hxYCMyqDLGF2mWBi2dDSMeBrDvhJ8Mj8GNAtxCFu9+lFiKxw3UZQ
9IaTIUDwc0dCuiuAsgdxwDmPN/9jAc2da7xNXUxob62o4Sls3G8VgEcjQikU80Ub
gur+yCud7ucs3/smBMkhFnMB0LOnS2MLSqUgvtkxAoGBALgk+sW9IlS4hAfQAzTW
wYmv1fqubTVddSe9qbo8eNe+Bv09iE4qLuWHupUGRG9JPY3Ig4oM1y9oN+1A8lf5
rfhZ1FUdmy4qJ5kY5DPFjL+wNYL1eKlxUOHbFUKqY5W4wqsYfHyrsGIyKsjgJkHx
HNjXY9pVnUuqajw+Z6pZfEu5AoGAG41ZS8bzJ/J8EQIpz+YNwIR+4WX5uVUhw1QH
M9tQabNjd7mX9NGUPLpKG9b2xpTL5ucKPoJOoLWhSEHavM5d34mqCzoDTH5/Qcpy
81XpzSW0LdaFyesYnilnVChbch4BbhO/B2dzfh+/KzkAkK/G239vgmKOVuphUifF
recctvECgYEAkuAbAvlZbk3vA8T2Pi0EdqxtAvoq6ZXLENjuRI4BPGySEiHhKFNM
a9Om+m5NIXAwbbI0ZU5xE0hNCZOF79SFFISzcD1+PHodgAU2dWJwl/Y9npXupkRj
A0DkFCHblTsA4eEAH+eK5x872m+RmrgJZc+tKr8X85g3ZGU0x77sWTg=
-----END RSA PRIVATE KEY -----
`
	rsapub=`
-----BEGIN RSA PUBLIC KEY -----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArlyVbgvuddsymapX0kWO
+8krzMwukijyhRU/VEahf1K3JqEP+gMitXjUtkIzLdhRGop34808y0cc1hvVJMsE
XmNR3YhNPjOVCyJXjCTqDrG/+IRW91edyEM1tKmkYJ/w9aHUoVLGHm0kOwSXSqSc
DdWj+0NVbsWzAQm6wwI/L0TRSH1cpzcBo0zXLwfBF4LV2OPPasuTUzLai4ACKJ53
7RxEelijHdGFs/LT5NpsrcfWKkMl582pW4pqv6POa/uack/mKdgdoDCApfAzks9B
I7N8LbWqStVSWFUWxq1DqodDlWDIiUZMW/OM1d82uhpmvOj30SgWcoHYOHpeqotN
lQIDAQAB
-----END RSA PUBLIC KEY -----
`
	aes=`f7e8b819l0ad0ccf9a9g8fc5e8c4765q`

	text=`
The Day You Went Away
　　歌手：M2m 专辑：Shades Of Purple
　　Well I wonder could it be
　　When I was dreaming about you baby
　　You were dreaming of me
　　Call me crazy
　　Call me blind
　　To still be suffering is stupid after all of this time
　　Did I lose my love to someone better
　　And does she love you like I do
　　I do, you know I really really do
　　Well hey
　　So much I need to say
　　Been lonely since the day
　　The day you went away
　　So sad but true
　　For me there's only you
　　Been crying since the day
　　the day you went away
　　I remember date and time
　　September twenty second
　　Sunday twenty five after nine
　　In the doorway with your case
　　No longer shouting at each other
　　There were tears on our faces
　　And we were letting go of something special
　　Something we'll never have again
　　I know, I guess I really really know
　　Why do we never know what we've got till it's gone
　　How could I carry on the day you went away
　　Cause I've been missing you so much I have to say
　　just one last dance
　　歌手：natural 专辑：it's only natural
　　Written by: Rob Tyger & Kay Denar
　　Just one last dance ...oh baby...
　　just one last dance
　　We meet in the night in the Spanish caf?
　　I look in your eyes just don't know
　　what to say
　　It feels like I'm drowning in salty water
　　A few hours left till the sun's gonna rise
　　tomorrow will come an it's time to realize
　　our love has finished forever
　　how I wish to come with you
　　(wish to come with you)
　　how I wish we make it through
　　(Chorus)
　　Just one last dance
　　before we say goodbye
　　when we sway and turn round and
　　round and round
　　it's like the first time
　　Just one more chance
　　hold me tight and keep me warm
　　cause the night is getting cold
　　and I don't know where I belong
　　Just one last dance
　　The wine and the lights and the
　　Spanish guitar
　　I'll never forget how romantic
　　they are
　　but I know, tomorrow I'll lose
　　the one I love
　　There's no way to come with you
　　it's the only thing to do
　　(Chorus 3x)
　　(until fade)
　　Just one last dance,
　　just one more chance,
　　just one last dance
　　life is cool
　　I never really tried to be positive
　　我从来没有努力去拥有积极的态度
　　I’m too damn busy being negative
　　我一直纵容自己消极的沉沦
　　So focused on what I get
　　如此看重得失
　　I never understand what it means to live
　　却从未明白过生活的实质
　　You know we all love to just complain
　　我们都喜欢抱怨
　　But maybe we should try to rearrange
　　但是也许我们应该重新整理心思
　　There’s always someone who’s got it worse than you
　　世界上总有比我们更糟糕的人或事
　　My life is so cool, my life is so cool
　　我的生活如此冷酷
　　Oh yeah, from a different point of view
　　当你从另外的角度去审视
　　My life is so cool, my life is so cool
　　我的生活如此冷酷
　　Oh yeah, from a different point of view
　　当你从另外的角度去审视
　　We’re all so busy trying to get ahead
　　我们都忙于不落人后,忙于提高自己的素质
　　Got a pillow of fears when we go to bed
　　甚至在睡觉时都担心恐惧钻进自己的被子
　　We’re never satisfied
　　我们从不满足
　　The grass is greener on the other side
　　总期待着更高更好的位置
　　So distracted with our jealousy
　　嫉妒总在前头指使
　　Forget it’s in our hands to stop the agony
　　却忘记了我们自己就可以使这一切停止
　　Will you ever be content on your side of the fence?
　　你难道就不能真正满足于上帝给我们的恩赐?
　　Maybe you’re the guy who needs a second chance
　　也许你需要人生的第二次选择
　　Maybe you’re the girl who’s never asked to dance
　　也许你从来未被大家重视
　　Maybe you’re a lonely soul
　　也许你有孤独的灵魂
　　A single mother scared and all alone
　　如同单身母亲心里恐惧的对生活的未知
　　Gotta remember we live what we choose
　　你要记住我们的生活是我们自己选择的方式
　　It’s not what you say, it’s what you do
　　少说多做永远是正确的path
　　And the life you want is the life you have to make
　　想过什么样的生活,就一定要为它奋斗不止
　　she
　　歌手：groove coverage 专辑：greatest hits
　　[ar:Groove Coverage]
　　[al:greatest hits]
　　She-Groove Coverage
　　She hangs out every day near by the beach
　　Havin' a HEINEKEN fallin' asleep
　　She looks so sexy when she's walking the sand
　　Nobody ever put a ring on her hand
　　Swim to the oceanshore fish in the sea
　　She is the story the story is she
　　She sings to the moon and the stars in the sky
　　Shining from high above you shouldn't ask why
　　She is the one that you never forget
　　She is the heaven-sent angel you met
　　Oh, she must be the reason why
　　God made a girl
　　She is so pretty all over the world
　　She puts the rhythm,
　　The beat in the drum
　　She comes in the morning
　　And the evening she's gone
　　Every little hour every second you live
　　Trust in eternity that's what she gives
　　She looks like Marilyn,
　　Walks like Suzanne
　　She talks like Monica and Marianne
　　She wins in everything that she might do
　　And she will respect you forever just you
　　She is the one that you never forget
　　She is the heaven-sent angel you met
　　Oh, she must be the reason why
　　God made a girl
　　She is so pretty all over the world
　　She is so pretty all over the world
　　She is so pretty
　　She is like you and me
　　Like them like we
　　She is in you and me
　　She is the one that you never forget
　　She is the heaven-sent angel you met
　　Oh, she must be the reason why
　　God made a girl
　　She is so pretty all over the world
　　(She is the one) She is the one
　　(That you never forget) That you never forget
　　She is the heaven-sent angel you met
　　She must be the reason why
　　God made a girl
　　She is so pretty all over the world (oh...)
　　Na na na na na ….
　　God Is A Girl
　　歌手：Groove Coverage 专辑：Booom 2003 The First
　　Remembering me,
　　Discover and see
　　All over the world,
　　She's known as a girl
　　To those who a free,
　　The mind shall be key
　　Forgotten as the past
　　'Cause history will last
　　God is a girl,
　　Wherever you are,
　　Do you believe it, can you recieve it?
　　God is a girl,
　　Whatever you say,
　　Do you believe it, can you recieve it?
　　God is a girl,
　　However you live,
　　Do you believe it, can you recieve it?
　　God is a girl,
　　She's only a girl,
　　Do you believe it, can you recieve it?
　　She wants to shine,
　　Forever in time,
　　She is so driven, she's always mine
　　Cleanly and free,
　　She wants you to be
　　A part of the future,
　　A girl like me
　　There is a sky,
　　Illuminating us, someone is out there
　　That we truly trust
　　There is a rainbow for you and me
　　A beautiful sunrise eternally
　　God is a girl
　　Wherever you are,
　　Do you believe it, can you recieve it?
　　God is a girl
　　Whatever you say,
　　Do you believe it, can you recieve it?
　　God is a girl
　　However you live,
　　Do you believe it, can you recieve it?
　　God is a girl
　　She's only a girl,
　　Do you believe it, can you recieve it?`
)

func TestPubEncrypt(t *testing.T) {
	src:=[]byte(text)
	byt,err:=PubEncrypt(src,[]byte(eccpub),[]byte(rsapub),[]byte(aes))
	fmt.Println(base64.StdEncoding.EncodeToString(byt),err)
	byt,err=PriDecrypt(byt,[]byte(eccpri),[]byte(rsapri),[]byte(aes))
	fmt.Println(err)
}

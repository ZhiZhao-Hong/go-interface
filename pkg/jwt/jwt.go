package jwt

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go-interface/pkg/log"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
	"time"
)

type Conf struct {
	JwtKey     string
	JwtExpHour int64
}

var (
	ERR_INVALID_TOKEN = fmt.Errorf("无效的token")
)

var (
	conf *Conf
)

func Init(JwtExpHour int64) {
	conf = &Conf{
		JwtKey:     "wondfo@2020",
		JwtExpHour: JwtExpHour,
	}
}

// GenerateEncodeToken return jwt token and error
func GenerateEncodeToken(m map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	for k, v := range m {
		claims[k] = v
	}
	// set expire time
	ttl := time.Duration(conf.JwtExpHour) * time.Hour
	claims["exp"] = time.Now().UTC().Add(ttl).Unix()
	// SignedString must be a []byte
	t, err := token.SignedString([]byte(conf.JwtKey))
	return t, err
}

type JwtMid struct {
	log *zap.SugaredLogger
}

func NewJwtMid() *JwtMid {
	return &JwtMid{log: log.Get()}
}

func (m *JwtMid) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ah := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")
		if len(ah) < 2 {
			m.log.Infof("ah len < 2, content: [%s]", ah)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if strings.LastIndex(ah[1], ",") == len(ah[1])-1 {
			m.log.Errorf("token以逗号结尾 [token=%v]", ah[1])
			ah[1] = strings.TrimRight(ah[1], ",")
		}

		_, err := m.decodeToken(ah[1])
		if err != nil {
			m.log.Infof("decode token error: [%s] [token=%v]", err.Error(), ah[1])
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func (m *JwtMid) getHashSum(token string) string {
	h := md5.New()
	_, _ = io.WriteString(h, token)
	return string(h.Sum(nil))
}

func (m *JwtMid) decodeToken(tokenString string) (ms map[string]interface{}, err error) {
	ts := strings.Split(tokenString, ".")
	if len(ts) != 3 || ts[1] == "" {
		m.log.Info("token格式错误 ", ts)
		err = ERR_INVALID_TOKEN
		return
	}
	payload, err := jwt.DecodeSegment(ts[1])
	if err != nil {
		m.log.Info("DecodeSegment fail ", string(payload), err)
		err = ERR_INVALID_TOKEN
		return
	}
	ms = make(map[string]interface{})
	if err = json.Unmarshal(payload, &ms); err != nil {
		m.log.Info("json Unmarshal fail ", string(payload), ms, err)
		err = ERR_INVALID_TOKEN
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.JwtKey), nil
	})
	if err != nil {
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && !(token.Valid) {
		err = fmt.Errorf("token无效或已过期")
		return
	}
	return
}

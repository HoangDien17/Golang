package middleware

import (
	db "employee/connect"
	"employee/models"
	"employee/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserByEmail(resource *db.Resource, email interface{}) (*models.User, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()
	var user models.User
	err := resource.DB.Collection("user").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return &user, err
}

func Authorization(resource *db.Resource, roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(500, gin.H{"status": false, "error": "Missing authorization token!"})
			c.Abort()
			return
		}
		mapClaim := jwt.MapClaims{}
		decode, error := jwt.ParseWithClaims(token, &mapClaim, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if error != nil || !decode.Valid {
			c.JSON(401, gin.H{"status": false, "error": "Unauthorized!"})
			c.Abort()
			return
		}

		user, err := GetUserByEmail(resource, mapClaim["email"])
		if err != nil {
			c.JSON(401, gin.H{"status": false, "error": "Unauthorized!"})
			c.Abort()
			return
		}

		if !utils.RoleContains(user.Roles, roles) {
			c.JSON(403, gin.H{"status": false, "error": "Forbidden!"})
			c.Abort()
			return
		}

		c.Set("user", map[string]interface{}{
			"id":    user.Id,
			"email": user.Email,
		})
		c.Next()
	}
}

package profile

import (
	"jeu-de-bourse/internal/env"
	"jeu-de-bourse/internal/models"

	"github.com/gin-gonic/gin"
)

type ProfileInfos struct {
	Username        string  `json:"username"`
	Valorisation    int     `json:"valorisation"`
	VarValorisation float64 `json:"var_valorisation"`
	DispoComptant   int     `json:"comptant"`
	DispoSRD        int     `json:"srd"`
}

func GetProfileInfos(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	// Get infos
	infos := ProfileInfos{
		Username: user.Username,
	}

	valorisation, err := user.Valorisation()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"status":  "Erreur interne",
		})
		return
	}
	infos.Valorisation = valorisation
	infos.VarValorisation = (float64(valorisation) / float64(env.GlobalConfig.StartComptant) * 100.0) - 100.0

	dispComptant, err := user.DisponibleComptant()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"status":  "Erreur interne",
		})
		return
	}

	infos.DispoComptant = dispComptant

	dispSRD, err := user.DisponibleSRD()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"status":  "Erreur interne",
		})
		return
	}

	infos.DispoSRD = dispSRD

	c.JSON(200, gin.H{
		"success": true,
		"infos":   infos,
	})
}

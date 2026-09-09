from django.db import models
 #======================================================================
# 
# Encapsulates data for model AdFormat
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdFormat Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AdFormat(Enum):   # A subclass of Enum
	Banner = 'Banner'
	Video = 'Video'
	Native = 'Native'
	Audio = 'Audio'
	Interstitial = 'Interstitial'
	RichMedia = 'RichMedia'
	SearchText = 'SearchText'
	SocialPost = 'SocialPost'
	CTVVideo = 'CTVVideo'

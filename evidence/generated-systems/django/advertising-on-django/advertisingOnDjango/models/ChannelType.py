from django.db import models
 #======================================================================
# 
# Encapsulates data for model ChannelType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ChannelType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ChannelType(Enum):   # A subclass of Enum
	Programmatic = 'Programmatic'
	Direct = 'Direct'
	Search = 'Search'
	Social = 'Social'
	Email = 'Email'
	Affiliate = 'Affiliate'
	DOOH = 'DOOH'

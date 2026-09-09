from django.db import models
 #======================================================================
# 
# Encapsulates data for model CaseOrigin
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CaseOrigin Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CaseOrigin(Enum):   # A subclass of Enum
	Email = 'Email'
	Phone = 'Phone'
	Web = 'Web'
	Chat = 'Chat'
	Social = 'Social'
	Community = 'Community'

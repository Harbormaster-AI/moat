from django.db import models
 #======================================================================
# 
# Encapsulates data for model NotificationChannel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NotificationChannel Declaration (enumerated type)
#======================================================================
from enum import Enum 
class NotificationChannel(Enum):   # A subclass of Enum
	Email = 'Email'
	SMS = 'SMS'
	Webhook = 'Webhook'
	Chat = 'Chat'

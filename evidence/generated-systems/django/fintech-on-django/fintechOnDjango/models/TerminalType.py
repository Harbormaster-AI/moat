from django.db import models
 #======================================================================
# 
# Encapsulates data for model TerminalType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerminalType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TerminalType(Enum):   # A subclass of Enum
	POS = 'POS'
	mPOS = 'mPOS'
	ECommerce = 'ECommerce'

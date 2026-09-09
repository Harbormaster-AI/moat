from django.db import models
 #======================================================================
# 
# Encapsulates data for model AccessLevel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccessLevel Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AccessLevel(Enum):   # A subclass of Enum
	View = 'View'
	Query = 'Query'
	Modify = 'Modify'
	Admin = 'Admin'

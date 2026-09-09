from django.db import models
 #======================================================================
# 
# Encapsulates data for model ConsentType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConsentType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ConsentType(Enum):   # A subclass of Enum
	Marketing = 'Marketing'
	Profiling = 'Profiling'
	Cookies = 'Cookies'
	Location = 'Location'
	Biometric = 'Biometric'

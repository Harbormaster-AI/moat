from django.db import models
 #======================================================================
# 
# Encapsulates data for model SoftwareUpdateType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SoftwareUpdateType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SoftwareUpdateType(Enum):   # A subclass of Enum
	SecurityPatch = 'SecurityPatch'
	FeatureUpdate = 'FeatureUpdate'
	BugFix = 'BugFix'
	FirmwareUpgrade = 'FirmwareUpgrade'

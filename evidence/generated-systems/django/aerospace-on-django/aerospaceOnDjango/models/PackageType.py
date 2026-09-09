from django.db import models
 #======================================================================
# 
# Encapsulates data for model PackageType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PackageType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PackageType(Enum):   # A subclass of Enum
	PerformancePack = 'PerformancePack'
	CabinPack = 'CabinPack'
	ConnectivityPack = 'ConnectivityPack'
	CompliancePack = 'CompliancePack'

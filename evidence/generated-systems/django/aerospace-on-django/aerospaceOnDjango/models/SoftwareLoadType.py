from django.db import models
 #======================================================================
# 
# Encapsulates data for model SoftwareLoadType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SoftwareLoadType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SoftwareLoadType(Enum):   # A subclass of Enum
	FlightDeckSoftware = 'FlightDeckSoftware'
	MaintenanceTools = 'MaintenanceTools'
	CabinIFE = 'CabinIFE'
	ConnectivityModem = 'ConnectivityModem'

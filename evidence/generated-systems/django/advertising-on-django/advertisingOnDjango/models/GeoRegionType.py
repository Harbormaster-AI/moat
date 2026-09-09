from django.db import models
 #======================================================================
# 
# Encapsulates data for model GeoRegionType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GeoRegionType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class GeoRegionType(Enum):   # A subclass of Enum
	Country = 'Country'
	State = 'State'
	Province = 'Province'
	City = 'City'
	DMA = 'DMA'
	PostalCode = 'PostalCode'

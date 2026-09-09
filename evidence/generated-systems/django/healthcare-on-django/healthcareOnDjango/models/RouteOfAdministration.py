from django.db import models
 #======================================================================
# 
# Encapsulates data for model RouteOfAdministration
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RouteOfAdministration Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RouteOfAdministration(Enum):   # A subclass of Enum
	Oral = 'Oral'
	Intravenous = 'Intravenous'
	Subcutaneous = 'Subcutaneous'
	Intramuscular = 'Intramuscular'
	Topical = 'Topical'
	Inhalation = 'Inhalation'

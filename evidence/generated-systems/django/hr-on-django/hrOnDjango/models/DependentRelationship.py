from django.db import models
 #======================================================================
# 
# Encapsulates data for model DependentRelationship
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DependentRelationship Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DependentRelationship(Enum):   # A subclass of Enum
	Spouse = 'Spouse'
	DomesticPartner = 'DomesticPartner'
	Child = 'Child'
	Other = 'Other'

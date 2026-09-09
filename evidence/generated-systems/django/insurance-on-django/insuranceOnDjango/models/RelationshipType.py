from django.db import models
 #======================================================================
# 
# Encapsulates data for model RelationshipType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RelationshipType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RelationshipType(Enum):   # A subclass of Enum
	Spouse = 'Spouse'
	Child = 'Child'
	Parent = 'Parent'
	Sibling = 'Sibling'
	BusinessPartner = 'BusinessPartner'
	Estate = 'Estate'
	Trust = 'Trust'
	Other = 'Other'

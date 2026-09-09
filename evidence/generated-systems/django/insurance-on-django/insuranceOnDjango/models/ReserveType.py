from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReserveType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReserveType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReserveType(Enum):   # A subclass of Enum
	Indemnity = 'Indemnity'
	Expense = 'Expense'
	Legal = 'Legal'
	Medical = 'Medical'

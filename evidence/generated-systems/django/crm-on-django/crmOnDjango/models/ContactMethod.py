from django.db import models
 #======================================================================
# 
# Encapsulates data for model ContactMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContactMethod Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ContactMethod(Enum):   # A subclass of Enum
	Email = 'Email'
	Phone = 'Phone'
	Mobile = 'Mobile'
	SMS = 'SMS'
	InPerson = 'InPerson'
	Web = 'Web'

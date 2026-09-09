from django.db import models
 #======================================================================
# 
# Encapsulates data for model DispositionAction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DispositionAction Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DispositionAction(Enum):   # A subclass of Enum
	Destroy = 'Destroy'
	TransferToArchive = 'TransferToArchive'
	Review = 'Review'
	SecureDelete = 'SecureDelete'
	ReturnToOwner = 'ReturnToOwner'

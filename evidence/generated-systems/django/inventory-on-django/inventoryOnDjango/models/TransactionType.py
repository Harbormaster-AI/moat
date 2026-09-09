from django.db import models
 #======================================================================
# 
# Encapsulates data for model TransactionType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransactionType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TransactionType(Enum):   # A subclass of Enum
	Receipt = 'Receipt'
	Issue = 'Issue'
	AdjustmentIncrease = 'AdjustmentIncrease'
	AdjustmentDecrease = 'AdjustmentDecrease'
	Reclassification = 'Reclassification'
	TransferOut = 'TransferOut'
	TransferIn = 'TransferIn'
	CountIncrease = 'CountIncrease'
	CountDecrease = 'CountDecrease'
	Putaway = 'Putaway'
	Pick = 'Pick'

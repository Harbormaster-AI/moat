from django.db import models
 #======================================================================
# 
# Encapsulates data for model ChargebackStage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ChargebackStage Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ChargebackStage(Enum):   # A subclass of Enum
	FirstChargeback = 'FirstChargeback'
	SecondChargeback = 'SecondChargeback'
	Arbitration = 'Arbitration'

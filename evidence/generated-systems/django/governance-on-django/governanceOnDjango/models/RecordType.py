from django.db import models
 #======================================================================
# 
# Encapsulates data for model RecordType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RecordType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RecordType(Enum):   # A subclass of Enum
	PolicyRecord = 'PolicyRecord'
	ContractRecord = 'ContractRecord'
	FinancialRecord = 'FinancialRecord'
	HRRecord = 'HRRecord'
	CustomerRecord = 'CustomerRecord'
	TechnicalRecord = 'TechnicalRecord'
	AuditRecord = 'AuditRecord'
	LegalRecord = 'LegalRecord'

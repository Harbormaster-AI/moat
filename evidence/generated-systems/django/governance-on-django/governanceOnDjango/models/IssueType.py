from django.db import models
 #======================================================================
# 
# Encapsulates data for model IssueType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class IssueType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class IssueType(Enum):   # A subclass of Enum
	ControlDeficiency = 'ControlDeficiency'
	ProcessGap = 'ProcessGap'
	ComplianceBreach = 'ComplianceBreach'
	SecurityIncident = 'SecurityIncident'
	DataQualityIssue = 'DataQualityIssue'
	ThirdPartyIssue = 'ThirdPartyIssue'

from django.db import models
 #======================================================================
# 
# Encapsulates data for model MetricType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MetricType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class MetricType(Enum):   # A subclass of Enum
	Impressions = 'Impressions'
	ViewableImpressions = 'ViewableImpressions'
	Clicks = 'Clicks'
	CTR = 'CTR'
	Reach = 'Reach'
	Frequency = 'Frequency'
	VideoStarts = 'VideoStarts'
	VideoCompletions = 'VideoCompletions'
	AvgViewTime = 'AvgViewTime'
	Conversions = 'Conversions'
	ViewThroughConversions = 'ViewThroughConversions'
	Spend = 'Spend'
	CPM = 'CPM'
	CPC = 'CPC'
	CPA = 'CPA'

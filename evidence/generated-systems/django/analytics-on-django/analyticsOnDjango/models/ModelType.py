from django.db import models
 #======================================================================
# 
# Encapsulates data for model ModelType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ModelType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ModelType(Enum):   # A subclass of Enum
	Classification = 'Classification'
	Regression = 'Regression'
	Clustering = 'Clustering'
	Forecasting = 'Forecasting'
	Ranking = 'Ranking'
	NLP = 'NLP'
	ComputerVision = 'ComputerVision'

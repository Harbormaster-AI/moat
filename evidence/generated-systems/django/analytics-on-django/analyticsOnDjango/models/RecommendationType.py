from django.db import models
 #======================================================================
# 
# Encapsulates data for model RecommendationType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RecommendationType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RecommendationType(Enum):   # A subclass of Enum
	Personalized = 'Personalized'
	Trending = 'Trending'
	SimilarItems = 'SimilarItems'
	FrequentlyBoughtTogether = 'FrequentlyBoughtTogether'
	ContentBased = 'ContentBased'

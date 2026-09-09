from django.db import models
 #======================================================================
# 
# Encapsulates data for model OptionCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OptionCategory Declaration (enumerated type)
#======================================================================
from enum import Enum 
class OptionCategory(Enum):   # A subclass of Enum
	Cabin = 'Cabin'
	Connectivity = 'Connectivity'
	Safety = 'Safety'
	Performance = 'Performance'
	Paint = 'Paint'
	FlightDeck = 'FlightDeck'

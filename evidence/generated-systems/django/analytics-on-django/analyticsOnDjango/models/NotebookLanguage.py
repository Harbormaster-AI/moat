from django.db import models
 #======================================================================
# 
# Encapsulates data for model NotebookLanguage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NotebookLanguage Declaration (enumerated type)
#======================================================================
from enum import Enum 
class NotebookLanguage(Enum):   # A subclass of Enum
	Python = 'Python'
	R = 'R'
	SQL = 'SQL'
	Julia = 'Julia'

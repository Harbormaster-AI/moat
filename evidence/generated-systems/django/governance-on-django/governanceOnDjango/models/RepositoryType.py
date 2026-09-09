from django.db import models
 #======================================================================
# 
# Encapsulates data for model RepositoryType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RepositoryType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RepositoryType(Enum):   # A subclass of Enum
	DocumentManagement = 'DocumentManagement'
	RecordsArchive = 'RecordsArchive'
	EmailArchive = 'EmailArchive'
	FileShare = 'FileShare'
	ContentServices = 'ContentServices'
	DataLake = 'DataLake'

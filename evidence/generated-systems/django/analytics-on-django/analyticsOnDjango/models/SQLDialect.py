from django.db import models
 #======================================================================
# 
# Encapsulates data for model SQLDialect
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SQLDialect Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SQLDialect(Enum):   # A subclass of Enum
	ANSI = 'ANSI'
	Postgres = 'Postgres'
	MySQL = 'MySQL'
	SQLServer = 'SQLServer'
	Oracle = 'Oracle'
	SparkSQL = 'SparkSQL'
	BigQuery = 'BigQuery'

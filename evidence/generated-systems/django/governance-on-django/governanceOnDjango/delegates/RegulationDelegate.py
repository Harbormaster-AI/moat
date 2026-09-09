from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Regulation import Regulation
from governanceOnDjango.models.Obligation import Obligation
from governanceOnDjango.models.ComplianceProgram import ComplianceProgram
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Regulation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RegulationDelegate Declaration
#======================================================================
class RegulationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, regulationId ):
		try:	
			regulation = Regulation.objects.filter(id=regulationId)
			return regulation.first();
		except Regulation.DoesNotExist:
			raise ProcessingError("Regulation with id " + str(regulationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, regulation):
		for model in serializers.deserialize("json", regulation):
			model.save()
			return model;

	def create(self, regulation):
		regulation.save()
		return regulation;

	def saveFromJson(self, regulation):
		for model in serializers.deserialize("json", regulation):
			model.save()
			return regulation;
	
	def save(self, regulation):
		regulation.save()
		return regulation;
	
	def delete(self, regulationId ):
		errMsg = "Failed to delete Regulation from db using id " + str(regulationId)
		
		try:
			regulation = Regulation.objects.get(id=regulationId)
			regulation.delete()
			return True
		except Regulation.DoesNotExist:
			raise ProcessingError("Regulation with id " + str(regulationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Regulation.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Regulation from db")
		except Exception:
			return None;
		
	def addObligations( self, regulationId, obligationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

		errMsg = "Failed to add elements " + str(obligationsIds) + " for Obligations on Regulation"

		try:
			# get the Regulation
			regulation = self.get( regulationId ).first()
				
			# split on a comma with no spaces
			idList = obligationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Obligation		
				obligation = ObligationDelegate().get(id).first();	
				# add the Obligation
				regulation.obligations.add(obligation)
				
			# save it		
			regulation.save()
			
			# reload and return the appropriate version
			return self.get( regulationId );
		except Regulation.DoesNotExist:
			raise ProcessingError(errMsg + " : Regulation with id " + str(regulationId) + " does not exist.")
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeObligations( self, regulationId, obligationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

		errMsg = "Failed to remove elements " + str(obligationsIds) + " for Obligations on Regulation"

		try:
			# get the Regulation
			regulation = self.get( regulationId ).first()
				
			# split on a comma with no spaces
			idList = obligationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Obligation		
				obligation = ObligationDelegate().get(id).first();	
				# add the Obligation
				regulation.obligations.remove(obligation)
				
			# save it		
			regulation.save()
			
			# reload and return the appropriate version
			return self.get( regulationId );
		except Regulation.DoesNotExist:
			raise ProcessingError(errMsg + " : Regulation with id " + str(regulationId) + " does not exist.")
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCompliancePrograms( self, regulationId, complianceProgramsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ComplianceProgramDelegate import ComplianceProgramDelegate

		errMsg = "Failed to add elements " + str(complianceProgramsIds) + " for CompliancePrograms on Regulation"

		try:
			# get the Regulation
			regulation = self.get( regulationId ).first()
				
			# split on a comma with no spaces
			idList = complianceProgramsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ComplianceProgram		
				complianceProgram = ComplianceProgramDelegate().get(id).first();	
				# add the ComplianceProgram
				regulation.compliancePrograms.add(complianceProgram)
				
			# save it		
			regulation.save()
			
			# reload and return the appropriate version
			return self.get( regulationId );
		except Regulation.DoesNotExist:
			raise ProcessingError(errMsg + " : Regulation with id " + str(regulationId) + " does not exist.")
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCompliancePrograms( self, regulationId, complianceProgramsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ComplianceProgramDelegate import ComplianceProgramDelegate

		errMsg = "Failed to remove elements " + str(complianceProgramsIds) + " for CompliancePrograms on Regulation"

		try:
			# get the Regulation
			regulation = self.get( regulationId ).first()
				
			# split on a comma with no spaces
			idList = complianceProgramsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ComplianceProgram		
				complianceProgram = ComplianceProgramDelegate().get(id).first();	
				# add the ComplianceProgram
				regulation.compliancePrograms.remove(complianceProgram)
				
			# save it		
			regulation.save()
			
			# reload and return the appropriate version
			return self.get( regulationId );
		except Regulation.DoesNotExist:
			raise ProcessingError(errMsg + " : Regulation with id " + str(regulationId) + " does not exist.")
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		

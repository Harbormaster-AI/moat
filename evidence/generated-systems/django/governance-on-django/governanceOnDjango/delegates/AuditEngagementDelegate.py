from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.AuditEngagement import AuditEngagement
from governanceOnDjango.models.AuditProgram import AuditProgram
from governanceOnDjango.models.BusinessUnit import BusinessUnit
from governanceOnDjango.models.ControlTest_ import ControlTest_
from governanceOnDjango.models.AuditWorkpaper import AuditWorkpaper
from governanceOnDjango.models.AuditFinding import AuditFinding
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AuditEngagement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditEngagementDelegate Declaration
#======================================================================
class AuditEngagementDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, auditEngagementId ):
		try:	
			auditEngagement = AuditEngagement.objects.filter(id=auditEngagementId)
			return auditEngagement.first();
		except AuditEngagement.DoesNotExist:
			raise ProcessingError("AuditEngagement with id " + str(auditEngagementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, auditEngagement):
		for model in serializers.deserialize("json", auditEngagement):
			model.save()
			return model;

	def create(self, auditEngagement):
		auditEngagement.save()
		return auditEngagement;

	def saveFromJson(self, auditEngagement):
		for model in serializers.deserialize("json", auditEngagement):
			model.save()
			return auditEngagement;
	
	def save(self, auditEngagement):
		auditEngagement.save()
		return auditEngagement;
	
	def delete(self, auditEngagementId ):
		errMsg = "Failed to delete AuditEngagement from db using id " + str(auditEngagementId)
		
		try:
			auditEngagement = AuditEngagement.objects.get(id=auditEngagementId)
			auditEngagement.delete()
			return True
		except AuditEngagement.DoesNotExist:
			raise ProcessingError("AuditEngagement with id " + str(auditEngagementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AuditEngagement.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AuditEngagement from db")
		except Exception:
			return None;
		
	def assignAuditProgram( self, auditEngagementId, auditProgramId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditProgramDelegate import AuditProgramDelegate

		errMsg = "Failed to assign element " + str(auditProgramId) + " for AuditProgram on AuditEngagement"

		try:
			# get the AuditEngagement from db
			auditEngagement = self.get( auditEngagementId ).first()	
			
			# get the AuditProgram from db
			auditProgram = AuditProgramDelegate().get(auditProgramId).first();
			
			# assign the AuditProgram		
			auditEngagement.auditProgram = auditProgram
			
			#save it
			auditEngagement.save()

			# reload and return the appropriate version					
			return self.get( auditEngagementId );
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(auditEngagementId) + " does not exist.")
		except AuditProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditProgram with id " + str(auditProgramId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAuditProgram( self, auditEngagementId ):
		errMsg = "Failed to unassign element " + str(auditProgramId) + " for AuditProgram on AuditEngagement"

		try:
			# get the AuditEngagement from db
			auditEngagement = self.get( auditEngagementId ).first()	
			
			# assign to None for unassignment
			auditEngagement.auditProgram = None			

			#save it
			auditEngagement.save()

			# reload and return the appropriate version					
			return self.get( auditEngagementId );
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(auditEngagementId) + " does not exist.")
		except Exception:
			return None;
		
	def addBusinessUnits( self, auditEngagementId, businessUnitsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.BusinessUnitDelegate import BusinessUnitDelegate

		errMsg = "Failed to add elements " + str(businessUnitsIds) + " for BusinessUnits on AuditEngagement"

		try:
			# get the AuditEngagement
			auditEngagement = self.get( auditEngagementId ).first()
				
			# split on a comma with no spaces
			idList = businessUnitsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BusinessUnit		
				businessUnit = BusinessUnitDelegate().get(id).first();	
				# add the BusinessUnit
				auditEngagement.businessUnits.add(businessUnit)
				
			# save it		
			auditEngagement.save()
			
			# reload and return the appropriate version
			return self.get( auditEngagementId );
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(auditEngagementId) + " does not exist.")
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBusinessUnits( self, auditEngagementId, businessUnitsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.BusinessUnitDelegate import BusinessUnitDelegate

		errMsg = "Failed to remove elements " + str(businessUnitsIds) + " for BusinessUnits on AuditEngagement"

		try:
			# get the AuditEngagement
			auditEngagement = self.get( auditEngagementId ).first()
				
			# split on a comma with no spaces
			idList = businessUnitsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BusinessUnit		
				businessUnit = BusinessUnitDelegate().get(id).first();	
				# add the BusinessUnit
				auditEngagement.businessUnits.remove(businessUnit)
				
			# save it		
			auditEngagement.save()
			
			# reload and return the appropriate version
			return self.get( auditEngagementId );
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(auditEngagementId) + " does not exist.")
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addControlTests( self, auditEngagementId, controlTestsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlTest_Delegate import ControlTest_Delegate

		errMsg = "Failed to add elements " + str(controlTestsIds) + " for ControlTests on AuditEngagement"

		try:
			# get the AuditEngagement
			auditEngagement = self.get( auditEngagementId ).first()
				
			# split on a comma with no spaces
			idList = controlTestsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ControlTest_		
				controlTest_ = ControlTest_Delegate().get(id).first();	
				# add the ControlTest_
				auditEngagement.controlTests.add(controlTest_)
				
			# save it		
			auditEngagement.save()
			
			# reload and return the appropriate version
			return self.get( auditEngagementId );
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(auditEngagementId) + " does not exist.")
		except ControlTest_.DoesNotExist:
			raise ProcessingError(errMsg + " : ControlTest_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeControlTests( self, auditEngagementId, controlTestsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlTest_Delegate import ControlTest_Delegate

		errMsg = "Failed to remove elements " + str(controlTestsIds) + " for ControlTests on AuditEngagement"

		try:
			# get the AuditEngagement
			auditEngagement = self.get( auditEngagementId ).first()
				
			# split on a comma with no spaces
			idList = controlTestsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ControlTest_		
				controlTest_ = ControlTest_Delegate().get(id).first();	
				# add the ControlTest_
				auditEngagement.controlTests.remove(controlTest_)
				
			# save it		
			auditEngagement.save()
			
			# reload and return the appropriate version
			return self.get( auditEngagementId );
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(auditEngagementId) + " does not exist.")
		except ControlTest_.DoesNotExist:
			raise ProcessingError(errMsg + " : ControlTest_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addWorkpapers( self, auditEngagementId, workpapersIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditWorkpaperDelegate import AuditWorkpaperDelegate

		errMsg = "Failed to add elements " + str(workpapersIds) + " for Workpapers on AuditEngagement"

		try:
			# get the AuditEngagement
			auditEngagement = self.get( auditEngagementId ).first()
				
			# split on a comma with no spaces
			idList = workpapersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AuditWorkpaper		
				auditWorkpaper = AuditWorkpaperDelegate().get(id).first();	
				# add the AuditWorkpaper
				auditEngagement.workpapers.add(auditWorkpaper)
				
			# save it		
			auditEngagement.save()
			
			# reload and return the appropriate version
			return self.get( auditEngagementId );
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(auditEngagementId) + " does not exist.")
		except AuditWorkpaper.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditWorkpaper does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWorkpapers( self, auditEngagementId, workpapersIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditWorkpaperDelegate import AuditWorkpaperDelegate

		errMsg = "Failed to remove elements " + str(workpapersIds) + " for Workpapers on AuditEngagement"

		try:
			# get the AuditEngagement
			auditEngagement = self.get( auditEngagementId ).first()
				
			# split on a comma with no spaces
			idList = workpapersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AuditWorkpaper		
				auditWorkpaper = AuditWorkpaperDelegate().get(id).first();	
				# add the AuditWorkpaper
				auditEngagement.workpapers.remove(auditWorkpaper)
				
			# save it		
			auditEngagement.save()
			
			# reload and return the appropriate version
			return self.get( auditEngagementId );
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(auditEngagementId) + " does not exist.")
		except AuditWorkpaper.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditWorkpaper does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addFindings( self, auditEngagementId, findingsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditFindingDelegate import AuditFindingDelegate

		errMsg = "Failed to add elements " + str(findingsIds) + " for Findings on AuditEngagement"

		try:
			# get the AuditEngagement
			auditEngagement = self.get( auditEngagementId ).first()
				
			# split on a comma with no spaces
			idList = findingsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AuditFinding		
				auditFinding = AuditFindingDelegate().get(id).first();	
				# add the AuditFinding
				auditEngagement.findings.add(auditFinding)
				
			# save it		
			auditEngagement.save()
			
			# reload and return the appropriate version
			return self.get( auditEngagementId );
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(auditEngagementId) + " does not exist.")
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFindings( self, auditEngagementId, findingsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditFindingDelegate import AuditFindingDelegate

		errMsg = "Failed to remove elements " + str(findingsIds) + " for Findings on AuditEngagement"

		try:
			# get the AuditEngagement
			auditEngagement = self.get( auditEngagementId ).first()
				
			# split on a comma with no spaces
			idList = findingsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AuditFinding		
				auditFinding = AuditFindingDelegate().get(id).first();	
				# add the AuditFinding
				auditEngagement.findings.remove(auditFinding)
				
			# save it		
			auditEngagement.save()
			
			# reload and return the appropriate version
			return self.get( auditEngagementId );
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(auditEngagementId) + " does not exist.")
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		

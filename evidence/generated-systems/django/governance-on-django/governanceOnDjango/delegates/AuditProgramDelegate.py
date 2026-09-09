from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.AuditProgram import AuditProgram
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.AuditEngagement import AuditEngagement
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AuditProgram
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditProgramDelegate Declaration
#======================================================================
class AuditProgramDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, auditProgramId ):
		try:	
			auditProgram = AuditProgram.objects.filter(id=auditProgramId)
			return auditProgram.first();
		except AuditProgram.DoesNotExist:
			raise ProcessingError("AuditProgram with id " + str(auditProgramId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, auditProgram):
		for model in serializers.deserialize("json", auditProgram):
			model.save()
			return model;

	def create(self, auditProgram):
		auditProgram.save()
		return auditProgram;

	def saveFromJson(self, auditProgram):
		for model in serializers.deserialize("json", auditProgram):
			model.save()
			return auditProgram;
	
	def save(self, auditProgram):
		auditProgram.save()
		return auditProgram;
	
	def delete(self, auditProgramId ):
		errMsg = "Failed to delete AuditProgram from db using id " + str(auditProgramId)
		
		try:
			auditProgram = AuditProgram.objects.get(id=auditProgramId)
			auditProgram.delete()
			return True
		except AuditProgram.DoesNotExist:
			raise ProcessingError("AuditProgram with id " + str(auditProgramId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AuditProgram.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AuditProgram from db")
		except Exception:
			return None;
		
	def assignOrganization( self, auditProgramId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on AuditProgram"

		try:
			# get the AuditProgram from db
			auditProgram = self.get( auditProgramId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			auditProgram.organization = organization
			
			#save it
			auditProgram.save()

			# reload and return the appropriate version					
			return self.get( auditProgramId );
		except AuditProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditProgram with id " + str(auditProgramId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, auditProgramId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on AuditProgram"

		try:
			# get the AuditProgram from db
			auditProgram = self.get( auditProgramId ).first()	
			
			# assign to None for unassignment
			auditProgram.organization = None			

			#save it
			auditProgram.save()

			# reload and return the appropriate version					
			return self.get( auditProgramId );
		except AuditProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditProgram with id " + str(auditProgramId) + " does not exist.")
		except Exception:
			return None;
		
	def addEngagements( self, auditProgramId, engagementsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditEngagementDelegate import AuditEngagementDelegate

		errMsg = "Failed to add elements " + str(engagementsIds) + " for Engagements on AuditProgram"

		try:
			# get the AuditProgram
			auditProgram = self.get( auditProgramId ).first()
				
			# split on a comma with no spaces
			idList = engagementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AuditEngagement		
				auditEngagement = AuditEngagementDelegate().get(id).first();	
				# add the AuditEngagement
				auditProgram.engagements.add(auditEngagement)
				
			# save it		
			auditProgram.save()
			
			# reload and return the appropriate version
			return self.get( auditProgramId );
		except AuditProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditProgram with id " + str(auditProgramId) + " does not exist.")
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEngagements( self, auditProgramId, engagementsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditEngagementDelegate import AuditEngagementDelegate

		errMsg = "Failed to remove elements " + str(engagementsIds) + " for Engagements on AuditProgram"

		try:
			# get the AuditProgram
			auditProgram = self.get( auditProgramId ).first()
				
			# split on a comma with no spaces
			idList = engagementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AuditEngagement		
				auditEngagement = AuditEngagementDelegate().get(id).first();	
				# add the AuditEngagement
				auditProgram.engagements.remove(auditEngagement)
				
			# save it		
			auditProgram.save()
			
			# reload and return the appropriate version
			return self.get( auditProgramId );
		except AuditProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditProgram with id " + str(auditProgramId) + " does not exist.")
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		

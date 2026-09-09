from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.BusinessUnit import BusinessUnit
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.AuditEngagement import AuditEngagement
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BusinessUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BusinessUnitDelegate Declaration
#======================================================================
class BusinessUnitDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, businessUnitId ):
		try:	
			businessUnit = BusinessUnit.objects.filter(id=businessUnitId)
			return businessUnit.first();
		except BusinessUnit.DoesNotExist:
			raise ProcessingError("BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, businessUnit):
		for model in serializers.deserialize("json", businessUnit):
			model.save()
			return model;

	def create(self, businessUnit):
		businessUnit.save()
		return businessUnit;

	def saveFromJson(self, businessUnit):
		for model in serializers.deserialize("json", businessUnit):
			model.save()
			return businessUnit;
	
	def save(self, businessUnit):
		businessUnit.save()
		return businessUnit;
	
	def delete(self, businessUnitId ):
		errMsg = "Failed to delete BusinessUnit from db using id " + str(businessUnitId)
		
		try:
			businessUnit = BusinessUnit.objects.get(id=businessUnitId)
			businessUnit.delete()
			return True
		except BusinessUnit.DoesNotExist:
			raise ProcessingError("BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BusinessUnit.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BusinessUnit from db")
		except Exception:
			return None;
		
	def assignOrganization( self, businessUnitId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on BusinessUnit"

		try:
			# get the BusinessUnit from db
			businessUnit = self.get( businessUnitId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			businessUnit.organization = organization
			
			#save it
			businessUnit.save()

			# reload and return the appropriate version					
			return self.get( businessUnitId );
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, businessUnitId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on BusinessUnit"

		try:
			# get the BusinessUnit from db
			businessUnit = self.get( businessUnitId ).first()	
			
			# assign to None for unassignment
			businessUnit.organization = None			

			#save it
			businessUnit.save()

			# reload and return the appropriate version					
			return self.get( businessUnitId );
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except Exception:
			return None;
		
	def addAudits( self, businessUnitId, auditsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditEngagementDelegate import AuditEngagementDelegate

		errMsg = "Failed to add elements " + str(auditsIds) + " for Audits on BusinessUnit"

		try:
			# get the BusinessUnit
			businessUnit = self.get( businessUnitId ).first()
				
			# split on a comma with no spaces
			idList = auditsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AuditEngagement		
				auditEngagement = AuditEngagementDelegate().get(id).first();	
				# add the AuditEngagement
				businessUnit.audits.add(auditEngagement)
				
			# save it		
			businessUnit.save()
			
			# reload and return the appropriate version
			return self.get( businessUnitId );
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAudits( self, businessUnitId, auditsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditEngagementDelegate import AuditEngagementDelegate

		errMsg = "Failed to remove elements " + str(auditsIds) + " for Audits on BusinessUnit"

		try:
			# get the BusinessUnit
			businessUnit = self.get( businessUnitId ).first()
				
			# split on a comma with no spaces
			idList = auditsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AuditEngagement		
				auditEngagement = AuditEngagementDelegate().get(id).first();	
				# add the AuditEngagement
				businessUnit.audits.remove(auditEngagement)
				
			# save it		
			businessUnit.save()
			
			# reload and return the appropriate version
			return self.get( businessUnitId );
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		

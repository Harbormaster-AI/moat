from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.JobFamily import JobFamily
from hrOnDjango.models.Organization import Organization
from hrOnDjango.models.JobProfile import JobProfile
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model JobFamily
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobFamilyDelegate Declaration
#======================================================================
class JobFamilyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, jobFamilyId ):
		try:	
			jobFamily = JobFamily.objects.filter(id=jobFamilyId)
			return jobFamily.first();
		except JobFamily.DoesNotExist:
			raise ProcessingError("JobFamily with id " + str(jobFamilyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, jobFamily):
		for model in serializers.deserialize("json", jobFamily):
			model.save()
			return model;

	def create(self, jobFamily):
		jobFamily.save()
		return jobFamily;

	def saveFromJson(self, jobFamily):
		for model in serializers.deserialize("json", jobFamily):
			model.save()
			return jobFamily;
	
	def save(self, jobFamily):
		jobFamily.save()
		return jobFamily;
	
	def delete(self, jobFamilyId ):
		errMsg = "Failed to delete JobFamily from db using id " + str(jobFamilyId)
		
		try:
			jobFamily = JobFamily.objects.get(id=jobFamilyId)
			jobFamily.delete()
			return True
		except JobFamily.DoesNotExist:
			raise ProcessingError("JobFamily with id " + str(jobFamilyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = JobFamily.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all JobFamily from db")
		except Exception:
			return None;
		
	def assignOrganization( self, jobFamilyId, organizationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on JobFamily"

		try:
			# get the JobFamily from db
			jobFamily = self.get( jobFamilyId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			jobFamily.organization = organization
			
			#save it
			jobFamily.save()

			# reload and return the appropriate version					
			return self.get( jobFamilyId );
		except JobFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : JobFamily with id " + str(jobFamilyId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, jobFamilyId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on JobFamily"

		try:
			# get the JobFamily from db
			jobFamily = self.get( jobFamilyId ).first()	
			
			# assign to None for unassignment
			jobFamily.organization = None			

			#save it
			jobFamily.save()

			# reload and return the appropriate version					
			return self.get( jobFamilyId );
		except JobFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : JobFamily with id " + str(jobFamilyId) + " does not exist.")
		except Exception:
			return None;
		
	def addJobProfiles( self, jobFamilyId, jobProfilesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobProfileDelegate import JobProfileDelegate

		errMsg = "Failed to add elements " + str(jobProfilesIds) + " for JobProfiles on JobFamily"

		try:
			# get the JobFamily
			jobFamily = self.get( jobFamilyId ).first()
				
			# split on a comma with no spaces
			idList = jobProfilesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the JobProfile		
				jobProfile = JobProfileDelegate().get(id).first();	
				# add the JobProfile
				jobFamily.jobProfiles.add(jobProfile)
				
			# save it		
			jobFamily.save()
			
			# reload and return the appropriate version
			return self.get( jobFamilyId );
		except JobFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : JobFamily with id " + str(jobFamilyId) + " does not exist.")
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeJobProfiles( self, jobFamilyId, jobProfilesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobProfileDelegate import JobProfileDelegate

		errMsg = "Failed to remove elements " + str(jobProfilesIds) + " for JobProfiles on JobFamily"

		try:
			# get the JobFamily
			jobFamily = self.get( jobFamilyId ).first()
				
			# split on a comma with no spaces
			idList = jobProfilesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the JobProfile		
				jobProfile = JobProfileDelegate().get(id).first();	
				# add the JobProfile
				jobFamily.jobProfiles.remove(jobProfile)
				
			# save it		
			jobFamily.save()
			
			# reload and return the appropriate version
			return self.get( jobFamilyId );
		except JobFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : JobFamily with id " + str(jobFamilyId) + " does not exist.")
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		

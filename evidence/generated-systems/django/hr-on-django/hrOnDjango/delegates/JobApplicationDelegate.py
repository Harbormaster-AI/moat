from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.JobApplication import JobApplication
from hrOnDjango.models.Candidate import Candidate
from hrOnDjango.models.JobRequisition import JobRequisition
from hrOnDjango.models.Screening import Screening
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model JobApplication
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobApplicationDelegate Declaration
#======================================================================
class JobApplicationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, jobApplicationId ):
		try:	
			jobApplication = JobApplication.objects.filter(id=jobApplicationId)
			return jobApplication.first();
		except JobApplication.DoesNotExist:
			raise ProcessingError("JobApplication with id " + str(jobApplicationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, jobApplication):
		for model in serializers.deserialize("json", jobApplication):
			model.save()
			return model;

	def create(self, jobApplication):
		jobApplication.save()
		return jobApplication;

	def saveFromJson(self, jobApplication):
		for model in serializers.deserialize("json", jobApplication):
			model.save()
			return jobApplication;
	
	def save(self, jobApplication):
		jobApplication.save()
		return jobApplication;
	
	def delete(self, jobApplicationId ):
		errMsg = "Failed to delete JobApplication from db using id " + str(jobApplicationId)
		
		try:
			jobApplication = JobApplication.objects.get(id=jobApplicationId)
			jobApplication.delete()
			return True
		except JobApplication.DoesNotExist:
			raise ProcessingError("JobApplication with id " + str(jobApplicationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = JobApplication.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all JobApplication from db")
		except Exception:
			return None;
		
	def assignCandidate( self, jobApplicationId, candidateId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CandidateDelegate import CandidateDelegate

		errMsg = "Failed to assign element " + str(candidateId) + " for Candidate on JobApplication"

		try:
			# get the JobApplication from db
			jobApplication = self.get( jobApplicationId ).first()	
			
			# get the Candidate from db
			candidate = CandidateDelegate().get(candidateId).first();
			
			# assign the Candidate		
			jobApplication.candidate = candidate
			
			#save it
			jobApplication.save()

			# reload and return the appropriate version					
			return self.get( jobApplicationId );
		except JobApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : JobApplication with id " + str(jobApplicationId) + " does not exist.")
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCandidate( self, jobApplicationId ):
		errMsg = "Failed to unassign element " + str(candidateId) + " for Candidate on JobApplication"

		try:
			# get the JobApplication from db
			jobApplication = self.get( jobApplicationId ).first()	
			
			# assign to None for unassignment
			jobApplication.candidate = None			

			#save it
			jobApplication.save()

			# reload and return the appropriate version					
			return self.get( jobApplicationId );
		except JobApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : JobApplication with id " + str(jobApplicationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRequisition( self, jobApplicationId, requisitionId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobRequisitionDelegate import JobRequisitionDelegate

		errMsg = "Failed to assign element " + str(requisitionId) + " for Requisition on JobApplication"

		try:
			# get the JobApplication from db
			jobApplication = self.get( jobApplicationId ).first()	
			
			# get the JobRequisition from db
			jobRequisition = JobRequisitionDelegate().get(requisitionId).first();
			
			# assign the Requisition		
			jobApplication.requisition = jobRequisition
			
			#save it
			jobApplication.save()

			# reload and return the appropriate version					
			return self.get( jobApplicationId );
		except JobApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : JobApplication with id " + str(jobApplicationId) + " does not exist.")
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(requisitionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRequisition( self, jobApplicationId ):
		errMsg = "Failed to unassign element " + str(requisitionId) + " for Requisition on JobApplication"

		try:
			# get the JobApplication from db
			jobApplication = self.get( jobApplicationId ).first()	
			
			# assign to None for unassignment
			jobApplication.jobRequisition = None			

			#save it
			jobApplication.save()

			# reload and return the appropriate version					
			return self.get( jobApplicationId );
		except JobApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : JobApplication with id " + str(jobApplicationId) + " does not exist.")
		except Exception:
			return None;
		
	def addScreenings( self, jobApplicationId, screeningsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.ScreeningDelegate import ScreeningDelegate

		errMsg = "Failed to add elements " + str(screeningsIds) + " for Screenings on JobApplication"

		try:
			# get the JobApplication
			jobApplication = self.get( jobApplicationId ).first()
				
			# split on a comma with no spaces
			idList = screeningsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Screening		
				screening = ScreeningDelegate().get(id).first();	
				# add the Screening
				jobApplication.screenings.add(screening)
				
			# save it		
			jobApplication.save()
			
			# reload and return the appropriate version
			return self.get( jobApplicationId );
		except JobApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : JobApplication with id " + str(jobApplicationId) + " does not exist.")
		except Screening.DoesNotExist:
			raise ProcessingError(errMsg + " : Screening does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeScreenings( self, jobApplicationId, screeningsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.ScreeningDelegate import ScreeningDelegate

		errMsg = "Failed to remove elements " + str(screeningsIds) + " for Screenings on JobApplication"

		try:
			# get the JobApplication
			jobApplication = self.get( jobApplicationId ).first()
				
			# split on a comma with no spaces
			idList = screeningsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Screening		
				screening = ScreeningDelegate().get(id).first();	
				# add the Screening
				jobApplication.screenings.remove(screening)
				
			# save it		
			jobApplication.save()
			
			# reload and return the appropriate version
			return self.get( jobApplicationId );
		except JobApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : JobApplication with id " + str(jobApplicationId) + " does not exist.")
		except Screening.DoesNotExist:
			raise ProcessingError(errMsg + " : Screening does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		

from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Interview import Interview
from hrOnDjango.models.JobRequisition import JobRequisition
from hrOnDjango.models.Candidate import Candidate
from hrOnDjango.models.Employee import Employee
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Interview
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InterviewDelegate Declaration
#======================================================================
class InterviewDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, interviewId ):
		try:	
			interview = Interview.objects.filter(id=interviewId)
			return interview.first();
		except Interview.DoesNotExist:
			raise ProcessingError("Interview with id " + str(interviewId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, interview):
		for model in serializers.deserialize("json", interview):
			model.save()
			return model;

	def create(self, interview):
		interview.save()
		return interview;

	def saveFromJson(self, interview):
		for model in serializers.deserialize("json", interview):
			model.save()
			return interview;
	
	def save(self, interview):
		interview.save()
		return interview;
	
	def delete(self, interviewId ):
		errMsg = "Failed to delete Interview from db using id " + str(interviewId)
		
		try:
			interview = Interview.objects.get(id=interviewId)
			interview.delete()
			return True
		except Interview.DoesNotExist:
			raise ProcessingError("Interview with id " + str(interviewId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Interview.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Interview from db")
		except Exception:
			return None;
		
	def assignRequisition( self, interviewId, requisitionId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobRequisitionDelegate import JobRequisitionDelegate

		errMsg = "Failed to assign element " + str(requisitionId) + " for Requisition on Interview"

		try:
			# get the Interview from db
			interview = self.get( interviewId ).first()	
			
			# get the JobRequisition from db
			jobRequisition = JobRequisitionDelegate().get(requisitionId).first();
			
			# assign the Requisition		
			interview.requisition = jobRequisition
			
			#save it
			interview.save()

			# reload and return the appropriate version					
			return self.get( interviewId );
		except Interview.DoesNotExist:
			raise ProcessingError(errMsg + " : Interview with id " + str(interviewId) + " does not exist.")
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(requisitionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRequisition( self, interviewId ):
		errMsg = "Failed to unassign element " + str(requisitionId) + " for Requisition on Interview"

		try:
			# get the Interview from db
			interview = self.get( interviewId ).first()	
			
			# assign to None for unassignment
			interview.jobRequisition = None			

			#save it
			interview.save()

			# reload and return the appropriate version					
			return self.get( interviewId );
		except Interview.DoesNotExist:
			raise ProcessingError(errMsg + " : Interview with id " + str(interviewId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCandidate( self, interviewId, candidateId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CandidateDelegate import CandidateDelegate

		errMsg = "Failed to assign element " + str(candidateId) + " for Candidate on Interview"

		try:
			# get the Interview from db
			interview = self.get( interviewId ).first()	
			
			# get the Candidate from db
			candidate = CandidateDelegate().get(candidateId).first();
			
			# assign the Candidate		
			interview.candidate = candidate
			
			#save it
			interview.save()

			# reload and return the appropriate version					
			return self.get( interviewId );
		except Interview.DoesNotExist:
			raise ProcessingError(errMsg + " : Interview with id " + str(interviewId) + " does not exist.")
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCandidate( self, interviewId ):
		errMsg = "Failed to unassign element " + str(candidateId) + " for Candidate on Interview"

		try:
			# get the Interview from db
			interview = self.get( interviewId ).first()	
			
			# assign to None for unassignment
			interview.candidate = None			

			#save it
			interview.save()

			# reload and return the appropriate version					
			return self.get( interviewId );
		except Interview.DoesNotExist:
			raise ProcessingError(errMsg + " : Interview with id " + str(interviewId) + " does not exist.")
		except Exception:
			return None;
		
	def addInterviewers( self, interviewId, interviewersIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to add elements " + str(interviewersIds) + " for Interviewers on Interview"

		try:
			# get the Interview
			interview = self.get( interviewId ).first()
				
			# split on a comma with no spaces
			idList = interviewersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Employee		
				employee = EmployeeDelegate().get(id).first();	
				# add the Employee
				interview.interviewers.add(employee)
				
			# save it		
			interview.save()
			
			# reload and return the appropriate version
			return self.get( interviewId );
		except Interview.DoesNotExist:
			raise ProcessingError(errMsg + " : Interview with id " + str(interviewId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInterviewers( self, interviewId, interviewersIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to remove elements " + str(interviewersIds) + " for Interviewers on Interview"

		try:
			# get the Interview
			interview = self.get( interviewId ).first()
				
			# split on a comma with no spaces
			idList = interviewersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Employee		
				employee = EmployeeDelegate().get(id).first();	
				# add the Employee
				interview.interviewers.remove(employee)
				
			# save it		
			interview.save()
			
			# reload and return the appropriate version
			return self.get( interviewId );
		except Interview.DoesNotExist:
			raise ProcessingError(errMsg + " : Interview with id " + str(interviewId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		

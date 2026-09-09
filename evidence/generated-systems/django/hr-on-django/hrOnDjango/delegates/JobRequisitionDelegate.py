from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.JobRequisition import JobRequisition
from hrOnDjango.models.Department import Department
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.JobProfile import JobProfile
from hrOnDjango.models.Candidate import Candidate
from hrOnDjango.models.Interview import Interview
from hrOnDjango.models.Offer import Offer
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model JobRequisition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobRequisitionDelegate Declaration
#======================================================================
class JobRequisitionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, jobRequisitionId ):
		try:	
			jobRequisition = JobRequisition.objects.filter(id=jobRequisitionId)
			return jobRequisition.first();
		except JobRequisition.DoesNotExist:
			raise ProcessingError("JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, jobRequisition):
		for model in serializers.deserialize("json", jobRequisition):
			model.save()
			return model;

	def create(self, jobRequisition):
		jobRequisition.save()
		return jobRequisition;

	def saveFromJson(self, jobRequisition):
		for model in serializers.deserialize("json", jobRequisition):
			model.save()
			return jobRequisition;
	
	def save(self, jobRequisition):
		jobRequisition.save()
		return jobRequisition;
	
	def delete(self, jobRequisitionId ):
		errMsg = "Failed to delete JobRequisition from db using id " + str(jobRequisitionId)
		
		try:
			jobRequisition = JobRequisition.objects.get(id=jobRequisitionId)
			jobRequisition.delete()
			return True
		except JobRequisition.DoesNotExist:
			raise ProcessingError("JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = JobRequisition.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all JobRequisition from db")
		except Exception:
			return None;
		
	def assignDepartment( self, jobRequisitionId, departmentId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

		errMsg = "Failed to assign element " + str(departmentId) + " for Department on JobRequisition"

		try:
			# get the JobRequisition from db
			jobRequisition = self.get( jobRequisitionId ).first()	
			
			# get the Department from db
			department = DepartmentDelegate().get(departmentId).first();
			
			# assign the Department		
			jobRequisition.department = department
			
			#save it
			jobRequisition.save()

			# reload and return the appropriate version					
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDepartment( self, jobRequisitionId ):
		errMsg = "Failed to unassign element " + str(departmentId) + " for Department on JobRequisition"

		try:
			# get the JobRequisition from db
			jobRequisition = self.get( jobRequisitionId ).first()	
			
			# assign to None for unassignment
			jobRequisition.department = None			

			#save it
			jobRequisition.save()

			# reload and return the appropriate version					
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignHiringManager( self, jobRequisitionId, hiringManagerId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(hiringManagerId) + " for HiringManager on JobRequisition"

		try:
			# get the JobRequisition from db
			jobRequisition = self.get( jobRequisitionId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(hiringManagerId).first();
			
			# assign the HiringManager		
			jobRequisition.hiringManager = employee
			
			#save it
			jobRequisition.save()

			# reload and return the appropriate version					
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(hiringManagerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignHiringManager( self, jobRequisitionId ):
		errMsg = "Failed to unassign element " + str(hiringManagerId) + " for HiringManager on JobRequisition"

		try:
			# get the JobRequisition from db
			jobRequisition = self.get( jobRequisitionId ).first()	
			
			# assign to None for unassignment
			jobRequisition.employee = None			

			#save it
			jobRequisition.save()

			# reload and return the appropriate version					
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRecruiter( self, jobRequisitionId, recruiterId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(recruiterId) + " for Recruiter on JobRequisition"

		try:
			# get the JobRequisition from db
			jobRequisition = self.get( jobRequisitionId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(recruiterId).first();
			
			# assign the Recruiter		
			jobRequisition.recruiter = employee
			
			#save it
			jobRequisition.save()

			# reload and return the appropriate version					
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(recruiterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRecruiter( self, jobRequisitionId ):
		errMsg = "Failed to unassign element " + str(recruiterId) + " for Recruiter on JobRequisition"

		try:
			# get the JobRequisition from db
			jobRequisition = self.get( jobRequisitionId ).first()	
			
			# assign to None for unassignment
			jobRequisition.employee = None			

			#save it
			jobRequisition.save()

			# reload and return the appropriate version					
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignJobProfile( self, jobRequisitionId, jobProfileId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobProfileDelegate import JobProfileDelegate

		errMsg = "Failed to assign element " + str(jobProfileId) + " for JobProfile on JobRequisition"

		try:
			# get the JobRequisition from db
			jobRequisition = self.get( jobRequisitionId ).first()	
			
			# get the JobProfile from db
			jobProfile = JobProfileDelegate().get(jobProfileId).first();
			
			# assign the JobProfile		
			jobRequisition.jobProfile = jobProfile
			
			#save it
			jobRequisition.save()

			# reload and return the appropriate version					
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile with id " + str(jobProfileId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignJobProfile( self, jobRequisitionId ):
		errMsg = "Failed to unassign element " + str(jobProfileId) + " for JobProfile on JobRequisition"

		try:
			# get the JobRequisition from db
			jobRequisition = self.get( jobRequisitionId ).first()	
			
			# assign to None for unassignment
			jobRequisition.jobProfile = None			

			#save it
			jobRequisition.save()

			# reload and return the appropriate version					
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Exception:
			return None;
		
	def addCandidates( self, jobRequisitionId, candidatesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CandidateDelegate import CandidateDelegate

		errMsg = "Failed to add elements " + str(candidatesIds) + " for Candidates on JobRequisition"

		try:
			# get the JobRequisition
			jobRequisition = self.get( jobRequisitionId ).first()
				
			# split on a comma with no spaces
			idList = candidatesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Candidate		
				candidate = CandidateDelegate().get(id).first();	
				# add the Candidate
				jobRequisition.candidates.add(candidate)
				
			# save it		
			jobRequisition.save()
			
			# reload and return the appropriate version
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCandidates( self, jobRequisitionId, candidatesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CandidateDelegate import CandidateDelegate

		errMsg = "Failed to remove elements " + str(candidatesIds) + " for Candidates on JobRequisition"

		try:
			# get the JobRequisition
			jobRequisition = self.get( jobRequisitionId ).first()
				
			# split on a comma with no spaces
			idList = candidatesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Candidate		
				candidate = CandidateDelegate().get(id).first();	
				# add the Candidate
				jobRequisition.candidates.remove(candidate)
				
			# save it		
			jobRequisition.save()
			
			# reload and return the appropriate version
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInterviews( self, jobRequisitionId, interviewsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.InterviewDelegate import InterviewDelegate

		errMsg = "Failed to add elements " + str(interviewsIds) + " for Interviews on JobRequisition"

		try:
			# get the JobRequisition
			jobRequisition = self.get( jobRequisitionId ).first()
				
			# split on a comma with no spaces
			idList = interviewsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Interview		
				interview = InterviewDelegate().get(id).first();	
				# add the Interview
				jobRequisition.interviews.add(interview)
				
			# save it		
			jobRequisition.save()
			
			# reload and return the appropriate version
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Interview.DoesNotExist:
			raise ProcessingError(errMsg + " : Interview does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInterviews( self, jobRequisitionId, interviewsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.InterviewDelegate import InterviewDelegate

		errMsg = "Failed to remove elements " + str(interviewsIds) + " for Interviews on JobRequisition"

		try:
			# get the JobRequisition
			jobRequisition = self.get( jobRequisitionId ).first()
				
			# split on a comma with no spaces
			idList = interviewsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Interview		
				interview = InterviewDelegate().get(id).first();	
				# add the Interview
				jobRequisition.interviews.remove(interview)
				
			# save it		
			jobRequisition.save()
			
			# reload and return the appropriate version
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Interview.DoesNotExist:
			raise ProcessingError(errMsg + " : Interview does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOffers( self, jobRequisitionId, offersIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OfferDelegate import OfferDelegate

		errMsg = "Failed to add elements " + str(offersIds) + " for Offers on JobRequisition"

		try:
			# get the JobRequisition
			jobRequisition = self.get( jobRequisitionId ).first()
				
			# split on a comma with no spaces
			idList = offersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Offer		
				offer = OfferDelegate().get(id).first();	
				# add the Offer
				jobRequisition.offers.add(offer)
				
			# save it		
			jobRequisition.save()
			
			# reload and return the appropriate version
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOffers( self, jobRequisitionId, offersIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OfferDelegate import OfferDelegate

		errMsg = "Failed to remove elements " + str(offersIds) + " for Offers on JobRequisition"

		try:
			# get the JobRequisition
			jobRequisition = self.get( jobRequisitionId ).first()
				
			# split on a comma with no spaces
			idList = offersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Offer		
				offer = OfferDelegate().get(id).first();	
				# add the Offer
				jobRequisition.offers.remove(offer)
				
			# save it		
			jobRequisition.save()
			
			# reload and return the appropriate version
			return self.get( jobRequisitionId );
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(jobRequisitionId) + " does not exist.")
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		

from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Report import Report
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.Visualization import Visualization
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.SemanticModel import SemanticModel
from analyticsOnDjango.models.BIQuery import BIQuery
from analyticsOnDjango.models.Tag import Tag
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Report
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReportDelegate Declaration
#======================================================================
class ReportDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, reportId ):
		try:	
			report = Report.objects.filter(id=reportId)
			return report.first();
		except Report.DoesNotExist:
			raise ProcessingError("Report with id " + str(reportId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, report):
		for model in serializers.deserialize("json", report):
			model.save()
			return model;

	def create(self, report):
		report.save()
		return report;

	def saveFromJson(self, report):
		for model in serializers.deserialize("json", report):
			model.save()
			return report;
	
	def save(self, report):
		report.save()
		return report;
	
	def delete(self, reportId ):
		errMsg = "Failed to delete Report from db using id " + str(reportId)
		
		try:
			report = Report.objects.get(id=reportId)
			report.delete()
			return True
		except Report.DoesNotExist:
			raise ProcessingError("Report with id " + str(reportId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Report.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Report from db")
		except Exception:
			return None;
		
	def assignWorkspace( self, reportId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on Report"

		try:
			# get the Report from db
			report = self.get( reportId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			report.workspace = analyticsWorkspace
			
			#save it
			report.save()

			# reload and return the appropriate version					
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, reportId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on Report"

		try:
			# get the Report from db
			report = self.get( reportId ).first()	
			
			# assign to None for unassignment
			report.analyticsWorkspace = None			

			#save it
			report.save()

			# reload and return the appropriate version					
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except Exception:
			return None;
		
	def addVisualizations( self, reportId, visualizationsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.VisualizationDelegate import VisualizationDelegate

		errMsg = "Failed to add elements " + str(visualizationsIds) + " for Visualizations on Report"

		try:
			# get the Report
			report = self.get( reportId ).first()
				
			# split on a comma with no spaces
			idList = visualizationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Visualization		
				visualization = VisualizationDelegate().get(id).first();	
				# add the Visualization
				report.visualizations.add(visualization)
				
			# save it		
			report.save()
			
			# reload and return the appropriate version
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVisualizations( self, reportId, visualizationsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.VisualizationDelegate import VisualizationDelegate

		errMsg = "Failed to remove elements " + str(visualizationsIds) + " for Visualizations on Report"

		try:
			# get the Report
			report = self.get( reportId ).first()
				
			# split on a comma with no spaces
			idList = visualizationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Visualization		
				visualization = VisualizationDelegate().get(id).first();	
				# add the Visualization
				report.visualizations.remove(visualization)
				
			# save it		
			report.save()
			
			# reload and return the appropriate version
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDatasets( self, reportId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on Report"

		try:
			# get the Report
			report = self.get( reportId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				report.datasets.add(dataSet)
				
			# save it		
			report.save()
			
			# reload and return the appropriate version
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, reportId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on Report"

		try:
			# get the Report
			report = self.get( reportId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				report.datasets.remove(dataSet)
				
			# save it		
			report.save()
			
			# reload and return the appropriate version
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSemanticModels( self, reportId, semanticModelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.SemanticModelDelegate import SemanticModelDelegate

		errMsg = "Failed to add elements " + str(semanticModelsIds) + " for SemanticModels on Report"

		try:
			# get the Report
			report = self.get( reportId ).first()
				
			# split on a comma with no spaces
			idList = semanticModelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SemanticModel		
				semanticModel = SemanticModelDelegate().get(id).first();	
				# add the SemanticModel
				report.semanticModels.add(semanticModel)
				
			# save it		
			report.save()
			
			# reload and return the appropriate version
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSemanticModels( self, reportId, semanticModelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.SemanticModelDelegate import SemanticModelDelegate

		errMsg = "Failed to remove elements " + str(semanticModelsIds) + " for SemanticModels on Report"

		try:
			# get the Report
			report = self.get( reportId ).first()
				
			# split on a comma with no spaces
			idList = semanticModelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SemanticModel		
				semanticModel = SemanticModelDelegate().get(id).first();	
				# add the SemanticModel
				report.semanticModels.remove(semanticModel)
				
			# save it		
			report.save()
			
			# reload and return the appropriate version
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addQueries( self, reportId, queriesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BIQueryDelegate import BIQueryDelegate

		errMsg = "Failed to add elements " + str(queriesIds) + " for Queries on Report"

		try:
			# get the Report
			report = self.get( reportId ).first()
				
			# split on a comma with no spaces
			idList = queriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BIQuery		
				bIQuery = BIQueryDelegate().get(id).first();	
				# add the BIQuery
				report.queries.add(bIQuery)
				
			# save it		
			report.save()
			
			# reload and return the appropriate version
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeQueries( self, reportId, queriesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BIQueryDelegate import BIQueryDelegate

		errMsg = "Failed to remove elements " + str(queriesIds) + " for Queries on Report"

		try:
			# get the Report
			report = self.get( reportId ).first()
				
			# split on a comma with no spaces
			idList = queriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BIQuery		
				bIQuery = BIQueryDelegate().get(id).first();	
				# add the BIQuery
				report.queries.remove(bIQuery)
				
			# save it		
			report.save()
			
			# reload and return the appropriate version
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTags( self, reportId, tagsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TagDelegate import TagDelegate

		errMsg = "Failed to add elements " + str(tagsIds) + " for Tags on Report"

		try:
			# get the Report
			report = self.get( reportId ).first()
				
			# split on a comma with no spaces
			idList = tagsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Tag		
				tag = TagDelegate().get(id).first();	
				# add the Tag
				report.tags.add(tag)
				
			# save it		
			report.save()
			
			# reload and return the appropriate version
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTags( self, reportId, tagsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TagDelegate import TagDelegate

		errMsg = "Failed to remove elements " + str(tagsIds) + " for Tags on Report"

		try:
			# get the Report
			report = self.get( reportId ).first()
				
			# split on a comma with no spaces
			idList = tagsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Tag		
				tag = TagDelegate().get(id).first();	
				# add the Tag
				report.tags.remove(tag)
				
			# save it		
			report.save()
			
			# reload and return the appropriate version
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		

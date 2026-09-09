from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.LineageNode import LineageNode
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Model import Model
from analyticsOnDjango.models.DataPipeline import DataPipeline
from analyticsOnDjango.models.Dashboard import Dashboard
from analyticsOnDjango.models.Report import Report
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model LineageNode
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LineageNodeDelegate Declaration
#======================================================================
class LineageNodeDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, lineageNodeId ):
		try:	
			lineageNode = LineageNode.objects.filter(id=lineageNodeId)
			return lineageNode.first();
		except LineageNode.DoesNotExist:
			raise ProcessingError("LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, lineageNode):
		for model in serializers.deserialize("json", lineageNode):
			model.save()
			return model;

	def create(self, lineageNode):
		lineageNode.save()
		return lineageNode;

	def saveFromJson(self, lineageNode):
		for model in serializers.deserialize("json", lineageNode):
			model.save()
			return lineageNode;
	
	def save(self, lineageNode):
		lineageNode.save()
		return lineageNode;
	
	def delete(self, lineageNodeId ):
		errMsg = "Failed to delete LineageNode from db using id " + str(lineageNodeId)
		
		try:
			lineageNode = LineageNode.objects.get(id=lineageNodeId)
			lineageNode.delete()
			return True
		except LineageNode.DoesNotExist:
			raise ProcessingError("LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = LineageNode.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all LineageNode from db")
		except Exception:
			return None;
		
	def assignWorkspace( self, lineageNodeId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on LineageNode"

		try:
			# get the LineageNode from db
			lineageNode = self.get( lineageNodeId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			lineageNode.workspace = analyticsWorkspace
			
			#save it
			lineageNode.save()

			# reload and return the appropriate version					
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, lineageNodeId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on LineageNode"

		try:
			# get the LineageNode from db
			lineageNode = self.get( lineageNodeId ).first()	
			
			# assign to None for unassignment
			lineageNode.analyticsWorkspace = None			

			#save it
			lineageNode.save()

			# reload and return the appropriate version					
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except Exception:
			return None;
		
	def addInputs( self, lineageNodeId, inputsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.LineageNodeDelegate import LineageNodeDelegate

		errMsg = "Failed to add elements " + str(inputsIds) + " for Inputs on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = inputsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LineageNode		
				lineageNode = LineageNodeDelegate().get(id).first();	
				# add the LineageNode
				lineageNode.inputs.add(lineageNode)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInputs( self, lineageNodeId, inputsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.LineageNodeDelegate import LineageNodeDelegate

		errMsg = "Failed to remove elements " + str(inputsIds) + " for Inputs on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = inputsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LineageNode		
				lineageNode = LineageNodeDelegate().get(id).first();	
				# add the LineageNode
				lineageNode.inputs.remove(lineageNode)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOutputs( self, lineageNodeId, outputsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.LineageNodeDelegate import LineageNodeDelegate

		errMsg = "Failed to add elements " + str(outputsIds) + " for Outputs on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = outputsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LineageNode		
				lineageNode = LineageNodeDelegate().get(id).first();	
				# add the LineageNode
				lineageNode.outputs.add(lineageNode)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOutputs( self, lineageNodeId, outputsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.LineageNodeDelegate import LineageNodeDelegate

		errMsg = "Failed to remove elements " + str(outputsIds) + " for Outputs on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = outputsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LineageNode		
				lineageNode = LineageNodeDelegate().get(id).first();	
				# add the LineageNode
				lineageNode.outputs.remove(lineageNode)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDatasets( self, lineageNodeId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				lineageNode.datasets.add(dataSet)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, lineageNodeId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				lineageNode.datasets.remove(dataSet)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addModels( self, lineageNodeId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to add elements " + str(modelsIds) + " for Models on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				lineageNode.models.add(model)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeModels( self, lineageNodeId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to remove elements " + str(modelsIds) + " for Models on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				lineageNode.models.remove(model)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPipelines( self, lineageNodeId, pipelinesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataPipelineDelegate import DataPipelineDelegate

		errMsg = "Failed to add elements " + str(pipelinesIds) + " for Pipelines on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = pipelinesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataPipeline		
				dataPipeline = DataPipelineDelegate().get(id).first();	
				# add the DataPipeline
				lineageNode.pipelines.add(dataPipeline)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePipelines( self, lineageNodeId, pipelinesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataPipelineDelegate import DataPipelineDelegate

		errMsg = "Failed to remove elements " + str(pipelinesIds) + " for Pipelines on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = pipelinesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataPipeline		
				dataPipeline = DataPipelineDelegate().get(id).first();	
				# add the DataPipeline
				lineageNode.pipelines.remove(dataPipeline)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDashboards( self, lineageNodeId, dashboardsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

		errMsg = "Failed to add elements " + str(dashboardsIds) + " for Dashboards on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = dashboardsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dashboard		
				dashboard = DashboardDelegate().get(id).first();	
				# add the Dashboard
				lineageNode.dashboards.add(dashboard)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDashboards( self, lineageNodeId, dashboardsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

		errMsg = "Failed to remove elements " + str(dashboardsIds) + " for Dashboards on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = dashboardsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dashboard		
				dashboard = DashboardDelegate().get(id).first();	
				# add the Dashboard
				lineageNode.dashboards.remove(dashboard)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReports( self, lineageNodeId, reportsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to add elements " + str(reportsIds) + " for Reports on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				lineageNode.reports.add(report)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReports( self, lineageNodeId, reportsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to remove elements " + str(reportsIds) + " for Reports on LineageNode"

		try:
			# get the LineageNode
			lineageNode = self.get( lineageNodeId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				lineageNode.reports.remove(report)
				
			# save it		
			lineageNode.save()
			
			# reload and return the appropriate version
			return self.get( lineageNodeId );
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		

class CoverageDefinitionsController < ApplicationController
  def index
    @coverageDefinitions = CoverageDefinition.all
  end
 
  def show
    @coverageDefinition = CoverageDefinition.find(params[:id])
  end
 
  def new
    @coverageDefinition = CoverageDefinition.new
  end
 
  def edit
    @coverageDefinition = CoverageDefinition.find(params[:id])
  end
 
  def create
    @coverageDefinition = CoverageDefinition.new(coverageDefinition_params)
 
    if @coverageDefinition.save
      redirect_to coverageDefinitions_path
    else
      render 'new'
    end
  end
 
  def update
    @coverageDefinition = CoverageDefinition.find(params[:id])
 
    if @coverageDefinition.update(coverageDefinition_params)
      redirect_to coverageDefinitions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @coverageDefinition = CoverageDefinition.find(params[:id])
    @coverageDefinition.destroy
    redirect_to coverageDefinitions_path
  end

 
  private
    def coverageDefinition_params
      params.require(:coverageDefinition).permit(:name, :defaultLimit, :defaultDeductible, :asMandatory, :CoverageType)
    end
end
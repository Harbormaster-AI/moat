class ComplianceRequirementsController < ApplicationController
  def index
    @complianceRequirements = ComplianceRequirement.all
  end
 
  def show
    @complianceRequirement = ComplianceRequirement.find(params[:id])
  end
 
  def new
    @complianceRequirement = ComplianceRequirement.new
  end
 
  def edit
    @complianceRequirement = ComplianceRequirement.find(params[:id])
  end
 
  def create
    @complianceRequirement = ComplianceRequirement.new(complianceRequirement_params)
 
    if @complianceRequirement.save
      redirect_to complianceRequirements_path
    else
      render 'new'
    end
  end
 
  def update
    @complianceRequirement = ComplianceRequirement.find(params[:id])
 
    if @complianceRequirement.update(complianceRequirement_params)
      redirect_to complianceRequirements_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @complianceRequirement = ComplianceRequirement.find(params[:id])
    @complianceRequirement.destroy
    redirect_to complianceRequirements_path
  end

 
  private
    def complianceRequirement_params
      params.require(:complianceRequirement).permit(:name, :source, :citation, :Applicability, :Status)
    end
end
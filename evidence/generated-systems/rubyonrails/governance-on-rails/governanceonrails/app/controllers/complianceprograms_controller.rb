class ComplianceProgramsController < ApplicationController
  def index
    @compliancePrograms = ComplianceProgram.all
  end
 
  def show
    @complianceProgram = ComplianceProgram.find(params[:id])
  end
 
  def new
    @complianceProgram = ComplianceProgram.new
  end
 
  def edit
    @complianceProgram = ComplianceProgram.find(params[:id])
  end
 
  def create
    @complianceProgram = ComplianceProgram.new(complianceProgram_params)
 
    if @complianceProgram.save
      redirect_to compliancePrograms_path
    else
      render 'new'
    end
  end
 
  def update
    @complianceProgram = ComplianceProgram.find(params[:id])
 
    if @complianceProgram.update(complianceProgram_params)
      redirect_to compliancePrograms_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @complianceProgram = ComplianceProgram.find(params[:id])
    @complianceProgram.destroy
    redirect_to compliancePrograms_path
  end

 
  private
    def complianceProgram_params
      params.require(:complianceProgram).permit(:name, :framework, :Status)
    end
end
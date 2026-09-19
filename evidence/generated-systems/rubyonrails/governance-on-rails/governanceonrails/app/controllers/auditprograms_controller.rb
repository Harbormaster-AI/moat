class AuditProgramsController < ApplicationController
  def index
    @auditPrograms = AuditProgram.all
  end
 
  def show
    @auditProgram = AuditProgram.find(params[:id])
  end
 
  def new
    @auditProgram = AuditProgram.new
  end
 
  def edit
    @auditProgram = AuditProgram.find(params[:id])
  end
 
  def create
    @auditProgram = AuditProgram.new(auditProgram_params)
 
    if @auditProgram.save
      redirect_to auditPrograms_path
    else
      render 'new'
    end
  end
 
  def update
    @auditProgram = AuditProgram.find(params[:id])
 
    if @auditProgram.update(auditProgram_params)
      redirect_to auditPrograms_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @auditProgram = AuditProgram.find(params[:id])
    @auditProgram.destroy
    redirect_to auditPrograms_path
  end

 
  private
    def auditProgram_params
      params.require(:auditProgram).permit(:name, :scope, :Cycle, :Status)
    end
end
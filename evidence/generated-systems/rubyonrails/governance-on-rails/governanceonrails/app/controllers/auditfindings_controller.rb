class AuditFindingsController < ApplicationController
  def index
    @auditFindings = AuditFinding.all
  end
 
  def show
    @auditFinding = AuditFinding.find(params[:id])
  end
 
  def new
    @auditFinding = AuditFinding.new
  end
 
  def edit
    @auditFinding = AuditFinding.find(params[:id])
  end
 
  def create
    @auditFinding = AuditFinding.new(auditFinding_params)
 
    if @auditFinding.save
      redirect_to auditFindings_path
    else
      render 'new'
    end
  end
 
  def update
    @auditFinding = AuditFinding.find(params[:id])
 
    if @auditFinding.update(auditFinding_params)
      redirect_to auditFindings_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @auditFinding = AuditFinding.find(params[:id])
    @auditFinding.destroy
    redirect_to auditFindings_path
  end

 
  private
    def auditFinding_params
      params.require(:auditFinding).permit(:title, :description, :dueDate, :Severity, :Status)
    end
end
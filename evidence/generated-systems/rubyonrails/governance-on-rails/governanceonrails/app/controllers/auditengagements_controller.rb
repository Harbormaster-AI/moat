class AuditEngagementsController < ApplicationController
  def index
    @auditEngagements = AuditEngagement.all
  end
 
  def show
    @auditEngagement = AuditEngagement.find(params[:id])
  end
 
  def new
    @auditEngagement = AuditEngagement.new
  end
 
  def edit
    @auditEngagement = AuditEngagement.find(params[:id])
  end
 
  def create
    @auditEngagement = AuditEngagement.new(auditEngagement_params)
 
    if @auditEngagement.save
      redirect_to auditEngagements_path
    else
      render 'new'
    end
  end
 
  def update
    @auditEngagement = AuditEngagement.find(params[:id])
 
    if @auditEngagement.update(auditEngagement_params)
      redirect_to auditEngagements_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @auditEngagement = AuditEngagement.find(params[:id])
    @auditEngagement.destroy
    redirect_to auditEngagements_path
  end

 
  private
    def auditEngagement_params
      params.require(:auditEngagement).permit(:title, :startDate, :endDate, :Status)
    end
end
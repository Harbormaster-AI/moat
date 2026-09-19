class ApprovalsController < ApplicationController
  def index
    @approvals = Approval.all
  end
 
  def show
    @approval = Approval.find(params[:id])
  end
 
  def new
    @approval = Approval.new
  end
 
  def edit
    @approval = Approval.find(params[:id])
  end
 
  def create
    @approval = Approval.new(approval_params)
 
    if @approval.save
      redirect_to approvals_path
    else
      render 'new'
    end
  end
 
  def update
    @approval = Approval.find(params[:id])
 
    if @approval.update(approval_params)
      redirect_to approvals_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @approval = Approval.find(params[:id])
    @approval.destroy
    redirect_to approvals_path
  end

 
  private
    def approval_params
      params.require(:approval).permit(:approverComment, :actionDate, :Status)
    end
end
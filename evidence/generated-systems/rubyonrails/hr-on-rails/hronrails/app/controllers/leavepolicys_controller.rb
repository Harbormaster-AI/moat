class LeavePolicysController < ApplicationController
  def index
    @leavePolicys = LeavePolicy.all
  end
 
  def show
    @leavePolicy = LeavePolicy.find(params[:id])
  end
 
  def new
    @leavePolicy = LeavePolicy.new
  end
 
  def edit
    @leavePolicy = LeavePolicy.find(params[:id])
  end
 
  def create
    @leavePolicy = LeavePolicy.new(leavePolicy_params)
 
    if @leavePolicy.save
      redirect_to leavePolicys_path
    else
      render 'new'
    end
  end
 
  def update
    @leavePolicy = LeavePolicy.find(params[:id])
 
    if @leavePolicy.update(leavePolicy_params)
      redirect_to leavePolicys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @leavePolicy = LeavePolicy.find(params[:id])
    @leavePolicy.destroy
    redirect_to leavePolicys_path
  end

 
  private
    def leavePolicy_params
      params.require(:leavePolicy).permit(:name, :accrualRate, :carryoverAllowed, :maxBalance, :LeaveCategory, :AccrualUnit)
    end
end